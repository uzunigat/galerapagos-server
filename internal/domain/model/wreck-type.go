package model

type WreckCardType string

const (
	RessourceCardType WreckCardType = "RESSOURCE"
	SpecialCardType   WreckCardType = "SPECIAL"
	PermanentCardType WreckCardType = "PERMANENT"
)

var WreckCardTypes = []WreckCardType{
	RessourceCardType,
	SpecialCardType,
	PermanentCardType,
}

func (WreckCardType) Enum() []interface{} {
	enums := []interface{}{}
	for _, element := range WreckCardTypes {
		enums = append(enums, element)
	}
	return enums
}
