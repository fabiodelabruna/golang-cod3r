package main

import "fmt"

type sporting interface {
	turnOnNitro()
}

type luxurious interface {
	park()
}

type sportingLuxurious interface {
	sporting
	luxurious
	// we can add more methods
}

type bmw7 struct {
}

func (b bmw7) turnOnNitro() {
	fmt.Println("Nitro...")
}

func (b bmw7) park() {
	fmt.Println("Park...")
}

func main() {
	var b sportingLuxurious = bmw7{}
	b.turnOnNitro()
	b.park()
}
