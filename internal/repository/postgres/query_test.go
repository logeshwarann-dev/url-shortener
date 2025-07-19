package postgres

import "testing"

func TestInsertUrlIntoDB(t *testing.T) {
	TestConnectToSQL(t)
	err := InsertRecordIntoDB("https://www.youtube.com/", "r8Uygm", "0")
	if err != nil {
		t.Logf("error in insert record: %v", err.Error())
		t.Fatal("DB Record Insert failed!")
	}
	t.Log("DB Record Insert operation successful!")
}

func TestDeleteUrlFromDB(t *testing.T) {
	TestConnectToSQL(t)
	if err := DeleteRecordInDB("roUygm"); err != nil {
		t.Logf("error in delete record: %v", err.Error())
		t.Fatal("DB Record Delete failed!")
	}
	t.Log("DB Record Deletion operation successful!")
}
