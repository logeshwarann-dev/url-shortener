package main

import (
	"log"

	"github.com/logeshwarann-dev/url-shortener/internal/api/router"
	"github.com/logeshwarann-dev/url-shortener/internal/repository/postgres"
	"github.com/logeshwarann-dev/url-shortener/pkg/utils"
)

var (
	serverHost string
	serverPort string
)

const envFilePath = "D:\\Workspace\\GoLang\\RoadMapProjects\\url-shortener\\cmd\\api\\.env"

func init() {
	utils.LoadEnv(envFilePath)
	serverHost, serverPort = router.LoadAPIEnv()
	postgres.LoadDBEnv()
	if err := postgres.ConnectToSQL(); err != nil {
		log.Panic("DB connection error")
		return
	}
}

func main() {
	router.Start(serverHost, serverPort)
}
