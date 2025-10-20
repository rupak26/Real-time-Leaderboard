package handler

import (
	"github.com/rupak26/Real-time-Leaderboard/config"

	"github.com/rupak26/Real-time-Leaderboard/internal/handler/laderbord_handler"
	"github.com/rupak26/Real-time-Leaderboard/internal/handler/user_handler"
    httpSwagger "github.com/swaggo/http-swagger"
	_"github.com/rupak26/Real-time-Leaderboard/docs"
	"github.com/rupak26/Real-time-Leaderboard/internal/middleware"
	"fmt"
	"net/http"
	"strconv"
)

type Server struct {
	// Storing as Dependecies 
	cnf              *config.Config
	laderbordHandler *laderbord_handler.Handler
	userHandler      *user_handler.Handler
}

func NewServer(
	cnf              *config.Config,
	laderbordHandler *laderbord_handler.Handler,
	userHandler      *user_handler.Handler,
) *Server {
	return &Server{
		 cnf              : cnf,
         laderbordHandler : laderbordHandler ,
		 userHandler      : userHandler,
	}
}

func (server *Server) Start() {
	manager := middleware.NewManager() 

	mux := http.NewServeMux() 

    mux.Handle("/swagger/", httpSwagger.WrapHandler)
	server.laderbordHandler.RegisterRouters(mux , manager)
	server.userHandler.RegisterRouters(mux , manager)
    
	//replacement 
	//mux.Handle("GET /users" ,middleware.Logger(http.HandlerFunc(handlers.GetUser)))
	//mux.Handle("GET /users/{id}",middleware.Logger(http.HandlerFunc(handlers.GetUserById)))
	//mux.Handle("POST /users" ,middleware.Logger(http.HandlerFunc(handlers.CreateUser)))
    
	addr := ":" + strconv.Itoa(int(server.cnf.HttpPort))
    
	fmt.Println("The server is running in port 8080" , addr) 


	err := http.ListenAndServe(addr , mux)
    
	if err != nil {
		fmt.Println("Server is facing issues" , err)
	}
}