package laderbord_handler

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/rupak26/Real-time-Leaderboard/domain"
	"github.com/rupak26/Real-time-Leaderboard/utils"
)			

type ReqLaderCreate struct {
    GameId   string `json:"game_id"`
	Score    int64  `json:"score"`
}

// @Summary Submit a new leaderboard score
// @Description Create a new leaderboard record for a user
// @Tags leaderboard
// @Accept  json
// @Produce  json
// @Param   score  body  domain.SubmitScore  true  "Score info"
// @Success 200 {object} domain.SubmitScore
// @Failure 400 {object} map[string]string
// @Router /submit-score [post]
func (h *Handler) CreateLaderScr( w http.ResponseWriter , r *http.Request ){
	var req ReqLaderCreate 

	decoder := json.NewDecoder(r.Body) 
	err := decoder.Decode(&req) 

    usr_name := r.Context().Value("user_name").(string)
    user_id  := r.Context().Value("user_id").(int)
	
	
	if err != nil {
		slog.Error("Give a valid json")
		fmt.Fprintln(w,"Give a valid json") 	
		return 
	}
	createLaderScr , err := h.svc.Create(domain.SubmitScore{
		UserId: user_id,
		UserName: usr_name,
		GameId: req.GameId,
		Score : req.Score,
	})
    
	if err != nil {
		slog.Error(err.Error())
		utils.WriteResponse(w , http.StatusInternalServerError , err)	
		return 
	}
    slog.Info("laderboard item created")
	utils.WriteResponse(w , http.StatusCreated , createLaderScr)
}