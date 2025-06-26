package channels

import (
	"bytes"
	"fmt"
	"os"
)

func BuffChanel() {
	// Create byte buffer
	var stdOutBuff bytes.Buffer
	defer stdOutBuff.WriteTo(os.Stdout)
	defer fmt.Printf("Now output all from the buffer\n")

	intStream := make(chan int, 4)
	go func() {
		//defer fmt.Println("intStream is closed")
		defer close(intStream)
		// Writing to buffer
		defer fmt.Println(&stdOutBuff, "Producer done")

		for i := 0; i < 17; i++ {
			// Writing to buffer
			fmt.Fprintf(&stdOutBuff, "Sending: %d\n", i)
			intStream <- i
		}
	}()

	for i := range intStream {
		// Writing to buffer
		fmt.Fprintf(&stdOutBuff, "Reciving: %d\n", i)
	}
}
