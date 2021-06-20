package http

import (
	"github.com/APIGateway/globals"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type MediaHandler struct {
	ContentURL string
}

func (handler *MediaHandler) Get(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r);
	id, _ := strconv.ParseUint(params["id"], 10, 64)

	requestURI := handler.ContentURL + "/api/media/" + strconv.FormatUint(id, 10)

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

func NewMediaHandler() *MediaHandler {
	return &MediaHandler{ContentURL: globals.GetContentMicroserviceUrl()}
}
