package dbutils

import (
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/uptrace/bun"
	"golang.org/x/exp/slices"
)

func MapStructToDBQueryWhere(dbQuery *bun.SelectQuery, queryModel interface{}) {
	mapStructToDBQueryWhereReflectStruct(dbQuery, queryModel)
}

func mapStructToDBQueryWhereReflectStruct(dbQuery *bun.SelectQuery, structToReflect interface{}) {
	reflectedStruct := reflect.TypeOf(structToReflect)
	reflectedValue := reflect.ValueOf(structToReflect)

	for i := 0; i < reflectedStruct.NumField(); i++ {
		fieldValue := reflectedValue.Field(i)
		fieldStruct := reflectedStruct.Field(i)

		if !ignoreField(fieldStruct) {
			if !fieldValue.IsZero() {
				if isFieldDeepStructure(fieldStruct, fieldValue) {
					if fieldValue.Type().Kind() == reflect.Ptr { // pointer
						mapStructToDBQueryWhereReflectStruct(dbQuery, fieldValue.Elem().Interface())
					} else {
						mapStructToDBQueryWhereReflectStruct(dbQuery, fieldValue.Interface())
					}
				} else {
					val := getFieldValueAsString(fieldValue)
					addtoQuery(dbQuery, fieldStruct, val)
				}
			}
		}
	}
}

func addtoQuery(dbQuery *bun.SelectQuery, fieldStruct reflect.StructField, searchValue string) {
	dbOperator := "="
	dbField := findTheMagicToConvertStructToDBField(fieldStruct.Name)

	dbTagContent, ok := fieldStruct.Tag.Lookup("db")
	if ok {
		splitDbTagContent := strings.Split(dbTagContent, ",")
		dbField = splitDbTagContent[0]
		if len(splitDbTagContent) > 1 {
			dbOperator = splitDbTagContent[1]
		}
	}

	if dbField == "-" {
		return
	}

	if slices.Contains([]string{"=", "IN"}, dbOperator) {
		searchValueArray := strings.Split(searchValue, ",")
		dbOperator = "="
		if len(searchValueArray) > 1 {
			dbOperator = "IN"
		}
	}

	if slices.Contains([]string{"IN", "NOT IN"}, dbOperator) {
		searchValueArray := strings.Split(searchValue, ",")
		dbQuery.Where(dbField+" "+dbOperator+" (?)", bun.In(searchValueArray))
	} else {
		dbQuery.Where(dbField+" "+dbOperator+" ?", searchValue)
	}
}

func ignoreField(fieldStruct reflect.StructField) bool {
	dbTagContent, ok := fieldStruct.Tag.Lookup("db")
	if ok {
		if strings.Split(dbTagContent, ",")[0] == "-" {
			return true
		}
	}
	return false
}

func isFieldDeepStructure(fieldStruct reflect.StructField, fieldValue reflect.Value) bool {
	if fieldValue.Type().Kind() == reflect.Ptr {
		if fieldStruct.Type.Elem().Kind() == reflect.Struct {
			if hasExportedFields(fieldValue.Elem().Interface()) {
				return true
			}
		}
	}

	if fieldStruct.Type.Kind() == reflect.Struct && hasExportedFields(fieldValue.Interface()) {
		return true
	}

	return false
}

func getFieldValueAsString(fieldValue reflect.Value) string {
	if fieldValue.Type().Kind() == reflect.Ptr { // pointer
		if fieldValue.Type().Elem().Kind() == reflect.Struct {
			return valueToString(fieldValue.Elem().Interface())
		}
		fieldValue = fieldValue.Elem()
	}

	if fieldValue.CanConvert(reflect.TypeOf("")) { // autoconvertion (does not work for Boolean)
		return valueToString(fieldValue)
	}

	if fieldValue.Type().Kind() == reflect.Struct { // structs without exported fields like time.Time
		return valueToString(fieldValue.Interface())
	}

	return valueToString(fieldValue.Interface())
}

func valueToString(value interface{}) string {
	if t, ok := value.(time.Time); ok {
		return t.Format(time.RFC3339)
	}
	return fmt.Sprintf("%v", value)
}

func hasExportedFields(structToReflect interface{}) bool {
	reflectedStruct := reflect.TypeOf(structToReflect)
	for i := 0; i < reflectedStruct.NumField(); i++ {
		if reflectedStruct.Field(i).PkgPath == "" {
			return true
		}
	}
	return false
}
