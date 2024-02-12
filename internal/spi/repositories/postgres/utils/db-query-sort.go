package dbutils

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/Audibene-GMBH/ta.go-hexagonal-skeletor/internal/domain/model"
	"github.com/uptrace/bun"
)

/* you need to write your own function if Struct and DB is different or dont follow "camel case" convertion */
func DbQuerySort(dbQuery *bun.SelectQuery, sortOnStruct interface{}, sorting model.Sorting) error {
	if sorting.SortBy == nil {
		return nil
	}

	if sorting.Sort == nil {
		defaultSortDirection := model.SortDirectionAsc
		sorting.Sort = &defaultSortDirection
	}

	sortByField, err := GetSortByField(*sorting.SortBy, sortOnStruct)
	if err != nil {
		return err
	}
	orderExpression := fmt.Sprintf("%s %s", findTheMagicToConvertStructToDBField(sortByField.Name), string(*sorting.Sort))
	dbQuery.Order(orderExpression)
	return nil
}

func findTheMagicToConvertStructToDBField(camelCase string) string {
	snakeCase := ""
	for i, char := range camelCase {
		if i > 0 && char >= 'A' && char <= 'Z' {
			snakeCase += "_"
		}
		snakeCase += string(char)
	}
	return strings.ToLower(snakeCase)
}

func GetSortByField(sortByFieldJsonFieldName string, sortOnStruct interface{}) (*reflect.StructField, error) {
	reflectedStruct := reflect.TypeOf(sortOnStruct).Elem()

	for i := 0; i < reflectedStruct.NumField(); i++ {
		field := reflectedStruct.Field(i)

		_, isSortable := field.Tag.Lookup("sortable")
		if isSortable {
			jsonName := field.Tag.Get("json")
			if jsonName == sortByFieldJsonFieldName {
				return &field, nil
			}
		}
	}

	return nil, fmt.Errorf("Sortable field with JSON name %s not found", sortByFieldJsonFieldName)
}
