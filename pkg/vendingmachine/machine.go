package vendingmachine

import (
	"momo.test.com/pkg/model"
)

type (
	Machine interface {
		ReceiveMoney(coins model.Coin)
		ReleaseProductAndChange(productID int) (*model.Bucket, error)
		Refund() map[model.Coin]int
	}

	CashInventory interface {
		GetStock(c model.Coin) (int, error)
		UpdateStock(c model.Coin, quantity int) error
	}

	ProductInventory interface {
		GetStock(productID int) (int, error)
		UpdateStock(productID int, quantity int) error
	}

	machine struct {
		userTotalMoney   int
		cashInventory    CashInventory
		productInventory ProductInventory
		products         map[int]*model.Product
	}
)

func NewMachine(products map[int]*model.Product, cashInventory CashInventory, productInventory ProductInventory) Machine {
	return &machine{
		cashInventory:    cashInventory,
		productInventory: productInventory,
		products:         products,
	}
}

func (m *machine) ReceiveMoney(coin model.Coin) {
	m.userTotalMoney += int(coin)
	stock, _ := m.cashInventory.GetStock(coin)
	m.cashInventory.UpdateStock(coin, stock+1)
}

func (m *machine) ReleaseProductAndChange(productID int) (*model.Bucket, error) {
	p, ok := m.products[productID]
	if !ok {
		return nil, ProductNotFoundError{
			productID: productID,
		}
	}
	totalMoney := m.userTotalMoney
	if totalMoney < p.GetPrice() {
		return nil, NotEnoughMoneyError{productPrice: p.Price, money: totalMoney}
	}
	if !m.hasItem(p) {
		return nil, OutOfProductStockError{productName: p.GetName()}
	}
	return m.releaseProductAndChanged(totalMoney, p)
}

func (m *machine) hasItem(p *model.Product) bool {
	stock, _ := m.productInventory.GetStock(p.GetID())
	return stock > 0
}

func (m *machine) releaseProductAndChanged(totalMoney int, p *model.Product) (*model.Bucket, error) {
	remain := totalMoney - p.GetPrice()

	changedCoins, err := m.getChange(remain)
	if err != nil {
		return nil, err
	}
	if !m.enoughStockCoins(changedCoins) {
		return nil, &NotSufficientChangedError{remain: remain}
	}
	m.updateCashInvetory(changedCoins)
	m.userTotalMoney = 0
	return model.NewBucket(p, changedCoins), nil
}

func (m *machine) enoughStockCoins(changedCoins map[model.Coin]int) bool {
	for coin, quantity := range changedCoins {
		stock, _ := m.cashInventory.GetStock(coin)
		if stock < quantity {
			return false
		}
	}
	return true
}

func (m *machine) updateCashInvetory(changedCoins map[model.Coin]int) {
	for coin, quantity := range changedCoins {
		stock, _ := m.cashInventory.GetStock(coin)
		m.cashInventory.UpdateStock(coin, stock-quantity)
	}
}

func (m *machine) getChange(remain int) (map[model.Coin]int, error) {
	result := make(map[model.Coin]int)
	amount := remain
	for amount > 0 {
		coin := model.Coin(-1)
		if amount >= int(model.TwoHundredThousand) {
			coin = model.TwoHundredThousand
		} else if amount >= int(model.OneHundredThousand) {
			coin = model.OneHundredThousand
		} else if amount >= int(model.FixtyThousand) {
			coin = model.FixtyThousand
		} else if amount >= int(model.TwentyThousand) {
			coin = model.TwentyThousand
		} else if amount >= int(model.TenThousand) {
			coin = model.TenThousand
		} else {
			// retur error cannot changed
			return nil, &NotSufficientChangedError{remain: remain}
		}

		if int(coin) != -1 {
			amount = amount - int(coin)
			quantity, ok := result[coin]
			if !ok {
				result[coin] = 1
			}
			result[coin] = quantity + 1
		}
	}
	return result, nil
}

func (m *machine) Refund() map[model.Coin]int {
	result, _ := m.getChange(m.userTotalMoney)
	m.userTotalMoney = 0
	return result
}
