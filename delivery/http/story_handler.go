package http

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/APIGateway/globals"
)

type StoryHandler struct {
	ContentURL string
}

func NewStoryHandler() *StoryHandler {
	return &StoryHandler{ContentURL: globals.GetContentMicroserviceUrl()}
}

func (handler *StoryHandler) Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	client := &http.Client{}
	req, _ := http.NewRequest("DELETE", handler.ContentURL+"/api/story/"+params["id"], nil)
	req.Header.Set("Authorization", r.Header.Get("Authorization"))
	resp, err := client.Do(req)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		WriteResponse(&w, resp)
	}
}
