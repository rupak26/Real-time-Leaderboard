package user_handler

import (
	"log/slog"
	"net/http"

	"github.com/rupak26/Real-time-Leaderboard/utils"
)

func (h *Handler) GetUser(w http.ResponseWriter , r *http.Request) {
	userList , err := h.svc.Get()
	
	if err != nil {
	   slog.Error(err.Error())	
       http.Error(w , "Internal Server Error" , http.StatusInternalServerError)
	   return
	}
	slog.Info("User info collection")
    utils.WriteResponse(w , http.StatusOK , userList)
}
