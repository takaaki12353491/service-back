package database

import (
	"os"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

func NewConnection() *gorm.DB {
	DBMS := "mysql"
	USER := os.Getenv("DB_USER")
	PASS := os.Getenv("DB_PASS")
	PROTOCOL := "tcp(takaaki:3306)"
	DBNAME := "cln_arch"
	OPTION := "?parseTime=true&loc=Asia%2FTokyo"
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + OPTION
	db, err := gorm.Open(DBMS, CONNECT)
	if err != nil {
		log.Panicln(err)
	}
	return db
}
