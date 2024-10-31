package core

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"runtime/debug"
	"strings"
	"sync"

	_ "github.com/lib/pq"
	"github.com/quill-sql/quill-go/pkg/db"
	"github.com/quill-sql/quill-go/pkg/models"
	"github.com/quill-sql/quill-go/pkg/utils"
)

type Quill struct {
	TargetConnection *db.CachedConnection
	baseURL          string
	config           struct {
		headers struct {
			authorization string
		}
	}
	httpClient *http.Client
}

func NewQuill(data models.QuillClientParams) (*Quill, error) {
	quill := &Quill{
		httpClient: &http.Client{}, // Initialize the HTTP client
	}

	// Set base URL
	if data.MetadataServerURL != nil {
		quill.baseURL = *data.MetadataServerURL
	} else {
		quill.baseURL = "http://localhost:8080"
	}

	// Set config headers
	quill.config.headers.authorization = fmt.Sprintf("Bearer %s", data.PrivateKey)

	// Handle database credentials
	var credentials interface{}
	var err error
	if data.DatabaseConnectionString != nil {
		credentials, err = db.GetDatabaseCredentials(data.DatabaseType, *data.DatabaseConnectionString)
	} else {
		credentials = data.DatabaseConfig
	}
	if err != nil {
		return nil, err
	}

	var emptyCache models.CacheCredentials
	cache := emptyCache
	if data.Cache != nil {
		cache = *data.Cache
	}
	// Initialize target connection
	connection, err := db.NewCachedConnection(data.DatabaseType, credentials, cache)

	if err != nil {
		return nil, err
	}
	if connection == nil {
		return nil, errors.New("failed to create CachedConnection: connection is nil")
	}

	quill.TargetConnection = connection

	return quill, err
}

// Query executes a query based on the provided parameters.
func (c *Quill) Query(params models.QuillQueryParams) (result models.QuillQueryResult, err error) {
	c.TargetConnection.OrgID = params.OrgId
	var responseMetadata interface{} = make(map[string]interface{})
	defer func() {
		if r := recover(); r != nil {
			errorMsg1 := fmt.Sprintf("Panic occurred: %v\nStack trace: %s", r, debug.Stack())
			// Log the error for better traceability
			fmt.Printf("[ERROR] %s", errorMsg1)
			errorMsg := fmt.Sprintf("%v", r)
			result = models.QuillQueryResult{
				Status: "error",
				Error:  &errorMsg,
			}
		}
	}()
	if params.Metadata.Task == "" {
		err := string("Missing task.")
		return models.QuillQueryResult{
			Status: "error",
			Error:  &err,
			Data:   make(map[string]interface{}),
		}, nil
	}

	preQueryResults := make(map[string]interface{})
	if params.Metadata.PreQueries != nil {
		preQueriesInterface := utils.ConvertStringToInterfaceSlice(*params.Metadata.PreQueries)
		preQueryResultsData, err := c.runQueries(preQueriesInterface, c.TargetConnection.DatabaseType, params.Metadata.DatabaseType, params.Metadata.RunQueryConfig)
		if err == nil && preQueryResultsData != nil {
			preQueryResults = preQueryResultsData
		}
		if err != nil {
			if params.Metadata.Task == "update-view" {
				payload := map[string]interface{}{
					"table":    params.Metadata.Name,
					"clientId": params.Metadata.ClientId,
					"error":    err.Error(),
				}
				_, _ = c.postQuill("set-broken-view", payload)
				errorMsg := err.Error()
				return models.QuillQueryResult{
					Status: "error",
					Error:  &errorMsg,
					Data:   responseMetadata,
				}, nil
			}
		}

	}
	if params.Metadata.RunQueryConfig != nil && params.Metadata.RunQueryConfig.OverridePost != nil && *params.Metadata.RunQueryConfig.OverridePost {
		return models.QuillQueryResult{
			Data:   preQueryResults,
			Status: "success",
		}, nil
	}
	payload := make(map[string]interface{})
	// Add metadata fields to the payload
	metadata, err := utils.StructToMap(params.Metadata)
	if err != nil {
		errorMsg := err.Error()
		return models.QuillQueryResult{
			Status: "error",
			Error:  &errorMsg,
			Data:   make(map[string]interface{}),
		}, nil
	}
	for k, v := range metadata {
		payload[k] = v
	}
	// Add preQueryResults fields to the payload
	for k, v := range preQueryResults {
		payload[k] = v
	}
	// Add additional fields
	payload["orgId"] = params.OrgId
	if params.Metadata.PreQueries != nil && len(*params.Metadata.PreQueries) > 0 {
		payload["viewQuery"] = strings.TrimSuffix((*params.Metadata.PreQueries)[0], ";")
	} else {
		payload["viewQuery"] = nil
	}
	if params.Filters != nil {
		// payload["sdkFilters"] = params.Filters
		// map elements of params.Filters with models.convertCustomFilter
		var baseFilters []models.BaseFilter
		for _, filter := range *params.Filters {
			baseFilters = append(baseFilters, models.ConvertCustomFilter(filter))
		}
		payload["sdkFilters"] = baseFilters
	}
	response, err := c.postQuill(params.Metadata.Task, payload)
	if err != nil {
		errorMsg := err.Error()
		return models.QuillQueryResult{
			Status: "error",
			Error:  &errorMsg,
			Data:   make(map[string]interface{}),
		}, nil
	}
	if response.Error != nil && *response.Error != "" {
		metadata := response.Metadata
		if metadata == nil {
			metadata = make(map[string]interface{})
		}
		return models.QuillQueryResult{
			Status: "error",
			Error:  response.Error,
			Data:   metadata,
		}, nil
	}
	if response.Metadata != nil {
		responseMetadata = response.Metadata
	}
	var additionalProcessing models.AdditionalProcessing
	// Check if runQueryConfig is present and not nil
	if config, ok := responseMetadata.(map[string]interface{})["runQueryConfig"]; ok && config != nil {
		configBytes, err := json.Marshal(config)
		if err != nil {
			errorMsg := err.Error()
			return models.QuillQueryResult{
				Status: "error",
				Error:  &errorMsg,
			}, nil
		}
		// Unmarshal JSON to AdditionalProcessing
		if err := json.Unmarshal(configBytes, &additionalProcessing); err != nil {
			errorMsg := err.Error()
			return models.QuillQueryResult{
				Status: "error",
				Error:  &errorMsg,
			}, nil
		}
	}
	results, err := c.runQueries(response.Queries, c.TargetConnection.DatabaseType, params.Metadata.DatabaseType, &additionalProcessing)
	if err != nil {
		errorMsg := err.Error()
		if responseMetadata == nil {
			responseMetadata = make(map[string]interface{})
		}
		return models.QuillQueryResult{
			Status: "error",
			Error:  &errorMsg,
			Data:   responseMetadata,
		}, nil
	}
	var additional models.AdditionalProcessing
	err = utils.MapToStruct(responseMetadata.(map[string]interface{}), &additional)
	if err != nil {
		errorMsg := err.Error()
		return models.QuillQueryResult{
			Status: "error",
			Error:  &errorMsg,
		}, nil
	}
	if results["mappedArray"] != nil && additional.ArrayToMap != nil {
		arrayToMap := additional.ArrayToMap
		mappedArray := results["mappedArray"].([]interface{})
		for index, array := range mappedArray {
			// Assuming responseMetadata is a map, adjust accordingly if not
			arr := responseMetadata.(map[string]interface{})[arrayToMap.ArrayName].([]interface{})
			arr[index].(map[string]interface{})[arrayToMap.Field] = array
		}
		delete(results, "mappedArray")
	}

	queryResults, ok := results["queryResults"].([]map[string]interface{})
	if !ok {
		queryResults = make([]map[string]interface{}, 0)
	}
	mappedArray, ok := results["mappedArray"].([]interface{})
	if !ok {
		mappedArray = []interface{}{}
	}

	columns, ok := results["columns"].([]models.Column)
	if !ok {
		columns = []models.Column{}
	}

	queryResult := models.QueryResults{
		QueryResults: queryResults,
		MappedArray:  mappedArray,
		Columns:      columns,
	}
	if len(queryResults) == 1 {
		// 	err = utils.MapToStruct(responseMetadata.(map[string]interface{}), &additional)
		var qResult models.QueryResult
		err = utils.MapToStruct(queryResults[0], &qResult)
		if err != nil {
			errorMsg := err.Error()
			return models.QuillQueryResult{
				Status: "error",
				Error:  &errorMsg,
			}, nil
		}
		// qResult := queryResults[0].(models.QueryResult)
		if qResult.Rows != nil {
			responseMetadata.(map[string]interface{})["rows"] = qResult.Rows
		}
		if qResult.Fields != nil {
			rowsBytes, err := json.Marshal(qResult.Fields)
			if err != nil {
				errorMsg := err.Error()
				return models.QuillQueryResult{
					Status: "error",
					Error:  &errorMsg,
				}, nil
			}

			var fields []db.Field
			if err := json.Unmarshal(rowsBytes, &fields); err != nil {
				errorMsg := err.Error()
				return models.QuillQueryResult{
					Status: "error",
					Error:  &errorMsg,
				}, nil
			}
			responseMetadata.(map[string]interface{})["fields"] = fields
		}
	}
	result = models.QuillQueryResult{
		Data:    responseMetadata,
		Queries: &queryResult,
		Status:  "success",
	}

	return result, nil
}

// RunQueries executes a series of queries based on the provided configuration.
func (c *Quill) runQueries(
	queries []interface{},
	pkDatabaseType string,
	databaseType *string,
	runQueryConfig *models.AdditionalProcessing,
) (map[string]interface{}, error) {
	var results = make(map[string]interface{})
	queryResults := make([]interface{}, 0)
	if queries == nil {
		return map[string]interface{}{
			"queryResults": queryResults,
		}, nil
	}
	if databaseType != nil && !strings.EqualFold(*databaseType, pkDatabaseType) {
		return map[string]interface{}{
			"dbMismatched":        true,
			"backendDatabaseType": pkDatabaseType,
			"queryResults":        queryResults,
		}, nil
	}
	if runQueryConfig != nil {
		if runQueryConfig.ArrayToMap != nil {
			stringQueries, err := utils.ConvertInterfaceToStringSlice(queries)
			if err != nil {
				return nil, err
			}
			mappedArray, err := db.MapQueries(stringQueries, c.TargetConnection)
			if err != nil {
				return nil, err
			}
			return map[string]interface{}{
				"queryResults": queryResults,
				"mappedArray":  mappedArray,
			}, nil
		}
		if runQueryConfig.GetColumns != nil {
			stringQueries, err := utils.ConvertInterfaceToStringSlice(queries)
			if err != nil {
				return nil, err
			}
			query := strings.TrimSuffix(stringQueries[0], ";") + " limit 1000"
			queryResult, err := c.TargetConnection.Query(query)
			if err != nil {
				return nil, err
			}
			columns := make([]utils.TableSchemaInfo, len(queryResult.(*db.QueryResults).Fields))
			for i, field := range queryResult.(*db.QueryResults).Fields {
				fieldName := field.Name
				columns[i] = utils.TableSchemaInfo{
					FieldType:   utils.ConvertTypeToPostgres(field.DataTypeID),
					Name:        field.Name,
					DisplayName: field.Name,
					IsVisible:   true,
					Field:       &fieldName,
				}
			}
			return map[string]interface{}{
				"columns": columns,
			}, nil
		}
		if runQueryConfig.GetColumnsForSchema != nil {
			var wg sync.WaitGroup
			resultsChannel := make(chan struct {
				index  int
				result map[string]interface{}
			}, len(queries))

			for i, table := range queries {
				wg.Add(1)
				go func(i int, t map[string]interface{}) {
					defer wg.Done()

					// Handle "viewQuery", "isSelectStar", "customFieldInfo"
					viewQuery, okViewQuery := t["viewQuery"]
					isSelectStar, okIsSelectStar := t["isSelectStar"]
					customFieldInfo, okCustomFieldInfo := t["customFieldInfo"]

					if (!okViewQuery || viewQuery == nil) ||
						((!okIsSelectStar || isSelectStar == nil) && (!okCustomFieldInfo || customFieldInfo == nil)) {

						// Create merged result
						mergedResult := make(map[string]interface{})
						for k, v := range t {
							mergedResult[k] = v
						}

						resultsChannel <- struct {
							index  int
							result map[string]interface{}
						}{
							index:  i,
							result: mergedResult,
						}
						return
					}

					// Query and handle "limit"
					limit := ""
					if runQueryConfig.LimitBy != nil && *runQueryConfig.LimitBy > 0 {
						limit = fmt.Sprintf(" limit %d", *runQueryConfig.LimitBy)
					}

					queryResult, err := c.TargetConnection.Query(fmt.Sprintf("%s %s", strings.TrimSuffix(t["viewQuery"].(string), ";"), limit))
					if err != nil {
						resultsChannel <- struct {
							index  int
							result map[string]interface{}
						}{
							index: i,
							result: map[string]interface{}{
								"error": fmt.Sprintf("Error fetching columns: %v", err),
							},
						}
						return
					}

					// Process columns
					columns := make([]utils.TableSchemaInfo, len(queryResult.(*db.QueryResults).Fields))
					for i, field := range queryResult.(*db.QueryResults).Fields {
						fieldName := field.Name
						columns[i] = utils.TableSchemaInfo{
							FieldType:   utils.ConvertTypeToPostgres(field.DataTypeID),
							Name:        field.Name,
							DisplayName: field.Name,
							IsVisible:   true,
							Field:       &fieldName,
						}
					}

					// Create merged result
					mergedResult := make(map[string]interface{})
					for k, v := range t {
						mergedResult[k] = v
					}
					mergedResult["columns"] = columns
					mergedResult["rows"] = queryResult.(*db.QueryResults).Rows

					// Send result with index
					resultsChannel <- struct {
						index  int
						result map[string]interface{}
					}{
						index:  i,
						result: mergedResult,
					}

				}(i, table.(map[string]interface{}))
			}
			wg.Wait()
			close(resultsChannel)
			// Collect results and ensure ordered output
			orderedResults := make([]map[string]interface{}, len(queries))
			for result := range resultsChannel {
				orderedResults[result.index] = result.result
			}

			var queryResultsList []map[string]interface{}
			queryResultsList = append(queryResultsList, orderedResults...)

			results["queryResults"] = queryResultsList
			// TODO test this portion
			if runQueryConfig.FieldsToRemove != nil {
				processedResults := make([]map[string]interface{}, len(queryResultsList))
				for i, result := range queryResultsList {
					processedResults[i] = utils.RemoveFields(result, runQueryConfig.FieldsToRemove)
				}
				results["queryResults"] = processedResults
			}

			return results, nil
		}
		if runQueryConfig.GetTables != nil {
			queryResult, err := db.GetTablesBySchemaByDatabase(
				c.TargetConnection.DatabaseType,
				c.TargetConnection.Pool,
				runQueryConfig.SchemaNames,
			)
			if err != nil {
				return nil, err
			}
			schemaInfo, err := db.GetColumnInfoBySchemaByDatabase(
				c.TargetConnection.DatabaseType,
				c.TargetConnection.Pool,
				queryResult,
			)
			if err != nil {
				return nil, err
			}
			return map[string]interface{}{
				"queryResults": schemaInfo,
			}, nil
		}
	}
	if runQueryConfig != nil && runQueryConfig.LimitThousand != nil && *runQueryConfig.LimitThousand {
		for i := range queries {
			queries[i] = strings.TrimSuffix(queries[i].(string), ";") + " limit 1000;"
		}
	} else if runQueryConfig != nil && runQueryConfig.LimitBy != nil && *runQueryConfig.LimitBy > 0 {
		for i := range queries {
			queries[i] = strings.TrimSuffix(queries[i].(string), ";") + fmt.Sprintf(" limit %d", *runQueryConfig.LimitBy)
		}
	}
	var queryResultsList = make([]map[string]interface{}, 0)
	for _, query := range queries {
		queryResult, err := c.TargetConnection.Query(query.(string))
		if err != nil {
			return nil, err
		}
		resultsMap, err := utils.StructToMap(queryResult)
		if err != nil {
			return nil, err
		}
		queryResultsList = append(queryResultsList, resultsMap)
	}
	results["queryResults"] = queryResultsList
	if runQueryConfig != nil && runQueryConfig.FieldsToRemove != nil {
		processedResults := make([]map[string]interface{}, len(queryResultsList))
		for i, result := range queryResultsList {
			processedResults[i] = utils.RemoveFields(result, runQueryConfig.FieldsToRemove)
		}
		results["queryResults"] = processedResults
	}
	if runQueryConfig != nil && runQueryConfig.ConvertDatatypes != nil {
		for i, result := range queryResultsList {
			var fields []db.Field
			if fieldData, ok := result["fields"]; ok {
				fieldBytes, err := json.Marshal(fieldData)
				if err != nil {
					fields = []db.Field{}
				} else {
					// Unmarshal the JSON into the []db.Field slice
					err = json.Unmarshal(fieldBytes, &fields)
					if err != nil {
						fields = []db.Field{}
					}
				}
			} else {
				// If result["fields"] doesn't exist or is not of the correct type, initialize as empty slice
				fields = []db.Field{}
			}

			// Proceed with processing the fields
			columns := make([]utils.TableSchemaInfo, len(fields))
			for j, field := range fields {
				fieldType := utils.ConvertTypeToPostgres(field.DataTypeID)
				fieldName := field.Name
				dataTypeId := field.DataTypeID
				columns[j] = utils.TableSchemaInfo{
					FieldType:   fieldType,
					IsVisible:   true,
					Field:       &fieldName,
					DisplayName: field.Name,
					Name:        field.Name,
					DataTypeID:  &dataTypeId,
				}
			}

			// Replace fields with processed columns
			result["fields"] = columns
			// Ensure the rows are preserved in the result, similar to the JavaScript version
			if rows, ok := result["rows"]; ok {
				queryResultsList[i]["rows"] = rows
			}
		}
		results["queryResults"] = queryResultsList
	}

	return results, nil
}

func (c *Quill) postQuill(path string, payload interface{}) (*models.QuillClientResponse, error) {
	// Serialize payload to JSON
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal payload: %w", err)
	}

	// Create POST request
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/sdk/%s", c.baseURL, path), bytes.NewReader(payloadBytes))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Add("Authorization", c.config.headers.authorization)
	req.Header.Set("Content-Type", "application/json")
	// Send request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()
	// Check status code
	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		var errorResp models.QuillClientResponse
		err := json.Unmarshal(bodyBytes, &errorResp)
		if err != nil {
			return nil, fmt.Errorf("unexpected error format: %s", string(bodyBytes))
		}
		return nil, fmt.Errorf("%s", *errorResp.Error)
	}

	// Decode response
	var response models.QuillClientResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("failed to decode response: %w. Response body: %s", err, string(bodyBytes))
	}

	return &response, nil
}

// close closes the target connection.
func (c *Quill) Close() error {
	if c.TargetConnection != nil {
		return c.TargetConnection.Close()
	}
	return nil
}
