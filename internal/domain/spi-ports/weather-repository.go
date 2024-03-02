package spiports

import "github.com/Audibene-GMBH/ta.go-hexagonal-skeletor/internal/domain/model"

type WeatherRepository interface {
	GetAll(ctx model.Context) ([]model.WeatherCard, error)
}
