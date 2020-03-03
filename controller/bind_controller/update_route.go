package bind_controller

import (
	"dns_api/bind"
	"fmt"
	"strconv"

	"github.com/labstack/echo/v4"
)

type BindControllerUpdateBehavior interface {
	updateRecord(int, bind.Record) error
}

func (bindControllerImp *BindControllerImp) updateRecord(position int, record bind.Record) error {
	return nil
	//err := bindControllerImp.bind.UpdateZoneRecord(position, record)

	//return err
}

func (bindController *BindController) Update(c echo.Context) error {
	position, err := strconv.Atoi(c.Param("pos"))

	if err != nil {
		fmt.Errorf("Error on parse param position to integer %v", err)
		return c.String(400, "Error on parse url param to integer")
	}

	record := bind.Record{}
	c.Bind(record)

	fmt.Printf("%s\n\n\n", record)

	err = bindController.bindControllerUpdateBehavior.updateRecord(position, record)

	if err != nil {
		switch err.(type) {
		//case errors.Validation:
		//	fmt.Errorf("Error on update record invalid content %v", err)
		//	return c.String(422, "Error on update record")
		default:
			fmt.Errorf("Error on update error unknown error %v", err)
			return c.String(500, "Error on update error unknown error")
		}
	}

	return c.String(200, "Record updated")
}
