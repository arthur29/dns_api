package bind_controller

import (
	"dns_api/bind"
	"fmt"

	"github.com/labstack/echo/v4"
)

func ListIndex(c echo.Context) error {
	gozoneArray, err := bind.ReadZoneFile()

	if err != nil {
		fmt.Errorf("Error on read dns zone %v", err)
		return c.String(500, "Error on read dns zone file")
	}
	return c.JSONPretty(200, gozoneArray, "  ")
}
