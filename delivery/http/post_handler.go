package http

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/APIGateway/globals"
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

func (handler *PostHandler) UploadMedia(w http.ResponseWriter, r *http.Request) {
	requestURI := handler.ContentURL + "/api/media/upload"
	body, bodyErr := ioutil.ReadAll(r.Body)
	if bodyErr != nil {
		log.Println(bodyErr)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	resp, err := http.Post(requestURI, "application/json", bytes.NewReader(body))
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	writeResponse(&w, resp)
}

func (handler *PostHandler) UploadPost(w http.ResponseWriter, r *http.Request) {
	requestURI := handler.ContentURL + "/api/post/upload"
	body, bodyErr := ioutil.ReadAll(r.Body)
	if bodyErr != nil {
		log.Println(bodyErr)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	resp, err := http.Post(requestURI, "application/json", bytes.NewReader(body))
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	writeResponse(&w, resp)
}

func (handler *PostHandler) UploadStory(w http.ResponseWriter, r *http.Request) {
	requestURI := handler.ContentURL + "/api/story/upload"
	body, bodyErr := ioutil.ReadAll(r.Body)
	if bodyErr != nil {
		log.Println(bodyErr)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	resp, err := http.Post(requestURI, "application/json", bytes.NewReader(body))
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	writeResponse(&w, resp)
}

func (handler *PostHandler) LikePost(w http.ResponseWriter, r *http.Request) {
	requestURI := handler.ContentURL + "/api/post/like"
	body, bodyErr := ioutil.ReadAll(r.Body)
	if bodyErr != nil {
		log.Println(bodyErr)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	resp, err := http.Post(requestURI, "application/json", bytes.NewReader(body))
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	writeResponse(&w, resp)
}

func (handler *PostHandler) CommentPost(w http.ResponseWriter, r *http.Request) {
	requestURI := handler.ContentURL + "/api/post/comment"
	body, bodyErr := ioutil.ReadAll(r.Body)
	if bodyErr != nil {
		log.Println(bodyErr)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	resp, err := http.Post(requestURI, "application/json", bytes.NewReader(body))
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	writeResponse(&w, resp)
}
