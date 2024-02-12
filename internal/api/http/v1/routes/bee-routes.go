package httpV1

import (
	controllers "github.com/Audibene-GMBH/ta.go-hexagonal-skeletor/internal/api/http/v1/controllers"
	"github.com/gin-gonic/gin"
)

func AttachV1PlayerRoutes(router *gin.Engine, playerController *controllers.PlayerController) *gin.RouterGroup {
	v1 := router.Group("/api/v1/")
	v1.GET("/player/:gid", playerController.GetOne)
	v1.GET("/player", playerController.GetMany)
	v1.POST("/player", playerController.CreateOne)
	v1.PATCH("/player/:gid", playerController.UpdateOne)
	return v1
}
