package domain

import "github.com/dgrijalva/jwt-go"

// Jwt Claims Struct
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
