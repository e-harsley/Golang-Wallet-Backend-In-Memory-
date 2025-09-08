## Golang Wallet Backend (In-Memory)

A simple Go project that models wallets and transfers between them using an in-memory repository and a small service layer. It includes unit tests for money arithmetic and transfer scenarios.

### Features
- **Money type**: Safe cents-based arithmetic (`utils.Money`).
- **Wallet model**: `id`, `owner`, `balance`.
- **Service layer**: Validates transfers, supports basic transaction semantics.
- **In-memory repository**: Thread-safe store with a transaction interface.
- **Unit tests**: Money operations and wallet transfers (success and error paths).

### Requirements
- Go 1.23+ (as specified in `go.mod`)

### Getting Started
1. Clone the repository
```bash
git clone https://github.com/e_harsley/golang_backend_test.git
cd golang_backend_test
```

2. Run the demo program
```bash
go run main.go
```
This will:
- Create two wallets (`alice`, `bob`)
- Print initial balances
- Transfer $25.00 from Alice to Bob
- Print final balances

Example output:
```text
Initial balances:
Alice: $100.00
Bob: $50.00

Transferring $25.00 from Alice to Bob...

Final balances:
Alice: $75.00
Bob: $75.00
```

### Running Tests
```bash
go test ./test
```
The tests cover:
- `utils.Money` creation, addition, subtraction, and edge cases
- `WalletService.Transfer` success, insufficient funds, invalid amount, and same-wallet cases

### Project Structure
```text
.
├── main.go                       # Demo program wiring repository + service
├── pkg
│   ├── models
│   │   └── wallet.go             # Wallet domain model
│   ├── repository
│   │   ├── wallet.go             # In-memory repository (CRUD + factory)
│   │   └── transaction.go        # Transaction interface + in-memory impl
│   └── services
│       └── wallet.go             # WalletService with Transfer/Create/Get
├── test
│   └── wallet_test.go            # Unit tests
└── utils
    ├── errors.go                 # Shared error values
    └── money.go                  # Money type and operations
```

### Key Concepts
- **Money type (`utils.Money`)**: Stores amounts in cents (int64) to avoid floating-point errors. Provides `Add`, `Subtract`, `GreaterThanOrEqual`, `IsPositive`, `String`, etc.
- **Transactions**: `repository.Transaction` interface with `GetByID`, `Update`, `Commit`, and `Rollback`. The in-memory implementation is lightweight; `Commit`/`Rollback` are no-ops to keep the example simple.
- **Service validation**: `Transfer` prevents self-transfers, non-positive amounts, and insufficient funds, and ensures both wallets exist before applying updates.

### Module Path
The module path is `github.com/e_harsley/golang_backend_test` (see `go.mod`). If you fork or rename the repository, update `go.mod` accordingly.

### Notes
- This project intentionally avoids external dependencies for clarity.
- The in-memory repository is suitable for demos and tests; swap it for a real database-backed implementation for production use.



