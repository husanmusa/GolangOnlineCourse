package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"

	pb "stock/proto"
)

func main() {
	conn, err := grpc.NewClient("localhost:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewStockServiceClient(conn)
	req := &pb.StockRequest{Symbol: "AAPL"}

	stream, err := client.GetStockPrices(context.Background(), req)
	if err != nil {
		log.Fatalf("Error calling GetStockPrices: %v", err)
	}

	for {
		priceUpdate, err := stream.Recv()
		if err != nil {
			break
		}
		fmt.Printf("Stock: %s | Price: %.2f | Time: %s\n", priceUpdate.GetSymbol(), priceUpdate.GetPrice(), priceUpdate.GetTimestamp())
	}
}
