package primitives

// Ref: https://en.wikipedia.org/wiki/Fibonacci_sequence
/*
	The select statement lets a goroutine wait on multiple communication operations.
	A select blocks until one of its cases can run, then it executes that case.
	It chooses one at random if multiple is ready
*/
func fibonacciAcc(c, quit chan int) int {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			//fmt.Println("quit")
			return y
		}
	}
}

func Fibonacci() int {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			//fmt.Println(<-c)
			// Pushing N signals into C channel;
			<-c
		}
		// Done: exiting
		quit <- 0
	}()
	return fibonacciAcc(c, quit)
}
