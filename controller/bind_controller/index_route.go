package bind_controller

import (
	"dns_api/bind"
	"fmt"

	"github.com/labstack/echo/v4"
)

type BindControllerIndexBehavior interface {
	searchRecords() ([]bind.Record, error)
}

func (bindControllerImp *BindControllerImp) searchRecords() ([]bind.Record, error) {
	err := bindControllerImp.bind.GetZoneRecords()

	return bindControllerImp.bind.ArrayRecords, err
}

func (bindController *BindController) Index(c echo.Context) error {
	array, err := bindController.bindControllerIndexBehavior.searchRecords()

	if err != nil {
		fmt.Errorf("Error on read dns zone %v", err)
		return c.String(500, "Error on read dns zone file")
	}
	return c.JSONPretty(200, array, "  ")
}
