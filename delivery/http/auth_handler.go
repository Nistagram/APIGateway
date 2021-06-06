package http

import (
	"bytes"
	"github.com/APIGateway/globals"
	"io/ioutil"
	"log"
	"net/http"
)

type AuthHandler struct {
	UsersURL   string
}

func (handler *AuthHandler) Login(w http.ResponseWriter, r *http.Request){
	body, bodyErr := ioutil.ReadAll(r.Body);
	if bodyErr != nil {
		log.Println(bodyErr)
		w.WriteHeader(http.StatusInternalServerError)
		return;
	}

	resp, err := http.Post(handler.UsersURL + "/api/auth/login", "application/json", bytes.NewReader(body));
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return;
	}
	writeResponse(&w, resp);
}

func (handler *AuthHandler) Register(w http.ResponseWriter, r *http.Request){
	body, bodyErr := ioutil.ReadAll(r.Body);
	if bodyErr != nil {
		log.Println(bodyErr)
		w.WriteHeader(http.StatusInternalServerError)
		return;
	}

	resp, err := http.Post(handler.UsersURL + "/api/auth/register", "application/json", bytes.NewReader(body));
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return;
	}
	writeResponse(&w, resp);
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{UsersURL: globals.GetUsersMicroserviceUrl()}
}
