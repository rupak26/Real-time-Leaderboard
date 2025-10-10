package laderbord_handler

import (
   "github.com/rupak26/Real-time-Leaderboard/domain"
)

type Service interface{
     Create(scr domain.SubmitScore)   (*domain.SubmitScore , error) 
	 GetScoreList(limit int64) (*[]domain.UserRanking , error)
	 GetIndividulScore(userId int) (*domain.UserRanking, error)
}