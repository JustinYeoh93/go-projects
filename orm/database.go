package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type database struct {
	db *gorm.DB
}

func newDatabase() *database {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return &database{
		db: db,
	}
}

func (d *database) initialMigration() {
	d.db.AutoMigrate(&User{})
}
