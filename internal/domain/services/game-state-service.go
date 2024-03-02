package services

import (
	"fmt"
	"math/rand"

	"github.com/Audibene-GMBH/ta.go-hexagonal-skeletor/internal/domain/model"
)

type GameStateService interface {
	SetupGame(ctx model.Context, numberOfPlayers int) (*model.Game, error)
}

type GameStateServices struct {
	WeatherService WeatherService
	GameService    GameService
}

type gameStateService struct {
	services GameStateServices
}

func NewGameStateService(services GameStateServices) GameStateService {
	return &gameStateService{services: services}
}

func (g *gameStateService) SetupGame(ctx model.Context, numberOfPlayers int) (*model.Game, error) {
	weatherCards, err := g.services.WeatherService.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	pickedWeatherCardsCards, err := g.pickWeatherCards(weatherCards)
	if err != nil {
		return nil, err
	}

	foodResources, err := g.getFoodResources(numberOfPlayers)

	if err != nil {
		return nil, err
	}

	waterResources, err := g.getWaterResources(numberOfPlayers)

	if err != nil {
		return nil, err
	}

	game := &model.Game{
		WeatherCards:   pickedWeatherCardsCards,
		FoodResources:  foodResources,
		WaterResources: waterResources,
	}

	return game, nil
}

func (g *gameStateService) pickWeatherCards(weatherCards []model.WeatherCard) ([]model.WeatherCard, error) {

	cards := weatherCards
	pickedCards := make([]model.WeatherCard, 0)
	limitPickBeforeFinalCard := 5

	for i := 0; i < limitPickBeforeFinalCard; i++ {
		index := rand.Intn(len(cards))

		if cards[index].IsFinalGame {
			i--
		}

		if cards[index].Quantity > 0 && !cards[index].IsFinalGame {
			pickedCards = append(pickedCards, cards[index])
			cards[index].Quantity--

			if cards[index].Quantity == 0 {
				cards = append(cards[:index], cards[index+1:]...)
			}
		}
	}

	for i := 6; i < 12; i++ {
		index := rand.Intn(len(cards))

		if cards[index].Quantity > 0 {
			pickedCards = append(pickedCards, cards[index])
			cards[index].Quantity--

			if cards[index].Quantity == 0 {
				cards = append(cards[:index], cards[index+1:]...)
			}
		}
	}

	return pickedCards, nil
}

func (g *gameStateService) getFoodResources(numberOfPlayers int) (int, error) {

	if numberOfPlayers < 3 {
		return 0, fmt.Errorf("not enough players to start the game")
	}

	switch numberOfPlayers {
	case 3:
		return 5, nil
	case 4:
		return 7, nil
	case 5:
		return 8, nil
	case 6:
		return 10, nil
	case 7:
		return 12, nil
	case 8:
		return 13, nil
	case 9:
		return 15, nil
	case 10:
		return 16, nil
	case 11:
		return 18, nil
	case 12:
		return 20, nil
	default:
		return 0, fmt.Errorf("something is wrong with the number of players")
	}
}

func (g *gameStateService) getWaterResources(numberOfPlayers int) (int, error) {

	if numberOfPlayers < 3 {
		return 0, fmt.Errorf("not enough players to start the game")
	}

	switch numberOfPlayers {
	case 3:
		return 6, nil
	case 4:
		return 8, nil
	case 5:
		return 10, nil
	case 6:
		return 12, nil
	case 7:
		return 14, nil
	case 8:
		return 16, nil
	case 9:
		return 18, nil
	case 10:
		return 20, nil
	case 11:
		return 22, nil
	case 12:
		return 24, nil
	default:
		return 0, fmt.Errorf("something is wrong with the number of players")
	}
}
