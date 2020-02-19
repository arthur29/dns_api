package bind_controller

import (
	"github.com/labstack/echo/v4"
)

func ListIndex(c echo.Context) error {
	return c.String(200, "Hello World")
}
