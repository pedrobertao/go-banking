package models

import (
	"time"
)

type TransactionType string

const (
	TransactionTypeDeposit  TransactionType = "deposit"
	TransactionTypeWithdraw TransactionType = "withdraw"
	TransactionTypeTransfer TransactionType = "transfer"
)

type Transaction struct {
	ID            int             `json:"id" db:"id"`
	FromAccountID *int            `json:"from_account_id" db:"from_account_id"`
	ToAccountID   int             `json:"to_account_id" db:"to_account_id"`
	Amount        float64         `json:"amount" db:"amount"`
	Type          TransactionType `json:"type" db:"type"`
	Description   string          `json:"description" db:"description"`
	CreatedAt     time.Time       `json:"created_at" db:"created_at"`
}

type CreateTransactionRequest struct {
	ToAccountID   int             `json:"to_account_id" validate:"required"`
	FromAccountID *int            `json:"from_account_id,omitempty"`
	Amount        float64         `json:"amount" validate:"required,gt=0"`
	Type          TransactionType `json:"type" validate:"required"`
	Description   string          `json:"description"`
}

type TransactionResponse struct {
	ID            int             `json:"id"`
	FromAccountID *int            `json:"from_account_id"`
	ToAccountID   int             `json:"to_account_id"`
	Amount        float64         `json:"amount"`
	Type          TransactionType `json:"type"`
	Description   string          `json:"description"`
	CreatedAt     time.Time       `json:"created_at"`
}