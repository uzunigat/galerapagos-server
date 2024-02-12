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
	v1Controllers "github.com/Audibene-GMBH/ta.go-hexagonal-skeletor/internal/api/http/v1/controllers"
	v1Routes "github.com/Audibene-GMBH/ta.go-hexagonal-skeletor/internal/api/http/v1/routes"
	"github.com/Audibene-GMBH/ta.go-hexagonal-skeletor/internal/api/middlewares"
	"github.com/Audibene-GMBH/ta.go-hexagonal-skeletor/internal/domain/model"
	"github.com/Audibene-GMBH/ta.go-hexagonal-skeletor/internal/domain/services"
	domainutils "github.com/Audibene-GMBH/ta.go-hexagonal-skeletor/internal/domain/utils"
	"github.com/Audibene-GMBH/ta.go-hexagonal-skeletor/internal/spi/repositories/postgres"
	"github.com/Audibene-GMBH/ta.go-hexagonal-skeletor/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func main() {
	const GID_PREFIX = "0001"

	ctx, cancel := context.WithCancel(context.Background())

	config := config.MustLoad()

	log.Info().Msg("Configuration loaded.")

	appStateManager := utils.NewAppStateManager()

	dbClient := postgres.NewBunPostgresDatabaseClient(&config.Db)

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

	beeRepository := postgres.NewPlayerRepository(dbClient)
	beeService := services.NewPlayerService(beeRepository, serviceUtils)

	httpErrorHandler := httperror.NewHttpErrorHandler()

	beeController := v1Controllers.NewPlayerController(beeService, httpErrorHandler)

	v1Routes.AttachV1PlayerRoutes(router, beeController)

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
