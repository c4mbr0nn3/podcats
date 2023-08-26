package db

import (
	"example/hello/config"
	"example/hello/db/models"
	"log"
	"path"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var database *gorm.DB

func Init() (*gorm.DB, error) {
	dbFile := config.GetConfig().GetString("database.file")
	dbPath := path.Join(config.GetConfig().GetString("database.path"), dbFile)
	log.Println(dbPath)
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		log.Panic("failed to connect database:", err)
		return nil, err
	}

	database = db
	return database, nil
}

func Migrate() {
	database.AutoMigrate(&models.Podcast{}, &models.PodcastItem{}, &models.User{}, &models.Migration{}, &models.Notification{})
	createAdminUserIfNotPresent()
}

func GetDb() *gorm.DB {
	return database
}

func createAdminUserIfNotPresent() {
	var user models.User
	database.First(&user, "username = ?", "root")
	if user.ID == 0 {
		log.Println("Creating admin user.")
		user = models.User{
			Email:    "admin@email.com",
			Name:     "I am",
			Surname:  "Root",
			Username: "root",
			Password: "changeme",
			IsAdmin:  true,
		}
		user.HashPassword()
		database.Create(&user)
	} else {
		log.Println("Admin user already present. Skipping creation.")
	}
}
