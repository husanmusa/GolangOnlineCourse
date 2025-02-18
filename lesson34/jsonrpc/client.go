package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc/jsonrpc"
)

// BotArgs is used to hold the command and message sent to the bot.

type BotArgs struct {
	Command string
	Message string
}

// BotReply contains the bot's response.

type BotReply struct {
	Response string
}

func main() {

	// Dial (connect) to the JSON-RPC bot server at localhost on port 1234.
	conn, err := net.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("Connection error:", err)
	}
	defer conn.Close()

	// Create a new JSON-RPC client using the established connection.
	client := jsonrpc.NewClient(conn)

	// Example 1: "greet" command.
	args := &BotArgs{
		Command: "greet",
		Message: "Alice",
	}

	var reply BotReply

	err = client.Call("BotService.Respond", args, &reply)

	if err != nil {
		log.Fatal("Error calling BotService.Respond:", err)
	}

	fmt.Println("Greet Response:", reply.Response)

	// Example 2: "echo" command.
	args = &BotArgs{
		Command: "echo",
		Message: "This is a test message",
	}

	err = client.Call("BotService.Respond", args, &reply)

	if err != nil {
		log.Fatal("Error calling BotService.Respond:", err)
	}

	fmt.Println("Echo Response:", reply.Response)

	// Example 3: Unknown command.

	args = &BotArgs{
		Command: "unknown",
		Message: "Some text",
	}

	err = client.Call("BotService.Respond", args, &reply)

	if err != nil {
		log.Fatal("Error calling BotService.Respond:", err)
	}

	fmt.Println("Unknown Command Response:", reply.Response)
}
