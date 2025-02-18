package main

import (
	"fmt"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"net"

	pb "chat/proto"
)

type server struct {
	pb.UnimplementedChatServiceServer
}

func (s *server) Chat(stream grpc.BidiStreamingServer[pb.ChatMessage, pb.ChatMessage]) error {
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error receiving message: %v", err)
		}

		fmt.Printf("[%s] %s: %s\n", msg.GetTimestamp(), msg.GetUser(), msg.GetMessage())

		response := &pb.ChatMessage{
			User:      "Server",
			Message:   fmt.Sprintf("Received: %s", msg.GetMessage()),
			Timestamp: time.Now().Format(time.RFC3339),
		}

		if err := stream.Send(response); err != nil {
			log.Fatalf("Error sending response: %v", err)
		}
	}
}

func main() {
	lis, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterChatServiceServer(s, &server{})
	reflection.Register(s)

	fmt.Println("Chat server running on port 50053...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
