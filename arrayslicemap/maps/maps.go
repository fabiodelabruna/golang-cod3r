package main

import "fmt"

func main() {

	// maps must be initialized
	// var approveds map[int]string
	approveds := make(map[int]string)

	approveds[12345678900] = "Maria"
	approveds[12345678911] = "Pedro"
	approveds[12345678922] = "Ana"
	fmt.Println(approveds)

	for cpf, name := range approveds {
		fmt.Printf("%s (CPF: %d)\n", name, cpf)
	}

	fmt.Println(approveds[12345678911])
	delete(approveds, 12345678911)
	fmt.Println(approveds[12345678911])

	employeesAndCompensation := map[string]float64{
		"José João":      11323.42,
		"Gabriela Silva": 15235.55,
		"Pedro Junior":   2342.00,
	}

	employeesAndCompensation["Rafael Filho"] = 12322.00
	delete(employeesAndCompensation, "Unknown")

	for name, compensation := range employeesAndCompensation {
		fmt.Println(name, compensation)
	}

	employeesByLetter := map[string]map[string]float64{
		"G": {
			"Gabriela Silva": 15235.55,
			"Guga Pereira":   11234.22,
		},
		"J": {
			"José João": 11323.42,
		},
		"P": {
			"Pedro Junior": 2342.00,
		},
	}

	delete(employeesByLetter, "P")

	for letter, employees := range employeesByLetter {
		fmt.Println(letter, employees)
	}

}
