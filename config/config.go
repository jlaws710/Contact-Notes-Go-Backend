package config

import (
	"fmt"
	"log"
	"notes-system/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	dsn := "postgres://postgres:postgres@localhost:5432/contacts"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	fmt.Println("Database connected!")

	DB.AutoMigrate(&model.Contact{}, &model.Note{})
}
