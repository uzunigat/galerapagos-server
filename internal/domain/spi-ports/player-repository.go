package spiports

import (
	apiports "github.com/Audibene-GMBH/ta.go-hexagonal-skeletor/internal/domain/api-ports"
	"github.com/Audibene-GMBH/ta.go-hexagonal-skeletor/internal/domain/model"
)

type CreatePlayerRequest struct {
	apiports.CreatePlayerRequest
	Gid string
}

type UpdatePlayerRequest struct {
	apiports.UpdatePlayerRequest
}

type PlayerRepository interface {
	CreateOne(ctx model.Context, createPlayerRequest CreatePlayerRequest) (*model.Player, error)
	GetOne(ctx model.Context, gid string) (*model.Player, error)
	GetOneByEmail(ctx model.Context, email string) (*model.Player, error)
	GetMany(ctx model.Context, query apiports.GetManyPlayersQuery) ([]model.Player, model.ResponseMeta, error)
	UpdateOne(ctx model.Context, gid string, updatePlayerRequest UpdatePlayerRequest) (*model.Player, error)
}
