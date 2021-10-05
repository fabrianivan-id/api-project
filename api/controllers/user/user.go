package user

import (
	"net/http"
	"strconv"

	"project-api/api/common"
	"project-api/models"

	echo "github.com/labstack/echo/v4"
)

type Controller struct {
	userModel models.UserModel
}

func NewController(userModel models.UserModel) *Controller {
	return &Controller{
		userModel,
	}
}

func (controller *Controller) GetAllUserController(c echo.Context) error {
	user, err := controller.userModel.GetAll()
	if err != nil {
		return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
	}

	return c.JSON(http.StatusOK, user)
}

func (controller *Controller) GetUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
	}

	user, err := controller.userModel.Get(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
	}

	response := GetUserResponse{
		Name:  user.Name,
		Email: user.Email,
	}

	return c.JSON(http.StatusOK, response)
}

func (controller *Controller) PostUserController(c echo.Context) error {
	// bind request value
	var userRequest PostUserRequest

	if err := c.Bind(&userRequest); err != nil {
		return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
	}

	user := models.User{
		Name:     userRequest.Name,
		Email:    userRequest.Email,
		Password: userRequest.Password,
	}

	_, err := controller.userModel.Insert(user)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse())
	}

	return c.JSON(http.StatusOK, common.NewSuccessOperationResponse())
}

func (controller *Controller) EditUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
	}

	// bind request value
	var userRequest EditUserRequest
	if err := c.Bind(&userRequest); err != nil {
		return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
	}

	user := models.User{
		Name:     userRequest.Name,
		Email:    userRequest.Email,
		Password: userRequest.Password,
	}

	if _, err := controller.userModel.Edit(user, id); err != nil {
		return c.JSON(http.StatusNotFound, common.NewBadRequestResponse())
	}

	return c.JSON(http.StatusOK, common.NewSuccessOperationResponse())
}

func (controller *Controller) DeleteUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
	}

	if _, err := controller.userModel.Delete(id); err != nil {
		return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
	}

	return c.JSON(http.StatusOK, common.NewSuccessOperationResponse())
}

func (controller *Controller) LoginUserController(c echo.Context) error {
	var userRequest LoginUserRequest

	if err := c.Bind(&userRequest); err != nil {
		return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
	}

	user, err := controller.userModel.Login(userRequest.Email, userRequest.Password)

	if err != nil {
		return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"token": user.Token,
	})
}
