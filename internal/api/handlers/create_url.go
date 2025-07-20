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
  "url": "https://www.example.com/some/long/url"
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
const DefaultAccessCount = 0

func CreateShortURL(ctx *gin.Context) {
	var urlDetails models.UrlInfo

	if err := ctx.ShouldBindJSON(&urlDetails); err != nil {
		log.Println("Error while binding json: ", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	urlDetails.ShortCode = GenerateUniqueCode(urlDetails.Url)
	err := postgres.InsertRecordIntoDB(urlDetails.Url, urlDetails.ShortCode, DefaultAccessCount)
	if err != nil {
		log.Println("Error in DB Record Insertion: ", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	if err := postgres.FetchRecordFromDB(urlDetails.ShortCode, &urlDetails); err != nil {
		log.Println("error while fetching url record from db: ", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	if utils.IsStringEmpty(urlDetails.ShortCode) {
		log.Println("record not properly inserted: ", urlDetails.ShortCode)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"id": utils.IntToStr(urlDetails.Id), "url": urlDetails.Url, "shortCode": urlDetails.ShortCode, "createdAt": urlDetails.CreatedAt, "updatedAt": urlDetails.UpdatedAt})

}

func GenerateUniqueCode(url string) string {
	shortCode := utils.GetShortCode(ShortCodeLength)
	return shortCode
}
