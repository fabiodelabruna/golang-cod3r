package main

import (
	"fmt"
	"time"
)

func gradeResult(n float32) string {
	var grade = int(n)

	switch grade {
	case 10:
		fallthrough
	case 9:
		return "A"
	case 8, 7:
		return "B"
	case 6, 5:
		return "C"
	case 4, 3:
		return "D"
	case 2, 1, 0:
		return "E"
	default:
		return "Invalid grade"
	}
}

func getType(i interface{}) string {
	switch i.(type) {
	case int:
		return "integer"
	case float32, float64:
		return "float"
	case string:
		return "string"
	case func():
		return "function"
	default:
		return "unknown"
	}
}

func main() {
	fmt.Println(gradeResult(9.8))
	fmt.Println(gradeResult(6.9))
	fmt.Println(gradeResult(2.1))
	fmt.Println(gradeResult(11))

	t := time.Now()
	switch { // switch true
	case t.Hour() < 12:
		fmt.Println("Good morning")
	case t.Hour() < 18:
		fmt.Println("Good afternoon")
	default:
		fmt.Println("Good night")
	}

	fmt.Println(getType(2.3))
	fmt.Println(getType(1))
	fmt.Println(getType("oops!"))
	fmt.Println(getType(func() {}))
	fmt.Println(getType(time.Now()))

}
