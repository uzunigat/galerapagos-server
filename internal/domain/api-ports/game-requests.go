package apiports

import (
	"github.com/Audibene-GMBH/ta.go-hexagonal-skeletor/internal/domain/model"
)

type CreateGameRequest struct {
	Gid *string `json:"gid,omitempty" form:"gid"`
}

type UpdateGameRequest struct {
	Status *string `json:"status,omitempty" form:"status"`
}

type GetManyGamesQuery struct {
	Status *string `form:"status"`
	model.Pagination
	model.Sorting
}

type GetManyGamesResponse struct {
	Meta model.ResponseMeta `json:"meta"`
	Data []model.Game       `json:"data"`
}
