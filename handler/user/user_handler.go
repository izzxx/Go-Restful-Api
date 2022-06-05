package user

import (
	"net/http"

	"github.com/izzxx/Go-Restful-Api/service/user"
	"github.com/izzxx/Go-Restful-Api/utility"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	UserService user.UserService
}

func (uh *UserHandler) Register(c echo.Context) error {
	var userRegister user.UserRegister
	err := c.Bind(&userRegister)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utility.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	response, err := uh.UserService.Register(c.Request().Context(), userRegister)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utility.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}

	return c.JSON(http.StatusOK, UserResponseApi{
		StatusCode: http.StatusOK,
		Message:    "success register",
		Data:       response,
	})
}

func (uh *UserHandler) Login(c echo.Context) error {
	var userLogin user.UserLogin
	err := c.Bind(&userLogin)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utility.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	response, err := uh.UserService.Login(c.Request().Context(), userLogin)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utility.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}

	return c.JSON(http.StatusOK, UserResponseApi{
		StatusCode: http.StatusOK,
		Message:    "success login",
		Data:       response,
	})
}

func (uh *UserHandler) UpdatePassword(c echo.Context) error {
	var userUpdatePassword user.UserUpdatePassword
	err := c.Bind(&userUpdatePassword)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utility.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	if err = uh.UserService.UpdatePassword(c.Request().Context(), userUpdatePassword); err != nil {
		return c.JSON(http.StatusInternalServerError, utility.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}

	return c.JSON(http.StatusOK, UserResponseApi{
		StatusCode: http.StatusOK,
		Message:    "success update password",
	})
}
