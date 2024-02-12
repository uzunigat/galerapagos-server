package api

import (
	"net/http"

	"github.com/Audibene-GMBH/ta.go-hexagonal-skeletor/internal/errors"
)

func NewUnauthorizedError(err error) *errors.CustomError {
	return errors.NewCustomError(http.StatusUnauthorized, "UNAUTHORIZED", err)
}

func NewForbiddenError(err error) *errors.CustomError {
	return errors.NewCustomError(http.StatusForbidden, "FORBIDDEN", err)
}
