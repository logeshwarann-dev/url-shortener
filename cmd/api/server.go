package main

import (
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
}

func main() {
	router.Start(serverHost, serverPort)
}
