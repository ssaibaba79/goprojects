package main

import "fmt"

type calculateDiscountPrice func(float32) float32

var giveway = false

// function closure
func storeDiscount(threshold, discount float32) calculateDiscountPrice {
	return func(price float32) float32 {

		if giveway {
			return 0
		} else if price > threshold {
			return price + (price - discount)
		}
		return price
	}
}

// function closure - forced early evaluation - using local variable
func storeDiscount2(threshold, discount float32) calculateDiscountPrice {
	gw := giveway
	return func(price float32) float32 {
		if gw {
			return 0
		} else if price > threshold {
			return price + (price - discount)
		}
		return price
	}
}

// function closure - forced early evaluation - using function argument
func storeDiscount3(threshold, discount float32, giveway bool) calculateDiscountPrice {
	return func(price float32) float32 {
		if giveway {
			return 0
		} else if price > threshold {
			return price + (price - discount)
		}
		return price
	}
}

type Item struct {
	Name  string
	Price float32
}

var dairy = map[string]Item{
	"milk":  {"Milk", 4.99},
	"egg":   {"Egg", 9.99},
	"bread": {"Bread", 4.28},
}

var veggies = map[string]Item{
	"carrot":  {"Carrots", 3.99},
	"spinach": {"Spinach", 2.99},
	"tomato":  {"Tomato", 1.28},
}

// method using the closure
func calculateDiscounts(items map[string]Item, calc calculateDiscountPrice) {
	for k, v := range items {
		fmt.Println("discount price for", k, calc(v.Price))
	}
}

// No early evaluation
func ClosureExample() {
	fmt.Println("ClosureExample")
	dairyDiscountCalc := storeDiscount(2.00, 0.25)
	veggiesDiscountCalc := storeDiscount(2.0, 0.15)

	calculateDiscounts(dairy, dairyDiscountCalc)
	calculateDiscounts(veggies, veggiesDiscountCalc)

	//close on value changed after evaluation of dairyDiscountCalc and dairyDiscountCalc
	giveway = true

	calculateDiscounts(dairy, dairyDiscountCalc)
	calculateDiscounts(veggies, veggiesDiscountCalc)
}

// Forced early evaluation using local variable
func ClosureExample2() {
	fmt.Println("ClosureExample2")
	giveway = false

	dairyDiscountCalc2 := storeDiscount2(2.0, 0.35)
	veggiesDiscountCalc2 := storeDiscount2(2.0, 0.18)

	calculateDiscounts(dairy, dairyDiscountCalc2)
	calculateDiscounts(veggies, veggiesDiscountCalc2)

	//closed on value changed after evaluation of dairyDiscountCalc and dairyDiscountCalc
	giveway = true

	calculateDiscounts(dairy, dairyDiscountCalc2)
	calculateDiscounts(veggies, veggiesDiscountCalc2)
}

// Forced early evaluation using function parameter
func ClosureExample3() {
	fmt.Println("ClosureExample3")
	giveway = false

	dairyDiscountCalc3 := storeDiscount3(2.0, 0.25, giveway)
	veggiesDiscountCalc3 := storeDiscount3(2.0, 0.15, giveway)

	calculateDiscounts(dairy, dairyDiscountCalc3)
	calculateDiscounts(veggies, veggiesDiscountCalc3)

	//closed on value changed after evaluation of dairyDiscountCalc and dairyDiscountCalc
	giveway = true

	calculateDiscounts(dairy, dairyDiscountCalc3)
	calculateDiscounts(veggies, veggiesDiscountCalc3)
}

func main() {
	ClosureExample()
	ClosureExample2()
	ClosureExample3()

}
