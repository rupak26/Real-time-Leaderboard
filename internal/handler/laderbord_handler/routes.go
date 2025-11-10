package laderbord_handler

import (
	"net/http"
	"github.com/rupak26/Real-time-Leaderboard/internal/middleware"
)

func (h *Handler) RegisterRouters(mux *http.ServeMux , manager *middleware.Manager) {
	mux.Handle("OPTIONS /submit-score", http.HandlerFunc(h.handleOptions))
    mux.Handle("OPTIONS /leaderboard", http.HandlerFunc(h.handleOptions))
    mux.Handle("OPTIONS /user-ranking/{id}", http.HandlerFunc(h.handleOptions))

	mux.Handle(
		"POST /submit-score" ,
		manager.With(
			http.HandlerFunc(h.CreateLaderScr),
			h.middleware.Authorization,
			h.middleware.EnableCORS,
	    ),
	)
	mux.Handle(
		"GET /leaderboard" ,
		manager.With(
			http.HandlerFunc(h.GetLaderScrList),
			h.middleware.Authorization,
			h.middleware.EnableCORS,
	    ),
	)
	mux.Handle(
		"GET /user-ranking/{id}",
		manager.With(
			http.HandlerFunc(h.GetIndividualLaderScr),
			h.middleware.Authorization,
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