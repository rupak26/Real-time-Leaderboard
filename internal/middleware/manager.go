package middleware

import (
	// "ecommerce/global_router"
	// "ecommerce/middleware"
	"net/http"
)

// type rupak int
// var number rupak

type Middleware func(http.Handler) http.Handler 

type Manager struct {
	globalmiddleware []Middleware 
}

func NewManager() *Manager {
    return &Manager{
		globalmiddleware: make([]Middleware, 0),
	}
}

func (mngr *Manager) Use(middlewares ...Middleware) {
	mngr.globalmiddleware = append(mngr.globalmiddleware, middlewares...)
}

func (mngr *Manager) With(next http.Handler , middlewares ...Middleware) http.Handler {
	n := next 
	for _,middleware := range middlewares {
		n = middleware(n) 
	}
	return  n 
}