// Code generated by protoc-gen-go. DO NOT EDIT.
// source: zhadan_server.proto

package ddproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Ignoring public import of common_srv_pokerPai from common_server_poker.proto

// Ignoring public import of zhadan_client_poker from zhadan_base.proto

// Ignoring public import of zhadan_user_bill from zhadan_base.proto

// Ignoring public import of zhadan_desk_option from zhadan_base.proto

// Ignoring public import of zhadan_srv_room from zhadan_base.proto

// Ignoring public import of zhadan_enum_protoid from zhadan_base.proto

// Ignoring public import of zhadan_enum_pokerType from zhadan_base.proto

// Ignoring public import of zhadan_enum_desk_status from zhadan_base.proto

// Ignoring public import of zhadan_gan_opt from zhadan_base.proto

// 打出去的牌
type ZhadanSrvPoker struct {
	Pais             []*CommonSrvPokerPai `protobuf:"bytes,1,rep,name=pais" json:"pais,omitempty"`
	Type             *ZhadanEnumPokerType `protobuf:"varint,2,opt,name=Type,enum=ddproto.ZhadanEnumPokerType" json:"Type,omitempty"`
	KeyValue         *int32               `protobuf:"varint,3,opt,name=keyValue" json:"keyValue,omitempty"`
	KeyLength        *int32               `protobuf:"varint,4,opt,name=keyLength" json:"keyLength,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *ZhadanSrvPoker) Reset()                    { *m = ZhadanSrvPoker{} }
func (m *ZhadanSrvPoker) String() string            { return proto.CompactTextString(m) }
func (*ZhadanSrvPoker) ProtoMessage()               {}
func (*ZhadanSrvPoker) Descriptor() ([]byte, []int) { return fileDescriptor58, []int{0} }

func (m *ZhadanSrvPoker) GetPais() []*CommonSrvPokerPai {
	if m != nil {
		return m.Pais
	}
	return nil
}

func (m *ZhadanSrvPoker) GetType() ZhadanEnumPokerType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return ZhadanEnumPokerType_ZHADAN_POKER_TYPE_OTHER
}

func (m *ZhadanSrvPoker) GetKeyValue() int32 {
	if m != nil && m.KeyValue != nil {
		return *m.KeyValue
	}
	return 0
}

func (m *ZhadanSrvPoker) GetKeyLength() int32 {
	if m != nil && m.KeyLength != nil {
		return *m.KeyLength
	}
	return 0
}

// desk 的信息
type ZhadanSrvDesk struct {
	DeskId         *int32                `protobuf:"varint,1,opt,name=deskId" json:"deskId,omitempty"`
	Pwd            *string               `protobuf:"bytes,2,opt,name=pwd" json:"pwd,omitempty"`
	GameNumber     *int32                `protobuf:"varint,3,opt,name=gameNumber" json:"gameNumber,omitempty"`
	RoomId         *int32                `protobuf:"varint,4,opt,name=roomId" json:"roomId,omitempty"`
	Status         *ZhadanEnumDeskStatus `protobuf:"varint,5,opt,name=status,enum=ddproto.ZhadanEnumDeskStatus" json:"status,omitempty"`
	CircleNo       *int32                `protobuf:"varint,7,opt,name=circleNo" json:"circleNo,omitempty"`
	CurrDeskScore  *int32                `protobuf:"varint,8,opt,name=currDeskScore" json:"currDeskScore,omitempty"`
	DeskScorePais  *ZhadanSrvPoker       `protobuf:"bytes,29,opt,name=deskScorePais" json:"deskScorePais,omitempty"`
	Owner          *uint32               `protobuf:"varint,11,opt,name=owner" json:"owner,omitempty"`
	IsDaikai       *bool                 `protobuf:"varint,12,opt,name=isDaikai" json:"isDaikai,omitempty"`
	DaikaiUser     *uint32               `protobuf:"varint,13,opt,name=daikaiUser" json:"daikaiUser,omitempty"`
	DeskOption     *ZhadanDeskOption     `protobuf:"bytes,14,opt,name=deskOption" json:"deskOption,omitempty"`
	LastActUser    *uint32               `protobuf:"varint,15,opt,name=lastActUser" json:"lastActUser,omitempty"`
	LastChupaiUser *uint32               `protobuf:"varint,16,opt,name=lastChupaiUser" json:"lastChupaiUser,omitempty"`
	CurrActUser    *uint32               `protobuf:"varint,26,opt,name=currActUser" json:"currActUser,omitempty"`
	CanGuo         *bool                 `protobuf:"varint,27,opt,name=canGuo" json:"canGuo,omitempty"`
	BaoZhuangUser  *uint32               `protobuf:"varint,28,opt,name=baoZhuangUser" json:"baoZhuangUser,omitempty"`
	IsOnDissolve   *bool                 `protobuf:"varint,18,opt,name=isOnDissolve" json:"isOnDissolve,omitempty"`
	DissolveTime   *int64                `protobuf:"varint,19,opt,name=dissolveTime" json:"dissolveTime,omitempty"`
	IsStart        *bool                 `protobuf:"varint,22,opt,name=isStart" json:"isStart,omitempty"`
	OneStartTime   *int64                `protobuf:"varint,23,opt,name=oneStartTime" json:"oneStartTime,omitempty"`
	AllStartTime   *int64                `protobuf:"varint,24,opt,name=allStartTime" json:"allStartTime,omitempty"`
	DissolveUser   *uint32               `protobuf:"varint,25,opt,name=dissolveUser" json:"dissolveUser,omitempty"`
	// 金币场
	IsCoinRoom       *bool  `protobuf:"varint,20,opt,name=isCoinRoom" json:"isCoinRoom,omitempty"`
	SurplusTime      *int32 `protobuf:"varint,21,opt,name=surplusTime" json:"surplusTime,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *ZhadanSrvDesk) Reset()                    { *m = ZhadanSrvDesk{} }
func (m *ZhadanSrvDesk) String() string            { return proto.CompactTextString(m) }
func (*ZhadanSrvDesk) ProtoMessage()               {}
func (*ZhadanSrvDesk) Descriptor() ([]byte, []int) { return fileDescriptor58, []int{1} }

func (m *ZhadanSrvDesk) GetDeskId() int32 {
	if m != nil && m.DeskId != nil {
		return *m.DeskId
	}
	return 0
}

func (m *ZhadanSrvDesk) GetPwd() string {
	if m != nil && m.Pwd != nil {
		return *m.Pwd
	}
	return ""
}

func (m *ZhadanSrvDesk) GetGameNumber() int32 {
	if m != nil && m.GameNumber != nil {
		return *m.GameNumber
	}
	return 0
}

func (m *ZhadanSrvDesk) GetRoomId() int32 {
	if m != nil && m.RoomId != nil {
		return *m.RoomId
	}
	return 0
}

func (m *ZhadanSrvDesk) GetStatus() ZhadanEnumDeskStatus {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return ZhadanEnumDeskStatus_ZHADAN_DESK_STATUS_WAIT_READY
}

func (m *ZhadanSrvDesk) GetCircleNo() int32 {
	if m != nil && m.CircleNo != nil {
		return *m.CircleNo
	}
	return 0
}

func (m *ZhadanSrvDesk) GetCurrDeskScore() int32 {
	if m != nil && m.CurrDeskScore != nil {
		return *m.CurrDeskScore
	}
	return 0
}

func (m *ZhadanSrvDesk) GetDeskScorePais() *ZhadanSrvPoker {
	if m != nil {
		return m.DeskScorePais
	}
	return nil
}

func (m *ZhadanSrvDesk) GetOwner() uint32 {
	if m != nil && m.Owner != nil {
		return *m.Owner
	}
	return 0
}

func (m *ZhadanSrvDesk) GetIsDaikai() bool {
	if m != nil && m.IsDaikai != nil {
		return *m.IsDaikai
	}
	return false
}

func (m *ZhadanSrvDesk) GetDaikaiUser() uint32 {
	if m != nil && m.DaikaiUser != nil {
		return *m.DaikaiUser
	}
	return 0
}

func (m *ZhadanSrvDesk) GetDeskOption() *ZhadanDeskOption {
	if m != nil {
		return m.DeskOption
	}
	return nil
}

func (m *ZhadanSrvDesk) GetLastActUser() uint32 {
	if m != nil && m.LastActUser != nil {
		return *m.LastActUser
	}
	return 0
}

func (m *ZhadanSrvDesk) GetLastChupaiUser() uint32 {
	if m != nil && m.LastChupaiUser != nil {
		return *m.LastChupaiUser
	}
	return 0
}

func (m *ZhadanSrvDesk) GetCurrActUser() uint32 {
	if m != nil && m.CurrActUser != nil {
		return *m.CurrActUser
	}
	return 0
}

func (m *ZhadanSrvDesk) GetCanGuo() bool {
	if m != nil && m.CanGuo != nil {
		return *m.CanGuo
	}
	return false
}

func (m *ZhadanSrvDesk) GetBaoZhuangUser() uint32 {
	if m != nil && m.BaoZhuangUser != nil {
		return *m.BaoZhuangUser
	}
	return 0
}

func (m *ZhadanSrvDesk) GetIsOnDissolve() bool {
	if m != nil && m.IsOnDissolve != nil {
		return *m.IsOnDissolve
	}
	return false
}

func (m *ZhadanSrvDesk) GetDissolveTime() int64 {
	if m != nil && m.DissolveTime != nil {
		return *m.DissolveTime
	}
	return 0
}

func (m *ZhadanSrvDesk) GetIsStart() bool {
	if m != nil && m.IsStart != nil {
		return *m.IsStart
	}
	return false
}

func (m *ZhadanSrvDesk) GetOneStartTime() int64 {
	if m != nil && m.OneStartTime != nil {
		return *m.OneStartTime
	}
	return 0
}

func (m *ZhadanSrvDesk) GetAllStartTime() int64 {
	if m != nil && m.AllStartTime != nil {
		return *m.AllStartTime
	}
	return 0
}

func (m *ZhadanSrvDesk) GetDissolveUser() uint32 {
	if m != nil && m.DissolveUser != nil {
		return *m.DissolveUser
	}
	return 0
}

func (m *ZhadanSrvDesk) GetIsCoinRoom() bool {
	if m != nil && m.IsCoinRoom != nil {
		return *m.IsCoinRoom
	}
	return false
}

func (m *ZhadanSrvDesk) GetSurplusTime() int32 {
	if m != nil && m.SurplusTime != nil {
		return *m.SurplusTime
	}
	return 0
}

// 用户属性
type ZhadanSrvUser struct {
	UserId           *uint32           `protobuf:"varint,1,opt,name=userId" json:"userId,omitempty"`
	Bill             *ZhadanUserBill   `protobuf:"bytes,2,opt,name=bill" json:"bill,omitempty"`
	IsOnline         *bool             `protobuf:"varint,3,opt,name=isOnline" json:"isOnline,omitempty"`
	Index            *int32            `protobuf:"varint,4,opt,name=index" json:"index,omitempty"`
	Pokers           *ZhadanSrvPoker   `protobuf:"bytes,5,opt,name=pokers" json:"pokers,omitempty"`
	OutPai           *ZhadanSrvPoker   `protobuf:"bytes,6,opt,name=out_pai,json=outPai" json:"out_pai,omitempty"`
	IsPass           *bool             `protobuf:"varint,7,opt,name=isPass" json:"isPass,omitempty"`
	DeskScore        *int32            `protobuf:"varint,16,opt,name=deskScore" json:"deskScore,omitempty"`
	IsBanker         *bool             `protobuf:"varint,17,opt,name=isBanker" json:"isBanker,omitempty"`
	BaozhuangStatus  *int32            `protobuf:"varint,18,opt,name=baozhuangStatus" json:"baozhuangStatus,omitempty"`
	ScorePais        *ZhadanSrvPoker   `protobuf:"bytes,22,opt,name=score_pais,json=scorePais" json:"score_pais,omitempty"`
	ExtScorePoker    []*ZhadanSrvPoker `protobuf:"bytes,23,rep,name=ext_score_poker,json=extScorePoker" json:"ext_score_poker,omitempty"`
	ExtScore         *int32            `protobuf:"varint,24,opt,name=ext_score,json=extScore" json:"ext_score,omitempty"`
	PaoScore         *int32            `protobuf:"varint,25,opt,name=paoScore" json:"paoScore,omitempty"`
	IsReady          *bool             `protobuf:"varint,8,opt,name=isReady" json:"isReady,omitempty"`
	WxInfo           *WeixinInfo       `protobuf:"bytes,10,opt,name=wxInfo" json:"wxInfo,omitempty"`
	DissolveState    *int32            `protobuf:"varint,11,opt,name=dissolveState" json:"dissolveState,omitempty"`
	IsRobot          *bool             `protobuf:"varint,12,opt,name=isRobot" json:"isRobot,omitempty"`
	IsOnWhiteList    *bool             `protobuf:"varint,13,opt,name=isOnWhiteList" json:"isOnWhiteList,omitempty"`
	WhiteWinRate     *int32            `protobuf:"varint,14,opt,name=whiteWinRate" json:"whiteWinRate,omitempty"`
	IsLeave          *bool             `protobuf:"varint,15,opt,name=isLeave" json:"isLeave,omitempty"`
	PaoNo            *int32            `protobuf:"varint,19,opt,name=paoNo" json:"paoNo,omitempty"`
	IsOnGaming       *bool             `protobuf:"varint,20,opt,name=isOnGaming" json:"isOnGaming,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *ZhadanSrvUser) Reset()                    { *m = ZhadanSrvUser{} }
func (m *ZhadanSrvUser) String() string            { return proto.CompactTextString(m) }
func (*ZhadanSrvUser) ProtoMessage()               {}
func (*ZhadanSrvUser) Descriptor() ([]byte, []int) { return fileDescriptor58, []int{2} }

func (m *ZhadanSrvUser) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *ZhadanSrvUser) GetBill() *ZhadanUserBill {
	if m != nil {
		return m.Bill
	}
	return nil
}

func (m *ZhadanSrvUser) GetIsOnline() bool {
	if m != nil && m.IsOnline != nil {
		return *m.IsOnline
	}
	return false
}

func (m *ZhadanSrvUser) GetIndex() int32 {
	if m != nil && m.Index != nil {
		return *m.Index
	}
	return 0
}

func (m *ZhadanSrvUser) GetPokers() *ZhadanSrvPoker {
	if m != nil {
		return m.Pokers
	}
	return nil
}

func (m *ZhadanSrvUser) GetOutPai() *ZhadanSrvPoker {
	if m != nil {
		return m.OutPai
	}
	return nil
}

func (m *ZhadanSrvUser) GetIsPass() bool {
	if m != nil && m.IsPass != nil {
		return *m.IsPass
	}
	return false
}

func (m *ZhadanSrvUser) GetDeskScore() int32 {
	if m != nil && m.DeskScore != nil {
		return *m.DeskScore
	}
	return 0
}

func (m *ZhadanSrvUser) GetIsBanker() bool {
	if m != nil && m.IsBanker != nil {
		return *m.IsBanker
	}
	return false
}

func (m *ZhadanSrvUser) GetBaozhuangStatus() int32 {
	if m != nil && m.BaozhuangStatus != nil {
		return *m.BaozhuangStatus
	}
	return 0
}

func (m *ZhadanSrvUser) GetScorePais() *ZhadanSrvPoker {
	if m != nil {
		return m.ScorePais
	}
	return nil
}

func (m *ZhadanSrvUser) GetExtScorePoker() []*ZhadanSrvPoker {
	if m != nil {
		return m.ExtScorePoker
	}
	return nil
}

func (m *ZhadanSrvUser) GetExtScore() int32 {
	if m != nil && m.ExtScore != nil {
		return *m.ExtScore
	}
	return 0
}

func (m *ZhadanSrvUser) GetPaoScore() int32 {
	if m != nil && m.PaoScore != nil {
		return *m.PaoScore
	}
	return 0
}

func (m *ZhadanSrvUser) GetIsReady() bool {
	if m != nil && m.IsReady != nil {
		return *m.IsReady
	}
	return false
}

func (m *ZhadanSrvUser) GetWxInfo() *WeixinInfo {
	if m != nil {
		return m.WxInfo
	}
	return nil
}

func (m *ZhadanSrvUser) GetDissolveState() int32 {
	if m != nil && m.DissolveState != nil {
		return *m.DissolveState
	}
	return 0
}

func (m *ZhadanSrvUser) GetIsRobot() bool {
	if m != nil && m.IsRobot != nil {
		return *m.IsRobot
	}
	return false
}

func (m *ZhadanSrvUser) GetIsOnWhiteList() bool {
	if m != nil && m.IsOnWhiteList != nil {
		return *m.IsOnWhiteList
	}
	return false
}

func (m *ZhadanSrvUser) GetWhiteWinRate() int32 {
	if m != nil && m.WhiteWinRate != nil {
		return *m.WhiteWinRate
	}
	return 0
}

func (m *ZhadanSrvUser) GetIsLeave() bool {
	if m != nil && m.IsLeave != nil {
		return *m.IsLeave
	}
	return false
}

func (m *ZhadanSrvUser) GetPaoNo() int32 {
	if m != nil && m.PaoNo != nil {
		return *m.PaoNo
	}
	return 0
}

func (m *ZhadanSrvUser) GetIsOnGaming() bool {
	if m != nil && m.IsOnGaming != nil {
		return *m.IsOnGaming
	}
	return false
}

// desk快照索引列表
type ZhadanSrvDeskSnapshotIdIndex struct {
	DeskId           []int32 `protobuf:"varint,1,rep,name=deskId" json:"deskId,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ZhadanSrvDeskSnapshotIdIndex) Reset()                    { *m = ZhadanSrvDeskSnapshotIdIndex{} }
func (m *ZhadanSrvDeskSnapshotIdIndex) String() string            { return proto.CompactTextString(m) }
func (*ZhadanSrvDeskSnapshotIdIndex) ProtoMessage()               {}
func (*ZhadanSrvDeskSnapshotIdIndex) Descriptor() ([]byte, []int) { return fileDescriptor58, []int{3} }

func (m *ZhadanSrvDeskSnapshotIdIndex) GetDeskId() []int32 {
	if m != nil {
		return m.DeskId
	}
	return nil
}

// 牌桌快照
type ZhadanSrvDeskSnapshot struct {
	DeskState        *ZhadanSrvDesk   `protobuf:"bytes,1,opt,name=deskState" json:"deskState,omitempty"`
	Users            []*ZhadanSrvUser `protobuf:"bytes,2,rep,name=users" json:"users,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *ZhadanSrvDeskSnapshot) Reset()                    { *m = ZhadanSrvDeskSnapshot{} }
func (m *ZhadanSrvDeskSnapshot) String() string            { return proto.CompactTextString(m) }
func (*ZhadanSrvDeskSnapshot) ProtoMessage()               {}
func (*ZhadanSrvDeskSnapshot) Descriptor() ([]byte, []int) { return fileDescriptor58, []int{4} }

func (m *ZhadanSrvDeskSnapshot) GetDeskState() *ZhadanSrvDesk {
	if m != nil {
		return m.DeskState
	}
	return nil
}

func (m *ZhadanSrvDeskSnapshot) GetUsers() []*ZhadanSrvUser {
	if m != nil {
		return m.Users
	}
	return nil
}

func init() {
	proto.RegisterType((*ZhadanSrvPoker)(nil), "ddproto.zhadan_srv_poker")
	proto.RegisterType((*ZhadanSrvDesk)(nil), "ddproto.zhadan_srv_desk")
	proto.RegisterType((*ZhadanSrvUser)(nil), "ddproto.zhadan_srv_user")
	proto.RegisterType((*ZhadanSrvDeskSnapshotIdIndex)(nil), "ddproto.zhadan_srv_desk_snapshot_id_index")
	proto.RegisterType((*ZhadanSrvDeskSnapshot)(nil), "ddproto.zhadan_srv_desk_snapshot")
}

func init() { proto.RegisterFile("zhadan_server.proto", fileDescriptor58) }

var fileDescriptor58 = []byte{
	// 949 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x55, 0xdd, 0x6e, 0xdb, 0x36,
	0x14, 0x9e, 0xea, 0xd8, 0x71, 0x98, 0x3a, 0x49, 0x99, 0x2e, 0x65, 0x7e, 0x56, 0x78, 0xc6, 0x30,
	0x18, 0x18, 0x16, 0x6c, 0xbe, 0x18, 0x0a, 0xf4, 0x62, 0xe8, 0x1a, 0xa0, 0x08, 0x10, 0x24, 0x06,
	0xd3, 0x2d, 0xc0, 0x6e, 0x04, 0xda, 0xe2, 0x62, 0xc2, 0x12, 0x29, 0x88, 0x54, 0xe2, 0xf4, 0x72,
	0x8f, 0xb4, 0x67, 0xd8, 0x4b, 0xed, 0x6e, 0x38, 0x87, 0x94, 0x2d, 0x19, 0x69, 0x73, 0x25, 0x7d,
	0x9f, 0xbe, 0xf3, 0x47, 0x9e, 0x73, 0x44, 0xf6, 0x3f, 0xcd, 0x44, 0x22, 0x74, 0x6c, 0x65, 0x71,
	0x27, 0x8b, 0xd3, 0xbc, 0x30, 0xce, 0xd0, 0xcd, 0x24, 0xc1, 0x97, 0xa3, 0xc3, 0xa9, 0xc9, 0x32,
	0x53, 0x7d, 0x8d, 0x73, 0x33, 0xaf, 0x34, 0x47, 0x2f, 0x82, 0xe1, 0x44, 0x58, 0xe9, 0xa9, 0xc1,
	0x3f, 0x11, 0xd9, 0xab, 0xdc, 0x15, 0x77, 0x5e, 0x4d, 0x7f, 0x22, 0x1b, 0xb9, 0x50, 0x96, 0x45,
	0xfd, 0xd6, 0x70, 0x7b, 0x74, 0x72, 0x1a, 0x5c, 0x9f, 0x56, 0x9e, 0x2b, 0xe1, 0x58, 0x28, 0x8e,
	0x4a, 0x3a, 0x22, 0x1b, 0x1f, 0x1f, 0x72, 0xc9, 0x9e, 0xf5, 0xa3, 0xe1, 0xce, 0xe8, 0xf5, 0xd2,
	0x22, 0xb8, 0x96, 0xba, 0xcc, 0xbc, 0x09, 0xa8, 0x38, 0x6a, 0xe9, 0x11, 0xe9, 0xce, 0xe5, 0xc3,
	0x1f, 0x22, 0x2d, 0x25, 0x6b, 0xf5, 0xa3, 0x61, 0x9b, 0x2f, 0x31, 0x3d, 0x21, 0x5b, 0x73, 0xf9,
	0x70, 0x21, 0xf5, 0xad, 0x9b, 0xb1, 0x0d, 0xfc, 0xb8, 0x22, 0x06, 0xff, 0x75, 0xc8, 0x6e, 0x2d,
	0xe9, 0x44, 0xda, 0x39, 0x3d, 0x20, 0x1d, 0x78, 0x9e, 0x27, 0x2c, 0x42, 0x79, 0x40, 0x74, 0x8f,
	0xb4, 0xf2, 0xfb, 0x04, 0x13, 0xdb, 0xe2, 0xf0, 0x4a, 0x5f, 0x13, 0x72, 0x2b, 0x32, 0x79, 0x59,
	0x66, 0x13, 0x59, 0x84, 0xc8, 0x35, 0x06, 0x3c, 0x15, 0xc6, 0x64, 0xe7, 0x49, 0x08, 0x1c, 0x10,
	0x7d, 0x43, 0x3a, 0xd6, 0x09, 0x57, 0x5a, 0xd6, 0xc6, 0x2a, 0xfb, 0x8f, 0x56, 0x09, 0x61, 0x63,
	0xaf, 0xe3, 0x41, 0x0f, 0x95, 0x4e, 0x55, 0x31, 0x4d, 0xe5, 0xa5, 0x61, 0x9b, 0xbe, 0xd2, 0x0a,
	0xd3, 0xef, 0x48, 0x6f, 0x5a, 0x16, 0xc5, 0x99, 0xb4, 0xf3, 0xeb, 0xa9, 0x29, 0x24, 0xeb, 0xa2,
	0xa0, 0x49, 0xd2, 0x5f, 0x49, 0x2f, 0xa9, 0xc0, 0x18, 0xae, 0xe6, 0x9b, 0x7e, 0x34, 0xdc, 0x1e,
	0x1d, 0xae, 0xa7, 0xb0, 0xbc, 0x1a, 0xde, 0xd4, 0xd3, 0x97, 0xa4, 0x6d, 0xee, 0xb5, 0x2c, 0xd8,
	0x76, 0x3f, 0x1a, 0xf6, 0xb8, 0x07, 0x90, 0x98, 0xb2, 0x67, 0x42, 0xcd, 0x85, 0x62, 0xcf, 0xfb,
	0xd1, 0xb0, 0xcb, 0x97, 0x18, 0x8e, 0x29, 0xc1, 0xb7, 0xdf, 0xad, 0x2c, 0x58, 0x0f, 0xcd, 0x6a,
	0x0c, 0x7d, 0x4b, 0x08, 0x84, 0xb8, 0xca, 0x9d, 0x32, 0x9a, 0xed, 0x60, 0x3e, 0xc7, 0xeb, 0xf9,
	0xe0, 0x69, 0x18, 0x94, 0xf0, 0x9a, 0x9c, 0xf6, 0xc9, 0x76, 0x2a, 0xac, 0x7b, 0x37, 0x75, 0xe8,
	0x7d, 0x17, 0xbd, 0xd7, 0x29, 0xfa, 0x3d, 0xd9, 0x01, 0xf8, 0x7e, 0x56, 0xe6, 0x21, 0x85, 0x3d,
	0x14, 0xad, 0xb1, 0xe0, 0x09, 0x8e, 0xaa, 0xf2, 0x74, 0xe4, 0x3d, 0xd5, 0x28, 0xb8, 0xcf, 0xa9,
	0xd0, 0x1f, 0x4a, 0xc3, 0x8e, 0xb1, 0xc4, 0x80, 0xe0, 0xe4, 0x27, 0xc2, 0xfc, 0x39, 0x2b, 0x85,
	0xbe, 0x45, 0xdb, 0x13, 0xb4, 0x6d, 0x92, 0x74, 0x40, 0x9e, 0x2b, 0x7b, 0xa5, 0xcf, 0x94, 0xb5,
	0x26, 0xbd, 0x93, 0x8c, 0xa2, 0x8f, 0x06, 0x07, 0x9a, 0x24, 0xbc, 0x7f, 0x54, 0x99, 0x64, 0xfb,
	0xfd, 0x68, 0xd8, 0xe2, 0x0d, 0x8e, 0x32, 0xb2, 0xa9, 0xec, 0xb5, 0x13, 0x85, 0x63, 0x07, 0xe8,
	0xa2, 0x82, 0x60, 0x6d, 0xb4, 0xc4, 0x77, 0xb4, 0x7e, 0xe5, 0xad, 0xeb, 0x1c, 0x68, 0x44, 0x9a,
	0xae, 0x34, 0xcc, 0x6b, 0xea, 0x5c, 0x3d, 0x0b, 0x2c, 0xe7, 0x10, 0xcb, 0x69, 0x70, 0x70, 0xa9,
	0xca, 0xbe, 0x37, 0x4a, 0x73, 0x63, 0x32, 0xf6, 0x12, 0x13, 0xa9, 0x31, 0x70, 0x9a, 0xb6, 0x2c,
	0xf2, 0xb4, 0xb4, 0x18, 0xe6, 0x6b, 0xec, 0xc5, 0x3a, 0x35, 0xf8, 0xb7, 0x39, 0x7b, 0x65, 0x38,
	0x61, 0x78, 0x86, 0xd9, 0xeb, 0xf1, 0x80, 0xe8, 0x8f, 0x64, 0x63, 0xa2, 0xd2, 0x14, 0x87, 0xef,
	0x91, 0x66, 0x05, 0x55, 0x0c, 0x02, 0x8e, 0x32, 0xdf, 0x8d, 0x57, 0x3a, 0x55, 0xda, 0x2f, 0x04,
	0xec, 0x46, 0x8f, 0xa1, 0x7f, 0x95, 0x4e, 0xe4, 0x22, 0xcc, 0xa4, 0x07, 0xf4, 0x67, 0xd2, 0xc1,
	0x6e, 0xf7, 0x23, 0xf9, 0xc5, 0x79, 0x08, 0x42, 0x3a, 0x22, 0x9b, 0xa6, 0x74, 0x71, 0x2e, 0x14,
	0xeb, 0x3c, 0x69, 0x63, 0x4a, 0x37, 0x16, 0x0a, 0xea, 0x53, 0x76, 0x2c, 0xac, 0xc5, 0xe9, 0xed,
	0xf2, 0x80, 0x60, 0x4b, 0x2d, 0xa7, 0x0c, 0xdb, 0xb3, 0xcd, 0x57, 0x84, 0x2f, 0xe7, 0x37, 0xa1,
	0xe7, 0xb2, 0x60, 0x2f, 0xaa, 0x72, 0x3c, 0xa6, 0x43, 0xb2, 0x3b, 0x11, 0xe6, 0x13, 0xb6, 0xd9,
	0xb5, 0x5f, 0x2a, 0x14, 0xed, 0xd7, 0x69, 0xfa, 0x86, 0x10, 0x0b, 0xee, 0x62, 0xdc, 0xc8, 0x07,
	0x4f, 0xa5, 0xbc, 0x65, 0x97, 0x23, 0xff, 0x8e, 0xec, 0xca, 0x85, 0x8b, 0x83, 0x35, 0x7c, 0x65,
	0xaf, 0x70, 0xa1, 0x7f, 0x69, 0x6b, 0xc8, 0x85, 0xf3, 0x4b, 0x03, 0x7f, 0x04, 0xc7, 0x64, 0x6b,
	0xe9, 0x02, 0x7b, 0xae, 0xcd, 0xbb, 0x95, 0x02, 0xea, 0xcb, 0x85, 0xf1, 0xc5, 0x1f, 0xfa, 0x6f,
	0x15, 0xf6, 0xdd, 0xce, 0xa5, 0x48, 0x1e, 0x70, 0x9f, 0x61, 0xb7, 0x23, 0xa4, 0x3f, 0x90, 0xce,
	0xfd, 0xe2, 0x5c, 0xff, 0x65, 0x18, 0xc1, 0x5a, 0xf6, 0x97, 0xc9, 0xdc, 0x48, 0xb5, 0x50, 0x1a,
	0x3e, 0xf1, 0x20, 0x81, 0x11, 0xad, 0xda, 0x17, 0x8e, 0x43, 0xe2, 0xf6, 0x6a, 0xf3, 0x26, 0x19,
	0x82, 0x99, 0x89, 0x71, 0x61, 0x89, 0x55, 0x10, 0xec, 0xa1, 0x83, 0x6e, 0x66, 0xca, 0xc9, 0x0b,
	0x65, 0x1d, 0xae, 0xb1, 0x2e, 0x6f, 0x92, 0x30, 0x38, 0xf7, 0x00, 0x6e, 0x94, 0xe6, 0x10, 0x64,
	0x07, 0x83, 0x34, 0x38, 0x1f, 0xe3, 0x42, 0x8a, 0x3b, 0x89, 0xcb, 0x0a, 0x63, 0x20, 0x84, 0xce,
	0xcc, 0x85, 0xb9, 0x34, 0x38, 0xf5, 0x6d, 0xee, 0x81, 0x1f, 0xb4, 0x2b, 0xfd, 0x41, 0x64, 0x4a,
	0xdf, 0xae, 0x06, 0xad, 0x62, 0x06, 0x6f, 0xc9, 0xb7, 0x6b, 0x7f, 0xb0, 0xd8, 0x6a, 0x91, 0xdb,
	0x99, 0x71, 0xb1, 0x4a, 0x62, 0xdf, 0xde, 0xf5, 0x7f, 0x5a, 0x6b, 0xf5, 0x4f, 0x1b, 0xfc, 0x1d,
	0x11, 0xf6, 0x39, 0x6b, 0xfa, 0x4b, 0x68, 0x4a, 0x3c, 0xaf, 0x08, 0xcf, 0x98, 0x3d, 0x76, 0xe1,
	0x20, 0xe2, 0x2b, 0x29, 0x3d, 0x25, 0x6d, 0x18, 0x48, 0xcb, 0x9e, 0x61, 0x93, 0x3c, 0x6a, 0x03,
	0x02, 0xee, 0x65, 0xe3, 0xaf, 0xc6, 0xd1, 0xff, 0x01, 0x00, 0x00, 0xff, 0xff, 0xbb, 0x7e, 0x22,
	0xc1, 0x8a, 0x08, 0x00, 0x00,
}