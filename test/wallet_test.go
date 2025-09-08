package test

import (
	"context"
	"errors"
	"github.com/e_harsley/golang_backend_test/pkg/models"
	"github.com/e_harsley/golang_backend_test/pkg/repository"
	"github.com/e_harsley/golang_backend_test/pkg/services"
	"github.com/e_harsley/golang_backend_test/utils"
	"testing"
)

func TestMoney(t *testing.T) {
	t.Run("valid money creation", func(t *testing.T) {
		money, err := utils.NewMoney(1000)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if money.Cents() != 1000 {
			t.Errorf("Expected 1000 cents, got %d", money.Cents())
		}
		if money.Dollars() != 10.0 {
			t.Errorf("Expected 10.0 dollars, got %.2f", money.Dollars())
		}
	})

	t.Run("negative money creation fails", func(t *testing.T) {
		_, err := utils.NewMoney(-100)
		if err == nil {
			t.Error("Expected error for negative amount")
		}
	})

	t.Run("money operations", func(t *testing.T) {
		money1, _ := utils.NewMoney(1000)
		money2, _ := utils.NewMoney(500)

		// Addition
		sum := money1.Add(money2)
		if sum.Cents() != 1500 {
			t.Errorf("Expected 1500 cents, got %d", sum.Cents())
		}

		// Subtraction
		diff, err := money1.Subtract(money2)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if diff.Cents() != 500 {
			t.Errorf("Expected 500 cents, got %d", diff.Cents())
		}

		// Insufficient funds
		_, err = money2.Subtract(money1)
		if err == nil {
			t.Error("Expected insufficient funds error")
		}
	})
}

func TestWalletService_Transfer_Success(t *testing.T) {
	repo := repository.NewInMemoryRepository()
	service := services.NewWalletService(repo)
	ctx := context.Background()

	// Create wallets
	initialBalance1, _ := utils.NewMoney(10000)
	initialBalance2, _ := utils.NewMoney(5000)

	wallet1 := &models.Wallet{ID: "wallet1", Owner: "Alice", Balance: initialBalance1}
	wallet2 := &models.Wallet{ID: "wallet2", Owner: "Bob", Balance: initialBalance2}

	repo.Create(ctx, wallet1)
	repo.Create(ctx, wallet2)

	// Transfer amount
	transferAmount, _ := utils.NewMoney(3000) // $30.00

	// Execute transfer
	err := service.Transfer(ctx, "wallet1", "wallet2", transferAmount)

	// Verify
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Check balances
	updatedWallet1, _ := service.GetWallet(ctx, "wallet1")
	updatedWallet2, _ := service.GetWallet(ctx, "wallet2")

	expectedBalance1, _ := utils.NewMoney(7000) // $70.00
	expectedBalance2, _ := utils.NewMoney(8000) // $80.00

	if updatedWallet1.Balance.Cents() != expectedBalance1.Cents() {
		t.Errorf("Expected wallet1 balance %d, got %d",
			expectedBalance1.Cents(), updatedWallet1.Balance.Cents())
	}

	if updatedWallet2.Balance.Cents() != expectedBalance2.Cents() {
		t.Errorf("Expected wallet2 balance %d, got %d",
			expectedBalance2.Cents(), updatedWallet2.Balance.Cents())
	}
}

func TestWalletService_Transfer_InsufficientFunds(t *testing.T) {
	repo := repository.NewInMemoryRepository()
	service := services.NewWalletService(repo)
	ctx := context.Background()

	// Create wallets
	initialBalance1, _ := utils.NewMoney(1000) // $10.00
	initialBalance2, _ := utils.NewMoney(5000) // $50.00

	wallet1 := &models.Wallet{ID: "wallet1", Owner: "Alice", Balance: initialBalance1}
	wallet2 := &models.Wallet{ID: "wallet2", Owner: "Bob", Balance: initialBalance2}

	repo.Create(ctx, wallet1)
	repo.Create(ctx, wallet2)

	// Transfer amount greater than balance
	transferAmount, _ := utils.NewMoney(3000) // $30.00

	// Execute transfer
	err := service.Transfer(ctx, "wallet1", "wallet2", transferAmount)

	// Verify error
	if err == nil {
		t.Error("Expected insufficient funds error")
	}
	if !errors.Is(err, utils.ErrInsufficientFunds) {
		t.Errorf("Expected ErrInsufficientFunds, got %v", err)
	}

	wallet1After, _ := service.GetWallet(ctx, "wallet1")
	wallet2After, _ := service.GetWallet(ctx, "wallet2")

	if wallet1After.Balance.Cents() != initialBalance1.Cents() {
		t.Error("Wallet1 balance should be unchanged after failed transfer")
	}
	if wallet2After.Balance.Cents() != initialBalance2.Cents() {
		t.Error("Wallet2 balance should be unchanged after failed transfer")
	}
}

func TestWalletService_Transfer_InvalidAmount(t *testing.T) {
	repo := repository.NewInMemoryRepository()
	service := services.NewWalletService(repo)
	ctx := context.Background()

	// Test zero amount
	zeroAmount, _ := utils.NewMoney(0)
	err := service.Transfer(ctx, "wallet1", "wallet2", zeroAmount)
	if !errors.Is(err, utils.ErrInvalidAmount) {
		t.Errorf("Expected ErrInvalidAmount for zero amount, got %v", err)
	}
}

func TestWalletService_Transfer_SameWallet(t *testing.T) {
	repo := repository.NewInMemoryRepository()
	service := services.NewWalletService(repo)
	ctx := context.Background()

	amount, _ := utils.NewMoney(1000)
	err := service.Transfer(ctx, "wallet1", "wallet1", amount)
	if !errors.Is(err, utils.ErrSameWallet) {
		t.Errorf("Expected ErrSameWallet, got %v", err)
	}
}
