package controllers

import "nws_api/middlewares"

func (server *Server) initializeRoutes() {

	// Home Routes
	server.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(server.Base)).Methods("GET").Name("Base URL")
	server.Router.HandleFunc("/api/v1", middlewares.SetMiddlewareJSON(server.Home)).Methods("GET").Name("API Base URl")

	// Ping Route
	server.Router.HandleFunc("/api/v1/ping", middlewares.SetMiddlewareJSON(server.Ping)).Methods("GET").Name("Ping Health Check")

	// Geocode Route
	server.Router.HandleFunc("/api/v1/geocode/{zipCode}", middlewares.SetMiddlewareJSON(server.GetGeocode)).Methods("GET").Name("Get Longitude and Latitude")

	// Forecast Routes
	server.Router.HandleFunc("/api/v1/points/{zipCode}", middlewares.SetMiddlewareJSON(server.GetPointsByZipCode)).Methods("GET").Name("Get Forecast by Zip Code")
	server.Router.HandleFunc("/api/v1/forecast/{zipCode}", middlewares.SetMiddlewareJSON(server.GetForecast)).Methods("GET").Name("Get Forecast by Zip Code")

	// Hourly Forecast Route
	server.Router.HandleFunc("/api/v1/forecast/hourly/{zipCode}", middlewares.SetMiddlewareJSON(server.GetForecastHourly)).Methods("GET").Name("Get Hourly Forecast by Zip Code")

}
