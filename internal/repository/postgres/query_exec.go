package postgres

import (
	"fmt"

	"github.com/logeshwarann-dev/url-shortener/internal/models"
)

const tableName = "url_info"

func InsertRecordIntoDB(longUrl string, shortCode string, accessCount string) error {
	db := GetDBConn()
	rawQuery := BuildInsertQuery(tableName, longUrl, shortCode, accessCount)
	tx := db.Exec(rawQuery)
	if tx.Error != nil {
		return fmt.Errorf("error while inserting record: %v", tx.Error.Error())
	}
	return nil
}

func UpdateRecordInDB() {

}

func FetchRecordFromDB(shortCode string, urlStruct *models.UrlInfo) error {
	db := GetDBConn()
	rawQuery := BuildFetchQuery(tableName, shortCode)
	tx := db.Raw(rawQuery).Scan(urlStruct)
	if tx.Error != nil {
		return fmt.Errorf("error while fetching record: %v", tx.Error.Error())
	}
	return nil
}

func DeleteRecordInDB(shortCode string) error {
	db := GetDBConn()
	rawQuery := BuildDeleteQuery(tableName, shortCode)
	tx := db.Exec(rawQuery)
	if tx.Error != nil {
		return fmt.Errorf("error while deleting record: %v", tx.Error.Error())
	}
	return nil
}

func FetchCountFromDB() {

}

func GetTotalEntriesInDB() {

}
