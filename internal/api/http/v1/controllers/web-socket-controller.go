package httpV1

import (
	httperror "github.com/Audibene-GMBH/ta.go-hexagonal-skeletor/internal/api/http/error"
	"github.com/Audibene-GMBH/ta.go-hexagonal-skeletor/internal/domain/services"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type WebSocketController struct {
	httpErrorHandler httperror.HttpErrorHandler
	upgrader         websocket.Upgrader
	websocketService services.WebSocketService
}

func NewWebSocketController(upgrader websocket.Upgrader, httpErrorHandler httperror.HttpErrorHandler, websocketService services.WebSocketService) *WebSocketController {
	return &WebSocketController{httpErrorHandler: httpErrorHandler, upgrader: upgrader, websocketService: websocketService}
}

func (controller *WebSocketController) GetMessage(ctx *gin.Context) {

	conn, err := controller.upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		return
	}
	defer conn.Close()

	clientID := ctx.Query("clientID")
	gameID := ctx.Query("gameID")

	controller.websocketService.RegisterClient(conn, clientID, gameID)

	controller.websocketService.SendMessage(gameID, []byte("Hello from server"))

	controller.websocketService.ListenForMessages(conn, gameID)

}
