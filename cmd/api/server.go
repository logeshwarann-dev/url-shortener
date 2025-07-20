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
		log.Panic("DB connection error")
		return
	}

	postgres.CreateTableIfNotExists()
}

func main() {
	router.Start(serverHost, serverPort)
}
