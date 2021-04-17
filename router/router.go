package router


import (
	"github.com/APIGateway/delivery/http/middleware"
	"github.com/gorilla/mux"
	httpHandler "github.com/APIGateway/delivery/http"
)

type Router struct {
	Router *mux.Router
}

func (r *Router) Initialize(){
	r.Router.Use(middleware.JSONMiddleware)
	registerUserHandler := httpHandler.NewRegisteredUserHandler()
	r.Router.HandleFunc("/api/users", registerUserHandler.GetAll).Methods("GET")
}