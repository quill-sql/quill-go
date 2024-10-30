package models

type QueryResults struct {
	QueryResults []map[string]interface{} `json:"queryResults"`
	MappedArray  []interface{}            `json:"mappedArray,omitempty"`
	Columns      []Column                 `json:"columns,omitempty"`
}

type QuillClientParams struct {
	PrivateKey               string
	DatabaseConnectionString *string
	DatabaseConfig           interface{}
	DatabaseType             string
	Cache                    *CacheCredentials
	MetadataServerURL        *string
}

type DashboardFilter struct {
	Table      string
	Field      string
	LabelField string
	Label      string
	FilterType string
	Options    []interface{}
	Query      string
}

type Request struct {
	Body RequestBody
}

type RequestBody struct {
	Metadata QuillRequestMetadata `json:"metadata"`
}

type DashboardConfig struct {
	Filters  []DashboardFilter
	Sections map[string][]DashboardItem `json:"sections"`
	Name     string
}

type Section struct {
	ID            string       `json:"_id"`
	ChartType     string       `json:"chartType"`
	ClientId      string       `json:"clientId"`
	Columns       []Column     `json:"columns"`
	CreatedAt     string       `json:"createdAt"` // or time.Time if you are going to parse it
	CustomerId    string       `json:"customerId"`
	DashboardName string       `json:"dashboardName"`
	DateField     DateField    `json:"dateField"`
	Name          string       `json:"name"`
	QueryString   string       `json:"queryString"`
	Template      bool         `json:"template"`
	UpdatedAt     string       `json:"updatedAt"` // or time.Time if you are going to parse it
	XAxisField    string       `json:"xAxisField"`
	XAxisFormat   string       `json:"xAxisFormat"`
	XAxisLabel    string       `json:"xAxisLabel"`
	YAxisFields   []YAxisField `json:"yAxisFields"`
	YAxisLabel    string       `json:"yAxisLabel"`
}

type DashboardItem struct {
	ID            string       `json:"_id"`
	ChartType     string       `json:"chartType"`
	ClientId      string       `json:"clientId"`
	Columns       []Column     `json:"columns"`
	CreatedAt     string       `json:"createdAt"` // or time.Time if you are going to parse it
	CustomerId    string       `json:"customerId"`
	DashboardName string       `json:"dashboardName"`
	DateField     DateField    `json:"dateField"`
	Name          string       `json:"name"`
	QueryString   string       `json:"queryString"`
	Template      bool         `json:"template"`
	UpdatedAt     string       `json:"updatedAt"` // or time.Time if you are going to parse it
	XAxisField    string       `json:"xAxisField"`
	XAxisFormat   string       `json:"xAxisFormat"`
	XAxisLabel    string       `json:"xAxisLabel"`
	YAxisFields   []YAxisField `json:"yAxisFields"`
	YAxisLabel    string       `json:"yAxisLabel"`
}

type Column struct {
	ID     string `json:"_id"`
	Field  string `json:"field"`
	Format string `json:"format"`
	Label  string `json:"label"`
}

type YAxisField struct {
	ID     string `json:"_id"`
	Field  string `json:"field"`
	Format string `json:"format"`
	Label  string `json:"label"`
}

type QuillQueryResult struct {
	Data    interface{}   `json:"data,omitempty"`
	Status  string        `json:"status"`
	Error   *string       `json:"error,omitempty"`
	Queries *QueryResults `json:"queries,omitempty"`
}

type Field struct {
	Name       string  `json:"name"`
	FieldType  *string `json:"fieldType"`
	DataTypeID int     `json:"dataTypeID"`
}
