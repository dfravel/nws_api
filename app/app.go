package app

import (
	"log"
	"nws_api/config"
	"nws_api/controllers"
	"nws_api/logger"
	"os"
)

var server = controllers.Server{}

// StartApplication .
func StartApplication() {

	config, err := config.LoadConfig()
	if err != nil {
		log.Panic("can't load the config file", err)

	}

	os.Setenv("GEOCODE_AUTH_TOKEN", config.Geocode.AuthToken)
	os.Setenv("GEOCODE_API_URL", config.Geocode.ApiURL)
	os.Setenv("FORECAST_API_URL", config.Forecast.ApiURL)

	server.Initialize(
		config.Server.Debug,
	)

	logger.Info("about to start the application")
	server.Run(config.Server.Port)

}
