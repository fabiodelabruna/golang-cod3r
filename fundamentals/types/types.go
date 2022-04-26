package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {

	// integer
	fmt.Println(1, 2, 1000)
	fmt.Println("integer literal is", reflect.TypeOf(32000))

	// uint8  - 1 byte
	// uint16 - 2 bytes
	// uint32 - 4 bytes
	// uint64 - 8 bytes

	// byte is a alias for uint8
	var b byte = 255
	fmt.Println("byte is", reflect.TypeOf(b))

	// int8, int16, int32, int64
	i1 := math.MaxInt64
	fmt.Println("max value of int is", i1)
	fmt.Println("type of i1 is", reflect.TypeOf(i1))

	// rune is a alias for uint32
	// represents a integer on unicode table
	var i2 rune = 'a'
	fmt.Println("rune is", reflect.TypeOf(i2))
	fmt.Println(i2)

	// floating point
	var x float32 = 49.99
	x = float32(79.99)
	fmt.Println("type of x is", reflect.TypeOf(x))
	fmt.Println("floating point literal is", reflect.TypeOf(49.99)) // float64

	// boolean
	bo := true
	fmt.Println("type if bo is", reflect.TypeOf(bo))
	fmt.Println(!bo)

	// string
	s1 := "Hi, my name is Fabio"
	fmt.Println(s1 + "!")
	fmt.Println("the length of the string is", len(s1))

	// multiple lines strings
	s2 := `Hi
	my
	name
	is
	Fabio`
	fmt.Println("the length of the string is", len(s2))

	// char ???
	char := 'a'
	fmt.Println("type if char is", reflect.TypeOf(char))
	fmt.Println(char)

}
