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
	followsHandler := httpHandler.NewFollowsHandler()

	r.Router.HandleFunc("/api/users", registerUserHandler.GetAll).Methods("GET", "OPTIONS")
	r.Router.HandleFunc("/api/user", registerUserHandler.Get).Methods("GET", "OPTIONS")
	r.Router.HandleFunc("/api/users/config/edit", registerUserHandler.EditConfig).Methods("POST", "OPTIONS")
	r.Router.HandleFunc("/api/search", searchHandler.Search).Methods("GET", "OPTIONS")

	usersSubRouter := r.Router.PathPrefix("/api/users").Subrouter()
	usersSubRouter.PathPrefix("/login").Handler(http.HandlerFunc(authHandler.Login)).Methods("POST", "OPTIONS")
	usersSubRouter.PathPrefix("/register").Handler(http.HandlerFunc(authHandler.Register)).Methods("POST", "OPTIONS")
	usersSubRouter.PathPrefix("/edit").Handler(http.HandlerFunc(registerUserHandler.Edit)).Methods("POST", "OPTIONS")
	usersSubRouter.PathPrefix("/info/{id:[0-9]+}").Handler(http.HandlerFunc(registerUserHandler.GetUserInfoById)).Methods("GET", "OPTIONS")
	usersSubRouter.PathPrefix("/info").Handler(http.HandlerFunc(registerUserHandler.GetUserInfo)).Methods("GET", "OPTIONS")

	contentSubRouter := r.Router.PathPrefix("/api/content").Subrouter()

	contentSubRouter.PathPrefix("/post/upload").Handler(http.HandlerFunc(postHandler.UploadPost)).Methods("POST", "OPTIONS")
	contentSubRouter.PathPrefix("/media/upload").Handler(http.HandlerFunc(postHandler.UploadMedia)).Methods("POST", "OPTIONS")
	contentSubRouter.PathPrefix("/story/upload").Handler(http.HandlerFunc(postHandler.UploadStory)).Methods("POST", "OPTIONS")
	contentSubRouter.PathPrefix("/post/{user_id:[0-9]+}").Handler(http.HandlerFunc(postHandler.GetByUserId)).Methods("GET", "OPTIONS")
	contentSubRouter.PathPrefix("/post").Handler(http.HandlerFunc(postHandler.GetAll)).Methods("GET", "OPTIONS")
	contentSubRouter.PathPrefix("/post/like").Handler(http.HandlerFunc(postHandler.LikePost)).Methods("POST", "OPTIONS")
	contentSubRouter.PathPrefix("/post/comment").Handler(http.HandlerFunc(postHandler.CommentPost)).Methods("POST", "OPTIONS")

	followsSubRouter := r.Router.PathPrefix("/api/follows").Subrouter()
	followsSubRouter.Path("/sendFollowRequest/{receiver_id:[0-9]+}").Handler(http.HandlerFunc(followsHandler.CreateFollowRequest)).Methods("POST", "OPTIONS")
	followsSubRouter.Path("/acceptFollowRequest/{sender_id:[0-9]+}").Handler(http.HandlerFunc(followsHandler.AcceptFollowRequest)).Methods("POST", "OPTIONS")
	followsSubRouter.Path("/rejectFollowRequest/{sender_id:[0-9]+}").Handler(http.HandlerFunc(followsHandler.RejectFollowRequest)).Methods("POST", "OPTIONS")
}
