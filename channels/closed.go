package channels

import "fmt"

func ReadFromClosed() {
	intStream := make(chan int)
	close(intStream)

	integer, ok := <-intStream
	fmt.Printf("Status: (%v): %v", integer, ok)
}
