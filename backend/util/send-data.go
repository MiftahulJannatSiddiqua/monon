package util

import (
	"encoding/json"
	"net/http"
)

func SendData(w http.ResponseWriter, statuscode int, data interface{}) {
	w.WriteHeader(statuscode)
	encoder := json.NewEncoder(w)
	encoder.Encode(data)
}

func SendError(w http.ResponseWriter, startusCode int, msg string) {
	w.WriteHeader(startusCode)
	encoder := json.NewEncoder(w)
	encoder.Encode(msg)
}