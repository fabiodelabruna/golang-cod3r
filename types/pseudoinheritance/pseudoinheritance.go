package main

import "fmt"

type car struct {
	name        string
	actualSpeed int
}

type ferrari struct {
	car     // anonymous fields
	nitroON bool
}

func main() {
	f := ferrari{}
	f.name = "F40"
	f.actualSpeed = 0
	f.nitroON = true

	fmt.Printf("The car %s is with nitro modo on? %v\n", f.name, f.nitroON)
	fmt.Println(f)
}
