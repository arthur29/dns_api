package main

import (
	"dns_api/controller/bind_controller"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/list", bind_controller.ListIndex)
	e.Logger.Fatal(e.Start(":9000"))
}
