/*
models/wallet.go

@author: Harsley Ekhorutomwen
@created: September 8, 2025

Domain model for the wallet entity used across the service and repositories.

Key Features:
    - Simple struct with ID, Owner, and Money balance
    - Decoupled from storage; suitable for in-memory or DB-backed repos

:module: pkg/models
:requires: utils.Money
*/
package models

import "github.com/e_harsley/golang_backend_test/utils"

type Wallet struct {
	ID      string      `json:"id"`
	Owner   string      `json:"owner"`
	Balance utils.Money `json:"balance"`
}
