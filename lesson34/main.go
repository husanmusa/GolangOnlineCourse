package main

import (
	"errors"
	"log"
	"net"
	"net/rpc"
)

// Args holds the arguments passed to the RPC methods.
type Args struct {
	A, B int
}

// Arith provides arithmetic operations.
type Arith int

// Multiply is an RPC method.
func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

// Divide is an RPC method.
func (t *Arith) Divide(args *Args, reply *int) error {
	if args.B == 0 {
		return errors.New("division by zero")
	}
	*reply = args.A / args.B
	return nil
}

func main() {
	arith := new(Arith)
	rpc.Register(arith) // Register the Arith service

	listener, err := net.Listen("tcp", ":1234") // Listen on port 1234
	if err != nil {
		log.Fatalf("Error listening: %v", err)
	}
	defer listener.Close()
	log.Println("Server started on port 1234")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Error accepting: %v", err)
			continue
		}
		go rpc.ServeConn(conn) // Handle the connection
	}
}
