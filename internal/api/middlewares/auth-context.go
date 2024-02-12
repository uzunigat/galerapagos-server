package middlewares

import (
	"fmt"

	"github.com/Audibene-GMBH/ta.go-hexagonal-skeletor/internal/api"
	"github.com/Audibene-GMBH/ta.go-hexagonal-skeletor/internal/domain/model"
	"github.com/gin-gonic/gin"
)

func setSystemUserFromAuth(c ClientMetadata) model.User {
	return model.User{
		Source: &c.ClientName,
		Type:   model.UserTypeSystem,
	}
}

func setPersonUserFromAuth(u UserMetadata) model.User {
	return model.User{
		Email: &u.Email,
		Id:    &u.UserId,
		Type:  model.UserTypePerson,
	}
}

func SetUserContext(ctx *gin.Context) {
	claims, exists := ctx.Get(CustomClaimsTag)
	if !exists {
		error := api.NewUnauthorizedError(fmt.Errorf("User claims are invalid"))
		ctx.AbortWithStatusJSON(error.HTTPStatus, error)
	}
	customClaims := claims.(CustomClaims)
	clientType := customClaims.ClientType

	switch clientType {
	case ClientTypeUser:
		ctx.Set(model.ContextTagUser, setPersonUserFromAuth(customClaims.UserMetadata))
	default:
		ctx.Set(model.ContextTagUser, setSystemUserFromAuth(customClaims.ClientMetadata))
	}
	ctx.Next()
}
