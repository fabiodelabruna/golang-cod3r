package main

import "fmt"

func buy(work1, work2 bool) (bool, bool, bool) {
	buy50InchesTV := work1 && work2
	buy32InchesTV := work1 != work2 // XOR
	buyIceCream := work1 || work2

	return buy50InchesTV, buy32InchesTV, buyIceCream
}

func main() {
	TV50, TV32, iceCream := buy(true, true)
	fmt.Printf("TV50: %t, TV32: %t, IceCream: %t, Healthy: %t", TV50, TV32, iceCream, !iceCream)
}
