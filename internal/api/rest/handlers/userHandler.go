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
		return rest.BadRequestError(ctx, "Please provide valid input")
	}

	token, err := h.service.Register(user)
	if err != nil {
		return rest.InternalError(ctx, err)
	}

	return rest.SuccessResponse(ctx, "Register success", fiber.Map{
		"token": token,
	})

}

func (h *UserHandler) Login(ctx *fiber.Ctx) error {

	loginInput := dto.UserLogin{}

	err := ctx.BodyParser(&loginInput)
	if err != nil {
		return rest.BadRequestError(ctx, "Please provide valid input")
	}

	token, err := h.service.Login(loginInput)
	if err != nil {
		return rest.ErrorMessage(ctx, http.StatusUnauthorized, err)
	}

	return rest.SuccessResponse(ctx, "Login success", fiber.Map{
		"token": token,
	})
}

func (h *UserHandler) GetVerificationCode(ctx *fiber.Ctx) error {
	user := h.service.Auth.GetCurrentUser(ctx)
	log.Println(user)

	code, err := h.service.GetVerificationCode(user)
	log.Println(err)

	if err != nil {
		return rest.InternalError(ctx, err)
	}

	return rest.SuccessResponse(ctx, "Success getting verification code", fiber.Map{
		"code": code,
	})
}

func (h *UserHandler) DoVerify(ctx *fiber.Ctx) error {
	user := h.service.Auth.GetCurrentUser(ctx)

	var req dto.VerificationCodeInput

	err := ctx.BodyParser(&req)

	if err != nil {
		return rest.BadRequestError(ctx, "Please provide valid input")
	}

	err = h.service.DoVerify(user.ID, req.Code)

	if err != nil {
		return rest.BadRequestError(ctx, "Please provide valid input")
	}

	return rest.SuccessResponse(ctx, "Verification successfully", fiber.Map{})
}

func (h *UserHandler) GetProfile(ctx *fiber.Ctx) error {
	user := h.service.Auth.GetCurrentUser(ctx)

	return rest.SuccessResponse(ctx, "Success get profile", fiber.Map{
		"user": user,
	})
}

func (h *UserHandler) CreateProfile(ctx *fiber.Ctx) error {
	user := h.service.Auth.GetCurrentUser(ctx)

	req := dto.ProfileInput{}

	if err := ctx.BodyParser(&req); err != nil {
		return rest.BadRequestError(ctx, "Please provide valid input")
	}
	log.Printf("User %v", user)

	err := h.service.CreateProfile(user.ID, req)

	if err != nil {
		return rest.InternalError(ctx, err)
	}

	return rest.SuccessResponse(ctx, "Profile created successfully", fiber.Map{
		"user": user,
	})
}

func (h *UserHandler) BecomeSeller(ctx *fiber.Ctx) error {

	user := h.service.Auth.GetCurrentUser(ctx)

	req := dto.BecomeSellerInput{}
	err := ctx.BodyParser(&req)
	if err != nil {
		return rest.BadRequestError(ctx, "Please provide valid input")
	}

	token, err := h.service.BecomeSeller(user.ID, req)
	if err != nil {
		return rest.ErrorMessage(ctx, http.StatusUnauthorized, err)
	}

	return rest.SuccessResponse(ctx, "Success become seller", fiber.Map{
		"token": token,
	})

}
