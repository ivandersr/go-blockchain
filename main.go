package main

import (
	"fmt"
	"log"

	"github.com/ivandersr/go-blockchain/blockchain"
	"github.com/ivandersr/go-blockchain/wallet"
)

func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {
	minersWallet := wallet.NewWallet()
	walletA := wallet.NewWallet()
	walletB := wallet.NewWallet()

	// wallet side
	t := wallet.NewTransaction(
		walletA.PrivateKey(),
		walletA.PublicKey(),
		walletA.BlockchainAddress(),
		walletB.BlockchainAddress(),
		1.0,
	)

	// blockchain side
	bc := blockchain.NewBlockchain(minersWallet.BlockchainAddress())
	isAdded := bc.AddTransaction(
		walletA.BlockchainAddress(),
		walletB.BlockchainAddress(),
		1.0, walletA.PublicKey(),
		t.GenerateSignature())
	fmt.Printf("Added? %v\n", isAdded)

	bc.Mine()
	bc.Print()

	fmt.Printf("A %.2f\n", bc.CalculateTotalAmount(walletA.BlockchainAddress()))
	fmt.Printf("B %.2f\n", bc.CalculateTotalAmount(walletB.BlockchainAddress()))
	fmt.Printf("Miner %.2f\n", bc.CalculateTotalAmount(minersWallet.BlockchainAddress()))
}
