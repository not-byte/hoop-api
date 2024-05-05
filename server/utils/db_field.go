package utils

import (
	"fmt"
	"reflect"
)

func FormatField(field reflect.Value) string {
	if field.Kind() == reflect.Ptr {
		if field.IsNil() {
			return "NULL"
		}
		field = field.Elem()
	}
	switch field.Kind() {
	case reflect.String:
		return fmt.Sprintf("'%s'", field.String())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return fmt.Sprintf("%d", field.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return fmt.Sprintf("%d", field.Uint())
	case reflect.Float32, reflect.Float64:
		return fmt.Sprintf("%f", field.Float())
	case reflect.Bool:
		return fmt.Sprintf("%t", field.Bool())
	default:
		return fmt.Sprintf("%v", field.Interface())
	}
}
