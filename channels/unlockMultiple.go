package channels

import (
	"fmt"
	"sync"
)

func Multiple() {
	begin := make(chan interface{})
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		//time.Sleep(1 * time.Second)
		go func(i int) {
			defer wg.Done()
			<-begin
			fmt.Printf("%v has begun\n", i)
		}(i)
	}

	fmt.Println("Unblocking goroutine")
	close(begin)
	wg.Wait()
}
