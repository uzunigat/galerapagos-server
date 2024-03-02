package postgres

import "github.com/Audibene-GMBH/ta.go-hexagonal-skeletor/internal/errors"

func NewUnkownDatabaseError(err error) *errors.CustomError {
	return errors.NewInternalServerError("UNKNOWN_DATABASE_ERROR", err)
}

func NewGameNotFoundError(err error) *errors.CustomError {
	return errors.NewRecordNotFoundError("GAME_NOT_FOUND", err)
}

func NewPlayerNotFoundError(err error) *errors.CustomError {
	return errors.NewRecordNotFoundError("PLAYER_NOT_FOUND", err)
}

func NewPlayerGameRelationNotFoundError(err error) *errors.CustomError {
	return errors.NewRecordNotFoundError("PLAYER_GAME_RELATION_NOT_FOUND", err)
}
