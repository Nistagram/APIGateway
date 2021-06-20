package http

import (
	"bytes"
	"github.com/APIGateway/globals"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type VerificationHandler struct {
	UsersURL   string
}

func NewVerificationHandler() *VerificationHandler {
	return &VerificationHandler{UsersURL: globals.GetUsersMicroserviceUrl()}
}

func (handler *VerificationHandler) CreateVerificationRequest(w http.ResponseWriter, r *http.Request){
	body, bodyErr := ioutil.ReadAll(r.Body);
	if bodyErr != nil {
		log.Println(bodyErr)
		w.WriteHeader(http.StatusInternalServerError)
		return;
	}

	req, err := http.NewRequest("POST", handler.UsersURL + "/api/verification/verify", bytes.NewReader(body))
	req.Header.Add("Authorization", r.Header.Get("Authorization"))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		writeResponse(&w, resp)
	}
}

func (handler *VerificationHandler) Accept(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r);
	id, _ := strconv.ParseUint(params["id"], 10, 64)

	client := &http.Client{}
	req, _ := http.NewRequest("POST", handler.UsersURL + "/api/verification/accept/" + strconv.FormatUint(id, 10), nil)
	req.Header.Set("Authorization", r.Header.Get("Authorization"))
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return;
	}
	writeResponse(&w, resp);
}

func (handler *VerificationHandler) Reject(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r);
	id, _ := strconv.ParseUint(params["id"], 10, 64)

	client := &http.Client{}
	req, _ := http.NewRequest("POST", handler.UsersURL + "/api/verification/reject/" + strconv.FormatUint(id, 10), nil)
	req.Header.Set("Authorization", r.Header.Get("Authorization"))
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return;
	}
	writeResponse(&w, resp);
}

func (handler *VerificationHandler) GetVerificationStatus(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r);
	id, _ := strconv.ParseUint(params["id"], 10, 64)

	client := &http.Client{}
	req, _ := http.NewRequest("GET", handler.UsersURL + "/api/verification/status/" + strconv.FormatUint(id, 10), nil)
	req.Header.Set("Authorization", r.Header.Get("Authorization"))
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return;
	}
	writeResponse(&w, resp);
}

func (handler *VerificationHandler) GetUnresolved(w http.ResponseWriter, r *http.Request){
	client := &http.Client{}
	req, _ := http.NewRequest("GET", handler.UsersURL + "/api/verification/unresolved", nil)
	req.Header.Set("Authorization", r.Header.Get("Authorization"))
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return;
	}
	writeResponse(&w, resp);
}
