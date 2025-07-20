package postgres

import (
	"fmt"

	"github.com/logeshwarann-dev/url-shortener/internal/models"
	"github.com/logeshwarann-dev/url-shortener/pkg/utils"
)

const tableName = "url_info"

func InsertRecordIntoDB(longUrl string, shortCode string, accessCount int) error {
	db := GetDBConn()

	rawQuery := BuildInsertQuery(tableName, longUrl, shortCode, utils.IntToStr(accessCount))
	tx := db.Exec(rawQuery)
	if tx.Error != nil {
		return fmt.Errorf("error while inserting record: %v", tx.Error.Error())
	}
	return nil
}

func UpdateRecordInDB(targetField string, newValue string, shortCode string) error {

	db := GetDBConn()
	isRowPresent, err := CheckIfRecordExists(shortCode)
	if err != nil {
		return err
	}
	if !isRowPresent {
		return fmt.Errorf("record doesn't exist in db")
	}
	rawQuery := BuildUpdateQuery(tableName, targetField, newValue, shortCode)
	tx := db.Exec(rawQuery)
	if tx.Error != nil {
		return fmt.Errorf("error while updating record: %v", tx.Error.Error())
	}
	return nil
}

func FetchRecordFromDB(shortCode string, urlStruct *models.UrlInfo) error {
	db := GetDBConn()
	isRowPresent, err := CheckIfRecordExists(shortCode)
	if err != nil {
		return err
	}
	if !isRowPresent {
		return fmt.Errorf("record doesn't exist in db")
	}
	rawQuery := BuildFetchQuery(tableName, shortCode)
	tx := db.Raw(rawQuery).Scan(urlStruct)
	if tx.Error != nil {
		return fmt.Errorf("error while fetching record: %v", tx.Error.Error())
	}
	return nil
}

func DeleteRecordInDB(shortCode string) error {
	db := GetDBConn()
	isRowPresent, err := CheckIfRecordExists(shortCode)
	if err != nil {
		return err
	}
	if !isRowPresent {
		return fmt.Errorf("record doesn't exist in db")
	}
	rawQuery := BuildDeleteQuery(tableName, shortCode)
	tx := db.Exec(rawQuery)
	if tx.Error != nil {
		return fmt.Errorf("error while deleting record: %v", tx.Error.Error())
	}
	return nil
}

func CheckIfRecordExists(shortCode string) (bool, error) {
	db := GetDBConn()
	selectQuery := BuildFetchQuery(tableName, shortCode)
	selectQuery = utils.RemoveCharFromString(selectQuery, ";", "")
	rawQuery := BuildRowCheckQuery(selectQuery)
	var found bool
	err := db.Raw(rawQuery).Row().Scan(&found)
	if err != nil {
		return false, fmt.Errorf("error while scanning record: %v", err.Error())
	}
	return found, nil
}
