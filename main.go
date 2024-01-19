package main

import (
	_ "你的專案名稱/docs" // Swagger 文檔

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	router := gin.Default()

	router.POST("/login", Login) // JWT 登入路由

	// CRUD 路由
	apiRoutes := router.Group("/api", AuthMiddleware())
	{
		apiRoutes.POST("/create", Create)
		apiRoutes.GET("/read/:id", Read)
		apiRoutes.PUT("/update/:id", Update)
		apiRoutes.DELETE("/delete/:id", Delete)
	}

	// Swagger 文檔路由
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run(":8080")
}
