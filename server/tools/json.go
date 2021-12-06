package tools

import (
	"encoding/json"
	"log"
	"net/http"
)

type errorObject struct {
	Message string `json:"message"`
}

func WriteJsonBadRequest(responseWriter http.ResponseWriter, message string) {
	writeJson(responseWriter, http.StatusBadRequest, errorObject{Message: message})
}

func WriteJsonInternalError(responseWriter http.ResponseWriter, message string) {
	writeJson(responseWriter, http.StatusInternalServerError, errorObject{Message: message})
}

func WriteJsonOk(responseWriter http.ResponseWriter, result interface{}) {
	writeJson(responseWriter, http.StatusOK, result)
}

func WriteJsonOkReplacement(responseWriter http.ResponseWriter) {
	writeJson(responseWriter, http.StatusNoContent, nil)
}

func writeJson(responseWriter http.ResponseWriter, status int, result interface{}) {
	responseWriter.Header().Set("content-type", "application/json")
	responseWriter.WriteHeader(status)
	if result != nil {
		if fault := json.NewEncoder(responseWriter).Encode(result); fault != nil {
			log.Printf("error: something went wrong while encoding json")
		}
	}
}
