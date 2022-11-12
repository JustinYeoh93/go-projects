package config

import (
	"example/src/domains"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Database *gorm.DB

func DbConnect() {
	dsn := "host=localhost user=postgres password=password dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Tokyo"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	Database = db

	runMigrations()
}

func runMigrations() {
	Database.AutoMigrate(&domains.Event{})
}
