package main

import "fmt"

func main() {
	var data [4]int
	data[0] = 100
	fmt.Println(data)

	list := [4]int{10, 20, 30, 40}
	fmt.Println(list)
	/// create a  array with max of 10 elements
	scores := make([]int, 0, 10)
	//scores[7] = 9033
	scores = append(scores, 5)
	scores = append(scores, 15)
	scores = scores[0:8]
	scores[7] = 9033
	fmt.Println(scores)
	c := cap(scores)
	fmt.Print(c)
	test()
	likes := make([]int, 10, 10)
	fmt.Println("likes : ", likes)
}

func test() {

}
