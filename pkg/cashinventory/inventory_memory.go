package cashinventory

import (
	"momo.test.com/pkg/model"
	"momo.test.com/pkg/vendingmachine"
)

type cashInventory struct {
	data map[model.Coin]int
}

func NewCashInventoryMem(data map[model.Coin]int) vendingmachine.CashInventory {
	return &cashInventory{
		data: data,
	}
}

func (cash *cashInventory) UpdateStock(coin model.Coin, stock int) error {
	cash.data[coin] = stock
	return nil
}

func (cash *cashInventory) GetStock(coin model.Coin) (int, error) {
	stock := cash.data[coin]
	return stock, nil
}
