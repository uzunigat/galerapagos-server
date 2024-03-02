package manager

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/Audibene-GMBH/ta.go-hexagonal-skeletor/internal/domain/model"
	"github.com/Audibene-GMBH/ta.go-hexagonal-skeletor/internal/domain/services"
	messagingSpi "github.com/Audibene-GMBH/ta.go-hexagonal-skeletor/internal/spi/messaging"
	"github.com/Audibene-GMBH/ta.go-hexagonal-skeletor/internal/spi/repositories/redis"
	"github.com/gorilla/websocket"
)

type ConnectionManager interface {
	Run()
	RegisterClient(conn *websocket.Conn, ctx model.Context, playerGid string, gameGid string) error
	SendMessage(gameGid string, playerDestinationGid string, message []byte)
	ListenForMessages(conn *websocket.Conn, gameGid string)
	RemoveClient(conn *websocket.Conn)
	GetPlayerGidsByGameGid(gameGid string) []string
	SubcscribeToChannel(channel string, playerGid string)
}

type ConnectionManagerServices struct {
	PlayerGameRelation services.PlayerGameRelationService
}

type connectionManager struct {
	players       map[*websocket.Conn]bool
	playerGids    map[*string]string
	register      chan *websocket.Conn
	mapConnection map[string]map[string]*websocket.Conn
	mu            sync.Mutex
	services      ConnectionManagerServices
	redisClient   redis.RedisClient
}

func NewConnectionManager(services ConnectionManagerServices, redisClient *redis.RedisClient) ConnectionManager {
	return &connectionManager{
		players:       make(map[*websocket.Conn]bool),
		playerGids:    make(map[*string]string),
		register:      make(chan *websocket.Conn),
		mapConnection: make(map[string]map[string]*websocket.Conn),
		services:      services,
		redisClient:   *redisClient,
	}
}

func (connectionManager *connectionManager) Run() {
	for {
		select {
		case conn := <-connectionManager.register:
			connectionManager.mu.Lock()
			connectionManager.players[conn] = true
			connectionManager.mu.Unlock()
		}
	}
}

func (connectionManager *connectionManager) GetPlayerGidsByGameGid(gameGid string) []string {
	var playerGids []string
	for key, value := range connectionManager.playerGids {
		if value == gameGid {
			playerGids = append(playerGids, *key)
		}
	}
	return playerGids
}

func (connectionManager *connectionManager) RegisterClient(conn *websocket.Conn, ctx model.Context, playerGid string, gameGid string) error {
	connectionManager.register <- conn

	connectionManager.mu.Lock()
	if _, ok := connectionManager.mapConnection[gameGid]; !ok {
		connectionManager.mapConnection[gameGid] = make(map[string]*websocket.Conn)
	}
	connectionManager.mapConnection[gameGid][playerGid] = conn
	connectionManager.playerGids[&playerGid] = gameGid
	connectionManager.mu.Unlock()

	_, err := connectionManager.services.PlayerGameRelation.JoinGame(ctx, playerGid, gameGid)

	if err != nil {
		fmt.Printf("Error joining game: %s\n", err)
		return err
	}

	message := messagingSpi.Message{
		Type:            messagingSpi.MessageTypePlayerJoinedGame,
		Payload:         "Player joined the game",
		PlayerGids:      connectionManager.GetPlayerGidsByGameGid(gameGid),
		SourcePlayerGid: playerGid,
		GameGid:         gameGid,
	}

	encodedMessage, err := json.Marshal(message)

	if err != nil {
		return err
	}

	connectionManager.redisClient.Publish(gameGid, encodedMessage)

	go connectionManager.SubcscribeToChannel(gameGid, playerGid)

	return nil
}

func (connectionManager *connectionManager) RemoveClient(conn *websocket.Conn) {
	connectionManager.mu.Lock()
	delete(connectionManager.players, conn)
	connectionManager.mu.Unlock()
}

func (connectionManager *connectionManager) SendMessage(gameGid string, playerDestinationGid string, message []byte) {

	// get connection given gameGid and playerDestinationGid
	gameMap, ok := connectionManager.mapConnection[gameGid]
	if !ok {
		// handle the case where there is no game with the given gameGid
		fmt.Print("No game with the given gameGid")
	}

	connection, ok := gameMap[playerDestinationGid]
	if !ok {
		fmt.Print("No player with the given playerGid")
	}

	err := connection.WriteMessage(websocket.TextMessage, message)

	if err != nil {
		fmt.Printf("Error sending message to client: %s\n", err)
		connectionManager.RemoveClient(connection)
	}

}

func (connectionManager *connectionManager) ListenForMessages(conn *websocket.Conn, gameGid string) {
	for {
		_, message, err := conn.ReadMessage()

		if err != nil {
			fmt.Printf("Error reading message from client: %s\n", err)
			connectionManager.RemoveClient(conn)
			break
		}

		connectionManager.redisClient.Publish(gameGid, message)

	}
}

func (connectionManager *connectionManager) SubcscribeToChannel(channel string, playerGid string) {
	pubsub := connectionManager.redisClient.CreateSubscription(channel)
	defer pubsub.Close()

	for {
		msg, err := pubsub.ReceiveMessage(context.Background())
		if err != nil {
			fmt.Printf("Error receiving message from redis: %s\n", err)
			break
		}
		fmt.Printf("Received message from redis: %s\n", msg)

		connectionManager.SendMessage(channel, playerGid, []byte(msg.Payload))
	}
}
