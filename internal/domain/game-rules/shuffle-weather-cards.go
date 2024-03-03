package gamerules

import (
	"math/rand"

	"github.com/Audibene-GMBH/ta.go-hexagonal-skeletor/internal/domain/model"
)

func ShuffleWeatherCards(weatherCards []model.WeatherCard) ([]model.WeatherCard, error) {

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
