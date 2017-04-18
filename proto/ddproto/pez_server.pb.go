// Code generated by protoc-gen-go.
// source: pez_server.proto
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

// Ignoring public import of CardInfo from mj_common.proto

// Ignoring public import of ErrorCode from mj_common.proto

// Ignoring public import of mj_enum_color from mj_common.proto

// Ignoring public import of mj_enum_gangType from mj_common.proto

// Ignoring public import of mj_enum_huType from mj_common.proto

// Ignoring public import of mj_enum_composeCardType from mj_common.proto

// Ignoring public import of mj_enum_paiType from mj_common.proto

// Ignoring public import of mj_enum_userGameStatus from mj_common.proto

// Ignoring public import of mj_enum_deskGameStatus from mj_common.proto

// Ignoring public import of MJRoomType from mj_common.proto

// 拼二张牌型
type PezEnum_PEZTYPE int32

const (
	PezEnum_PEZTYPE_PEZTYPE_MAXPAI PezEnum_PEZTYPE = 1
	PezEnum_PEZTYPE_PEZTYPE_DUIZI  PezEnum_PEZTYPE = 2
	PezEnum_PEZTYPE_PEZTYPE_LAIZI  PezEnum_PEZTYPE = 3
)

var PezEnum_PEZTYPE_name = map[int32]string{
	1: "PEZTYPE_MAXPAI",
	2: "PEZTYPE_DUIZI",
	3: "PEZTYPE_LAIZI",
}
var PezEnum_PEZTYPE_value = map[string]int32{
	"PEZTYPE_MAXPAI": 1,
	"PEZTYPE_DUIZI":  2,
	"PEZTYPE_LAIZI":  3,
}

func (x PezEnum_PEZTYPE) Enum() *PezEnum_PEZTYPE {
	p := new(PezEnum_PEZTYPE)
	*p = x
	return p
}
func (x PezEnum_PEZTYPE) String() string {
	return proto.EnumName(PezEnum_PEZTYPE_name, int32(x))
}
func (x *PezEnum_PEZTYPE) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(PezEnum_PEZTYPE_value, data, "PezEnum_PEZTYPE")
	if err != nil {
		return err
	}
	*x = PezEnum_PEZTYPE(value)
	return nil
}
func (PezEnum_PEZTYPE) EnumDescriptor() ([]byte, []int) { return fileDescriptor39, []int{0} }

func init() {
	proto.RegisterEnum("ddproto.PezEnum_PEZTYPE", PezEnum_PEZTYPE_name, PezEnum_PEZTYPE_value)
}

var fileDescriptor39 = []byte{
	// 128 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x28, 0x48, 0xad, 0x8a,
	0x2f, 0x4e, 0x2d, 0x2a, 0x4b, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4f, 0x49,
	0x01, 0x33, 0xa4, 0x84, 0x93, 0xf3, 0x73, 0x73, 0xf3, 0xf3, 0xe2, 0x93, 0x73, 0x32, 0x53, 0xf3,
	0x4a, 0x20, 0xb2, 0x52, 0xfc, 0xb9, 0x59, 0xf1, 0x10, 0x71, 0x88, 0x80, 0x96, 0x0f, 0xc4, 0x88,
	0xd4, 0xbc, 0xd2, 0xdc, 0xf8, 0x00, 0xd7, 0xa8, 0x90, 0xc8, 0x00, 0x57, 0x21, 0x21, 0x2e, 0x3e,
	0x28, 0x33, 0xde, 0xd7, 0x31, 0x22, 0xc0, 0xd1, 0x53, 0x80, 0x51, 0x48, 0x90, 0x8b, 0x17, 0x26,
	0xe6, 0x12, 0xea, 0x19, 0xe5, 0x29, 0xc0, 0x84, 0x2c, 0xe4, 0xe3, 0x08, 0x12, 0x62, 0x0e, 0x60,
	0x08, 0x60, 0x04, 0x04, 0x00, 0x00, 0xff, 0xff, 0x32, 0xd5, 0x00, 0x5f, 0x93, 0x00, 0x00, 0x00,
}
