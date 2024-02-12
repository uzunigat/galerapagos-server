package services

import (
	apiports "github.com/Audibene-GMBH/ta.go-hexagonal-skeletor/internal/domain/api-ports"
	"github.com/Audibene-GMBH/ta.go-hexagonal-skeletor/internal/domain/model"
	spiports "github.com/Audibene-GMBH/ta.go-hexagonal-skeletor/internal/domain/spi-ports"
)

type PlayerService interface {
	GetOne(ctx model.Context, gid string) (*model.Player, error)
	GetMany(ctx model.Context, query apiports.GetManyPlayersQuery) ([]model.Player, model.ResponseMeta, error)
	CreateOne(ctx model.Context, createBeeRequest apiports.CreatePlayerRequest) (*model.Player, error)
	UpdateOne(ctx model.Context, gid string, updateBeeRequest apiports.UpdatePlayerRequest) (*model.Player, error)
}

type playerService struct {
	repository   spiports.PlayerRepository
	validator    model.ServiceValidator
	gidGenerator model.GidGenerator
}

func NewPlayerService(repository spiports.PlayerRepository, utils model.DomainUtils) PlayerService {
	return &playerService{
		repository:   repository,
		validator:    utils.Validator,
		gidGenerator: utils.GidGenerator,
	}
}

func (service *playerService) GetOne(ctx model.Context, gid string) (*model.Player, error) {
	return service.repository.GetOne(ctx, gid)
}

func (service *playerService) GetMany(ctx model.Context, query apiports.GetManyPlayersQuery) ([]model.Player, model.ResponseMeta, error) {
	validationErrors := service.validator.ValidateStruct(query)
	if validationErrors != nil {
		return []model.Player{}, model.ResponseMeta{}, validationErrors
	}

	return service.repository.GetMany(ctx, query)
}

func (service *playerService) CreateOne(ctx model.Context, createPlayerRequest apiports.CreatePlayerRequest) (*model.Player, error) {
	validationErrors := service.validator.ValidateStruct(createPlayerRequest)
	if validationErrors != nil {
		return nil, validationErrors
	}

	createPlayerRequest.Gid = service.gidGenerator.GenerateIfEmpty(createPlayerRequest.Gid)
	spiCreateBeeRequest := spiports.CreatePlayerRequest{
		CreatePlayerRequest: createPlayerRequest,
		Gid:                 *createPlayerRequest.Gid,
	}

	createdBee, err := service.repository.CreateOne(ctx, spiCreateBeeRequest)

	if err != nil {
		return nil, err
	}

	return createdBee, nil
}

func (service *playerService) UpdateOne(ctx model.Context, gid string, updatePlayerRequest apiports.UpdatePlayerRequest) (*model.Player, error) {
	validationErrors := service.validator.ValidateStruct(updatePlayerRequest)
	if validationErrors != nil {
		return nil, validationErrors
	}

	spiUpdatePlayerRequest := spiports.UpdatePlayerRequest{
		UpdatePlayerRequest: updatePlayerRequest,
	}

	updatedBee, err := service.repository.UpdateOne(ctx, gid, spiUpdatePlayerRequest)

	if err != nil {
		return nil, err
	}

	return updatedBee, nil

}
