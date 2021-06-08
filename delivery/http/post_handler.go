package http

import (
	"bytes"
	json2 "encoding/json"
	"github.com/APIGateway/globals"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

type PostHandler struct {
	ContentURL string
	UsersURL string
}

func NewPostHandler() *PostHandler {
	return &PostHandler{ContentURL: globals.GetContentMicroserviceUrl(), UsersURL: globals.GetUsersMicroserviceUrl()}
}

func (handler *PostHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	requestURI := handler.ContentURL + "/api/post"
	log.Println("Request URI: " + requestURI)
	resp, err := http.Get(requestURI)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		WriteResponse(&w, resp)
	}
}


func (handler *PostHandler) getUserIdFromSession(token string) (uint64, error){
	requestURI := handler.UsersURL + "/api/userId"
	client := &http.Client{}
	req, _ := http.NewRequest("GET", requestURI, nil)
	req.Header.Set("Authorization", token)
	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}

	var UserId uint64
	json2.NewDecoder(resp.Body).Decode(&UserId)
	log.Print("JSON Data: ")
	log.Println(UserId)

	return UserId, nil
}

func (handler *PostHandler) LikePost(w http.ResponseWriter, r *http.Request) {
	bearToken := r.Header.Get("Authorization")
	userId, err := handler.getUserIdFromSession(bearToken)

	params := url.Values{}
	params.Add("userId", strconv.Itoa(int(userId)))

	requestURI := handler.ContentURL + "/api/post/like?" + params.Encode()
	body, bodyErr := ioutil.ReadAll(r.Body);
	if bodyErr != nil {
		log.Println(bodyErr)
		w.WriteHeader(http.StatusInternalServerError)
		return;
	}
	resp, err := http.Post(requestURI, "application/json", bytes.NewReader(body));
	log.Println(resp)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return;
	}
	writeResponse(&w, resp);
}

func (handler *PostHandler) CommentPost(w http.ResponseWriter, r *http.Request) {
	requestURI := handler.ContentURL + "/api/post/comment"
	body, bodyErr := ioutil.ReadAll(r.Body);
	if bodyErr != nil {
		log.Println(bodyErr)
		w.WriteHeader(http.StatusInternalServerError)
		return;
	}
	resp, err := http.Post(requestURI, "application/json", bytes.NewReader(body));
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return;
	}
	writeResponse(&w, resp);
}