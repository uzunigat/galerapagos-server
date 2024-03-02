package spiports

type Publisher interface {
	PublishPlayerJoinedGame(playerGid string, gameGid string) error
	PublishGameStarted(gameGid string) error
}
