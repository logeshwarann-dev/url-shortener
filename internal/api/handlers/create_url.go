package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/logeshwarann-dev/url-shortener/internal/models"
	"github.com/logeshwarann-dev/url-shortener/pkg/utils"
)

/*

POST /shorten

{
	url: "https://www.google.com"
}
*/

var (
	IdCounter int
)

func CreateShortURL(ctx *gin.Context) {
	var urlDetails models.UrlInfo

	if err := ctx.ShouldBindJSON(&urlDetails); err != nil {
		log.Fatal("Error while binding json: ", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
	}

	urlDetails.ShortCode = GenerateUniqueCode(urlDetails.Url)

	ctx.JSON(http.StatusCreated, gin.H{"message": "New Short URL created successfully"})
}

func GenerateUniqueCode(url string) string {
	shortCode := utils.EncodeString(url)
	return shortCode
}
