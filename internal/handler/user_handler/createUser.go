package user_handler

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/rupak26/Real-time-Leaderboard/domain"
	"github.com/rupak26/Real-time-Leaderboard/utils"
)

type CreateUserReq struct {
	UserName string     `json:"username"`
	Email string        `json:"email"`
	Password string     `json:"password"`
}
type ResponseBody struct {
	Data    interface{} `json:"data,omitempty"`
}


func (h *Handler) CreateUser(w http.ResponseWriter , r *http.Request) {
	
	var newUser CreateUserReq

	decoder := json.NewDecoder(r.Body) 
	err := decoder.Decode(&newUser) 
    
	if err != nil {
		fmt.Fprintln(w,"Give a valid json") 
		return 
	}
	
	createUser , err := h.svc.Create(domain.User{
		UserName: newUser.UserName,
		Email: newUser.Email,
		Password: newUser.Password,
	})
	
	if err != nil {
		slog.Error(err.Error())
		http.Error(w , "Internal Server Error" , http.StatusInternalServerError)
		return
	}
	slog.Info("User Created")
	
	api_response := ApiResponse {
		Status: 201,
		Message: "User Created Successfully",
        Data: createUser,
	}

	utils.WriteResponse(w , http.StatusCreated , api_response)
}
