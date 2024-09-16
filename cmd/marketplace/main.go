package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/whatthehanan/go-ddd/internal/application/services"
	"github.com/whatthehanan/go-ddd/internal/infrastructure/db/postgres"
	"github.com/whatthehanan/go-ddd/internal/interface/api/rest"
)

func loadEnvironmentVariables() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	loadEnvironmentVariables()

	// initialize database connection and repositories
	gormDB := postgres.NewConnection(postgres.ConnectionConfig{ShouldMigrate: true})
	productRepo := postgres.NewGormProductRepository(gormDB)
	sellerRepo := postgres.NewGormSellerRepository(gormDB)

	// initialize services
	productService := services.NewProductService(productRepo, sellerRepo)
	sellerService := services.NewSellerService(sellerRepo)

	// initialize http server and controllers
	e := echo.New()
	rest.NewProductController(e, productService)
	rest.NewSellerController(e, sellerService)
	port := ":8080"
	if err := e.Start("127.0.0.1" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
