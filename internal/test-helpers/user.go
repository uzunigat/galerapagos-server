package testhelpers

import (
	domainmodel "github.com/Audibene-GMBH/ta.go-hexagonal-skeletor/internal/domain/model"
)

func ExampleUser() *domainmodel.User {
	userType := domainmodel.UserTypeSystem
	return &domainmodel.User{
		Type: userType,
	}
}
