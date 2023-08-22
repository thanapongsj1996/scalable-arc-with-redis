package postgres

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func InitDB(connectionStr string) {
	var err error

	db, err = gorm.Open(postgres.Open(connectionStr), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal(err)
	}

}

func GetDB() *gorm.DB {
	return db
}

func CloseDB() {
	_db, _ := db.DB()
	_db.Close()
}
