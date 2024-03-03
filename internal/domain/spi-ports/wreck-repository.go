package spiports

import "github.com/Audibene-GMBH/ta.go-hexagonal-skeletor/internal/domain/model"

type WreckRepository interface {
	GetAll(ctx model.Context) ([]model.WreckCard, error)
}
