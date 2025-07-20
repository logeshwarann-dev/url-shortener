package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/logeshwarann-dev/url-shortener/internal/repository/postgres"
	"github.com/logeshwarann-dev/url-shortener/pkg/utils"
)

func DeleteShortURL(ctx *gin.Context) {
	targetShortCode := ctx.Param("shortCode")
	if utils.IsStringEmpty(targetShortCode) {
		log.Println("Recieved Empty Short Code")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Request"})
		return
	}

	err := postgres.DeleteRecordInDB(targetShortCode)
	if err != nil {
		if utils.StringContains(err.Error(), "record doesn't exist") {
			log.Println("Error in DeleteShortURL handler: ", err.Error())
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
			return
		}
		log.Println("Error while deleting shortCode: ", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	ctx.JSON(http.StatusNoContent, gin.H{"message": "record has been deleted"})
}
