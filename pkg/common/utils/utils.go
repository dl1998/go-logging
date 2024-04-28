// Package utils contains utility functions.
package utils

import (
	"fmt"
	"net/http"
	"reflect"
)

// StructToMap converts a struct public fields to a map.
func StructToMap(object interface{}) map[string]interface{} {
	objectType := reflect.TypeOf(object)
	objectValue := reflect.ValueOf(object)

	var data = make(map[string]interface{})
	for index := 0; index < objectType.NumField(); index++ {
		field := objectType.Field(index)
		if field.PkgPath == "" {
			value := objectValue.Field(index).Interface()
			switch convertedValue := value.(type) {
			case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64, bool, string:
				data[field.Name] = value
			case http.Header:
				for key, value := range convertedValue {
					fieldName := fmt.Sprintf("Header.%s", key)
					data[fieldName] = value
				}
			case map[string]interface{}, []interface{}:
			default:
				data[field.Name] = fmt.Sprintf("%v", value)
			}
		}
	}

	return data
}
