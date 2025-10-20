package laderbord_handler

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/rupak26/Real-time-Leaderboard/utils"
)

func (h *Handler) GetIndividualLaderScr(w http.ResponseWriter , r *http.Request) {
     id := r.PathValue("id") 

	 scrId , err := strconv.Atoi(id)

	if err != nil {
		slog.Error(err.Error())
		utils.WriteResponse(w , http.StatusInternalServerError , "Enter a valid scr id")
		return
	}

	scr , err := h.svc.GetIndividulScore(scrId)
	if err != nil {
		slog.Error(err.Error())
		utils.Send_erros(w , "Scores not found" , http.StatusNotFound)
		return
	}
	slog.Info("User Info collection")
	utils.WriteResponse(w , http.StatusOK , scr)
}