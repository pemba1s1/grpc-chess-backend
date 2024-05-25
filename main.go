package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync"

	"github.com/google/uuid"
	chess "github.com/pemba1s1/grpc-chess-backend/chess"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Room struct {
	RoomId          string
	Player_1        *chess.Player
	Player_2        *chess.Player
	Player_1_Stream *chess.Chess_ListenToRoomServer
	Player_2_Stream *chess.Chess_ListenToRoomServer
}

type myChessServer struct {
	rooms map[string]*Room
	ch    chan chessChannel
	mu    sync.Mutex
	chess.UnimplementedChessServer
}

type chessChannel struct {
	RoomId string
	From   *chess.Coordinate
	To     *chess.Coordinate
	Player *chess.Player
}

func (s *myChessServer) CreateRoom(ctx context.Context, r *chess.CreateRoomRequest) (*chess.RoomResponse, error) {
	room := &Room{
		RoomId: uuid.New().String(),
		Player_1: &chess.Player{
			Name:  r.Player_1.Name,
			Color: r.Player_1.Color,
		},
		Player_1_Stream: nil,
	}

	s.rooms[room.RoomId] = room
	log.Printf("Creating Room, %v \n", room.RoomId)
	return &chess.RoomResponse{RoomId: room.RoomId, Status: "Room Created"}, nil
}

func (s *myChessServer) JoinRoom(ctx context.Context, r *chess.JoinRoomRequest) (*chess.RoomResponse, error) {
	room, ok := s.rooms[r.RoomId]
	if !ok {
		return &chess.RoomResponse{Status: "Room Not Found"}, nil
	}
	player_2 := &chess.Player{
		Name:  r.Player_2.Name,
		Color: r.Player_2.Color,
	}
	room.Player_2 = player_2
	room.Player_2_Stream = nil
	chessChannel := &chessChannel{RoomId: room.RoomId, Player: player_2}
	log.Printf("%v is joining Room, %v \n", player_2.Name, room.RoomId)
	if room.Player_1_Stream != nil {
		s.mu.Lock()
		s.ch <- *chessChannel
		s.mu.Unlock()
	}
	return &chess.RoomResponse{RoomId: r.RoomId, Status: "Room Joined"}, nil
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
	return &chess.RoomResponse{RoomId: room.RoomId, Player_1: room.Player_1, Player_2: room.Player_2, Status: "Room Found"}, nil
}

func (s *myChessServer) Moves(ctx context.Context, r *chess.MoveRequest) (*chess.MoveResponse, error) {
	moveResponse := &chess.MoveResponse{Move: r.Move, Player: r.Player}
	chessChannel := &chessChannel{RoomId: r.RoomId, From: r.Move.From, To: r.Move.To, Player: r.Player}
	log.Printf("Move Received from %v \n", r.Player.Name)
	log.Printf("Move From (%v, %v) to (%v, %v) \n", r.Move.From.X, r.Move.From.Y, r.Move.To.X, r.Move.To.Y)
	s.mu.Lock()
	s.ch <- *chessChannel
	s.mu.Unlock()
	return moveResponse, nil
}

func (s *myChessServer) ListenToRoom(r *chess.MoveRequest, stream chess.Chess_ListenToRoomServer) error {
	room, ok := s.rooms[r.RoomId]
	if !ok {
		return fmt.Errorf("Room Not Found")
	}
	log.Printf("%v Listening to Room %v \n", r.Player.Name, r.RoomId)
	// send board, current player
	if r.Player.Color == chess.Color_WHITE {
		room.Player_1_Stream = &stream
	}
	if r.Player.Color == chess.Color_BLACK {
		room.Player_2_Stream = &stream
	}

	for chessChannel := range s.ch {
		s.mu.Lock()
		roomFromCh, ok := s.rooms[chessChannel.RoomId]
		if !ok {
			return fmt.Errorf("room not found")
		}
		if chessChannel.Player == nil {
			return fmt.Errorf("player not found")
		}
		roomResponseStream := &chess.RoomResponseStream{
			RoomId: chessChannel.RoomId,
			Player: chessChannel.Player,
		}
		if chessChannel.From != nil && chessChannel.To != nil {
			roomResponseStream.Move = &chess.Move{From: chessChannel.From, To: chessChannel.To}
		}
		log.Printf("Sending message from %v \n", chessChannel.Player.Name)
		if chessChannel.Player.Color == chess.Color_WHITE {
			if err := (*roomFromCh.Player_2_Stream).Send(roomResponseStream); err != nil {
				log.Printf("Cannot send message to %v, %v", roomFromCh.Player_2.Name, err)
			}
		}
		if chessChannel.Player.Color == chess.Color_BLACK {
			if err := (*roomFromCh.Player_1_Stream).Send(roomResponseStream); err != nil {
				log.Printf("Cannot send message to %v, %v", roomFromCh.Player_1.Name, err)
			}
		}
		s.mu.Unlock()
	}
	log.Printf("%v stopped listening to room %v", r.Player.Name, r.RoomId)
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
	service.ch = make(chan chessChannel)
	chess.RegisterChessServer(serverRegistrar, service)
	log.Printf("Server Starting On Port 8080")
	err = serverRegistrar.Serve(lis)

	if err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
