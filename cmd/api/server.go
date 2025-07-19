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

func init() {
	serverHost, serverPort = router.LoadAPIEnv()
	postgres.LoadDBEnv()
	if err := postgres.ConnectToSQL(); err != nil {
		log.Panic("DB connection error")
	}
}

func main() {
	router.Start(serverHost, serverPort)
}
