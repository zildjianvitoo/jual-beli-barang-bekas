package handlers

import (
	"errors"
	"jual-beli-barang-bekas/internal/api/rest"
	"jual-beli-barang-bekas/internal/repository"
	"jual-beli-barang-bekas/internal/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type OrderHandler struct {
	service service.OrderService
}

func SetupOrderRoutes(rh *rest.RestHandler) {

	app := rh.App

	service := service.OrderService{
		Repo:     repository.NewOrderRepository(rh.DB),
		CartRepo: repository.NewCartRepository(rh.DB),
		Auth:     rh.Auth,
		Config:   rh.Config,
	}

	handler := OrderHandler{
		service: service,
	}

	app.Post("/orders", rh.Auth.Authorize, handler.CreateOrder)
	app.Get("/orders", rh.Auth.Authorize, handler.GetOrders)
	app.Get("/orders/:id", rh.Auth.Authorize, handler.GetOrderById)
}

func (h *OrderHandler) CreateOrder(ctx *fiber.Ctx) error {
	user := h.service.Auth.GetCurrentUser(ctx)
	order, err := h.service.CreateOrder(user)

	if err != nil {
		return rest.InternalError(ctx, errors.New("unable to create order"))
	}

	return rest.SuccessResponse(ctx, "Order created successfully", order)

}

func (h *OrderHandler) GetOrders(ctx *fiber.Ctx) error {
	user := h.service.Auth.GetCurrentUser(ctx)

	orders, err := h.service.GetOrders(user)
	if err != nil {
		return rest.InternalError(ctx, err)
	}

	return rest.SuccessResponse(ctx, "Success get orders", orders)

}

func (h *OrderHandler) GetOrderById(ctx *fiber.Ctx) error {
	orderId, _ := strconv.Atoi(ctx.Params("id"))
	user := h.service.Auth.GetCurrentUser(ctx)

	order, err := h.service.GetOrderById(uint(orderId), user.ID)
	if err != nil {
		return rest.InternalError(ctx, err)
	}

	return rest.SuccessResponse(ctx, "Success get order by id", order)

}
