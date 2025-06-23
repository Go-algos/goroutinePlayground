package main

import "fmt"

func main() {
	fmt.Println("Goroutine examples")

	// 6. Deadlock
	//deadlock()
	// above will cause next error
	// => fatal error: all goroutines are asleep - deadlock!

	// 5: Once
	once()

	// 4. New Cond/Mutex
	//nCond()

	// 3: Mutex and RWMutex - holds r/w level lock
	//m()

	// 2: Select operator
	//var res int = fibonacci()
	//fmt.Println(res)

	// 1: Basic example
	//iterate()
}
