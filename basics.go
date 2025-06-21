package main

import (
	"fmt"
	"github.com/google/uuid"
	"runtime"
	"sync"
)

func iterate() {
	// Create an unbuffered channel
	c := make(chan string)
	var wg sync.WaitGroup
	salutations := []string{"Hello", "greetings", "Good day"}
	for _, salutation := range salutations {
		wg.Add(1)
		// what it should be
		go runner(&salutation, c, &wg)
		// read from a channel

	}
	// read from a channel
	for i := 0; i < len(salutations); i++ {
		id := <-c // receive from c
		fmt.Println("With id", id)
	}
	wg.Wait()
	fmt.Println("All done")
}

func runner(input *string, c chan string, wg *sync.WaitGroup) {
	id := uuid.New().String()

	defer wg.Done()

	fmt.Println("Exec_id:", id)
	fmt.Println(input)
	var s runtime.MemStats
	runtime.ReadMemStats(&s)
	fmt.Println(s)
	// Send back to main goroutine ID
	c <- id
}
