package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Version        string 
	ServiceName    string 
	HttpPort       int64
}

var configuration Config

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

	configuration = Config{
		Version: version,
		ServiceName: service_name,
		HttpPort: port,
	}
}

func GetConfig() Config {
	loadConfig()
	return configuration
}