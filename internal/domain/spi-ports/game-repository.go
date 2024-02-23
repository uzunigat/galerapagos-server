package spiports

import (
	apiports "github.com/Audibene-GMBH/ta.go-hexagonal-skeletor/internal/domain/api-ports"
	"github.com/Audibene-GMBH/ta.go-hexagonal-skeletor/internal/domain/model"
)

type CreateGameRequest struct {
	apiports.CreateGameRequest
	Status model.GameStatus
}

type UpdateGameRequest struct {
	apiports.UpdateGameRequest
}

type GameRepository interface {
	CreateNewGame(ctx model.Context, createGameRequest CreateGameRequest) (*model.Game, error)
	GetOne(ctx model.Context, gid string) (*model.Game, error)
	GetMany(ctx model.Context, query apiports.GetManyGamesQuery) ([]model.Game, model.ResponseMeta, error)
	UpdateOne(ctx model.Context, gid string, updateGameRequest UpdateGameRequest) (*model.Game, error)
}
