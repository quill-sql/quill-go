package quill

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"

	_ "github.com/lib/pq"
)

type Client struct {
	PrivateKey               string
	DatabaseConnectionString string
}

type QuillClientParams struct {
	PrivateKey               string
	DatabaseConnectionString string
}

func NewClient(clientParams QuillClientParams) *Client {

	client := &Client{
		PrivateKey:               clientParams.PrivateKey,
		DatabaseConnectionString: clientParams.DatabaseConnectionString,
	}

	return client
}

type RequestMetadata struct {
	Task string
	// a query to be run
	Query string
	// a report to be fetched
	ID      string
	Filters []Filter
	// dashboard item fields
	Name          string
	XAxisField    string
	YAxisFields   []FormattedColumn
	XAxisLabel    string
	XAxisFormat   FieldFormat
	YAxisLabel    string
	ChartType     string
	DashboardName string
	Columns       []FormattedColumn
	DateField     DateField
	Template      bool
}

type Filter struct {
	// ex: building
	Table string
	// ex: buildingId
	Field string
	// ex: buildingName - for the case that it's not unique
	LabelField string
	// label above field: "Building name"
	Label string
	// NUMBER_RANGE
	// SQL_LIKE: field like '%thing'
	// SQL_IN: field in (option1, option2, option3)
	// DATE_RANGE
	// BOOLEAN lmao
	FilterType string
	Options    []interface{}
	Query      string
}

type FieldFormat string

const (
	WholeNumber      FieldFormat = "whole_number"
	OneDecimalPlace  FieldFormat = "one_decimal_place"
	TwoDecimalPlaces FieldFormat = "two_decimal_places"
	DollarAmount     FieldFormat = "dollar_amount"
	MMMYYYY          FieldFormat = "MMM_yyyy"
	MMMDDYYYY        FieldFormat = "MMM_dd_yyyy"
	MMMDDMMMDD       FieldFormat = "MMM_dd-MMM_dd"
	MMMDDHHMMAPPM    FieldFormat = "MMM_dd_hh:mm_ap_pm"
	HHAPPM           FieldFormat = "hh_ap_pm"
	Percent          FieldFormat = "percent"
	StringFormat     FieldFormat = "string"
)

type FormattedColumn struct {
	Label     string
	Field     string
	ChartType string
	Format    FieldFormat
}

type Request struct {
	Body RequestBody
}

type RequestBody struct {
	Metadata RequestMetadata `json:"metadata"`
}

type QueryParams struct {
	OrgID       string
	Metadata    RequestMetadata
	Environment string
}

type DashboardConfig struct {
	Filters  []Filter
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

type DateField struct {
	Field string `json:"field"`
	Table string `json:"table"`
}

type YAxisField struct {
	ID     string `json:"_id"`
	Field  string `json:"field"`
	Format string `json:"format"`
	Label  string `json:"label"`
}

type QueryResult struct {
	rows   []interface{}
	fields []interface{}
}

type Field struct {
	Name       string `json:"name"`
	FieldType  string `json:"fieldType"`
	DataTypeID int    `json:"dataTypeID"`
}

func (c *Client) Query(orgID string, metadata RequestMetadata) (map[string]interface{}, error) {

	db, err := sql.Open("postgres", c.DatabaseConnectionString)
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	if metadata.Task == "config" {
		var dashConfig map[string]interface{}

		// Prepare request URL and headers
		reqUrl, err := url.Parse("https://quill-344421.uc.r.appspot.com/config")
		if err != nil {
			return nil, err
		}
		q := reqUrl.Query()
		q.Add("orgId", orgID)
		q.Add("name", metadata.Name)
		reqUrl.RawQuery = q.Encode()

		req, err := http.NewRequest(http.MethodGet, reqUrl.String(), nil)

		if err != nil {
			return nil, err
		}
		req.Header.Add("Authorization", "Bearer "+c.PrivateKey)

		// Send request
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		// Parse response
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		json.Unmarshal(body, &dashConfig)

		return dashConfig, nil
	}

	if metadata.Task == "create" {

		reqUrl, err := url.Parse("https://quill-344421.uc.r.appspot.com/item")
		if err != nil {
			return nil, err
		}
		q := reqUrl.Query()
		q.Add("orgId", orgID)
		reqUrl.RawQuery = q.Encode()
		jsonData, err := json.Marshal(metadata)
		if err != nil {
			return nil, err
		}

		req, err := http.NewRequest(http.MethodPost, reqUrl.String(), bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")

		if err != nil {
			return nil, err
		}
		req.Header.Add("Authorization", "Bearer "+c.PrivateKey)

		var itemBody map[string]interface{}

		// Send request
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		// Parse response
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		json.Unmarshal(body, &itemBody)

		return itemBody, nil
	}

	if metadata.Task == "item" {
		// const resp = await axios.get(
		//     "https://quill-344421.uc.r.appspot.com/selfhostitem",
		//     {
		//       params: {
		//         id,
		//         orgId,
		//       },
		//       headers: {
		//         Authorization: `Bearer ${privateKey}`,
		//       },
		//     }
		//   );
		reqUrl, err := url.Parse("https://quill-344421.uc.r.appspot.com/selfhostitem")
		if err != nil {
			return nil, err
		}
		q := reqUrl.Query()
		q.Add("orgId", orgID)
		q.Add("id", metadata.ID)
		reqUrl.RawQuery = q.Encode()

		req, err := http.NewRequest(http.MethodGet, reqUrl.String(), nil)

		if err != nil {
			return nil, err
		}
		req.Header.Add("Authorization", "Bearer "+c.PrivateKey)

		// Send request
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		var itemBody map[string]interface{}
		// Parse response
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		json.Unmarshal(body, &itemBody)

		query, ok := itemBody["queryString"].(string)
		if !ok {
			log.Fatal("query is not a string")
		}
		queryResult, err := queryDatabase(db, query, orgID, c.PrivateKey)
		if err != nil {
			log.Fatal(err)
		}
		itemBody["rows"] = queryResult["rows"]
		itemBody["fields"] = queryResult["fields"]
		return itemBody, nil

	}

	if metadata.Task == "query" {
		queryResult, err := queryDatabase(db, metadata.Query, orgID, c.PrivateKey)
		if err != nil {
			log.Fatal(err)
		}
		return queryResult, nil
	}

	return nil, nil
}

func queryDatabase(db *sql.DB, query string, orgID string, privateKey string) (map[string]interface{}, error) {
	queryResult := make(map[string]interface{})
	reqUrl, err := url.Parse("https://quill-344421.uc.r.appspot.com/validate")
	if err != nil {
		return nil, err
	}
	q := reqUrl.Query()
	q.Add("orgId", orgID)
	reqUrl.RawQuery = q.Encode()

	queryInBody := map[string]string{
		"query": query,
	}
	jsonData, err := json.Marshal(queryInBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, reqUrl.String(), bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+privateKey)

	var validationBody map[string]interface{}

	// Send request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Parse response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(body, &validationBody)

	// run query on db

	query, ok := validationBody["query"].(string)
	if !ok {
		log.Fatal("query is not a string")
	}

	fieldToRemove, ok := validationBody["fieldToRemove"].(string)
	if !ok {
		log.Fatal("fieldToRemove is not a string")
	}

	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		log.Fatal(err)
	}
	columnTypes, err := rows.ColumnTypes()
	if err != nil {
		log.Fatal(err)
	}
	var fields []Field

	typeRows, err := db.Query("SELECT oid, typname FROM pg_type")
	if err != nil {
		log.Fatal(err)
	}
	defer typeRows.Close()

	typeData := make(map[string]int)

	for typeRows.Next() {
		var oid int
		var typname string
		err = typeRows.Scan(&oid, &typname)
		if err != nil {
			log.Fatal(err)
		}
		typeData[typname] = oid
	}

	for _, column := range columnTypes {
		// Append the column info to the slice
		// fmt.Println(column.DatabaseTypeName())
		fields = append(fields, Field{
			Name:       column.Name(),
			FieldType:  strings.ToLower(column.DatabaseTypeName()),
			DataTypeID: typeData[strings.ToLower(column.DatabaseTypeName())],
		})
	}

	// Create a slice of interface{}'s to represent each column,
	// and a second slice to contain pointers to each item in the columns slice.
	values := make([]interface{}, len(columns))
	pointers := make([]interface{}, len(columns))
	for i := range values {
		pointers[i] = &values[i]
	}

	var rowsData []map[string]interface{}

	// Loop through rows
	for rows.Next() {
		err := rows.Scan(pointers...)
		if err != nil {
			log.Fatal(err)
		}

		row := make(map[string]interface{})
		for i, colName := range columns {
			val := values[i]

			// Check if the value is nil (NULL value)
			if b, ok := val.([]byte); ok {
				row[colName] = string(b)
			} else {
				row[colName] = val
			}
		}
		delete(row, fieldToRemove)

		rowsData = append(rowsData, row)

	}

	queryResult["rows"] = rowsData

	queryResult["fields"] = fields

	return queryResult, nil
}
