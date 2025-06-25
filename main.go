package main

import (
	"fmt"
	"goroutinePlayground/channels"
)

func main() {
	fmt.Println("Goroutine examples")

	// 10: read from a closed channel
	channels.ReadFromClosed()

	// 9: simple channel
	//channels.SimpleChannel()

	// 8: netWorkDemon && pool usage to improve performance

	// 7.2 Pool
	//pool2()

	// 7.1 Pool
	//pool()

	// 6. Deadlock
	//deadlock()
	// above will cause next error
	// => fatal error: all goroutines are asleep - deadlock!

	// 5: Once
	//once()

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
