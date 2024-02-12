package env

import (
	"fmt"
	"os"
	"reflect"
	"strconv"

	"github.com/joho/godotenv"
)

func Load(filenames ...string) (err error) {
	err = godotenv.Load(filenames...)
	if err != nil {
		return fmt.Errorf("Could not load .env file: %s", err)
	}
	return
}

func Marshal(element interface{}) (err error) {
	value := reflect.ValueOf(element)

	if value.Kind() != reflect.Ptr || value.IsNil() {
		return fmt.Errorf("%s is not a pointer", value.Kind())
	}

	err = traverseElement(element, true)

	return err
}

func traverseElement(element interface{}, isFinalLoad bool) (err error) {
	elementValue := reflect.ValueOf(element).Elem()
	elementType := reflect.TypeOf(element).Elem()

	for i := 0; i < elementType.NumField(); i++ {
		field := elementType.Field(i)
		fieldValue := elementValue.Field(i)
		if field.Type.Kind() == reflect.Struct {
			err = traverseElement(fieldValue.Addr().Interface(), isFinalLoad)
		} else {
			err = processField(field, fieldValue, isFinalLoad)
		}
		if err != nil {
			return err
		}
	}
	return
}

func processField(field reflect.StructField, fieldValue reflect.Value, isFinalLoad bool) (err error) {

	if !fieldValue.IsZero() {
		return
	}

	envTag, ok := field.Tag.Lookup("env")
	if !ok {
		return
	}

	if !fieldValue.CanSet() {
		return fmt.Errorf("field '%s' cannot be set", field.Name)
	}

	envValue, ok := os.LookupEnv(envTag)
	if ok {
		fieldValue.Set(reflect.ValueOf(envValue))
		return
	}

	if !isFinalLoad {
		return
	}

	defaultValue, ok := field.Tag.Lookup("default")
	if ok {
		fieldValue.Set(reflect.ValueOf(defaultValue))
		return
	}

	optionalTag, ok := field.Tag.Lookup("optional")
	if !ok {
		return fmt.Errorf("Field '%s' is required, but value could not be determined.", field.Name)
	}

	isOptional, err := strconv.ParseBool(optionalTag)

	if err != nil {
		return fmt.Errorf("Optional tag for field '%s' can not be casted into a boolean.", field.Name)
	}
	if !isOptional {
		return fmt.Errorf("Field '%s' is required, but value could not be determined.", field.Name)
	}
	return
}
