package handlers

import (
	"jual-beli-barang-bekas/internal/api/rest"
	"jual-beli-barang-bekas/internal/dto"
	"jual-beli-barang-bekas/internal/repository"
	"jual-beli-barang-bekas/internal/service"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	service service.UserService
}

func SetupUserRoutes(rh *rest.RestHandler) {

	app := rh.App

	repo := repository.NewUserRepository(rh.DB)

	service := service.UserService{
		Repo:   repo,
		Auth:   rh.Auth,
		Config: rh.Config,
	}

	handler := UserHandler{
		service: service,
	}

	// Public endpoint
	publicRoutes := app.Group("/users")

	publicRoutes.Post("/register", handler.Register)
	publicRoutes.Post("/login", handler.Login)

	// Private endpoint
	privateRoutes := app.Group("/users", rh.Auth.Authorize)

	privateRoutes.Get("/verify", handler.GetVerificationCode)
	privateRoutes.Post("/verify", handler.DoVerify)
	privateRoutes.Get("/profile", handler.GetProfile)
	privateRoutes.Post("/profile", handler.CreateProfile)

	privateRoutes.Post("/become-seller", handler.BecomeSeller)

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
		"data": fiber.Map{
			"token": token,
		},
	})
}

func (h *UserHandler) Login(ctx *fiber.Ctx) error {

	loginInput := dto.UserLogin{}

	err := ctx.BodyParser(&loginInput)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Please provide valid inputs",
		})
	}

	token, err := h.service.Login(loginInput)
	if err != nil {
		return ctx.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"message": "Wrong email/password",
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Login success",
		"data": fiber.Map{
			"token": token,
		},
	})
}

func (h *UserHandler) GetVerificationCode(ctx *fiber.Ctx) error {
	user := h.service.Auth.GetCurrentUser(ctx)
	log.Println(user)

	code, err := h.service.GetVerificationCode(user)
	log.Println(err)

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "Unable to generate verification code",
		})
	}

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Success getting verification code",
		"data": fiber.Map{
			"code": code,
		},
	})
}

func (h *UserHandler) DoVerify(ctx *fiber.Ctx) error {
	user := h.service.Auth.GetCurrentUser(ctx)

	var req dto.VerificationCodeInput

	err := ctx.BodyParser(&req)

	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "Please provide a valid input",
		})
	}

	err = h.service.DoVerify(user.ID, req.Code)

	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Verified success",
	})
}

func (h *UserHandler) GetProfile(ctx *fiber.Ctx) error {
	user := h.service.Auth.GetCurrentUser(ctx)

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Success get profile",
		"data": fiber.Map{
			"user": user,
		},
	})
}

func (h *UserHandler) CreateProfile(ctx *fiber.Ctx) error {
	user := h.service.Auth.GetCurrentUser(ctx)
	req := dto.ProfileInput{}
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "Please provide a valid input",
		})
	}
	log.Printf("User %v", user)

	err := h.service.CreateProfile(user.ID, req)

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "Unable to create profile",
		})
	}

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Profile created successfully",
	})
}

func (h *UserHandler) BecomeSeller(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Success become seller",
	})
}
