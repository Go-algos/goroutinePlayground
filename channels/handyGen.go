package channels

import (
	"fmt"
	//"math/rand"
)

func Handy() {
	repeat := func(
		done <-chan interface{},
		values ...interface{},
	) <-chan interface{} {
		valueStream := make(chan interface{})
		go func() {
			defer close(valueStream)
			for {
				for _, value := range values {
					select {
					case <-done:
						return
					case valueStream <- value:
					}
				}
			}
		}()
		return valueStream
	}

	take := func(
		done <-chan interface{},
		valueStream <-chan interface{},
		num int,
	) <-chan interface{} {
		takeStream := make(chan interface{})
		go func() {
			defer close(takeStream)
			for i := 0; i < num; i++ {
				select {
				case <-done:
					return
					// something new: here is the way to read a readonly channel
				case takeStream <- <-valueStream:
				}
			}
		}()
		return takeStream
	}

	//done := make(chan interface{})
	//defer close(done)
	//
	//for num := range take(done, repeat(done, 1), 10) {
	//	fmt.Printf("%v ", num)
	//}

	//repeatFn := func(
	//	done <-chan interface{},
	//	fn func() interface{},
	//) <-chan interface{} {
	//	valueStream := make(chan interface{})
	//	go func() {
	//		defer close(valueStream)
	//		for {
	//			select {
	//			case <-done:
	//				return
	//			case valueStream <- fn():
	//			}
	//		}
	//	}()
	//	return valueStream
	//}

	//done := make(chan interface{})
	//defer close(done)
	//
	//rand := func() interface{} { return rand.Int() }
	//for num := range take(done, repeatFn(done, rand), 10) {
	//	fmt.Println(num)
	//}

	toString := func(
		done <-chan interface{},
		valueStream <-chan interface{},
	) <-chan string {
		stringStream := make(chan string)
		go func() {
			defer close(stringStream)
			for v := range valueStream {
				select {
				case <-done:
					return
				case stringStream <- v.(string):
				}
			}
		}()
		return stringStream
	}

	done := make(chan interface{})
	defer close(done)

	var message string
	for token := range toString(done, take(done, repeat(done, "I", "am,"), 5)) {
		message += token
	}

	fmt.Printf("Message: %s", message)

}
