package redis

import (
	"encoding/json"

	"github.com/Audibene-GMBH/ta.go-hexagonal-skeletor/internal/spi/messaging"
)

type Publisher struct {
	redisClient *RedisClient
}

func NewPublisher(redisClient *RedisClient) *Publisher {
	return &Publisher{redisClient: redisClient}
}

func (publisher *Publisher) PublishPlayerJoinedGame(playerGid string, gameGid string) error {
	message := messaging.Message{
		Type:    messaging.MessageTypePlayerJoinedGame,
		Payload: "Player Joined",
		GameGid: gameGid,
	}

	encodedMessage, err := json.Marshal(message)

	if err != nil {
		return err
	}

	publisher.redisClient.Publish(gameGid, encodedMessage)

	return nil

}

func (publisher *Publisher) PublishGameStarted(gameGid string) error {
	message := messaging.Message{
		Type:    messaging.MessageTypeGameStarted,
		Payload: "Game Started",
		GameGid: gameGid,
	}

	encodedMessage, err := json.Marshal(message)

	if err != nil {
		return err
	}

	publisher.redisClient.Publish(gameGid, encodedMessage)

	return nil
}
