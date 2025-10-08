package utils

import (
	"encoding/json"
	"net/http"
)


func WriteResponse(w http.ResponseWriter , status int , data interface{}) error {
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}