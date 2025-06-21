package main

import (
	"fmt"
	"github.com/google/uuid"
	"runtime"
	"sync"
)

func iterate() {
	var wg sync.WaitGroup
	salutations := []string{"Hello", "greetings", "Good day"}
	for _, salutation := range salutations {
		wg.Add(1)
		// what it should be
		go runner(&salutation, &wg)
	}

	wg.Wait()
	fmt.Println("All done")
}

func runner(input *string, wg *sync.WaitGroup) {
	defer wg.Done()

	id := uuid.New().String()
	fmt.Println("Exec_id:", id)
	fmt.Println(input)
	var s runtime.MemStats
	runtime.ReadMemStats(&s)
	fmt.Println(s)
}
