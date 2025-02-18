package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

// BotArgs indicates the structure for arguments used in bot communication.
type BotArgs struct {
	Command string
	Message string
}

// BotReply holds the response data for the bot.
type BotReply struct {
	Response string
}

// BotService provides functionalities of our bot.
type BotService struct{}

// Respond processes commands and returns an appropriate response.
// It belongs to the BotService class.

func (b *BotService) Respond(args *BotArgs, reply *BotReply) error {
	switch args.Command {
	case "greet":
		reply.Response = fmt.Sprintf("Hello, %s!", args.Message)
	case "echo":
		reply.Response = "Echo: " + args.Message
	default:
		reply.Response = "Unknown command: " + args.Command
	}
	return nil
}

func main() {

	// Register the BotService with the RPC server.
	err := rpc.Register(new(BotService))

	if err != nil {
		log.Fatal("Error registering BotService:", err)
	}

	// Start listening on TCP port 1234.
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("Listener error:", err)
	}

	fmt.Println("Bot JSON-RPC server listening on port 1234...")

	// Accept incoming connections and serve them using JSON-RPC.
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Failed to accept connection:", err)
			continue
		}

		// Serve each connection in a new goroutine.
		go jsonrpc.ServeConn(conn)
	}
}
