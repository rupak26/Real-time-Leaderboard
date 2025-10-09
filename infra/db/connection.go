package db

import (
	"github.com/rupak26/Real-time-Leaderboard/config"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	//"golang.org/x/text/number"
)

func GetConnectionString(cnf *config.DBConfig) string {
	 connString := fmt.Sprintf(
		"user = %s password = %s host = %s port = %d dbname = %s",
		cnf.User , cnf.Password , cnf.Host , cnf.Port , cnf.Name,
	 )
	 if !cnf.EnableSSLMODE {
		connString += " sslmode=disable"
	 }
     return connString
}

func NewConnection(cnf *config.DBConfig) (*sqlx.DB , error) {
    dbSource := GetConnectionString(cnf)
	
	db , err := sqlx.Connect("postgres" , dbSource)
    
	if err != nil {
		fmt.Println(err) 
		return nil , err
	}
	return db , nil
}