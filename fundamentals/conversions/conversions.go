package main

import (
	"fmt"
	"strconv"
)

func main() {

	x := 2.4
	y := 2
	fmt.Println(x / float64(y))

	grade := 6.9
	finalGrade := int(grade)
	fmt.Println(finalGrade)

	// atention...
	fmt.Println("Teste " + string(97))

	// int to string
	fmt.Println("Teste " + strconv.Itoa(123))

	// string to int
	num, _ := strconv.Atoi("123")
	fmt.Println(num - 122)

	// string to bool
	b, _ := strconv.ParseBool("true") // accepts 0 and 1
	if b {
		fmt.Println("TRUE")
	}
}
