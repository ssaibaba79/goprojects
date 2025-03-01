package store

import "fmt"

var storeDiscount = discount{10.0, 0.05}

type Product struct {
	Name, Category string
	price          float64
}

func NewProduct(name, category string, price float64) *Product {
	return &Product{name, category, price}
}
func (p *Product) SetPrice(newPrice float64) error {
	if newPrice < 0 {
		return fmt.Errorf("price cannot be less than 0")
	}
	p.price = newPrice
	return nil
}

func (p *Product) Price() float64 {

	d := storeDiscount.calculateDiscount(p.price)
	return p.price - d
}
