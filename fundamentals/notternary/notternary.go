package main

import "fmt"

// Golang does not have ternary operator
func getResult(grade float32) string {
	if grade >= 6 {
		return "Approved"
	}
	return "Disapproved"
}

func main() {
	fmt.Println(getResult(6.2))
}
