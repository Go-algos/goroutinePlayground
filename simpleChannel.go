package main

import (
	"fmt"
	"strconv"
)

func simpleChannel() {
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

	// Type conversion8
	for i := 0; i < msgCount.(int); i++ {
		msg := <-queue
		fmt.Println("Msg:", msg)
	}
	//wg.Wait()
	fmt.Println("All messages extracted ...")
}
