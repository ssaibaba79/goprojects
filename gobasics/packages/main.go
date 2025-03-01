package main

import (
	"fmt"
	_ "packages/data"
	cfmt "packages/fmt"
	"packages/store"

	"github.com/fatih/color"
)

func main() {

	//s := store.Product{Name: "Eraser", Category: "Stationary", Price: 0.99}
	p := store.NewProduct("Eraser", "Stationary", 0.99)
	fmt.Println(p.Name, p.Category, p.Price())
	p.SetPrice(0.78)
	fmt.Println(p.Name, p.Category, p.Price())
	if err := p.SetPrice(-5.0); err != nil {
		fmt.Println(err)
	}

	p = store.NewProduct("Sneaker", "Apparel", 99.0)
	fmt.Println(p.Name, p.Category, cfmt.ToCurrency(p.Price()))

	color.Green("Name:" + p.Name)
}
