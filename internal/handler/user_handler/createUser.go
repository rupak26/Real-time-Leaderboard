package user_handler

import (
	"github.com/rupak26/Real-time-Leaderboard/domain"
	"github.com/rupak26/Real-time-Leaderboard/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

type CreateUserReq struct {
	UserName string     `json:"username"`
	Email string        `json:"email"`
	Password string     `json:"password"`
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
		http.Error(w , "Internal Server Error" , http.StatusInternalServerError)
		return
	}
	
	utils.WriteResponse(w , http.StatusCreated , createUser)
}
