package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/rupak26/Real-time-Leaderboard/domain"
	"github.com/rupak26/Real-time-Leaderboard/laderbord"
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
		INSERT INTO leaderboard (user_id, game_id, score)
		VALUES ($1, $2, $3)
		RETURNING id, user_id, game_id, score
	`

	var result domain.SubmitScore

	if err := r.db.QueryRowx(query, scr.UserId, scr.GameId, scr.Score).
		StructScan(&result); err != nil {
		return nil, err
	}

	return &result, nil
}


func (r *ladrRepo) GetIndividulScore(userId int) (*domain.UserRanking, error) {
    var score domain.UserRanking
	query := `SELECT * FROM leaderboard WHERE user_id = $1`
	err := r.db.Get(&score, query, userId) 
	if err != nil {
		return nil, err
	}
	return &score, nil
}

func (r *ladrRepo) GetScoreList(page int64, limit int64) (*[]domain.UserRanking, error) {
    var scores []domain.UserRanking
	query := `SELECT * FROM leaderboard`
	err := r.db.Select(&scores, query)
	if err != nil {
		return nil, err
	}
	return &scores, nil
}