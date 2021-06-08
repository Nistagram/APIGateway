package router

import (
	httpHandler "github.com/APIGateway/delivery/http"
	"github.com/APIGateway/delivery/http/middleware"
	"github.com/gorilla/mux"
	"net/http"
)

type Router struct {
	Router *mux.Router
}

func (r *Router) Initialize() {
	r.Router.Use(middleware.JSONMiddleware)
	r.Router.Use(middleware.CORSMiddleware)
	registerUserHandler := httpHandler.NewRegisteredUserHandler()
	searchHandler := httpHandler.NewSearchHandler()
	authHandler := httpHandler.NewAuthHandler()
	postHandler := httpHandler.NewPostHandler()

	r.Router.HandleFunc("/api/users", registerUserHandler.GetAll).Methods("GET")

	r.Router.HandleFunc("/api/search", searchHandler.Search).Methods("GET")

	usersSubRouter := r.Router.PathPrefix("/api/users").Subrouter();
	usersSubRouter.PathPrefix("/login").Handler(http.HandlerFunc(authHandler.Login)).Methods("POST", "OPTIONS")
	usersSubRouter.PathPrefix("/register").Handler(http.HandlerFunc(authHandler.Register))

	contentSubRouter := r.Router.PathPrefix("/api/content").Subrouter();
	contentSubRouter.PathPrefix("/post").Handler(http.HandlerFunc(postHandler.GetAll)).Methods("GET", "OPTIONS")
	contentSubRouter.PathPrefix("/post/like").Handler(http.HandlerFunc(postHandler.LikePost)).Methods("POST", "OPTIONS")
	contentSubRouter.PathPrefix("/post/comment").Handler(http.HandlerFunc(postHandler.CommentPost)).Methods("POST", "OPTIONS")
}
