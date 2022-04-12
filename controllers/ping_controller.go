package controllers

import (
	"net/http"

	"nws_api/responses"
)

// Ping checks to see if the API is up and running
func (server *Server) Ping(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Pong")

}
