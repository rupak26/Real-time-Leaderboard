package cmd

import (
	"fmt"
	"os"
	"github.com/rupak26/Real-time-Leaderboard/config"
	"github.com/rupak26/Real-time-Leaderboard/infra/db"
	"github.com/rupak26/Real-time-Leaderboard/internal/handler/laderbord_handler"
	"github.com/rupak26/Real-time-Leaderboard/internal/handler/user_handler"
	"github.com/rupak26/Real-time-Leaderboard/internal/middleware"
	"github.com/rupak26/Real-time-Leaderboard/laderbord"
	"github.com/rupak26/Real-time-Leaderboard/repository"
	"github.com/rupak26/Real-time-Leaderboard/users"
	"github.com/rupak26/Real-time-Leaderboard/internal/handler"
)


// @title Real-time Leaderboard API
// @version 1.0
// @description API for managing leaderboard and users
// @host localhost:8080
func Server() {
	cnf := config.GetConfig() 

	dbCon , err := db.NewConnection(cnf.DB)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
    
    err = db.MigrateDB(dbCon , "./migrations")
    if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	rdb := db.InitRedis(*cnf.RedisCon)
	laderRepo := repository.NewLaderbordRepo(rdb)
	userRepo := repository.NewUserRepo(dbCon)
   
    //domains
    usrSvc := users.NewService(userRepo)
    laderSvc := laderbord.NewService(laderRepo)

	middleware := middleware.NewMiddleware(cnf)
	
    laderBordHandler := laderbord_handler.NewHandler(middleware , laderSvc)
	userHandler := user_handler.NewHandler(usrSvc)

    server := handler.NewServer(cnf , laderBordHandler , userHandler)
	server.Start()
}