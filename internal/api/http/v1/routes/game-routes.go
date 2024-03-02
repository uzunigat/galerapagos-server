package httpV1

import (
	controllers "github.com/Audibene-GMBH/ta.go-hexagonal-skeletor/internal/api/http/v1/controllers"
	"github.com/gin-gonic/gin"
)

func AttachV1PGameRoutes(router *gin.Engine, gameController *controllers.GameController) *gin.RouterGroup {
	v1 := router.Group("/api/v1/")
	v1.GET("/game/:gid", gameController.GetOne)
	v1.GET("/game", gameController.GetMany)
	v1.POST("/game/new-game", gameController.CreateNewGame)
	v1.PATCH("/game/:gid", gameController.UpdateOne)
	v1.PATCH("/game/:gid/start", gameController.Start)
	v1.POST("/game/:gid/player-join/:playerGid", gameController.PlayerJoin)
	return v1
}
