package channels

import (
	"fmt"
	"time"
)

func ProgressRunner() {
	done := make(chan interface{})
	go func() {
		time.Sleep(time.Second * 5)
		close(done)
	}()

	workCount := 0

	//NonBlocking call
loop:
	for {
		select {
		case <-done:
			break loop
		default:
		}
		workCount++
		time.Sleep(time.Second * 1)
		fmt.Printf("ProgressOnWork: %d\n", workCount)
	}

	fmt.Println("Goroutine progress work count:", workCount)
}
