package main

import (
	"fmt"
	"sync"
	"time"
)

// Mutext example
func m() {

	var count int
	var lock sync.Mutex

	increment := func() {
		lock.Lock()
		defer lock.Unlock()
		count++
		time.Sleep(300 * time.Millisecond)
		fmt.Printf("Incrementing %d\n", count)
	}

	decrement := func() {
		lock.Lock()
		defer lock.Unlock()
		count--
		time.Sleep(300 * time.Millisecond)
		fmt.Printf("Decrementing %d\n", count)
	}

	var arithmetic sync.WaitGroup
	for i := 0; i < 10; i++ {
		arithmetic.Add(1)
		go func() {
			defer arithmetic.Done()
			increment()
		}()
	}

	for i := 0; i < 10; i++ {
		arithmetic.Add(1)
		go func() {
			defer arithmetic.Done()
			decrement()
		}()
	}

	arithmetic.Wait()
	fmt.Println("Final result:", count)
}
