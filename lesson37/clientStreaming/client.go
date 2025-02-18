package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"google.golang.org/grpc"

	pb "fileUpload/proto"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewFileUploadServiceClient(conn)

	stream, err := client.UploadFile(context.Background())
	if err != nil {
		log.Fatalf("Failed to upload file: %v", err)
	}

	file, err := os.Open("text.txt")
	if err != nil {
		log.Fatalf("Could not open file: %v", err)
	}
	defer file.Close()

	buffer := make([]byte, 1024) // 1KB chunks
	for {
		n, err := file.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error reading file: %v", err)
		}

		err = stream.Send(&pb.FileChunk{Data: buffer[:n]})
		if err != nil {
			log.Fatalf("Failed to send chunk: %v", err)
		}
	}

	resp, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Failed to receive response: %v", err)
	}

	fmt.Printf("Upload Status: %v\n", resp.GetMessage())
}
