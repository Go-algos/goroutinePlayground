package main

import (
	"fmt"
	"github.com/google/uuid"
	"sync"
)

func iterate() {
	var wg sync.WaitGroup
	salutations := []string{"Hello", "greetings", "Good day"}
	for _, salutation := range salutations {
		wg.Add(1)
		go func() {
			id := uuid.New().String()
			fmt.Println("Exec_id:", id)
			defer wg.Done()
			fmt.Println(salutation)
		}()
	}

	wg.Wait()
	fmt.Println("All done")
}
