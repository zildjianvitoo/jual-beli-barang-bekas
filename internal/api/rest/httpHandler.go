package rest

import (
	"jual-beli-barang-bekas/internal/helper"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type RestHandler struct {
	App  *fiber.App
	DB   *gorm.DB
	Auth helper.Auth
}
