package main

import (
	"fmt"
	"goroutinePlayground/primitives"
)

func main() {
	fmt.Println("Goroutine examples")

	// 10: read from a closed channel
	//channels.ReadFromClosed()

	// 9: simple channel
	//channels.SimpleChannel()

	// 8: netWorkDemon && pool usage to improve performance
	//demon.StartNetworkDemon()
	// 7.2 Pool
	//pool2()

	// 7.1 Pool
	primitives.Pool2()

	// 6. Deadlock
	//primitives.Deadlock() // on the second attempt after import/refactor: compiles - but deadlock!!
	// above will cause next error
	// => fatal error: all goroutines are asleep - deadlock!

	// 5: Once
	primitives.Once()

	// 4. New Cond/Mutex
	primitives.NCond()

	// 3: Mutex and RWMutex - holds r/w level lock
	//primitives.M()

	// 2: Select operator
	//var res int = primitives.Fibonacci()
	//fmt.Println(res)

	// 1: Basic example
	//primitives.Iterate()
}
