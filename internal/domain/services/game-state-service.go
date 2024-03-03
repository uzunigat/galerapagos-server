package services

import (
	gamerules "github.com/Audibene-GMBH/ta.go-hexagonal-skeletor/internal/domain/game-rules"
	"github.com/Audibene-GMBH/ta.go-hexagonal-skeletor/internal/domain/model"
)

type GameStateService interface {
	SetupGame(ctx model.Context, players []model.Player) (*model.Game, error)
}

type GameStateServices struct {
	Weather WeatherService
	Game    GameService
	Wreck   WreckService
}

type gameStateService struct {
	services GameStateServices
}

func NewGameStateService(services GameStateServices) GameStateService {
	return &gameStateService{services: services}
}

func (g *gameStateService) SetupGame(ctx model.Context, players []model.Player) (*model.Game, error) {
	weatherCards, err := g.services.Weather.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	wreckCards, err := g.services.Wreck.GetAll(ctx)

	if err != nil {
		return nil, err
	}

	wreckCardGids := make([]string, 0)

	for _, card := range wreckCards {
		wreckCardGids = append(wreckCardGids, card.Gid)
	}

	pickedWeatherCardsCards, err := gamerules.ShuffleWeatherCards(weatherCards)

	if err != nil {
		return nil, err
	}

	foodResources, err := gamerules.GetFoodResources(len(players))

	if err != nil {
		return nil, err
	}

	waterResources, err := gamerules.GetWaterResources(len(players))

	if err != nil {
		return nil, err
	}

	game := &model.Game{
		WeatherCards:   pickedWeatherCardsCards,
		FoodResources:  foodResources,
		WaterResources: waterResources,
		WreckCardGids:  wreckCardGids,
		PlayerTurns:    players,
	}

	return game, nil
}
