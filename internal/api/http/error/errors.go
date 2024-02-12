package httperror

import (
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
)

type HttpErrorHandler interface {
	Handle(ctx *gin.Context, err error)
}

type httpErrorHandler struct{}

func NewHttpErrorHandler() HttpErrorHandler {
	return &httpErrorHandler{}
}

func (httpError *httpErrorHandler) Handle(ctx *gin.Context, err error) {
	if err == nil {
		return
	}

	apiError := httpError.mapErrorToApiError(err)
	ctx.JSON(apiError.HTTPStatus, apiError)
}

func (httpError *httpErrorHandler) mapErrorToApiError(err error) ApiError {
	reflectedStruct := reflect.TypeOf(err)
	reflectedValue := reflect.ValueOf(err)

	if reflectedStruct.Kind() == reflect.Ptr {
		reflectedStruct = reflectedStruct.Elem()
		reflectedValue = reflectedValue.Elem()
	}

	apiError := NewApiError(http.StatusInternalServerError, "INTERNAL_SERVER_ERROR", err.Error())

	/*
		it takes the HTTPStatuts from DomainError if it exist
		the domain should not know anything about http (its against hexacon)
		best would be to have a "type" and map it to httpstatus
	*/
	if httpStatusField, ok := reflectedStruct.FieldByName("HTTPStatus"); ok {
		if httpStatusField.Type.Kind() == reflect.Int {
			apiError.HTTPStatus = int(reflectedValue.FieldByName("HTTPStatus").Int())
		}
	}

	if _, ok := reflectedStruct.FieldByName("Code"); ok {
		apiError.Code = reflectedValue.FieldByName("Code").String()
	}

	return *apiError
}
