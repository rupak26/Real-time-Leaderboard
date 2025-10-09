package repository

import (
	"fmt"
	"context"
	"github.com/rupak26/Real-time-Leaderboard/domain"
	"github.com/rupak26/Real-time-Leaderboard/laderbord"
	"github.com/redis/go-redis/v9"
	"strconv"
)

type LaderRepo interface {
    laderbord.LaderbordRepo // Embeeding
}

type ladrRepo struct {
	rdb *redis.Client
	ctx context.Context
	key string
}


func NewLaderbordRepo(rdb *redis.Client) LaderRepo {
     repo := &ladrRepo{
		rdb: rdb,
		ctx: context.Background(),
		key: "leaderboard",
	 }
	 return repo
}


func (r *ladrRepo) Create(scr domain.SubmitScore) (*domain.SubmitScore, error) {

	err := r.rdb.ZAdd(r.ctx, r.key, redis.Z{
		Score:  float64(scr.Score),
		Member: fmt.Sprintf("%s:%s", scr.UserName, scr.GameId), 
	}).Err()
	if err != nil {
		return nil, err
	}

	return &scr, nil
}


func (r *ladrRepo) GetIndividulScore(userId int) (int64 , error) {
    var username string
	err := r.db.Get(&username, "SELECT username FROM users WHERE id = $1", userId)
	if err != nil {
		return 0, fmt.Errorf("failed to fetch username: %w", err)
	}

	// 2️⃣ Fetch game_id from leaderboard table
	var gameId int
	err = r.db.Get(&gameId, "SELECT game_id FROM leaderboard WHERE user_id = $1", userId)
	if err != nil {
		return 0, fmt.Errorf("failed to fetch game_id: %w", err)
	}
	 

	 rank , err := r.rdb.ZRank(r.ctx , r.key , "jhon:cricket__1").Result()
	 if err != nil {
		return 0, err
	 }
	 fmt.Println(userid)
     return rank, nil
}

func (r *ladrRepo) GetScoreList(limit int64) (*[]domain.UserRanking, error) {
	var userRank []domain.UserRanking

	allrec, err := r.rdb.ZRevRangeWithScores(r.ctx, r.key, 0, limit-1).Result()
	if err != nil {
		return nil, err
	}
	for i, rec := range allrec {
		userRank = append(userRank, domain.UserRanking{
			UserName: rec.Member.(string), 
			Score:    rec.Score,
			Rank:     int64(i + 1), // rank starts at 1
		})
	}

	return &userRank, nil
}
