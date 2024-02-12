package model

import (
	"github.com/Audibene-GMBH/ta.go-hexagonal-skeletor/internal/errors"

	"github.com/swaggest/openapi-go/openapi3"
)

type BaseOperationRequest struct {
	Reflector *openapi3.Reflector
	Security  map[string][]string
	Tags      []string
}

type OperationRequest struct {
	Operation            *openapi3.Operation
	Method               string
	Model                interface{}
	Summary              string
	Query                interface{}
	BaseOperationRequest BaseOperationRequest
	Path                 string
	ErrorResponse        *errors.CustomError
}
