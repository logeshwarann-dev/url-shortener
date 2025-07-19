package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/logeshwarann-dev/url-shortener/internal/models"
	"github.com/logeshwarann-dev/url-shortener/internal/repository/postgres"
	"github.com/logeshwarann-dev/url-shortener/pkg/utils"
)

/*

POST /shorten

{
  "url": "https://www.example.com/some/updated/url"
}

Response:
{
  "id": "1",
  "url": "https://www.example.com/some/updated/url",
  "shortCode": "abc123",
  "createdAt": "2021-09-01T12:00:00Z",
  "updatedAt": "2021-09-01T12:30:00Z"
}
*/

const ShortCodeLength = 6

func CreateShortURL(ctx *gin.Context) {
	var urlDetails models.UrlInfo

	if err := ctx.ShouldBindJSON(&urlDetails); err != nil {
		log.Println("Error while binding json: ", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
	}

	urlDetails.ShortCode = GenerateUniqueCode(urlDetails.Url)
	err := postgres.InsertRecordIntoDB(urlDetails.Url, urlDetails.ShortCode, utils.IntToStr(urlDetails.AccessCount))
	if err != nil {
		log.Println("Error in DB Record Insertion: ", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "New Short URL created successfully", "ShortenUrl": urlDetails})
}

func GenerateUniqueCode(url string) string {
	shortCode := utils.GetShortCode(ShortCodeLength)
	return shortCode
}
