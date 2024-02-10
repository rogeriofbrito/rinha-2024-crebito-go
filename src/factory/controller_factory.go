package factory

import (
	"github.com/rogeriofbrito/rinha-2024-crebito-go/src/infra/controller"
	controller_adapter "github.com/rogeriofbrito/rinha-2024-crebito-go/src/infra/controller/adapter"
)

func NewControllerAdapter() controller_adapter.IControllerAdapter {
	// for now, only Fiber controller implemented
	return newFiberControllerAdapter()
}

func newFiberControllerAdapter() controller_adapter.FiberControllerAdapter {
	return controller_adapter.FiberControllerAdapter{
		TransactionController: newTransactionController(),
		App:                   newFiberApp(),
	}
}

func newTransactionController() controller.TransactionController {
	return controller.TransactionController{}
}
