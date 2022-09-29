package models

import (
	"fmt"
)

type Slot struct {
	Selector string
	Product  Product
	Quantity int
}

type VendingMachine struct {
	Slots      []Slot
	TotalFunds int
	IsRunning  bool
}

func buildSelectorOptions() []string {
	selectorOptions := []string{"A1", "A2", "A3", "A4", "B1", "B2", "B3", "B4", "C1", "C2", "C3", "C4", "D1", "D2", "D3", "D4"}
	return selectorOptions
}

func (vendingMachine *VendingMachine) StartVendingMachine(products []RawProduct) {
	vendingMachine.AddProductsToSlots(products)
	vendingMachine.TotalFunds = 0
}

func (vendingMachine *VendingMachine) AddProductsToSlots(rawProducts []RawProduct) {
	selectorOptions := buildSelectorOptions()

	vendingMachine.Slots = make([]Slot, len(rawProducts))
	for i := range vendingMachine.Slots {
		vendingMachine.Slots[i] = Slot{
			Selector: selectorOptions[i],
		}
	}
}

func (vendingMachine *VendingMachine) InsertFunds(funds int) {
	vendingMachine.TotalFunds += funds
}

func (vendingMachine *VendingMachine) ReturnFunds() int {
	changeToReturn := vendingMachine.TotalFunds
	vendingMachine.TotalFunds = 0

	return changeToReturn
}

func (vendingMachine *VendingMachine) BuyItem(selector string) (bool, error) {

	for i := range vendingMachine.Slots {

		if vendingMachine.Slots[i].Selector == selector {
			if vendingMachine.Slots[i].Quantity > 0 {
				vendingMachine.Slots[i].Quantity -= 1
			} else {
				return false, fmt.Errorf("insufficient quantity")
			}
		}
	}

	return true, nil
}
