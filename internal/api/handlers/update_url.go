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

}

func UpdateOrginalUrl(newUrl string, shortCode string) error {

	if err := postgres.UpdateRecordInDB(models.UrlField, newUrl, shortCode); err != nil {
		return fmt.Errorf("failed updating access count: %v", err.Error())
	}
	return nil
}
