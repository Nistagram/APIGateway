package http

import (
	"bytes"
	"github.com/APIGateway/globals"
	"io/ioutil"
	"log"
	"net/http"
)

type PostHandler struct {
	ContentURL string
}

func NewPostHandler() *PostHandler {
	return &PostHandler{ContentURL: globals.GetContentMicroserviceUrl()}
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

func (handler *PostHandler) LikePost(w http.ResponseWriter, r *http.Request) {
	requestURI := handler.ContentURL + "/api/post/like"
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