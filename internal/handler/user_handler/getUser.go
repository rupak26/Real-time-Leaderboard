package user_handler

import (
	"net/http"
	"github.com/rupak26/Real-time-Leaderboard/utils"
)

func (h *Handler) GetUser(w http.ResponseWriter , r *http.Request) {
	userList , err := h.svc.Get()
	
	if err != nil {
       http.Error(w , "Internal Server Error" , http.StatusInternalServerError)
	   return
	}
    utils.WriteResponse(w , http.StatusOK , userList)
}
