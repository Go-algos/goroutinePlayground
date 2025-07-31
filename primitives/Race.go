package primitives

import (
	"fmt"
	"sync"
)

var counter int
var mutex sync.Mutex

func increment(wg *sync.WaitGroup, id int) {
	defer wg.Done()
	mutex.Lock()
	counter++
	fmt.Printf("Counter: %d\n", counter)
	mutex.Unlock()
	if id == 1 {
		counter++
	}
}

/*
How to detect data race:
~go run -race .

Ref: race detection
https://go.dev/doc/articles/race_detector
*/
func Race() {
	var wg sync.WaitGroup
	wg.Add(4)

	go increment(&wg, 0)
	go increment(&wg, 1)
	go increment(&wg, 2)
	go increment(&wg, 4)

	wg.Wait()
	fmt.Printf("Final: %d\n", counter)
}

func Race2() {
	var wg sync.WaitGroup
	wg.Add(5)
	var i int
	for i = 0; i < 5; i++ {
		go func() {
			fmt.Println(i) // Not the 'i' you are looking for.
			wg.Done()
		}()
	}
	wg.Wait()
}
