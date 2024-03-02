package model

type WreckCard struct {
	Gid         string        `json:"gid"`
	Name        string        `json:"name"`
	Type        WreckCardType `json:"type"`
	Description string        `json:"description"`
	IsDiscarded bool          `json:"isDiscarded"`
}
