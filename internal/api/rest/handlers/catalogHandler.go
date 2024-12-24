package handlers

import (
	"jual-beli-barang-bekas/internal/api/rest"
	"jual-beli-barang-bekas/internal/dto"
	"jual-beli-barang-bekas/internal/repository"
	"jual-beli-barang-bekas/internal/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type CatalogHandler struct {
	service service.CatalogService
}

func SetupCatalogRoutes(rh *rest.RestHandler) {

	app := rh.App

	service := service.CatalogService{
		Repo:   repository.NewCatalogRepository(rh.DB),
		Auth:   rh.Auth,
		Config: rh.Config,
	}
	handler := CatalogHandler{
		service: service,
	}

	// Public
	// Listing products and categories
	app.Get("/products", handler.GetProducts)
	app.Get("/products/:id", handler.GetProductById)
	app.Get("/categories", handler.GetCategories)
	app.Get("/categories/:id", handler.GetCategoryById)

	// Private
	// Manage products and categories
	selRoutes := app.Group("/seller", rh.Auth.AuthorizeSeller)
	// Categories
	selRoutes.Post("/categories", handler.CreateCategory)
	selRoutes.Patch("/categories/:id", handler.EditCategory)
	selRoutes.Delete("/categories/:id", handler.DeleteCategory)

	// Products
	selRoutes.Post("/products", handler.CreateProduct)
	selRoutes.Get("/products", handler.GetProducts)
	selRoutes.Get("/products/:id", handler.GetProductById)
	selRoutes.Put("/products/:id", handler.EditProduct)

	selRoutes.Delete("/products/:id", handler.DeleteProduct)
}

func (h CatalogHandler) GetCategories(ctx *fiber.Ctx) error {

	cats, err := h.service.GetCategories()
	if err != nil {
		return rest.ErrorMessage(ctx, 404, err)
	}

	return rest.SuccessResponse(ctx, "Success get categories", cats)
}
func (h CatalogHandler) GetCategoryById(ctx *fiber.Ctx) error {

	id, _ := strconv.Atoi(ctx.Params("id"))

	cat, err := h.service.GetCategory(id)
	if err != nil {
		return rest.ErrorMessage(ctx, 404, err)
	}

	return rest.SuccessResponse(ctx, "Success get category", cat)
}

func (h CatalogHandler) CreateCategory(ctx *fiber.Ctx) error {

	req := dto.CreateCategoryRequest{}

	err := ctx.BodyParser(&req)
	if err != nil {
		return rest.BadRequestError(ctx, "Create category request is not valid")
	}

	err = h.service.CreateCategory(req)
	if err != nil {
		return rest.InternalError(ctx, err)
	}

	return rest.SuccessResponse(ctx, "Category created successfully", nil)
}

func (h CatalogHandler) EditCategory(ctx *fiber.Ctx) error {

	id, _ := strconv.Atoi(ctx.Params("id"))

	req := dto.CreateCategoryRequest{}

	err := ctx.BodyParser(&req)
	if err != nil {
		return rest.BadRequestError(ctx, "Update category request is not valid")
	}

	updatedCat, err := h.service.EditCategory(id, req)
	if err != nil {
		return rest.InternalError(ctx, err)
	}

	return rest.SuccessResponse(ctx, "Success edit category", updatedCat)
}

func (h CatalogHandler) DeleteCategory(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))

	err := h.service.DeleteCategory(id)
	if err != nil {
		return rest.InternalError(ctx, err)
	}

	return rest.SuccessResponse(ctx, "Category delete successfully", nil)
}

func (h CatalogHandler) CreateProduct(ctx *fiber.Ctx) error {

	req := dto.CreateProductRequest{}
	err := ctx.BodyParser(&req)
	if err != nil {
		return rest.BadRequestError(ctx, "Create product request is not valid")
	}

	user := h.service.Auth.GetCurrentUser(ctx)
	err = h.service.CreateProduct(req, user)
	if err != nil {
		return rest.InternalError(ctx, err)
	}

	return rest.SuccessResponse(ctx, "Product created successfully", nil)
}

func (h CatalogHandler) GetProducts(ctx *fiber.Ctx) error {

	products, err := h.service.GetProducts()
	if err != nil {
		return rest.ErrorMessage(ctx, 404, err)
	}

	return rest.SuccessResponse(ctx, "Success get products", products)
}

func (h CatalogHandler) GetProductById(ctx *fiber.Ctx) error {

	id, _ := strconv.Atoi(ctx.Params("id"))

	product, err := h.service.GetProductById(id)
	if err != nil {
		return rest.BadRequestError(ctx, "Product not found")
	}

	return rest.SuccessResponse(ctx, "Success get product", product)
}

func (h CatalogHandler) EditProduct(ctx *fiber.Ctx) error {

	id, _ := strconv.Atoi(ctx.Params("id"))

	req := dto.CreateProductRequest{}

	err := ctx.BodyParser(&req)
	if err != nil {
		return rest.BadRequestError(ctx, "Edit product request is not valid")
	}

	user := h.service.Auth.GetCurrentUser(ctx)

	product, err := h.service.EditProduct(id, req, user)
	if err != nil {
		return rest.InternalError(ctx, err)
	}

	return rest.SuccessResponse(ctx, "Success edit product", product)
}

func (h CatalogHandler) DeleteProduct(ctx *fiber.Ctx) error {

	id, _ := strconv.Atoi(ctx.Params("id"))

	user := h.service.Auth.GetCurrentUser(ctx)
	err := h.service.DeleteProduct(id, user)

	return rest.SuccessResponse(ctx, "Success delete product ", err)
}
