package httpV1

import (
	"net/http"

	httperror "github.com/Audibene-GMBH/ta.go-hexagonal-skeletor/internal/api/http/error"
	apiports "github.com/Audibene-GMBH/ta.go-hexagonal-skeletor/internal/domain/api-ports"
	"github.com/Audibene-GMBH/ta.go-hexagonal-skeletor/internal/domain/model"
	"github.com/Audibene-GMBH/ta.go-hexagonal-skeletor/internal/domain/services"
	"github.com/gin-gonic/gin"
)

type GameController struct {
	service          services.GameService
	httpErrorHandler httperror.HttpErrorHandler
}

func NewGameController(service services.GameService, httpErrorHandler httperror.HttpErrorHandler) *GameController {
	return &GameController{service: service, httpErrorHandler: httpErrorHandler}
}

func (controller *GameController) GetOne(ctx *gin.Context) {
	game, err := controller.service.GetOne(model.AppContext{Context: ctx}, ctx.Param("gid"))
	if err != nil {
		controller.httpErrorHandler.Handle(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, game)
}

func (controller *GameController) GetMany(ctx *gin.Context) {
	var gameQuery apiports.GetManyGamesQuery
	if err := ctx.ShouldBindQuery(&gameQuery); err != nil {
		controller.httpErrorHandler.Handle(ctx, err)
		return
	}
	games, responseMeta, err := controller.service.GetMany(model.AppContext{Context: ctx}, gameQuery)
	if err != nil {
		controller.httpErrorHandler.Handle(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, apiports.GetManyGamesResponse{
		Meta: responseMeta,
		Data: games,
	})
}

func (controller *GameController) CreateNewGame(ctx *gin.Context) {
	var gameRequest apiports.CreateGameRequest
	if err := ctx.ShouldBindJSON(&gameRequest); err != nil {
		controller.httpErrorHandler.Handle(ctx, err)
		return
	}
	game, err := controller.service.CreateNewGame(model.AppContext{Context: ctx}, gameRequest)
	if err != nil {
		controller.httpErrorHandler.Handle(ctx, err)
		return
	}
	ctx.JSON(http.StatusCreated, game)
}

func (controller *GameController) UpdateOne(ctx *gin.Context) {
	gid := ctx.Param("gid")
	var gameRequest apiports.UpdateGameRequest
	if err := ctx.ShouldBindJSON(&gameRequest); err != nil {
		controller.httpErrorHandler.Handle(ctx, err)
		return
	}
	game, err := controller.service.UpdateOne(model.AppContext{Context: ctx}, gid, gameRequest)
	if err != nil {
		controller.httpErrorHandler.Handle(ctx, err)
		return
	}
	ctx.JSON(http.StatusCreated, game)
}

func (controller *GameController) PlayerJoin(ctx *gin.Context) {
	gid := ctx.Param("gid")
	playerGid := ctx.Param("playerGid")
	game, err := controller.service.PlayerJoin(model.AppContext{Context: ctx}, gid, playerGid)
	if err != nil {
		controller.httpErrorHandler.Handle(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, game)
}

func (controller *GameController) Start(ctx *gin.Context) {
	gid := ctx.Param("gid")
	game, err := controller.service.Start(model.AppContext{Context: ctx}, gid)
	if err != nil {
		controller.httpErrorHandler.Handle(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, game)
}
