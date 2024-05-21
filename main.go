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

func (s *myChessServer) CreateRoom(context.Context, *chess.CreateRoomRequest) (*chess.RoomResponse, error) {
	return &chess.RoomResponse{}, nil
}

func (s *myChessServer) JoinRoom(context.Context, *chess.JoinRoomRequest) (*chess.RoomResponse, error) {
	return &chess.RoomResponse{}, nil
}

func (s *myChessServer) Moves(stream chess.Chess_MovesServer) error {
	return nil
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
