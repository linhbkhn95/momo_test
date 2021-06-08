package productinventory

import "momo.test.com/pkg/vendingmachine"

type productInventoryMomery struct {
	data map[int]int
}

func NewProductInventoryMomery(data map[int]int) vendingmachine.ProductInventory {
	return &productInventoryMomery{
		data: data,
	}
}

func (i *productInventoryMomery) GetStock(productID int) (int, error) {
	if i.data == nil {
		return 0, nil
	}
	stock := i.data[productID]
	return stock, nil
}
func (i *productInventoryMomery) UpdateStock(productID int, stockQuantity int) error {
	if stockQuantity < 0 {
		return nil
	}
	i.data[productID] = stockQuantity
	return nil
}
