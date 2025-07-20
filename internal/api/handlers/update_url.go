package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/logeshwarann-dev/url-shortener/internal/models"
	"github.com/logeshwarann-dev/url-shortener/internal/repository/postgres"
	"github.com/logeshwarann-dev/url-shortener/pkg/utils"
)

func UpdateShortURL(ctx *gin.Context) {
	targetShortCode := ctx.Param("shortCode")
	if utils.IsStringEmpty(targetShortCode) {
		log.Println("Recieved Empty Short Code")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Request"})
		return
	}

	var newUrlDetails models.UrlInfo
	if err := ctx.ShouldBindJSON(&newUrlDetails); err != nil {
		log.Println("Error in binding JSON of UpdateShortURL request: ", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	if err := UpdateOrginalUrl(newUrlDetails.Url, targetShortCode); err != nil {
		if utils.StringContains(err.Error(), "record doesn't exist") {
			log.Println("Error in UpdateOrginalUrl handler: ", err.Error())
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
			return
		}
		log.Println("Error in UpdateOrginalUrl handler: ", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	if err := UpdateAccessCount(DefaultAccessCount, targetShortCode); err != nil {
		log.Println("Error in RetrieveShortURL handler: ", err.Error())
	}

	if err := postgres.FetchRecordFromDB(targetShortCode, &newUrlDetails); err != nil {
		log.Println("error while fetching url record from db: ", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"id": utils.IntToStr(newUrlDetails.Id), "url": newUrlDetails.Url, "shortCode": newUrlDetails.ShortCode, "createdAt": newUrlDetails.CreatedAt, "updatedAt": newUrlDetails.UpdatedAt})

}

func UpdateOrginalUrl(newUrl string, shortCode string) error {
	if err := postgres.UpdateRecordInDB(models.UrlField, newUrl, shortCode); err != nil {
		return fmt.Errorf("failed updating access count: %v", err.Error())
	}
	return nil
}
