// Code generated by protoc-gen-go.
// source: erddz_desk.proto
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

// 房主解散房间(未开局)
type ErddzReqDissolveDesk struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserId           *uint32      `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *ErddzReqDissolveDesk) Reset()                    { *m = ErddzReqDissolveDesk{} }
func (m *ErddzReqDissolveDesk) String() string            { return proto.CompactTextString(m) }
func (*ErddzReqDissolveDesk) ProtoMessage()               {}
func (*ErddzReqDissolveDesk) Descriptor() ([]byte, []int) { return fileDescriptor19, []int{0} }

func (m *ErddzReqDissolveDesk) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *ErddzReqDissolveDesk) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

// 解散房间回复
type ErddzAckDissolveDesk struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserId           *uint32      `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	DeskId           *int32       `protobuf:"varint,3,opt,name=deskId" json:"deskId,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *ErddzAckDissolveDesk) Reset()                    { *m = ErddzAckDissolveDesk{} }
func (m *ErddzAckDissolveDesk) String() string            { return proto.CompactTextString(m) }
func (*ErddzAckDissolveDesk) ProtoMessage()               {}
func (*ErddzAckDissolveDesk) Descriptor() ([]byte, []int) { return fileDescriptor19, []int{1} }

func (m *ErddzAckDissolveDesk) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *ErddzAckDissolveDesk) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *ErddzAckDissolveDesk) GetDeskId() int32 {
	if m != nil && m.DeskId != nil {
		return *m.DeskId
	}
	return 0
}

// 准备游戏
type ErddzReqReady struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserId           *uint32      `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *ErddzReqReady) Reset()                    { *m = ErddzReqReady{} }
func (m *ErddzReqReady) String() string            { return proto.CompactTextString(m) }
func (*ErddzReqReady) ProtoMessage()               {}
func (*ErddzReqReady) Descriptor() ([]byte, []int) { return fileDescriptor19, []int{2} }

func (m *ErddzReqReady) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *ErddzReqReady) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

// 准备游戏的结果
type ErddzAckReady struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	UserId           *uint32      `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *ErddzAckReady) Reset()                    { *m = ErddzAckReady{} }
func (m *ErddzAckReady) String() string            { return proto.CompactTextString(m) }
func (*ErddzAckReady) ProtoMessage()               {}
func (*ErddzAckReady) Descriptor() ([]byte, []int) { return fileDescriptor19, []int{3} }

func (m *ErddzAckReady) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *ErddzAckReady) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

// 赢牌信息：谁赢了多少
type ErddzBaseWinCoinInfo struct {
	NickName         *string            `protobuf:"bytes,1,opt,name=nickName" json:"nickName,omitempty"`
	UserId           *uint32            `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	BaseValue        *int32             `protobuf:"varint,3,opt,name=baseValue" json:"baseValue,omitempty"`
	WinCoin          *int64             `protobuf:"varint,4,opt,name=winCoin" json:"winCoin,omitempty"`
	Coin             *int64             `protobuf:"varint,5,opt,name=coin" json:"coin,omitempty"`
	IsDiZhu          *bool              `protobuf:"varint,6,opt,name=isDiZhu" json:"isDiZhu,omitempty"`
	Rate             *int32             `protobuf:"varint,7,opt,name=rate" json:"rate,omitempty"`
	Description      *string            `protobuf:"bytes,8,opt,name=description" json:"description,omitempty"`
	Bomb             *int32             `protobuf:"varint,9,opt,name=bomb" json:"bomb,omitempty"`
	IsSpring         *bool              `protobuf:"varint,10,opt,name=isSpring" json:"isSpring,omitempty"`
	HandPokers       []*ClientBasePoker `protobuf:"bytes,11,rep,name=handPokers" json:"handPokers,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *ErddzBaseWinCoinInfo) Reset()                    { *m = ErddzBaseWinCoinInfo{} }
func (m *ErddzBaseWinCoinInfo) String() string            { return proto.CompactTextString(m) }
func (*ErddzBaseWinCoinInfo) ProtoMessage()               {}
func (*ErddzBaseWinCoinInfo) Descriptor() ([]byte, []int) { return fileDescriptor19, []int{4} }

func (m *ErddzBaseWinCoinInfo) GetNickName() string {
	if m != nil && m.NickName != nil {
		return *m.NickName
	}
	return ""
}

func (m *ErddzBaseWinCoinInfo) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *ErddzBaseWinCoinInfo) GetBaseValue() int32 {
	if m != nil && m.BaseValue != nil {
		return *m.BaseValue
	}
	return 0
}

func (m *ErddzBaseWinCoinInfo) GetWinCoin() int64 {
	if m != nil && m.WinCoin != nil {
		return *m.WinCoin
	}
	return 0
}

func (m *ErddzBaseWinCoinInfo) GetCoin() int64 {
	if m != nil && m.Coin != nil {
		return *m.Coin
	}
	return 0
}

func (m *ErddzBaseWinCoinInfo) GetIsDiZhu() bool {
	if m != nil && m.IsDiZhu != nil {
		return *m.IsDiZhu
	}
	return false
}

func (m *ErddzBaseWinCoinInfo) GetRate() int32 {
	if m != nil && m.Rate != nil {
		return *m.Rate
	}
	return 0
}

func (m *ErddzBaseWinCoinInfo) GetDescription() string {
	if m != nil && m.Description != nil {
		return *m.Description
	}
	return ""
}

func (m *ErddzBaseWinCoinInfo) GetBomb() int32 {
	if m != nil && m.Bomb != nil {
		return *m.Bomb
	}
	return 0
}

func (m *ErddzBaseWinCoinInfo) GetIsSpring() bool {
	if m != nil && m.IsSpring != nil {
		return *m.IsSpring
	}
	return false
}

func (m *ErddzBaseWinCoinInfo) GetHandPokers() []*ClientBasePoker {
	if m != nil {
		return m.HandPokers
	}
	return nil
}

// 本局结果(广播)
type ErddzBcCurrentResult struct {
	Header           *ProtoHeader            `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	WinCoinInfo      []*ErddzBaseWinCoinInfo `protobuf:"bytes,2,rep,name=winCoinInfo" json:"winCoinInfo,omitempty"`
	XXX_unrecognized []byte                  `json:"-"`
}

func (m *ErddzBcCurrentResult) Reset()                    { *m = ErddzBcCurrentResult{} }
func (m *ErddzBcCurrentResult) String() string            { return proto.CompactTextString(m) }
func (*ErddzBcCurrentResult) ProtoMessage()               {}
func (*ErddzBcCurrentResult) Descriptor() ([]byte, []int) { return fileDescriptor19, []int{5} }

func (m *ErddzBcCurrentResult) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *ErddzBcCurrentResult) GetWinCoinInfo() []*ErddzBaseWinCoinInfo {
	if m != nil {
		return m.WinCoinInfo
	}
	return nil
}

type ErddzBaseEndLotteryInfo struct {
	UserId           *uint32 `protobuf:"varint,1,opt,name=userId" json:"userId,omitempty"`
	NickName         *string `protobuf:"bytes,2,opt,name=nickName" json:"nickName,omitempty"`
	BigWin           *bool   `protobuf:"varint,3,opt,name=bigWin" json:"bigWin,omitempty"`
	IsOwner          *bool   `protobuf:"varint,4,opt,name=isOwner" json:"isOwner,omitempty"`
	WinCoin          *int64  `protobuf:"varint,5,opt,name=winCoin" json:"winCoin,omitempty"`
	MaxWinCoin       *int32  `protobuf:"varint,6,opt,name=maxWinCoin" json:"maxWinCoin,omitempty"`
	CountBomb        *int32  `protobuf:"varint,7,opt,name=countBomb" json:"countBomb,omitempty"`
	CountWin         *int32  `protobuf:"varint,8,opt,name=countWin" json:"countWin,omitempty"`
	CountLose        *int32  `protobuf:"varint,9,opt,name=countLose" json:"countLose,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ErddzBaseEndLotteryInfo) Reset()                    { *m = ErddzBaseEndLotteryInfo{} }
func (m *ErddzBaseEndLotteryInfo) String() string            { return proto.CompactTextString(m) }
func (*ErddzBaseEndLotteryInfo) ProtoMessage()               {}
func (*ErddzBaseEndLotteryInfo) Descriptor() ([]byte, []int) { return fileDescriptor19, []int{6} }

func (m *ErddzBaseEndLotteryInfo) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *ErddzBaseEndLotteryInfo) GetNickName() string {
	if m != nil && m.NickName != nil {
		return *m.NickName
	}
	return ""
}

func (m *ErddzBaseEndLotteryInfo) GetBigWin() bool {
	if m != nil && m.BigWin != nil {
		return *m.BigWin
	}
	return false
}

func (m *ErddzBaseEndLotteryInfo) GetIsOwner() bool {
	if m != nil && m.IsOwner != nil {
		return *m.IsOwner
	}
	return false
}

func (m *ErddzBaseEndLotteryInfo) GetWinCoin() int64 {
	if m != nil && m.WinCoin != nil {
		return *m.WinCoin
	}
	return 0
}

func (m *ErddzBaseEndLotteryInfo) GetMaxWinCoin() int32 {
	if m != nil && m.MaxWinCoin != nil {
		return *m.MaxWinCoin
	}
	return 0
}

func (m *ErddzBaseEndLotteryInfo) GetCountBomb() int32 {
	if m != nil && m.CountBomb != nil {
		return *m.CountBomb
	}
	return 0
}

func (m *ErddzBaseEndLotteryInfo) GetCountWin() int32 {
	if m != nil && m.CountWin != nil {
		return *m.CountWin
	}
	return 0
}

func (m *ErddzBaseEndLotteryInfo) GetCountLose() int32 {
	if m != nil && m.CountLose != nil {
		return *m.CountLose
	}
	return 0
}

// 牌局结束(广播)
type ErddzBcEndLottery struct {
	Header           *ProtoHeader               `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	CoinInfo         []*ErddzBaseEndLotteryInfo `protobuf:"bytes,2,rep,name=coinInfo" json:"coinInfo,omitempty"`
	XXX_unrecognized []byte                     `json:"-"`
}

func (m *ErddzBcEndLottery) Reset()                    { *m = ErddzBcEndLottery{} }
func (m *ErddzBcEndLottery) String() string            { return proto.CompactTextString(m) }
func (*ErddzBcEndLottery) ProtoMessage()               {}
func (*ErddzBcEndLottery) Descriptor() ([]byte, []int) { return fileDescriptor19, []int{7} }

func (m *ErddzBcEndLottery) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *ErddzBcEndLottery) GetCoinInfo() []*ErddzBaseEndLotteryInfo {
	if m != nil {
		return m.CoinInfo
	}
	return nil
}

func init() {
	proto.RegisterType((*ErddzReqDissolveDesk)(nil), "ddproto.erddz_req_dissolveDesk")
	proto.RegisterType((*ErddzAckDissolveDesk)(nil), "ddproto.erddz_ack_dissolveDesk")
	proto.RegisterType((*ErddzReqReady)(nil), "ddproto.erddz_req_ready")
	proto.RegisterType((*ErddzAckReady)(nil), "ddproto.erddz_ack_ready")
	proto.RegisterType((*ErddzBaseWinCoinInfo)(nil), "ddproto.erddz_base_winCoinInfo")
	proto.RegisterType((*ErddzBcCurrentResult)(nil), "ddproto.erddz_bc_currentResult")
	proto.RegisterType((*ErddzBaseEndLotteryInfo)(nil), "ddproto.erddz_base_endLotteryInfo")
	proto.RegisterType((*ErddzBcEndLottery)(nil), "ddproto.erddz_bc_endLottery")
}

var fileDescriptor19 = []byte{
	// 460 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x92, 0xdf, 0x6e, 0xd3, 0x30,
	0x14, 0xc6, 0x49, 0xd3, 0xb5, 0xe9, 0x29, 0xb0, 0x91, 0x22, 0x64, 0x76, 0xc3, 0x14, 0x71, 0xc1,
	0x55, 0x2f, 0x10, 0x4f, 0x30, 0x26, 0xb1, 0x49, 0xd3, 0xa8, 0x40, 0x62, 0x12, 0x17, 0x44, 0xa9,
	0x7d, 0x58, 0xad, 0x36, 0x76, 0x66, 0x3b, 0x8c, 0xf1, 0x6a, 0xbc, 0x14, 0x8f, 0xc0, 0xb1, 0xd3,
	0xa4, 0xad, 0xd8, 0xc5, 0x04, 0xdc, 0x54, 0x3d, 0xff, 0x3e, 0x7f, 0xdf, 0xaf, 0x85, 0x03, 0x34,
	0x42, 0xfc, 0xc8, 0x05, 0xda, 0xe5, 0xb4, 0x32, 0xda, 0xe9, 0x74, 0x28, 0x44, 0xf8, 0x72, 0x38,
	0xe1, 0xba, 0x2c, 0xb5, 0xca, 0xf9, 0x4a, 0xa2, 0x72, 0xcd, 0x34, 0xbb, 0x80, 0x67, 0xcd, 0x85,
	0xc1, 0xeb, 0x5c, 0x48, 0x6b, 0xf5, 0xea, 0x1b, 0x9e, 0xd0, 0x75, 0xfa, 0x12, 0x06, 0x0b, 0x2c,
	0x04, 0x1a, 0x16, 0x1d, 0x45, 0xaf, 0xc6, 0xaf, 0x9f, 0x4e, 0xd7, 0x42, 0xd3, 0x99, 0xff, 0x3c,
	0x0d, 0xb3, 0xf4, 0x31, 0x0c, 0x6a, 0x8b, 0xe6, 0x4c, 0xb0, 0x1e, 0x6d, 0x3d, 0xca, 0xbe, 0xb4,
	0x7a, 0x05, 0x5f, 0xfe, 0x07, 0x3d, 0x5f, 0xfb, 0x2c, 0x54, 0xc7, 0x54, 0xef, 0x65, 0xef, 0x60,
	0x7f, 0xe3, 0xd7, 0xd0, 0xcd, 0xed, 0x5f, 0x1a, 0xed, 0x84, 0xbc, 0xd1, 0x7f, 0x11, 0xfa, 0x15,
	0xb5, 0x91, 0xe7, 0x85, 0xc5, 0xfc, 0x46, 0xaa, 0xb7, 0x5a, 0xaa, 0x33, 0xf5, 0x55, 0xa7, 0x07,
	0x90, 0x28, 0xc9, 0x97, 0x17, 0x45, 0x89, 0x41, 0x72, 0xf4, 0x47, 0xbc, 0x27, 0x30, 0xf2, 0x57,
	0x9f, 0x8a, 0x55, 0x8d, 0x4d, 0xc2, 0x74, 0x1f, 0x86, 0x6b, 0x0d, 0xd6, 0xa7, 0x46, 0x9c, 0x3e,
	0x84, 0x3e, 0xf7, 0xd5, 0x5e, 0xa8, 0x68, 0x2c, 0xed, 0x89, 0xfc, 0xbc, 0xa8, 0xd9, 0x80, 0x1a,
	0x89, 0x1f, 0x9b, 0xc2, 0x21, 0x1b, 0x86, 0xeb, 0x09, 0x8c, 0x89, 0x17, 0x37, 0xb2, 0x72, 0x52,
	0x2b, 0x96, 0x84, 0x57, 0x69, 0x65, 0xae, 0xcb, 0x39, 0x1b, 0x85, 0x15, 0x72, 0x25, 0xed, 0xc7,
	0xca, 0x48, 0x75, 0xc5, 0x20, 0x48, 0x4c, 0x01, 0x16, 0x85, 0x12, 0x33, 0xbd, 0x44, 0x63, 0xd9,
	0xf8, 0x28, 0xa6, 0xf0, 0x87, 0x5d, 0xf8, 0xe6, 0xff, 0xd2, 0xa4, 0xab, 0xfc, 0x4a, 0xe6, 0xba,
	0xc4, 0x3c, 0xe7, 0xb5, 0x31, 0x34, 0xfe, 0x80, 0xb6, 0x5e, 0xb9, 0x7b, 0x22, 0x7c, 0x03, 0xe3,
	0x2d, 0x4c, 0x84, 0xc2, 0x3f, 0xf8, 0xa2, 0x5b, 0xbd, 0x9b, 0x66, 0xf6, 0x33, 0x82, 0xe7, 0x5b,
	0x23, 0x54, 0xe2, 0x5c, 0x3b, 0x87, 0xe6, 0x36, 0xb0, 0xde, 0x90, 0x8d, 0x02, 0xd9, 0x6d, 0xf6,
	0xbd, 0x96, 0xfd, 0x5c, 0x5e, 0x5d, 0x12, 0xc9, 0x38, 0xa4, 0x0e, 0x24, 0xdf, 0xdf, 0x28, 0x32,
	0xdb, 0x6f, 0x1b, 0x2d, 0xf9, 0x86, 0x75, 0x0a, 0x50, 0x16, 0xdf, 0x2f, 0xd7, 0xbd, 0x41, 0xa0,
	0x47, 0xbf, 0x18, 0xd7, 0xb5, 0x72, 0xc7, 0x1e, 0xe8, 0xb0, 0x05, 0x1a, 0x5a, 0x5e, 0x3a, 0xd9,
	0x59, 0x3a, 0xd7, 0x16, 0x1b, 0xea, 0xd9, 0x35, 0x4c, 0x3a, 0x66, 0x1b, 0xeb, 0xf7, 0x06, 0x96,
	0xf0, 0x5d, 0x5a, 0xd9, 0x5d, 0xb4, 0x76, 0x91, 0x1c, 0xf7, 0x4e, 0xe3, 0xd9, 0x83, 0xdf, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x31, 0xdf, 0xbf, 0xfb, 0x12, 0x04, 0x00, 0x00,
}
