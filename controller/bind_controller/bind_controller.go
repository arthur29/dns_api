package bind_controller

import (
	"dns_api/bind"
	"fmt"

	"github.com/labstack/echo/v4"
)

type BindController struct {
	bindControllerBehavior BindControllerBehavior
}

type BindControllerBehavior interface {
	searchRecords() ([]bind.Record, error)
}

type BindControllerImp struct {
	bind bind.Bind
}

func (bindControllerImp *BindControllerImp) searchRecords() ([]bind.Record, error) {
	err := bindControllerImp.bind.GetZoneRecords()

	return bindControllerImp.bind.ArrayRecords, err
}

func InitializeBindController() BindController {
	var bindController BindController
	var bindControllerImp = new(BindControllerImp)

	bindControllerImp.bind = bind.InitializeBind()
	bindController.bindControllerBehavior = bindControllerImp

	return bindController
}

func (bindController *BindController) Index(c echo.Context) error {
	array, err := bindController.bindControllerBehavior.searchRecords()

	if err != nil {
		fmt.Errorf("Error on read dns zone %v", err)
		return c.String(500, "Error on read dns zone file")
	}
	return c.JSONPretty(200, array, "  ")
}
