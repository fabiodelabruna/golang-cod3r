package main

import (
	"fmt"
	"math"
)

func main() {
	a := 3
	b := 2

	fmt.Println("SUM:", a+b)
	fmt.Println("SUBTRACTION:", a-b)
	fmt.Println("DIVISION:", a/b)
	fmt.Println("MULTIPLICATION:", a*b)
	fmt.Println("MODULE:", a%b)

	// bitwise
	fmt.Println("AND:", a&b) // 11 & 10 = 10
	fmt.Println("OR:", a|b)  // 11 | 10 = 11
	fmt.Println("XOR:", a^b) // 11 ^ 10 = 01

	c := 3.0
	d := 2.0

	// other operations using math
	fmt.Println("MAX:", math.Max(c, d))
	fmt.Println("MIN:", math.Min(c, d))
	fmt.Println("POW:", math.Pow(c, d))

}
