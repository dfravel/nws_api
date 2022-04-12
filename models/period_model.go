package models

type Period struct {
	Number           int         `json:"number"`
	Name             string      `json:"name"`
	StartTime        string      `json:"startTime"`
	EndTime          string      `json:"endTime"`
	IsDaytime        bool        `json:"isDaytime"`
	Temperature      int         `json:"temperature"`
	TemperatureUnit  string      `json:"temperatureUnit"`
	TemperatureTrend interface{} `json:"temperatureTrend"`
	WindSpeed        string      `json:"windSpeed"`
	WindDirection    string      `json:"windDirection"`
	Icon             string      `json:"icon"`
	ShortForecast    string      `json:"shortForecast"`
	DetailedForecast string      `json:"detailedForecast"`
}
