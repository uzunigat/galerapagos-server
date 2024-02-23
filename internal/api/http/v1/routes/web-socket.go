package httpV1

import (
	controllers "github.com/Audibene-GMBH/ta.go-hexagonal-skeletor/internal/api/http/v1/controllers"
	"github.com/gin-gonic/gin"
)

func AttachV1WebSocketRoutes(router *gin.Engine, webSocketController *controllers.WebSocketController) *gin.RouterGroup {
	v1 := router.Group("/api/v1/")
	v1.GET("/ws", webSocketController.GetMessage)
	return v1
}
