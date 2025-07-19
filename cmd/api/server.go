package main

import "github.com/logeshwarann-dev/url-shortener/internal/api/router"

func main() {
	router.Start("localhost", "8080")
}
