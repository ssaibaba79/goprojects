package main

type Product struct {
	Name, Category string
	Price          float64
}

var ProductList = []*Product{
	{"Bread", "Bakery", 3.29},
	{"Cookies", "Bakery", 4.25},
	{"Milk", "Dairy", 4.19},
	{"Eggs", "Dairy", 6.29},
	{"Tide", "Detergent", 13.99},
}

type ProductGroup []*Product

type ProductData = map[string]ProductGroup

var Products = make(ProductData)

func init() {
	for _, p := range ProductList {
		if _, ok := Products[p.Category]; ok {
			Products[p.Category] = append(Products[p.Category], p)
		} else {
			Products[p.Category] = ProductGroup{p}
		}
	}
}

func main() {

}
