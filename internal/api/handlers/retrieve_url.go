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

func RetrieveShortURL(ctx *gin.Context) {
	targetShortCode := ctx.Param("shortCode")
	if len(targetShortCode) == 0 {
		log.Println("Recieved Empty Short Code")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Request"})
		return
	}
	var urlRecord models.UrlInfo
	if err := postgres.FetchRecordFromDB(targetShortCode, &urlRecord); err != nil {
		log.Println("error while fetching url record from db: ", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	if utils.IsStringEmpty(urlRecord.ShortCode) {
		log.Println("No record found: ", urlRecord.ShortCode)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "record not found"})
		return
	}

	if err := UpdateAccessCount(urlRecord.AccessCount, urlRecord.ShortCode); err != nil {
		log.Println("Error in RetrieveShortURL handler: ", err.Error())
	}

	ctx.JSON(http.StatusOK, gin.H{"id": utils.IntToStr(urlRecord.Id), "url": urlRecord.Url, "shortCode": urlRecord.ShortCode, "createdAt": urlRecord.CreatedAt, "updatedAt": urlRecord.UpdatedAt})

}

func RetrieveShortURLStats(ctx *gin.Context) {
	targetShortCode := ctx.Param("shortCode")
	if utils.IsStringEmpty(targetShortCode) {
		log.Println("Recieved Empty Short Code")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Request"})
		return
	}
	var urlRecord models.UrlInfo
	if err := postgres.FetchRecordFromDB(targetShortCode, &urlRecord); err != nil {
		log.Println("error while fetching url record from db: ", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	if utils.IsStringEmpty(urlRecord.ShortCode) {
		log.Println("No record found: ", urlRecord.ShortCode)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "record not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"id": utils.IntToStr(urlRecord.Id), "url": urlRecord.Url, "shortCode": urlRecord.ShortCode, "createdAt": urlRecord.CreatedAt, "updatedAt": urlRecord.UpdatedAt, "accessCount": urlRecord.AccessCount})

}

func UpdateAccessCount(currentAccessCnt int, shortCode string) error {
	newCount := utils.IntToStr(currentAccessCnt + 1)
	if err := postgres.UpdateRecordInDB(models.AccessCountField, newCount, shortCode); err != nil {
		return fmt.Errorf("failed updating access count: %v", err.Error())
	}
	return nil
}
