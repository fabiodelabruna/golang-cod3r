package main

import "fmt"

func change(p1, p2 int) (second, first int) {
	second, first = p2, p1
	return
}

func main() {
	r1, r2 := change(1, 2)
	fmt.Printf("%d %d\n", r1, r2)
}
