package main

import "fmt"

func inc1(n int) {
	n++
}

func inc2(n *int) {
	*n++
}

func main() {
	n := 1

	inc1(n)
	fmt.Println("by value:", n)

	inc2(&n)
	fmt.Println("by reference:", n)
}
