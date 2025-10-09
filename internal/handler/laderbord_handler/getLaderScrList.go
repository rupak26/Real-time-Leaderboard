package laderbord_handler

import (

	"net/http"
	"strconv"

	"github.com/rupak26/Real-time-Leaderboard/utils"
)


func (h *Handler) GetLaderScrList (w http.ResponseWriter , r *http.Request) {
	page := r.URL.Query().Get("page")
	limit := r.URL.Query().Get("limit") 

	pg, _ := strconv.ParseInt(page, 10, 64)
	lmt, _ := strconv.ParseInt(limit, 10, 64)
	if pg == 0 {
		pg = 1 
	}
	if lmt == 0 {
		lmt = 10
	}
    laderList , err := h.svc.GetScoreList(pg , lmt) 
	if err != nil {
       utils.WriteResponse(w ,http.StatusInternalServerError , "Internal Server Error" )
	   return
	}
     
    utils.SendPage(w , laderList , pg , lmt)
}