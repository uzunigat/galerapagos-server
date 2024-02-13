package model

type Game struct {
	Gid       string `json:"gid" sortable:""`
	Status    string `json:"status" sortable:""`
	CreatedAt string `json:"createdAt" sortable:""`
	UpdatedAt string `json:"updatedAt" sortable:""`
}
