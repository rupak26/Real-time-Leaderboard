package utils

import (
	"encoding/json"
	"net/http"
	"github.com/rupak26/Real-time-Leaderboard/domain"
)

func WriteResponse(w http.ResponseWriter , status int , data interface{}) error {
	w.WriteHeader(status)
//	return json.NewEncoder(w).Encode(data)
    response := domain.ApiResponse {
		Success: true,
		Message: "Successfully Done",
		Data: data,
		StatusCode: status,
	}
	return json.NewEncoder(w).Encode(response)
}