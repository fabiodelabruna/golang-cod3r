package main

import (
	"fmt"
	"strings"
)

type person struct {
	name    string
	surname string
}

func (p person) getFullName() string {
	return p.name + " " + p.surname
}

func (p *person) setFullName(fullName string) {
	parts := strings.Split(fullName, " ")
	p.name = parts[0]
	p.surname = parts[1]
}

func main() {
	p1 := person{"Peter", "Jordan"}
	fmt.Println(p1.getFullName())

	p1.setFullName("Davy Jones")
	fmt.Println(p1.getFullName())
}
