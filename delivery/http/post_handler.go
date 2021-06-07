package http

import (
	"bytes"
	"github.com/APIGateway/globals"
	"io/ioutil"
	"log"
	"net/http"
)

type ContentHandler struct {
	ContentURL string
}

func NewContentHandler() *ContentHandler {
	return &ContentHandler{ContentURL: globals.GetContentMicroserviceUrl()}
}

func (handler *ContentHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	requestURI := handler.ContentURL + "/api/post"
	resp, err := http.Get(requestURI)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		WriteResponse(&w, resp)
	}
}

func (handler *ContentHandler) LikePost(w http.ResponseWriter, r *http.Request) {
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

func (handler *ContentHandler) CommentPost(w http.ResponseWriter, r *http.Request) {
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