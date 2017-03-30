// Code generated by protoc-gen-go.
// source: zjh_serever.proto
// DO NOT EDIT!

package ddproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Ignoring public import of common_srv_pokerPai from common_server_poker.proto

// 三张牌的类型
type ZjhEnum_ZJHTYPE int32

const (
	ZjhEnum_ZJHTYPE_ZJHTYPE_GAOPAI   ZjhEnum_ZJHTYPE = 1
	ZjhEnum_ZJHTYPE_ZJHTYPE_DUIZI    ZjhEnum_ZJHTYPE = 2
	ZjhEnum_ZJHTYPE_ZJHTYPE_LIANZI   ZjhEnum_ZJHTYPE = 3
	ZjhEnum_ZJHTYPE_ZJHTYPE_QING     ZjhEnum_ZJHTYPE = 4
	ZjhEnum_ZJHTYPE_ZJHTYPE_QINGLIAN ZjhEnum_ZJHTYPE = 5
	ZjhEnum_ZJHTYPE_ZJHTYPE_BAOZI    ZjhEnum_ZJHTYPE = 6
)

var ZjhEnum_ZJHTYPE_name = map[int32]string{
	1: "ZJHTYPE_GAOPAI",
	2: "ZJHTYPE_DUIZI",
	3: "ZJHTYPE_LIANZI",
	4: "ZJHTYPE_QING",
	5: "ZJHTYPE_QINGLIAN",
	6: "ZJHTYPE_BAOZI",
}
var ZjhEnum_ZJHTYPE_value = map[string]int32{
	"ZJHTYPE_GAOPAI":   1,
	"ZJHTYPE_DUIZI":    2,
	"ZJHTYPE_LIANZI":   3,
	"ZJHTYPE_QING":     4,
	"ZJHTYPE_QINGLIAN": 5,
	"ZJHTYPE_BAOZI":    6,
}

func (x ZjhEnum_ZJHTYPE) Enum() *ZjhEnum_ZJHTYPE {
	p := new(ZjhEnum_ZJHTYPE)
	*p = x
	return p
}
func (x ZjhEnum_ZJHTYPE) String() string {
	return proto.EnumName(ZjhEnum_ZJHTYPE_name, int32(x))
}
func (x *ZjhEnum_ZJHTYPE) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ZjhEnum_ZJHTYPE_value, data, "ZjhEnum_ZJHTYPE")
	if err != nil {
		return err
	}
	*x = ZjhEnum_ZJHTYPE(value)
	return nil
}
func (ZjhEnum_ZJHTYPE) EnumDescriptor() ([]byte, []int) { return fileDescriptor30, []int{0} }

// 用户游戏状态
type ZjhEnumUserStatus int32

const (
	ZjhEnumUserStatus_zjh_S_GAMING   ZjhEnumUserStatus = 1
	ZjhEnumUserStatus_zjh_S_STANDING ZjhEnumUserStatus = 2
	ZjhEnumUserStatus_zjh_S_SITED    ZjhEnumUserStatus = 3
	ZjhEnumUserStatus_zjh_S_READY    ZjhEnumUserStatus = 4
	ZjhEnumUserStatus_zjh_S_OFFLINE  ZjhEnumUserStatus = 5
)

var ZjhEnumUserStatus_name = map[int32]string{
	1: "zjh_S_GAMING",
	2: "zjh_S_STANDING",
	3: "zjh_S_SITED",
	4: "zjh_S_READY",
	5: "zjh_S_OFFLINE",
}
var ZjhEnumUserStatus_value = map[string]int32{
	"zjh_S_GAMING":   1,
	"zjh_S_STANDING": 2,
	"zjh_S_SITED":    3,
	"zjh_S_READY":    4,
	"zjh_S_OFFLINE":  5,
}

func (x ZjhEnumUserStatus) Enum() *ZjhEnumUserStatus {
	p := new(ZjhEnumUserStatus)
	*p = x
	return p
}
func (x ZjhEnumUserStatus) String() string {
	return proto.EnumName(ZjhEnumUserStatus_name, int32(x))
}
func (x *ZjhEnumUserStatus) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ZjhEnumUserStatus_value, data, "ZjhEnumUserStatus")
	if err != nil {
		return err
	}
	*x = ZjhEnumUserStatus(value)
	return nil
}
func (ZjhEnumUserStatus) EnumDescriptor() ([]byte, []int) { return fileDescriptor30, []int{1} }

// 打出去的牌
type ZjhSrvPoker struct {
	KeyValue         []int32              `protobuf:"varint,1,rep,name=keyValue" json:"keyValue,omitempty"`
	Pais             []*CommonSrvPokerPai `protobuf:"bytes,2,rep,name=pais" json:"pais,omitempty"`
	Type             *ZjhEnum_ZJHTYPE     `protobuf:"varint,3,opt,name=type,enum=ddproto.ZjhEnum_ZJHTYPE" json:"type,omitempty"`
	UserId           *uint32              `protobuf:"varint,9,opt,name=userId" json:"userId,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *ZjhSrvPoker) Reset()                    { *m = ZjhSrvPoker{} }
func (m *ZjhSrvPoker) String() string            { return proto.CompactTextString(m) }
func (*ZjhSrvPoker) ProtoMessage()               {}
func (*ZjhSrvPoker) Descriptor() ([]byte, []int) { return fileDescriptor30, []int{0} }

func (m *ZjhSrvPoker) GetKeyValue() []int32 {
	if m != nil {
		return m.KeyValue
	}
	return nil
}

func (m *ZjhSrvPoker) GetPais() []*CommonSrvPokerPai {
	if m != nil {
		return m.Pais
	}
	return nil
}

func (m *ZjhSrvPoker) GetType() ZjhEnum_ZJHTYPE {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return ZjhEnum_ZJHTYPE_ZJHTYPE_GAOPAI
}

func (m *ZjhSrvPoker) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

// 账单
type ZjhSrvBill struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *ZjhSrvBill) Reset()                    { *m = ZjhSrvBill{} }
func (m *ZjhSrvBill) String() string            { return proto.CompactTextString(m) }
func (*ZjhSrvBill) ProtoMessage()               {}
func (*ZjhSrvBill) Descriptor() ([]byte, []int) { return fileDescriptor30, []int{1} }

// 用户的游戏数据
type ZjhSrv_GameData struct {
	Pai              *ZjhSrvPoker `protobuf:"bytes,1,opt,name=pai" json:"pai,omitempty"`
	Bill             *int64       `protobuf:"varint,2,opt,name=bill" json:"bill,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *ZjhSrv_GameData) Reset()                    { *m = ZjhSrv_GameData{} }
func (m *ZjhSrv_GameData) String() string            { return proto.CompactTextString(m) }
func (*ZjhSrv_GameData) ProtoMessage()               {}
func (*ZjhSrv_GameData) Descriptor() ([]byte, []int) { return fileDescriptor30, []int{2} }

func (m *ZjhSrv_GameData) GetPai() *ZjhSrvPoker {
	if m != nil {
		return m.Pai
	}
	return nil
}

func (m *ZjhSrv_GameData) GetBill() int64 {
	if m != nil && m.Bill != nil {
		return *m.Bill
	}
	return 0
}

// 用户属性
type ZjhSrvUser struct {
	UserId           *uint32            `protobuf:"varint,1,opt,name=userId" json:"userId,omitempty"`
	Coin             *int64             `protobuf:"varint,2,opt,name=coin" json:"coin,omitempty"`
	RoomId           *int32             `protobuf:"varint,3,opt,name=roomId" json:"roomId,omitempty"`
	DeskId           *int32             `protobuf:"varint,4,opt,name=deskId" json:"deskId,omitempty"`
	GameNumber       *int32             `protobuf:"varint,5,opt,name=gameNumber" json:"gameNumber,omitempty"`
	Data             *ZjhSrv_GameData   `protobuf:"bytes,6,opt,name=data" json:"data,omitempty"`
	Status           *ZjhEnumUserStatus `protobuf:"varint,7,opt,name=status,enum=ddproto.ZjhEnumUserStatus" json:"status,omitempty"`
	IsDiscard        *bool              `protobuf:"varint,8,opt,name=isDiscard" json:"isDiscard,omitempty"`
	IsWatch          *bool              `protobuf:"varint,9,opt,name=isWatch" json:"isWatch,omitempty"`
	IsOnline         *bool              `protobuf:"varint,10,opt,name=isOnline" json:"isOnline,omitempty"`
	IsLost           *bool              `protobuf:"varint,11,opt,name=isLost" json:"isLost,omitempty"`
	Index            *int32             `protobuf:"varint,12,opt,name=index" json:"index,omitempty"`
	DashangNum       *int32             `protobuf:"varint,13,opt,name=dashangNum" json:"dashangNum,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *ZjhSrvUser) Reset()                    { *m = ZjhSrvUser{} }
func (m *ZjhSrvUser) String() string            { return proto.CompactTextString(m) }
func (*ZjhSrvUser) ProtoMessage()               {}
func (*ZjhSrvUser) Descriptor() ([]byte, []int) { return fileDescriptor30, []int{3} }

func (m *ZjhSrvUser) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *ZjhSrvUser) GetCoin() int64 {
	if m != nil && m.Coin != nil {
		return *m.Coin
	}
	return 0
}

func (m *ZjhSrvUser) GetRoomId() int32 {
	if m != nil && m.RoomId != nil {
		return *m.RoomId
	}
	return 0
}

func (m *ZjhSrvUser) GetDeskId() int32 {
	if m != nil && m.DeskId != nil {
		return *m.DeskId
	}
	return 0
}

func (m *ZjhSrvUser) GetGameNumber() int32 {
	if m != nil && m.GameNumber != nil {
		return *m.GameNumber
	}
	return 0
}

func (m *ZjhSrvUser) GetData() *ZjhSrv_GameData {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *ZjhSrvUser) GetStatus() ZjhEnumUserStatus {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return ZjhEnumUserStatus_zjh_S_GAMING
}

func (m *ZjhSrvUser) GetIsDiscard() bool {
	if m != nil && m.IsDiscard != nil {
		return *m.IsDiscard
	}
	return false
}

func (m *ZjhSrvUser) GetIsWatch() bool {
	if m != nil && m.IsWatch != nil {
		return *m.IsWatch
	}
	return false
}

func (m *ZjhSrvUser) GetIsOnline() bool {
	if m != nil && m.IsOnline != nil {
		return *m.IsOnline
	}
	return false
}

func (m *ZjhSrvUser) GetIsLost() bool {
	if m != nil && m.IsLost != nil {
		return *m.IsLost
	}
	return false
}

func (m *ZjhSrvUser) GetIndex() int32 {
	if m != nil && m.Index != nil {
		return *m.Index
	}
	return 0
}

func (m *ZjhSrvUser) GetDashangNum() int32 {
	if m != nil && m.DashangNum != nil {
		return *m.DashangNum
	}
	return 0
}

// desk 的信息
type ZjhSrvDesk struct {
	DeskId           *int32               `protobuf:"varint,1,opt,name=deskId" json:"deskId,omitempty"`
	GameNumber       *int32               `protobuf:"varint,2,opt,name=gameNumber" json:"gameNumber,omitempty"`
	RoomId           *int32               `protobuf:"varint,3,opt,name=roomId" json:"roomId,omitempty"`
	AllPokers        []*CommonSrvPokerPai `protobuf:"bytes,4,rep,name=allPokers" json:"allPokers,omitempty"`
	LastUser         *uint32              `protobuf:"varint,5,opt,name=lastUser" json:"lastUser,omitempty"`
	LastWiner        *uint32              `protobuf:"varint,6,opt,name=lastWiner" json:"lastWiner,omitempty"`
	IsGamming        *bool                `protobuf:"varint,7,opt,name=isGamming" json:"isGamming,omitempty"`
	CuurBaseValue    *int64               `protobuf:"varint,8,opt,name=cuurBaseValue" json:"cuurBaseValue,omitempty"`
	MinUser          *int32               `protobuf:"varint,9,opt,name=minUser" json:"minUser,omitempty"`
	BaseValue        *int64               `protobuf:"varint,10,opt,name=baseValue" json:"baseValue,omitempty"`
	MaxValue         *int64               `protobuf:"varint,11,opt,name=maxValue" json:"maxValue,omitempty"`
	GameStatus       *int32               `protobuf:"varint,12,opt,name=gameStatus" json:"gameStatus,omitempty"`
	GamerNum         *int32               `protobuf:"varint,13,opt,name=gamerNum" json:"gamerNum,omitempty"`
	CircleNo         *int32               `protobuf:"varint,14,opt,name=circleNo" json:"circleNo,omitempty"`
	XuepinBaseValue  *int64               `protobuf:"varint,15,opt,name=xuepinBaseValue" json:"xuepinBaseValue,omitempty"`
	CoinCount        *int64               `protobuf:"varint,16,opt,name=coinCount" json:"coinCount,omitempty"`
	XuepinStartIndex *int32               `protobuf:"varint,17,opt,name=xuepinStartIndex" json:"xuepinStartIndex,omitempty"`
	XuepinChipCount  *int32               `protobuf:"varint,18,opt,name=xuepinChipCount" json:"xuepinChipCount,omitempty"`
	MinCoin          *int64               `protobuf:"varint,19,opt,name=minCoin" json:"minCoin,omitempty"`
	DaShangChip      *int64               `protobuf:"varint,20,opt,name=daShangChip" json:"daShangChip,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *ZjhSrvDesk) Reset()                    { *m = ZjhSrvDesk{} }
func (m *ZjhSrvDesk) String() string            { return proto.CompactTextString(m) }
func (*ZjhSrvDesk) ProtoMessage()               {}
func (*ZjhSrvDesk) Descriptor() ([]byte, []int) { return fileDescriptor30, []int{4} }

func (m *ZjhSrvDesk) GetDeskId() int32 {
	if m != nil && m.DeskId != nil {
		return *m.DeskId
	}
	return 0
}

func (m *ZjhSrvDesk) GetGameNumber() int32 {
	if m != nil && m.GameNumber != nil {
		return *m.GameNumber
	}
	return 0
}

func (m *ZjhSrvDesk) GetRoomId() int32 {
	if m != nil && m.RoomId != nil {
		return *m.RoomId
	}
	return 0
}

func (m *ZjhSrvDesk) GetAllPokers() []*CommonSrvPokerPai {
	if m != nil {
		return m.AllPokers
	}
	return nil
}

func (m *ZjhSrvDesk) GetLastUser() uint32 {
	if m != nil && m.LastUser != nil {
		return *m.LastUser
	}
	return 0
}

func (m *ZjhSrvDesk) GetLastWiner() uint32 {
	if m != nil && m.LastWiner != nil {
		return *m.LastWiner
	}
	return 0
}

func (m *ZjhSrvDesk) GetIsGamming() bool {
	if m != nil && m.IsGamming != nil {
		return *m.IsGamming
	}
	return false
}

func (m *ZjhSrvDesk) GetCuurBaseValue() int64 {
	if m != nil && m.CuurBaseValue != nil {
		return *m.CuurBaseValue
	}
	return 0
}

func (m *ZjhSrvDesk) GetMinUser() int32 {
	if m != nil && m.MinUser != nil {
		return *m.MinUser
	}
	return 0
}

func (m *ZjhSrvDesk) GetBaseValue() int64 {
	if m != nil && m.BaseValue != nil {
		return *m.BaseValue
	}
	return 0
}

func (m *ZjhSrvDesk) GetMaxValue() int64 {
	if m != nil && m.MaxValue != nil {
		return *m.MaxValue
	}
	return 0
}

func (m *ZjhSrvDesk) GetGameStatus() int32 {
	if m != nil && m.GameStatus != nil {
		return *m.GameStatus
	}
	return 0
}

func (m *ZjhSrvDesk) GetGamerNum() int32 {
	if m != nil && m.GamerNum != nil {
		return *m.GamerNum
	}
	return 0
}

func (m *ZjhSrvDesk) GetCircleNo() int32 {
	if m != nil && m.CircleNo != nil {
		return *m.CircleNo
	}
	return 0
}

func (m *ZjhSrvDesk) GetXuepinBaseValue() int64 {
	if m != nil && m.XuepinBaseValue != nil {
		return *m.XuepinBaseValue
	}
	return 0
}

func (m *ZjhSrvDesk) GetCoinCount() int64 {
	if m != nil && m.CoinCount != nil {
		return *m.CoinCount
	}
	return 0
}

func (m *ZjhSrvDesk) GetXuepinStartIndex() int32 {
	if m != nil && m.XuepinStartIndex != nil {
		return *m.XuepinStartIndex
	}
	return 0
}

func (m *ZjhSrvDesk) GetXuepinChipCount() int32 {
	if m != nil && m.XuepinChipCount != nil {
		return *m.XuepinChipCount
	}
	return 0
}

func (m *ZjhSrvDesk) GetMinCoin() int64 {
	if m != nil && m.MinCoin != nil {
		return *m.MinCoin
	}
	return 0
}

func (m *ZjhSrvDesk) GetDaShangChip() int64 {
	if m != nil && m.DaShangChip != nil {
		return *m.DaShangChip
	}
	return 0
}

// room 的信息
type ZjhSrvRoom struct {
	RoomId           *int32  `protobuf:"varint,1,opt,name=roomId" json:"roomId,omitempty"`
	RoomType         *int32  `protobuf:"varint,2,opt,name=roomType" json:"roomType,omitempty"`
	RoomLevel        *int32  `protobuf:"varint,3,opt,name=roomLevel" json:"roomLevel,omitempty"`
	RoomTitle        *string `protobuf:"bytes,4,opt,name=roomTitle" json:"roomTitle,omitempty"`
	BaseValue        *int64  `protobuf:"varint,5,opt,name=baseValue" json:"baseValue,omitempty"`
	MaxValue         *int64  `protobuf:"varint,6,opt,name=maxValue" json:"maxValue,omitempty"`
	XuepinBaseValue  *int64  `protobuf:"varint,7,opt,name=xuepinBaseValue" json:"xuepinBaseValue,omitempty"`
	MinCoin          *int64  `protobuf:"varint,8,opt,name=minCoin" json:"minCoin,omitempty"`
	DaShangChip      *int64  `protobuf:"varint,9,opt,name=daShangChip" json:"daShangChip,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ZjhSrvRoom) Reset()                    { *m = ZjhSrvRoom{} }
func (m *ZjhSrvRoom) String() string            { return proto.CompactTextString(m) }
func (*ZjhSrvRoom) ProtoMessage()               {}
func (*ZjhSrvRoom) Descriptor() ([]byte, []int) { return fileDescriptor30, []int{5} }

func (m *ZjhSrvRoom) GetRoomId() int32 {
	if m != nil && m.RoomId != nil {
		return *m.RoomId
	}
	return 0
}

func (m *ZjhSrvRoom) GetRoomType() int32 {
	if m != nil && m.RoomType != nil {
		return *m.RoomType
	}
	return 0
}

func (m *ZjhSrvRoom) GetRoomLevel() int32 {
	if m != nil && m.RoomLevel != nil {
		return *m.RoomLevel
	}
	return 0
}

func (m *ZjhSrvRoom) GetRoomTitle() string {
	if m != nil && m.RoomTitle != nil {
		return *m.RoomTitle
	}
	return ""
}

func (m *ZjhSrvRoom) GetBaseValue() int64 {
	if m != nil && m.BaseValue != nil {
		return *m.BaseValue
	}
	return 0
}

func (m *ZjhSrvRoom) GetMaxValue() int64 {
	if m != nil && m.MaxValue != nil {
		return *m.MaxValue
	}
	return 0
}

func (m *ZjhSrvRoom) GetXuepinBaseValue() int64 {
	if m != nil && m.XuepinBaseValue != nil {
		return *m.XuepinBaseValue
	}
	return 0
}

func (m *ZjhSrvRoom) GetMinCoin() int64 {
	if m != nil && m.MinCoin != nil {
		return *m.MinCoin
	}
	return 0
}

func (m *ZjhSrvRoom) GetDaShangChip() int64 {
	if m != nil && m.DaShangChip != nil {
		return *m.DaShangChip
	}
	return 0
}

func init() {
	proto.RegisterType((*ZjhSrvPoker)(nil), "ddproto.zjh_srv_poker")
	proto.RegisterType((*ZjhSrvBill)(nil), "ddproto.zjh_srv_bill")
	proto.RegisterType((*ZjhSrv_GameData)(nil), "ddproto.zjh_srv_GameData")
	proto.RegisterType((*ZjhSrvUser)(nil), "ddproto.zjh_srv_user")
	proto.RegisterType((*ZjhSrvDesk)(nil), "ddproto.zjh_srv_desk")
	proto.RegisterType((*ZjhSrvRoom)(nil), "ddproto.zjh_srv_room")
	proto.RegisterEnum("ddproto.ZjhEnum_ZJHTYPE", ZjhEnum_ZJHTYPE_name, ZjhEnum_ZJHTYPE_value)
	proto.RegisterEnum("ddproto.ZjhEnumUserStatus", ZjhEnumUserStatus_name, ZjhEnumUserStatus_value)
}

func init() { proto.RegisterFile("zjh_serever.proto", fileDescriptor30) }

var fileDescriptor30 = []byte{
	// 751 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x54, 0xcd, 0x6e, 0xf2, 0x46,
	0x14, 0xad, 0xb1, 0x01, 0x73, 0xc1, 0x30, 0x38, 0x69, 0x3b, 0xa9, 0xb2, 0x40, 0x74, 0x11, 0x14,
	0x55, 0xa9, 0x94, 0x37, 0x20, 0x81, 0x50, 0x57, 0x14, 0x68, 0x21, 0x8d, 0xc2, 0x06, 0x4d, 0xf0,
	0x28, 0x4c, 0xe3, 0x3f, 0x79, 0x6c, 0x94, 0xf4, 0x01, 0xaa, 0x3e, 0x46, 0x5f, 0xa2, 0x9b, 0x3e,
	0x5d, 0x75, 0xc7, 0x36, 0x3f, 0x0a, 0x8b, 0x6f, 0xe7, 0x39, 0x73, 0xe7, 0xcc, 0xb9, 0xe7, 0xdc,
	0x31, 0xb4, 0xff, 0xfc, 0x63, 0xb3, 0x92, 0x3c, 0xe6, 0x5b, 0x1e, 0xdf, 0x44, 0x71, 0x98, 0x84,
	0x76, 0xd5, 0x75, 0xd5, 0xc7, 0x77, 0x17, 0xeb, 0xd0, 0xf7, 0xc3, 0x00, 0xb7, 0xb7, 0x3c, 0x5e,
	0x45, 0xe1, 0x5b, 0x51, 0xd3, 0xfd, 0x4b, 0x03, 0x4b, 0x9d, 0x8c, 0xb7, 0x19, 0x6e, 0x13, 0x30,
	0xdf, 0xf8, 0xc7, 0xef, 0xcc, 0x4b, 0x39, 0xd5, 0x3a, 0x7a, 0xaf, 0x6c, 0x5f, 0x83, 0x11, 0x31,
	0x21, 0x69, 0xa9, 0xa3, 0xf7, 0xea, 0xb7, 0x97, 0x37, 0x39, 0xed, 0x4d, 0xc1, 0x5a, 0x1c, 0x9d,
	0x31, 0x61, 0x5f, 0x81, 0x91, 0x7c, 0x44, 0x9c, 0xea, 0x1d, 0xad, 0xd7, 0xbc, 0xbd, 0xd8, 0xd5,
	0xe2, 0x1d, 0x3c, 0x48, 0xfd, 0xd5, 0xf2, 0xe7, 0x9f, 0x16, 0xcf, 0xb3, 0xa1, 0xdd, 0x84, 0x4a,
	0x2a, 0x79, 0xec, 0xb8, 0xb4, 0xd6, 0xd1, 0x7a, 0x56, 0xb7, 0x09, 0x8d, 0x42, 0xc7, 0x8b, 0xf0,
	0xbc, 0xee, 0x10, 0x48, 0xb1, 0x1e, 0x31, 0x9f, 0x0f, 0x58, 0xc2, 0xec, 0xef, 0x41, 0x8f, 0x98,
	0xa0, 0x5a, 0x47, 0xeb, 0xd5, 0x6f, 0xbf, 0x39, 0xe2, 0xde, 0xeb, 0x6f, 0x80, 0x81, 0x04, 0xb4,
	0xd4, 0xd1, 0x7a, 0x7a, 0xf7, 0x9f, 0xd2, 0x9e, 0x17, 0xef, 0x3b, 0xb8, 0x17, 0x69, 0x2c, 0x2c,
	0x5f, 0x87, 0x22, 0xc8, 0xca, 0x71, 0x37, 0x0e, 0x43, 0xdf, 0x71, 0x55, 0x03, 0x65, 0x5c, 0xbb,
	0x5c, 0xbe, 0x39, 0x2e, 0x35, 0xd4, 0xda, 0x06, 0x78, 0x65, 0x3e, 0x9f, 0xa4, 0xfe, 0x0b, 0x8f,
	0x69, 0x59, 0x61, 0x57, 0x60, 0xb8, 0x2c, 0x61, 0xb4, 0xa2, 0x64, 0x5d, 0x7c, 0x92, 0xb5, 0x93,
	0xff, 0x03, 0x54, 0x64, 0xc2, 0x92, 0x54, 0xd2, 0xaa, 0x72, 0xe7, 0xf2, 0xb3, 0x3b, 0x28, 0x6d,
	0xae, 0x6a, 0xec, 0x36, 0xd4, 0x84, 0x1c, 0x08, 0xb9, 0x66, 0xb1, 0x4b, 0xcd, 0x8e, 0xd6, 0x33,
	0xed, 0x16, 0x54, 0x85, 0x7c, 0x62, 0xc9, 0x7a, 0xa3, 0x4c, 0x33, 0x31, 0x2b, 0x21, 0xa7, 0x81,
	0x27, 0x02, 0x4e, 0x41, 0x21, 0x4d, 0xa8, 0x08, 0x39, 0x0e, 0x65, 0x42, 0xeb, 0x6a, 0x6d, 0x41,
	0x59, 0x04, 0x2e, 0x7f, 0xa7, 0x8d, 0x42, 0xbf, 0xcb, 0xe4, 0x86, 0x05, 0xaf, 0x93, 0xd4, 0xa7,
	0x16, 0x62, 0xdd, 0x7f, 0xf5, 0xbd, 0x45, 0xd8, 0xec, 0x41, 0xd3, 0xda, 0x89, 0xa6, 0x4b, 0x85,
	0x31, 0x47, 0x46, 0xfd, 0x08, 0x35, 0xe6, 0x79, 0x33, 0x4c, 0x40, 0x52, 0xe3, 0x0b, 0x06, 0x85,
	0x80, 0xe9, 0x31, 0x99, 0x3c, 0xca, 0xdc, 0x47, 0x0b, 0x1b, 0x46, 0xe4, 0x49, 0x04, 0x3c, 0x56,
	0x66, 0x5a, 0x99, 0x07, 0x23, 0xe6, 0xfb, 0x22, 0x78, 0x55, 0xa6, 0x99, 0xf6, 0xd7, 0x60, 0xad,
	0xd3, 0x34, 0xbe, 0x63, 0x92, 0x67, 0x33, 0x6a, 0xaa, 0xe0, 0x5a, 0x50, 0xf5, 0x45, 0xa0, 0xd8,
	0x6a, 0x4a, 0x50, 0x1b, 0x6a, 0x2f, 0xbb, 0x1a, 0x50, 0x35, 0x04, 0x4c, 0x9f, 0xbd, 0x67, 0x48,
	0x5d, 0x21, 0x79, 0x67, 0x99, 0xe3, 0xb9, 0x45, 0x04, 0x4c, 0xc4, 0xe2, 0x9d, 0x41, 0x88, 0xac,
	0x45, 0xbc, 0xf6, 0xf8, 0x24, 0xa4, 0x4d, 0x85, 0x7c, 0x0b, 0xad, 0xf7, 0x94, 0x47, 0x22, 0xd8,
	0xcb, 0x68, 0x29, 0xc2, 0x36, 0xd4, 0x70, 0x9a, 0xee, 0xc3, 0x34, 0x48, 0x28, 0x51, 0x10, 0x05,
	0x92, 0xd5, 0xce, 0x13, 0x16, 0x27, 0x8e, 0x0a, 0xa3, 0x7d, 0xcc, 0x72, 0xbf, 0x11, 0x51, 0x76,
	0xc4, 0x56, 0x1b, 0x59, 0x33, 0xf7, 0x38, 0x96, 0x67, 0x8a, 0xe3, 0x0c, 0xea, 0x2e, 0x9b, 0x63,
	0x6c, 0x58, 0x4a, 0xcf, 0xd5, 0x68, 0xff, 0xa7, 0xed, 0x73, 0xc3, 0x2c, 0x0e, 0x32, 0xd1, 0x0a,
	0xdd, 0xb8, 0x5e, 0xe0, 0x7b, 0x2c, 0x15, 0xa6, 0x20, 0x32, 0xe6, 0x5b, 0xee, 0xe5, 0xc1, 0xe5,
	0xd0, 0x42, 0x24, 0x1e, 0x57, 0x43, 0x5e, 0x3b, 0xb6, 0xae, 0xfc, 0xc9, 0xba, 0x8a, 0x42, 0x4e,
	0x58, 0x50, 0x3d, 0x48, 0x42, 0x89, 0x37, 0x4f, 0x89, 0xc7, 0x78, 0xf4, 0xeb, 0xbf, 0xb5, 0xec,
	0x7d, 0x1f, 0xfd, 0x13, 0x6c, 0x68, 0xe6, 0x9f, 0xab, 0x51, 0x7f, 0x3a, 0xeb, 0x3b, 0x44, 0xb3,
	0xdb, 0x60, 0x15, 0xd8, 0xe0, 0xd1, 0x59, 0x3a, 0xa4, 0x74, 0x58, 0x36, 0x76, 0xfa, 0x93, 0xa5,
	0x43, 0x50, 0x60, 0xa3, 0xc0, 0x7e, 0x75, 0x26, 0x23, 0x62, 0xd8, 0xe7, 0x40, 0x0e, 0x11, 0xac,
	0x24, 0xe5, 0x43, 0xba, 0xbb, 0xfe, 0x74, 0xe9, 0x90, 0xca, 0x75, 0x04, 0x67, 0xa7, 0xde, 0x1f,
	0xc9, 0xdc, 0x9d, 0xaf, 0x46, 0xfd, 0x5f, 0x90, 0x51, 0xc3, 0x7b, 0x33, 0x64, 0xbe, 0xe8, 0x4f,
	0x06, 0x88, 0x95, 0xec, 0x16, 0xd4, 0x73, 0xcc, 0x59, 0x0c, 0x07, 0x44, 0xdf, 0x03, 0xbf, 0x0d,
	0xfb, 0x83, 0x67, 0x62, 0xe0, 0x8d, 0x19, 0x30, 0x7d, 0x78, 0x18, 0x3b, 0x93, 0x21, 0x29, 0xcf,
	0xbe, 0xfa, 0x3f, 0x00, 0x00, 0xff, 0xff, 0x5c, 0x33, 0x91, 0x1f, 0xae, 0x05, 0x00, 0x00,
}
