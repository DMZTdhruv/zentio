package config

import (
	"fmt"
	"log"

	"github.com/personal-project/zentio/internal/schema"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDb() {
	databaseUrl, err := GetEnv("DOCKER_DB_URL")
	if err != nil {
		log.Fatal(err)
	}
	db, err := gorm.Open(postgres.Open(databaseUrl), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&schema.User{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
	DB = db
	fmt.Println("Successfully connected to db")
}

func GetDb() *gorm.DB {
	return DB
}
