package model

type Product struct {
	Name  string
	Price int
	ID    int
}

func (p *Product) GetPrice() int {
	return p.Price
}

func (p *Product) GetName() string {
	return p.Name
}

func (p *Product) GetID() int {
	return p.ID
}
