package db

import (
	"context"
	"log/slog"
	"github.com/redis/go-redis/v9"
	"github.com/rupak26/Real-time-Leaderboard/config"
)

var Ctx = context.Background()

func InitRedis(cnf config.RedisConfig) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cnf.Addr,                   // Redis default port
		Password: cnf.Password,               // no password set
		DB:       int(cnf.RedisDb),           // use default DB
	})
	slog.Info("Redis Connection implemented")
	return rdb
}