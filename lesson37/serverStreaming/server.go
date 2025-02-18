package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"net"

	pb "stock/proto"
)

type server struct {
	pb.UnimplementedStockServiceServer
}

func (s *server) GetStockPrices(req *pb.StockRequest, stream pb.StockService_GetStockPricesServer) error {
	symbol := req.GetSymbol()
	for i := 0; i < 10; i++ { // Simulate 10 updates
		price := rand.Float64()*100 + 100 // Random price between 100-200
		resp := &pb.StockPrice{
			Symbol:    symbol,
			Price:     price,
			Timestamp: time.Now().Format(time.RFC3339),
		}
		if err := stream.Send(resp); err != nil {
			return err
		}
		time.Sleep(2 * time.Second) // Simulate real-time delay
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterStockServiceServer(s, &server{})
	reflection.Register(s)

	fmt.Println("Stock server running on port 50052...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
