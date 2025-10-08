package user_handler

import (
	"github.com/rupak26/Real-time-Leaderboard/domain"
)

type Service interface {
	Create(u domain.User) (*domain.User , error)
	Get() ([]domain.User, error)
	Find(email , password string) (*domain.User , error)
}

