package db

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"regexp"
	"strings"

	"cloud.google.com/go/bigquery"
	"github.com/quill-sql/quill-go/pkg/utils"
	"google.golang.org/api/option"
)

type BigQueryConfig struct {
	DatasetName string
	ProjectID   string
	Credentials map[string]interface{}
}

type ColumnInfo struct {
	ColumnName  string
	DataTypeID  int
	DisplayName string
	FieldType   string
}

type TableColumnInfo struct {
	TableName   string
	DisplayName string
	Columns     []ColumnInfo
}

func FormatBigQueryConfig(connectionString string) (BigQueryConfig, error) {
	jsonStartIndex := strings.Index(connectionString, "{")
	if jsonStartIndex == -1 {
		return BigQueryConfig{}, errors.New("invalid input string. No JSON data found")
	}

	datasetName := strings.TrimSpace(connectionString[:jsonStartIndex])
	jsonString := connectionString[jsonStartIndex:]

	var serviceAccount map[string]interface{}
	if err := json.Unmarshal([]byte(jsonString), &serviceAccount); err != nil {
		return BigQueryConfig{}, errors.New("failed to parse JSON string: " + err.Error())
	}

	projectID, ok := serviceAccount["project_id"].(string)
	if !ok || projectID == "" {
		return BigQueryConfig{}, errors.New("invalid service account JSON. Required fields are missing")
	}

	privateKey, ok := serviceAccount["private_key"].(string)
	if !ok || privateKey == "" {
		return BigQueryConfig{}, errors.New("invalid service account JSON. Required fields are missing")
	}

	return BigQueryConfig{
		DatasetName: datasetName,
		ProjectID:   projectID,
		Credentials: serviceAccount,
	}, nil
}

func ConnectToBigQuery(config BigQueryConfig) (*bigquery.Client, error) {
	ctx := context.Background()
	credentialsJSON, err := json.Marshal(config.Credentials)
	if err != nil {
		return nil, err
	}

	client, err := bigquery.NewClient(ctx, config.ProjectID, option.WithCredentialsJSON(credentialsJSON))
	if err != nil {
		return nil, err
	}

	return client, nil
}

func RunQueryBigQuery(client *bigquery.Client, sql string) (*QueryResults, error) {
	query := client.Query(sql)
	it, err := query.Read(context.TODO())
	if err != nil {
		return nil, err
	}

	var rows []map[string]interface{}
	for {
		var row map[string]bigquery.Value
		err := it.Next(&row)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		convertedRow := make(map[string]interface{})
		for key, value := range row {
			convertedRow[key] = value
		}
		rows = append(rows, convertedRow)
	}

	if len(rows) == 0 {
		return &QueryResults{Fields: []Field{}, Rows: []map[string]interface{}{}}, nil
	}

	var fields []Field
	for name := range rows[0] {
		fields = append(fields, Field{Name: name, DataTypeID: 1043})
	}

	for i := range fields {
		for _, row := range rows {
			if row[fields[i].Name] == nil {
				continue
			}
			fields[i].DataTypeID = inferType(row[fields[i].Name])
			break
		}
	}

	return &QueryResults{Fields: fields, Rows: rows}, nil
}

func GetSchemaBigQuery(client *bigquery.Client) ([]string, error) {
	ctx := context.Background()
	it := client.Datasets(ctx)
	var filtered []string

	for {
		dataset, err := it.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		if dataset.DatasetID != "" {
			filtered = append(filtered, dataset.DatasetID)
		}
	}

	return filtered, nil
}

func GetTablesBySchemaBigQuery(client *bigquery.Client, schemaNames []string) ([]Table, error) {
	ctx := context.Background()
	var allTables []Table

	for _, schema := range schemaNames {
		sql := fmt.Sprintf(
			"SELECT table_name FROM %s.INFORMATION_SCHEMA.TABLES WHERE table_type = 'BASE TABLE' OR table_type = 'VIEW' OR table_type = 'MATERIALIZED VIEW'", schema)

		query := client.Query(sql)
		it, err := query.Read(ctx)
		if err != nil {
			return nil, err
		}

		for {
			var row map[string]bigquery.Value
			err := it.Next(&row)
			if err == io.EOF {
				break
			}
			if err != nil {
				return nil, err
			}

			tableName, ok := row["table_name"].(string)
			if !ok {
				return nil, fmt.Errorf("failed to retrieve table_name from schema %s", schema)
			}

			allTables = append(allTables, Table{
				TableName:  tableName,
				SchemaName: schema,
			})
		}
	}

	return allTables, nil
}

func GetColumnsByTableBigQuery(client *bigquery.Client, schemaName, tableName string) ([]string, error) {
	ctx := context.Background()
	sql := fmt.Sprintf(
		"SELECT column_name FROM %s.INFORMATION_SCHEMA.COLUMNS WHERE table_name = '%s'",
		schemaName, tableName)

	query := client.Query(sql)
	it, err := query.Read(ctx)
	if err != nil {
		return nil, err
	}

	var columns []string
	for {
		var row map[string]bigquery.Value
		err := it.Next(&row)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		columnName, ok := row["column_name"].(string)
		if !ok {
			return nil, fmt.Errorf("failed to retrieve column_name from table %s", tableName)
		}

		columns = append(columns, columnName)
	}

	return columns, nil
}

func GetForeignKeysBigQuery(connection *bigquery.Client, schemaName, tableName, primaryKey string) ([]string, error) {
	depluralizedTableName := utils.Depluralize(tableName)
	sql := fmt.Sprintf(`SELECT column_name FROM %s.INFORMATION_SCHEMA.COLUMNS
		WHERE table_name != '%s'
		AND (column_name = '%s'
		OR column_name = '%s_%s'
		OR column_name = '%s%s')`, schemaName, tableName, primaryKey, depluralizedTableName, primaryKey, depluralizedTableName, utils.Capitalize(primaryKey))

	results, err := RunQueryBigQuery(connection, sql)
	if err != nil {
		return nil, err
	}

	foreignKeys := filterBigQueryForeignKeys(results.Rows)

	if len(foreignKeys) == 0 {
		sql = fmt.Sprintf(`SELECT column_name FROM %s.INFORMATION_SCHEMA.COLUMNS
			WHERE table_name != '%s'
			AND (column_name LIKE '%s%%'
			OR column_name LIKE '%%_id'
			OR column_name LIKE '%%Id'
			OR column_name LIKE '%%_%s'
			OR column_name LIKE '%%%s')`, schemaName, tableName, depluralizedTableName, primaryKey, utils.Capitalize(primaryKey))

		results, err := RunQueryBigQuery(connection, sql)
		if err != nil {
			return nil, err
		}

		foreignKeys = uniqueForeignKeys(results.Rows)
	}

	return foreignKeys, nil
}

func filterBigQueryForeignKeys(rows []map[string]interface{}) []string {
	var foreignKeys []string
	for _, row := range rows {
		columnName, ok := row["column_name"].(string)
		if ok && columnName != "id" && columnName != "_id_" {
			foreignKeys = append(foreignKeys, columnName)
		}
	}
	return unique(foreignKeys)
}

func uniqueForeignKeys(rows []map[string]interface{}) []string {
	var foreignKeys []string
	for _, row := range rows {
		columnName, ok := row["column_name"].(string)
		if ok {
			foreignKeys = append(foreignKeys, columnName)
		}
	}
	return unique(foreignKeys)
}

func unique(keys []string) []string {
	keysMap := make(map[string]struct{})
	for _, key := range keys {
		keysMap[key] = struct{}{}
	}
	uniqueKeys := make([]string, 0, len(keysMap))
	for key := range keysMap {
		uniqueKeys = append(uniqueKeys, key)
	}
	return uniqueKeys
}

func convertBigQueryTypeToPostgresOID(bigQueryType string) int {
	typeToOidMap := map[string]int{
		"VARCHAR":   1043,
		"INTEGER":   23,
		"FLOAT":     700,
		"TIMESTAMP": 1114,
		"DATE":      1082,
		"BOOL":      16,
	}

	postgresType := strings.ToUpper(bigQueryType)
	if oid, exists := typeToOidMap[postgresType]; exists {
		return oid
	}
	return 1043
}

func GetSchemaColumnInfoBigQuery(
	connection *bigquery.Client,
	//schemaName string,
	tableNames []Table,
) ([]TableColumnInfo, error) {

	allColumns := make([]TableColumnInfo, 0, len(tableNames))

	for _, tableName := range tableNames {
		query := fmt.Sprintf(`
			SELECT column_name AS columnName, data_type AS dataType
			FROM %s.INFORMATION_SCHEMA.COLUMNS
			WHERE table_name = '%s'
			ORDER BY ordinal_position`,
			tableName.SchemaName, tableName.TableName)

		results, err := RunQueryBigQuery(connection, query)
		if err != nil {
			return nil, err
		}

		tableInfo := TableColumnInfo{
			TableName:   fmt.Sprintf("%s.%s", tableName.SchemaName, tableName.TableName),
			DisplayName: fmt.Sprintf("%s.%s", tableName.SchemaName, tableName.TableName),
			Columns:     make([]ColumnInfo, len(results.Rows)),
		}

		for i, row := range results.Rows {
			tableInfo.Columns[i] = ColumnInfo{
				ColumnName:  row["columnName"].(string),
				DisplayName: row["columnName"].(string),
				DataTypeID:  convertBigQueryTypeToPostgresOID(row["dataType"].(string)),
				FieldType:   row["dataType"].(string),
			}
		}

		allColumns = append(allColumns, tableInfo)
	}

	return allColumns, nil
}

// inferType determines the type based on the value.
func inferType(elem interface{}) int {
	switch v := elem.(type) {
	case int, int8, int16, int32, int64:
		return 23 // integer
	case float32, float64:
		return 700 // real
	case map[string]interface{}:
		if value, ok := v["value"].(string); ok {
			if matched, _ := regexp.MatchString(`^\d{4}-\d{2}-\d{2}$`, value); matched {
				return 1082 // YYYY-MM-DD format
			}
			if matched, _ := regexp.MatchString(`^\d{2}/\d{2}/\d{2,4}$`, value); matched {
				return 1082 // MM/DD/YYYY or MM/DD/YY format
			}
			if matched, _ := regexp.MatchString(`^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?$`, value); matched {
				return 1114 // YYYY-MM-DDTHH:MM:SS[.fraction] format
			}
			if matched, _ := regexp.MatchString(`^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?Z$`, value); matched {
				return 1184 // YYYY-MM-DDTHH:MM:SS[.fraction]Z format
			}
			if matched, _ := regexp.MatchString(`^\d{2}:\d{2}:\d{2}$`, value); matched {
				return 1083 // HH:MM:SS format
			}
		}
		return 1043 // varchar
	case string:
		if matched, _ := regexp.MatchString(`^\d{4}-\d{2}-\d{2}$`, v); matched {
			return 1082 // date
		}
		if matched, _ := regexp.MatchString(`^\d{2}/\d{2}/\d{2,4}$`, v); matched {
			return 1082 // date
		}
		if matched, _ := regexp.MatchString(`^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?$`, v); matched {
			return 1114 // timestamp without timezone
		}
		if matched, _ := regexp.MatchString(`^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?Z$`, v); matched {
			return 1184 // timestamp with timezone
		}
		if matched, _ := regexp.MatchString(`^\d{2}:\d{2}:\d{2}$`, v); matched {
			return 1083 // time
		}
		return 1043 // varchar
	default:
		return 1043 // default or unknown type
	}
}
