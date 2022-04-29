package main

import "fmt"

type item struct {
	productID int
	amount    int
	price     float64
}

type request struct {
	userID int
	items  []item
}

func (r request) totalValue() float64 {
	total := 0.0
	for _, item := range r.items {
		total += item.price * float64(item.amount)
	}
	return total
}

func main() {
	request := request{
		userID: 1,
		items: []item{
			{1, 2, 12.10},
			{2, 1, 23.49},
			{11, 100, 3.13},
		},
	}
	fmt.Printf("Request value: %.2f", request.totalValue())
}
