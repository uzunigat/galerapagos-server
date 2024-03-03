package apiports

import (
	"github.com/Audibene-GMBH/ta.go-hexagonal-skeletor/internal/domain/model"
)

type CreateGameRequest struct {
	Gid *string `json:"gid,omitempty" form:"gid"`
}

type StartGameRequest struct {
	Status         *model.GameStatus   `json:"status,omitempty" form:"status"`
	FoodResources  int                 `json:"food_resources" form:"food_resources"`
	WaterResources int                 `json:"water_resources" form:"water_resources"`
	WeatherCards   []model.WeatherCard `json:"weather_cards" form:"weather_cards"`
	WreckCardGids  []string            `json:"wreck_card_gids" form:"wreck_card_gids"`
	PlayerTurns    []model.Player      `json:"player_order_gids" form:"player_order_gids"`
}

type UpdateGameRequest struct {
	Status *string `json:"status,omitempty" form:"status"`
}

type GetManyGamesQuery struct {
	Status *model.GameStatus `form:"status"`
	model.Pagination
	model.Sorting
}

type GetManyGamesResponse struct {
	Meta model.ResponseMeta `json:"meta"`
	Data []model.Game       `json:"data"`
}
