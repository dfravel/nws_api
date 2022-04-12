package config

import "github.com/spf13/viper"

// Configurations .
type Configurations struct {
	Server   ServerConfigurations
	Geocode  GeocodeConfigurations
	Site     SiteConfigurations
	Logging  LogConfigurations
	Forecast ForecastConfigurations
}

// SiteConfigurations .
type SiteConfigurations struct {
	Title string
}

// GeocodeConfigurations allows us to accept a zip code from the end user.
// We're using a 3rd party platform (geocode.xyz) to convert the zip code to latitude/longitude values .
type GeocodeConfigurations struct {
	ApiURL    string
	AuthToken string
}

// ForecastConfigurations manages the connection to the NWS API endpoint
type ForecastConfigurations struct {
	ApiURL string
}

// LogConfigurations allows us to define the default level of logging
type LogConfigurations struct {
	LogLevel  string
	LogOutput string
}

// ServerConfigurations for the port that we're running the API on and weather or not we're running in debug mode
type ServerConfigurations struct {
	Port  string
	Debug bool
}

// LoadConfig reads in the config.yaml file at the root
func LoadConfig() (config Configurations, err error) {
	// Set the file name of the configurations file
	viper.SetConfigName("config")

	// Set the path to look for the configurations file
	viper.AddConfigPath(".")

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	// We want to read in any yaml files as our config files
	viper.SetConfigType("yaml")

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return

}
