package primitives

import (
	"fmt"
	"sync"
)

func Deadlock() {
	var onceA, onceB sync.Once

	var initB func()

	initA := func() {
		fmt.Println("initA")
		onceB.Do(initB)
	}

	initB = func() {
		fmt.Println("initB")
		onceA.Do(initA)
	}

	onceA.Do(initA)
}
