package postgres

import "testing"

func TestInsertUrlIntoDB(t *testing.T) {
	TestConnectToSQL(t)
	err := InsertRecordIntoDB("https://www.youtube.com/", "r8Uygm", "0")
	if err != nil {
		t.Logf("error in insert record: %v", err.Error())
		t.Fatal("DB Insert failed!")
	}
	t.Log("DB Insert operation successful!")
}
