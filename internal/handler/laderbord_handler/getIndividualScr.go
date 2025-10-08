package laderbord_handler

import (
	"net/http"
	"strconv"
	"github.com/rupak26/Real-time-Leaderboard/utils"
)

func (h *Handler) GetIndividualLaderScr(w http.ResponseWriter , r *http.Request) {
     id := r.PathValue("id") 

	 scrId , err := strconv.Atoi(id)

	if err != nil {
		utils.WriteResponse(w , http.StatusInternalServerError , "Enter a valid scr id")
		return
	}

	scr , err := h.svc.GetIndividulScore(scrId)
	if err != nil {
		utils.Send_erros(w , "Product not found" , http.StatusNotFound)
		return
	}
	utils.WriteResponse(w , http.StatusOK , scr)
}