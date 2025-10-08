package cmd

import (
	"fmt"
    "os"
	"github.com/rupak26/Real-time-Leaderboard/config"
	"github.com/rupak26/Real-time-Leaderboard/infra/db"
)

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
	
}