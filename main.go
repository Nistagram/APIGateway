package main

import (
	routing "github.com/APIGateway/router"
	"github.com/gorilla/mux"
	"github.com/APIGateway/globals"
	"log"
	"net/http"
)

func main() {
	server := &http.Server{
		Addr:              ":" + globals.ServerPort,
		Handler:           getPreparedRouter().Router,
	}
	log.Println("Listening...")
	server.ListenAndServe()
}

func getPreparedRouter() routing.Router{
	router := routing.Router{Router: &mux.Router{}}
	router.Initialize()
	return router
}