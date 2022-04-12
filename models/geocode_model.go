package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

type GeocodeResponse struct {
	Success  bool `json:"success"`
	Location []struct {
		ZipCode      string `json:"zipCode"`
		Country      string `json:"country"`
		CountryCode2 string `json:"countryCode2"`
		CountryCode3 string `json:"countryCode3"`
		State        string `json:"state"`
		StateCode2   string `json:"stateCode2"`
		City         string `json:"city"`
		County       string `json:"county"`
		Latitude     string `json:"latitude"`
		Longitude    string `json:"longitude"`
		TimeZone     string `json:"timeZone"`
	} `json:"location"`
}

func (gc *GeocodeResponse) GetGeocodeByZipCode(zipCode string) (*GeocodeResponse, error) {

	// build the URL that we'll call from https://thezipcodes.com
	// the querystring parameters are zipCode, countryCode and apiKey
	// zipCode is an input parameter, countryCode will be US and apiKey is saved in the config
	host := os.Getenv("GEOCODE_API_URL")
	apiKey := os.Getenv("GEOCODE_AUTH_TOKEN")
	urlParams := fmt.Sprintf("search?zipCode=%s&countryCode=US&apiKey=%s", zipCode, apiKey)

	baseURL, err := url.Parse(host + urlParams)
	if err != nil {
		return &GeocodeResponse{}, err
	}

	req, err := http.NewRequest("GET", baseURL.String(), nil)
	if err != nil {
		return &GeocodeResponse{}, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return &GeocodeResponse{}, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return &GeocodeResponse{}, err
	}

	var geocodeResponseObject GeocodeResponse
	json.Unmarshal(body, &geocodeResponseObject)

	return &geocodeResponseObject, err
}
