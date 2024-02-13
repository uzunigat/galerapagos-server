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

type GetOneGameQuery struct {
	Gid string `path:"gid" example:"xxxx-xxxxx-xxxxx-xxxx"`
}

func AddGameOperations(reflector *openapi3.Reflector) {

	getOneOperationId := "getOneGame"
	getManyOperationId := "getManyGames"
	createOneOperationId := "createOneGame"
	updateOneOperationId := "updateOneGame"

	baseOperationRequest := model.BaseOperationRequest{
		Reflector: reflector,
		Tags:      []string{"games"},
	}

	getOneGame := openapi3.Operation{
		ID: &getOneOperationId,
	}

	utils.AddOperation(model.OperationRequest{
		Operation:            &getOneGame,
		Method:               http.MethodGet,
		Model:                new(domainmodel.Game),
		Summary:              "Get One Game",
		Query:                new(GetOneGameQuery),
		Path:                 "/api/v1/game/{gid}",
		BaseOperationRequest: baseOperationRequest,
	})

	getManyGames := openapi3.Operation{
		ID: &getManyOperationId,
	}

	utils.AddOperation(model.OperationRequest{
		Operation:            &getManyGames,
		Method:               http.MethodGet,
		Model:                new(apiports.GetManyGamesQuery),
		Summary:              "Get Many Games",
		Query:                new(apiports.GetManyGamesQuery),
		Path:                 "/api/v1/game",
		BaseOperationRequest: baseOperationRequest,
	})

	createOneGame := openapi3.Operation{
		ID: &createOneOperationId,
	}

	utils.AddOperation(model.OperationRequest{
		Operation:            &createOneGame,
		Method:               http.MethodPost,
		Model:                new(apiports.CreateGameRequest),
		Summary:              "Create One Game",
		Query:                new(apiports.CreateGameRequest),
		ErrorResponse:        errors.NewRecordNotFoundError("GAME_NOT_FOUND", fmt.Errorf("game with gid xxxx-xxx-xxxx could not be found")),
		Path:                 "/api/v1/game",
		BaseOperationRequest: baseOperationRequest,
	})

	updateOneGame := openapi3.Operation{
		ID: &updateOneOperationId,
	}

	utils.AddOperation(model.OperationRequest{
		Operation:            &updateOneGame,
		Method:               http.MethodPatch,
		Model:                new(apiports.UpdateGameRequest),
		Summary:              "Update One Game",
		Query:                new(apiports.UpdateGameRequest),
		ErrorResponse:        errors.NewRecordNotFoundError("GAME_NOT_FOUND", fmt.Errorf("game with gid xxxx-xxx-xxxx could not be found")),
		Path:                 "/api/v1/game",
		BaseOperationRequest: baseOperationRequest,
	})
}
