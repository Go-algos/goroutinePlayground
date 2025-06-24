package main

import (
	"fmt"
	"github.com/google/uuid"
	"sync"
)

func pool2() {
	var numCalcsCreated int
	myPool := &sync.Pool{
		New: func() interface{} {
			instanceID := uuid.NewString()
			numCalcsCreated += 1
			mem := make([]byte, 1024)
			fmt.Println("Create a new.", instanceID)
			return &mem
		},
	}

	// Seed the pool with 4kB
	myPool.Put(myPool.New())
	myPool.Put(myPool.New())
	myPool.Put(myPool.New())
	myPool.Put(myPool.New())

	const numWorkers = 1024 * 1024
	var wg sync.WaitGroup
	wg.Add(numWorkers)

	for i := numWorkers; i > 0; i-- {
		go func() {
			defer wg.Done()
			mem := myPool.Get().(*[]byte)
			defer myPool.Put(mem)
		}()
	}

	wg.Wait()
	fmt.Printf("%d calculation created", numCalcsCreated)

}
