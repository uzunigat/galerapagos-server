package model

type Game struct {
	Gid            string        `json:"gid" sortable:""`
	Status         string        `json:"status" sortable:""`
	RaftLevel      int           `json:"raftLevel" sortable:""`
	WeatherCards   []WeatherCard `json:"weatherCards" sortable:""`
	FoodResources  int           `json:"foodLevel" sortable:""`
	WaterResources int           `json:"waterLevel" sortable:""`
	CreatedAt      string        `json:"createdAt" sortable:""`
	UpdatedAt      string        `json:"updatedAt" sortable:""`
}
