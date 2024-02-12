package utils

import (
	"fmt"

	"github.com/Audibene-GMBH/ta.go-hexagonal-skeletor/internal/errors"
	"github.com/go-playground/validator/v10"
)

func FormatValidationErrors(errs error) *errors.CustomError {
	if _, ok := errs.(*validator.InvalidValidationError); ok {
		return errors.NewInvalidPayloadError("INVALID_PAYLOAD", errs)
	}

	err := errs.(validator.ValidationErrors)[0]

	return errors.NewInvalidPayloadError("INVALID_PAYLOAD", fmt.Errorf("Value '%s' for attribute '%s' is not of type: %s", err.Value(), err.Field(), err.Tag()))
}
