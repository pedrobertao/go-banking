package router

import (
	"go-banking/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App, accHandler *handlers.AccountHandler, txHandler *handlers.TransactionHandler) {
	api := app.Group("/api/v1")

	api.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "ok",
			"message": "Banking API is running",
		})
	})

	accounts := api.Group("/accounts")
	accounts.Post("/", accHandler.CreateAccount)
	accounts.Get("/", accHandler.GetAllAccounts)
	accounts.Get("/:id", accHandler.GetAccount)

	transactions := api.Group("/transactions")
	transactions.Post("/", txHandler.CreateTransaction)
	transactions.Get("/", txHandler.GetTransactions)
}
