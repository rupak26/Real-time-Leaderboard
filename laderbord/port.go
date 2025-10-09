package laderbord

import (
	"github.com/rupak26/Real-time-Leaderboard/internal/handler/laderbord_handler"
    "github.com/rupak26/Real-time-Leaderboard/domain"
)

type Service interface{
    laderbord_handler.Service // Embedding
}

type LaderbordRepo interface {
    Create(scr domain.SubmitScore)   (*domain.SubmitScore , error) 
	GetScoreList(limit int64) (*[]domain.UserRanking , error)
	GetIndividulScore(userId int)    (int64 , error)
}


