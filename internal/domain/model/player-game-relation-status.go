package model

type PlayerGameRelationStatus string

const (
	PlayerGameRelationConnected PlayerGameRelationStatus = "CREATED"
)

var PlayerGameRelationStatuses = []PlayerGameRelationStatus{
	PlayerGameRelationConnected,
}

func (PlayerGameRelationStatus) Enum() []interface{} {
	enums := []interface{}{}
	for _, element := range PlayerGameRelationStatuses {
		enums = append(enums, element)
	}
	return enums
}
