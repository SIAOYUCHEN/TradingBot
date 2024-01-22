package main

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var jwtKey = []byte("TradingBot")

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		claims := &Claims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil {
			status := http.StatusInternalServerError
			message := "Error processing request"
			if err == jwt.ErrSignatureInvalid {
				status = http.StatusUnauthorized
				message = "Signature is invalid"
			} else if ve, ok := err.(*jwt.ValidationError); ok {
				if ve.Errors&jwt.ValidationErrorMalformed != 0 {
					message = "Token is malformed"
					status = http.StatusBadRequest
				} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
					message = "Token is either expired or not active yet"
					status = http.StatusUnauthorized
				} else {
					message = "Couldn't handle this token:"
				}
			}
			c.JSON(status, gin.H{"error": message})
			c.Abort()
			return
		}

		if !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token is invalid"})
			c.Abort()
			return
		}

		c.Next()
	}
}

// Login handles user authentication and returns a JWT token
// @Summary User login
// @Description Authenticate user and return JWT token if successful
// @Tags Authentication
// @Accept  json
// @Produce  json
// @Param loginBody body LoginRequest true "Login credentials"
// @Success 200 {object} map[string]interface{} "token: JWT Token"
// @Failure 400 {object} map[string]string "error: Bad request"
// @Failure 401 {object} map[string]string "error: Unauthorized"
// @Router /login [post]
func Login(c *gin.Context) {

	var request LoginRequest

	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, "Bad request")
		return
	}

	user, err := GetUserByUsername(db, request.Username)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusUnauthorized, "Credentials are not correct")
			return
		}
		c.JSON(http.StatusInternalServerError, "Database error")
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Credentials are not correct"})
		return
	}

	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		Username: request.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		c.JSON(http.StatusInternalServerError, "Could not create token")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})
}
