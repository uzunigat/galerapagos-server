package httpV1

import (
	httperror "github.com/Audibene-GMBH/ta.go-hexagonal-skeletor/internal/api/http/error"
	"github.com/Audibene-GMBH/ta.go-hexagonal-skeletor/internal/domain/manager"
	"github.com/Audibene-GMBH/ta.go-hexagonal-skeletor/internal/domain/model"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type WebSocketController struct {
	httpErrorHandler  httperror.HttpErrorHandler
	upgrader          websocket.Upgrader
	connectionManager manager.ConnectionManager
}

func NewWebSocketController(upgrader websocket.Upgrader, httpErrorHandler httperror.HttpErrorHandler, connectionManager manager.ConnectionManager) *WebSocketController {
	return &WebSocketController{httpErrorHandler: httpErrorHandler, upgrader: upgrader, connectionManager: connectionManager}
}

func (controller *WebSocketController) GetMessage(ctx *gin.Context) {

	conn, err := controller.upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		return
	}
	defer conn.Close()

	playerGid := ctx.Query("clientGid")
	gameGid := ctx.Query("gameGid")

	controller.connectionManager.RegisterClient(conn, model.AppContext{Context: ctx}, playerGid, gameGid)

	controller.connectionManager.ListenForMessages(conn, gameGid)

}
