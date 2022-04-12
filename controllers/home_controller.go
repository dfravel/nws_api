package controllers

import (
	"net/http"

	"nws_api/responses"
)

// Home welcomes you to the API
func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome To the CrossnoKaye National Weather Service API")
}


// Base welcomes you to the Base URL
func (server *Server) Base(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome To the CrossnoKaye Base URL")
}