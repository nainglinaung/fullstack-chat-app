package main

import "time"

type User struct {
	ID        string `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Username  string `gorm:"uniqueIndex; not null"`
	Password  string `gorm:"not null"`
	CreatedAt time.Time
}
