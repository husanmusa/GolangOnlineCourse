package app

import (
	"auth/pkg/handler"
	pb "auth/proto"
	"google.golang.org/grpc"
	"log"
	"log/slog"
	"net"
	"net/rpc"
)

type Server struct {
	RPC *handler.AuthController
}

func (s *Server) Run() error {
	err := rpc.Register(s.RPC)
	if err != nil {
		slog.Error(err.Error())
		panic(err)
	}

	l, err := net.Listen("tcp", ":8545") // Standard Ethereum JSON-RPC port
	if err != nil {
		log.Fatal("listen error:", err)
	}
	defer l.Close()
	server := grpc.NewServer()

	pb.RegisterAuthServiceServer(server, s.RPC)

	log.Printf("server listening at %v", l.Addr())
	if err := server.Serve(l); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	return nil
}
