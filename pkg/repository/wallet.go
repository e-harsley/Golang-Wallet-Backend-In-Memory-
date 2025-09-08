/*
repository/wallet.go

@author: Harsley Ekhorutomwen
@created: September 8, 2025

Thread-safe in-memory repository implementation for wallets with basic CRUD
and a transaction factory.

Key Features:
    - RWMutex-protected map-based storage
    - Copy-on-read/write to avoid external mutation
    - Provides BeginTransaction to supply a Transaction implementation

:module: pkg/repository
:requires: models, utils, sync
*/
package repository

import (
	"context"
	"fmt"
	"github.com/e_harsley/golang_backend_test/pkg/models"
	"github.com/e_harsley/golang_backend_test/utils"
	"sync"
)

type (
	WalletRepository interface {
		GetByID(ctx context.Context, id string) (*models.Wallet, error)
		Update(ctx context.Context, wallet *models.Wallet) error
		Create(ctx context.Context, wallet *models.Wallet) error
		BeginTransaction(ctx context.Context) (Transaction, error)
	}

	InMemoryRepository struct {
		mu      sync.RWMutex
		wallets map[string]*models.Wallet
	}
)

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{
		wallets: make(map[string]*models.Wallet),
	}
}

func (r *InMemoryRepository) GetByID(ctx context.Context, id string) (*models.Wallet, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	wallet, exists := r.wallets[id]
	if !exists {
		return nil, utils.ErrWalletNotFound
	}

	walletCopy := *wallet
	return &walletCopy, nil
}

func (r *InMemoryRepository) Update(ctx context.Context, wallet *models.Wallet) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.wallets[wallet.ID]; !exists {
		return utils.ErrWalletNotFound
	}

	walletCopy := *wallet
	r.wallets[wallet.ID] = &walletCopy
	return nil
}

func (r *InMemoryRepository) Create(ctx context.Context, wallet *models.Wallet) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.wallets[wallet.ID]; exists {
		return fmt.Errorf("wallet with ID %s already exists", wallet.ID)
	}

	walletCopy := *wallet
	r.wallets[wallet.ID] = &walletCopy
	return nil
}

func (r *InMemoryRepository) BeginTransaction(ctx context.Context) (Transaction, error) {
	return &InMemoryTransaction{repo: r}, nil
}
