package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("==", "bannana" == "bannana")
	fmt.Println("!=", 3 != 2)
	fmt.Println(">", 3 > 2)
	fmt.Println("<", 3 < 2)
	fmt.Println(">=", 3 >= 2)
	fmt.Println("<=", 3 <= 2)

	d1 := time.Unix(0, 0)
	d2 := time.Unix(0, 0)
	fmt.Println("dates", d1 == d2)
	fmt.Println("dates", d1.Equal(d2))

	type Person struct {
		Name string
	}

	p1 := Person{"Fabio"}
	p2 := Person{"Fabio"}
	fmt.Println("people", p1 == p2)

}
