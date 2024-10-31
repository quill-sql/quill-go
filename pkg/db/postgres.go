package db

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"

	_ "github.com/lib/pq"
	"github.com/quill-sql/quill-go/pkg/assets"
	"github.com/quill-sql/quill-go/pkg/utils"
)

type PostgresConnectionConfig struct {
	ConnectionString string
	SSL              *struct {
		RejectUnauthorized bool
		CA                 *string
		Key                *string
		Cert               *string
	}
}

// ConnectToPostgres establishes a connection to a PostgreSQL database.
func ConnectToPostgres(config PostgresConnectionConfig) (*sql.DB, error) {
	connStr := config.ConnectionString
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return db, nil
}

// DisconnectFromPostgres closes the connection to the PostgreSQL database.
func DisconnectFromPostgres(db *sql.DB) error {
	if err := db.Close(); err != nil {
		log.Fatalf("Failed to close connection: %v", err)
		return err
	}
	return nil
}

// RunQueryPostgres executes a SQL query on the PostgreSQL database.
func RunQueryPostgres(db *sql.DB, query string) (*QueryResults, error) {
	rows, err := db.Query(query)
	if err != nil {
		cleanErrMsg := strings.TrimPrefix(err.Error(), "pq: ")
		return nil, errors.New(cleanErrMsg)
	}
	defer rows.Close()

	cols, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	colTypes, err := rows.ColumnTypes()
	if err != nil {
		return nil, err
	}
	results := &QueryResults{}
	results.Fields = make([]Field, len(cols))
	results.Rows = make([]map[string]interface{}, 0)
	for i, col := range cols {
		dataTypeID := 0
		if sqlType := colTypes[i].DatabaseTypeName(); sqlType != "" {
			dataTypeID = getDataTypeID(sqlType)
		}
		results.Fields[i] = Field{Name: col, DataTypeID: dataTypeID}
	}
	for rows.Next() {
		rowData := make([]interface{}, len(cols))
		rowPointers := make([]interface{}, len(cols))

		for i := range rowData {
			rowPointers[i] = &rowData[i]
		}

		if err := rows.Scan(rowPointers...); err != nil {
			return nil, err
		}

		rowMap := make(map[string]interface{})
		for i, col := range cols {
			value := rowData[i]
			switch v := value.(type) {
			case []byte:
				// Try to convert []byte to string for numeric types
				str := string(v)
				if f, err := strconv.ParseFloat(str, 64); err == nil {
					rowMap[col] = f
				} else {
					rowMap[col] = str
				}
			default:
				rowMap[col] = v
			}
		}

		results.Rows = append(results.Rows, rowMap)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return results, nil
}

func getDataTypeID(sqlType string) int {
	// This is a simplified mapping. You may need to expand this based on your needs.
	var pgType int = 1043

	for _, pgTypeStruct := range assets.PGTypes {
		if pgTypeStruct.Typname == strings.ToLower(sqlType) {
			pgType = pgTypeStruct.OID
			break
		}
	}

	return pgType
}

// GetSchemasPostgres retrieves the list of schemas from the PostgreSQL database.
func GetSchemasPostgres(db *sql.DB) ([]string, error) {
	sql := `SELECT schema_name FROM information_schema.schemata 
    WHERE schema_name NOT LIKE 'pg_%' AND schema_name != 'information_schema';`
	results, err := RunQueryPostgres(db, sql)
	if err != nil {
		return nil, err
	}

	schemas := make([]string, len(results.Rows))
	for i, row := range results.Rows {
		schemas[i] = row["schema_name"].(string)
	}

	return schemas, nil
}

// GetTablesBySchemaPostgres retrieves the tables for the given schemas from the PostgreSQL database.
func GetTablesBySchemaPostgres(db *sql.DB, schemaNames []string) ([]Table, error) {
	var allTables []Table
	for _, schema := range schemaNames {
		sql := fmt.Sprintf(`SELECT table_name, table_schema FROM information_schema.tables WHERE table_schema = '%s'`, schema)
		results, err := RunQueryPostgres(db, sql)
		if err != nil {
			return nil, err
		}

		for _, row := range results.Rows {
			allTables = append(allTables, Table{
				TableName:  row["table_name"].(string),
				SchemaName: row["table_schema"].(string),
			})
		}
	}

	return allTables, nil
}

// GetColumnsByTablePostgres retrieves the columns for a given table from the PostgreSQL database.
func GetColumnsByTablePostgres(db *sql.DB, schemaName, tableName string) ([]string, error) {
	sql := fmt.Sprintf(`SELECT column_name FROM information_schema.columns WHERE table_schema = '%s' and table_name = '%s'`, schemaName, tableName)
	results, err := RunQueryPostgres(db, sql)
	if err != nil {
		return nil, err
	}

	columns := make([]string, len(results.Rows))
	for i, row := range results.Rows {
		columns[i] = row["column_name"].(string)
	}

	return columns, nil
}

// GetForeignKeysPostgres retrieves the foreign keys for a given table from the PostgreSQL database.
func GetForeignKeysPostgres(db *sql.DB, schemaName, tableName, primaryKey string) ([]string, error) {
	depluralizedTableName := utils.Depluralize(tableName)
	sql := fmt.Sprintf(`SELECT column_name FROM information_schema.columns 
  WHERE table_schema = '%s' 
  and table_name != '%s' 
  and (column_name = '%s' 
    or column_name = '%s_%s' 
    or column_name = '%s%s')`, schemaName, tableName, primaryKey, depluralizedTableName, primaryKey, depluralizedTableName, utils.Capitalize(primaryKey))
	results, err := RunQueryPostgres(db, sql)
	if err != nil {
		return nil, err
	}

	foreignKeysString := make([]string, len(results.Rows))
	for i, row := range results.Rows {
		foreignKeysString[i] = row["column_name"].(string)
	}

	filteredKeys := filterPostgresForeignKeys(foreignKeysString)

	if len(filteredKeys) == 0 {
		sql = fmt.Sprintf(`SELECT column_name FROM information_schema.columns 
      WHERE table_schema = '%s' 
      and table_name != '%s' 
      and (column_name like '%s%%' 
      or column_name like '%%\\_id' 
      or column_name like '%%Id' 
      or column_name like '%%\\_%s' 
      or column_name like '%%%s')`, schemaName, tableName, tableName, primaryKey, utils.Capitalize(primaryKey))
		results, err := RunQueryPostgres(db, sql)
		if err != nil {
			return nil, err
		}

		for i, row := range results.Rows {
			foreignKeysString[i] = row["column_name"].(string)
		}

		filteredKeys = filterPostgresForeignKeys(foreignKeysString)
	}

	return filteredKeys, nil
}

// GetSchemaColumnInfoPostgres retrieves column information for a given schema from the PostgreSQL database.
func GetSchemaColumnInfoPostgres(db *sql.DB /*schemaName string,*/, tableNames []Table) ([]map[string]interface{}, error) {
	var allColumns []map[string]interface{}
	for _, tableName := range tableNames {
		query := fmt.Sprintf(`
      SELECT column_name as "columnName", udt_name as "fieldType"
      FROM information_schema.columns
      WHERE table_schema = '%s' 
      AND table_name = '%s'
      ORDER BY ordinal_position;`, tableName.SchemaName, tableName.TableName)
		results, err := RunQueryPostgres(db, query)
		if err != nil {
			return nil, err
		}

		columns := make([]map[string]interface{}, len(results.Rows))
		for i, row := range results.Rows {
			var pgType int = 1043

			for _, pgTypeStruct := range assets.PGTypes {
				if pgTypeStruct.Typname == row["fieldType"].(string) {
					pgType = pgTypeStruct.OID
					break
				}
			}

			columns[i] = map[string]interface{}{
				"columnName":  row["columnName"],
				"dataTypeID":  pgType,
				"fieldType":   row["fieldType"],
				"displayName": row["columnName"],
			}
		}

		allColumns = append(allColumns, map[string]interface{}{
			"tableName":   fmt.Sprintf("%s.%s", tableName.SchemaName, tableName.TableName),
			"displayName": fmt.Sprintf("%s.%s", tableName.SchemaName, tableName.TableName),
			"columns":     columns,
		})
	}

	return allColumns, nil
}

func FormatPostgresConfig(connectionString string) (PostgresConnectionConfig, error) {
	config := PostgresConnectionConfig{
		ConnectionString: connectionString,
		SSL: &struct {
			RejectUnauthorized bool
			CA                 *string
			Key                *string
			Cert               *string
		}{
			RejectUnauthorized: false,
		},
	}
	return config, nil
}

// Helper function to filter foreign keys.
func filterPostgresForeignKeys(keys []string) []string {
	filtered := []string{}
	for _, key := range keys {
		if key != "id" && key != "_id_" {
			filtered = append(filtered, key)
		}
	}
	return filtered
}
