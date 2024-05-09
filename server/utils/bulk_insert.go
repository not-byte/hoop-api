package utils

import (
	"fmt"
	"reflect"
	"strings"
)

/*
This function leverages the assumption that JSON property names directly correspond to database column names. For example:

	type Player struct {
		FirstName *string `json:"first_name"`
		LastName  *string `json:"last_name"`
		Age       *int    `json:"age"`
		Number    *int    `json:"number"`
	}

Under this assumption, the struct's JSON tags ('first_name', 'last_name', 'age', 'number') are used as the column names in the database. This approach significantly enhances automation by eliminating the need for manually specifying column names when creating bulk inserts. It is a strong assumption that simplifies integration but requires consistent naming conventions between JSON representations and database schemas.

I know this exposes column names through API but it is so nicely automated that i do not mind for this particular app
*/

func BulkInsert[T any](unsavedRows []*T, tableName string) (string, error) {

	slice := reflect.ValueOf(unsavedRows)
	if slice.Kind() != reflect.Slice {
		return "", fmt.Errorf("unsavedRows must be a slice")
	}

	if slice.Len() == 0 {
		return "", fmt.Errorf("unsavedRows slice is empty")
	}

	firstRow := slice.Index(0).Elem()
	if firstRow.Kind() != reflect.Struct {
		return "", fmt.Errorf("element in the slice is not a struct")
	}

	columnCount := firstRow.NumField()
	valStrings := make([]string, 0, slice.Len())
	columnNames := make([]string, 0, columnCount)
	t := reflect.TypeOf(*unsavedRows[0])
	for i := 0; i < t.NumField(); i++ {
		columnNames = append(columnNames, t.Field(i).Tag.Get("json"))
	}

	for i := 0; i < slice.Len(); i++ {
		el := slice.Index(i).Elem()
		if el.Kind() != reflect.Struct {
			return "", fmt.Errorf("element %d in the slice is not a struct", i)
		}

		placeholders := make([]string, columnCount)
		for j := 0; j < columnCount; j++ {
			field := el.Field(j)
			if !field.CanInterface() {
				return "", fmt.Errorf("cannot interface field %d of element %d", j, i)
			}

			placeholders[j] = FormatField(field)
		}
		valStrings = append(valStrings, fmt.Sprintf("(%s)", strings.Join(placeholders, ", ")))
	}

	columns := fmt.Sprintf("(%s)", strings.Join(columnNames, ", "))
	stmt := fmt.Sprintf("INSERT INTO %s %s VALUES %s", tableName, columns, strings.Join(valStrings, ", "))
	return stmt, nil
}

func BulkUpdate[T any](updatedRows []*T, tableName string, condition string) (string, error) {
	slice := reflect.ValueOf(updatedRows)
	if slice.Kind() != reflect.Slice {
		return "", fmt.Errorf("updatedRows must be a slice")
	}

	if slice.Len() == 0 {
		return "", fmt.Errorf("updatedRows slice is empty")
	}

	firstRow := slice.Index(0).Elem()
	if firstRow.Kind() != reflect.Struct {
		return "", fmt.Errorf("element in the slice is not a struct")
	}

	columnCount := firstRow.NumField()
	columnNames := make([]string, 0, columnCount)
	t := reflect.TypeOf(*updatedRows[0])
	for i := 0; i < t.NumField(); i++ {
		columnNames = append(columnNames, t.Field(i).Tag.Get("json"))
	}

	valStrings := make([]string, 0, slice.Len())
	for i := 0; i < slice.Len(); i++ {
		el := slice.Index(i).Elem()
		if el.Kind() != reflect.Struct {
			return "", fmt.Errorf("element %d in the slice is not a struct", i)
		}

		placeholders := make([]string, columnCount)
		for j := 0; j < columnCount; j++ {
			field := el.Field(j)
			if !field.CanInterface() {
				return "", fmt.Errorf("cannot interface field %d of element %d", j, i)
			}

			placeholders[j] = fmt.Sprintf("%s = %s", columnNames[j], FormatField(field))
		}
		valStrings = append(valStrings, strings.Join(placeholders, ", "))
	}

	stmt := fmt.Sprintf("UPDATE %s SET %s WHERE %s", tableName, strings.Join(valStrings, ", "), condition)
	return stmt, nil
}
