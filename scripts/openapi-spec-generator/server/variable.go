package server

import "github.com/swaggest/openapi-go/openapi3"

func NewServerVariable(name string, description *string, defaultValue string, values []string) openapi3.ServerVariable {
	return openapi3.ServerVariable{
		Description: description,
		Default:     defaultValue,
		Enum:        values,
	}
}
