package main

import "fmt"

func printResult(grade float32) {
	if grade >= 7 {
		fmt.Println("Approved with grade", grade)
	} else {
		fmt.Println("Disapproved with grade", grade)
	}
}

func main() {
	printResult(7.3)
	printResult(4.1)
}
