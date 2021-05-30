package main

import (
	"log"
	"net/http"

	"github.com/APIGateway/globals"
	routing "github.com/APIGateway/router"
	"github.com/gorilla/mux"
)

func main() {
	server := &http.Server{
		Addr:    ":" + globals.Port,
		Handler: getPreparedRouter().Router,
	}
	log.Println("[API GATEWAY] Listening on " + globals.Port)
	server.ListenAndServe()
}

func getPreparedRouter() routing.Router {
	router := routing.Router{Router: &mux.Router{}}
	router.Initialize()
	return router
}
