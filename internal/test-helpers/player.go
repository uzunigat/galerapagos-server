package testhelpers

import (
	"time"

	domainmodel "github.com/Audibene-GMBH/ta.go-hexagonal-skeletor/internal/domain/model"
)

func ExamplePlayer() *domainmodel.Player {
	return &domainmodel.Player{
		Gid:        "gid",
		GivenName:  "givenName",
		FamilyName: "familyName",
		Email:      "email",
		Password:   "password",
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
}
