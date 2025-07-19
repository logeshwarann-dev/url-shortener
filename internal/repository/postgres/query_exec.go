package postgres

import "fmt"

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

func FetchRecordFromDB() {

}

func DeleteRecordInDB() {

}

func FetchCountFromDB() {

}

func GetTotalEntriesInDB() {

}
