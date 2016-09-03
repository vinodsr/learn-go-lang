package main

import "fmt"

func main() {
	var name string

	fmt.Println("Enter the name")
	fmt.Scanf("%s", &name)
	fmt.Printf("Hello " + name)
	fmt.Println(eat())

}

func eat() (bool, int) {
	return true, 10
}
