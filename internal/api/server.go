package api

import (
	"jual-beli-barang-bekas/config"

	"github.com/gofiber/fiber/v2"
)

func StartServer(config config.AppConfig) {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Halo fiber")
	})

	app.Get("/kocak", func(c *fiber.Ctx) error {
		return c.JSON("aaa")
	})

	app.Listen(config.ServerPort)
}
