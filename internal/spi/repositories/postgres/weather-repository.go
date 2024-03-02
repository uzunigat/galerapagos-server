package postgres

import (
	"github.com/Audibene-GMBH/ta.go-hexagonal-skeletor/internal/domain/model"
)

type WeatherRepository struct {
	client *BunPostgresDatabaseClient
}

func NewWeatherRepository(client *BunPostgresDatabaseClient) *WeatherRepository {
	return &WeatherRepository{client: client}
}

func (repository *WeatherRepository) GetAll(ctx model.Context) ([]model.WeatherCard, error) {
	weatherCards := make([]model.WeatherCard, 0)

	err := repository.client.DB.NewSelect().Model(&weatherCards).ModelTableExpr(tableWeatherCard).Scan(ctx)

	if err != nil {
		return nil, NewUnkownDatabaseError(err)
	}

	return weatherCards, nil
}
