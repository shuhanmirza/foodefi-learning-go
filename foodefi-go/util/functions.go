package util

import (
	"fmt"
)

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func IsValidEventFieldType(fieldType string) bool {
	return stringInSlice(fieldType, GetAllEventFieldTypes())
}

func ConvertFieldValueInterfaceToString(value interface{}, eventFieldType string) (string, error) {
	switch eventFieldType {
	case EventFieldString:
		switch value.(type) {
		case string:
			return value.(string), nil
		default:
			return "", &InvalidEventFieldType{}
		}
	case EventFieldNumber:
		switch value.(type) {
		case float64:
			return fmt.Sprintf("%g", value.(float64)), nil
		case int64:
			return fmt.Sprintf("%d", value.(int64)), nil
		default:
			return "", &InvalidEventFieldType{}
		}
	case EventFieldBoolean:
		switch value.(type) {
		case bool:
			return fmt.Sprintf("%t", value.(bool)), nil
		default:
			return "", &InvalidEventFieldType{}
		}
	}
	return "", &InvalidEventFieldType{}
}
