// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common_client_award.proto

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

// Ignoring public import of common_req_kickout from common_client.proto

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

// Ignoring public import of common_req_upload_location from common_client.proto

// Ignoring public import of common_bc_leaveTimeout from common_client.proto

// Ignoring public import of common_desk_by_agent from common_client.proto

// Ignoring public import of common_req_list_coin_desk from common_client.proto

// Ignoring public import of common_ack_list_coin_desk from common_client.proto

// Ignoring public import of CommonCoinDeskInfo from common_client.proto

// Ignoring public import of createCoinDeskInfo from common_client.proto

// Ignoring public import of common_enum_reg from common_client.proto

// Ignoring public import of common_enum_os_type from common_client.proto

// Ignoring public import of common_enum_pokerColor from common_client.proto

// 请求得到在线奖励
type AwardReqOnline struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *AwardReqOnline) Reset()                    { *m = AwardReqOnline{} }
func (m *AwardReqOnline) String() string            { return proto.CompactTextString(m) }
func (*AwardReqOnline) ProtoMessage()               {}
func (*AwardReqOnline) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

func (m *AwardReqOnline) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

type WardAckOnline struct {
	Coin             *int64 `protobuf:"varint,1,opt,name=coin" json:"coin,omitempty"`
	Duration         *int64 `protobuf:"varint,2,opt,name=duration" json:"duration,omitempty"`
	ChangeCoin       *int64 `protobuf:"varint,3,opt,name=changeCoin" json:"changeCoin,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *WardAckOnline) Reset()                    { *m = WardAckOnline{} }
func (m *WardAckOnline) String() string            { return proto.CompactTextString(m) }
func (*WardAckOnline) ProtoMessage()               {}
func (*WardAckOnline) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{1} }

func (m *WardAckOnline) GetCoin() int64 {
	if m != nil && m.Coin != nil {
		return *m.Coin
	}
	return 0
}

func (m *WardAckOnline) GetDuration() int64 {
	if m != nil && m.Duration != nil {
		return *m.Duration
	}
	return 0
}

func (m *WardAckOnline) GetChangeCoin() int64 {
	if m != nil && m.ChangeCoin != nil {
		return *m.ChangeCoin
	}
	return 0
}

// 请求得到在线奖励
type AwardReqOnlineInfo struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *AwardReqOnlineInfo) Reset()                    { *m = AwardReqOnlineInfo{} }
func (m *AwardReqOnlineInfo) String() string            { return proto.CompactTextString(m) }
func (*AwardReqOnlineInfo) ProtoMessage()               {}
func (*AwardReqOnlineInfo) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{2} }

func (m *AwardReqOnlineInfo) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

type AwardAckOnlineInfo struct {
	Duration         *int64 `protobuf:"varint,1,opt,name=duration" json:"duration,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *AwardAckOnlineInfo) Reset()                    { *m = AwardAckOnlineInfo{} }
func (m *AwardAckOnlineInfo) String() string            { return proto.CompactTextString(m) }
func (*AwardAckOnlineInfo) ProtoMessage()               {}
func (*AwardAckOnlineInfo) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{3} }

func (m *AwardAckOnlineInfo) GetDuration() int64 {
	if m != nil && m.Duration != nil {
		return *m.Duration
	}
	return 0
}

// 请求得到新手奖励
type AwardReqGetNewUser struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *AwardReqGetNewUser) Reset()                    { *m = AwardReqGetNewUser{} }
func (m *AwardReqGetNewUser) String() string            { return proto.CompactTextString(m) }
func (*AwardReqGetNewUser) ProtoMessage()               {}
func (*AwardReqGetNewUser) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{4} }

func (m *AwardReqGetNewUser) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

// 回复
type AwardAckGetNewUser struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *AwardAckGetNewUser) Reset()                    { *m = AwardAckGetNewUser{} }
func (m *AwardAckGetNewUser) String() string            { return proto.CompactTextString(m) }
func (*AwardAckGetNewUser) ProtoMessage()               {}
func (*AwardAckGetNewUser) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{5} }

func (m *AwardAckGetNewUser) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func init() {
	proto.RegisterType((*AwardReqOnline)(nil), "ddproto.award_req_online")
	proto.RegisterType((*WardAckOnline)(nil), "ddproto.ward_ack_online")
	proto.RegisterType((*AwardReqOnlineInfo)(nil), "ddproto.award_req_onlineInfo")
	proto.RegisterType((*AwardAckOnlineInfo)(nil), "ddproto.award_ack_onlineInfo")
	proto.RegisterType((*AwardReqGetNewUser)(nil), "ddproto.award_req_getNewUser")
	proto.RegisterType((*AwardAckGetNewUser)(nil), "ddproto.award_ack_getNewUser")
}

func init() { proto.RegisterFile("common_client_award.proto", fileDescriptor4) }

var fileDescriptor4 = []byte{
	// 219 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x92, 0x4c, 0xce, 0xcf, 0xcd,
	0xcd, 0xcf, 0x8b, 0x4f, 0xce, 0xc9, 0x4c, 0xcd, 0x2b, 0x89, 0x4f, 0x2c, 0x4f, 0x2c, 0x4a, 0xd1,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4f, 0x49, 0x01, 0x33, 0xa4, 0x84, 0x51, 0xd4, 0x40,
	0x64, 0x95, 0x1c, 0xb8, 0x04, 0xc0, 0x8a, 0xe3, 0x8b, 0x52, 0x0b, 0xe3, 0xf3, 0xf3, 0x72, 0x32,
	0xf3, 0x52, 0x85, 0x74, 0xb8, 0xd8, 0x32, 0x52, 0x13, 0x53, 0x52, 0x8b, 0x24, 0x18, 0x15, 0x18,
	0x35, 0xb8, 0x8d, 0x44, 0xf4, 0xa0, 0x46, 0xe8, 0x05, 0x80, 0x48, 0x0f, 0xb0, 0x5c, 0x10, 0x54,
	0x8d, 0x52, 0x22, 0x17, 0x3f, 0xd8, 0x80, 0xc4, 0xe4, 0x6c, 0x98, 0x01, 0x42, 0x5c, 0x2c, 0xc9,
	0xf9, 0x99, 0x79, 0x60, 0xed, 0xcc, 0x41, 0x60, 0xb6, 0x90, 0x14, 0x17, 0x47, 0x4a, 0x69, 0x51,
	0x62, 0x49, 0x66, 0x7e, 0x9e, 0x04, 0x13, 0x58, 0x1c, 0xce, 0x17, 0x92, 0xe3, 0xe2, 0x4a, 0xce,
	0x48, 0xcc, 0x4b, 0x4f, 0x75, 0x06, 0xe9, 0x62, 0x06, 0xcb, 0x22, 0x89, 0x28, 0xb9, 0x70, 0x89,
	0xa0, 0x3b, 0xd2, 0x33, 0x2f, 0x2d, 0x9f, 0x44, 0x87, 0x1a, 0xc1, 0x4c, 0x41, 0xb8, 0x14, 0x6c,
	0x0a, 0xb2, 0xcb, 0x18, 0x51, 0x5d, 0x86, 0x6a, 0x73, 0x7a, 0x6a, 0x89, 0x5f, 0x6a, 0x79, 0x68,
	0x71, 0x6a, 0x11, 0x89, 0x36, 0xbb, 0x20, 0xdb, 0x4c, 0xae, 0x29, 0x01, 0x0c, 0x80, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x7f, 0x65, 0x25, 0xdb, 0xe6, 0x01, 0x00, 0x00,
}
