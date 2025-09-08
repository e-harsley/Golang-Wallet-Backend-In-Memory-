/*
repository/failing_repository.go

@author: Harsley Ekhorutomwen
@created: September 8, 2025

Repository implementation that intentionally fails for all operations. Useful
for testing error paths in services.

Key Features:
    - Deterministic failures on Get, Update, Create, and BeginTransaction

:module: pkg/repository
:requires: models
*/
package repository

import (
	"context"
	"errors"
	"github.com/e_harsley/golang_backend_test/pkg/models"
)

type FailingRepository struct{}

func NewFailingRepository() *FailingRepository {
	return &FailingRepository{}
}

func (f *FailingRepository) GetByID(ctx context.Context, id string) (*models.Wallet, error) {
	return nil, errors.New("failing repository: get operation failed")
}

func (f *FailingRepository) Update(ctx context.Context, wallet *models.Wallet) error {
	return errors.New("failing repository: update operation failed")
}

func (f *FailingRepository) Create(ctx context.Context, wallet *models.Wallet) error {
	return errors.New("failing repository: create operation failed")
}

func (f *FailingRepository) BeginTransaction(ctx context.Context) (Transaction, error) {
	return nil, errors.New("failing repository: transaction failed")
}
