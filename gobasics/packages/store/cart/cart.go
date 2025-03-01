package cart

import "packages/store"

type Cart struct {
	CustomerName string
	Items        []store.Product
}

func (c *Cart) GetTotal() (total float64) {
	for _, v := range c.Items {
		total += v.Price()
	}
	return
}
