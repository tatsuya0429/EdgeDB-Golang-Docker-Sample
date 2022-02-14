package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tatsuya0429/edgedb_golang_docker_sample/src/repositories"
)

func SignIn(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")
	repo := repositories.NewUserRepository()
	user, err := repo.GetUserByUsername(c.Request().Context(), username)
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusUnauthorized, "Invalid username or password")
	}
	if user.Password != password {
		return c.JSON(http.StatusUnauthorized, "Invalid username or password")
	}
	return c.JSON(http.StatusOK, user)
}

func SignUp(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")
	repo := repositories.NewUserRepository()
	err := repo.CreateUser(c.Request().Context(), username, password)
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, "Bad Request")
	}
	return c.JSON(http.StatusOK, "Success!!")
}

func GetUsers(c echo.Context) error {
	repo := repositories.NewUserRepository()
	users, err := repo.GetAll(c.Request().Context())
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, "Bad Request")
	}
	return c.JSON(http.StatusOK, users)
}
