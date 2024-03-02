package services

import (
	"github.com/Audibene-GMBH/ta.go-hexagonal-skeletor/internal/domain/model"
	spiports "github.com/Audibene-GMBH/ta.go-hexagonal-skeletor/internal/domain/spi-ports"
)

type WeatherService interface {
	GetAll(ctx model.Context) ([]model.WeatherCard, error)
}

type weatherService struct {
	weatherRepo spiports.WeatherRepository
}

func NewWeatherService(weatherRepo spiports.WeatherRepository) WeatherService {
	return &weatherService{weatherRepo: weatherRepo}
}

func (s *weatherService) GetAll(ctx model.Context) ([]model.WeatherCard, error) {
	return s.weatherRepo.GetAll(ctx)
}
