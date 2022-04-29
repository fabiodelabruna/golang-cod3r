package main

import "fmt"

func getApprovedValue(number int) int {
	defer fmt.Println("END")

	if number >= 5000 {
		fmt.Println("Value is greater than 5000")
		return 5000
	}

	fmt.Println("Value is lower than 5000")
	return number
}

func main() {
	fmt.Println(getApprovedValue(6000))
}
