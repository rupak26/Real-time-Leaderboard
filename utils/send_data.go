package utils

import (
	"encoding/json"
	"net/http"
)


func Send_erros(w http.ResponseWriter , message string , status int) error {
   w.WriteHeader(status)
   return json.NewEncoder(w).Encode(message)
}