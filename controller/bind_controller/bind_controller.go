package bind_controller

import (
	"dns_api/bind"
)

type BindController struct {
	bindControllerBehavior BindControllerBehavior
}

type BindControllerBehavior interface {
	searchRecords() ([]bind.Record, error)
	updateRecord(int, bind.Record) error
}

type BindControllerImp struct {
	bind bind.Bind
}

func (bindControllerImp *BindControllerImp) searchRecords() ([]bind.Record, error) {
	err := bindControllerImp.bind.GetZoneRecords()

	return bindControllerImp.bind.ArrayRecords, err
}

func (bindControllerImp *BindControllerImp) updateRecord(position int, record bind.Record) error {
	err := bindControllerImp.bind.UpdateZoneRecord(position, record)

	return err
}

func InitializeBindController() BindController {
	var bindController BindController
	var bindControllerImp = new(BindControllerImp)

	bind := bind.InitializeBind()
	bindControllerImp.bind = bind

	bindController.bindControllerBehavior = bindControllerImp

	return bindController
}
