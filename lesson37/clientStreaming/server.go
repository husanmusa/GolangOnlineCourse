package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"net"

	pb "fileUpload/proto"
)

type server struct {
	pb.UnimplementedFileUploadServiceServer
}

func (s *server) UploadFile(stream pb.FileUploadService_UploadFileServer) error {
	file, err := os.Create("uploaded_file.yaml")
	if err != nil {
		return err
	}
	defer file.Close()

	for {
		chunk, err := stream.Recv()
		if err == io.EOF {
			// End of file, send response
			return stream.SendAndClose(&pb.UploadStatus{
				Success: true,
				Message: "File uploaded successfully!",
			})
		}
		if err != nil {
			return err
		}
		fmt.Println(len(chunk.GetData()))
		_, err = file.Write(chunk.GetData())
		if err != nil {
			return err
		}
	}
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterFileUploadServiceServer(s, &server{})
	reflection.Register(s)

	fmt.Println("Server is running on port 50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
