package main

import (
	"log"

	"go-banking/config"
	"go-banking/internal/database"
	"go-banking/internal/handlers"
	"go-banking/internal/middleware"
	"go-banking/internal/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

	db, err := database.New(cfg)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}
			return c.Status(code).JSON(fiber.Map{
				"error": err.Error(),
			})
		},
	})

	app.Use(cors.New())
	app.Use(middleware.Logger())

	accountHandler := handlers.NewAccountHandler(db)
	transactionHandler := handlers.NewTransactionHandler(db)

	router.Setup(app, accountHandler, transactionHandler)

	log.Printf("Server starting on port %s", cfg.Port)
	log.Fatal(app.Listen(":" + cfg.Port))
}
