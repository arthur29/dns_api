package bind_controller

import (
	"dns_api/bind"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/wpalmer/gozone"
)

type BindControllerDI struct {
	bindControllerBehavior BindControllerBehavior
}

type BindControllerBehavior interface {
	searchRecords() ([]gozone.Record, error)
}

type BindController struct{}

func (bindController *BindController) searchRecords() ([]gozone.Record, error) {
	return bind.ReadZoneFile()
}

func InitializeBindController() BindControllerDI {
	var bindControllerDI BindControllerDI

	bindControllerDI.bindControllerBehavior = &BindController{}

	return bindControllerDI
}

func (bindControllerDI *BindControllerDI) ListIndex(c echo.Context) error {
	gozoneArray, err := bindControllerDI.bindControllerBehavior.searchRecords()

	if err != nil {
		fmt.Errorf("Error on read dns zone %v", err)
		return c.String(500, "Error on read dns zone file")
	}
	return c.JSONPretty(200, gozoneArray, "  ")
}
