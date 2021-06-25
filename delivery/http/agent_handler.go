package http

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/APIGateway/globals"
)

type AgentHandler struct {
	UsersURL string
}

func NewAgentHandler() *AgentHandler {
	return &AgentHandler{UsersURL: globals.GetUsersMicroserviceUrl()}
}

func (handler *AgentHandler) GetAllRegistrationRequests(w http.ResponseWriter, r *http.Request) {
	requestURI := handler.UsersURL + "/api/agent/registration-request"

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

func (handler *AgentHandler) CreateRegistrationRequest(w http.ResponseWriter, r *http.Request) {
	requestURI := handler.UsersURL + "/api/agent/registration-request"

	body, bodyErr := ioutil.ReadAll(r.Body)
	if bodyErr != nil {
		log.Println(bodyErr)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	client := &http.Client{}
	req, _ := http.NewRequest("POST", requestURI, bytes.NewReader(body))
	//req.Header.Set("Authorization", r.Header.Get("Authorization"))
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		WriteResponse(&w, resp)
	}
}
