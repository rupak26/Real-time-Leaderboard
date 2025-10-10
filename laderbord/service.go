package laderbord

import (
    "github.com/rupak26/Real-time-Leaderboard/domain"
)

type service struct {
	ladrrepo LaderbordRepo
}

func NewService(ladrrepo LaderbordRepo) Service {
    return &service{
		ladrrepo: ladrrepo,
	}
}

func (svc *service) Create(scr domain.SubmitScore) (*domain.SubmitScore, error) {
	 return svc.ladrrepo.Create(scr)
}
func (svc *service) GetIndividulScore(userId int) (*domain.UserRanking, error) {
	 return svc.ladrrepo.GetIndividulScore(userId)
}
func (svc *service) GetScoreList(limit int64) (*[]domain.UserRanking, error) {
	 return svc.ladrrepo.GetScoreList(limit)
}