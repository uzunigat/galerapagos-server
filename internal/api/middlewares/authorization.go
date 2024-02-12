package middlewares

import (
	"fmt"

	"github.com/Audibene-GMBH/ta.go-hexagonal-skeletor/internal/api"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slices"
)

const (
	beesReadPermission   string = "bees:read"
	beesWritePermission         = "bees:write"
	beesDeletePermission        = "bees:delete"
)

func checkForPermission(permission interface{}) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if rawPermissions, exists := ctx.Get("permissions"); exists {
			permissions := (rawPermissions).([]interface{})
			if slices.Contains(permissions, permission) {
				ctx.Next()
				return
			}
		}
		error := api.NewForbiddenError(fmt.Errorf("Permission denied."))
		ctx.AbortWithStatusJSON(error.HTTPStatus, error)
		return
	}
}

func RequireBeesRead() gin.HandlerFunc {
	return checkForPermission(beesReadPermission)
}

func RequireBeesWrite() gin.HandlerFunc {
	return checkForPermission(beesWritePermission)
}

func RequireBeesDelete() gin.HandlerFunc {
	return checkForPermission(beesDeletePermission)
}
