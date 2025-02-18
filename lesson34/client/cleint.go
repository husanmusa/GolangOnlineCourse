package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Args struct {
	A, B int
}

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234") // Connect to the RPC server
	if err != nil {
		log.Fatalf("Error dialing: %v", err)
	}

	args := Args{A: 10, B: 2}
	var reply int

	// Call the Multiply method
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatalf("Error calling Arith.Multiply: %v", err)
	}
	fmt.Printf("Arith.Multiply: %d*%d=%d\\n", args.A, args.B, reply)

	// Call the Divide method
	err = client.Call("Arith.Divide", args, &reply)
	if err != nil {
		log.Fatalf("Error calling Arith.Divide: %v", err)
	}
	fmt.Printf("Arith.Divide: %d/%d=%d\\n", args.A, args.B, reply)
}
