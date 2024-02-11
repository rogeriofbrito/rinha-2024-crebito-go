package main

import (
	"github.com/rogeriofbrito/rinha-2024-crebito-go/src/di"
	controller_adapter "github.com/rogeriofbrito/rinha-2024-crebito-go/src/infra/controller/adapter"
)

func main() {
	dic := di.GetDiContainer()
	ca := dic.Get("fiber-controller-adapter").(controller_adapter.FiberControllerAdapter)
	if err := ca.Start(dic); err != nil {
		panic(err)
	}
}
