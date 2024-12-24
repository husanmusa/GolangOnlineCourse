package main

import (
    "fmt"
    "net"
)

func main() {
    // Listen on UDP port 8080
    addr, err := net.ResolveUDPAddr("udp", ":8080")
    if err != nil {
        fmt.Println("Error resolving address:", err)
        return
    }

    // Create UDP connection
    conn, err := net.ListenUDP("udp", addr)
    if err != nil {
        fmt.Println("Error listening:", err)
        return
    }
    defer conn.Close()

    fmt.Println("Server listening on port 8080")

    // Buffer for incoming data
    buffer := make([]byte, 1024)

    for {
        // Read from connection
        n, addr, err := conn.ReadFromUDP(buffer)
        if err != nil {
            fmt.Println("Error reading:", err)
            continue
        }

        // Print received message
        fmt.Printf("Received message from %s: %s\n", addr.String(), string(buffer[:n]))

        // Echo back to client
        _, err = conn.WriteToUDP([]byte("Message received"), addr)
        if err != nil {
            fmt.Println("Error writing:", err)
        }
    }
}
