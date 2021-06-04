package http

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"github.com/APIGateway/globals"
)

type SearchHandler struct {
	UsersURL   string
	ContentURL string
}

func NewSearchHandler() *SearchHandler {
	return &SearchHandler{UsersURL: globals.GetUsersMicroserviceUrl(), ContentURL: globals.GetContentMicroserviceUrl()}
}

func WriteResponse(w *http.ResponseWriter, resp *http.Response) {
	defer resp.Body.Close()
	if responseBody, err := ioutil.ReadAll(resp.Body); err != nil {
		log.Println(err)
		(*w).WriteHeader(http.StatusInternalServerError)
	} else {
		(*w).WriteHeader(resp.StatusCode)
		(*w).Write(responseBody)
	}
}

func (handler *SearchHandler) Search(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	hashtag := r.URL.Query().Get("hashtag")
	location := r.URL.Query().Get("location")

	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	pageSize, err := strconv.Atoi(r.URL.Query().Get("page_size"))
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	params := url.Values{}
	params.Add("page", strconv.Itoa(page))
	params.Add("page_size", strconv.Itoa(pageSize))

	path := "/api/search?"
	var requestURI string

	if username != "" {
		params.Add("username", username)
		requestURI = globals.GetUsersMicroserviceUrl() + path + params.Encode()
	} else if hashtag != "" {
		params.Add("hashtag", hashtag)
		requestURI = globals.GetContentMicroserviceUrl() + path + params.Encode()
	} else if location != "" {
		params.Add("location", location)
		requestURI = globals.GetContentMicroserviceUrl() + path + params.Encode()
	} else {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	resp, err := http.Get(requestURI)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		WriteResponse(&w, resp)
	}
}
