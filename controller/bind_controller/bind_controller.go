package bind_controller

import (
	"dns_api/bind"
)

type BindController struct {
	bindControllerIndexBehavior  BindControllerIndexBehavior
	bindControllerUpdateBehavior BindControllerUpdateBehavior
}

type BindControllerImp struct {
	bind bind.Bind
}

func InitializeBindController() BindController {
	var bindController BindController
	var bindControllerImp = new(BindControllerImp)

	bind := bind.InitializeBind()
	bindControllerImp.bind = bind

	bindController.bindControllerIndexBehavior = bindControllerImp
	bindController.bindControllerUpdateBehavior = bindControllerImp

	return bindController
}
