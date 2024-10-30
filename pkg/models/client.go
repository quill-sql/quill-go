package models

// Client represents the configuration and settings for a database client.
type Client struct {
	Name                            string
	DatabaseConnectionString        string
	ETLDatabaseConnectionString     string
	StagingDatabaseConnectionString string
	CustomerTableTitleFieldName     string
	DatabaseType                    string
	CustomerFieldName               string
	CustomerTableFieldName          string
	CustomerTableName               string
	CustomerView                    string
	CustomerFieldType               string
	UseSSL                          bool
	ServerCA                        string
	ClientCert                      string
	ClientKey                       string
	DefaultQuery                    string
	IgnoreDarkMode                  bool
	DomainName                      string
	HideSqlEditor                   bool
	AdminCustomerID                 string
	StagingAdminCustomerID          string
	CacheCloudConfig                CacheCloudConfig
}

// CacheCloudConfig holds the configuration for cache-related settings.
type CacheCloudConfig struct {
	Type    CacheType
	Default interface{}
}

// CacheType holds the cache-related boolean settings.
type CacheType struct {
	CacheQueries bool
}
