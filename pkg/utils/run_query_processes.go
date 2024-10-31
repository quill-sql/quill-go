package utils

import (
	"encoding/json"
	"fmt"
)

type TableSchemaInfo struct {
	FieldType   string  `json:"fieldType"`
	Name        string  `json:"name"`
	DisplayName string  `json:"displayName"`
	IsVisible   bool    `json:"isVisible"`
	Field       *string `json:"field,omitempty"`
	DataTypeID  *int    `json:"dataTypeID,omitempty"`
}

func RemoveFields(queryResults map[string]interface{}, fieldsToRemove []string) map[string]interface{} {
	var filteredFields []interface{}
	var filteredRows []interface{}
	if fields, ok := queryResults["fields"].([]interface{}); ok {
		// Filter the fields
		for _, field := range fields {
			fieldMap := field.(map[string]interface{})
			name := fieldMap["name"].(string)
			if !contains(fieldsToRemove, name) {
				filteredFields = append(filteredFields, field)
			}
		}
	}

	if rows, ok := queryResults["rows"].([]interface{}); ok {
		filteredRows = make([]interface{}, len(rows))
		for i, row := range rows {
			rowMap := row.(map[string]interface{})
			for _, field := range fieldsToRemove {
				delete(rowMap, field)
			}
			filteredRows[i] = rowMap
		}
	}

	return map[string]interface{}{
		"fields": filteredFields,
		"rows":   filteredRows,
	}
}

// Helper function to check if a slice contains a string
func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func StructToMap(obj interface{}) (map[string]interface{}, error) {
	var result map[string]interface{}
	data, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

func MapToStruct(m map[string]interface{}, obj interface{}) error {
	// Convert the map to JSON
	data, err := json.Marshal(m)
	if err != nil {
		return err
	}

	// Unmarshal the JSON into the provided struct
	err = json.Unmarshal(data, &obj)
	return err
}

func ConvertMapToStructList[T any](maps []map[string]interface{}) ([]T, error) {
	var structs []T
	for _, m := range maps {
		data, err := json.Marshal(m)
		if err != nil {
			return nil, err
		}

		var obj T
		err = json.Unmarshal(data, &obj)
		if err != nil {
			return nil, err
		}

		structs = append(structs, obj)
	}
	return structs, nil
}

func ConvertInterfaceToStringSlice(queries []interface{}) ([]string, error) {
	stringQueries := make([]string, len(queries))
	for i, query := range queries {
		str, ok := query.(string)
		if !ok {
			return nil, fmt.Errorf("query at index %d is not a string", i)
		}
		stringQueries[i] = str
	}
	return stringQueries, nil
}

func ConvertStringToInterfaceSlice(strs []string) []interface{} {
	result := make([]interface{}, len(strs))
	for i, v := range strs {
		result[i] = v
	}
	return result
}
