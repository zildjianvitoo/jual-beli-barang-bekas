package handlers

import (
	"jual-beli-barang-bekas/internal/api/rest"
	"jual-beli-barang-bekas/internal/dto"
	"jual-beli-barang-bekas/internal/repository"
	"jual-beli-barang-bekas/internal/service"

	"github.com/gofiber/fiber/v2"
)

type CartHandler struct {
	service service.CartService
}

func SetupCartRoutes(rh *rest.RestHandler) {
	app := rh.App

	service := service.CartService{
		Repo:   repository.NewCartRepository(rh.DB),
		Auth:   rh.Auth,
		Config: rh.Config,
	}

	handler := CartHandler{
		service: service,
	}

	app.Get("/cart", rh.Auth.Authorize, handler.GetCart)
	app.Post("/cart", rh.Auth.Authorize, handler.AddItemToCart)
}

func (h *CartHandler) GetCart(ctx *fiber.Ctx) error {
	user := h.service.Auth.GetCurrentUser(ctx)
	cart, err := h.service.GetCart(user.ID)

	if err != nil {
		return rest.ErrorMessage(ctx, 404, err)
	}

	return rest.SuccessResponse(ctx, "Success get cart", cart)
}

func (h *CartHandler) AddItemToCart(ctx *fiber.Ctx) error {
	req := dto.CreateCartRequest{}

	err := ctx.BodyParser(&req)
	if err != nil {
		rest.BadRequestError(ctx, "Please provide valid input")
	}

	user := h.service.Auth.GetCurrentUser(ctx)

	cartItems, err := h.service.CreateCart(req, user)
	if err != nil {
		return rest.InternalError(ctx, err)
	}

	return rest.SuccessResponse(ctx, "Cart created successfully", cartItems)
}
