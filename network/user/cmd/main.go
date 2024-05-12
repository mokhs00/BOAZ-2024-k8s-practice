package main

import (
	"github.com/labstack/echo/v4"
	"user/handler"
)

func main() {
	e := echo.New()

	e.GET("/users/:id", handler.GetUserHandler)

	e.Logger.Fatal(e.Start(":8080"))
}
