package database

import (
	"os"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

func NewConnection() *gorm.DB {
	DBMS := "mysql"
	USER := os.Getenv("SERVICE_MYSQL_USER")
	PASS := os.Getenv("SERVICE_MYSQL_PASSWORD")
	PROTOCOL := "tcp(localhost:3306)"
	DBNAME := os.Getenv("SERVICE_MYSQL_DATABASE")
	OPTION := "?parseTime=true&loc=Asia%2FTokyo"
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + OPTION
	db, err := gorm.Open(DBMS, CONNECT)
	if err != nil {
		log.Panicln(err)
	}
	return db
}
