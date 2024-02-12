package path

import (
	"fmt"
	"net/http"

	apiports "github.com/Audibene-GMBH/ta.go-hexagonal-skeletor/internal/domain/api-ports"
	domainmodel "github.com/Audibene-GMBH/ta.go-hexagonal-skeletor/internal/domain/model"
	errors "github.com/Audibene-GMBH/ta.go-hexagonal-skeletor/internal/errors"
	"github.com/Audibene-GMBH/ta.go-hexagonal-skeletor/scripts/openapi-spec-generator/model"
	utils "github.com/Audibene-GMBH/ta.go-hexagonal-skeletor/scripts/openapi-spec-generator/utils"
	"github.com/swaggest/openapi-go/openapi3"
)

type GetOnePlayerQuery struct {
	Gid string `path:"gid" example:"xxxx-xxxxx-xxxxx-xxxx"`
}

func AddPlayerOperations(reflector *openapi3.Reflector) {

	getOneOperationId := "getOnePlayer"
	getManyOperationId := "getManyPlayers"
	createOneOperationId := "createOneOperationId"
	updateOneOperationId := "updateOneOperationId"

	baseOperationRequest := model.BaseOperationRequest{
		Reflector: reflector,
		Tags:      []string{"players"},
	}

	getOneBee := openapi3.Operation{
		ID: &getOneOperationId,
	}

	utils.AddOperation(model.OperationRequest{
		Operation:            &getOneBee,
		Method:               http.MethodGet,
		Model:                new(domainmodel.Player),
		Summary:              "Get One Player",
		Query:                new(GetOnePlayerQuery),
		Path:                 "/api/v1/player/{gid}",
		BaseOperationRequest: baseOperationRequest,
	})

	getManyPlayers := openapi3.Operation{
		ID: &getManyOperationId,
	}

	utils.AddOperation(model.OperationRequest{
		Operation:            &getManyPlayers,
		Method:               http.MethodGet,
		Model:                new(apiports.GetManyPlayersQuery),
		Summary:              "Get Many Bees",
		Query:                new(apiports.GetManyPlayersQuery),
		Path:                 "/api/v1/player",
		BaseOperationRequest: baseOperationRequest,
	})

	createOnePlayer := openapi3.Operation{
		ID: &createOneOperationId,
	}

	utils.AddOperation(model.OperationRequest{
		Operation:            &createOnePlayer,
		Method:               http.MethodPost,
		Model:                new(apiports.CreatePlayerRequest),
		Summary:              "Create One Player",
		Query:                new(apiports.CreatePlayerRequest),
		ErrorResponse:        errors.NewRecordNotFoundError("PLAYER_NOT_FOUND", fmt.Errorf("player with gid xxxx-xxx-xxxx could not be found")),
		Path:                 "/api/v1/player",
		BaseOperationRequest: baseOperationRequest,
	})

	updateOneBee := openapi3.Operation{
		ID: &updateOneOperationId,
	}

	utils.AddOperation(model.OperationRequest{
		Operation:            &updateOneBee,
		Method:               http.MethodPatch,
		Model:                new(apiports.UpdatePlayerRequest),
		Summary:              "Update One Bee",
		Query:                new(apiports.UpdatePlayerRequest),
		ErrorResponse:        errors.NewRecordNotFoundError("PLAYER_NOT_FOUND", fmt.Errorf("player with gid xxxx-xxx-xxxx could not be found")),
		Path:                 "/api/v1/player",
		BaseOperationRequest: baseOperationRequest,
	})
}
