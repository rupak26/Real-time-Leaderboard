package user_handler

import (
	"net/http"
	"github.com/rupak26/Real-time-Leaderboard/internal/middleware"
)

func (h *Handler) RegisterRouters(mux *http.ServeMux , manager *middleware.Manager) {
	mux.Handle("OPTIONS /users", http.HandlerFunc(h.handleOptions))
    mux.Handle("OPTIONS /users/register", http.HandlerFunc(h.handleOptions))
    mux.Handle("OPTIONS /users/login", http.HandlerFunc(h.handleOptions))

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

func (h *Handler) handleOptions(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
    w.WriteHeader(http.StatusOK)
}