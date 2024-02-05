package main

import (
	"TradingBot/docs"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/mehdihadeli/go-mediatr"

	userCommand "TradingBot/application/handler/user/command"
	userQuery "TradingBot/application/handler/user/query"
	controller "TradingBot/controller"

	createUser "TradingBot/domain/createUser"
	deleteUser "TradingBot/domain/deleteUser"
	getAllUsers "TradingBot/domain/getAllUsers"
	getUser "TradingBot/domain/getUser"
	login "TradingBot/domain/login"
	updateUserEmail "TradingBot/domain/updateUserEmail"
	"TradingBot/infrastructure"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	ginSwagger "github.com/swaggo/echo-swagger"
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
	dsn := "user:password@tcp(127.0.0.1:3306)/tradingbot?charset=utf8mb4&parseTime=True&loc=Local"

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	echo := echo.New()

	echo.Validator = &controller.CustomValidator{Validator: validator.New()}

	userRepo := infrastructure.NewGormUserRepository(db)

	loginHandler := userCommand.NewLoginHandler(userRepo)

	createUserHandler := userCommand.NewCreateUserHandler(userRepo)

	getAllUsersHandler := userQuery.NewGetAllUsersHandler(userRepo)

	getUserHandler := userQuery.NewGetUserHandler(userRepo)

	deleteUserHandler := userCommand.NewDeleteUserHandler(userRepo)

	updateUserEmailHandler := userCommand.NewUpdateUserEmailHandler(userRepo)

	err := mediatr.RegisterRequestHandler[*login.LoginCommand, *login.LoginResponse](loginHandler)
	if err != nil {
		panic("LoginHandler Error")
	}

	err = mediatr.RegisterRequestHandler[*getAllUsers.GetAllUsersQuery, *getAllUsers.GetUserAllResponse](getAllUsersHandler)
	if err != nil {
		panic("GetAllUsersHandler Error")
	}

	err = mediatr.RegisterRequestHandler[*getUser.GetUserQuery, *getUser.GetUserResponse](getUserHandler)
	if err != nil {
		panic("GetUserHandler Error")
	}

	err = mediatr.RegisterRequestHandler[*createUser.CreateUserCommand, *createUser.CreateResponse](createUserHandler)
	if err != nil {
		panic("CreateUserHandler Error")
	}

	err = mediatr.RegisterRequestHandler[*deleteUser.DeleteUserCommand, *deleteUser.DeleteUserResponse](deleteUserHandler)
	if err != nil {
		panic("DeleteUserHandler Error")
	}

	err = mediatr.RegisterRequestHandler[*updateUserEmail.UpdateUserEmailCommand, *updateUserEmail.UpdateUserEmailResponse](updateUserEmailHandler)
	if err != nil {
		panic("UpdateUserEmailHandler Error")
	}

	controllerInstance := controller.NewUserController(echo)

	controller.MapUserRoutes(echo, controllerInstance)

	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Title = "TradingBot Service Api"
	docs.SwaggerInfo.Description = "TradingBot Service Api."

	echo.GET("/swagger/*any", ginSwagger.WrapHandler)

	go func() {
		if err := echo.Start(":8080"); err != nil {
			panic("Error starting Server")
		}
	}()

	<-ctx.Done()

	if err := echo.Shutdown(ctx); err != nil {
		panic("(Shutdown) err")
	}
}
