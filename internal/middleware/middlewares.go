package middleware

import (
	"github.com/rupak26/Real-time-Leaderboard/config"
)

type Middlewares struct {
	cnf *config.Config
}

func NewMiddleware(cnf *config.Config) *Middlewares {
	return &Middlewares{
		cnf: cnf,
	}
}