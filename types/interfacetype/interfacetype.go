package main

import "fmt"

type course struct {
	name string
}

func main() {
	var anything interface{}
	fmt.Println(anything)

	anything = 3
	fmt.Println(anything)

	type dynamic interface{}

	var anything2 dynamic = "Ooops"
	fmt.Println(anything2)

	anything2 = true
	fmt.Println(anything2)

	anything2 = course{"Golang course!"}
	fmt.Println(anything2)
}
