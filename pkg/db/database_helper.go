package db

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"cloud.google.com/go/bigquery"
)

type DatabaseConnection interface{} // TODO: Define this type

type QueryResults struct {
	Fields []Field                  `json:"fields"`
	Rows   []map[string]interface{} `json:"rows"`
}

type Field struct {
	Name       string `json:"name"`
	DataTypeID int    `json:"dataTypeID"`
}

type Table struct {
	TableName  string
	SchemaName string
}

func GetDatabaseCredentials(databaseType, connectionString string) (interface{}, error) {
	switch strings.ToLower(databaseType) {
	case "postgresql":
		return FormatPostgresConfig(connectionString)
	case "bigquery":
		return FormatBigQueryConfig(connectionString)
	default:
		return nil, fmt.Errorf("unsupported database type: %s", databaseType)
	}
}

func ConnectToDatabase(databaseType string, config interface{}) (DatabaseConnection, error) {
	switch strings.ToLower(databaseType) {
	case "postgresql":
		return ConnectToPostgres(config.(PostgresConnectionConfig))
	case "bigquery":
		return ConnectToBigQuery(config.(BigQueryConfig))
	default:
		return nil, fmt.Errorf("unsupported database type: %s", databaseType)
	}
}

func RunQueryByDatabase(databaseType string, connection DatabaseConnection, query string) (*QueryResults, error) {
	switch strings.ToLower(databaseType) {
	case "postgresql":
		return RunQueryPostgres(connection.(*sql.DB), query)
	case "bigquery":
		return RunQueryBigQuery(connection.(*bigquery.Client), query)
	default:
		return nil, fmt.Errorf("unsupported database type: %s", databaseType)
	}
}

func DisconnectFromDatabase(databaseType string, connection DatabaseConnection) error {
	switch strings.ToLower(databaseType) {
	case "postgresql":
		return DisconnectFromPostgres(connection.(*sql.DB))
	case "bigquery":
		return nil // BigQuery does not need to be disconnected
	default:
		return fmt.Errorf("unsupported database type: %s", databaseType)
	}
}

func GetSchemasByDatabase(databaseType string, connection DatabaseConnection) ([]string, error) {
	switch strings.ToLower(databaseType) {
	case "postgresql":
		return GetSchemasPostgres(connection.(*sql.DB))
	case "bigquery":
		return GetSchemaBigQuery(connection.(*bigquery.Client))
	default:
		return nil, fmt.Errorf("unsupported database type: %s", databaseType)
	}
}

func GetTablesBySchemaByDatabase(
	databaseType string,
	connection DatabaseConnection,
	schemaName interface{},
) ([]Table, error) {
	switch strings.ToLower(databaseType) {
	case "postgres", "postgresql":
		return GetTablesBySchemaPostgres(connection.(*sql.DB), schemaName.([]string))
	case "bigquery":
		return GetTablesBySchemaBigQuery(connection.(*bigquery.Client), schemaName.([]string))
	default:
		return nil, fmt.Errorf("unsupported database type: %s", databaseType)
	}
}

func GetColumnsByTableByDatabase(
	ctx context.Context,
	databaseType string,
	connection DatabaseConnection,
	schemaName string,
	tableName string,
) ([]string, error) {
	switch strings.ToLower(databaseType) {
	case "postgres", "postgresql":
		return GetColumnsByTablePostgres(connection.(*sql.DB), schemaName, tableName)
	case "bigquery":
		return GetColumnsByTableBigQuery(connection.(*bigquery.Client), schemaName, tableName)
	default:
		return nil, fmt.Errorf("unsupported database type: %s", databaseType)
	}
}

func GetForeignKeysByDatabase(
	ctx context.Context,
	databaseType string,
	connection DatabaseConnection,
	schemaName string,
	tableName string,
	primaryKey string,
) ([]string, error) {
	switch strings.ToLower(databaseType) {
	case "postgres", "postgresql":
		return GetForeignKeysPostgres(connection.(*sql.DB), schemaName, tableName, primaryKey)
	case "bigquery":
		return GetForeignKeysBigQuery(connection.(*bigquery.Client), schemaName, tableName, primaryKey)
	default:
		return nil, fmt.Errorf("unsupported database type: %s", databaseType)
	}
}

func GetColumnInfoBySchemaByDatabase(
	databaseType string,
	connection DatabaseConnection,
	//schemaName string, // unused
	tables []Table,
) (interface{}, error) {
	switch strings.ToLower(databaseType) {
	case "postgres", "postgresql":
		return GetSchemaColumnInfoPostgres(connection.(*sql.DB) /*schemaName,*/, tables)
	case "bigquery":
		return GetSchemaColumnInfoBigQuery(connection.(*bigquery.Client) /*schemaName,*/, tables)
	default:
		return nil, fmt.Errorf("unsupported database type: %s", databaseType)
	}
}
