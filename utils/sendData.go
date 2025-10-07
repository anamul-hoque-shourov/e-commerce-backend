package utils

import (
	"encoding/json"
	"net/http"
)

func SendData(res http.ResponseWriter, data any, statusCode int) {
	res.WriteHeader(statusCode)
	json.NewEncoder(res).Encode(data)
}

func SendError(res http.ResponseWriter, msg string, statusCode int) {
	res.WriteHeader(statusCode)
	json.NewEncoder(res).Encode(msg)
}
