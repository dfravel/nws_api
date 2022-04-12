package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

type PointResponse struct {
	Context    []interface{} `json:"@context"`
	ID         string        `json:"id"`
	Type       string        `json:"type"`
	Geometry   Geometry      `json:"geometry"`
	Properties struct {
		ID                  string `json:"@id"`
		Type                string `json:"@type"`
		Cwa                 string `json:"cwa"`
		ForecastOffice      string `json:"forecastOffice"`
		GridID              string `json:"gridId"`
		GridX               int    `json:"gridX"`
		GridY               int    `json:"gridY"`
		Forecast            string `json:"forecast"`
		ForecastHourly      string `json:"forecastHourly"`
		ForecastGridData    string `json:"forecastGridData"`
		ObservationStations string `json:"observationStations"`
		RelativeLocation    struct {
			Type     string `json:"type"`
			Geometry struct {
				Type        string    `json:"type"`
				Coordinates []float64 `json:"coordinates"`
			} `json:"geometry"`
			Properties struct {
				City     string `json:"city"`
				State    string `json:"state"`
				Distance struct {
					UnitCode string  `json:"unitCode"`
					Value    float64 `json:"value"`
				} `json:"distance"`
				Bearing struct {
					UnitCode string `json:"unitCode"`
					Value    int    `json:"value"`
				} `json:"bearing"`
			} `json:"properties"`
		} `json:"relativeLocation"`
		ForecastZone    string `json:"forecastZone"`
		County          string `json:"county"`
		FireWeatherZone string `json:"fireWeatherZone"`
		TimeZone        string `json:"timeZone"`
		RadarStation    string `json:"radarStation"`
	} `json:"properties"`
}

type ForecastResponse struct {
	Context    []interface{} `json:"@context"`
	Type       string        `json:"type"`
	Geometry   Geometry      `json:"geometry"`
	Properties Property      `json:"properties"`
}

func (p *PointResponse) GetForecastByLatLong(latitude string, longitude string) (*PointResponse, error) {

	host := os.Getenv("FORECAST_API_URL")
	urlParams := fmt.Sprintf("points/%s,%s", latitude, longitude)

	baseURL, err := url.Parse(host + urlParams)
	if err != nil {
		return &PointResponse{}, err
	}

	req, err := http.NewRequest("GET", baseURL.String(), nil)
	if err != nil {
		return &PointResponse{}, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return &PointResponse{}, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return &PointResponse{}, err
	}

	var pointResponseObject PointResponse
	json.Unmarshal(body, &pointResponseObject)

	return &pointResponseObject, err
}

func (f *ForecastResponse) GetForecast(forecastURL string) (*ForecastResponse, error) {

	req, err := http.NewRequest("GET", forecastURL, nil)
	if err != nil {
		return &ForecastResponse{}, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return &ForecastResponse{}, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return &ForecastResponse{}, err
	}

	var forecastResponseObject ForecastResponse
	json.Unmarshal(body, &forecastResponseObject)

	return &forecastResponseObject, err
}

func (f *ForecastResponse) GetForecastByHour(hourlyURL string) (*ForecastResponse, error) {

	req, err := http.NewRequest("GET", hourlyURL, nil)
	if err != nil {
		return &ForecastResponse{}, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return &ForecastResponse{}, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return &ForecastResponse{}, err
	}

	var forecastResponseObject ForecastResponse
	json.Unmarshal(body, &forecastResponseObject)

	return &forecastResponseObject, err
}
