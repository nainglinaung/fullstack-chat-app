package main

import "time"

type User struct {
	ID        string `gorm:"primaryKey"`
	UserName  string `gorm:"unique"`
	Password  string
	CreatedAt time.Time
}
