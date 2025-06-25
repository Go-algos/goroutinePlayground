package primitives

import (
	"fmt"
	"github.com/google/uuid"
	"sync"
)

type ArbitraryResource struct {
	id string
}

func Pool() {
	myPool := sync.Pool{
		New: func() interface{} {
			instance := ArbitraryResource{id: uuid.NewString()}
			fmt.Println("Create a new instance.", instance)
			return instance
		},
	}

	// We are operating with two instances in total
	myPool.Get()
	instanceA := myPool.Get()
	myPool.Put(instanceA)
	instanceB := myPool.Get()
	myPool.Put(instanceB)
	myPool.Get()
}
