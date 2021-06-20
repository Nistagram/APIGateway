package http

import (
	"github.com/APIGateway/globals"
	"log"
	"net/http"
)

type BusinessCategoryHandler struct {
	UsersURL   string
}

func NewBusinessCategoryHandlerHandler() *BusinessCategoryHandler {
	return &BusinessCategoryHandler{UsersURL: globals.GetUsersMicroserviceUrl()}
}

func (handler *BusinessCategoryHandler) GetAll(w http.ResponseWriter, r *http.Request){
	client := &http.Client{}
	req, _ := http.NewRequest("GET", handler.UsersURL + "/api/business-category", nil)
	req.Header.Set("Authorization", r.Header.Get("Authorization"))
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return;
	}

	writeResponse(&w, resp);
}