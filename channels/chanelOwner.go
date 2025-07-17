package channels

import (
	"fmt"
	"math/rand"
	"time"
)

/*
Wrap channel owner and consumption in the lexical confinement
Note: '<-chan' is a readonly channel
*/
func RunOwner() {
	fmt.Println("Some example ..")

	chanOwner := func() <-chan int {
		results := make(chan int, 5)
		// Some processing function
		go func() {
			defer close(results)
			for i := 0; i < 5; i++ {
				randomDelay := rand.Intn(5)
				fmt.Printf("Delay %d seconds ...\n", randomDelay)
				time.Sleep(time.Duration(randomDelay) * time.Second)
				results <- rand.Intn(100)
			}
		}()
		return results
	}

	consumer := func(results <-chan int) {
		for result := range results {
			fmt.Printf("Result: %d\n", result)
		}
		fmt.Println("Done with receiving ...")
	}

	results := chanOwner()
	// results <- 5 This would fail as we have a read-only channel returned
	consumer(results)
}
