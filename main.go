package main

import (
	_ "TradingBot/docs" // Swagger 文档

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

// @title TradingBot API
// @version 1
// @description This is a sample server for TradingBot.
// @host localhost:8080
// @BasePath /
// @schemes http
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	router := gin.Default()

	router.POST("/login", Login) // JWT 登入路由

	apiRoutes := router.Group("/api", AuthMiddleware())
	{
		apiRoutes.POST("/create", CreateUser)
		apiRoutes.GET("/read/:id", ReadUser)
		apiRoutes.PUT("/update/email/:id", UpdateUserEmail)
		apiRoutes.DELETE("/delete/:id", DeleteUser)
		apiRoutes.GET("/users", GetAllUsers)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	dsn := "user:password@tcp(127.0.0.1:3306)/tradingbot?charset=utf8mb4&parseTime=True&loc=Local"

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	router.Run(":8080")
}
