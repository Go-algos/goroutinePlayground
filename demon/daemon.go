package demon

import (
	"fmt"
	"log"
	"net"
	"sync"
	"time"
)

// Simulate connection to the remote service
func connectToService() interface{} {
	time.Sleep(1 * time.Second)
	return struct{}{}
}

func warmServiceConnCache() *sync.Pool {
	p := &sync.Pool{
		New: connectToService,
	}
	for i := 0; i < 10; i++ {
		p.Put(p.New())
	}
	return p
}

func StartNetworkDemon() *sync.WaitGroup {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		connPool := warmServiceConnCache()

		server, err := net.Listen("tcp", ":8080")
		if err != nil {
			log.Fatalf("Cannot listen %v", err)
		}
		defer server.Close()
		wg.Done()

		for {
			conn, err := server.Accept()
			if err != nil {
				log.Printf("Cannot accept connection%v", err)
				continue
			}
			//connectToService()
			svcConn := connPool.Get()
			fmt.Printf("Accept connection from %v\n", svcConn)
			connPool.Put(svcConn)
			conn.Close()
		}
	}()
	return &wg
}
