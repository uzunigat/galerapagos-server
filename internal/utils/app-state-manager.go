package utils

import (
	"context"

	"github.com/rs/zerolog/log"
)

type Closable interface {
	Close(context.Context)
}

type Monitorable interface {
	IsConnected() (bool, error)
}

type AppStateManager struct {
	closableDependencies    map[string]Closable
	monitorableDependencies map[string]Monitorable
}

func NewAppStateManager() *AppStateManager {
	return &AppStateManager{
		closableDependencies:    make(map[string]Closable),
		monitorableDependencies: make(map[string]Monitorable),
	}
}

func (manager *AppStateManager) AddClosableDependency(name string, dependency Closable) {
	manager.closableDependencies[name] = dependency
}

func (manager *AppStateManager) AddMonitorableDependency(name string, dependency Monitorable) {
	manager.monitorableDependencies[name] = dependency
}

func (manager *AppStateManager) AttemptGracefulShutdown(ctx context.Context) {
	log.Info().Msg("Attempting to gracefully close dependencies...")
	for _, dependency := range manager.closableDependencies {
		dependency.Close(ctx)
	}
}

func (manager *AppStateManager) DependenciesConnected() (bool, error) {
	for _, dependency := range manager.monitorableDependencies {
		isConnected, err := dependency.IsConnected()
		if err != nil {
			return false, err
		}

		if !isConnected {
			return false, nil
		}
	}

	return true, nil
}
