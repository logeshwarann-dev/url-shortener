package router

import (
	"github.com/gin-gonic/gin"
	"github.com/logeshwarann-dev/url-shortener/internal/api/handlers"
	"github.com/logeshwarann-dev/url-shortener/pkg/utils"
)

func LoadAPIEnv() (string, string) {
	host := utils.GetEnv("SERVER_HOST")
	port := utils.GetEnv("SERVER_PORT")
	return host, port
}

func Start(host string, port string) {
	router := gin.Default()
	addr := host + ":" + port
	RegisterRoutes(router)
	router.Run(addr)
}

func RegisterRoutes(router *gin.Engine) {
	router.POST("/shorten", handlers.CreateShortURL)
	router.GET("/shorten/:shortCode", handlers.RetrieveShortURL)
	router.PUT("/shorten/:shortCode", handlers.UpdateShortURL)
	router.DELETE("/shorten/:shortCode", handlers.DeleteShortURL)
	router.GET("/shorten/:shortCode/stats", handlers.RetrieveShortURLStats)
}
