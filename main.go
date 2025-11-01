package main

import (
	"fmt"
	"log"

	"github.com/ivandersr/go-blockchain/blockchain"
)

func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {
	blockchainAddress := "blockchain_address_X"
	bc := blockchain.NewBlockchain(blockchainAddress)
	bc.Print()

	bc.AddTransaction("a", "b", 1.0)

	bc.Mine()
	bc.Print()

	bc.AddTransaction("x", "y", 2.0)
	bc.AddTransaction("c", "d", 3.0)

	bc.Mine()
	bc.Print()

	fmt.Printf("C %.2f\n", bc.CalculateTotalAmount("c"))
	fmt.Printf("D %.2f\n", bc.CalculateTotalAmount("d"))
	fmt.Printf("X %.2f\n", bc.CalculateTotalAmount("x"))
	fmt.Printf("Y %.2f\n", bc.CalculateTotalAmount("y"))
}
