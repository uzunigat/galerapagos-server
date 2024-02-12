package postgres

import "github.com/Audibene-GMBH/ta.go-hexagonal-skeletor/internal/errors"

func NewUnkownDatabaseError(err error) *errors.CustomError {
	return errors.NewInternalServerError("UNKNOWN_DATABASE_ERROR", err)
}

func NewBeeNotFoundError(err error) *errors.CustomError {
	return errors.NewRecordNotFoundError("BEE_NOT_FOUND", err)
}
