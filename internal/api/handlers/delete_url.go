package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/logeshwarann-dev/url-shortener/internal/models"
	"github.com/logeshwarann-dev/url-shortener/internal/repository/postgres"
)

func DeleteShortURL(ctx *gin.Context) {
	var deleteReq models.UrlInfo

	if err := ctx.ShouldBindJSON(&deleteReq); err != nil {
		log.Println("Error binding json: ", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
	}

	err := postgres.DeleteRecordInDB(deleteReq.ShortCode)
	if err != nil {
		log.Println("Error while deleting shortCode: ", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
	}
	ctx.JSON(http.StatusAccepted, gin.H{"message": "record has been deleted"})
}
