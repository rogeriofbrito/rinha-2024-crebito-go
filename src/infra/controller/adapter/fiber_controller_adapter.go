package controller_adapter

import (
	"fmt"
	"os"
	"strconv"

	"github.com/go-playground/validator"
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
	req := controller_model.CreateTransactionRequestModel{}
	if err := c.BodyParser(&req); err != nil {
		c.Status(fiber.StatusUnprocessableEntity)
		return nil
	}

	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		c.Status(fiber.StatusUnprocessableEntity)
		return nil
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

	res, err := fc.Tc.CreateTransaction(scdic, clientId, req)
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

	c.JSON(res)

	return nil
}

func (fc FiberControllerAdapter) getStatement(dic di.Container, c *fiber.Ctx) error {
	clientId, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return err
	}

	scdic, err := dic.SubContainer()
	if err != nil {
		return err
	}
	defer scdic.Delete()

	st, err := fc.Tc.GetStatement(scdic, clientId)
	if err != nil {
		switch err.Error() {
		case "client not found":
			c.Status(fiber.StatusNotFound)
			return nil
		}

		c.Status(fiber.StatusInternalServerError)
		return nil
	}

	c.JSON(st)

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

	fc.App.Get("/clientes/:id/extrato", func(c *fiber.Ctx) error {
		log.WithFields(log.Fields{
			"originalUrl": c.OriginalURL(),
			"method":      c.Method(),
		}).Info("Handling request...")
		return fc.getStatement(dic, c)
	})

	return fc.App.Listen(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
