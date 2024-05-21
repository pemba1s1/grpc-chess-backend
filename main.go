package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"

	"github.com/google/uuid"
	"github.com/pemba1s1/chess-backend/chess"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Room struct {
	RoomId          string
	Player_1        chess.Player
	Player_2        chess.Player
	Player_1_Stream *chess.Chess_MovesServer
	Player_2_Stream *chess.Chess_MovesServer
}

type myChessServer struct {
	rooms map[string]*Room
	chess.UnimplementedChessServer
}

func (s *myChessServer) CreateRoom(ctx context.Context, r *chess.CreateRoomRequest) (*chess.RoomResponse, error) {
	fmt.Printf("Creating Room, %v", r.Player_1.Color)
	room := &Room{
		RoomId: uuid.New().String(),
		Player_1: chess.Player{
			Name:  r.Player_1.Name,
			Color: r.Player_1.Color,
		},
		Player_1_Stream: nil,
	}

	s.rooms[room.RoomId] = room
	return &chess.RoomResponse{RoomId: room.RoomId, Status: "Room Created"}, nil
}

func (s *myChessServer) JoinRoom(ctx context.Context, r *chess.JoinRoomRequest) (*chess.RoomResponse, error) {
	room, ok := s.rooms[r.RoomId]
	if !ok {
		return &chess.RoomResponse{Status: "Room Not Found"}, nil
	}
	room.Player_2 = chess.Player{
		Name:  r.Player_2.Name,
		Color: r.Player_2.Color,
	}
	room.Player_2_Stream = nil
	return &chess.RoomResponse{RoomId: room.RoomId, Status: "Room Joined"}, nil
}

func (s *myChessServer) GetRooms(context.Context, *chess.GetRoomRequest) (*chess.GetRoomsResponse, error) {
	roomsResponse := &chess.GetRoomsResponse{}
	for _, room := range s.rooms {
		roomsResponse.Rooms = append(roomsResponse.Rooms, &chess.RoomResponse{
			RoomId: room.RoomId,
		})
	}
	return roomsResponse, nil
}

func (s *myChessServer) BroadcastMove(room *Room, in *chess.MoveRequest) {
	if in.Player.Color == chess.Color_WHITE && room.Player_2_Stream != nil {
		broadcast(room.Player_2_Stream, &chess.MoveResponse{Move: in.Move, Player: in.Player})
	} else if in.Player.Color == chess.Color_BLACK && room.Player_1_Stream != nil {
		broadcast(room.Player_1_Stream, &chess.MoveResponse{Move: in.Move, Player: in.Player})
	}
}

func broadcast(stream *chess.Chess_MovesServer, moveResponse *chess.MoveResponse) {
	err := (*stream).Send(moveResponse)
	if err != nil {
		fmt.Printf("Cannot broadcast message, %v", err)
	}
}

func (s *myChessServer) Moves(stream chess.Chess_MovesServer) error {
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		room, ok := s.rooms[in.RoomId]
		if !ok {
			return fmt.Errorf("Room Not Found")
		}
		if in.Player.Color == chess.Color_WHITE && room.Player_1_Stream == nil {
			room.Player_1_Stream = &stream
		}
		if in.Player.Color == chess.Color_BLACK && room.Player_2_Stream == nil {
			room.Player_2_Stream = &stream
		}
		s.BroadcastMove(room, in)
		stream.Send(&chess.MoveResponse{Move: in.Move, Player: in.Player})
	}
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	serverRegistrar := grpc.NewServer()
	reflection.Register(serverRegistrar)
	service := &myChessServer{}
	service.rooms = make(map[string]*Room)
	chess.RegisterChessServer(serverRegistrar, service)
	fmt.Printf("Server Starting On Port 8080")
	err = serverRegistrar.Serve(lis)

	if err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
