package response

import (
	"encoding/json"
	"log"
	"net/http"
)

type JsonResponse struct {
}

func NewJsonResponse() *JsonResponse {
	return &JsonResponse{}
}
func (j JsonResponse) Send(w http.ResponseWriter, code int, message string, payload interface{}) {
	res := Pattern{
		Data:    payload,
		Message: message,
		Status:  code,
	}
	response, err := json.Marshal(res)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		log.Printf("Failed to encode response: %v", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, err = w.Write(response)
	if err != nil {
		log.Printf("Failed to write response: %v", err)
	}
}
func (j JsonResponse) Created(w http.ResponseWriter, payload interface{}) {
	j.Send(w, http.StatusCreated, "CREATED WITH SUCESS!", payload)
}
func (j JsonResponse) Ok(w http.ResponseWriter, payload interface{}) {
	j.Send(w, http.StatusOK, "OK!", payload)
}
func (j JsonResponse) BadRequest(w http.ResponseWriter, payload interface{}) {
	j.Send(w, http.StatusBadRequest, "BAD REQUEST", payload)
}
func (j JsonResponse) NotFound(w http.ResponseWriter, payload interface{}) {
	j.Send(w, http.StatusNotFound, "NOT FOUND", payload)
}
func (j JsonResponse) InternalServerError(w http.ResponseWriter, payload interface{}) {
	j.Send(w, http.StatusInternalServerError, "INTERNAL SERVER ERROR", payload)
}
