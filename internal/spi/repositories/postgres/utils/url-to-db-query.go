package dbutils

import (
	"reflect"
	"strings"

	"github.com/uptrace/bun"
)

func UrlToDbQuery(dbQuery *bun.SelectQuery, urlQuery interface{}) error {
	queryValue := reflect.ValueOf(urlQuery)
	queryField := reflect.TypeOf(urlQuery)

	for i := 0; i < queryField.NumField(); i++ {
		field := queryField.Field(i)
		fieldValue := queryValue.Field(i)
		if !fieldValue.IsZero() {
			dbTag, ok := field.Tag.Lookup("db")
			if ok {
				if field.Type != reflect.TypeOf("") {
					dbQuery.Where(dbTag+" = ?", fieldValue.Interface())
					return nil
				}
				queryValueString := fieldValue.Interface().(string)
				if strings.Contains(queryValueString, ",") {
					queryValueArray := strings.Split(queryValueString, ",")
					dbQuery.Where(dbTag+" IN (?)", bun.In(queryValueArray))
				} else {
					dbQuery.Where(dbTag+" = ?", queryValueString)
				}
			}
		}
	}

	return nil
}
