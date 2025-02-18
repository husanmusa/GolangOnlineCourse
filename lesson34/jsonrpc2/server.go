package main

import (
	"net"

	"net/rpc"

	"net/rpc/jsonrpc"

	"log"
)

// BlockchainNode represents a monitored node

type BlockchainNode struct {
	peersConnected int
	blockHeight    int64
}

// NodeStatusResponse matches Ethereum client JSON-RPC format

type NodeStatusResponse struct {
	Peers       int   `json:"peers"`
	BlockHeight int64 `json:"latestBlock"`
	SyncStatus  bool  `json:"isSyncing"`
}

// RPC exposed method

func (n *BlockchainNode) GetStatus(_ struct{}, reply *NodeStatusResponse) error {
	*reply = NodeStatusResponse{
		Peers:       n.peersConnected,
		BlockHeight: n.blockHeight,
		SyncStatus:  n.blockHeight < getNetworkHeight(),
	}

	n.peersConnected++ // Simulate peer changes

	n.blockHeight += 4 // Simulate block mining

	return nil

}

func getNetworkHeight() int64 { return 19876543 } // Mock external data

func main() {

	rpc.Register(&BlockchainNode{peersConnected: 12, blockHeight: 19876400})

	l, err := net.Listen("tcp", ":8545") // Standard Ethereum JSON-RPC port

	if err != nil {

		log.Fatal("listen error:", err)

	}

	defer l.Close()

	log.Println("JSON-RPC server running on port 8545 (ETH mainnet standard)")

	for {

		conn, err := l.Accept()

		if err != nil {

			log.Printf("accept error: %v", err)

			continue

		}

		go jsonrpc.ServeConn(conn) // Concurrent request handling

	}

}
