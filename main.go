package main

import (
	"encoding/csv"
	"fmt"
	"log"

	"github.com/nickater/go-vending-machine/models"
	"github.com/nickater/go-vending-machine/ui"
	"github.com/nickater/go-vending-machine/utils"
)

func fetchProducts() []models.RawProduct {
	productsFile, _ := utils.CheckFile("./products.csv")
	defer productsFile.Close()

	csvReader := csv.NewReader(productsFile)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	rawProducts := utils.DecipherProductCSV(data)

	return rawProducts
}

func fetchWallet() models.Wallet {
	wallet := models.InitializeWallet()
	return wallet
}

func main() {
	fmt.Println("Welcome to the Circus of Value!")
	wallet := fetchWallet()
	vendingMachine := models.VendingMachine{}
	products := fetchProducts()
	vendingMachine.StartVendingMachine(products)
	vendingMachine.BuyItem("A1")
	fmt.Println(wallet.GetAvailableFunds())
	fmt.Println(vendingMachine.TotalFunds)
	vendingMachine.InsertFunds(1200)
	fmt.Println(vendingMachine.TotalFunds)
	funds := vendingMachine.ReturnFunds()
	fmt.Println(vendingMachine.TotalFunds)
	fmt.Println("change", funds)
	ui.DisplayMainMenu()
}
