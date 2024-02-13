package controller_adapter

import (
	"strconv"

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

	clientId, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return err
	}

	scdic, err := dic.SubContainer()
	if err != nil {
		return err
	}
	defer scdic.Delete()

	tm, err = fc.Tc.CreateTransaction(scdic, clientId, tm)
	if err != nil {
		switch err.Error() {
		case "client not found":
			c.Status(fiber.StatusNotFound)
			return nil
		case "invalid transaction":
			c.Status(fiber.StatusUnprocessableEntity)
			return nil
		}

		c.Status(fiber.StatusInternalServerError)
		return nil
	}

	c.JSON(struct {
		Limit   int64 `json:"limite"`
		Balance int64 `json:"saldo"`
	}{
		Limit:   99,
		Balance: 99,
	})

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
