package router

import (
	"github.com/gin-gonic/gin"
	"github.com/logeshwarann-dev/url-shortener/internal/api/handlers"
)

func Start(host string, port string) {
	router := gin.Default()
	addr := host + ":" + port
	RegisterRoutes(router)
	router.Run(addr)
}

func RegisterRoutes(router *gin.Engine) {
	router.POST("/shorten", handlers.CreateShortURL)
	router.GET("/shorten/:id", handlers.RetrieveShortURL)
	router.PUT("/shorten/:id", handlers.UpdateShortURL)
	router.DELETE("/shorten/:id", handlers.DeleteShortURL)
	router.GET("/shorten/:id/stats", handlers.RetrieveShortURLStats)
}
