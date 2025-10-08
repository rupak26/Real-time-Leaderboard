package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"github.com/joho/godotenv"
)

type DBConfig struct {
	Host          string
	Port          int
	Name          string
	User          string
	Password      string
	EnableSSLMODE bool
}

type Config struct {
	Version        string 
	ServiceName    string 
	HttpPort       int64
	SecretKey      string
	DB            *DBConfig
}

var configuration *Config

func loadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	version := os.Getenv("VERSION") 
	
	if version == "" {
       log.Fatal("Version is required") 
	   os.Exit(1)
	}

	service_name := os.Getenv("SERVICE_NAME") 

	if service_name == "" {
		log.Fatal("Service Name is required")
		os.Exit(1)
	}

	http_port := os.Getenv("HTTP_PORT") 

	if http_port == "" {
		log.Fatal("Http Port is required") 
		os.Exit(1) 
	}
    
	port , err := strconv.ParseInt(http_port , 10 , 64)
    
	if err != nil {
		fmt.Println("Port must be integer")
		os.Exit(1)
	}

	jwtSecretkey := os.Getenv("JWT_SECRET")
    if jwtSecretkey == "" {
	   fmt.Println("JWTSecretkey is required")
	   os.Exit(1)
	}
    
	dbhost := os.Getenv("DB_HOST")
	if dbhost == "" {
	   fmt.Println("Host is required")
	   os.Exit(1)	
	}
	
	dbPort := os.Getenv("DB_PORT")
	dbPrt , err := strconv.ParseInt(dbPort , 10 , 64)
    
	if err != nil {
		fmt.Println("Port must be integer")
		os.Exit(1)
	}
	
	
	dbuser := os.Getenv("DB_USER")
	if dbuser == "" {
	   fmt.Println("User is required")
	   os.Exit(1)	
	}
	
	dbname := os.Getenv("DB_NAME")
	if dbname == "" {
	   fmt.Println("Name is required")
	   os.Exit(1)	
	}
	
	dbpassword := os.Getenv("DB_PASSWORD")
	if dbpassword == "" {
	   fmt.Println("Password is required")
	   os.Exit(1)	
	}
	
	enableSslMode := os.Getenv("ENABLE_SSL_MODE")
	
	enableSslModes , err := strconv.ParseBool(enableSslMode) 
    if err != nil {
		fmt.Println("Invalid Enalbel SSL Mode")
		os.Exit(1)
	}
   
    dbConfig := &DBConfig{
		Host: dbhost,
		Port: int(dbPrt),
		Name: dbname,
		User: dbuser,
		Password: dbpassword,
		EnableSSLMODE: enableSslModes,
	}
    
	configuration = &Config{
		Version: version,
		ServiceName: service_name,
		HttpPort: port,
		SecretKey: jwtSecretkey,
		DB:        dbConfig,
	}
}

func GetConfig() * Config {
	if configuration == nil {
	   loadConfig()
	}
	return configuration
}