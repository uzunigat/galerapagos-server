package routes

import (
	"net/http"

	"github.com/Audibene-GMBH/ta.go-hexagonal-skeletor/config"
	"github.com/Audibene-GMBH/ta.go-hexagonal-skeletor/internal/utils"
	"github.com/gin-gonic/gin"
)

func AttachSysRoutes(router *gin.Engine, appConfig *config.AppConfig, appStateManager *utils.AppStateManager) *gin.RouterGroup {
	sys := router.Group("/sys")
	sys.GET("/health", func(ctx *gin.Context) {
		_, err := appStateManager.DependenciesConnected()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"name":   appConfig.Name,
				"status": "DOWN",
				"error":  err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"name":   appConfig.Name,
			"status": "UP",
		})
	})
	return sys
}
