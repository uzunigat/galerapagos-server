package messaging

type Message struct {
	PlayerGids      []string    `json:"players"`
	GameGid         string      `json:"gameGid"`
	Type            MessageType `json:"type"`
	SourcePlayerGid string      `json:"sourcePlayerGid"`
	Payload         string      `json:"payload"`
}
