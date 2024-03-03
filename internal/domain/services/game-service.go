package services

import (
	apiports "github.com/Audibene-GMBH/ta.go-hexagonal-skeletor/internal/domain/api-ports"
	"github.com/Audibene-GMBH/ta.go-hexagonal-skeletor/internal/domain/model"
	spiports "github.com/Audibene-GMBH/ta.go-hexagonal-skeletor/internal/domain/spi-ports"
	"github.com/Audibene-GMBH/ta.go-hexagonal-skeletor/internal/spi/repositories/redis"
)

type GameService interface {
	GetOne(ctx model.Context, gid string) (*model.Game, error)
	GetMany(ctx model.Context, query apiports.GetManyGamesQuery) ([]model.Game, model.ResponseMeta, error)
	CreateNewGame(ctx model.Context, createGameRequest apiports.CreateGameRequest) (*model.Game, error)
	UpdateOne(ctx model.Context, gid string, updatePlayerRequest apiports.UpdateGameRequest) (*model.Game, error)
	PlayerJoin(ctx model.Context, gid string, playerGid string) (*model.PlayerGameRelation, error)
	Start(ctx model.Context, gameGid string, startGameRequest apiports.StartGameRequest) (*model.Game, error)
	End(ctx model.Context, gameGid string) (*model.Game, error)
}

type GameServices struct {
	PlayerGameRelation PlayerGameRelationService
}

type gameService struct {
	services     GameServices
	repository   spiports.GameRepository
	validator    model.ServiceValidator
	gidGenerator model.GidGenerator
	publisher    spiports.Publisher
}

func NewGameService(repository spiports.GameRepository, utils model.DomainUtils, services GameServices, redisClient *redis.RedisClient, publisher spiports.Publisher) GameService {
	return &gameService{
		repository:   repository,
		validator:    utils.Validator,
		gidGenerator: utils.GidGenerator,
		services:     services,
		publisher:    publisher,
	}
}

func (service *gameService) GetOne(ctx model.Context, gid string) (*model.Game, error) {
	return service.repository.GetOne(ctx, gid)
}

func (service *gameService) GetMany(ctx model.Context, query apiports.GetManyGamesQuery) ([]model.Game, model.ResponseMeta, error) {
	validationErrors := service.validator.ValidateStruct(query)
	if validationErrors != nil {
		return []model.Game{}, model.ResponseMeta{}, validationErrors
	}

	return service.repository.GetMany(ctx, query)
}

func (service *gameService) CreateNewGame(ctx model.Context, createGameRequest apiports.CreateGameRequest) (*model.Game, error) {
	validationErrors := service.validator.ValidateStruct(createGameRequest)
	if validationErrors != nil {
		return nil, validationErrors
	}

	createGameRequest.Gid = service.gidGenerator.GenerateIfEmpty(createGameRequest.Gid)
	spiCreateGameRequest := spiports.CreateGameRequest{
		CreateGameRequest: createGameRequest,
		Status:            model.GameCreated,
	}

	createdGame, err := service.repository.CreateNewGame(ctx, spiCreateGameRequest)

	if err != nil {
		return nil, err
	}

	return createdGame, nil
}

func (service *gameService) UpdateOne(ctx model.Context, gid string, updateGameRequest apiports.UpdateGameRequest) (*model.Game, error) {
	validationErrors := service.validator.ValidateStruct(updateGameRequest)
	if validationErrors != nil {
		return nil, validationErrors
	}

	spiUpdateGameRequest := spiports.UpdateGameRequest{
		UpdateGameRequest: updateGameRequest,
	}

	updatedGame, err := service.repository.UpdateOne(ctx, gid, spiUpdateGameRequest)

	if err != nil {
		return nil, err
	}

	return updatedGame, nil

}

func (service *gameService) PlayerJoin(ctx model.Context, gid string, playerGid string) (*model.PlayerGameRelation, error) {
	return service.services.PlayerGameRelation.JoinGame(ctx, playerGid, gid)
}

func (service *gameService) Start(ctx model.Context, gameGid string, startGameRequest apiports.StartGameRequest) (*model.Game, error) {

	spiRequest := spiports.StartGameRequest{
		Status:         model.GameStatus(*startGameRequest.Status),
		FoodResources:  startGameRequest.FoodResources,
		WaterResources: startGameRequest.WaterResources,
		WeatherCards:   startGameRequest.WeatherCards,
		WreckCardGids:  startGameRequest.WreckCardGids,
		PlayerTurns:    startGameRequest.PlayerTurns,
	}

	game, error := service.repository.Start(ctx, gameGid, spiRequest)

	if error != nil {
		return nil, error
	}

	service.publisher.PublishGameStarted(game.Gid)

	return game, nil
}

func (service *gameService) End(ctx model.Context, gameGid string) (*model.Game, error) {

	spiEndRequest := spiports.EndGameRequest{
		Status: model.GameFinished,
	}

	return service.repository.End(ctx, gameGid, spiEndRequest)
}
