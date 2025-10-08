package repository

import (
	"github.com/rupak26/Real-time-Leaderboard/laderbord"
	"github.com/jmoiron/sqlx"
	"github.com/rupak26/Real-time-Leaderboard/domain"
)

type LaderRepo interface {
    laderbord.LaderbordRepo // Embeeding
}

type ladrRepo struct {
    db *sqlx.DB
}


func NewLaderbordRepo(db *sqlx.DB) LaderRepo {
     repo := &ladrRepo{
		db : db ,
	 }
	 return repo
}


func (r *ladrRepo) Create(scr domain.SubmitScore) (*domain.SubmitScore, error) {
     query := `
		INSERT INTO scores (
			game_id,
			score
		) VALUES (
			:game_id,
			:score
		)
		RETURNING id;
    `
}
func (r *ladrRepo) GetIndividulScore(userId int) (*domain.UserRanking, error) {

}
func (r *ladrRepo) GetScoreList(page int64, limit int64) (*[]domain.UserRanking, error) {

}