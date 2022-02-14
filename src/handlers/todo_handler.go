package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tatsuya0429/edgedb_golang_docker_sample/src/repositories"
)

func AddTodo(c echo.Context) error {
	userId := c.FormValue("user_id")
	title := c.FormValue("title")
	description := c.FormValue("description")
	deadline := c.FormValue("deadline")
	repo := repositories.NewTodoRepository()
	err := repo.CreateTodo(c.Request().Context(), title, description, deadline, userId)
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, "Bad Request")
	}
	return c.JSON(http.StatusOK, "Success!!")
}

func GetAllTodos(c echo.Context) error {
	repo := repositories.NewTodoRepository()
	todos, err := repo.GetAll(c.Request().Context())
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, "Bad Request")
	}
	return c.JSON(http.StatusOK, todos)
}