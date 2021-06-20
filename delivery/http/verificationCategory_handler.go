package http

import (
	"github.com/APIGateway/globals"
	"log"
	"net/http"
)

type VerificationCategoryHandler struct {
	UsersURL   string
}

func NewVerificationCategoryHandler() *VerificationCategoryHandler {
	return &VerificationCategoryHandler{UsersURL: globals.GetUsersMicroserviceUrl()}
}

func (handler *VerificationCategoryHandler) GetAll(w http.ResponseWriter, r *http.Request){
	client := &http.Client{}
	req, _ := http.NewRequest("GET", handler.UsersURL + "/api/verification/category", nil)
	req.Header.Set("Authorization", r.Header.Get("Authorization"))
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return;
	}

	writeResponse(&w, resp);
}
