package postgres

import "fmt"

func BuildInsertQuery(table string, longUrl string, shortCode string, accessCount string) string {
	return fmt.Sprintf("INSERT INTO %v (url, short_code, access_count) VALUES ('%v', '%v', '%v') ;", table, longUrl, shortCode, accessCount)
}

func BuildDeleteQuery(table string, shortCode string) string {
	return fmt.Sprintf("DELETE FROM %v WHERE short_code = '%v' ;", table, shortCode)
}

func BuildFetchQuery(table string, shortCode string) string {
	return fmt.Sprintf("SELECT * FROM %v WHERE short_code = '%v' ;", table, shortCode)
}

func BuildRowCheckQuery(selectQuery string) string {

	return fmt.Sprintf("SELECT EXISTS(%v) AS FOUND ;", selectQuery)
}

func BuildUpdateQuery(table string, targetField string, updatedValue string, shortCode string) string {
	return fmt.Sprintf("UPDATE %v SET %v = '%v' WHERE short_code = '%v' ;", table, targetField, updatedValue, shortCode)
}
