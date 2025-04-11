package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func initDB() {
	var err error

	DB, err = gorm.Open(sqlite.Open("chat.db"), &gorm.Config{})

	if err != nil {
		panic("fail to connect database")
	}

	DB.AutoMigrate(&User{})

}
