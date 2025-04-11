package db

import "gorm.io/gorm"

type GormDB struct {
	DB *gorm.DB
}

func (g *GormDB) CreateUser(user *User) error {
	return g.DB.Create(user).Error
}

func (g *GormDB) GetUserByUsername(username string) (*User, error) {
	var user User
	if err := g.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
