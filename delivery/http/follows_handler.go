package http

import (
	"github.com/APIGateway/globals"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type FollowsHandler struct {
	UsersURL string
}

func(handler *FollowsHandler) CreateFollowRequest(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r);
	receiver_id, _ := strconv.ParseUint(params["receiver_id"], 10, 64)

	client := &http.Client{}
	req, _ := http.NewRequest("POST", handler.UsersURL + "/api/follows/sendFollowRequest/" + strconv.FormatUint(receiver_id, 10), nil)
	req.Header.Set("Authorization", r.Header.Get("Authorization"))
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return;
	}
	writeResponse(&w, resp);
}

func (handler *FollowsHandler) AcceptFollowRequest(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r);
	receiver_id, _ := strconv.ParseUint(params["sender_id"], 10, 64)

	client := &http.Client{}
	req, _ := http.NewRequest("POST", handler.UsersURL + "/api/follows/acceptFollowRequest/" + strconv.FormatUint(receiver_id, 10), nil)
	req.Header.Set("Authorization", r.Header.Get("Authorization"))
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return;
	}
	writeResponse(&w, resp);
}

func (handler *FollowsHandler) RejectFollowRequest(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r);
	receiver_id, _ := strconv.ParseUint(params["sender_id"], 10, 64)

	client := &http.Client{}
	req, _ := http.NewRequest("POST", handler.UsersURL + "/api/follows/rejectFollowRequest/" + strconv.FormatUint(receiver_id, 10), nil)
	req.Header.Set("Authorization", r.Header.Get("Authorization"))
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return;
	}
	writeResponse(&w, resp);
}

func (handler *FollowsHandler) IsFollowing(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r);
	id, _ := strconv.ParseUint(params["id"], 10, 64)

	client := &http.Client{}
	req, _ := http.NewRequest("GET", handler.UsersURL + "/api/follows/isFollowing/" + strconv.FormatUint(id, 10), nil)
	req.Header.Set("Authorization", r.Header.Get("Authorization"))
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return;
	}
	writeResponse(&w, resp);
}



func NewFollowsHandler() *FollowsHandler {
	return &FollowsHandler{UsersURL: globals.GetUsersMicroserviceUrl()}
}