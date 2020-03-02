package main

import (
	"dns_api/controller/bind_controller"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	bindController := bind_controller.InitializeBindController()
	e.GET("/", bindController.Index)
	e.Logger.Fatal(e.Start(":9000"))
}
