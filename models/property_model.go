package models

import "time"

type Property struct {
	Updated           time.Time `json:"updated"`
	Units             string    `json:"units"`
	ForecastGenerator string    `json:"forecastGenerator"`
	GeneratedAt       time.Time `json:"generatedAt"`
	UpdateTime        time.Time `json:"updateTime"`
	ValidTimes        string    `json:"validTimes"`
	Elevation         Elevation `json:"elevation"`
	Periods           []Period  `json:"periods"`
}
