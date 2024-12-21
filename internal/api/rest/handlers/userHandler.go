package handlers

import (
	"jual-beli-barang-bekas/internal/api/rest"
	"jual-beli-barang-bekas/internal/dto"
	"jual-beli-barang-bekas/internal/service"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	service service.UserService
}

func SetupUserRoutes(rh *rest.RestHandler) {

	app := rh.App

	service := service.UserService{}

	handler := UserHandler{
		service: service,
	}

	// Public endpoint
	app.Post("/register", handler.Register)
	app.Post("/login", handler.Login)

	// Private endpoint
	app.Get("/verify", handler.GetVerificationCode)
	app.Post("/verify", handler.DoVerify)
	app.Get("/profile", handler.GetProfile)
	app.Post("/profile", handler.CreateProfile)

	app.Post("/become-seller", handler.BecomeSeller)

}

func (h *UserHandler) Register(ctx *fiber.Ctx) error {
	user := dto.UserRegister{}

	err := ctx.BodyParser(&user)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Please provide valid inputs",
		})
	}

	token, err := h.service.Register(user)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error on signup",
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Register success",
		"data":    token,
	})
}

func (h *UserHandler) Login(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Login success",
	})
}

func (h *UserHandler) GetVerificationCode(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Verification code sent",
	})
}

func (h *UserHandler) DoVerify(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Verify success",
	})
}

func (h *UserHandler) GetProfile(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Success get profile",
	})
}

func (h *UserHandler) CreateProfile(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Success create profile",
	})
}

func (h *UserHandler) BecomeSeller(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Success become seller",
	})
}
