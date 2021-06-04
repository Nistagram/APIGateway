package router

import (
	httpHandler "github.com/APIGateway/delivery/http"
	"github.com/APIGateway/delivery/http/middleware"
	"github.com/gorilla/mux"
)

type Router struct {
	Router *mux.Router
}

func (r *Router) Initialize() {
	r.Router.Use(middleware.JSONMiddleware)
	registerUserHandler := httpHandler.NewRegisteredUserHandler()
	searchHandler := httpHandler.NewSearchHandler()

	r.Router.HandleFunc("/api/users", registerUserHandler.GetAll).Methods("GET")

	r.Router.HandleFunc("/api/search", searchHandler.Search).Methods("GET")
}
