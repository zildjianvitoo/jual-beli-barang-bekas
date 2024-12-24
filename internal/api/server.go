package api

import (
	"jual-beli-barang-bekas/config"
	"jual-beli-barang-bekas/internal/api/rest"
	"jual-beli-barang-bekas/internal/api/rest/handlers"
	"jual-beli-barang-bekas/internal/domain"
	"jual-beli-barang-bekas/internal/helper"
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func StartServer(config config.AppConfig) {
	app := fiber.New()

	db, err := gorm.Open(postgres.Open(config.DatasourceName), &gorm.Config{})

	if err != nil {
		log.Fatalf("Database not connected %v\n", err)
	}

	err = db.AutoMigrate(&domain.User{}, &domain.BankAccount{})

	if err != nil {
		log.Fatalf("Error migration %v", err)
	}

	appSecret := helper.SetupAuth(config.AppSecret)

	restHandler := &rest.RestHandler{
		App:    app,
		DB:     db,
		Auth:   appSecret,
		Config: config,
	}

	SetupRoutes(restHandler)

	app.Listen(config.ServerPort)
}

func SetupRoutes(rh *rest.RestHandler) {
	handlers.SetupUserRoutes(rh)
	handlers.SetupCartRoutes(rh)
	handlers.SetupOrderRoutes(rh)
	//	Transaction
	//	Catalog
}
