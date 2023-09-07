package models

type Store struct {
	ID         int     `json:"id"`
	Store_name string  `json:"store_name"`
	Rating     int     `json:"rating"`
	Longitude  float64 `json:"longitude"`
	Latitude   float64 `json:"latitude"`
	Address    string  `json:"address"`
}
