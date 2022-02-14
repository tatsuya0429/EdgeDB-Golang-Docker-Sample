package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/tatsuya0429/edgedb_golang_docker_sample/src/handlers"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Echo World!!")
	})
	e.GET("/users", handlers.GetUsers)
	e.POST("/signup", handlers.SignUp)
	e.POST("/signin", handlers.SignIn)

	e.GET("/todos", handlers.GetAllTodos)
	e.POST("/todos", handlers.AddTodo)

	e.Logger.Fatal(e.Start(":5000"))
}
