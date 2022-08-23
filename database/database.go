package database

import (
	"log"

	"app/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Instance *gorm.DB
var err error

func Connect(connectionString string) {
	Instance, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		panic("Cannot connect to DB")
	}
	log.Println("Connected to Database...")
}

func InitialiseDB() {
	Instance.AutoMigrate(&models.OU{})
	Instance.AutoMigrate(&models.Endpoint{})
	Instance.AutoMigrate(&models.Credential{})
	Instance.AutoMigrate(&models.User{})
	log.Println("Database Migration Completed...")
}
