package main

import (
	"fmt"

	"momo.test.com/pkg/cashinventory"
	"momo.test.com/pkg/model"
	"momo.test.com/pkg/productinventory"
	"momo.test.com/pkg/vendingmachine"
)

func main() {
	coke := model.Product{
		Name:  "Coke",
		Price: 10000,
		ID:    1,
	}
	pesi := model.Product{
		Name:  "Pesi",
		Price: 10000,
		ID:    2,
	}
	soda := model.Product{
		Name:  "Soda",
		Price: 20000,
		ID:    3,
	}
	products := map[int]*model.Product{
		coke.ID: &coke,
		pesi.ID: &pesi,
		soda.ID: &soda,
	}
	productInventoryData := map[int]int{
		coke.ID: 5,
		pesi.ID: 7,
		soda.ID: 10,
	}
	cashInventoryData := map[model.Coin]int{
		model.TenThousand:        5,
		model.TwentyThousand:     7,
		model.FixtyThousand:      10,
		model.OneHundredThousand: 10,
		model.TwoHundredThousand: 10,
	}

	cashInventory := cashinventory.NewCashInventoryMem(cashInventoryData)
	productInventory := productinventory.NewProductInventoryMomery(productInventoryData)

	vendingMachine := vendingmachine.NewMachine(products, cashInventory, productInventory)
	fmt.Println("vending machine receives 100000")

	vendingMachine.ReceiveMoney(model.OneHundredThousand)
	bucket, err := vendingMachine.ReleaseProductAndChange(coke.ID)
	if err != nil {
		fmt.Printf("err when process cause by %v", err)

	}
	if bucket != nil {
		fmt.Println("release product", bucket.GetProduct().GetName())
		fmt.Println("remaining change", bucket.GetCoins())
		fmt.Println("remaining change total", toMoney(bucket.GetCoins()))
	}

	fmt.Println("vending machine receives 100000")

	vendingMachine.ReceiveMoney(model.OneHundredThousand)

	fmt.Println("cancel the request and receive a refund", vendingMachine.Refund())

}

func toMoney(coins map[model.Coin]int) int {
	result := 0

	for coin, quantity := range coins {
		result += int(coin) * quantity
	}
	return result
}
