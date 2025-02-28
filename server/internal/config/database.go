package config

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDb() {
	databaseUrl, err := GetEnv("DB_URL")
	if err != nil {
		log.Fatal(err)
	}
	db, err := gorm.Open(postgres.Open(databaseUrl), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	DB = db
	fmt.Println("Successfully connected to db")
}

func GetDb() *gorm.DB {
	return DB
}
