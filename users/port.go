package users


import (
	"github.com/rupak26/Real-time-Leaderboard/domain"
	"github.com/rupak26/Real-time-Leaderboard/internal/handler/user_handler"

)

type Service interface {
    user_handler.Service //embedding
}


type UserRepo interface {
	Create(u domain.User) (*domain.User , error)
	Get() ([]domain.User, error)
	Find(email , password string) (*domain.User , error)
}
