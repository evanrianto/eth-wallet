package main

import (
	"crypto/ecdsa"
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run walletgen.go <command>")
		fmt.Println("\nCommands:\n  generate: Generates a new Ethereum wallet keypair")
		return
	}

	switch os.Args[1] {
	case "generate":
		generateKeyPair()
	default:
		fmt.Println("Invalid command. Available command is 'generate'.")
	}
}

// generateKeyPair generates a new Ethereum wallet key pair and prints it
func generateKeyPair() {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatalf("Error generating Ethereum key pair: %v", err)
	}

	privateKeyBytes := crypto.FromECDSA(privateKey)
	fmt.Printf("Private Key: %x\n", privateKeyBytes)

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey) // Corrected the type assertion here
	if !ok {
		log.Fatal("Error casting public key to ECDSA")
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	fmt.Printf("Public Key: %x\n", publicKeyBytes)
}
