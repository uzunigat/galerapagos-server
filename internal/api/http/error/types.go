package httperror

type ApiError struct {
	HTTPStatus int    `json:"-"`
	Code       string `json:"code"`
	Message    string `json:"message"`
}

func NewApiError(httpStatus int, code, message string) *ApiError {
	return &ApiError{
		HTTPStatus: httpStatus,
		Code:       code,
		Message:    message,
	}
}

func (err *ApiError) Error() string {
	return err.Message
}
