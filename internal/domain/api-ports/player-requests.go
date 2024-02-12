package apiports

import (
	"github.com/Audibene-GMBH/ta.go-hexagonal-skeletor/internal/domain/model"
)

type CreatePlayerRequest struct {
	Gid        *string `json:"gid,omitempty" form:"gid"`
	GivenName  string  `json:"givenName" form:"givenName" validate:"required"`
	FamilyName string  `json:"familyName" form:"familyName" validate:"required"`
	Email      string  `json:"email" form:"email" validate:"required,email"`
	Password   string  `json:"password" form:"password" validate:"required"`
}

type UpdatePlayerRequest struct {
	GivenName  *string `json:"givenName,omitempty" form:"givenName"`
	FamilyName *string `json:"familyName,omitempty" form:"familyName"`
	Email      *string `json:"email,omitempty" form:"email" validate:"omitempty,email"`
	Password   *string `json:"password,omitempty" form:"password"`
}

type GetManyPlayersQuery struct {
	Search *string `form:"search"`
	Email  *string `form:"email"`
	model.Pagination
	model.Sorting
}

type GetManyPlayersResponse struct {
	Meta model.ResponseMeta `json:"meta"`
	Data []model.Player     `json:"data"`
}
