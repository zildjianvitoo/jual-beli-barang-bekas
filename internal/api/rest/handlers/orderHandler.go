package handlers

import (
	"jual-beli-barang-bekas/internal/api/rest"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type OrderHandler struct {
	// sv
}

func SetupOrderRoutes(rh *rest.RestHandler) {

	app := rh.App

	handler := OrderHandler{}

	app.Get("/order", handler.GetOrders)
	app.Get("/order/:id", handler.GetOrderById)
}

func (h *OrderHandler) GetOrders(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Success Get all orders",
	})
}

func (h *OrderHandler) GetOrderById(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Success get order by id",
	})
}
