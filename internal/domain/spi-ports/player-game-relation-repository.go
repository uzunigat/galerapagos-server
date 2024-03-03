package spiports

import "github.com/Audibene-GMBH/ta.go-hexagonal-skeletor/internal/domain/model"

type CreateOnePlayerGameRelationRequest struct {
	Gid       string
	PlayerGid string
	GameGid   string
	Status    model.PlayerGameRelationStatus
}

type UpdateOnePlayerGameRelationRequest struct {
	Status *model.PlayerGameRelationStatus
}

type PlayerGameRelationRepository interface {
	GetOne(ctx model.Context, playerGid string, gameGid string) (*model.PlayerGameRelation, error)
	CreateOne(ctx model.Context, request CreateOnePlayerGameRelationRequest) (*model.PlayerGameRelation, error)
	UpdateOne(ctx model.Context, gid string, status UpdateOnePlayerGameRelationRequest) (*model.PlayerGameRelation, error)
	GetPlayersByGameGid(ctx model.Context, gameGid string) ([]model.Player, error)
}
