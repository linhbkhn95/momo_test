package vendingmachine

import "fmt"

type NotEnoughMoneyError struct {
	money        int
	productPrice int
}

func (e NotEnoughMoneyError) Error() string {
	return fmt.Sprintf("not enough money with current money=%d for product price=%d", e.money, e.productPrice)
}

type NotSufficientChangedError struct {
	remain int
}

func (e NotSufficientChangedError) Error() string {
	return fmt.Sprintf("not sufficient change in cashinventory with remain=%d", e.remain)
}

type OutOfProductStockError struct {
	productName string
}

func (e OutOfProductStockError) Error() string {
	return fmt.Sprintf("out of stock of product =%s", e.productName)
}

type ProductNotFoundError struct {
	productID int
}

func (e ProductNotFoundError) Error() string {
	return fmt.Sprintf("not found product =%d", e.productID)
}
