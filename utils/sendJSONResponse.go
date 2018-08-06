package utils

import (
	"encoding/json"
	"net/http"
)

// SendJSONResponse this function will give Content-Type JSON on given response
func SendJSONResponse(res http.ResponseWriter, error interface{}, message interface{}, data interface{}, httpCode int) {
	res.Header().Set("Content-Type", "application/json;charset=UTF-8")
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.WriteHeader(httpCode)

	response := &ResponseJSON{
		Error:   error,
		Data:    data,
		Message: message,
	}

	if err := json.NewEncoder(res).Encode(response); err != nil {
		panic(err)
	}
}
