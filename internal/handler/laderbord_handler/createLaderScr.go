package laderbord_handler

import (
	"fmt"
	"net/http"

	"encoding/json"

	"github.com/rupak26/Real-time-Leaderboard/domain"
	"github.com/rupak26/Real-time-Leaderboard/utils"
	
)			

type ReqLaderCreate struct {
    GameId   string `json:"game_id"`
	Score    int64  `json:"score"`
}

func (h *Handler) CreateLaderScr( w http.ResponseWriter , r *http.Request ){
	var req ReqLaderCreate 

	decoder := json.NewDecoder(r.Body) 
	err := decoder.Decode(&req) 

	if err != nil {
		fmt.Fprintln(w,"Give a valid json") 	
		return 
	}

	createLaderScr , err := h.svc.Create(domain.SubmitScore{
		GameId: req.GameId,
		Score : req.Score,
	})
	utils.WriteResponse(w , http.StatusCreated , createLaderScr)
}