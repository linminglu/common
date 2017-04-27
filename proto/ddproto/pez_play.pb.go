// Code generated by protoc-gen-go.
// source: pez_play.proto
// DO NOT EDIT!

package ddproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Ignoring public import of ProtoHeader from common_client.proto

// Ignoring public import of Heartbeat from common_client.proto

// Ignoring public import of ServerInfo from common_client.proto

// Ignoring public import of QuickConn from common_client.proto

// Ignoring public import of AckQuickConn from common_client.proto

// Ignoring public import of WeixinInfo from common_client.proto

// Ignoring public import of common_req_reg from common_client.proto

// Ignoring public import of common_req_reg_via_input from common_client.proto

// Ignoring public import of common_ack_reg from common_client.proto

// Ignoring public import of common_req_gameLogin from common_client.proto

// Ignoring public import of common_req_gameLogin_via_input from common_client.proto

// Ignoring public import of common_ack_gameLogin from common_client.proto

// Ignoring public import of common_req_qrLogin from common_client.proto

// Ignoring public import of common_ack_qrLogin from common_client.proto

// Ignoring public import of common_req_qrWxInfo from common_client.proto

// Ignoring public import of common_ack_qrWxInfo from common_client.proto

// Ignoring public import of common_ack_reconnect from common_client.proto

// Ignoring public import of common_req_reconnect from common_client.proto

// Ignoring public import of common_req_gameState from common_client.proto

// Ignoring public import of common_ack_gameState from common_client.proto

// Ignoring public import of common_req_logout from common_client.proto

// Ignoring public import of common_ack_logout from common_client.proto

// Ignoring public import of common_req_feedback from common_client.proto

// Ignoring public import of client_base_poker from common_client.proto

// Ignoring public import of common_req_message from common_client.proto

// Ignoring public import of common_bc_message from common_client.proto

// Ignoring public import of common_req_notice from common_client.proto

// Ignoring public import of common_ack_notice from common_client.proto

// Ignoring public import of common_req_enterAgentMode from common_client.proto

// Ignoring public import of common_ack_enterAgentMode from common_client.proto

// Ignoring public import of common_req_quitAgentMode from common_client.proto

// Ignoring public import of common_ack_quitAgentMode from common_client.proto

// Ignoring public import of common_req_leaveDesk from common_client.proto

// Ignoring public import of common_ack_leaveDesk from common_client.proto

// Ignoring public import of common_bc_kickout from common_client.proto

// Ignoring public import of common_req_allowance from common_client.proto

// Ignoring public import of common_ack_allowance from common_client.proto

// Ignoring public import of common_req_applyDissolve from common_client.proto

// Ignoring public import of common_bc_applyDissolve from common_client.proto

// Ignoring public import of common_req_applyDissolveBack from common_client.proto

// Ignoring public import of common_ack_applyDissolveBack from common_client.proto

// Ignoring public import of common_ack_timeout from common_client.proto

// Ignoring public import of common_bc_userBreak from common_client.proto

// Ignoring public import of common_req_clickStatistic from common_client.proto

// Ignoring public import of common_req_offline from common_client.proto

// Ignoring public import of common_enum_reg from common_client.proto

// Ignoring public import of common_enum_os_type from common_client.proto

// Ignoring public import of common_enum_pokerColor from common_client.proto

// Ignoring public import of pez_base_PaiInfo from pez_base.proto

// Ignoring public import of pez_base_PlayConf from pez_base.proto

// Ignoring public import of pez_base_RoomTypeInfo from pez_base.proto

// Ignoring public import of pez_base_PaiValue from pez_base.proto

// Ignoring public import of pez_base_PlayerCard from pez_base.proto

// Ignoring public import of pez_base_PlayerInfo from pez_base.proto

// Ignoring public import of pez_base_DeskGameInfo from pez_base.proto

// Ignoring public import of pez_enum_protoId from pez_base.proto

// Ignoring public import of pez_enum_ErrorCode from pez_base.proto

// Ignoring public import of pez_enum_Option from pez_base.proto

// Ignoring public import of pez_RoomType from pez_base.proto

// Ignoring public import of pez_enum_mjColor from pez_base.proto

// Ignoring public import of pez_enum_PaiType from pez_base.proto

// Ignoring public import of pez_enum_Base from pez_base.proto

// Ignoring public import of pez_enum_Bet from pez_base.proto

// Ignoring public import of pez_enum_UserGameStatus from pez_base.proto

// Ignoring public import of pez_enum_DeskGameStatus from pez_base.proto

// Ignoring public import of pez_enum_TuoziType from pez_base.proto

// Ignoring public import of common_enum_game from common_enum.proto

// Ignoring public import of COMMON_ENUM_ROOMTYPE from common_enum.proto

// Ignoring public import of COMMON_ENUM_GAMESTATUS from common_enum.proto

// Ignoring public import of COMMON_ENUM_RELEASETAG from common_enum.proto

// Ignoring public import of COMMON_ENUM_KICKOUT from common_enum.proto

// Ignoring public import of COMMON_ENUM_APPLYDISSOLVE from common_enum.proto

// Ignoring public import of BTN_TYPE from common_enum.proto

// Ignoring public import of COMMON_ENUM_ERROR_TYPE_ENTERCOINROOM from common_enum.proto

// Ignoring public import of pez_DissolveDesk from pez_desk.proto

// Ignoring public import of pez_AckDissolveDesk from pez_desk.proto

// Ignoring public import of pez_Ready from pez_desk.proto

// Ignoring public import of pez_AckReady from pez_desk.proto

// Ignoring public import of pez_EndLotteryInfo from pez_desk.proto

// Ignoring public import of pez_SendCurrentResult from pez_desk.proto

// Ignoring public import of EndLottery from pez_desk.proto

// Ignoring public import of pez_SendEndLottery from pez_desk.proto

// Ignoring public import of pez_tuozi from pez_desk.proto

// Ignoring public import of pez_UserBean from pez_desk.proto

// Ignoring public import of pez_Bill from pez_desk.proto

// 链接类型
type PEZ_RECONNECT_TYPE int32

const (
	PEZ_RECONNECT_TYPE_PEZ_NORMAL    PEZ_RECONNECT_TYPE = 1
	PEZ_RECONNECT_TYPE_PEZ_RECONNECT PEZ_RECONNECT_TYPE = 2
)

var PEZ_RECONNECT_TYPE_name = map[int32]string{
	1: "PEZ_NORMAL",
	2: "PEZ_RECONNECT",
}
var PEZ_RECONNECT_TYPE_value = map[string]int32{
	"PEZ_NORMAL":    1,
	"PEZ_RECONNECT": 2,
}

func (x PEZ_RECONNECT_TYPE) Enum() *PEZ_RECONNECT_TYPE {
	p := new(PEZ_RECONNECT_TYPE)
	*p = x
	return p
}
func (x PEZ_RECONNECT_TYPE) String() string {
	return proto.EnumName(PEZ_RECONNECT_TYPE_name, int32(x))
}
func (x *PEZ_RECONNECT_TYPE) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(PEZ_RECONNECT_TYPE_value, data, "PEZ_RECONNECT_TYPE")
	if err != nil {
		return err
	}
	*x = PEZ_RECONNECT_TYPE(value)
	return nil
}
func (PEZ_RECONNECT_TYPE) EnumDescriptor() ([]byte, []int) { return fileDescriptor38, []int{0} }

// 积分
type PezUserCoinBean struct {
	UserId           *uint32 `protobuf:"varint,1,opt,name=userId" json:"userId,omitempty"`
	Coin             *int64  `protobuf:"varint,2,opt,name=coin" json:"coin,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *PezUserCoinBean) Reset()                    { *m = PezUserCoinBean{} }
func (m *PezUserCoinBean) String() string            { return proto.CompactTextString(m) }
func (*PezUserCoinBean) ProtoMessage()               {}
func (*PezUserCoinBean) Descriptor() ([]byte, []int) { return fileDescriptor38, []int{0} }

func (m *PezUserCoinBean) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *PezUserCoinBean) GetCoin() int64 {
	if m != nil && m.Coin != nil {
		return *m.Coin
	}
	return 0
}

// 开局（接收服务端消息）
type Pez_Opening struct {
	Header           *ProtoHeader       `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	CurrPlayCount    *int32             `protobuf:"varint,2,opt,name=CurrPlayCount" json:"CurrPlayCount,omitempty"`
	UserCoinBeans    []*PezUserCoinBean `protobuf:"bytes,3,rep,name=userCoinBeans" json:"userCoinBeans,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *Pez_Opening) Reset()                    { *m = Pez_Opening{} }
func (m *Pez_Opening) String() string            { return proto.CompactTextString(m) }
func (*Pez_Opening) ProtoMessage()               {}
func (*Pez_Opening) Descriptor() ([]byte, []int) { return fileDescriptor38, []int{1} }

func (m *Pez_Opening) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Pez_Opening) GetCurrPlayCount() int32 {
	if m != nil && m.CurrPlayCount != nil {
		return *m.CurrPlayCount
	}
	return 0
}

func (m *Pez_Opening) GetUserCoinBeans() []*PezUserCoinBean {
	if m != nil {
		return m.UserCoinBeans
	}
	return nil
}

// 发牌
type Pez_DealCards struct {
	Header           *ProtoHeader    `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Banker           *uint32         `protobuf:"varint,2,opt,name=banker" json:"banker,omitempty"`
	Dice1            *int32          `protobuf:"varint,3,opt,name=dice1" json:"dice1,omitempty"`
	Dice2            *int32          `protobuf:"varint,4,opt,name=dice2" json:"dice2,omitempty"`
	LastScore        []*PezLastScore `protobuf:"bytes,5,rep,name=lastScore" json:"lastScore,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *Pez_DealCards) Reset()                    { *m = Pez_DealCards{} }
func (m *Pez_DealCards) String() string            { return proto.CompactTextString(m) }
func (*Pez_DealCards) ProtoMessage()               {}
func (*Pez_DealCards) Descriptor() ([]byte, []int) { return fileDescriptor38, []int{2} }

func (m *Pez_DealCards) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Pez_DealCards) GetBanker() uint32 {
	if m != nil && m.Banker != nil {
		return *m.Banker
	}
	return 0
}

func (m *Pez_DealCards) GetDice1() int32 {
	if m != nil && m.Dice1 != nil {
		return *m.Dice1
	}
	return 0
}

func (m *Pez_DealCards) GetDice2() int32 {
	if m != nil && m.Dice2 != nil {
		return *m.Dice2
	}
	return 0
}

func (m *Pez_DealCards) GetLastScore() []*PezLastScore {
	if m != nil {
		return m.LastScore
	}
	return nil
}

type PezLastScore struct {
	UserId           *uint32 `protobuf:"varint,1,opt,name=userId" json:"userId,omitempty"`
	Score            *int64  `protobuf:"varint,2,opt,name=score" json:"score,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *PezLastScore) Reset()                    { *m = PezLastScore{} }
func (m *PezLastScore) String() string            { return proto.CompactTextString(m) }
func (*PezLastScore) ProtoMessage()               {}
func (*PezLastScore) Descriptor() ([]byte, []int) { return fileDescriptor38, []int{3} }

func (m *PezLastScore) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *PezLastScore) GetScore() int64 {
	if m != nil && m.Score != nil {
		return *m.Score
	}
	return 0
}

// 押注
type Pez_Bet struct {
	UserId           *uint32 `protobuf:"varint,1,opt,name=userId" json:"userId,omitempty"`
	BetNum           *int64  `protobuf:"varint,2,opt,name=betNum" json:"betNum,omitempty"`
	Time             *int32  `protobuf:"varint,3,opt,name=time" json:"time,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Pez_Bet) Reset()                    { *m = Pez_Bet{} }
func (m *Pez_Bet) String() string            { return proto.CompactTextString(m) }
func (*Pez_Bet) ProtoMessage()               {}
func (*Pez_Bet) Descriptor() ([]byte, []int) { return fileDescriptor38, []int{4} }

func (m *Pez_Bet) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *Pez_Bet) GetBetNum() int64 {
	if m != nil && m.BetNum != nil {
		return *m.BetNum
	}
	return 0
}

func (m *Pez_Bet) GetTime() int32 {
	if m != nil && m.Time != nil {
		return *m.Time
	}
	return 0
}

type Pez_AckBet struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserId           *uint32      `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	BetCount         *int64       `protobuf:"varint,3,opt,name=betCount" json:"betCount,omitempty"`
	Time             *int32       `protobuf:"varint,4,opt,name=time" json:"time,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *Pez_AckBet) Reset()                    { *m = Pez_AckBet{} }
func (m *Pez_AckBet) String() string            { return proto.CompactTextString(m) }
func (*Pez_AckBet) ProtoMessage()               {}
func (*Pez_AckBet) Descriptor() ([]byte, []int) { return fileDescriptor38, []int{5} }

func (m *Pez_AckBet) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Pez_AckBet) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *Pez_AckBet) GetBetCount() int64 {
	if m != nil && m.BetCount != nil {
		return *m.BetCount
	}
	return 0
}

func (m *Pez_AckBet) GetTime() int32 {
	if m != nil && m.Time != nil {
		return *m.Time
	}
	return 0
}

type Pez_BCOpenPai struct {
	UserId           *uint32            `protobuf:"varint,1,opt,name=userId" json:"userId,omitempty"`
	HandPai          []*PezBase_PaiInfo `protobuf:"bytes,2,rep,name=handPai" json:"handPai,omitempty"`
	KeyValue         *int32             `protobuf:"varint,3,opt,name=keyValue" json:"keyValue,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *Pez_BCOpenPai) Reset()                    { *m = Pez_BCOpenPai{} }
func (m *Pez_BCOpenPai) String() string            { return proto.CompactTextString(m) }
func (*Pez_BCOpenPai) ProtoMessage()               {}
func (*Pez_BCOpenPai) Descriptor() ([]byte, []int) { return fileDescriptor38, []int{6} }

func (m *Pez_BCOpenPai) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *Pez_BCOpenPai) GetHandPai() []*PezBase_PaiInfo {
	if m != nil {
		return m.HandPai
	}
	return nil
}

func (m *Pez_BCOpenPai) GetKeyValue() int32 {
	if m != nil && m.KeyValue != nil {
		return *m.KeyValue
	}
	return 0
}

// 发送游戏信息(广播)
type Pez_SendGameInfo struct {
	Header *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	// 1. 首先是牌桌的玩家数据（玩家数据包括其id昵称筹码头像等基本信息，其手牌数据，以及自己打开牌的数据，还有状态是否已经押注了，玩家在整局的总输赢）
	PlayerInfo []*PezBase_PlayerInfo `protobuf:"bytes,2,rep,name=playerInfo" json:"playerInfo,omitempty"`
	// 2. 桌面信息（包括：游戏是否结束，当前哪个玩家未押注，倒计时剩余时间）
	DeskGameInfo *PezBase_DeskGameInfo `protobuf:"bytes,3,opt,name=deskGameInfo" json:"deskGameInfo,omitempty"`
	//
	SenderUserId     *uint32 `protobuf:"varint,4,opt,name=senderUserId" json:"senderUserId,omitempty"`
	IsReconnect      *bool   `protobuf:"varint,5,opt,name=isReconnect" json:"isReconnect,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Pez_SendGameInfo) Reset()                    { *m = Pez_SendGameInfo{} }
func (m *Pez_SendGameInfo) String() string            { return proto.CompactTextString(m) }
func (*Pez_SendGameInfo) ProtoMessage()               {}
func (*Pez_SendGameInfo) Descriptor() ([]byte, []int) { return fileDescriptor38, []int{7} }

func (m *Pez_SendGameInfo) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Pez_SendGameInfo) GetPlayerInfo() []*PezBase_PlayerInfo {
	if m != nil {
		return m.PlayerInfo
	}
	return nil
}

func (m *Pez_SendGameInfo) GetDeskGameInfo() *PezBase_DeskGameInfo {
	if m != nil {
		return m.DeskGameInfo
	}
	return nil
}

func (m *Pez_SendGameInfo) GetSenderUserId() uint32 {
	if m != nil && m.SenderUserId != nil {
		return *m.SenderUserId
	}
	return 0
}

func (m *Pez_SendGameInfo) GetIsReconnect() bool {
	if m != nil && m.IsReconnect != nil {
		return *m.IsReconnect
	}
	return false
}

func init() {
	proto.RegisterType((*PezUserCoinBean)(nil), "ddproto.pez_user_coin_bean")
	proto.RegisterType((*Pez_Opening)(nil), "ddproto.pez_Opening")
	proto.RegisterType((*Pez_DealCards)(nil), "ddproto.pez_DealCards")
	proto.RegisterType((*PezLastScore)(nil), "ddproto.pez_lastScore")
	proto.RegisterType((*Pez_Bet)(nil), "ddproto.pez_Bet")
	proto.RegisterType((*Pez_AckBet)(nil), "ddproto.pez_AckBet")
	proto.RegisterType((*Pez_BCOpenPai)(nil), "ddproto.pez_BCOpenPai")
	proto.RegisterType((*Pez_SendGameInfo)(nil), "ddproto.pez_SendGameInfo")
	proto.RegisterEnum("ddproto.PEZ_RECONNECT_TYPE", PEZ_RECONNECT_TYPE_name, PEZ_RECONNECT_TYPE_value)
}

var fileDescriptor38 = []byte{
	// 574 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0x5f, 0x8b, 0xd3, 0x4e,
	0x14, 0xfd, 0x25, 0x69, 0xb7, 0xfb, 0xbb, 0xdd, 0x2c, 0xdd, 0xb1, 0x94, 0x58, 0x45, 0x4a, 0xf0,
	0xa1, 0x88, 0x14, 0xac, 0x82, 0x2f, 0x0a, 0xb6, 0xd9, 0xe2, 0x2e, 0xb8, 0x6d, 0x98, 0x5d, 0x05,
	0x7d, 0x09, 0xd3, 0xe4, 0xea, 0x86, 0x26, 0x93, 0x92, 0x3f, 0x60, 0x7d, 0xf4, 0xc9, 0x0f, 0xe2,
	0x77, 0xf0, 0xeb, 0xc9, 0x4c, 0xa6, 0xd9, 0x86, 0xb5, 0x0f, 0x7d, 0x09, 0x73, 0xce, 0x9c, 0x73,
	0x73, 0xee, 0x9d, 0x19, 0x38, 0x5d, 0xe3, 0x0f, 0x6f, 0x1d, 0xb1, 0xcd, 0x68, 0x9d, 0x26, 0x79,
	0x42, 0x5a, 0x41, 0x20, 0x17, 0xfd, 0x07, 0x7e, 0x12, 0xc7, 0x09, 0xf7, 0xfc, 0x28, 0x44, 0x9e,
	0x97, 0xbb, 0x7d, 0xa9, 0x5e, 0xb2, 0x0c, 0x15, 0x3e, 0x53, 0x22, 0xe4, 0x45, 0xbc, 0x2b, 0x09,
	0x30, 0x5b, 0x95, 0xd8, 0x7e, 0x07, 0x44, 0x30, 0x45, 0x86, 0xa9, 0xe7, 0x27, 0x21, 0xf7, 0x96,
	0xc8, 0x38, 0xe9, 0xc1, 0x91, 0x60, 0x2e, 0x03, 0x4b, 0x1b, 0x68, 0x43, 0x93, 0x2a, 0x44, 0x08,
	0x34, 0x84, 0xc8, 0xd2, 0x07, 0xda, 0xd0, 0xa0, 0x72, 0x6d, 0xff, 0xd6, 0xa0, 0x2d, 0x4a, 0x2c,
	0xd6, 0xc8, 0x43, 0xfe, 0x8d, 0x3c, 0x87, 0xa3, 0x5b, 0x64, 0x01, 0xa6, 0xd2, 0xdb, 0x1e, 0x77,
	0x47, 0x2a, 0xf3, 0xc8, 0x15, 0xdf, 0x0b, 0xb9, 0x47, 0x95, 0x86, 0x3c, 0x05, 0xd3, 0x29, 0xd2,
	0xd4, 0x8d, 0xd8, 0xc6, 0x49, 0x0a, 0x9e, 0xcb, 0xd2, 0x4d, 0x5a, 0x27, 0xc9, 0x04, 0x4c, 0x91,
	0xc0, 0x49, 0x42, 0x3e, 0x45, 0xc6, 0x33, 0xcb, 0x18, 0x18, 0xc3, 0xf6, 0xf8, 0x51, 0x55, 0xfa,
	0x7e, 0x0f, 0xb4, 0xee, 0xb0, 0xff, 0x68, 0x60, 0x0a, 0xd5, 0x39, 0xb2, 0xc8, 0x61, 0x69, 0x90,
	0x1d, 0x18, 0xb4, 0x07, 0x47, 0x4b, 0xc6, 0x57, 0x98, 0xca, 0x84, 0x26, 0x55, 0x88, 0x74, 0xa1,
	0x19, 0x84, 0x3e, 0xbe, 0xb0, 0x0c, 0x19, 0xbc, 0x04, 0x5b, 0x76, 0x6c, 0x35, 0xee, 0xd8, 0x31,
	0x79, 0x05, 0xff, 0x47, 0x2c, 0xcb, 0xaf, 0xfd, 0x24, 0x45, 0xab, 0x29, 0x5b, 0xe8, 0xd5, 0x5a,
	0xa8, 0x76, 0xe9, 0x9d, 0xd0, 0x7e, 0x5b, 0x06, 0xaf, 0x88, 0xbd, 0xa7, 0xd3, 0x85, 0x66, 0x26,
	0x4b, 0x97, 0xc7, 0x53, 0x02, 0xfb, 0x0a, 0x5a, 0xc2, 0x3e, 0xc5, 0x7c, 0xaf, 0x51, 0xf4, 0x86,
	0xf9, 0xbc, 0x88, 0x95, 0x53, 0x21, 0x71, 0xdc, 0x79, 0x18, 0xa3, 0x6a, 0x4d, 0xae, 0xed, 0x9f,
	0x1a, 0x80, 0xa8, 0x37, 0xf1, 0x57, 0xa2, 0xe4, 0xc1, 0x43, 0x54, 0x01, 0xf4, 0x5a, 0x80, 0x3e,
	0x1c, 0x2f, 0x31, 0x2f, 0x2f, 0x80, 0x21, 0x23, 0x54, 0xb8, 0x0a, 0xd1, 0xd8, 0x09, 0xf1, 0xbd,
	0x1c, 0xc9, 0xd4, 0x11, 0x97, 0xce, 0x65, 0xe1, 0xde, 0xce, 0x5e, 0x42, 0xeb, 0x96, 0xf1, 0xc0,
	0x65, 0xa1, 0xa5, 0xcb, 0x79, 0x3f, 0xac, 0xcd, 0x5b, 0xbc, 0x15, 0xcf, 0x65, 0xe1, 0x25, 0xff,
	0x9a, 0xd0, 0xad, 0x52, 0xa4, 0x59, 0xe1, 0xe6, 0x13, 0x8b, 0x8a, 0x6d, 0xeb, 0x15, 0xb6, 0x7f,
	0xe9, 0xd0, 0x11, 0xce, 0x6b, 0xe4, 0xc1, 0x7b, 0x16, 0xa3, 0x70, 0x1e, 0x38, 0x84, 0x37, 0x00,
	0xe2, 0x45, 0x63, 0x2a, 0xbc, 0x2a, 0xd6, 0xe3, 0x7f, 0xc4, 0xaa, 0x34, 0x74, 0x47, 0x4f, 0xa6,
	0x70, 0x22, 0x9e, 0xef, 0xf6, 0xdf, 0x32, 0x60, 0x7b, 0xfc, 0xe4, 0xbe, 0xff, 0x7c, 0x47, 0x45,
	0x6b, 0x1e, 0x62, 0xc3, 0x49, 0x86, 0x3c, 0xc0, 0xf4, 0x63, 0x39, 0xb3, 0x86, 0x9c, 0x59, 0x8d,
	0x23, 0x03, 0x68, 0x87, 0x19, 0x45, 0x3f, 0xe1, 0x1c, 0xfd, 0xdc, 0x6a, 0x0e, 0xb4, 0xe1, 0x31,
	0xdd, 0xa5, 0x9e, 0xbd, 0x06, 0xe2, 0xce, 0xbe, 0x78, 0x74, 0xe6, 0x2c, 0xe6, 0xf3, 0x99, 0x73,
	0xe3, 0xdd, 0x7c, 0x76, 0x67, 0xe4, 0x14, 0x40, 0xb0, 0xf3, 0x05, 0xbd, 0x9a, 0x7c, 0xe8, 0x68,
	0xe4, 0x0c, 0xcc, 0x9a, 0xaa, 0xa3, 0x4f, 0xf5, 0x0b, 0xc3, 0xfd, 0xcf, 0xd5, 0x5c, 0xdd, 0x35,
	0xfe, 0x06, 0x00, 0x00, 0xff, 0xff, 0x6e, 0x21, 0x42, 0xae, 0xe1, 0x04, 0x00, 0x00,
}
