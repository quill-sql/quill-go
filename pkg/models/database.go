package models

type DatabaseCredentials struct {
	DatabaseConnectionString        string
	StagingDatabaseConnectionString string
	DatabaseType                    string
}
