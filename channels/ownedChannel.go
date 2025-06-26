package channels

import "fmt"

func Owned() {
	// return channel from the function
	chanOwner := func() <-chan int {
		resultStream := make(chan int, 5)

		// Populate channel within a goroutine
		go func() {
			defer close(resultStream)
			for i := 0; i <= 10; i++ {
				resultStream <- i
			}
		}()
		return resultStream
	}

	resultStream := chanOwner()
	for res := range resultStream {
		fmt.Printf("Recieved: %d\n", res)
	}
	fmt.Println("Done receiving!")
}
