package main

import (
	"fmt"
	"net"
	"os"
)

const socketPath = "/tmp/benchmark.sock"

func handleConnection(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 1024)
	for {
		_, err := conn.Read(buf)
		if err != nil {
			break
		}
		_, _ = conn.Write([]byte("OK\n"))
	}
}

func main() {
	os.Remove(socketPath) // Ensure the socket file doesn't already exist

	listener, err := net.Listen("unix", socketPath)
	if err != nil {
		panic(err)
	}
	defer listener.Close()
	fmt.Println("Unix Socket Server listening on", socketPath)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Accept error:", err)
			continue
		}
		go handleConnection(conn)
	}
}
