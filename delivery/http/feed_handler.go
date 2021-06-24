package http

import (
	"bytes"
	"io/ioutil"
	"net/http"

	"github.com/APIGateway/globals"
)

type FeedHandler struct {
	ContentURL string
	UsersURL   string
}

func NewFeedHandler() *FeedHandler {
	return &FeedHandler{ContentURL: globals.GetContentMicroserviceUrl(), UsersURL: globals.GetUsersMicroserviceUrl()}
}

func (handler *FeedHandler) GetFeedPosts(w http.ResponseWriter, r *http.Request) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", handler.UsersURL+"/api/follows/followings", nil)
	req.Header.Set("Authorization", r.Header.Get("Authorization"))
	followings, err_followings := client.Do(req)

	if err_followings != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else {

		body, _ := ioutil.ReadAll(followings.Body)
		req, _ := http.NewRequest("GET", handler.ContentURL+"/api/feed?page="+r.URL.Query().Get("page")+"&page_size="+r.URL.Query().Get("page_size"), bytes.NewReader(body))
		resp, err := client.Do(req)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		} else {
			writeResponse(&w, resp)
		}
	}
}
