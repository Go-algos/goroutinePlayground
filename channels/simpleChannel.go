package channels

import (
	"fmt"
	"reflect"
	"strconv"
)

// https://stackoverflow.com/questions/20170275/how-to-find-the-type-of-an-object-in-go
// Type detection using a switch statement
func typeof(v interface{}) string {
	switch v.(type) {
	case int:
		return "int"
	case float64:
		return "float64"
	//... etc
	default:
		return "unknown"
	}
}

func SimpleChannel() {
	//var wg sync.WaitGroup
	queue := make(chan interface{})
	//wg.Add(1)

	go func() {
		//defer wg.Done()
		dynamicLen := 10
		queue <- dynamicLen
		for i := 0; i < dynamicLen; i++ {
			queue <- "Message " + strconv.Itoa(i)
		}
	}()

	// Extract fist element -> N - number of messages
	msgCount := <-queue

	// Detect type
	fmt.Println(reflect.TypeOf(msgCount))
	fmt.Println(typeof(msgCount))

	// Type conversion8
	for i := 0; i < msgCount.(int); i++ {
		msg := <-queue
		fmt.Println("Msg:", msg)
	}
	//wg.Wait()
	fmt.Println("All messages extracted ...")
}
