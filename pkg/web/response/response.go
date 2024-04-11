package response

import "net/http"

type Response interface {
	Created(w http.ResponseWriter, payload interface{})
	Ok(w http.ResponseWriter, payload interface{})
	BadRequest(w http.ResponseWriter, payload interface{})
	NotFound(w http.ResponseWriter, payload interface{})
	InternalServerError(w http.ResponseWriter, payload interface{})
	InvalidParameters(w http.ResponseWriter, payload interface{})
}

type Pattern struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
