package http

import (
	"github.com/APIGateway/globals"
	"io/ioutil"
	"log"
	"net/http"
)

type ResponseError struct {
	Message string `json:"message"`
}

type RegisteredUserHandler struct {
	Url string
}

func NewRegisteredUserHandler() *RegisteredUserHandler{
	return &RegisteredUserHandler{globals.GetUsersMicroserviceUrl()}
}


func (handler *RegisteredUserHandler) GetAll(w http.ResponseWriter, r *http.Request){
	resp, err := http.Get(handler.Url + "/api/users")
	if err != nil{
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		writeResponse(&w, resp);
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

