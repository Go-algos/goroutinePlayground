package main

import "fmt"

func main() {
	fmt.Println("Goroutine examples")
	// 2: Select operator
	var res int = fibonacci()
	fmt.Println(res)

	// 1: Basic example
	//iterate()
}
