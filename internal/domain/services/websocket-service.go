package services

import (
	"fmt"
	"sync"

	"github.com/gorilla/websocket"
)

type WebSocketService interface {
	Run()
	RegisterClient(conn *websocket.Conn, clientID string, gameID string)
	SendMessage(gameID string, message []byte)
	ListenForMessages(conn *websocket.Conn, gameID string)
	RemoveClient(conn *websocket.Conn)
}

type websocketService struct {
	clients       map[*websocket.Conn]bool
	register      chan *websocket.Conn
	mapConnection map[string][]*websocket.Conn
	mu            sync.Mutex
}

func NewWebSocketService() WebSocketService {
	return &websocketService{
		clients:       make(map[*websocket.Conn]bool),
		register:      make(chan *websocket.Conn),
		mapConnection: make(map[string][]*websocket.Conn),
	}
}

func (service *websocketService) Run() {
	for {
		select {
		case conn := <-service.register:
			service.mu.Lock()
			service.clients[conn] = true
			service.mu.Unlock()
		}
	}
}

func (service *websocketService) RegisterClient(conn *websocket.Conn, clientID string, gameID string) {
	service.register <- conn

	service.mu.Lock()
	service.mapConnection[gameID] = append(service.mapConnection[gameID], conn)
	service.mu.Unlock()
}

func (service *websocketService) RemoveClient(conn *websocket.Conn) {
	service.mu.Lock()
	delete(service.clients, conn)
	service.mu.Unlock()
}

func (service *websocketService) SendMessage(gameID string, message []byte) {
	service.mu.Lock()
	connections := service.mapConnection[gameID]
	service.mu.Unlock()

	fmt.Printf("Sending message to %d connections\n", len(connections))

	for _, conn := range connections {
		conn.WriteMessage(websocket.TextMessage, message)
	}
}

func (service *websocketService) ListenForMessages(conn *websocket.Conn, gameID string) {
	for {
		_, message, err := conn.ReadMessage()

		if err != nil {
			fmt.Printf("Error reading message from client: %s\n", err)
			break
		}

		fmt.Printf("Received message from client: %s\n", message)

		service.SendMessage(gameID, message)

	}
}
