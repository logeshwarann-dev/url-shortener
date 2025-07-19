package postgres

import (
	"testing"

	"github.com/logeshwarann-dev/url-shortener/pkg/utils"
)

func setDBEnv() {
	utils.SetEnv("DB_USER", "admin")
	utils.SetEnv("DB_HOST", "localhost")
	utils.SetEnv("DB_PORT", "5432")
	utils.SetEnv("DB_NAME", "url_db")
	utils.SetEnv("DB_PWD", "admin@123")
	utils.SetEnv("DB_SSL_MODE", "disable")
	utils.SetEnv("TIME_ZONE", "Asia/Shanghai")
}
func TestConnectToSQL(t *testing.T) {
	setDBEnv()
	LoadDBEnv()
	if err := ConnectToSQL(); err != nil {
		t.Error("Error connecting to DB: ", err.Error())
		t.Fatalf("DB connection failed")
	}
	t.Logf("DB connection successful!")
}
