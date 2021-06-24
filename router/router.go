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
	verificationCategoryHandler := httpHandler.NewVerificationCategoryHandler()
	businessCategoryHandler := httpHandler.NewBusinessCategoryHandlerHandler()
	verificationHandler := httpHandler.NewVerificationHandler()
	mediaHandler := httpHandler.NewMediaHandler()
	feedHandler := httpHandler.NewFeedHandler()
	reportHandler := httpHandler.NewReportHandler()
	storyHandler := httpHandler.NewStoryHandler()

	r.Router.HandleFunc("/api/users", registerUserHandler.GetAll).Methods("GET", "OPTIONS")
	r.Router.HandleFunc("/api/user", registerUserHandler.Get).Methods("GET", "OPTIONS")
	r.Router.HandleFunc("/api/user/{id:[0-9]+}", registerUserHandler.Delete).Methods("DELETE", "OPTIONS")
	r.Router.HandleFunc("/api/users/config/edit", registerUserHandler.EditConfig).Methods("POST", "OPTIONS")
	r.Router.HandleFunc("/api/search", searchHandler.Search).Methods("GET", "OPTIONS")
	r.Router.HandleFunc("/api/block/user/{user_id:[0-9]+}", registerUserHandler.Block).Methods("GET", "OPTIONS")
	r.Router.HandleFunc("/api/mute/user/{user_id:[0-9]+}", registerUserHandler.Mute).Methods("GET", "OPTIONS")
	r.Router.HandleFunc("/api/unblock/user/{user_id:[0-9]+}", registerUserHandler.Unblock).Methods("GET", "OPTIONS")
	r.Router.HandleFunc("/api/unmute/user/{user_id:[0-9]+}", registerUserHandler.Unmute).Methods("GET", "OPTIONS")
	r.Router.HandleFunc("/api/feed", feedHandler.GetFeedPosts).Methods("GET", "OPTIONS")

	usersSubRouter := r.Router.PathPrefix("/api/users").Subrouter()
	usersSubRouter.PathPrefix("/login").Handler(http.HandlerFunc(authHandler.Login)).Methods("POST", "OPTIONS")
	usersSubRouter.PathPrefix("/register").Handler(http.HandlerFunc(authHandler.Register)).Methods("POST", "OPTIONS")
	usersSubRouter.PathPrefix("/edit").Handler(http.HandlerFunc(registerUserHandler.Edit)).Methods("POST", "OPTIONS")
	usersSubRouter.PathPrefix("/info/{id:[0-9]+}").Handler(http.HandlerFunc(registerUserHandler.GetUserInfoById)).Methods("GET", "OPTIONS")
	usersSubRouter.PathPrefix("/info").Handler(http.HandlerFunc(registerUserHandler.GetUserInfo)).Methods("GET", "OPTIONS")

	contentSubRouter := r.Router.PathPrefix("/api/content").Subrouter()

	contentSubRouter.PathPrefix("/post/is/isLiked").Handler(http.HandlerFunc(postHandler.IsLikedPost)).Methods("POST", "OPTIONS")
	contentSubRouter.PathPrefix("/post/upload").Handler(http.HandlerFunc(postHandler.UploadPost)).Methods("POST", "OPTIONS")
	contentSubRouter.PathPrefix("/media/upload").Handler(http.HandlerFunc(postHandler.UploadMedia)).Methods("POST", "OPTIONS")
	contentSubRouter.PathPrefix("/story/upload").Handler(http.HandlerFunc(postHandler.UploadStory)).Methods("POST", "OPTIONS")
	contentSubRouter.PathPrefix("/story/{id:[0-9]+}").Handler(http.HandlerFunc(storyHandler.Delete)).Methods("DELETE", "OPTIONS")
	contentSubRouter.PathPrefix("/post/user/{user_id:[0-9]+}").Handler(http.HandlerFunc(postHandler.GetByUserId)).Methods("GET", "OPTIONS")
	contentSubRouter.PathPrefix("/post/liked/{id:[0-9]+}").Handler(http.HandlerFunc(postHandler.GetAllLiked)).Methods("GET", "OPTIONS")
	contentSubRouter.PathPrefix("/post/taggedIn/{id:[0-9]+}").Handler(http.HandlerFunc(postHandler.GetAllTaggedIn)).Methods("GET", "OPTIONS")
	contentSubRouter.PathPrefix("/post/{id:[0-9]+}").Handler(http.HandlerFunc(postHandler.GetById)).Methods("GET", "OPTIONS")
	contentSubRouter.PathPrefix("/post/{id:[0-9]+}").Handler(http.HandlerFunc(postHandler.Delete)).Methods("DELETE", "OPTIONS")
	contentSubRouter.PathPrefix("/post").Handler(http.HandlerFunc(postHandler.GetAll)).Methods("GET", "OPTIONS")
	contentSubRouter.PathPrefix("/post/like").Handler(http.HandlerFunc(postHandler.LikePost)).Methods("POST", "OPTIONS")
	contentSubRouter.PathPrefix("/post/comment").Handler(http.HandlerFunc(postHandler.CommentPost)).Methods("POST", "OPTIONS")
	contentSubRouter.PathPrefix("/post/report").Handler(http.HandlerFunc(postHandler.ReportPost)).Methods("POST", "OPTIONS")
	//contentSubRouter.HandleFunc("/api/media/{id:[0-9]+}", mediaHandler.Get).Methods("GET", "OPTIONS")
	contentSubRouter.PathPrefix("/media/{id:[0-9]+}").Handler(http.HandlerFunc(mediaHandler.Get)).Methods("GET", "OPTIONS")
	followsSubRouter := r.Router.PathPrefix("/api/follows").Subrouter()
	followsSubRouter.Path("/sendFollowRequest/{receiver_id:[0-9]+}").Handler(http.HandlerFunc(followsHandler.CreateFollowRequest)).Methods("POST", "OPTIONS")
	followsSubRouter.Path("/acceptFollowRequest/{sender_id:[0-9]+}").Handler(http.HandlerFunc(followsHandler.AcceptFollowRequest)).Methods("POST", "OPTIONS")
	followsSubRouter.Path("/rejectFollowRequest/{sender_id:[0-9]+}").Handler(http.HandlerFunc(followsHandler.RejectFollowRequest)).Methods("POST", "OPTIONS")
	followsSubRouter.Path("/unfollow/{id:[0-9]+}").Handler(http.HandlerFunc(followsHandler.Unfollow)).Methods("POST", "OPTIONS")
	followsSubRouter.HandleFunc("/isFollowing/{id:[0-9]+}", followsHandler.IsFollowing).Methods("GET", "OPTIONS")
	followsSubRouter.HandleFunc("/isFollowRequestPending/{id:[0-9]+}", followsHandler.IsFollowRequestPending).Methods("GET", "OPTIONS")
	followsSubRouter.HandleFunc("/withdrawFollowRequest/{id:[0-9]+}", followsHandler.WithdrawFollowRequest).Methods("POST", "OPTIONS")
	verificationSubRouter := r.Router.PathPrefix("/api/verification").Subrouter()
	verificationSubRouter.Path("/category").Handler(http.HandlerFunc(verificationCategoryHandler.GetAll))
	verificationSubRouter.Path("/verify").Handler(http.HandlerFunc(verificationHandler.CreateVerificationRequest))
	verificationSubRouter.Path("/status/{id}").Handler(http.HandlerFunc(verificationHandler.GetVerificationStatus)).Methods("GET", "OPTIONS")
	verificationSubRouter.Path("/accept/{id}").Handler(http.HandlerFunc(verificationHandler.Accept)).Methods("POST", "OPTIONS")
	verificationSubRouter.Path("/reject/{id}").Handler(http.HandlerFunc(verificationHandler.Reject)).Methods("POST", "OPTIONS")
	verificationSubRouter.Path("/unresolved").Handler(http.HandlerFunc(verificationHandler.GetUnresolved)).Methods("GET", "OPTIONS")
	businessCategorySubRouter := r.Router.PathPrefix("/api/business-category").Subrouter()
	businessCategorySubRouter.Path("").Handler(http.HandlerFunc(businessCategoryHandler.GetAll))

	r.Router.HandleFunc("/api/report-types", reportHandler.GetAllTypes).Methods("GET", "OPTIONS")
	r.Router.HandleFunc("/api/report", reportHandler.GetAll).Methods("GET", "OPTIONS")
	r.Router.HandleFunc("/api/report/{id:[0-9]+}", reportHandler.Delete).Methods("DELETE", "OPTIONS")
}
