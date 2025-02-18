package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"

	pb "chat/proto"
)

func main() {
	conn, err := grpc.NewClient("localhost:50053", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewChatServiceClient(conn)
	stream, err := client.Chat(context.Background())
	if err != nil {
		log.Fatalf("Error creating stream: %v", err)
	}

	go func() {
		for {
			resp, err := stream.Recv()
			if err != nil {
				log.Fatalf("Error receiving message: %v", err)
			}
			fmt.Printf("[%s] %s: %s\n", resp.GetTimestamp(), resp.GetUser(), resp.GetMessage())
		}
	}()

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter messages (Ctrl+C to quit):")
	for scanner.Scan() {
		msg := scanner.Text()
		err = stream.Send(&pb.ChatMessage{
			User:      "Client",
			Message:   msg,
			Timestamp: time.Now().Format(time.RFC3339),
		})
		if err != nil {
			return
		}
	}
}
