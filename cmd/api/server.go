package main

import (
	"log"

	"github.com/logeshwarann-dev/url-shortener/internal/api/router"
	"github.com/logeshwarann-dev/url-shortener/internal/repository/postgres"
)

var (
	serverHost string
	serverPort string
)

// const envFilePath = "D:\\Workspace\\GoLang\\RoadMapProjects\\url-shortener\\cmd\\api\\.env"

func init() {
	// utils.LoadEnv(envFilePath) // Only for Local Testing
	serverHost, serverPort = router.LoadAPIEnv()
	postgres.LoadDBEnv()
	if err := postgres.ConnectToSQL(); err != nil {
		log.Panicf("DB connection error: %v", err.Error())
		return
	}

	if err := postgres.CreateTableIfNotExists(); err != nil {
		log.Panicf("Error in table creation: %v", err.Error())
		return
	}
}

func main() {
	router.Start(serverHost, serverPort)
}
