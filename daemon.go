package main

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

func startNetworkDemon() *sync.WaitGroup {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
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
			connectToService()
			fmt.Printf("Accept connection from %v\n", conn.RemoteAddr())
			conn.Close()
		}
	}()
	return &wg
}
