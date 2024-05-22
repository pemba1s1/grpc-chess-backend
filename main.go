package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/google/uuid"
	chess "github.com/pemba1s1/grpc-chess-backend/chess"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Room struct {
	RoomId          string
	Player_1        chess.Player
	Player_2        chess.Player
	Player_1_Stream *chess.Chess_ListenToRoomServer
	Player_2_Stream *chess.Chess_ListenToRoomServer
}

type myChessServer struct {
	rooms map[string]*Room
	chess.UnimplementedChessServer
}

func (s *myChessServer) CreateRoom(ctx context.Context, r *chess.CreateRoomRequest) (*chess.RoomResponse, error) {
	room := &Room{
		RoomId: uuid.New().String(),
		Player_1: chess.Player{
			Name:  r.Player_1.Name,
			Color: r.Player_1.Color,
		},
		Player_1_Stream: nil,
	}

	s.rooms[room.RoomId] = room
	fmt.Println("Creating Room,")
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
	fmt.Printf("Joining Room, %v", room)
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

func (s *myChessServer) GetRoomInfo(ctx context.Context, r *chess.GetRoomRequest) (*chess.RoomResponse, error) {
	room, ok := s.rooms[r.RoomId]
	if !ok {
		return &chess.RoomResponse{Status: "Room Not Found"}, nil
	}
	return &chess.RoomResponse{RoomId: room.RoomId, Player_1: &room.Player_1, Player_2: &room.Player_2, Status: "Room Found"}, nil
}

// func (s *myChessServer) BroadcastMove(room *Room, in *chess.MoveRequest) {
// 	if in.Player.Color == chess.Color_WHITE && room.Player_2_Stream != nil {
// 		broadcast(room.Player_2_Stream, &chess.MoveResponse{Move: in.Move, Player: in.Player})
// 	} else if in.Player.Color == chess.Color_BLACK && room.Player_1_Stream != nil {
// 		broadcast(room.Player_1_Stream, &chess.MoveResponse{Move: in.Move, Player: in.Player})
// 	}
// }

// func broadcast(stream *chess.Chess_MovesServer, moveResponse *chess.MoveResponse) {
// 	err := (*stream).Send(moveResponse)
// 	if err != nil {
// 		fmt.Printf("Cannot broadcast message, %v", err)
// 	}
// }

func (s *myChessServer) Moves(context.Context, *chess.MoveRequest) (*chess.MoveResponse, error) {
	// for {
	// 	in, err := stream.Recv()
	// 	if err == io.EOF {
	// 		return nil
	// 	}
	// 	if err != nil {
	// 		return err
	// 	}
	// 	room, ok := s.rooms[in.RoomId]
	// 	if !ok {
	// 		return fmt.Errorf("Room Not Found")
	// 	}
	// 	if in.Player.Color == chess.Color_WHITE && room.Player_1_Stream == nil {
	// 		room.Player_1_Stream = &stream
	// 	}
	// 	if in.Player.Color == chess.Color_BLACK && room.Player_2_Stream == nil {
	// 		room.Player_2_Stream = &stream
	// 	}
	// 	s.BroadcastMove(room, in)
	// 	stream.Send(&chess.MoveResponse{Move: in.Move, Player: in.Player})
	// }
	return &chess.MoveResponse{}, nil
}

func (s *myChessServer) ListenToRoom(r *chess.GetRoomRequest, stream chess.Chess_ListenToRoomServer) error {
	fmt.Println("Listening to Room")
	a := []string{"one", "two", "three", "four", "five"}
	for _, element := range a {
		// Do something with each element
		fmt.Println(element)
		c := &chess.RoomResponse{}
		c.RoomId = element
		if err := stream.Send(c); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", ":8082")
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
