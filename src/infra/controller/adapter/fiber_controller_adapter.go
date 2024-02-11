package controller_adapter

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rogeriofbrito/rinha-2024-crebito-go/src/infra/controller"
	controller_model "github.com/rogeriofbrito/rinha-2024-crebito-go/src/infra/controller/model"
	"github.com/sarulabs/di"
	log "github.com/sirupsen/logrus"
)

type FiberControllerAdapter struct {
	Tc  controller.TransactionController
	App *fiber.App
}

func (fc FiberControllerAdapter) createTransaction(dic di.Container, c *fiber.Ctx) error {
	tm := controller_model.TransactionModel{}
	if err := c.BodyParser(&tm); err != nil {
		return err
	}

	scdic, err := dic.SubContainer()
	if err != nil {
		return err
	}
	defer scdic.Delete()

	b, err := fc.Tc.CreateTransaction(scdic, tm)
	if err != nil {
		return err
	}

	c.JSON(b)

	return nil
}

func (fc FiberControllerAdapter) Start(dic di.Container) error {
	fc.App.Post("/clientes/:id/transacoes", func(c *fiber.Ctx) error {
		log.WithFields(log.Fields{
			"originalUrl": c.OriginalURL(),
			"method":      c.Method(),
		}).Info("Handling request...")
		return fc.createTransaction(dic, c)
	})

	return fc.App.Listen("127.0.0.1:3000")
}
