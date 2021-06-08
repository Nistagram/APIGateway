package router

import (
	"net/http"

	httpHandler "github.com/APIGateway/delivery/http"
	"github.com/APIGateway/delivery/http/middleware"
	"github.com/gorilla/mux"
)

type Router struct {
	Router *mux.Router
}

func (r *Router) Initialize() {
	r.Router.Use(middleware.CORSMiddleware)
	r.Router.Use(middleware.JSONMiddleware)
	registerUserHandler := httpHandler.NewRegisteredUserHandler()
	searchHandler := httpHandler.NewSearchHandler()
	authHandler := httpHandler.NewAuthHandler()
	postHandler := httpHandler.NewPostHandler()

	r.Router.HandleFunc("/api/users", registerUserHandler.GetAll).Methods("GET")
	r.Router.HandleFunc("/api/user", registerUserHandler.Get).Methods("GET", "OPTIONS")

	r.Router.HandleFunc("/api/search", searchHandler.Search).Methods("GET")

	usersSubRouter := r.Router.PathPrefix("/api/users").Subrouter()
	usersSubRouter.PathPrefix("/login").Handler(http.HandlerFunc(authHandler.Login)).Methods("POST", "OPTIONS")
	usersSubRouter.PathPrefix("/register").Handler(http.HandlerFunc(authHandler.Register))
	usersSubRouter.PathPrefix("/edit").Handler(http.HandlerFunc(registerUserHandler.Edit)).Methods("POST", "OPTIONS")

	contentSubRouter := r.Router.PathPrefix("/api/content").Subrouter()
	contentSubRouter.PathPrefix("/post").Handler(http.HandlerFunc(postHandler.GetAll)).Methods("GET")
	contentSubRouter.PathPrefix("/post/like").Handler(http.HandlerFunc(postHandler.GetAll)).Methods("POST", "OPTIONS")
	contentSubRouter.PathPrefix("/post/comment").Handler(http.HandlerFunc(postHandler.GetAll)).Methods("POST", "OPTIONS")
	contentSubRouter.PathPrefix("/post/upload").Handler(http.HandlerFunc(postHandler.UploadPost)).Methods("POST", "OPTIONS")
	contentSubRouter.PathPrefix("/media/upload").Handler(http.HandlerFunc(postHandler.UploadMedia)).Methods("POST", "OPTIONS")
	contentSubRouter.PathPrefix("/story/upload").Handler(http.HandlerFunc(postHandler.UploadStory)).Methods("POST", "OPTIONS")
}
