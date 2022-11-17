package db

import (
	"example/hello/db/models"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var database *gorm.DB

func Init() (*gorm.DB, error) {
	//dbFile := "sqlite.db"
	//dbPath := path.Join(config.GetConfig().GetString("database.path"), dbFile)
	//log.Println(dbPath)
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Panic("failed to connect database:", err)
		return nil, err
	}
	//localDb, _ := db.DB()
	//localDb.SetMaxIdleConns(10)
	database = db
	return database, nil
}

func Migrate() {
	database.AutoMigrate(&models.Podcast{}, &models.PodcastItem{}, &models.User{}, &models.Migration{})
}

func GetDb() *gorm.DB {
	return database
}
