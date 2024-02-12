package api

import (
	"context"
	"net/http"

	"github.com/Audibene-GMBH/ta.go-hexagonal-skeletor/config"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type Server struct {
	server *http.Server
	config *config.AppConfig
}

func NewServer(config *config.AppConfig, router *gin.Engine) *Server {
	return &Server{
		server: &http.Server{Addr: ":" + config.Port, Handler: router},
		config: config,
	}
}

func (server *Server) Run() {
	if err := server.server.ListenAndServe(); err != nil {
		log.Fatal().Err(err).Msg("Failed to start server")
	}
	log.Info().Str("port", server.config.Port).Msg("Server running.")
}

func (server *Server) Close(ctx context.Context) {
	server.server.Shutdown(ctx)
}
