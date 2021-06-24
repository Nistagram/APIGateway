package http

import (
	"log"
	"net/http"

	"github.com/APIGateway/globals"
	"github.com/gorilla/mux"
)

type ReportHandler struct {
	ContentUrl string
}

func NewReportHandler() *ReportHandler {
	return &ReportHandler{ContentUrl: globals.GetContentMicroserviceUrl()}
}

func (handler *ReportHandler) GetAllTypes(w http.ResponseWriter, r *http.Request) {
	requestURI := handler.ContentUrl + "/api/report-types"

	resp, err := http.Get(requestURI)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		WriteResponse(&w, resp)
	}
}

func (handler *ReportHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", handler.ContentUrl+"/api/report?page="+r.URL.Query().Get("page")+"&page_size="+r.URL.Query().Get("page_size"), nil)
	req.Header.Set("Authorization", r.Header.Get("Authorization"))
	resp, err := client.Do(req)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		WriteResponse(&w, resp)
	}
}

func (handler *ReportHandler) Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	client := &http.Client{}
	req, _ := http.NewRequest("DELETE", handler.ContentUrl+"/api/report/"+params["id"], nil)
	req.Header.Set("Authorization", r.Header.Get("Authorization"))
	resp, err := client.Do(req)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		WriteResponse(&w, resp)
	}
}
