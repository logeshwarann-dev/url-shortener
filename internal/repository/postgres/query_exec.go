package postgres

import (
	"fmt"
	"log"

	"github.com/logeshwarann-dev/url-shortener/internal/models"
	"github.com/logeshwarann-dev/url-shortener/pkg/utils"
	"gorm.io/gorm"
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

func CreateTableIfNotExists() error {
	db := GetDBConn()
	tableExists, err := CheckIfTableExists(db, tableName)
	if err != nil {
		return err
	}
	if tableExists {
		log.Printf("%v already exists\n", tableName)
		return nil
	}
	creatTableQuery := BuildCreateTableQuery()
	tx := db.Exec(creatTableQuery)
	if tx.Error != nil {
		return fmt.Errorf("error while creating table: %v", tx.Error.Error())
	}
	return nil
}

func CheckIfTableExists(db *gorm.DB, table string) (bool, error) {
	var isTablePresent bool
	tableCheckQuery := BuildTableCheckQuery(table)
	tx := db.Raw(tableCheckQuery).Scan(&isTablePresent)
	if tx.Error != nil {
		return false, fmt.Errorf("error in scanning for table existence: %v", tx.Error.Error())
	}
	return isTablePresent, nil

}
