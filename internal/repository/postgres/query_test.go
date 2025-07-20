package postgres

import (
	"testing"

	"github.com/logeshwarann-dev/url-shortener/internal/models"
)

func TestInsertUrlIntoDB(t *testing.T) {
	TestConnectToSQL(t)
	err := InsertRecordIntoDB("https://www.youtube.com/", "r9Uygm", "0")
	if err != nil {
		t.Logf("error in insert record: %v", err.Error())
		t.Fatal("DB Record Insert failed!")
	}
	t.Log("DB Record Insert operation successful!")
}

func TestDeleteUrlFromDB(t *testing.T) {
	TestConnectToSQL(t)
	if err := DeleteRecordInDB("r9Uygm"); err != nil {
		t.Logf("error in delete record: %v", err.Error())
		t.Fatal("DB Record Delete failed!")
	}
	t.Log("DB Record Deletion operation successful!")
}

func TestFetchUrlFromDB(t *testing.T) {
	TestConnectToSQL(t)
	var urlRecord models.UrlInfo
	if err := FetchRecordFromDB("r9Uygm", &urlRecord); err != nil {
		t.Logf("error in delete record: %v", err.Error())
		t.Fatal("DB Record Delete failed!")
	}
	t.Log("URL Record: ", urlRecord)
	t.Log("DB Record Deletion operation successful!")

}
