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

	tradeCommand "TradingBot/application/handler/trade/command"
	tradeQuery "TradingBot/application/handler/trade/query"
	userCommand "TradingBot/application/handler/user/command"
	userQuery "TradingBot/application/handler/user/query"

	controller "TradingBot/controller"

	createUser "TradingBot/domain/user/createUser"
	deleteUser "TradingBot/domain/user/deleteUser"
	getAllUsers "TradingBot/domain/user/getAllUsers"
	getUser "TradingBot/domain/user/getUser"
	login "TradingBot/domain/user/login"
	updateUserEmail "TradingBot/domain/user/updateUserEmail"

	createTrade "TradingBot/domain/trade/createTrade"
	deleteTrade "TradingBot/domain/trade/deleteTrade"
	getAllTrade "TradingBot/domain/trade/getAllTrade"
	getTrade "TradingBot/domain/trade/getTrade"
	"TradingBot/infrastructure"

	"github.com/go-playground/validator"
	"github.com/go-redis/redis/v8"
	"github.com/labstack/echo/v4"
	ginSwagger "github.com/swaggo/echo-swagger"
)

var (
	db  *gorm.DB
	err error
	rdb *redis.Client
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

	rdb = initRedis()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	echo := echo.New()

	customValidator := controller.NewCustomValidator(validator.New())

	echo.Validator = customValidator

	userRepo := infrastructure.NewGormUserRepository(db)

	tradeRepo := infrastructure.NewRedisTradeRepository(rdb)

	loginHandler := userCommand.NewLoginHandler(userRepo)

	createUserHandler := userCommand.NewCreateUserHandler(userRepo)

	getAllUsersHandler := userQuery.NewGetAllUsersHandler(userRepo)

	getUserHandler := userQuery.NewGetUserHandler(userRepo)

	deleteUserHandler := userCommand.NewDeleteUserHandler(userRepo)

	updateUserEmailHandler := userCommand.NewUpdateUserEmailHandler(userRepo)

	createTradeHandler := tradeCommand.NewCreateTradeHandler(tradeRepo)

	getTradeHandler := tradeQuery.NewGetTradeHandler(tradeRepo)

	getAllTradeHandler := tradeQuery.NewGetAllTradeHandler(tradeRepo)

	deleteTradeHandler := tradeCommand.NewDeleteTradeHandler(tradeRepo)

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

	err = mediatr.RegisterRequestHandler[*createTrade.CreateTradeCommand, *createTrade.CreateTradeResponse](createTradeHandler)
	if err != nil {
		panic("CreateTradeHandler Error")
	}

	err = mediatr.RegisterRequestHandler[*getTrade.GetTradeQuery, *getTrade.GetTradeResponse](getTradeHandler)
	if err != nil {
		panic("GetTradeHandler Error")
	}

	err = mediatr.RegisterRequestHandler[*getAllTrade.GetAllTradeQuery, *getAllTrade.GetAllTradeResponse](getAllTradeHandler)
	if err != nil {
		panic("GetAllTradeHandler Error")
	}

	err = mediatr.RegisterRequestHandler[*deleteTrade.DeleteTradeCommand, *deleteTrade.DeleteTradeResponse](deleteTradeHandler)
	if err != nil {
		panic("DeleteTradeHandler Error")
	}

	baseController := controller.NewBaseController(customValidator)

	controllerInstance := controller.NewUserController(echo, baseController)

	tradeInstance := controller.NewTradeController(echo, baseController)

	controller.MapUserRoutes(echo, controllerInstance)

	controller.MapTradeRoutes(echo, tradeInstance)

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

func initRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	ctx := context.Background()
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		panic("Unable to connect to Redis: " + err.Error())
	}

	return rdb
}
