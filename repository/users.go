package repository

import (
	"github.com/rupak26/Real-time-Leaderboard/domain"
	"github.com/rupak26/Real-time-Leaderboard/users"	
	"github.com/jmoiron/sqlx"
)

type UserRepo interface {
    users.UserRepo   // embedding 
}

type ReqLogin struct {
	Email string      `json:"email"`
	Password string   `json:"password"`
}

type userRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) UserRepo {
	return &userRepo{
		db: db,
	}
}

func (r *userRepo) Create(usr domain.User) (*domain.User , error) {
    query := `
		INSERT INTO users (
		  username, 
		  email, 
		  password, 
		) VALUES (
		  :username , 
		  :email , 
		  :password ,
		)
		RETURNING id;
	`

	var id int
	rows , err := r.db.NamedQuery(query, usr) 
	if err != nil {
		return nil , err
	}
	defer rows.Close()
	if rows.Next(){
		rows.Scan(&id)
	}
	
	usr.ID = id
	return &usr , nil
}


func (r *userRepo) Find(email, password string) (*domain.User, error) {
	var user domain.User
	query := `SELECT * FROM users WHERE email = $1 AND password = $2`
	err := r.db.Get(&user, query, email, password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepo) Get() ([]domain.User, error) {
	var users []domain.User
	query := `SELECT * FROM users`
	err := r.db.Select(&users, query)
	if err != nil {
		return nil, err
	}
	return users, nil
}
