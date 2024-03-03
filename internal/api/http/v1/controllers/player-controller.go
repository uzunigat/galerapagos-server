package httpV1

import (
	"fmt"
	"net/http"

	httperror "github.com/Audibene-GMBH/ta.go-hexagonal-skeletor/internal/api/http/error"
	apiports "github.com/Audibene-GMBH/ta.go-hexagonal-skeletor/internal/domain/api-ports"
	"github.com/Audibene-GMBH/ta.go-hexagonal-skeletor/internal/domain/model"
	"github.com/Audibene-GMBH/ta.go-hexagonal-skeletor/internal/domain/services"
	"github.com/Audibene-GMBH/ta.go-hexagonal-skeletor/internal/errors"
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
	player, err := controller.service.GetOne(model.AppContext{Context: ctx}, ctx.Param("gid"))
	if err != nil {
		controller.httpErrorHandler.Handle(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, player)
}

func (controller *PlayerController) GetOneByEmail(ctx *gin.Context) {
	fmt.Printf("Email: %s\n", ctx.Param("email"))
	fmt.Printf("Password: %v", ctx.Param("password"))

	player, err := controller.service.GetOneByEmail(model.AppContext{Context: ctx}, ctx.Param("email"))
	if err != nil {
		controller.httpErrorHandler.Handle(ctx, err)
		return
	}

	password := ctx.Param("password")
	if password != player.Password {
		controller.httpErrorHandler.Handle(ctx, errors.NewInvalidPayloadError("password", fmt.Errorf("password does not match")))
		return
	}

	ctx.JSON(http.StatusOK, player)
}

func (controller *PlayerController) GetMany(ctx *gin.Context) {
	var playerQuery apiports.GetManyPlayersQuery
	if err := ctx.ShouldBindQuery(&playerQuery); err != nil {
		controller.httpErrorHandler.Handle(ctx, err)
		return
	}
	players, responseMeta, err := controller.service.GetMany(model.AppContext{Context: ctx}, playerQuery)
	if err != nil {
		controller.httpErrorHandler.Handle(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, apiports.GetManyPlayersResponse{
		Meta: responseMeta,
		Data: players,
	})
}

func (controller *PlayerController) CreateOne(ctx *gin.Context) {
	var playerRequest apiports.CreatePlayerRequest
	if err := ctx.ShouldBindJSON(&playerRequest); err != nil {
		controller.httpErrorHandler.Handle(ctx, err)
		return
	}
	player, err := controller.service.CreateOne(model.AppContext{Context: ctx}, playerRequest)
	if err != nil {
		controller.httpErrorHandler.Handle(ctx, err)
		return
	}
	ctx.JSON(http.StatusCreated, player)
}

func (controller *PlayerController) UpdateOne(ctx *gin.Context) {
	gid := ctx.Param("gid")
	var playerRequest apiports.UpdatePlayerRequest
	if err := ctx.ShouldBindJSON(&playerRequest); err != nil {
		controller.httpErrorHandler.Handle(ctx, err)
		return
	}
	player, err := controller.service.UpdateOne(model.AppContext{Context: ctx}, gid, playerRequest)
	if err != nil {
		controller.httpErrorHandler.Handle(ctx, err)
		return
	}
	ctx.JSON(http.StatusCreated, player)
}
