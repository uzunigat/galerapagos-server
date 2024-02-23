package model

type GameStatus string

const (
	GameCreated    GameStatus = "CREATED"
	GameInProgress GameStatus = "IN_PROGRESS"
	GameFinished   GameStatus = "FINISHED"
)

var GameStatuses = []GameStatus{
	GameCreated,
}

func (GameStatus) Enum() []interface{} {
	enums := []interface{}{}
	for _, element := range GameStatuses {
		enums = append(enums, element)
	}
	return enums
}
