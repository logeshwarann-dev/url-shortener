package postgres

import "fmt"

func BuildInsertQuery(table string, longUrl string, shortCode string, accessCount string) string {
	return fmt.Sprintf("INSERT INTO %v (url, short_code, access_count) VALUES ('%v', '%v', '%v') ;", table, longUrl, shortCode, accessCount)
}

func BuildDeleteQuery() {

}

func BuildFetchQuery() {

}

func BuildUpdateQuery() {

}
