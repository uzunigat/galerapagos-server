package messaging

type MessageType string

const (
	PlayerJoinedGame MessageType = "PLAYER_JOINED_GAME"
	PlayerLeftGame   MessageType = "PLAYER_LEFT_GAME"
)

var MessageTypes = []MessageType{
	PlayerJoinedGame,
	PlayerLeftGame,
}

func (MessageType) Enum() []interface{} {
	enums := []interface{}{}
	for _, element := range MessageTypes {
		enums = append(enums, element)
	}
	return enums
}
