package database

import (
	"golang-crud-rest-api/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Instance *gorm.DB

var err error

func Connect(connectionString string) {
	Instance, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{SkipDefaultTransaction: true})
	if err != nil {
		log.Fatal(err)
		panic("Connot connect to database!")
	}
	log.Println("Connected to database...")
}

func Migrate() {
	Instance.AutoMigrate(&models.Product{})
	log.Println("Database Migration Completed...")
}
