package controller

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/mehdihadeli/go-mediatr"

	createUser "TradingBot/domain/createUser"
	deleteUser "TradingBot/domain/deleteUser"
	getAllUsers "TradingBot/domain/getAllUsers"
	getUser "TradingBot/domain/getUser"
	login "TradingBot/domain/login"
	updateUserEmail "TradingBot/domain/updateUserEmail"
)

type CustomValidator struct {
	Validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}

type UserController struct {
	echo *echo.Echo
}

func NewUserController(echo *echo.Echo) *UserController {
	return &UserController{echo: echo}
}

// Login handles the user login request.
// @Summary User Login
// @Description Handles user login and returns a JWT token
// @Tags Authentication
// @Accept  json
// @Produce  json
// @Param loginBody body login.LoginCommand true "Login Credentials"
// @Success 200 {object} login.LoginResponse "JWT Token returned on successful authentication"
// @Failure 400 {object} string "Bad Request"
// @Failure 401 {object} string "Unauthorized"
// @Router /api/v1/login [post]
func (uc *UserController) Login() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		request := &login.LoginCommand{}
		if err := ctx.Bind(request); err != nil {
			return ctx.JSON(http.StatusBadRequest, "Bad request")
		}

		if err := ctx.Validate(request); err != nil {
			return ctx.JSON(http.StatusBadRequest, err.Error())
		}

		result, err := mediatr.Send[*login.LoginCommand, *login.LoginResponse](ctx.Request().Context(), request)
		if err != nil {
			return ctx.JSON(http.StatusUnauthorized, err.Error())
		}

		return ctx.JSON(http.StatusOK, result)
	}
}

// GetAllUsers returns a list of all users.
// @Summary Get all users
// @Description Retrieve a list of all users
// @Tags Users
// @Accept  json
// @Produce  json
// @Success 200 {array} getAllUsers.GetUserAllResponse
// @Router /api/v1/users [get]
// @Security ApiKeyAuth
func (uc *UserController) GetAllUsers() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		request := &getAllUsers.GetAllUsersQuery{}
		if err := ctx.Bind(request); err != nil {
			return ctx.JSON(http.StatusBadRequest, "Bad request")
		}

		result, err := mediatr.Send[*getAllUsers.GetAllUsersQuery, *getAllUsers.GetUserAllResponse](ctx.Request().Context(), request)
		if err != nil {
			return ctx.JSON(http.StatusUnauthorized, err.Error())
		}

		return ctx.JSON(http.StatusOK, result)
	}
}

// @Summary Get User by ID
// @Description Retrieve a single user by ID
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path uint true "User ID"
// @Success 200 {object} getUser.UserResponse "Successful retrieval of user information"
// @Failure 400 {string} string "Bad request"
// @Failure 401 {string} string "Unauthorized"
// @Failure 404 {string} string "User not found"
// @Router /api/v1/user/{id} [get]
// @Security ApiKeyAuth
func (uc *UserController) GetUser() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		idStr := ctx.Param("id")

		if strings.HasPrefix(idStr, "/") {
			idStr = idStr[1:]
		}

		id, err := strconv.Atoi(idStr)
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, "Invalid user ID")
		}

		request := &getUser.GetUserQuery{
			Id: uint(id),
		}

		result, err := mediatr.Send[*getUser.GetUserQuery, *getUser.GetUserResponse](ctx.Request().Context(), request)
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, err.Error())
		}

		return ctx.JSON(http.StatusOK, result)
	}
}

// CreateUser creates a new user
// @Summary Create new user
// @Description Create a new user with username, password, and email
// @Tags user
// @Accept json
// @Produce json
// @Param createUserCommand body createUser.CreateUserCommand true "Create User Command"
// @Success 200 {object} createUser.CreateResponse "User created successfully"
// @Failure 400 {string} string "Bad request - invalid input"
// @Failure 401 {string} string "Unauthorized - invalid credentials"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/v1/create/user [post]
// @Security ApiKeyAuth
func (uc *UserController) CreateUser() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		request := &createUser.CreateUserCommand{}
		if err := ctx.Bind(request); err != nil {
			return ctx.JSON(http.StatusBadRequest, "Bad request")
		}

		if request.Email != nil && *request.Email == "" {
			return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "Email cannot be an empty string"})
		}

		if err := ctx.Validate(request); err != nil {
			return ctx.JSON(http.StatusBadRequest, err.Error())
		}

		result, err := mediatr.Send[*createUser.CreateUserCommand, *createUser.CreateResponse](ctx.Request().Context(), request)
		if err != nil {
			return ctx.JSON(http.StatusUnauthorized, err.Error())
		}

		return ctx.JSON(http.StatusOK, result)
	}
}

// DeleteUser removes a user by ID
// @Summary Delete a user
// @Description Deletes a user with the specified ID from the database
// @Tags users
// @Accept json
// @Produce json
// @Param id path uint true "User ID"
// @Success 200 {object} deleteUser.DeleteUserResponse "User deleted successfully"
// @Failure 400 {string} string "Invalid user ID"
// @Failure 404 {string} string "User not found"
// @Failure 500 {string} string "Internal Server Error - Error deleting user"
// @Router /api/v1/delete/user/{id} [delete]
// @Security ApiKeyAuth
func (uc *UserController) DeleteUser() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		idStr := ctx.Param("id")

		if strings.HasPrefix(idStr, "/") {
			idStr = idStr[1:]
		}

		id, err := strconv.Atoi(idStr)
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, "Invalid user ID")
		}

		request := &deleteUser.DeleteUserCommand{
			Id: uint(id),
		}

		result, err := mediatr.Send[*deleteUser.DeleteUserCommand, *deleteUser.DeleteUserResponse](ctx.Request().Context(), request)
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, err.Error())
		}

		return ctx.JSON(http.StatusOK, result)
	}
}

// UpdateUserEmail update user email by ID
// @Summary Update user email
// @Description Update user email by user ID
// @Tags Users
// @Accept json
// @Produce json
// @Param id path uint true "User ID"
// @Param email body updateUserEmail.Email true "Email Object"
// @Success 200 {object} string "Email updated successfully"
// @Failure 400 {object} string "Invalid request"
// @Failure 404 {object} string "User not found"
// @Failure 500 {object} string "Internal server error"
// @Router /api/v1/users/{id}/email [put]
// @Security ApiKeyAuth
func (uc *UserController) UpdateUserEmail() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, "Invalid user ID")
		}

		var request updateUserEmail.UpdateUserEmailCommand
		if err := ctx.Bind(&request); err != nil {
			return ctx.JSON(http.StatusBadRequest, "Invalid request body")
		}
		request.Id = uint(id)

		if err := ctx.Validate(request); err != nil {
			return ctx.JSON(http.StatusBadRequest, err.Error())
		}

		result, err := mediatr.Send[*updateUserEmail.UpdateUserEmailCommand, *updateUserEmail.UpdateUserEmailResponse](ctx.Request().Context(), &request)
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, err.Error())
		}

		return ctx.JSON(http.StatusOK, result)
	}
}
