# Go Banking System

A simple banking system API built with Go, Fiber v2, and PostgreSQL.

## Features

- Create bank accounts
- Deposit money
- Withdraw money
- Transfer money between accounts
- View account details and transaction history

## Project Structure

```
go-banking/
├── cmd/api/                 # Application entry point
│   └── main.go
├── config/                  # Configuration
│   └── config.go
├── internal/
│   ├── database/           # Database connection
│   │   └── database.go
│   ├── handlers/           # HTTP handlers
│   │   ├── account.go
│   │   └── transaction.go
│   ├── middleware/         # HTTP middleware
│   │   └── logger.go
│   └── models/             # Data models
│       ├── account.go
│       └── transaction.go
├── migrations/             # Database migrations
│   ├── 001_create_accounts_table.sql
│   └── 002_create_transactions_table.sql
├── pkg/utils/              # Utility functions
├── .env.example            # Environment variables example
├── go.mod                  # Go module
└── README.md
```

## API Endpoints

### Health Check
- `GET /api/v1/health` - Check API status

### Accounts
- `POST /api/v1/accounts` - Create a new account
- `GET /api/v1/accounts` - Get all accounts
- `GET /api/v1/accounts/:id` - Get account by ID

### Transactions
- `POST /api/v1/transactions` - Create a new transaction (deposit/withdraw/transfer)
- `GET /api/v1/transactions` - Get all transactions
- `GET /api/v1/transactions?account_id=:id` - Get transactions for specific account

## Setup

1. Copy environment file:
   ```bash
   cp .env.example .env
   ```

2. Update database credentials in `.env`

3. Create PostgreSQL database:
   ```sql
   CREATE DATABASE banking_db;
   ```

4. Run migrations (manually execute SQL files in `migrations/` folder)

5. Install dependencies:
   ```bash
   go mod tidy
   ```

6. Run the application:
   ```bash
   go run cmd/api/main.go
   ```

## Example Usage

### Create Account
```bash
curl -X POST http://localhost:3000/api/v1/accounts \
  -H "Content-Type: application/json" \
  -d '{"name": "John Doe", "email": "john@example.com"}'
```

### Deposit Money
```bash
curl -X POST http://localhost:3000/api/v1/transactions \
  -H "Content-Type: application/json" \
  -d '{"to_account_id": 1, "amount": 100.00, "type": "deposit", "description": "Initial deposit"}'
```

### Transfer Money
```bash
curl -X POST http://localhost:3000/api/v1/transactions \
  -H "Content-Type: application/json" \
  -d '{"from_account_id": 1, "to_account_id": 2, "amount": 50.00, "type": "transfer", "description": "Transfer to friend"}'
```