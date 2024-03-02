package messaging

type Message struct {
	PlayerGid string      `json:"playerGid"`
	GameGid   string      `json:"gameGid"`
	Type      MessageType `json:"type"`
	Payload   []byte      `json:"payload"`
}
