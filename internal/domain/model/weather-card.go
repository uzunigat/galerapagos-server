package model

type WeatherCard struct {
	Gid         string `json:"gid"`
	Name        string `json:"name"`
	WaterLevel  int    `json:"water_level"`
	IsFinalGame bool   `json:"is_final_game"`
	Quantity    int    `json:"quantity"`
}
