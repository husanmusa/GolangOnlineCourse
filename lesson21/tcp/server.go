package main

import (
	"fmt"
	"net"
)

func handleClient(conn net.Conn) {
	defer conn.Close()

	// Read data from client
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading:", err)
		return
	}
	// Print received message
	fmt.Printf("Received message: %s\n", string(buffer[:n]))

	// Respond to client
	conn.Write([]byte("Message received\n"))
}

func main() {
	// Start server on port 8080
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server listening on port 8080")

	// Accept incoming connections
	for {
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("Error accepting:", err)
		return
	}
	// Handle each connection concurrently
	go handleClient(conn)
	// handleClient(conn)
	}
}
