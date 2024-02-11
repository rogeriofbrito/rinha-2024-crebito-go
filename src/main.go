package main

import (
	"github.com/rogeriofbrito/rinha-2024-crebito-go/src/di"
	controller_adapter "github.com/rogeriofbrito/rinha-2024-crebito-go/src/infra/controller/adapter"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetLevel(log.DebugLevel)
	dic := di.GetDiContainer()
	ca := dic.Get("fiber-controller-adapter").(controller_adapter.FiberControllerAdapter)
	if err := ca.Start(dic); err != nil {
		panic(err)
	}
}
