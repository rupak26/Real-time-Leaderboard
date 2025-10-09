package users

import (
	"github.com/rupak26/Real-time-Leaderboard/domain"
)

type service struct {
	usrRepo UserRepo 
}

func NewService(usrRepo UserRepo) Service {
    return &service{
		usrRepo: usrRepo,
	}
}

func (svc *service) Create(user domain.User) (*domain.User, error) {
	return  svc.usrRepo.Create(user)
}

func (svc *service) Find(email string , password string) (*domain.User , error) {
	usr , err := svc.usrRepo.Find(email , password) 
	if err != nil {
		return nil , err
	}
	return usr , nil
}
func (svc *service) Get() ([]domain.User , error) {
	usr , err := svc.usrRepo.Get() 
	if err != nil {
		return nil , err
	}
	return usr , nil
}