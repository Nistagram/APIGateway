package http

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/APIGateway/globals"
)

type AdvertisementHandler struct {
	AdvertisementURL string
}

func NewAdvertisementHandler() *AdvertisementHandler {
	return &AdvertisementHandler{AdvertisementURL: globals.GetAdvertisementMicroserviceUrl()}
}

func (handler *AdvertisementHandler) Query(w http.ResponseWriter, r *http.Request) {
	requestURI := handler.AdvertisementURL + "/query"
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
