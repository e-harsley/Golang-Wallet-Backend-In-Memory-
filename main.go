package main

import (
	"context"
	"fmt"
	"github.com/e_harsley/golang_backend_test/pkg/repository"
	"github.com/e_harsley/golang_backend_test/pkg/services"
	"github.com/e_harsley/golang_backend_test/utils"
	"log"
)

func main() {
	repo := repository.NewInMemoryRepository()
	service := services.NewWalletService(repo)
	ctx := context.Background()

	// Create wallets
	balance1, _ := utils.NewMoney(10000)
	balance2, _ := utils.NewMoney(5000)

	err := service.CreateWallet(ctx, "alice", "Alice Smith", balance1)
	if err != nil {
		log.Fatalf("Failed to create Alice's wallet: %v", err)
	}

	err = service.CreateWallet(ctx, "bob", "Bob Jones", balance2)
	if err != nil {
		log.Fatalf("Failed to create Bob's wallet: %v", err)
	}

	aliceWallet, _ := service.GetWallet(ctx, "alice")
	bobWallet, _ := service.GetWallet(ctx, "bob")

	fmt.Printf("Initial balances:\n")
	fmt.Printf("Alice: %s\n", aliceWallet.Balance.String())
	fmt.Printf("Bob: %s\n", bobWallet.Balance.String())

	transferAmount, _ := utils.NewMoney(2500)
	fmt.Printf("\nTransferring %s from Alice to Bob...\n", transferAmount.String())

	err = service.Transfer(ctx, "alice", "bob", transferAmount)
	if err != nil {
		log.Fatalf("Transfer failed: %v", err)
	}

	aliceWallet, _ = service.GetWallet(ctx, "alice")
	bobWallet, _ = service.GetWallet(ctx, "bob")

	fmt.Printf("\nFinal balances:\n")
	fmt.Printf("Alice: %s\n", aliceWallet.Balance.String())
	fmt.Printf("Bob: %s\n", bobWallet.Balance.String())

}
