package database

import (
	"fmt"

	"github.com/randijulio13/gofiber/pkg/logger"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func GetDatabase() *gorm.DB {
	return db
}

func OpenConnection(dsn string) *gorm.DB {
	log := logger.GetLogger()
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func GetDsn(dbHost string, dbPort string, dbUser string, dbPass string, dbName string) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPass, dbHost, dbPort, dbName)
}
