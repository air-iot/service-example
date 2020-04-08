package main

import (
	"net/http"

	"github.com/air-iot/service"
	"github.com/labstack/echo/v4"
)

func main() {
	app := service.NewApp()
	app.GetServer().GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK,"hello world")
	})
	app.Run()
}
