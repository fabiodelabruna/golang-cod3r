package main

import "fmt"

type product struct {
	name     string
	price    float64
	discount float64
}

// method: function with receiver
func (p product) priceWithDiscount() float64 {
	return p.price * (1 - p.discount)
}

func main() {
	product1 := product{
		name:     "Pen",
		price:    1.79,
		discount: 0.05,
	}

	fmt.Println(product1, product1.priceWithDiscount())

	product2 := product{"Notebook", 1999.0, 0.10}
	fmt.Println(product2, product2.priceWithDiscount())
}
