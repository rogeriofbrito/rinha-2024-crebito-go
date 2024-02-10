package main

import "github.com/rogeriofbrito/rinha-2024-crebito-go/src/factory"

func main() {
	ca := factory.NewControllerAdapter()
	if err := ca.Start(); err != nil {
		panic(err)
	}
}
