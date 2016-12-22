// Code generated by protoc-gen-go.
// source: common_client_award.proto
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

// Ignoring public import of common_req_gameLogin from common_client.proto

// Ignoring public import of common_ack_gameLogin from common_client.proto

// Ignoring public import of common_req_logout from common_client.proto

// Ignoring public import of common_ack_logout from common_client.proto

// Ignoring public import of common_req_feedback from common_client.proto

// Ignoring public import of client_base_poker from common_client.proto

// Ignoring public import of common_req_message from common_client.proto

// Ignoring public import of common_bc_message from common_client.proto

// Ignoring public import of common_req_notice from common_client.proto

// Ignoring public import of common_ack_notice from common_client.proto

// Ignoring public import of common_enum_pokerColor from common_client.proto

// 请求得到在线奖励
type AwardReqOnline struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *AwardReqOnline) Reset()                    { *m = AwardReqOnline{} }
func (m *AwardReqOnline) String() string            { return proto.CompactTextString(m) }
func (*AwardReqOnline) ProtoMessage()               {}
func (*AwardReqOnline) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *AwardReqOnline) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

type WardAckOnline struct {
	Coin             *int64 `protobuf:"varint,1,opt,name=coin" json:"coin,omitempty"`
	Duration         *int64 `protobuf:"varint,2,opt,name=duration" json:"duration,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *WardAckOnline) Reset()                    { *m = WardAckOnline{} }
func (m *WardAckOnline) String() string            { return proto.CompactTextString(m) }
func (*WardAckOnline) ProtoMessage()               {}
func (*WardAckOnline) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

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

func init() {
	proto.RegisterType((*AwardReqOnline)(nil), "ddproto.award_req_online")
	proto.RegisterType((*WardAckOnline)(nil), "ddproto.ward_ack_online")
}

var fileDescriptor1 = []byte{
	// 146 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x92, 0x4c, 0xce, 0xcf, 0xcd,
	0xcd, 0xcf, 0x8b, 0x4f, 0xce, 0xc9, 0x4c, 0xcd, 0x2b, 0x89, 0x4f, 0x2c, 0x4f, 0x2c, 0x4a, 0xd1,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4f, 0x49, 0x01, 0x33, 0xa4, 0x84, 0x51, 0xd4, 0x40,
	0x64, 0x95, 0x2c, 0xb8, 0x04, 0xc0, 0x8a, 0xe3, 0x8b, 0x52, 0x0b, 0xe3, 0xf3, 0xf3, 0x72, 0x32,
	0xf3, 0x52, 0x85, 0x54, 0xb8, 0xd8, 0x32, 0x52, 0x13, 0x53, 0x52, 0x8b, 0x24, 0x18, 0x15, 0x18,
	0x35, 0xb8, 0x8d, 0x44, 0xf4, 0xa0, 0x46, 0xe8, 0x05, 0x80, 0x48, 0x0f, 0xb0, 0x9c, 0x92, 0x21,
	0x17, 0x3f, 0x58, 0x63, 0x62, 0x72, 0x36, 0x4c, 0x23, 0x0f, 0x17, 0x4b, 0x72, 0x7e, 0x66, 0x1e,
	0x58, 0x1b, 0xb3, 0x90, 0x00, 0x17, 0x47, 0x4a, 0x69, 0x51, 0x62, 0x49, 0x66, 0x7e, 0x9e, 0x04,
	0x13, 0x48, 0x24, 0x80, 0x01, 0x10, 0x00, 0x00, 0xff, 0xff, 0x86, 0xcf, 0xb0, 0x46, 0xa8, 0x00,
	0x00, 0x00,
}