package main

import (
	"fmt"
)

func main() {
	fmt.Println("Pointers example")
	emp := employee{
		name: "Vinod"}
	calcAge(&emp)
	emp.print()
	fmt.Println(emp.age)
}

func (e employee) print() {
	fmt.Println("Hello Employee ")
}

type employee struct {
	name string
	age  int
}

func calcAge(e *employee) {
	e.age = 90
}
