package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// @Summary Get all users
// @Description Get a list of all users
// @Tags Users
// @Produce  json
// @Success 200 {array} User "List of all users"
// @Failure 500 {object} map[string]string "error: Internal server error"
// @Router /api/users [get]
// @Security ApiKeyAuth
func GetAllUsers(c *gin.Context) {
	var users []UserAll
	if result := db.Select("ID", "Username", "Email").Find(&users); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving users"})
		return
	}
	c.JSON(http.StatusOK, users)
}

// @Summary CreateUser
// @Description CreateUser a new user
// @Tags Users
// @Accept  json
// @Produce  json
// @Param user body UserCreate true "User data"
// @Success 201 {object} User "User created"
// @Failure 400 "Bad request"
// @Router /api/create [post]
// @Security ApiKeyAuth
func CreateUser(c *gin.Context) {
	var request UserCreate

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		return
	}

	var existingUser User
	if err := db.Where("username = ?", request.Username).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username already exists"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Password encryption failed"})
		return
	}

	newUser := User{
		Username: request.Username,
		Password: string(hashedPassword),
		Email:    request.Email,
	}

	if result := db.Create(&newUser); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	userResponse := UserAll{
		Username: newUser.Username,
		Email:    newUser.Email,
	}

	c.JSON(http.StatusCreated, userResponse)
}

// @Summary Get user
// @Description Get a user by id
// @Tags Users
// @Produce  json
// @Param id path uint true "User ID"
// @Success 200 {object} User "User found"
// @Failure 404 "Not found"
// @Router /api/read/{id} [get]
// @Security ApiKeyAuth
func ReadUser(c *gin.Context) {
	id := c.Param("id")
	var user UserAll
	if result := db.Select("ID", "Username", "Email").First(&user, id); result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

// @Summary UpdateUserEmail
// @Description UpdateUserEmail a user by id
// @Tags Users
// @Accept  json
// @Produce  json
// @Param id path uint true "User ID"
// @Param user body UserEmailUpdate true "User data"
// @Success 200 "User updated"
// @Failure 404 "Not found"
// @Router /api/update/email/{id} [put]
// @Security ApiKeyAuth
func UpdateUserEmail(c *gin.Context) {
	id := c.Param("id")

	var request UserEmailUpdate
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request", "details": err.Error()})
		return
	}

	var existingUser User
	if err := db.First(&existingUser, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if result := db.Model(&User{}).Where("id = ?", id).Update("email", request.Email).Error; result != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User email updated"})
}

// @Summary Delete user
// @Description DeleteUser a user by id
// @Tags Users
// @Produce  json
// @Param id path uint true "User ID"
// @Success 200 {object} map[string]string "message: User deleted"
// @Failure 404 {object} map[string]string "error: User not found"
// @Failure 500 {object} map[string]string "error: Internal server error"
// @Router /api/delete/{id} [delete]
// @Security ApiKeyAuth
func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	var user User
	if err := db.First(&user, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if result := db.Delete(&user).Error; result != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}
