package db

import (
	"example/hello/config"
	"fmt"
	"log"
	"path"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var database *gorm.DB

func Init() (*gorm.DB, error) {
	dbFile := "sqlite.db"
	dbPath := path.Join(config.GetConfig().GetString("dbPath"), dbFile)
	log.Println(dbPath)
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		fmt.Println("db err:", err)
		return nil, err
	}
	localDb, _ := db.DB()
	localDb.SetMaxIdleConns(10)
	database = db
	return database, nil
}
