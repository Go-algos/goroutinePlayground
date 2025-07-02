package channels

import (
	"fmt"
	"time"
)

func Start() {
	start := time.Now()
	c := make(chan interface{})
	go func() {
		fmt.Println("Start waiting for 5 seconds...")
		time.Sleep(time.Second * 5)
		close(c)
	}()

	fmt.Println("Blocking on read ...")
	select {
	case <-c:
		fmt.Printf("Unblocked %v later.\n", time.Since(start))
	}
}
