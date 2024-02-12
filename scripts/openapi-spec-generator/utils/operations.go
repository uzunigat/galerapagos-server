package utils

import (
	"net/http"

	model "github.com/Audibene-GMBH/ta.go-hexagonal-skeletor/scripts/openapi-spec-generator/model"
)

func AddOperation(or model.OperationRequest) {

	if or.BaseOperationRequest.Tags != nil {
		or.Operation.WithTags(or.BaseOperationRequest.Tags...)
	}
	or.Operation.WithSummary(or.Summary)
	or.Operation.WithDescription(or.Summary)

	handleError(or.BaseOperationRequest.Reflector.SetRequest(or.Operation, or.Query, or.Method))
	handleError(or.BaseOperationRequest.Reflector.SetJSONResponse(or.Operation, or.Model, http.StatusOK))
	if or.ErrorResponse != nil {
		handleError(or.BaseOperationRequest.Reflector.SetJSONResponse(or.Operation, or.ErrorResponse.Error, http.StatusNotFound))
	}
	handleError(or.BaseOperationRequest.Reflector.Spec.AddOperation(or.Method, or.Path, *or.Operation))
}
