package models

type Weather struct {
	CloudPercentage int     `json:"cloud_pct"`
	Temperature     int     `json:"temp"`
	FeelsLike       int     `json:"feels_like"`
	Humidity        int     `json:"humidity"`
	MinTemperature  int     `json:"min_temp"`
	MaxTemperature  int     `json:"max_temp"`
	WindSpeed       float64 `json:"wind_speed"`
	WindDegrees     int     `json:"wind_degrees"`
	Sunrise         int     `json:"sunrise"`
	Sunset          int     `json:"sunset"`
}
