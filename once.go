package main

import (
	"fmt"
	"sync"
)

func once() {
	var count int

	increment := func() {
		count++
	}

	var once sync.Once
	var increments sync.WaitGroup
	increments.Add(100)

	for i := 0; i < 100; i++ {
		go func() {
			defer increments.Done()
			fmt.Println("Attempt to call once many times $d", i)
			once.Do(increment)
		}()
	}

	increments.Wait()
	fmt.Println("Actually called N-times =>", count)
}
