package main

import (
	"github.com/dgrijalva/jwt-go"
)

// Jwt Claims Struct
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// Login Api Input Struct
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// User Struct In Database
type User struct {
	ID       uint `gorm:"primaryKey"`
	Username string
	Password string
	Email    *string
}

// User All Struct
type UserAll struct {
	ID       uint `gorm:"primaryKey"`
	Username string
	Email    *string
}

type UserCreate struct {
	Username string
	Password string
	Email    *string
}

type UserEmailUpdate struct {
	Email string `json:"email" binding:"required,email"`
}

func (UserAll) TableName() string {
	return "users"
}
