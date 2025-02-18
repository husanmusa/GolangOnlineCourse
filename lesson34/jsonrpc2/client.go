package main

import (
	"log"
	"net"
	"net/rpc/jsonrpc"
	"os"
	"time"
)

type StatusRequest struct{} // Empty request params

func main() {

	conn, err := net.Dial("tcp", "localhost:8545")
	if err != nil {
		log.Fatal("dial error:", err)
	}

	defer conn.Close()

	client := jsonrpc.NewClient(conn)

	for range time.Tick(15 * time.Second) { // Poll every 15s
		var status struct {
			Peers       int
			BlockHeight int64
			SyncStatus  bool
		}

		err := client.Call("BlockchainNode.GetStatus", StatusRequest{}, &status)

		if err != nil {

			log.Printf("RPC error: %v", err)

			continue

		}

		log.Printf("Node status: %d peers | Block %d | Syncing: %v",
			status.Peers, status.BlockHeight, status.SyncStatus)

		if status.BlockHeight > 19876500 {
			log.Println("Node synchronized with network")
			os.Exit(0)
		}
	}
}
