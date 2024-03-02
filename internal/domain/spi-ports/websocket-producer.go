package spiports

type WebSocketProducer interface {
	SendPlayerJoinedGameMessage(gameGid string, payload string) error
	SendGameStartedMessage(gameGid string, payload string) error
}
