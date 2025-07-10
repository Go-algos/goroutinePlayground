package channels

import (
	"fmt"
	"time"
)

func MultiChannel() {
	start := time.Now()
	c1 := make(chan interface{})
	//close(c1)
	c2 := make(chan interface{})
	//close(c2)

	var c1Count, c2Count int
	for i := 0; i < 1000; i++ {
		select {
		case <-c1:
			c1Count++
		case <-c2:
			c2Count++
		// Option 2: default
		default:
			fmt.Printf("In default %v\n\n\n", time.Since(start))
			// Option 1 =>
			//case <-time.After(time.Second):
			//	fmt.Println("Timeout ...")
		}
	}

	fmt.Printf("Count number: %d\n - %d", c1Count, c2Count)
}
