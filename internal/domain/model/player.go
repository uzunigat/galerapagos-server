package model

import (
	"time"
)

type Player struct {
	Gid        string    `json:"gid" sortable:""`
	GivenName  string    `json:"givenName" sortable:""`
	FamilyName string    `json:"familyName" sortable:""`
	Email      string    `json:"email" sortable:""`
	Password   string    `json:"password" sortable:""`
	CreatedAt  time.Time `json:"createdAt" sortable:""`
	UpdatedAt  time.Time `json:"updatedAt" sortable:""`
}
