package main

import (
	_ "service-back/docs"
	"service-back/infra/server"
	"time"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	log "github.com/sirupsen/logrus"
)

// @title Service API
// @version 1.0
// @description This is a personal project
// @license.name takaaki12353491
// @host localhost:8080
// @BasePath
func main() {
	// time zone setting
	const LOCATION = "Asia/Tokyo"
	loc, err := time.LoadLocation(LOCATION)
	if err != nil {
		loc = time.FixedZone(LOCATION, 9*60*60)
	}
	time.Local = loc
	// Log settings
	log.SetFormatter(&log.JSONFormatter{
		PrettyPrint: true,
	})
	log.SetReportCaller(true)

	server.Start()
}
