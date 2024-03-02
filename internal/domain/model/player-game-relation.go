package model

type PlayerGameRelation struct {
	Gid       string                   `json:"gid"`
	PlayerGid string                   `json:"playerGid"`
	GameGid   string                   `json:"gameGid"`
	Status    PlayerGameRelationStatus `json:"status"`
}
