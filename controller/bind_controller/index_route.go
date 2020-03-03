package bind_controller

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

func (bindController *BindController) Index(c echo.Context) error {
	array, err := bindController.bindControllerBehavior.searchRecords()

	if err != nil {
		fmt.Errorf("Error on read dns zone %v", err)
		return c.String(500, "Error on read dns zone file")
	}
	return c.JSONPretty(200, array, "  ")
}
