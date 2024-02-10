package factory

import (
	"github.com/gofiber/fiber/v2"
)

func newFiberApp() *fiber.App {
	return fiber.New()
}
