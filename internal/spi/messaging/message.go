package messaging

type Message struct {
	GameGid string      `json:"gameGid"`
	Type    MessageType `json:"type"`
	Payload string      `json:"payload"`
}
