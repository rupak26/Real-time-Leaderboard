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
		Member: fmt.Sprintf("%d",scr.UserId), 
	}).Err()
	if err != nil {
		return nil, err
	}
    
	userKey := fmt.Sprintf("user:%d", scr.UserId)
    _, err = r.rdb.HSet(r.ctx, userKey, map[string]interface{}{
        "username": scr.UserName,
        "game_id":  scr.GameId,
    }).Result()
    if err != nil {
        return nil , err
    }

	return &scr, nil
}


func (r *ladrRepo) GetIndividulScore(userId int) (*domain.UserRanking , error) {
    userid := strconv.Itoa(userId)
	rank , err := r.rdb.ZRank(r.ctx , r.key , userid).Result()
	if err == redis.Nil {
        return nil, fmt.Errorf("user not found in leaderboard")
    } else if err != nil {
        return nil, err
    }
	score, err := r.rdb.ZScore(r.ctx, r.key, strconv.Itoa(userId)).Result()
    if err != nil {
        return nil, err
    }

    // 3️⃣ Get user info
    userKey := fmt.Sprintf("user:%d", userId)
    userData, err := r.rdb.HGetAll(r.ctx, userKey).Result()
    if err != nil {
        return nil, err
    }

    userRank := &domain.UserRanking{
        UserName: userData["username"],
        Score:    float64(score),
        Rank:     rank + 1,
    }

    return userRank, nil
}



func (r *ladrRepo) GetScoreList(limit int64) (*[]domain.UserRanking, error) {
	var rankings []domain.UserRanking

	players, err := r.rdb.ZRevRangeWithScores(r.ctx, r.key, 0, limit-1).Result()
	if err != nil {
		return nil, err
	}
	
    for i, player := range players {
        userId, _ := strconv.Atoi(player.Member.(string))
        userKey := fmt.Sprintf("user:%d", userId)
        userData, _ := r.rdb.HGetAll(r.ctx, userKey).Result()

        rankings = append(rankings, domain.UserRanking{
            UserName: userData["username"],
            Score:    float64(player.Score),
            Rank:     int64(i + 1),
        })
    }
	return &rankings, nil
}
