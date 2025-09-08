/*
repository/mock_repository.go

@author: Harsley Ekhorutomwen
@created: September 8, 2025

Configurable in-memory mock repository for controlled success and failure
scenarios in tests.

Key Features:
    - Toggleable errors for Get, Update, Create, and transactions
    - In-memory storage with copy semantics

:module: pkg/repository
:requires: models, utils
*/
package repository

import (
	"context"
	"github.com/e_harsley/golang_backend_test/pkg/models"
	"github.com/e_harsley/golang_backend_test/utils"
)

type (
	MockRepository struct {
		wallets     map[string]*models.Wallet
		GetError    error
		UpdateError error
		CreateError error
		TxError     error
	}
	MockTransaction struct {
		repo *MockRepository
	}
)

func NewMockRepository() *MockRepository {
	return &MockRepository{
		wallets: make(map[string]*models.Wallet),
	}
}

func (m *MockRepository) GetByID(ctx context.Context, id string) (*models.Wallet, error) {
	if m.GetError != nil {
		return nil, m.GetError
	}

	wallet, exists := m.wallets[id]
	if !exists {
		return nil, utils.ErrWalletNotFound
	}

	walletCopy := *wallet
	return &walletCopy, nil
}

func (m *MockRepository) Update(ctx context.Context, wallet *models.Wallet) error {
	if m.UpdateError != nil {
		return m.UpdateError
	}

	if _, exists := m.wallets[wallet.ID]; !exists {
		return utils.ErrWalletNotFound
	}

	walletCopy := *wallet
	m.wallets[wallet.ID] = &walletCopy
	return nil
}

func (m *MockRepository) Create(ctx context.Context, wallet *models.Wallet) error {
	if m.CreateError != nil {
		return m.CreateError
	}

	walletCopy := *wallet
	m.wallets[wallet.ID] = &walletCopy
	return nil
}

func (m *MockRepository) BeginTransaction(ctx context.Context) (Transaction, error) {
	if m.TxError != nil {
		return nil, m.TxError
	}
	return &MockTransaction{repo: m}, nil
}

// MOCK TRANSACTIONS

func (tx *MockTransaction) GetByID(ctx context.Context, id string) (*models.Wallet, error) {
	return tx.repo.GetByID(ctx, id)
}

func (tx *MockTransaction) Update(ctx context.Context, wallet *models.Wallet) error {
	return tx.repo.Update(ctx, wallet)
}

func (tx *MockTransaction) Commit(ctx context.Context) error {
	return nil
}

func (tx *MockTransaction) Rollback(ctx context.Context) error {
	return nil
}
