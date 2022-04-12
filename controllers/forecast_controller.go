package controllers

import (
	"github.com/gorilla/mux"
	"net/http"
	"nws_api/logger"
	"nws_api/models"
	"nws_api/responses"
)

// GetGeocode returns the longitude and latitude for a US Zip code. Not required for the task, but leaving it in here for testing
func (server *Server) GetGeocode(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	zipCode := vars["zipCode"]

	geocode := models.GeocodeResponse{}

	logger.GetLogger().Printf("%s:%s", "getting geocode by US zip code", zipCode)
	geoCodeResponse, err := geocode.GetGeocodeByZipCode(zipCode)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, geoCodeResponse)
}

// GetPointsByZipCode returns the point information for a US Zip Code - all forecast URLs are part of the response
func (server *Server) GetPointsByZipCode(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	zipCode := vars["zipCode"]

	geocode := models.GeocodeResponse{}

	// convert the zip code to a geocode object
	// this will allow us to grab that latitude and longitude of the zip code
	logger.GetLogger().Printf("%s:%s", "getting geocode by US zip code", zipCode)
	geoCodeResponse, err := geocode.GetGeocodeByZipCode(zipCode)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	logger.GetLogger().Printf("%s:%s", "getting points by US zip code", zipCode)

	point := models.PointResponse{}
	pointResponse, err := point.GetForecastByLatLong(geoCodeResponse.Location[0].Latitude, geoCodeResponse.Location[0].Longitude)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, pointResponse)
}

// GetForecast will run all the required functions to return a multi-day forecast for a given US zipcode
func (server *Server) GetForecast(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	zipCode := vars["zipCode"]

	geocode := models.GeocodeResponse{}

	// convert the zip code to a geocode object
	// this will allow us to grab that latitude and longitude of the zip code
	geoCodeResponse, err := geocode.GetGeocodeByZipCode(zipCode)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	// confirm what we're getting back from the Geocode function
	// fmt.Printf("%s: %v", "GeoCode Response", geoCodeResponse)

	point := models.PointResponse{}
	pointResponse, err := point.GetForecastByLatLong(geoCodeResponse.Location[0].Latitude, geoCodeResponse.Location[0].Longitude)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	// print the API Urls to the console to confirm we have accurate URLs
	logger.GetLogger().Printf("%s:%s", "Forecast Hourly URL", pointResponse.Properties.ForecastHourly)
	logger.GetLogger().Printf("%s:%s", "Forecast URL", pointResponse.Properties.Forecast)

	forecast := models.ForecastResponse{}

	forecastResponse, err := forecast.GetForecast(pointResponse.Properties.Forecast)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, forecastResponse)
}

// GetForecastHourly will run all the required functions to return an hourly forecast for a given US zipcode
func (server *Server) GetForecastHourly(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	zipCode := vars["zipCode"]

	geocode := models.GeocodeResponse{}

	// convert the zip code to a geocode object
	// this will allow us to grab that latitude and longitude of the zip code
	geoCodeResponse, err := geocode.GetGeocodeByZipCode(zipCode)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	// confirm what we're getting back from the Geocode function
	// fmt.Printf("%s: %v", "GeoCode Response", geoCodeResponse)

	point := models.PointResponse{}
	pointResponse, err := point.GetForecastByLatLong(geoCodeResponse.Location[0].Latitude, geoCodeResponse.Location[0].Longitude)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	// print the API Urls to the console to confirm we have accurate URLs
	logger.GetLogger().Printf("%s:%s", "Forecast Hourly URL", pointResponse.Properties.ForecastHourly)
	logger.GetLogger().Printf("%s:%s", "Forecast URL", pointResponse.Properties.Forecast)

	forecast := models.ForecastResponse{}

	forecastResponse, err := forecast.GetForecastByHour(pointResponse.Properties.ForecastHourly)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, forecastResponse)
}
