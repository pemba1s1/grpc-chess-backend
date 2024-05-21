package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/pemba1s1/chess-backend/chess"
	"google.golang.org/grpc"
)

type myChessServer struct {
	chess.UnimplementedChessServer
}

func (s *myChessServer) NewGame(context.Context, *chess.NewGameRequest) (*chess.NewGameResponse, error) {
	return &chess.NewGameResponse{}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	serverRegistrar := grpc.NewServer()
	service := &myChessServer{}
	chess.RegisterChessServer(serverRegistrar, service)
	fmt.Printf("Server Starting On Port 8080")
	err = serverRegistrar.Serve(lis)

	if err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
