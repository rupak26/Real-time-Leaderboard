package user_handler

import (
	"net/http"
	"github.com/rupak26/Real-time-Leaderboard/internal/middleware"
)

func (h *Handler) RegisterRouters(mux *http.ServeMux , manager *middleware.Manager) {
	mux.Handle(
		"GET /users" ,
		manager.With(
			http.HandlerFunc(h.GetUser),
			h.middleware.EnableCORS,
	    ),
    )
	mux.Handle(
		"POST /users/register" ,
		manager.With(
			http.HandlerFunc(h.CreateUser),
			h.middleware.EnableCORS,
	    ),
	)
	mux.Handle(
		"POST /users/login" ,
		manager.With(
			http.HandlerFunc(h.Login),
			h.middleware.EnableCORS,
	    ),
	)
}