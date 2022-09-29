package models

import (
	"log"
	"os"
	"strconv"
)

type Wallet struct {
	balance int
}

func (wallet *Wallet) GetAvailableFunds() int {
	return wallet.balance
}

func (wallet *Wallet) AddFunds(funds int) {
	wallet.balance += funds
}

func (wallet *Wallet) RemoveFunds(funds int) {
	wallet.balance -= funds
}

func InitializeWallet() Wallet {
	content, err := os.ReadFile("wallet.txt")

	if err != nil {
		log.Fatal(err)
	}
	fundsString := string(content)
	funds, _ := strconv.Atoi(fundsString)
	wallet := Wallet{}
	wallet.AddFunds(funds)

	return wallet
}
