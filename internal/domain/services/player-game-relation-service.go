package services

import (
	"fmt"

	"github.com/Audibene-GMBH/ta.go-hexagonal-skeletor/internal/domain/model"
	spiports "github.com/Audibene-GMBH/ta.go-hexagonal-skeletor/internal/domain/spi-ports"
	"github.com/Audibene-GMBH/ta.go-hexagonal-skeletor/internal/errors"
)

type PlayerGameRelationService interface {
	JoinGame(ctx model.Context, playerGid string, gameGid string) (*model.PlayerGameRelation, error)
}
type playerGameRelationService struct {
	repository   spiports.PlayerGameRelationRepository
	gidGenerator model.GidGenerator
}

func NewPlayerGameRelationService(repository spiports.PlayerGameRelationRepository, gidGenerator model.GidGenerator) PlayerGameRelationService {
	return &playerGameRelationService{repository: repository, gidGenerator: gidGenerator}
}

func (service *playerGameRelationService) JoinGame(ctx model.Context, playerGid string, gameGid string) (*model.PlayerGameRelation, error) {

	player, err := service.repository.GetOne(ctx, playerGid, gameGid)

	if err != nil {

		fmt.Printf("Error: %v\n", err)

		if err.(*errors.CustomError).Code == "PLAYER_GAME_RELATION_NOT_FOUND" {
			relation := spiports.CreateOnePlayerGameRelationRequest{
				Gid:       service.gidGenerator.Generate(),
				PlayerGid: playerGid,
				GameGid:   gameGid,
				Status:    model.PlayerGameRelationConnected,
			}

			return service.repository.CreateOne(ctx, relation)
		}

		return nil, err

	}

	return player, nil
}
