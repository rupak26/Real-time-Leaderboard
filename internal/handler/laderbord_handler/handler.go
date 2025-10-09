package laderbord_handler

import "github.com/rupak26/Real-time-Leaderboard/internal/middleware"


type Handler struct {
	middleware *middleware.Middlewares
	svc Service
}

func NewHandler( middlewares *middleware.Middlewares, svc Service , ) *Handler {
	return &Handler{
		middleware: middlewares,
		svc: svc,
	}
}