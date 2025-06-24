package main

import (
	"io"
	"net"
	"testing"
)

func init() {
	deamonStarted := startNetworkDemon()
	deamonStarted.Wait()
}

func BenchmarkNetworkRequest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		conn, err := net.Dial("tcp", "127.0.0.1:8080")
		if err != nil {
			b.Fatalf("Cannot dial host %v", err)
		}
		if _, err = io.ReadAll(conn); err != nil {
			b.Fatalf("Cannot read %v", err)
		}
		conn.Close()
	}
}
