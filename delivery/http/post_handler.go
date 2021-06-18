package http

import (
	"bytes"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/APIGateway/globals"

	json2 "encoding/json"

	"net/url"
	"strconv"
)

type PostHandler struct {
	ContentURL string
	UsersURL   string
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

func (handler *PostHandler) GetAllLiked(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r);
	id, _ := strconv.ParseUint(params["id"], 10, 64)

	requestURI := handler.ContentURL + "/api/post/liked/" + strconv.FormatUint(id, 10)

	client := &http.Client{}
	req, _ := http.NewRequest("GET", requestURI, nil)
	req.Header.Set("Authorization", r.Header.Get("Authorization"))
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		WriteResponse(&w, resp)
	}
}

func (handler *PostHandler) GetAllTaggedIn(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r);
	id, _ := strconv.ParseUint(params["id"], 10, 64)

	requestURI := handler.ContentURL + "/api/post/taggedIn/" + strconv.FormatUint(id, 10)

	client := &http.Client{}
	req, _ := http.NewRequest("GET", requestURI, nil)
	req.Header.Set("Authorization", r.Header.Get("Authorization"))
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		WriteResponse(&w, resp)
	}
}

func (handler *PostHandler) GetByUserId(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r);
	user_id, _ := strconv.ParseUint(params["user_id"], 10, 64)

	requestURI := handler.ContentURL + "/api/post/" + strconv.FormatUint(user_id, 10)

	client := &http.Client{}
	req, _ := http.NewRequest("GET", requestURI, nil)
	req.Header.Set("Authorization", r.Header.Get("Authorization"))
	resp, err := client.Do(req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else {
		writeResponse(&w, resp)
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

func (handler *PostHandler) getUserIdFromSession(token string) (uint64, error) {
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
	body, bodyErr := ioutil.ReadAll(r.Body)
	if bodyErr != nil {
		log.Println(bodyErr)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp, err := http.Post(requestURI, "application/json", bytes.NewReader(body))
	log.Println(resp)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	writeResponse(&w, resp)
}

func (handler *PostHandler) CommentPost(w http.ResponseWriter, r *http.Request) {
	bearToken := r.Header.Get("Authorization")
	userId, err := handler.getUserIdFromSession(bearToken)

	params := url.Values{}
	params.Add("userId", strconv.Itoa(int(userId)))

	requestURI := handler.ContentURL + "/api/post/comment?" + params.Encode()
	body, bodyErr := ioutil.ReadAll(r.Body);

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

func (handler *PostHandler) ReportPost(w http.ResponseWriter, r *http.Request) {
	bearToken := r.Header.Get("Authorization")
	userId, err := handler.getUserIdFromSession(bearToken)

	params := url.Values{}
	params.Add("userId", strconv.Itoa(int(userId)))

	requestURI := handler.ContentURL + "/api/post/report?" + params.Encode()
	body, bodyErr := ioutil.ReadAll(r.Body);

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
