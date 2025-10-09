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
    var userRank domain.UserRanking
	
	query := `
		SELECT username, score, rank
		FROM (
			SELECT 
				u.username,
				l.user_id,
				l.score,
				RANK() OVER (ORDER BY l.score DESC) AS rank
			FROM leaderboard l
			JOIN users u ON u.id = l.user_id
		) ranked
		WHERE ranked.user_id = $1;
	`
	err := r.db.Get(&userRank, query , userId) 
	if err != nil {
		return nil, err
	}
	return &userRank, nil
}

func (r *ladrRepo) GetScoreList(page int64, limit int64) (*[]domain.UserRanking, error) {
    var scores []domain.UserRanking
	offset := (page - 1) * limit
	query := `
		SELECT 
			u.username,
			l.score,
			RANK() OVER (ORDER BY l.score DESC) AS rank
		FROM leaderboard l
		JOIN users u ON u.id = l.user_id
		ORDER BY rank
		LIMIT $1 OFFSET $2;
	`
	err := r.db.Select(&scores, query , limit , offset)
	if err != nil {
		return nil, err
	}
	return &scores, nil
}