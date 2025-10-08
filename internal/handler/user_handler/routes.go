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
	    ),
    )
	mux.Handle(
		"POST /users" ,
		manager.With(
			http.HandlerFunc(h.CreateUser),
	    ),
	)
	mux.Handle(
		"POST /users/login" ,
		manager.With(
			http.HandlerFunc(h.Login),
	    ),
	)
}