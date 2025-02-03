package main

import (
	"fmt"
	"net"
	"time"
)

const socketPath = "/tmp/benchmark.sock"

func main() {
	conn, err := net.Dial("unix", socketPath)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	requests := 100000
	start := time.Now()

	for i := 0; i < requests; i++ {
		_, err := conn.Write([]byte("PING\n"))
		if err != nil {
			panic(err)
		}
		buf := make([]byte, 1024)
		_, err = conn.Read(buf)
		if err != nil {
			panic(err)
		}
	}

	duration := time.Since(start)
	fmt.Printf("Unix Socket Throughput: %d requests in %v (%.2f req/sec)\n",
		requests, duration, float64(requests)/duration.Seconds())
}
