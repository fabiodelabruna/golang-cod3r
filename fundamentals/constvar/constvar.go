package main

import (
	"fmt"
	"math"
)

func main() {
	const PI float64 = 3.1415
	var radius = 3.2 // float type inferred by compiler

	area := PI * math.Pow(radius, 2)
	fmt.Println("The area of the circumference is:", area)

	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 4
	)

	fmt.Println(a, b, c, d)

	var e, f bool = true, false
	fmt.Println(e, f)

	g, h, i := 2, false, "ooops!"
	fmt.Println(g, h, i)

}
