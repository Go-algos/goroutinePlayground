package main

import (
	"fmt"
	"sync"
	"time"
)

func nCond() {
	c := sync.NewCond(&sync.Mutex{})
	queue := make([]interface{}, 0, 10)

	removeFromTheQueue := func(delay time.Duration) {
		time.Sleep(delay)
		c.L.Lock()
		queue = queue[1:]
		fmt.Println("Removed from the queue", queue)
		c.L.Unlock()
		c.Signal()
	}

	for i := 0; i < 10; i++ {
		c.L.Lock()
		for len(queue) == 2 {
			c.Wait()
		}
		fmt.Println("Adding to the queue")
		queue = append(queue, struct{}{})
		fmt.Println("State", queue)

		go removeFromTheQueue(2 * time.Second)
		c.L.Unlock()
	}
}
