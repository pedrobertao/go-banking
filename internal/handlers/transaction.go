package handlers

import (
	"database/sql"
	"time"

	"github.com/gofiber/fiber/v2"
	"go-banking/internal/database"
	"go-banking/internal/models"
)

type TransactionHandler struct {
	db *database.DB
}

func NewTransactionHandler(db *database.DB) *TransactionHandler {
	return &TransactionHandler{db: db}
}

func (h *TransactionHandler) CreateTransaction(c *fiber.Ctx) error {
	var req models.CreateTransactionRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if req.Amount <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Amount must be greater than 0",
		})
	}

	tx, err := h.db.Begin()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to start transaction",
		})
	}
	defer tx.Rollback()

	switch req.Type {
	case models.TransactionTypeDeposit:
		_, err = tx.Exec(
			"UPDATE accounts SET balance = balance + $1, updated_at = $2 WHERE id = $3",
			req.Amount, time.Now(), req.ToAccountID,
		)
	case models.TransactionTypeWithdraw:
		var balance float64
		err = tx.QueryRow("SELECT balance FROM accounts WHERE id = $1", req.ToAccountID).Scan(&balance)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Account not found",
			})
		}
		if balance < req.Amount {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Insufficient funds",
			})
		}
		_, err = tx.Exec(
			"UPDATE accounts SET balance = balance - $1, updated_at = $2 WHERE id = $3",
			req.Amount, time.Now(), req.ToAccountID,
		)
	case models.TransactionTypeTransfer:
		if req.FromAccountID == nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "From account ID is required for transfers",
			})
		}
		var balance float64
		err = tx.QueryRow("SELECT balance FROM accounts WHERE id = $1", *req.FromAccountID).Scan(&balance)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "From account not found",
			})
		}
		if balance < req.Amount {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Insufficient funds",
			})
		}
		_, err = tx.Exec(
			"UPDATE accounts SET balance = balance - $1, updated_at = $2 WHERE id = $3",
			req.Amount, time.Now(), *req.FromAccountID,
		)
		if err == nil {
			_, err = tx.Exec(
				"UPDATE accounts SET balance = balance + $1, updated_at = $2 WHERE id = $3",
				req.Amount, time.Now(), req.ToAccountID,
			)
		}
	default:
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid transaction type",
		})
	}

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update account balance",
		})
	}

	var transactionID int
	err = tx.QueryRow(
		"INSERT INTO transactions (from_account_id, to_account_id, amount, type, description, created_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id",
		req.FromAccountID, req.ToAccountID, req.Amount, req.Type, req.Description, time.Now(),
	).Scan(&transactionID)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create transaction record",
		})
	}

	if err = tx.Commit(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to commit transaction",
		})
	}

	transaction := models.TransactionResponse{
		ID:            transactionID,
		FromAccountID: req.FromAccountID,
		ToAccountID:   req.ToAccountID,
		Amount:        req.Amount,
		Type:          req.Type,
		Description:   req.Description,
		CreatedAt:     time.Now(),
	}

	return c.Status(fiber.StatusCreated).JSON(transaction)
}

func (h *TransactionHandler) GetTransactions(c *fiber.Ctx) error {
	accountID := c.Query("account_id")
	
	var query string
	var args []interface{}
	
	if accountID != "" {
		query = "SELECT id, from_account_id, to_account_id, amount, type, description, created_at FROM transactions WHERE from_account_id = $1 OR to_account_id = $1 ORDER BY created_at DESC"
		args = append(args, accountID)
	} else {
		query = "SELECT id, from_account_id, to_account_id, amount, type, description, created_at FROM transactions ORDER BY created_at DESC"
	}

	rows, err := h.db.Query(query, args...)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve transactions",
		})
	}
	defer rows.Close()

	var transactions []models.TransactionResponse
	for rows.Next() {
		var transaction models.TransactionResponse
		var fromAccountID sql.NullInt32
		err := rows.Scan(
			&transaction.ID,
			&fromAccountID,
			&transaction.ToAccountID,
			&transaction.Amount,
			&transaction.Type,
			&transaction.Description,
			&transaction.CreatedAt,
		)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to scan transaction",
			})
		}
		
		if fromAccountID.Valid {
			fromID := int(fromAccountID.Int32)
			transaction.FromAccountID = &fromID
		}
		
		transactions = append(transactions, transaction)
	}

	return c.JSON(transactions)
}