package channels

import (
	"fmt"
	"sync"
	"time"
)

func Ranging() {
	// TODO: try to use buffered channel!!!
	// => make(chan int, N)
	wg := sync.WaitGroup{}
	wg.Add(1)
	intStream := make(chan int, 15)
	go func() {
		defer wg.Done()
		defer close(intStream)
		for i := 1; i <= 10; i++ {
			fmt.Printf("Push item into the stream: %d\n", i)
			intStream <- i
			time.Sleep(1 * time.Second)
		}
	}()

	wg.Wait()
	fmt.Println("Unblocking ... \nRead from the stream")
	for item := range intStream {
		fmt.Printf("Read item: %d\n", item)
	}
}
