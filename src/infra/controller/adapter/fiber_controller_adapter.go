package controller_adapter

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rogeriofbrito/rinha-2024-crebito-go/src/infra/controller"
	controller_model "github.com/rogeriofbrito/rinha-2024-crebito-go/src/infra/controller/model"
)

type FiberControllerAdapter struct {
	Tc  controller.TransactionController
	App *fiber.App
}

func (fc FiberControllerAdapter) createTransaction(c *fiber.Ctx) error {
	tm := controller_model.TransactionModel{}
	if err := c.BodyParser(&tm); err != nil {
		return err
	}

	b, err := fc.Tc.CreateTransaction(tm)
	if err != nil {
		return err
	}

	c.JSON(b)

	return nil
}

func (fc FiberControllerAdapter) Start() error {
	fc.App.Post("/clientes/:id/transacoes", func(c *fiber.Ctx) error {
		return fc.createTransaction(c)
	})

	return fc.App.Listen("127.0.0.1:3000")
}
