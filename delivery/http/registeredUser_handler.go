package http

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/APIGateway/globals"
)

type ResponseError struct {
	Message string `json:"message"`
}

type RegisteredUserHandler struct {
	Url string
}

func NewRegisteredUserHandler() *RegisteredUserHandler {
	return &RegisteredUserHandler{globals.GetUsersMicroserviceUrl()}
}

func (handler *RegisteredUserHandler) Get(w http.ResponseWriter, r *http.Request) {
	req, err := http.NewRequest("GET", handler.Url+"/api/users/user", nil)
	req.Header.Add("Authorization", r.Header.Get("Authorization"))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		writeResponse(&w, resp)
	}
}

func (handler *RegisteredUserHandler) Edit(w http.ResponseWriter, r *http.Request) {
	body, bodyErr := ioutil.ReadAll(r.Body)
	if bodyErr != nil {
		log.Println(bodyErr)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp, err := http.Post(handler.Url+"/api/users/edit", "application/json", bytes.NewReader(body))
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	writeResponse(&w, resp)
}

func (handler *RegisteredUserHandler) EditConfig(w http.ResponseWriter, r *http.Request) {
	body, bodyErr := ioutil.ReadAll(r.Body)
	if bodyErr != nil {
		log.Println(bodyErr)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp, err := http.Post(handler.Url+"/api/users/config/edit", "application/json", bytes.NewReader(body))
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	writeResponse(&w, resp)
}

func (handler *RegisteredUserHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get(handler.Url + "/api/users")
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		writeResponse(&w, resp)
	}
}

func (handler *RegisteredUserHandler) GetUserInfo(w http.ResponseWriter, r *http.Request) {
	bearToken := r.Header.Get("Authorization")
	requestURI := handler.Url + "/api/info"
	client := &http.Client{}
	req, _ := http.NewRequest("GET", requestURI, nil)
	req.Header.Set("Authorization", bearToken)
	resp, err := client.Do(req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else {
		writeResponse(&w, resp)
	}
}

func writeResponse(w *http.ResponseWriter, resp *http.Response) {
	defer resp.Body.Close()
	if responseBody, err := ioutil.ReadAll(resp.Body); err != nil {
		log.Println(err)
		(*w).WriteHeader(http.StatusInternalServerError)
	} else {
		(*w).WriteHeader(resp.StatusCode)
		(*w).Write(responseBody)
	}
}
