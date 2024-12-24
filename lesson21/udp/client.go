package main

import (
    "fmt"
    "net"
)

func main() {
    // Resolve UDP server address
    serverAddr, err := net.ResolveUDPAddr("udp", "localhost:8080")
    if err != nil {
        fmt.Println("Error resolving address:", err)
        return
    }

    // Create UDP connection
    conn, err := net.DialUDP("udp", nil, serverAddr)
    if err != nil {
        fmt.Println("Error dialing:", err)
        return
    }
    defer conn.Close()

    // Send message to server
    message := []byte("Hello, server!")
    _, err = conn.Write(message)
    if err != nil {
        fmt.Println("Error writing:", err)
        return
    }

    // // Read response from server
    buffer := make([]byte, 1024)
    n, _, err := conn.ReadFromUDP(buffer)
    if err != nil {
        fmt.Println("Error reading response:", err)
        return
    }

    // Print response from server
    fmt.Printf("Response from server: %s\n", string(buffer[:n]))
}
