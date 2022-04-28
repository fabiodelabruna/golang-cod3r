package main

import "fmt"

func avg(numbers ...float64) float64 {
	total := 0.0
	for _, n := range numbers {
		total += n
	}
	return total / float64(len(numbers))
}

func printApproveds(approveds ...string) {
	fmt.Println("Approved list:")
	for i, approved := range approveds {
		fmt.Printf("%d) %s\n", i+1, approved)
	}
}

func main() {
	fmt.Printf("Avg: %.2f\n", avg(6.4, 5.8, 8.6, 5.7, 9.9))
	approveds := []string{"Maria", "João", "Pedro"}
	printApproveds(approveds...)
}
