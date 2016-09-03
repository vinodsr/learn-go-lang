package main

import (
	"fmt"
)

func main() {
	fmt.Println("Pointers example")
	emp := employee{
		name: "Vinod"}
	calcAge(&emp)
	emp.b = &address{
		pin: 100}
	emp.address = &address{
		pin: 695123}
	emp.print()
	emp.printAddress()
	fmt.Println(emp.age)
}

func (e employee) print() {
	fmt.Println("Hello Employee ")
}

func (a address) printAddress() {
	fmt.Println("Print Adddress")
	fmt.Println(a.pin)
}

type address struct {
	pin int
}
type employee struct {
	name string
	*address
	b   *address
	age int
}

func calcAge(e *employee) {
	e.age = 90
}
