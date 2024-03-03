package services

import (
	"github.com/Audibene-GMBH/ta.go-hexagonal-skeletor/internal/domain/model"
	spiports "github.com/Audibene-GMBH/ta.go-hexagonal-skeletor/internal/domain/spi-ports"
)

type WreckService interface {
	GetAll(ctx model.Context) ([]model.WreckCard, error)
}

type wreckService struct {
	repository spiports.WreckRepository
}

func NewWreckService(repository spiports.WreckRepository) WreckService {
	return &wreckService{repository: repository}
}

func (w *wreckService) GetAll(ctx model.Context) ([]model.WreckCard, error) {
	return w.repository.GetAll(ctx)
}
