// Code generated by protoc-gen-go.
// source: zjh_base.proto
// DO NOT EDIT!

package ddproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type EProtoId int32

const (
	// //////////////////////////////////
	EProtoId_ZJH_PID_HEARTBEAT           EProtoId = 0
	EProtoId_ZJH_PID_QUICKCOnn           EProtoId = 1
	EProtoId_ZJH_PID_QUICKCONNA_ACK      EProtoId = 2
	EProtoId_ZJH_PID_GAME_LOGIN          EProtoId = 3
	EProtoId_ZJH_PID_GAME_LOGIN_ACK      EProtoId = 4
	EProtoId_ZJH_PID_GET_ROOM_LIST       EProtoId = 5
	EProtoId_ZJH_PID_ENTER_ROOM_LIST_ACK EProtoId = 6
	EProtoId_ZJH_PID_AUTO_ENTER_DESK     EProtoId = 7
	EProtoId_ZJH_PID_AUTO_ENTER_DESK_ACK EProtoId = 8
	EProtoId_ZJH_PID_SEND_GAMEINFO       EProtoId = 9
)

var EProtoId_name = map[int32]string{
	0: "ZJH_PID_HEARTBEAT",
	1: "ZJH_PID_QUICKCOnn",
	2: "ZJH_PID_QUICKCONNA_ACK",
	3: "ZJH_PID_GAME_LOGIN",
	4: "ZJH_PID_GAME_LOGIN_ACK",
	5: "ZJH_PID_GET_ROOM_LIST",
	6: "ZJH_PID_ENTER_ROOM_LIST_ACK",
	7: "ZJH_PID_AUTO_ENTER_DESK",
	8: "ZJH_PID_AUTO_ENTER_DESK_ACK",
	9: "ZJH_PID_SEND_GAMEINFO",
}
var EProtoId_value = map[string]int32{
	"ZJH_PID_HEARTBEAT":           0,
	"ZJH_PID_QUICKCOnn":           1,
	"ZJH_PID_QUICKCONNA_ACK":      2,
	"ZJH_PID_GAME_LOGIN":          3,
	"ZJH_PID_GAME_LOGIN_ACK":      4,
	"ZJH_PID_GET_ROOM_LIST":       5,
	"ZJH_PID_ENTER_ROOM_LIST_ACK": 6,
	"ZJH_PID_AUTO_ENTER_DESK":     7,
	"ZJH_PID_AUTO_ENTER_DESK_ACK": 8,
	"ZJH_PID_SEND_GAMEINFO":       9,
}

func (x EProtoId) Enum() *EProtoId {
	p := new(EProtoId)
	*p = x
	return p
}
func (x EProtoId) String() string {
	return proto.EnumName(EProtoId_name, int32(x))
}
func (x *EProtoId) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(EProtoId_value, data, "EProtoId")
	if err != nil {
		return err
	}
	*x = EProtoId(value)
	return nil
}
func (EProtoId) EnumDescriptor() ([]byte, []int) { return fileDescriptor10, []int{0} }

// 玩家当前状态
type ZjhEnumPlayerGameStatus int32

const (
	ZjhEnumPlayerGameStatus_ZJH_TEMP ZjhEnumPlayerGameStatus = 0
)

var ZjhEnumPlayerGameStatus_name = map[int32]string{
	0: "ZJH_TEMP",
}
var ZjhEnumPlayerGameStatus_value = map[string]int32{
	"ZJH_TEMP": 0,
}

func (x ZjhEnumPlayerGameStatus) Enum() *ZjhEnumPlayerGameStatus {
	p := new(ZjhEnumPlayerGameStatus)
	*p = x
	return p
}
func (x ZjhEnumPlayerGameStatus) String() string {
	return proto.EnumName(ZjhEnumPlayerGameStatus_name, int32(x))
}
func (x *ZjhEnumPlayerGameStatus) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ZjhEnumPlayerGameStatus_value, data, "ZjhEnumPlayerGameStatus")
	if err != nil {
		return err
	}
	*x = ZjhEnumPlayerGameStatus(value)
	return nil
}
func (ZjhEnumPlayerGameStatus) EnumDescriptor() ([]byte, []int) { return fileDescriptor10, []int{1} }

// 桌面状态
type ZjhEnumDeskState int32

const (
	ZjhEnumDeskState_DESK_IS_GAMING ZjhEnumDeskState = 1
	ZjhEnumDeskState_DESK_IS_WAIT   ZjhEnumDeskState = 2
)

var ZjhEnumDeskState_name = map[int32]string{
	1: "DESK_IS_GAMING",
	2: "DESK_IS_WAIT",
}
var ZjhEnumDeskState_value = map[string]int32{
	"DESK_IS_GAMING": 1,
	"DESK_IS_WAIT":   2,
}

func (x ZjhEnumDeskState) Enum() *ZjhEnumDeskState {
	p := new(ZjhEnumDeskState)
	*p = x
	return p
}
func (x ZjhEnumDeskState) String() string {
	return proto.EnumName(ZjhEnumDeskState_name, int32(x))
}
func (x *ZjhEnumDeskState) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ZjhEnumDeskState_value, data, "ZjhEnumDeskState")
	if err != nil {
		return err
	}
	*x = ZjhEnumDeskState(value)
	return nil
}
func (ZjhEnumDeskState) EnumDescriptor() ([]byte, []int) { return fileDescriptor10, []int{2} }

// 房间玩家状态
type ZjhEnumUserState int32

const (
	ZjhEnumUserState_USER_IS_GAMING ZjhEnumUserState = 1
	ZjhEnumUserState_USER_IS_STAND  ZjhEnumUserState = 2
	ZjhEnumUserState_USER_IS_SITED  ZjhEnumUserState = 3
)

var ZjhEnumUserState_name = map[int32]string{
	1: "USER_IS_GAMING",
	2: "USER_IS_STAND",
	3: "USER_IS_SITED",
}
var ZjhEnumUserState_value = map[string]int32{
	"USER_IS_GAMING": 1,
	"USER_IS_STAND":  2,
	"USER_IS_SITED":  3,
}

func (x ZjhEnumUserState) Enum() *ZjhEnumUserState {
	p := new(ZjhEnumUserState)
	*p = x
	return p
}
func (x ZjhEnumUserState) String() string {
	return proto.EnumName(ZjhEnumUserState_name, int32(x))
}
func (x *ZjhEnumUserState) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ZjhEnumUserState_value, data, "ZjhEnumUserState")
	if err != nil {
		return err
	}
	*x = ZjhEnumUserState(value)
	return nil
}
func (ZjhEnumUserState) EnumDescriptor() ([]byte, []int) { return fileDescriptor10, []int{3} }

// 房间类型
type ZjhEnumRoomType int32

const (
	ZjhEnumRoomType_ROOM_TYPE_FRIEND   ZjhEnumRoomType = 1
	ZjhEnumRoomType_ROOM_TYPE_NORMAL   ZjhEnumRoomType = 2
	ZjhEnumRoomType_ROOM_TYPE_REDBLACK ZjhEnumRoomType = 3
)

var ZjhEnumRoomType_name = map[int32]string{
	1: "ROOM_TYPE_FRIEND",
	2: "ROOM_TYPE_NORMAL",
	3: "ROOM_TYPE_REDBLACK",
}
var ZjhEnumRoomType_value = map[string]int32{
	"ROOM_TYPE_FRIEND":   1,
	"ROOM_TYPE_NORMAL":   2,
	"ROOM_TYPE_REDBLACK": 3,
}

func (x ZjhEnumRoomType) Enum() *ZjhEnumRoomType {
	p := new(ZjhEnumRoomType)
	*p = x
	return p
}
func (x ZjhEnumRoomType) String() string {
	return proto.EnumName(ZjhEnumRoomType_name, int32(x))
}
func (x *ZjhEnumRoomType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ZjhEnumRoomType_value, data, "ZjhEnumRoomType")
	if err != nil {
		return err
	}
	*x = ZjhEnumRoomType(value)
	return nil
}
func (ZjhEnumRoomType) EnumDescriptor() ([]byte, []int) { return fileDescriptor10, []int{4} }

func init() {
	proto.RegisterEnum("ddproto.EProtoId", EProtoId_name, EProtoId_value)
	proto.RegisterEnum("ddproto.ZjhEnumPlayerGameStatus", ZjhEnumPlayerGameStatus_name, ZjhEnumPlayerGameStatus_value)
	proto.RegisterEnum("ddproto.ZjhEnumDeskState", ZjhEnumDeskState_name, ZjhEnumDeskState_value)
	proto.RegisterEnum("ddproto.ZjhEnumUserState", ZjhEnumUserState_name, ZjhEnumUserState_value)
	proto.RegisterEnum("ddproto.ZjhEnumRoomType", ZjhEnumRoomType_name, ZjhEnumRoomType_value)
}

var fileDescriptor10 = []byte{
	// 351 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x5c, 0x91, 0xc1, 0x6e, 0xe2, 0x30,
	0x14, 0x45, 0x19, 0x98, 0x19, 0x98, 0x27, 0x06, 0x19, 0xab, 0x50, 0x51, 0x16, 0x5d, 0x37, 0x8b,
	0x7e, 0x40, 0x77, 0x86, 0x98, 0xe0, 0x92, 0x38, 0x69, 0x6c, 0x84, 0xda, 0x8d, 0x95, 0x2a, 0x96,
	0xaa, 0xb6, 0x10, 0x94, 0xc0, 0x82, 0xfe, 0x4b, 0xff, 0xb5, 0x49, 0x68, 0x82, 0xcb, 0xce, 0x3a,
	0xf7, 0x9d, 0x67, 0xe7, 0x06, 0x7a, 0x1f, 0xaf, 0x2f, 0xea, 0x39, 0xca, 0xf4, 0xed, 0x36, 0x4d,
	0x76, 0x09, 0x6e, 0xc7, 0x71, 0x79, 0xb0, 0x3e, 0x9b, 0xd0, 0xa1, 0x41, 0x71, 0x64, 0x31, 0x1e,
	0x40, 0xff, 0xe9, 0x7e, 0xae, 0x02, 0x66, 0xab, 0x39, 0x25, 0xa1, 0x9c, 0x50, 0x22, 0x51, 0xc3,
	0xc4, 0x0f, 0x4b, 0x36, 0x5d, 0x4c, 0xfd, 0xcd, 0x06, 0xfd, 0xc2, 0x57, 0x30, 0x3c, 0xc3, 0x9c,
	0x13, 0x45, 0xa6, 0x0b, 0xd4, 0xc4, 0x43, 0xc0, 0x55, 0xe6, 0x10, 0x8f, 0x2a, 0xd7, 0x77, 0x18,
	0x47, 0x2d, 0xd3, 0x39, 0xf1, 0xd2, 0xf9, 0x8d, 0x47, 0x30, 0xa8, 0x33, 0x2a, 0x55, 0xe8, 0xfb,
	0x9e, 0x72, 0x99, 0x90, 0xe8, 0x0f, 0xbe, 0x86, 0x71, 0x15, 0x51, 0x2e, 0x69, 0x78, 0x0a, 0x4b,
	0xf7, 0x2f, 0x1e, 0xc3, 0x65, 0x35, 0x40, 0x96, 0xd2, 0xff, 0x9e, 0xb2, 0xa9, 0x58, 0xa0, 0xb6,
	0x69, 0x9f, 0x85, 0xa5, 0xdd, 0x31, 0x6f, 0x16, 0x94, 0x1f, 0x9f, 0xc6, 0xf8, 0xcc, 0x47, 0xff,
	0xac, 0x1b, 0x18, 0x15, 0xd5, 0xe9, 0xcd, 0x7e, 0xad, 0xb6, 0xef, 0xd1, 0x41, 0xa7, 0x4e, 0xb4,
	0xd6, 0x62, 0x17, 0xed, 0xf6, 0x19, 0xee, 0x42, 0xa7, 0xf0, 0x24, 0xf5, 0x02, 0xd4, 0xb0, 0xee,
	0x00, 0xd7, 0xa3, 0xb1, 0xce, 0xde, 0x8a, 0x21, 0x8d, 0x31, 0xf4, 0xca, 0x9b, 0x98, 0x28, 0xd6,
	0x32, 0xee, 0xe4, 0xcd, 0x21, 0xe8, 0x56, 0x6c, 0x45, 0x98, 0x44, 0x4d, 0x8b, 0x1b, 0xee, 0x3e,
	0xd3, 0x69, 0xed, 0x2e, 0x45, 0xfe, 0x54, 0xd3, 0xed, 0xc3, 0xff, 0x8a, 0x09, 0x49, 0xb8, 0x9d,
	0x97, 0x6d, 0x22, 0x26, 0xa9, 0x8d, 0x5a, 0xd6, 0x0a, 0xfa, 0xf5, 0xbe, 0x34, 0x49, 0xd6, 0xf2,
	0xb0, 0xd5, 0xf8, 0x02, 0x50, 0xd9, 0x9b, 0x7c, 0x0c, 0xa8, 0x9a, 0x85, 0x2c, 0xff, 0xd4, 0x7c,
	0xe1, 0x0f, 0xca, 0xfd, 0xd0, 0x23, 0xee, 0xf1, 0x07, 0x9e, 0x68, 0x48, 0xed, 0x89, 0x5b, 0x54,
	0xd5, 0xfa, 0x0a, 0x00, 0x00, 0xff, 0xff, 0xd3, 0x5b, 0x4e, 0x0c, 0x49, 0x02, 0x00, 0x00,
}
