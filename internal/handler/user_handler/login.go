package user_handler

import (
	"github.com/rupak26/Real-time-Leaderboard/config"
	//	"github.com/rupak26/Real-time-Leaderboard/domain"
	"encoding/json"
	"net/http"

	"github.com/rupak26/Real-time-Leaderboard/utils"
)

type ReqLogin struct {
	Email    string   `json:"email"`
	Password string   `json:"password"`
}

type ApiResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}




func (h *Handler) Login(w http.ResponseWriter , r *http.Request) {
	
	var reqLogin ReqLogin
    
	decoder := json.NewDecoder(r.Body) 
	err := decoder.Decode(&reqLogin) 
   
	key := config.GetConfig().SecretKey

	if err != nil {
		http.Error(w,"Invalid Request Data" , http.StatusBadRequest)
		return 
	}

	user , err := h.svc.Find(reqLogin.Email , reqLogin.Password) 
	if err != nil {
		http.Error(w , "Invalid Credentials" , http.StatusBadRequest)
		return 
	}

	jwt , err := utils.CreateJwt(key , utils.Payload{
		Sub: 1,
		UserId: user.ID,
		UserName: user.UserName,
		Email: user.Email,
		Password: reqLogin.Password,
	})
	
		
	if err != nil {
		utils.WriteResponse(w , http.StatusInternalServerError , "Internal Server Error")
		return 
	}
	
	response := ApiResponse {
		Status: 201,
		Message: "Access Token",
		Data: jwt,
	}

	utils.WriteResponse(w , http.StatusCreated , response)
}




