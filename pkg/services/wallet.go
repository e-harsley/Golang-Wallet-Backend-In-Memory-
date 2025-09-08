/*
services/wallet.go

@author: Harsley Ekhorutomwen
@created: September 8, 2025

Business logic for wallet operations including create, get, and transfer.

Key Features:
    - Validates transfer rules (positive amount, distinct wallets, sufficient funds)
    - Uses repository transaction interface for atomic updates
    - Returns domain-specific errors from utils

:module: pkg/services
:requires: repository, models, utils
*/
package services

import (
	"context"
	"fmt"
	"github.com/e_harsley/golang_backend_test/pkg/models"
	"github.com/e_harsley/golang_backend_test/pkg/repository"
	"github.com/e_harsley/golang_backend_test/utils"
)

type WalletService struct {
	repo repository.WalletRepository
}

func NewWalletService(repo repository.WalletRepository) *WalletService {
	return &WalletService{
		repo: repo,
	}
}

func (s *WalletService) Transfer(ctx context.Context, fromID, toID string, amount utils.Money) error {
	if fromID == toID {
		return utils.ErrSameWallet
	}
	if !amount.IsPositive() {
		return utils.ErrInvalidAmount
	}

	tx, err := s.repo.BeginTransaction(ctx)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback(ctx)

	fromWallet, err := tx.GetByID(ctx, fromID)
	if err != nil {
		return fmt.Errorf("failed to get sender wallet: %w", err)
	}
	if fromWallet == nil {
		return utils.ErrWalletNotFound
	}

	toWallet, err := tx.GetByID(ctx, toID)
	if err != nil {
		return fmt.Errorf("failed to get recipient wallet: %w", err)
	}
	if toWallet == nil {
		return utils.ErrWalletNotFound
	}

	if !fromWallet.Balance.GreaterThanOrEqual(amount) {
		return utils.ErrInsufficientFunds
	}

	newFromBalance, err := fromWallet.Balance.Subtract(amount)
	if err != nil {
		return fmt.Errorf("failed to deduct from sender: %w", err)
	}

	newToBalance := toWallet.Balance.Add(amount)

	fromWallet.Balance = newFromBalance
	toWallet.Balance = newToBalance

	if err := tx.Update(ctx, fromWallet); err != nil {
		return fmt.Errorf("failed to update sender wallet: %w", err)
	}

	if err := tx.Update(ctx, toWallet); err != nil {
		return fmt.Errorf("failed to update recipient wallet: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func (s *WalletService) CreateWallet(ctx context.Context, id, owner string, initialBalance utils.Money) error {
	wallet := &models.Wallet{
		ID:      id,
		Owner:   owner,
		Balance: initialBalance,
	}
	return s.repo.Create(ctx, wallet)
}

func (s *WalletService) GetWallet(ctx context.Context, id string) (*models.Wallet, error) {
	return s.repo.GetByID(ctx, id)
}
