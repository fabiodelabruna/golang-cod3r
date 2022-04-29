package main

import "fmt"

type grade float64

func (n grade) between(start, end float64) bool {
	return float64(n) >= start && float64(n) <= end
}

func gradeToConcept(n grade) string {
	if n.between(9.0, 10) {
		return "A"
	} else if n.between(7.0, 8.99) {
		return "B"
	} else if n.between(5.0, 6.99) {
		return "C"
	} else if n.between(3.0, 4.99) {
		return "D"
	} else {
		return "E"
	}
}

func main() {
	fmt.Println(gradeToConcept(9.8))
	fmt.Println(gradeToConcept(6.4))
	fmt.Println(gradeToConcept(3.2))
}
