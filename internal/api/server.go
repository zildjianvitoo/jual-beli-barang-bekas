package api

import (
	"jual-beli-barang-bekas/config"
	"jual-beli-barang-bekas/internal/api/rest"
	"jual-beli-barang-bekas/internal/api/rest/handlers"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func StartServer(config config.AppConfig) {
	app := fiber.New()

	app.Get("/health", HealthCheck)

	restHandler := &rest.RestHandler{
		App: app,
	}

	SetupRoutes(restHandler)

	app.Listen(config.ServerPort)
}

func HealthCheck(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Health check success",
	})
}

func SetupRoutes(rh *rest.RestHandler) {
	handlers.SetupUserRoutes(rh)
	handlers.SetupCartRoutes(rh)
	handlers.SetupOrderRoutes(rh)
	//	Transaction
	//	Catalog
}
