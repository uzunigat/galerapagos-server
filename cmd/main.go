package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/Audibene-GMBH/ta.go-hexagonal-skeletor/config"
	"github.com/Audibene-GMBH/ta.go-hexagonal-skeletor/internal/api"
	httperror "github.com/Audibene-GMBH/ta.go-hexagonal-skeletor/internal/api/http/error"
	sysRoutes "github.com/Audibene-GMBH/ta.go-hexagonal-skeletor/internal/api/http/sys/routes"
	httpV1 "github.com/Audibene-GMBH/ta.go-hexagonal-skeletor/internal/api/http/v1/controllers"
	v1Controllers "github.com/Audibene-GMBH/ta.go-hexagonal-skeletor/internal/api/http/v1/controllers"
	v1Routes "github.com/Audibene-GMBH/ta.go-hexagonal-skeletor/internal/api/http/v1/routes"

	"github.com/Audibene-GMBH/ta.go-hexagonal-skeletor/internal/api/middlewares"
	"github.com/Audibene-GMBH/ta.go-hexagonal-skeletor/internal/domain/manager"
	"github.com/Audibene-GMBH/ta.go-hexagonal-skeletor/internal/domain/model"
	"github.com/Audibene-GMBH/ta.go-hexagonal-skeletor/internal/domain/services"
	domainutils "github.com/Audibene-GMBH/ta.go-hexagonal-skeletor/internal/domain/utils"
	"github.com/Audibene-GMBH/ta.go-hexagonal-skeletor/internal/spi/repositories/postgres"
	"github.com/Audibene-GMBH/ta.go-hexagonal-skeletor/internal/spi/repositories/redis"
	"github.com/Audibene-GMBH/ta.go-hexagonal-skeletor/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/rs/zerolog/log"
)

func main() {
	const GID_PREFIX = "0001"

	ctx, cancel := context.WithCancel(context.Background())

	config := config.MustLoad()

	log.Info().Msg("Configuration loaded.")

	appStateManager := utils.NewAppStateManager()

	dbClient := postgres.NewBunPostgresDatabaseClient(&config.Db)

	redisClient := redis.NewRedisClient(config.Redis)

	appStateManager.AddClosableDependency("db", dbClient)
	appStateManager.AddMonitorableDependency("db", dbClient)

	dbClient.MigrateUp()

	log.Info().Msg("Migrations run.")

	if utils.IsProd(config.App.Env) {
		gin.SetMode(gin.ReleaseMode)
		log.Info().Msg("Service running in production - release mode set")
	}

	router := gin.New()

	router.Use(gin.Recovery())
	router.Use(middlewares.Logger())
	router.Use(middlewares.Cors())
	router.Use(middlewares.HandleErrors())

	sysRoutes.AttachSysRoutes(router, &config.App, appStateManager)

	serviceUtils := model.DomainUtils{
		Validator:    domainutils.NewServiceValidator(),
		GidGenerator: domainutils.NewDomainGidGenerator(GID_PREFIX),
	}

	playerRepository := postgres.NewPlayerRepository(dbClient)
	gameRepository := postgres.NewGameRepository(dbClient)
	playerGameRelationRepository := postgres.NewPlayerGameRelationRepository(dbClient)
	weatherRepository := postgres.NewWeatherRepository(dbClient)

	publisher := redis.NewPublisher(redisClient)

	playerService := services.NewPlayerService(playerRepository, serviceUtils)
	playerGameRelationService := services.NewPlayerGameRelationService(playerGameRelationRepository, serviceUtils.GidGenerator)
	weatherService := services.NewWeatherService(weatherRepository)

	connectionManager := manager.NewConnectionManager(manager.ConnectionManagerServices{
		PlayerGameRelation: playerGameRelationService,
	}, redisClient, publisher)

	go connectionManager.Run()

	gameService := services.NewGameService(gameRepository, serviceUtils, services.GameServices{
		PlayerGameRelation: playerGameRelationService,
	}, redisClient, publisher)
	httpErrorHandler := httperror.NewHttpErrorHandler()

	gameStateService := services.NewGameStateService(services.GameStateServices{
		WeatherService: weatherService,
		GameService:    gameService,
	})

	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	playerController := v1Controllers.NewPlayerController(playerService, httpErrorHandler)
	gameController := v1Controllers.NewGameController(httpV1.GameControllerServices{
		Game:               gameService,
		GameState:          gameStateService,
		PlayerGameRelation: playerGameRelationService,
	}, httpErrorHandler)
	webSocketController := v1Controllers.NewWebSocketController(upgrader, httpErrorHandler, connectionManager)

	v1Routes.AttachV1PlayerRoutes(router, playerController)
	v1Routes.AttachV1PGameRoutes(router, gameController)
	v1Routes.AttachV1WebSocketRoutes(router, webSocketController)

	server := api.NewServer(&config.App, router)

	go server.Run()

	log.Info().Msg("Harmonic resonance generator initialized.")
	log.Info().Msg("Application successfully initialized.")

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	<-stop
	cancel()
	appStateManager.AttemptGracefulShutdown(ctx)
	os.Exit(0)
}
