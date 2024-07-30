package rest

import (
	"strconv"

	interfaces "github.com/Cthulhu-tech/microservice/internal/application/interface"
	"github.com/Cthulhu-tech/microservice/internal/interface/api/rest/dto/mapper"
	"github.com/Cthulhu-tech/microservice/internal/interface/api/rest/dto/request"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	"net/http"
)

type UserController struct {
	service interfaces.UserService
}

func NewUserController(e *echo.Echo, service interfaces.UserService) *UserController {
	controller := &UserController{
		service: service,
	}

	e.POST("/api/v1/user", controller.CreateUserController)
	e.GET("/api/v1/user", controller.GetAllUsersController)
	e.GET("/api/v1/user/:id", controller.GetUserByIdController)
	e.PUT("/api/v1/user", controller.PutUserController)
	e.DELETE("/api/v1/user/:id", controller.DeleteUserController)

	return controller
}

func (sc *UserController) CreateUserController(c echo.Context) error {
	var createUserRequest request.CreateUserRequest

	if err := c.Bind(&createUserRequest); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	UserCommand, err := createUserRequest.ToCreateUserCommand()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	commandResult, err := sc.service.CreateUser(UserCommand)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	response := mapper.ToUserResponse(commandResult.Result)

	return c.JSON(http.StatusCreated, response)
}

func (sc *UserController) GetAllUsersController(c echo.Context) error {
	offsetRaw := c.QueryParam("offset")
	countRaw := c.QueryParam("count")
	count, err := strconv.Atoi(countRaw)
	if err != nil {
		count=25
	}
	offset, err := strconv.Atoi(offsetRaw)
	if err != nil {
		offset=0
	}
	Users, err := sc.service.FindAllUsers(offset, count)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	response := mapper.ToUserListResponse(Users.Result)

	return c.JSON(http.StatusOK, response)
}

func (sc *UserController) GetUserByIdController(c echo.Context) error {

	idRaw := c.Request().URL.Path[len("/api/v1/Users/"):]

	id, err := uuid.Parse(idRaw)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	User, err := sc.service.FindUserById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	if User == nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "User not found",
		})
	}

	response := mapper.ToUserResponse(User.Result)

	return c.JSON(http.StatusOK, response)
}

func (sc *UserController) PutUserController(c echo.Context) error {
	var updateUserRequest request.UpdateUserRequest

	if err := c.Bind(&updateUserRequest); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	updateUserCommand, err := updateUserRequest.ToUpdateUserCommand()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	commandResult, err := sc.service.UpdateUser(updateUserCommand)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	response := mapper.ToUserResponse(commandResult.Result)

	return c.JSON(http.StatusOK, response)
}

func (sc *UserController) DeleteUserController(c echo.Context) error {
	idRaw := c.Request().URL.Path[len("/api/v1/users/"):]

	id, err := uuid.Parse(idRaw)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	err = sc.service.DeleteUser(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.NoContent(http.StatusNoContent)
}