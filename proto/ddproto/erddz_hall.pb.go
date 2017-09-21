// Code generated by protoc-gen-go. DO NOT EDIT.
// source: erddz_hall.proto

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

// Ignoring public import of erddz_base_roomTypeInfo from erddz_base.proto

// Ignoring public import of erddz_base_playerInfo from erddz_base.proto

// Ignoring public import of erddz_base_playerRateInfo from erddz_base.proto

// Ignoring public import of erddz_base_commonRateInfo from erddz_base.proto

// Ignoring public import of erddz_base_timerInfo from erddz_base.proto

// Ignoring public import of erddz_base_deskInfo from erddz_base.proto

// Ignoring public import of erddz_enum_protoId from erddz_base.proto

// Ignoring public import of erddz_enum_errorCode from erddz_base.proto

// Ignoring public import of erddz_enum_paiType from erddz_base.proto

// Ignoring public import of erddz_enum_actType from erddz_base.proto

// Ignoring public import of erddz_enum_deskStatus from erddz_base.proto

// Ignoring public import of erddz_enum_playerStatus from erddz_base.proto

// Ignoring public import of erddz_enum_roomType from erddz_base.proto

// Ignoring public import of erddz_enum_coinRoomLevel from erddz_base.proto

// 创建房间
type ErddzReqCreateDesk struct {
	Header           *ProtoHeader           `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	RoomTypeInfo     *ErddzBaseRoomTypeInfo `protobuf:"bytes,2,opt,name=roomTypeInfo" json:"roomTypeInfo,omitempty"`
	XXX_unrecognized []byte                 `json:"-"`
}

func (m *ErddzReqCreateDesk) Reset()                    { *m = ErddzReqCreateDesk{} }
func (m *ErddzReqCreateDesk) String() string            { return proto.CompactTextString(m) }
func (*ErddzReqCreateDesk) ProtoMessage()               {}
func (*ErddzReqCreateDesk) Descriptor() ([]byte, []int) { return fileDescriptor24, []int{0} }

func (m *ErddzReqCreateDesk) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *ErddzReqCreateDesk) GetRoomTypeInfo() *ErddzBaseRoomTypeInfo {
	if m != nil {
		return m.RoomTypeInfo
	}
	return nil
}

// 创建房间回复的信息
type ErddzAckCreateDesk struct {
	Header           *ProtoHeader           `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	DeskId           *int32                 `protobuf:"varint,2,opt,name=deskId" json:"deskId,omitempty"`
	Password         *string                `protobuf:"bytes,3,opt,name=password" json:"password,omitempty"`
	CreateFee        *int64                 `protobuf:"varint,5,opt,name=createFee" json:"createFee,omitempty"`
	RoomTypeInfo     *ErddzBaseRoomTypeInfo `protobuf:"bytes,6,opt,name=roomTypeInfo" json:"roomTypeInfo,omitempty"`
	XXX_unrecognized []byte                 `json:"-"`
}

func (m *ErddzAckCreateDesk) Reset()                    { *m = ErddzAckCreateDesk{} }
func (m *ErddzAckCreateDesk) String() string            { return proto.CompactTextString(m) }
func (*ErddzAckCreateDesk) ProtoMessage()               {}
func (*ErddzAckCreateDesk) Descriptor() ([]byte, []int) { return fileDescriptor24, []int{1} }

func (m *ErddzAckCreateDesk) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *ErddzAckCreateDesk) GetDeskId() int32 {
	if m != nil && m.DeskId != nil {
		return *m.DeskId
	}
	return 0
}

func (m *ErddzAckCreateDesk) GetPassword() string {
	if m != nil && m.Password != nil {
		return *m.Password
	}
	return ""
}

func (m *ErddzAckCreateDesk) GetCreateFee() int64 {
	if m != nil && m.CreateFee != nil {
		return *m.CreateFee
	}
	return 0
}

func (m *ErddzAckCreateDesk) GetRoomTypeInfo() *ErddzBaseRoomTypeInfo {
	if m != nil {
		return m.RoomTypeInfo
	}
	return nil
}

// 客户端请求进入 room, 服务器返回DdzSendGameInfo
type ErddzReqEnterDesk struct {
	Header           *ProtoHeader            `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	PassWord         *string                 `protobuf:"bytes,2,opt,name=PassWord" json:"PassWord,omitempty"`
	UserId           *uint32                 `protobuf:"varint,3,opt,name=userId" json:"userId,omitempty"`
	RoomType         *ErddzEnumRoomType      `protobuf:"varint,4,opt,name=roomType,enum=ddproto.ErddzEnumRoomType" json:"roomType,omitempty"`
	CoinRoomLevel    *ErddzEnumCoinRoomLevel `protobuf:"varint,5,opt,name=coinRoomLevel,enum=ddproto.ErddzEnumCoinRoomLevel" json:"coinRoomLevel,omitempty"`
	XXX_unrecognized []byte                  `json:"-"`
}

func (m *ErddzReqEnterDesk) Reset()                    { *m = ErddzReqEnterDesk{} }
func (m *ErddzReqEnterDesk) String() string            { return proto.CompactTextString(m) }
func (*ErddzReqEnterDesk) ProtoMessage()               {}
func (*ErddzReqEnterDesk) Descriptor() ([]byte, []int) { return fileDescriptor24, []int{2} }

func (m *ErddzReqEnterDesk) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *ErddzReqEnterDesk) GetPassWord() string {
	if m != nil && m.PassWord != nil {
		return *m.PassWord
	}
	return ""
}

func (m *ErddzReqEnterDesk) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *ErddzReqEnterDesk) GetRoomType() ErddzEnumRoomType {
	if m != nil && m.RoomType != nil {
		return *m.RoomType
	}
	return ErddzEnumRoomType_ERDDZ_FRIEND
}

func (m *ErddzReqEnterDesk) GetCoinRoomLevel() ErddzEnumCoinRoomLevel {
	if m != nil && m.CoinRoomLevel != nil {
		return *m.CoinRoomLevel
	}
	return ErddzEnumCoinRoomLevel_ERDDZ_LV1
}

type ErddzAckEnterDesk struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *ErddzAckEnterDesk) Reset()                    { *m = ErddzAckEnterDesk{} }
func (m *ErddzAckEnterDesk) String() string            { return proto.CompactTextString(m) }
func (*ErddzAckEnterDesk) ProtoMessage()               {}
func (*ErddzAckEnterDesk) Descriptor() ([]byte, []int) { return fileDescriptor24, []int{3} }

func (m *ErddzAckEnterDesk) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func init() {
	proto.RegisterType((*ErddzReqCreateDesk)(nil), "ddproto.erddz_req_createDesk")
	proto.RegisterType((*ErddzAckCreateDesk)(nil), "ddproto.erddz_ack_createDesk")
	proto.RegisterType((*ErddzReqEnterDesk)(nil), "ddproto.erddz_req_enterDesk")
	proto.RegisterType((*ErddzAckEnterDesk)(nil), "ddproto.erddz_ack_enterDesk")
}

func init() { proto.RegisterFile("erddz_hall.proto", fileDescriptor24) }

var fileDescriptor24 = []byte{
	// 331 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x91, 0xbd, 0x4e, 0xc3, 0x30,
	0x14, 0x85, 0x71, 0x4b, 0x4b, 0x6b, 0x68, 0x85, 0xdc, 0x0a, 0x45, 0x51, 0x87, 0xd0, 0x29, 0x03,
	0xea, 0xd0, 0x89, 0x19, 0x2a, 0x68, 0x25, 0x86, 0xc8, 0x42, 0x62, 0x8c, 0x4c, 0x7c, 0x51, 0xab,
	0x24, 0x76, 0xb0, 0x53, 0x10, 0x3c, 0x02, 0xcf, 0xc8, 0x93, 0x30, 0xa1, 0xc4, 0xf9, 0x21, 0xc0,
	0x94, 0x25, 0xca, 0xb9, 0x3e, 0x27, 0x3e, 0xdf, 0x0d, 0x3e, 0x05, 0xc5, 0xf9, 0xbb, 0xbf, 0x65,
	0x51, 0xb4, 0x48, 0x94, 0x4c, 0x25, 0x39, 0xe2, 0x3c, 0x7f, 0xb1, 0x27, 0x81, 0x8c, 0x63, 0x29,
	0xfc, 0x20, 0xda, 0x81, 0x48, 0xcd, 0xa9, 0x5d, 0xf8, 0x1f, 0x99, 0x06, 0x33, 0x99, 0x7f, 0x20,
	0x3c, 0x35, 0x43, 0x05, 0xcf, 0x7e, 0xa0, 0x80, 0xa5, 0xb0, 0x02, 0x1d, 0x92, 0x0b, 0xdc, 0xdf,
	0x02, 0xe3, 0xa0, 0x2c, 0xe4, 0x20, 0xf7, 0x78, 0x39, 0x5d, 0x14, 0x5f, 0x5e, 0x78, 0xd9, 0x73,
	0x9d, 0x9f, 0xd1, 0xc2, 0x43, 0x56, 0xf8, 0x44, 0x49, 0x19, 0xdf, 0xbf, 0x25, 0xb0, 0x11, 0x4f,
	0xd2, 0xea, 0xe4, 0x19, 0xa7, 0xca, 0xd4, 0xf7, 0xfa, 0x3f, 0x7d, 0xb4, 0x91, 0x9a, 0x7f, 0x56,
	0x65, 0x58, 0x10, 0xb6, 0x2f, 0x73, 0x86, 0xfb, 0x1c, 0x74, 0xb8, 0xe1, 0x79, 0x8d, 0x1e, 0x2d,
	0x14, 0xb1, 0xf1, 0x20, 0x61, 0x5a, 0xbf, 0x4a, 0xc5, 0xad, 0xae, 0x83, 0xdc, 0x21, 0xad, 0x34,
	0x99, 0xe1, 0xa1, 0xb9, 0xef, 0x06, 0xc0, 0xea, 0x39, 0xc8, 0xed, 0xd2, 0x7a, 0xf0, 0x07, 0xaf,
	0xdf, 0x0a, 0xef, 0x0b, 0xe1, 0x49, 0xbd, 0x6b, 0x10, 0x29, 0xa8, 0x16, 0x74, 0x36, 0x1e, 0x78,
	0x4c, 0xeb, 0x87, 0x8c, 0xa2, 0x63, 0x28, 0x4a, 0x9d, 0x91, 0xef, 0x35, 0xa8, 0x8d, 0xe1, 0x1b,
	0xd1, 0x42, 0x91, 0x4b, 0x3c, 0x28, 0x9b, 0x58, 0x87, 0x0e, 0x72, 0xc7, 0xcb, 0xd9, 0xaf, 0xee,
	0x20, 0xf6, 0x71, 0xd5, 0x9d, 0x56, 0x6e, 0x72, 0x8b, 0x47, 0x81, 0xdc, 0x09, 0x2a, 0x65, 0x7c,
	0x07, 0x2f, 0x10, 0xe5, 0xbb, 0x19, 0x2f, 0xcf, 0xff, 0x8b, 0x37, 0x8c, 0xb4, 0x99, 0x9b, 0x5f,
	0x97, 0xec, 0xd9, 0xaf, 0x6d, 0xc9, 0x7e, 0xd5, 0x59, 0x77, 0xbd, 0x03, 0x0f, 0x7d, 0x07, 0x00,
	0x00, 0xff, 0xff, 0xbf, 0x6a, 0x9e, 0x04, 0xf8, 0x02, 0x00, 0x00,
}
