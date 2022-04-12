package controllers

import (
	"fmt"
	"log"
	"net/http"
	"nws_api/logger"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// Server .
type Server struct {
	Router *mux.Router
}

// Initialize .
func (server *Server) Initialize(Debug bool) {

	server.Router = mux.NewRouter()

	server.initializeRoutes()

	// if we're in debug mode, walk the routes and print them to the console
	if Debug {
		server.Router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
			tpl, err1 := route.GetPathTemplate()
			met, err2 := route.GetMethods()
			name := route.GetName()
			fmt.Println(tpl, err1, met, err2, name)
			return nil
		})
	}

}

// Run .
func (server *Server) Run(addr string) {
	logger.GetLogger().Printf("Listening to port %s", addr)

	// logger.Printf("Listening to port %s", addr)
	header := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "HEAD", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"*"})

	log.Fatal(http.ListenAndServe(addr, handlers.CORS(header, methods, origins)(server.Router)))
}
