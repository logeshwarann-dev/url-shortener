package postgres

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/logeshwarann-dev/url-shortener/pkg/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	dbUser     string
	dbPwd      string
	dbHost     string
	dbPort     string
	dbName     string
	dbSslMode  string
	dbTimeZone string
	DbConn     *gorm.DB
)

func LoadDBEnv() {
	dbUser = utils.GetEnv("DB_USER")
	dbPwd = utils.GetEnv("DB_PWD")
	dbHost = utils.GetEnv("DB_HOST")
	dbPort = utils.GetEnv("DB_PORT")
	dbName = utils.GetEnv("DB_NAME")
	dbSslMode = utils.GetEnv("DB_SSL_MODE")
	dbTimeZone = utils.GetEnv("Time_Zone")
}

func GetDatabaseConnectionString() string {
	return fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=%s TimeZone=%v", dbUser, dbPwd, dbName, dbHost, dbPort, dbSslMode, dbTimeZone)
}

func ConnectToSQL() error {
	dsn := GetDatabaseConnectionString()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("Error connection to DB: ", err.Error())
		return err
	}
	dbase, err := db.DB()
	if err != nil {
		log.Println("Unable to get db instance: ", err.Error())
		return err
	}
	if err := PingCheck(dbase); err != nil {
		return err
	}
	log.Println("DB Connection successful!")
	DbConn = db
	return nil
}

func PingCheck(db *sql.DB) error {
	if err := db.Ping(); err != nil {
		return fmt.Errorf("unable to ping database: %v", err)
	}
	return nil
}

func GetDBConn() *gorm.DB {
	return DbConn
}
