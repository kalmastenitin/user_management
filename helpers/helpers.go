package helpers

import (
	"encoding/json"
	"log"
	"net/http"
)

type ErrorResponse struct {
	StatusCode   int    `json:"status"`
	ErrorMessage string `json:"message"`
}

func GetError(err error, w http.ResponseWriter, code int) {
	log.Println(err)
	var response = ErrorResponse{
		ErrorMessage: err.Error(),
		StatusCode:   code,
	}

	message, _ := json.Marshal(response)
	w.WriteHeader(response.StatusCode)
	w.Write(message)
}

func GetErrorCustom(err string, w http.ResponseWriter, code int) {
	var response = ErrorResponse{
		ErrorMessage: err,
		StatusCode:   code,
	}

	message, _ := json.Marshal(response)
	w.WriteHeader(response.StatusCode)
	w.Write(message)
}
