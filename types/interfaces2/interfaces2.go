package main

import "fmt"

type sporting interface {
	turnOnNitro()
}

type ferrari struct {
	model   string
	nitroON bool
	speed   int
}

func (f *ferrari) turnOnNitro() {
	f.nitroON = true
}

func main() {
	car1 := ferrari{"F40", false, 0}
	car1.turnOnNitro()

	var car2 sporting = &ferrari{"F40", false, 0}
	car2.turnOnNitro()

	fmt.Println(car1, car2)
}
