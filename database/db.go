package database

import (
	"final-project-acgm/models"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func StartDB() *gorm.DB {
	var (
		host     = os.Getenv("DB_HOST")
		user     = os.Getenv("DB_USER")
		password = os.Getenv("DB_PASS")
		dbPort   = os.Getenv("DB_PORT")
		dbname   = os.Getenv("DB_NAME")
		timeZone = os.Getenv("DB_TIMEZONE")
		dsn      = ""
		err      error
	)

	dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s", host, user, password, dbname, dbPort, timeZone)

	if db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{FullSaveAssociations: true}); err != nil {
		log.Fatal("Error connecting to database: ", err)
	}

	if err = db.AutoMigrate(models.User{}, models.Locker{}, models.LockerRent{}, models.LockerRentDetail{}); err != nil {
		log.Fatal("Error migrating database: ", err.Error())
	}

	return db
}

func GetDB() *gorm.DB {
	if os.Getenv("DB_DEBUG") == "true" {
		db = db.Debug()
	}

	return db
}
