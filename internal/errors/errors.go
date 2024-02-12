package errors

import (
	"net/http"
)

// RequestError holds the message string and http code
// TODO: Make the CustomError immutable
type CustomError struct {
	HTTPStatus int    `json:"-"`
	Code       string `json:"code"`
	Message    string `json:"message"`
	Retryable  bool   `json:"-"`
}

func (customError *CustomError) Error() string {
	return customError.Message
}

func (customError *CustomError) IsRetryable() bool {
	return customError.Retryable
}

func NewCustomError(httpStatus int, code string, err error) *CustomError {
	return &CustomError{
		HTTPStatus: httpStatus,
		Code:       code,
		Message:    err.Error(),
		Retryable:  true,
	}
}

func NewInternalServerError(code string, err error) *CustomError {
	return &CustomError{
		HTTPStatus: http.StatusInternalServerError,
		Code:       code,
		Message:    err.Error(),
		Retryable:  false,
	}
}

func NewRecordNotFoundError(code string, err error) *CustomError {
	return &CustomError{
		HTTPStatus: http.StatusNotFound,
		Code:       code,
		Message:    err.Error(),
		Retryable:  true,
	}
}

func NewBadRequest(code string, err error) *CustomError {
	return &CustomError{
		HTTPStatus: http.StatusBadRequest,
		Code:       code,
		Message:    err.Error(),
		Retryable:  false,
	}
}

func NewInvalidPayloadError(code string, err error) *CustomError {
	return &CustomError{
		HTTPStatus: http.StatusBadRequest,
		Code:       code,
		Message:    err.Error(),
		Retryable:  false,
	}
}
