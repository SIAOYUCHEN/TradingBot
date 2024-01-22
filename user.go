package main

import (
	"gorm.io/gorm"
)

func GetUserByUsername(db *gorm.DB, username string) (User, error) {
	var user User
	result := db.Where("Username = ?", username).First(&user)
	return user, result.Error
}
