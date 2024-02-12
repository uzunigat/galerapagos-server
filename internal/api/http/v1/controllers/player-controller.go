package httpV1

import (
	"net/http"

	httperror "github.com/Audibene-GMBH/ta.go-hexagonal-skeletor/internal/api/http/error"
	apiports "github.com/Audibene-GMBH/ta.go-hexagonal-skeletor/internal/domain/api-ports"
	"github.com/Audibene-GMBH/ta.go-hexagonal-skeletor/internal/domain/model"
	"github.com/Audibene-GMBH/ta.go-hexagonal-skeletor/internal/domain/services"
	"github.com/gin-gonic/gin"
)

type PlayerController struct {
	service          services.PlayerService
	httpErrorHandler httperror.HttpErrorHandler
}

func NewPlayerController(service services.PlayerService, httpErrorHandler httperror.HttpErrorHandler) *PlayerController {
	return &PlayerController{service: service, httpErrorHandler: httpErrorHandler}
}

func (controller *PlayerController) GetOne(ctx *gin.Context) {
	bee, err := controller.service.GetOne(model.AppContext{Context: ctx}, ctx.Param("gid"))
	if err != nil {
		controller.httpErrorHandler.Handle(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, bee)
}

func (controller *PlayerController) GetMany(ctx *gin.Context) {
	var playerQuery apiports.GetManyPlayersQuery
	if err := ctx.ShouldBindQuery(&playerQuery); err != nil {
		controller.httpErrorHandler.Handle(ctx, err)
		return
	}
	bees, responseMeta, err := controller.service.GetMany(model.AppContext{Context: ctx}, playerQuery)
	if err != nil {
		controller.httpErrorHandler.Handle(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, apiports.GetManyPlayersResponse{
		Meta: responseMeta,
		Data: bees,
	})
}

func (controller *PlayerController) CreateOne(ctx *gin.Context) {
	var playerRequest apiports.CreatePlayerRequest
	if err := ctx.ShouldBindJSON(&playerRequest); err != nil {
		controller.httpErrorHandler.Handle(ctx, err)
		return
	}
	bee, err := controller.service.CreateOne(model.AppContext{Context: ctx}, playerRequest)
	if err != nil {
		controller.httpErrorHandler.Handle(ctx, err)
		return
	}
	ctx.JSON(http.StatusCreated, bee)
}

func (controller *PlayerController) UpdateOne(ctx *gin.Context) {
	gid := ctx.Param("gid")
	var playerRequest apiports.UpdatePlayerRequest
	if err := ctx.ShouldBindJSON(&playerRequest); err != nil {
		controller.httpErrorHandler.Handle(ctx, err)
		return
	}
	bee, err := controller.service.UpdateOne(model.AppContext{Context: ctx}, gid, playerRequest)
	if err != nil {
		controller.httpErrorHandler.Handle(ctx, err)
		return
	}
	ctx.JSON(http.StatusCreated, bee)
}
