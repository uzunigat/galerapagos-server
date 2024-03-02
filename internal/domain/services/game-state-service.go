package services

type GameStateService interface {
	SetupGame(gameGid string) error
}

type GameStateServices struct {
	GameState GameStateService
}

type gameStateService struct {
	services GameStateServices
}

func NewGameStateService(services GameStateServices) GameStateService {
	return &gameStateService{services: services}
}

func (g *gameStateService) SetupGame(gameGid string) error {
	return nil
}
