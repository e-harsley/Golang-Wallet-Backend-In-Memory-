/*
repository/transaction.go

@author: Harsley Ekhorutomwen
@created: September 8, 2025

Transaction interface and in-memory implementation used by the repository
and service layer to encapsulate atomic updates.

Key Features:
    - Interface with GetByID, Update, Commit, Rollback
    - In-memory implementation is a no-op for Commit/Rollback

:module: pkg/repository
:requires: models
*/
package repository

import (
	"context"
	"github.com/e_harsley/golang_backend_test/pkg/models"
)

type (
	Transaction interface {
		GetByID(ctx context.Context, id string) (*models.Wallet, error)
		Update(ctx context.Context, wallet *models.Wallet) error
		Commit(ctx context.Context) error
		Rollback(ctx context.Context) error
	}

	InMemoryTransaction struct {
		repo *InMemoryRepository
	}
)

func (tx *InMemoryTransaction) GetByID(ctx context.Context, id string) (*models.Wallet, error) {
	return tx.repo.GetByID(ctx, id)
}

func (tx *InMemoryTransaction) Update(ctx context.Context, wallet *models.Wallet) error {
	return tx.repo.Update(ctx, wallet)
}

func (tx *InMemoryTransaction) Commit(ctx context.Context) error {
	return nil
}

func (tx *InMemoryTransaction) Rollback(ctx context.Context) error {
	return nil
}
