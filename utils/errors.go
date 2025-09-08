package utils

import "errors"

var (
	ErrWalletNotFound    = errors.New("wallet not found")
	ErrInsufficientFunds = errors.New("insufficient funds")
	ErrInvalidAmount     = errors.New("invalid transfer amount")
	ErrSameWallet        = errors.New("cannot transfer to the same wallet")
)
