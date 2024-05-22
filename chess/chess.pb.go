// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.1
// source: chess/chess.proto

package chess

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PieceType int32

const (
	PieceType_NONE   PieceType = 0
	PieceType_KING   PieceType = 1
	PieceType_QUEEN  PieceType = 2
	PieceType_ROOK   PieceType = 3
	PieceType_BISHOP PieceType = 4
	PieceType_KNIGHT PieceType = 5
	PieceType_PAWN   PieceType = 6
)

// Enum value maps for PieceType.
var (
	PieceType_name = map[int32]string{
		0: "NONE",
		1: "KING",
		2: "QUEEN",
		3: "ROOK",
		4: "BISHOP",
		5: "KNIGHT",
		6: "PAWN",
	}
	PieceType_value = map[string]int32{
		"NONE":   0,
		"KING":   1,
		"QUEEN":  2,
		"ROOK":   3,
		"BISHOP": 4,
		"KNIGHT": 5,
		"PAWN":   6,
	}
)

func (x PieceType) Enum() *PieceType {
	p := new(PieceType)
	*p = x
	return p
}

func (x PieceType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PieceType) Descriptor() protoreflect.EnumDescriptor {
	return file_chess_chess_proto_enumTypes[0].Descriptor()
}

func (PieceType) Type() protoreflect.EnumType {
	return &file_chess_chess_proto_enumTypes[0]
}

func (x PieceType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PieceType.Descriptor instead.
func (PieceType) EnumDescriptor() ([]byte, []int) {
	return file_chess_chess_proto_rawDescGZIP(), []int{0}
}

type Color int32

const (
	Color_WHITE Color = 0
	Color_BLACK Color = 1
)

// Enum value maps for Color.
var (
	Color_name = map[int32]string{
		0: "WHITE",
		1: "BLACK",
	}
	Color_value = map[string]int32{
		"WHITE": 0,
		"BLACK": 1,
	}
)

func (x Color) Enum() *Color {
	p := new(Color)
	*p = x
	return p
}

func (x Color) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Color) Descriptor() protoreflect.EnumDescriptor {
	return file_chess_chess_proto_enumTypes[1].Descriptor()
}

func (Color) Type() protoreflect.EnumType {
	return &file_chess_chess_proto_enumTypes[1]
}

func (x Color) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Color.Descriptor instead.
func (Color) EnumDescriptor() ([]byte, []int) {
	return file_chess_chess_proto_rawDescGZIP(), []int{1}
}

type Player struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Color Color  `protobuf:"varint,2,opt,name=color,proto3,enum=Color" json:"color,omitempty"`
}

func (x *Player) Reset() {
	*x = Player{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chess_chess_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Player) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Player) ProtoMessage() {}

func (x *Player) ProtoReflect() protoreflect.Message {
	mi := &file_chess_chess_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Player.ProtoReflect.Descriptor instead.
func (*Player) Descriptor() ([]byte, []int) {
	return file_chess_chess_proto_rawDescGZIP(), []int{0}
}

func (x *Player) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Player) GetColor() Color {
	if x != nil {
		return x.Color
	}
	return Color_WHITE
}

type CreateRoomRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Player_1 *Player `protobuf:"bytes,1,opt,name=player_1,json=player1,proto3" json:"player_1,omitempty"`
}

func (x *CreateRoomRequest) Reset() {
	*x = CreateRoomRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chess_chess_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRoomRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRoomRequest) ProtoMessage() {}

func (x *CreateRoomRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chess_chess_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRoomRequest.ProtoReflect.Descriptor instead.
func (*CreateRoomRequest) Descriptor() ([]byte, []int) {
	return file_chess_chess_proto_rawDescGZIP(), []int{1}
}

func (x *CreateRoomRequest) GetPlayer_1() *Player {
	if x != nil {
		return x.Player_1
	}
	return nil
}

type RoomResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomId   string  `protobuf:"bytes,1,opt,name=roomId,proto3" json:"roomId,omitempty"`
	Status   string  `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	Player_1 *Player `protobuf:"bytes,3,opt,name=player_1,json=player1,proto3" json:"player_1,omitempty"`
	Player_2 *Player `protobuf:"bytes,4,opt,name=player_2,json=player2,proto3" json:"player_2,omitempty"`
}

func (x *RoomResponse) Reset() {
	*x = RoomResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chess_chess_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoomResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomResponse) ProtoMessage() {}

func (x *RoomResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chess_chess_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomResponse.ProtoReflect.Descriptor instead.
func (*RoomResponse) Descriptor() ([]byte, []int) {
	return file_chess_chess_proto_rawDescGZIP(), []int{2}
}

func (x *RoomResponse) GetRoomId() string {
	if x != nil {
		return x.RoomId
	}
	return ""
}

func (x *RoomResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *RoomResponse) GetPlayer_1() *Player {
	if x != nil {
		return x.Player_1
	}
	return nil
}

func (x *RoomResponse) GetPlayer_2() *Player {
	if x != nil {
		return x.Player_2
	}
	return nil
}

type JoinRoomRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomId   string  `protobuf:"bytes,1,opt,name=roomId,proto3" json:"roomId,omitempty"`
	Player_2 *Player `protobuf:"bytes,2,opt,name=player_2,json=player2,proto3" json:"player_2,omitempty"`
}

func (x *JoinRoomRequest) Reset() {
	*x = JoinRoomRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chess_chess_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinRoomRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinRoomRequest) ProtoMessage() {}

func (x *JoinRoomRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chess_chess_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinRoomRequest.ProtoReflect.Descriptor instead.
func (*JoinRoomRequest) Descriptor() ([]byte, []int) {
	return file_chess_chess_proto_rawDescGZIP(), []int{3}
}

func (x *JoinRoomRequest) GetRoomId() string {
	if x != nil {
		return x.RoomId
	}
	return ""
}

func (x *JoinRoomRequest) GetPlayer_2() *Player {
	if x != nil {
		return x.Player_2
	}
	return nil
}

type GetRoomsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rooms []*RoomResponse `protobuf:"bytes,1,rep,name=rooms,proto3" json:"rooms,omitempty"`
}

func (x *GetRoomsResponse) Reset() {
	*x = GetRoomsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chess_chess_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRoomsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRoomsResponse) ProtoMessage() {}

func (x *GetRoomsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chess_chess_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRoomsResponse.ProtoReflect.Descriptor instead.
func (*GetRoomsResponse) Descriptor() ([]byte, []int) {
	return file_chess_chess_proto_rawDescGZIP(), []int{4}
}

func (x *GetRoomsResponse) GetRooms() []*RoomResponse {
	if x != nil {
		return x.Rooms
	}
	return nil
}

type Coordinate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X int32 `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
	Y int32 `protobuf:"varint,2,opt,name=y,proto3" json:"y,omitempty"`
}

func (x *Coordinate) Reset() {
	*x = Coordinate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chess_chess_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Coordinate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Coordinate) ProtoMessage() {}

func (x *Coordinate) ProtoReflect() protoreflect.Message {
	mi := &file_chess_chess_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Coordinate.ProtoReflect.Descriptor instead.
func (*Coordinate) Descriptor() ([]byte, []int) {
	return file_chess_chess_proto_rawDescGZIP(), []int{5}
}

func (x *Coordinate) GetX() int32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *Coordinate) GetY() int32 {
	if x != nil {
		return x.Y
	}
	return 0
}

type Move struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From *Coordinate `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To   *Coordinate `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
}

func (x *Move) Reset() {
	*x = Move{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chess_chess_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Move) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Move) ProtoMessage() {}

func (x *Move) ProtoReflect() protoreflect.Message {
	mi := &file_chess_chess_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Move.ProtoReflect.Descriptor instead.
func (*Move) Descriptor() ([]byte, []int) {
	return file_chess_chess_proto_rawDescGZIP(), []int{6}
}

func (x *Move) GetFrom() *Coordinate {
	if x != nil {
		return x.From
	}
	return nil
}

func (x *Move) GetTo() *Coordinate {
	if x != nil {
		return x.To
	}
	return nil
}

type MoveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomId      string    `protobuf:"bytes,1,opt,name=roomId,proto3" json:"roomId,omitempty"`
	Player      *Player   `protobuf:"bytes,2,opt,name=player,proto3" json:"player,omitempty"`
	Move        *Move     `protobuf:"bytes,3,opt,name=move,proto3" json:"move,omitempty"`
	KilledPiece PieceType `protobuf:"varint,4,opt,name=killedPiece,proto3,enum=PieceType" json:"killedPiece,omitempty"`
}

func (x *MoveRequest) Reset() {
	*x = MoveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chess_chess_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MoveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MoveRequest) ProtoMessage() {}

func (x *MoveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chess_chess_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MoveRequest.ProtoReflect.Descriptor instead.
func (*MoveRequest) Descriptor() ([]byte, []int) {
	return file_chess_chess_proto_rawDescGZIP(), []int{7}
}

func (x *MoveRequest) GetRoomId() string {
	if x != nil {
		return x.RoomId
	}
	return ""
}

func (x *MoveRequest) GetPlayer() *Player {
	if x != nil {
		return x.Player
	}
	return nil
}

func (x *MoveRequest) GetMove() *Move {
	if x != nil {
		return x.Move
	}
	return nil
}

func (x *MoveRequest) GetKilledPiece() PieceType {
	if x != nil {
		return x.KilledPiece
	}
	return PieceType_NONE
}

type MoveResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Move        *Move     `protobuf:"bytes,1,opt,name=move,proto3" json:"move,omitempty"`
	Player      *Player   `protobuf:"bytes,2,opt,name=player,proto3" json:"player,omitempty"`
	KilledPiece PieceType `protobuf:"varint,3,opt,name=killedPiece,proto3,enum=PieceType" json:"killedPiece,omitempty"`
}

func (x *MoveResponse) Reset() {
	*x = MoveResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chess_chess_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MoveResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MoveResponse) ProtoMessage() {}

func (x *MoveResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chess_chess_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MoveResponse.ProtoReflect.Descriptor instead.
func (*MoveResponse) Descriptor() ([]byte, []int) {
	return file_chess_chess_proto_rawDescGZIP(), []int{8}
}

func (x *MoveResponse) GetMove() *Move {
	if x != nil {
		return x.Move
	}
	return nil
}

func (x *MoveResponse) GetPlayer() *Player {
	if x != nil {
		return x.Player
	}
	return nil
}

func (x *MoveResponse) GetKilledPiece() PieceType {
	if x != nil {
		return x.KilledPiece
	}
	return PieceType_NONE
}

type GetRoomRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomId string `protobuf:"bytes,1,opt,name=roomId,proto3" json:"roomId,omitempty"`
}

func (x *GetRoomRequest) Reset() {
	*x = GetRoomRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chess_chess_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRoomRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRoomRequest) ProtoMessage() {}

func (x *GetRoomRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chess_chess_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRoomRequest.ProtoReflect.Descriptor instead.
func (*GetRoomRequest) Descriptor() ([]byte, []int) {
	return file_chess_chess_proto_rawDescGZIP(), []int{9}
}

func (x *GetRoomRequest) GetRoomId() string {
	if x != nil {
		return x.RoomId
	}
	return ""
}

type RoomResponseStream struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Room *RoomResponse `protobuf:"bytes,1,opt,name=room,proto3" json:"room,omitempty"`
	Move *MoveResponse `protobuf:"bytes,2,opt,name=move,proto3" json:"move,omitempty"`
}

func (x *RoomResponseStream) Reset() {
	*x = RoomResponseStream{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chess_chess_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoomResponseStream) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomResponseStream) ProtoMessage() {}

func (x *RoomResponseStream) ProtoReflect() protoreflect.Message {
	mi := &file_chess_chess_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomResponseStream.ProtoReflect.Descriptor instead.
func (*RoomResponseStream) Descriptor() ([]byte, []int) {
	return file_chess_chess_proto_rawDescGZIP(), []int{10}
}

func (x *RoomResponseStream) GetRoom() *RoomResponse {
	if x != nil {
		return x.Room
	}
	return nil
}

func (x *RoomResponseStream) GetMove() *MoveResponse {
	if x != nil {
		return x.Move
	}
	return nil
}

var File_chess_chess_proto protoreflect.FileDescriptor

var file_chess_chess_proto_rawDesc = []byte{
	0x0a, 0x11, 0x63, 0x68, 0x65, 0x73, 0x73, 0x2f, 0x63, 0x68, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x3a, 0x0a, 0x06, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x1c, 0x0a, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x06, 0x2e, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x52, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x22,
	0x37, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x31,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52,
	0x07, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x31, 0x22, 0x86, 0x01, 0x0a, 0x0c, 0x52, 0x6f, 0x6f,
	0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x6f, 0x6f,
	0x6d, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x22, 0x0a, 0x08, 0x70, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x5f, 0x31, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x50, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x52, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x31, 0x12, 0x22, 0x0a,
	0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x32, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x07, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x32, 0x22, 0x4d, 0x0a, 0x0f, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x08,
	0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07,
	0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x32,
	0x22, 0x37, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x05, 0x72, 0x6f, 0x6f, 0x6d, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x52, 0x05, 0x72, 0x6f, 0x6f, 0x6d, 0x73, 0x22, 0x28, 0x0a, 0x0a, 0x43, 0x6f, 0x6f,
	0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x01, 0x79, 0x22, 0x44, 0x0a, 0x04, 0x4d, 0x6f, 0x76, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x66,
	0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x43, 0x6f, 0x6f, 0x72,
	0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x1b, 0x0a, 0x02,
	0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x43, 0x6f, 0x6f, 0x72, 0x64,
	0x69, 0x6e, 0x61, 0x74, 0x65, 0x52, 0x02, 0x74, 0x6f, 0x22, 0x8f, 0x01, 0x0a, 0x0b, 0x4d, 0x6f,
	0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x6f, 0x6f,
	0x6d, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49,
	0x64, 0x12, 0x1f, 0x0a, 0x06, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x07, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x06, 0x70, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x12, 0x19, 0x0a, 0x04, 0x6d, 0x6f, 0x76, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x05, 0x2e, 0x4d, 0x6f, 0x76, 0x65, 0x52, 0x04, 0x6d, 0x6f, 0x76, 0x65, 0x12, 0x2c, 0x0a,
	0x0b, 0x6b, 0x69, 0x6c, 0x6c, 0x65, 0x64, 0x50, 0x69, 0x65, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x0a, 0x2e, 0x50, 0x69, 0x65, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b,
	0x6b, 0x69, 0x6c, 0x6c, 0x65, 0x64, 0x50, 0x69, 0x65, 0x63, 0x65, 0x22, 0x78, 0x0a, 0x0c, 0x4d,
	0x6f, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x04, 0x6d,
	0x6f, 0x76, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x4d, 0x6f, 0x76, 0x65,
	0x52, 0x04, 0x6d, 0x6f, 0x76, 0x65, 0x12, 0x1f, 0x0a, 0x06, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52,
	0x06, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x2c, 0x0a, 0x0b, 0x6b, 0x69, 0x6c, 0x6c, 0x65,
	0x64, 0x50, 0x69, 0x65, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0a, 0x2e, 0x50,
	0x69, 0x65, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x6b, 0x69, 0x6c, 0x6c, 0x65, 0x64,
	0x50, 0x69, 0x65, 0x63, 0x65, 0x22, 0x28, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6f, 0x6d,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x22,
	0x5a, 0x0a, 0x12, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x21, 0x0a, 0x04, 0x72, 0x6f, 0x6f, 0x6d, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x52, 0x04, 0x72, 0x6f, 0x6f, 0x6d, 0x12, 0x21, 0x0a, 0x04, 0x6d, 0x6f, 0x76, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x4d, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x04, 0x6d, 0x6f, 0x76, 0x65, 0x2a, 0x56, 0x0a, 0x09, 0x50,
	0x69, 0x65, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45,
	0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x4b, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05,
	0x51, 0x55, 0x45, 0x45, 0x4e, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x52, 0x4f, 0x4f, 0x4b, 0x10,
	0x03, 0x12, 0x0a, 0x0a, 0x06, 0x42, 0x49, 0x53, 0x48, 0x4f, 0x50, 0x10, 0x04, 0x12, 0x0a, 0x0a,
	0x06, 0x4b, 0x4e, 0x49, 0x47, 0x48, 0x54, 0x10, 0x05, 0x12, 0x08, 0x0a, 0x04, 0x50, 0x41, 0x57,
	0x4e, 0x10, 0x06, 0x2a, 0x1d, 0x0a, 0x05, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x09, 0x0a, 0x05,
	0x57, 0x48, 0x49, 0x54, 0x45, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x42, 0x4c, 0x41, 0x43, 0x4b,
	0x10, 0x01, 0x32, 0xa8, 0x02, 0x0a, 0x05, 0x43, 0x68, 0x65, 0x73, 0x73, 0x12, 0x31, 0x0a, 0x0a,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x12, 0x12, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d,
	0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x2d, 0x0a, 0x08, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x6f, 0x6f, 0x6d, 0x12, 0x10, 0x2e, 0x4a, 0x6f,
	0x69, 0x6e, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e,
	0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x30,
	0x0a, 0x08, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x73, 0x12, 0x0f, 0x2e, 0x47, 0x65, 0x74,
	0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x47, 0x65,
	0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x26, 0x0a, 0x05, 0x4d, 0x6f, 0x76, 0x65, 0x73, 0x12, 0x0c, 0x2e, 0x4d, 0x6f, 0x76, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x4d, 0x6f, 0x76, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2f, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x52,
	0x6f, 0x6f, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0f, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6f,
	0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x0c, 0x4c, 0x69, 0x73,
	0x74, 0x65, 0x6e, 0x54, 0x6f, 0x52, 0x6f, 0x6f, 0x6d, 0x12, 0x0f, 0x2e, 0x47, 0x65, 0x74, 0x52,
	0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x52, 0x6f, 0x6f,
	0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x42, 0x29, 0x5a,
	0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x65, 0x6d, 0x62,
	0x61, 0x31, 0x73, 0x31, 0x2f, 0x63, 0x68, 0x65, 0x73, 0x73, 0x2d, 0x62, 0x61, 0x63, 0x6b, 0x65,
	0x6e, 0x64, 0x2f, 0x63, 0x68, 0x65, 0x73, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chess_chess_proto_rawDescOnce sync.Once
	file_chess_chess_proto_rawDescData = file_chess_chess_proto_rawDesc
)

func file_chess_chess_proto_rawDescGZIP() []byte {
	file_chess_chess_proto_rawDescOnce.Do(func() {
		file_chess_chess_proto_rawDescData = protoimpl.X.CompressGZIP(file_chess_chess_proto_rawDescData)
	})
	return file_chess_chess_proto_rawDescData
}

var file_chess_chess_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_chess_chess_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_chess_chess_proto_goTypes = []interface{}{
	(PieceType)(0),             // 0: PieceType
	(Color)(0),                 // 1: Color
	(*Player)(nil),             // 2: Player
	(*CreateRoomRequest)(nil),  // 3: CreateRoomRequest
	(*RoomResponse)(nil),       // 4: RoomResponse
	(*JoinRoomRequest)(nil),    // 5: JoinRoomRequest
	(*GetRoomsResponse)(nil),   // 6: GetRoomsResponse
	(*Coordinate)(nil),         // 7: Coordinate
	(*Move)(nil),               // 8: Move
	(*MoveRequest)(nil),        // 9: MoveRequest
	(*MoveResponse)(nil),       // 10: MoveResponse
	(*GetRoomRequest)(nil),     // 11: GetRoomRequest
	(*RoomResponseStream)(nil), // 12: RoomResponseStream
}
var file_chess_chess_proto_depIdxs = []int32{
	1,  // 0: Player.color:type_name -> Color
	2,  // 1: CreateRoomRequest.player_1:type_name -> Player
	2,  // 2: RoomResponse.player_1:type_name -> Player
	2,  // 3: RoomResponse.player_2:type_name -> Player
	2,  // 4: JoinRoomRequest.player_2:type_name -> Player
	4,  // 5: GetRoomsResponse.rooms:type_name -> RoomResponse
	7,  // 6: Move.from:type_name -> Coordinate
	7,  // 7: Move.to:type_name -> Coordinate
	2,  // 8: MoveRequest.player:type_name -> Player
	8,  // 9: MoveRequest.move:type_name -> Move
	0,  // 10: MoveRequest.killedPiece:type_name -> PieceType
	8,  // 11: MoveResponse.move:type_name -> Move
	2,  // 12: MoveResponse.player:type_name -> Player
	0,  // 13: MoveResponse.killedPiece:type_name -> PieceType
	4,  // 14: RoomResponseStream.room:type_name -> RoomResponse
	10, // 15: RoomResponseStream.move:type_name -> MoveResponse
	3,  // 16: Chess.CreateRoom:input_type -> CreateRoomRequest
	5,  // 17: Chess.JoinRoom:input_type -> JoinRoomRequest
	11, // 18: Chess.GetRooms:input_type -> GetRoomRequest
	9,  // 19: Chess.Moves:input_type -> MoveRequest
	11, // 20: Chess.GetRoomInfo:input_type -> GetRoomRequest
	11, // 21: Chess.ListenToRoom:input_type -> GetRoomRequest
	4,  // 22: Chess.CreateRoom:output_type -> RoomResponse
	4,  // 23: Chess.JoinRoom:output_type -> RoomResponse
	6,  // 24: Chess.GetRooms:output_type -> GetRoomsResponse
	10, // 25: Chess.Moves:output_type -> MoveResponse
	4,  // 26: Chess.GetRoomInfo:output_type -> RoomResponse
	4,  // 27: Chess.ListenToRoom:output_type -> RoomResponse
	22, // [22:28] is the sub-list for method output_type
	16, // [16:22] is the sub-list for method input_type
	16, // [16:16] is the sub-list for extension type_name
	16, // [16:16] is the sub-list for extension extendee
	0,  // [0:16] is the sub-list for field type_name
}

func init() { file_chess_chess_proto_init() }
func file_chess_chess_proto_init() {
	if File_chess_chess_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_chess_chess_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Player); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_chess_chess_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRoomRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_chess_chess_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoomResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_chess_chess_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinRoomRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_chess_chess_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRoomsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_chess_chess_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Coordinate); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_chess_chess_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Move); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_chess_chess_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MoveRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_chess_chess_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MoveResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_chess_chess_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRoomRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_chess_chess_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoomResponseStream); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_chess_chess_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_chess_chess_proto_goTypes,
		DependencyIndexes: file_chess_chess_proto_depIdxs,
		EnumInfos:         file_chess_chess_proto_enumTypes,
		MessageInfos:      file_chess_chess_proto_msgTypes,
	}.Build()
	File_chess_chess_proto = out.File
	file_chess_chess_proto_rawDesc = nil
	file_chess_chess_proto_goTypes = nil
	file_chess_chess_proto_depIdxs = nil
}
