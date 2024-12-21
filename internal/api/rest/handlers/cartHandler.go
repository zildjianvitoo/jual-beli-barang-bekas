package handlers

import (
	"jual-beli-barang-bekas/internal/api/rest"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type CartHandler struct {
	// Cart sv
}

func SetupCartRoutes(rh *rest.RestHandler) {
	app := rh.App

	handler := CartHandler{}

	app.Get("/cart", handler.GetCart)
	app.Post("/cart", handler.AddItemToCart)
}

func (h *CartHandler) GetCart(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Success get cart",
	})
}

func (h *CartHandler) AddItemToCart(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Success add item to cart",
	})
}
