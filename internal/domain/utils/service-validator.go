package domainutils

import (
	"fmt"

	"github.com/Audibene-GMBH/ta.go-hexagonal-skeletor/internal/errors"
	validator "github.com/go-playground/validator/v10"
)

type ServiceValidator struct {
	validator *validator.Validate
}

func NewServiceValidator() *ServiceValidator {
	return &ServiceValidator{
		validator: validator.New(),
	}
}

func (v *ServiceValidator) ValidateStruct(s interface{}) error {
	err := v.validator.Struct(s)
	if err != nil {
		return v.formatValidationErrors(err)
	}
	return nil
}

func (v *ServiceValidator) formatValidationErrors(errs error) *errors.CustomError {
	if _, ok := errs.(*validator.InvalidValidationError); ok {
		return errors.NewInvalidPayloadError("INVALID_PAYLOAD", errs)
	}

	err := errs.(validator.ValidationErrors)[0]

	return errors.NewInvalidPayloadError("INVALID_PAYLOAD", fmt.Errorf("value '%s' for attribute '%s' is not of type: %s", err.Value(), err.Field(), err.Tag()))
}
