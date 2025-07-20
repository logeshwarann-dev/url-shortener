package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/logeshwarann-dev/url-shortener/internal/repository/postgres"
)

func DeleteShortURL(ctx *gin.Context) {
	targetShortCode := ctx.Param("shortCode")
	if len(targetShortCode) == 0 {
		log.Println("Recieved Empty Short Code")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Request"})
		return
	}

	err := postgres.DeleteRecordInDB(targetShortCode)
	if err != nil {
		log.Println("Error while deleting shortCode: ", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	ctx.JSON(http.StatusNoContent, gin.H{"message": "record has been deleted"})
}
