package messaging

type MessageType string

const (
	MessageTypePlayerJoinedGame MessageType = "PLAYER_JOINED_GAME"
	MessageTypePlayerLeftGame   MessageType = "PLAYER_LEFT_GAME"
	MessageTypeGameStarted      MessageType = "GAME_STARTED"
)

var MessageTypes = []MessageType{
	MessageTypePlayerJoinedGame,
	MessageTypePlayerLeftGame,
	MessageTypeGameStarted,
}

func (MessageType) Enum() []interface{} {
	enums := []interface{}{}
	for _, element := range MessageTypes {
		enums = append(enums, element)
	}
	return enums
}
