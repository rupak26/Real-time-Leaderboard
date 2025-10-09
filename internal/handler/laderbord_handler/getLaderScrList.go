package laderbord_handler

import (

	"net/http"
	"strconv"

	"github.com/rupak26/Real-time-Leaderboard/utils"
)


func (h *Handler) GetLaderScrList (w http.ResponseWriter , r *http.Request) {	
	limit := r.URL.Query().Get("limit") 

	lmt, _ := strconv.ParseInt(limit, 10, 64)
	if lmt == 0 {
		lmt = 10
	}
    laderList , err := h.svc.GetScoreList(lmt) 
	if err != nil {
       utils.WriteResponse(w ,http.StatusInternalServerError , "Internal Server Error" )
	   return
	}
     
    utils.SendPage(w , laderList , lmt)
}