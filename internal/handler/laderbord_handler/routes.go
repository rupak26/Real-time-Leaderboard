package laderbord_handler

import (
	"net/http"
	"github.com/rupak26/Real-time-Leaderboard/internal/middleware"
)

func (h *Handler) RegisterRouters(mux *http.ServeMux , manager *middleware.Manager) {
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