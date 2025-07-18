package handlers

import (
	"database/sql"
	"time"

	"github.com/gofiber/fiber/v2"
	"go-banking/internal/database"
	"go-banking/internal/models"
)

type AccountHandler struct {
	db *database.DB
}

func NewAccountHandler(db *database.DB) *AccountHandler {
	return &AccountHandler{db: db}
}

func (h *AccountHandler) CreateAccount(c *fiber.Ctx) error {
	var req models.CreateAccountRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if req.Name == "" || req.Email == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Name and email are required",
		})
	}

	var accountID int
	err := h.db.QueryRow(
		"INSERT INTO accounts (name, email, created_at, updated_at) VALUES ($1, $2, $3, $4) RETURNING id",
		req.Name, req.Email, time.Now(), time.Now(),
	).Scan(&accountID)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create account",
		})
	}

	account := models.AccountResponse{
		ID:      accountID,
		Name:    req.Name,
		Email:   req.Email,
		Balance: 0.0,
	}

	return c.Status(fiber.StatusCreated).JSON(account)
}

func (h *AccountHandler) GetAccount(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Account ID is required",
		})
	}

	var account models.Account
	err := h.db.QueryRow(
		"SELECT id, name, email, balance, created_at, updated_at FROM accounts WHERE id = $1",
		id,
	).Scan(&account.ID, &account.Name, &account.Email, &account.Balance, &account.CreatedAt, &account.UpdatedAt)

	if err == sql.ErrNoRows {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Account not found",
		})
	}

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve account",
		})
	}

	response := models.AccountResponse{
		ID:      account.ID,
		Name:    account.Name,
		Email:   account.Email,
		Balance: account.Balance,
	}

	return c.JSON(response)
}

func (h *AccountHandler) GetAllAccounts(c *fiber.Ctx) error {
	rows, err := h.db.Query("SELECT id, name, email, balance FROM accounts ORDER BY id")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve accounts",
		})
	}
	defer rows.Close()

	var accounts []models.AccountResponse
	for rows.Next() {
		var account models.AccountResponse
		err := rows.Scan(&account.ID, &account.Name, &account.Email, &account.Balance)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to scan account",
			})
		}
		accounts = append(accounts, account)
	}

	return c.JSON(accounts)
}