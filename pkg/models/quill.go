package models

import (
	"time"
)

// type OrgID interface{}

// func (o *OrgID) UnmarshalJSON(data []byte) error {
// 	var id int
// 	if err := json.Unmarshal(data, &id); err == nil {
// 		*o = OrgID(id)
// 		return nil
// 	}

// 	var idStr string
// 	if err := json.Unmarshal(data, &idStr); err == nil {
// 		intID, err := strconv.Atoi(idStr)
// 		if err != nil {
// 			return err
// 		}
// 		*o = OrgID(intID)
// 		return nil
// 	}

// 	return fmt.Errorf("OrgID must be an int or a string representing an int")
// }

type DateField struct {
	Table string `json:"table,omitempty"`
	Field string `json:"field,omitempty"`
}

type AdditionalProcessing struct {
	GetSchema           *bool    `json:"getSchema,omitempty"`
	GetColumns          *bool    `json:"getColumns,omitempty"`
	GetColumnsForSchema *bool    `json:"getColumnsForSchema,omitempty"`
	GetTables           *bool    `json:"getTables,omitempty"`
	Schema              *string  `json:"schema,omitempty"`
	SchemaNames         []string `json:"schemaNames,omitempty"`
	Table               *string  `json:"table,omitempty"`
	FieldsToRemove      []string `json:"fieldsToRemove,omitempty"`
	ArrayToMap          *struct {
		ArrayName string
		Field     string
	} `json:"arrayToMap,omitempty"`
	OverridePost     *bool `json:"overridePost,omitempty"`
	ConvertDatatypes *bool `json:"convertDatatypes,omitempty"`
	LimitThousand    *bool `json:"limitThousand,omitempty"`
	LimitBy          *int  `json:"limitBy,omitempty"`
}

type QuillRequest struct {
	OrgID    *string              `json:"orgId"`
	Metadata QuillRequestMetadata `json:"metadata"`
}

type QuillRequestMetadata struct {
	Task                    string                    `json:"task,omitempty"`
	Queries                 *[]string                 `json:"queries,omitempty"`
	PreQueries              *[]string                 `json:"preQueries,omitempty"`
	RunQueryConfig          *AdditionalProcessing     `json:"runQueryConfig,omitempty"`
	Query                   *string                   `json:"query,omitempty"`
	Id                      *string                   `json:"id,omitempty"`
	Filters                 *[]interface{}            `json:"filters,omitempty"`
	Name                    *string                   `json:"name,omitempty"`
	XAxisField              *string                   `json:"xAxisField,omitempty"`
	YAxisFields             *[]FormattedColumn        `json:"yAxisFields,omitempty"`
	XAxisLabel              *string                   `json:"xAxisLabel,omitempty"`
	XAxisFormat             *FieldFormat              `json:"xAxisFormat,omitempty"`
	YAxisLabel              *string                   `json:"yAxisLabel,omitempty"`
	ChartType               *string                   `json:"chartType,omitempty"`
	DashboardName           *string                   `json:"dashboardName,omitempty"`
	Columns                 *[]FormattedColumn        `json:"columns,omitempty"`
	DateField               *DateField                `json:"dateField,omitempty"`
	Template                *bool                     `json:"template,omitempty"`
	ClientId                *string                   `json:"clientId,omitempty"`
	Deleted                 *bool                     `json:"deleted,omitempty"`
	DatabaseType            *string                   `json:"databaseType,omitempty"`
	OrgID                   interface{}               `json:"orgId,omitempty"` // sometimes int, sometimes string
	DashboardItemId         *string                   `json:"dashboardItemId,omitempty"`
	Ast                     map[string]interface{}    `json:"ast,omitempty"`
	PublicKey               *string                   `json:"publicKey,omitempty"`
	Pivot                   *bool                     `json:"pivot,omitempty"`
	UseNewNodeSql           *bool                     `json:"useNewNodeSql,omitempty"`
	UseUpdatedDataGathering *bool                     `json:"useUpdatedDataGathering,omitempty"`
	AdditionalProcessing    map[string]interface{}    `json:"additionalProcessing,omitempty"`
	ReportId                *string                   `json:"reportId,omitempty"`
	Tables                  *[]string                 `json:"tables,omitempty"`
	Table                   *string                   `json:"table,omitempty"`
	RemoveCustomerField     *bool                     `json:"removeCustomerField,omitempty"`
	GetCustomFields         *bool                     `json:"getCustomFields,omitempty"`
	CustomFieldsByTable     *[]map[string]interface{} `json:"customFieldsByTable,omitempty"`
	TemplateReportIds       *[]string                 `json:"templateReportIds,omitempty"`
	ReferencedTables        *[]string                 `json:"referencedTables,omitempty"`
	ReferencedColumns       map[string]interface{}    `json:"referencedColumns,omitempty"`
}

type QuillQueryParams struct {
	OrgId    string
	Metadata QuillRequestMetadata
	Filters  *[]Filter
}

type QuillConfig struct {
	PrivateKey string
	DB         *DatabaseCredentials
	Cache      *CacheCredentials
}

type QuillClientResponse struct {
	Queries        []interface{}         `json:"queries"`
	Metadata       interface{}           `json:"metadata"`
	RunQueryConfig *AdditionalProcessing `json:"runQueryConfig,omitempty"`
	Error          *string               `json:"error,omitempty"`
}

// Mapable interface represents a cache that can get and set values
type RedisMapable interface {
	Get(key string) (string, error)
	Set(key string, value interface{}, expiration time.Duration) error
}

type QueryResult struct {
	Fields []map[string]interface{} `json:"fields"`
	Rows   []map[string]interface{} `json:"rows"`
}

// PgError represents a PostgreSQL error
type PgError struct {
	Message  string
	Detail   string
	Hint     string
	Position string
}

func (e *PgError) Error() string {
	return e.Message
}
