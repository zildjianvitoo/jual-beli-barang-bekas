package api

import (
	"fmt"
	"jual-beli-barang-bekas/config"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func StartServer(config config.AppConfig) {
	app := fiber.New()

	fmt.Println("Server berjalan")

	app.Get("/health", HealthCheck)

	app.Listen(config.ServerPort)
}

func HealthCheck(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Health Check Success",
	})
}
