package model

type Bucket struct {
	coins   map[Coin]int
	product *Product
}

func NewBucket(p *Product, coins map[Coin]int) *Bucket {
	return &Bucket{
		coins:   coins,
		product: p,
	}
}

func (b Bucket) GetProduct() *Product {
	return b.product
}

func (b Bucket) GetCoins() map[Coin]int {
	return b.coins
}
