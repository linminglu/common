// Code generated by protoc-gen-go. DO NOT EDIT.
// source: phz_base.proto

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

// Ignoring public import of common_enum_reg from common_client.proto

// Ignoring public import of common_enum_os_type from common_client.proto

// Ignoring public import of common_enum_pokerColor from common_client.proto

type PhzEnumProtoId int32

const (
	PhzEnumProtoId_PHZ_PID_HEARTBEAT             PhzEnumProtoId = 0
	PhzEnumProtoId_PHZ_PID_CREATEDESK_REQ        PhzEnumProtoId = 1
	PhzEnumProtoId_PHZ_PID_ENTERDESK_REQ         PhzEnumProtoId = 2
	PhzEnumProtoId_PHZ_PID_DESK_ACK              PhzEnumProtoId = 3
	PhzEnumProtoId_PHZ_PID_GAMEINFO_REQ          PhzEnumProtoId = 4
	PhzEnumProtoId_PHZ_PID_GAMEINFO              PhzEnumProtoId = 5
	PhzEnumProtoId_PHZ_PID_READY_REQ             PhzEnumProtoId = 6
	PhzEnumProtoId_PHZ_PID_READY_ACK             PhzEnumProtoId = 7
	PhzEnumProtoId_PHZ_PID_OPENING               PhzEnumProtoId = 8
	PhzEnumProtoId_PHZ_PID_SENDCARDS             PhzEnumProtoId = 9
	PhzEnumProtoId_PHZ_PID_OUTCARD_REQ           PhzEnumProtoId = 10
	PhzEnumProtoId_PHZ_PID_OUTCARD_ACK           PhzEnumProtoId = 11
	PhzEnumProtoId_PHZ_PID_MOPAI                 PhzEnumProtoId = 12
	PhzEnumProtoId_PHZ_PID_OVERTURN              PhzEnumProtoId = 13
	PhzEnumProtoId_PHZ_PID_CANPENG               PhzEnumProtoId = 14
	PhzEnumProtoId_PHZ_PID_PENG_REQ              PhzEnumProtoId = 15
	PhzEnumProtoId_PHZ_PID_PENG_ACK              PhzEnumProtoId = 16
	PhzEnumProtoId_PHZ_PID_CHI_REQ               PhzEnumProtoId = 17
	PhzEnumProtoId_PHZ_PID_CANCHI                PhzEnumProtoId = 18
	PhzEnumProtoId_PHZ_PID_CHI_ACK               PhzEnumProtoId = 19
	PhzEnumProtoId_PHZ_PID_BIPAI_REQ             PhzEnumProtoId = 20
	PhzEnumProtoId_PHZ_PID_BIPAI_ACK             PhzEnumProtoId = 21
	PhzEnumProtoId_PHZ_PID_TIPAI                 PhzEnumProtoId = 22
	PhzEnumProtoId_PHZ_PID_WEIPAI                PhzEnumProtoId = 23
	PhzEnumProtoId_PHZ_PID_PAOPAI                PhzEnumProtoId = 24
	PhzEnumProtoId_PHZ_PID_HUPAI_REQ             PhzEnumProtoId = 25
	PhzEnumProtoId_PHZ_PID_CANHUPAI              PhzEnumProtoId = 26
	PhzEnumProtoId_PHZ_PID_HUPAI_ACK             PhzEnumProtoId = 27
	PhzEnumProtoId_PHZ_PID_PASS_REQ              PhzEnumProtoId = 28
	PhzEnumProtoId_PHZ_PID_PASS_ACK              PhzEnumProtoId = 29
	PhzEnumProtoId_PHZ_PID_CURRENTRESULT         PhzEnumProtoId = 30
	PhzEnumProtoId_PHZ_PID_ENDRESULT             PhzEnumProtoId = 31
	PhzEnumProtoId_PHZ_PID_APPLYDISSOLVE_REQ     PhzEnumProtoId = 32
	PhzEnumProtoId_PHZ_PID_APPLYDISSOLVE_ACK     PhzEnumProtoId = 33
	PhzEnumProtoId_PHZ_PID_APPLYDISSOLVEBACK_REQ PhzEnumProtoId = 34
	PhzEnumProtoId_PHZ_PID_APPLYDISSOLVEBACK_ACK PhzEnumProtoId = 35
	PhzEnumProtoId_PHZ_PID_DISSOLVEDESK_REQ      PhzEnumProtoId = 36
	PhzEnumProtoId_PHZ_PID_DISSOLVEDESK_ACK      PhzEnumProtoId = 37
	PhzEnumProtoId_PHZ_PID_USERBREAK             PhzEnumProtoId = 38
	PhzEnumProtoId_PHZ_PID_BREAKTIMEOUT          PhzEnumProtoId = 39
	PhzEnumProtoId_PHZ_PID_LEAVEDESK_REQ         PhzEnumProtoId = 40
	PhzEnumProtoId_PHZ_PID_LEAVEDESK_ACK         PhzEnumProtoId = 41
	PhzEnumProtoId_PHZ_PID_KICKUSER              PhzEnumProtoId = 42
	PhzEnumProtoId_PHZ_PID_KICKUSER_BC           PhzEnumProtoId = 43
	PhzEnumProtoId_PHZ_PID_MESSAGE_REQ           PhzEnumProtoId = 44
	PhzEnumProtoId_PHZ_PID_MESSAGE_BC            PhzEnumProtoId = 45
)

var PhzEnumProtoId_name = map[int32]string{
	0:  "PHZ_PID_HEARTBEAT",
	1:  "PHZ_PID_CREATEDESK_REQ",
	2:  "PHZ_PID_ENTERDESK_REQ",
	3:  "PHZ_PID_DESK_ACK",
	4:  "PHZ_PID_GAMEINFO_REQ",
	5:  "PHZ_PID_GAMEINFO",
	6:  "PHZ_PID_READY_REQ",
	7:  "PHZ_PID_READY_ACK",
	8:  "PHZ_PID_OPENING",
	9:  "PHZ_PID_SENDCARDS",
	10: "PHZ_PID_OUTCARD_REQ",
	11: "PHZ_PID_OUTCARD_ACK",
	12: "PHZ_PID_MOPAI",
	13: "PHZ_PID_OVERTURN",
	14: "PHZ_PID_CANPENG",
	15: "PHZ_PID_PENG_REQ",
	16: "PHZ_PID_PENG_ACK",
	17: "PHZ_PID_CHI_REQ",
	18: "PHZ_PID_CANCHI",
	19: "PHZ_PID_CHI_ACK",
	20: "PHZ_PID_BIPAI_REQ",
	21: "PHZ_PID_BIPAI_ACK",
	22: "PHZ_PID_TIPAI",
	23: "PHZ_PID_WEIPAI",
	24: "PHZ_PID_PAOPAI",
	25: "PHZ_PID_HUPAI_REQ",
	26: "PHZ_PID_CANHUPAI",
	27: "PHZ_PID_HUPAI_ACK",
	28: "PHZ_PID_PASS_REQ",
	29: "PHZ_PID_PASS_ACK",
	30: "PHZ_PID_CURRENTRESULT",
	31: "PHZ_PID_ENDRESULT",
	32: "PHZ_PID_APPLYDISSOLVE_REQ",
	33: "PHZ_PID_APPLYDISSOLVE_ACK",
	34: "PHZ_PID_APPLYDISSOLVEBACK_REQ",
	35: "PHZ_PID_APPLYDISSOLVEBACK_ACK",
	36: "PHZ_PID_DISSOLVEDESK_REQ",
	37: "PHZ_PID_DISSOLVEDESK_ACK",
	38: "PHZ_PID_USERBREAK",
	39: "PHZ_PID_BREAKTIMEOUT",
	40: "PHZ_PID_LEAVEDESK_REQ",
	41: "PHZ_PID_LEAVEDESK_ACK",
	42: "PHZ_PID_KICKUSER",
	43: "PHZ_PID_KICKUSER_BC",
	44: "PHZ_PID_MESSAGE_REQ",
	45: "PHZ_PID_MESSAGE_BC",
}
var PhzEnumProtoId_value = map[string]int32{
	"PHZ_PID_HEARTBEAT":             0,
	"PHZ_PID_CREATEDESK_REQ":        1,
	"PHZ_PID_ENTERDESK_REQ":         2,
	"PHZ_PID_DESK_ACK":              3,
	"PHZ_PID_GAMEINFO_REQ":          4,
	"PHZ_PID_GAMEINFO":              5,
	"PHZ_PID_READY_REQ":             6,
	"PHZ_PID_READY_ACK":             7,
	"PHZ_PID_OPENING":               8,
	"PHZ_PID_SENDCARDS":             9,
	"PHZ_PID_OUTCARD_REQ":           10,
	"PHZ_PID_OUTCARD_ACK":           11,
	"PHZ_PID_MOPAI":                 12,
	"PHZ_PID_OVERTURN":              13,
	"PHZ_PID_CANPENG":               14,
	"PHZ_PID_PENG_REQ":              15,
	"PHZ_PID_PENG_ACK":              16,
	"PHZ_PID_CHI_REQ":               17,
	"PHZ_PID_CANCHI":                18,
	"PHZ_PID_CHI_ACK":               19,
	"PHZ_PID_BIPAI_REQ":             20,
	"PHZ_PID_BIPAI_ACK":             21,
	"PHZ_PID_TIPAI":                 22,
	"PHZ_PID_WEIPAI":                23,
	"PHZ_PID_PAOPAI":                24,
	"PHZ_PID_HUPAI_REQ":             25,
	"PHZ_PID_CANHUPAI":              26,
	"PHZ_PID_HUPAI_ACK":             27,
	"PHZ_PID_PASS_REQ":              28,
	"PHZ_PID_PASS_ACK":              29,
	"PHZ_PID_CURRENTRESULT":         30,
	"PHZ_PID_ENDRESULT":             31,
	"PHZ_PID_APPLYDISSOLVE_REQ":     32,
	"PHZ_PID_APPLYDISSOLVE_ACK":     33,
	"PHZ_PID_APPLYDISSOLVEBACK_REQ": 34,
	"PHZ_PID_APPLYDISSOLVEBACK_ACK": 35,
	"PHZ_PID_DISSOLVEDESK_REQ":      36,
	"PHZ_PID_DISSOLVEDESK_ACK":      37,
	"PHZ_PID_USERBREAK":             38,
	"PHZ_PID_BREAKTIMEOUT":          39,
	"PHZ_PID_LEAVEDESK_REQ":         40,
	"PHZ_PID_LEAVEDESK_ACK":         41,
	"PHZ_PID_KICKUSER":              42,
	"PHZ_PID_KICKUSER_BC":           43,
	"PHZ_PID_MESSAGE_REQ":           44,
	"PHZ_PID_MESSAGE_BC":            45,
}

func (x PhzEnumProtoId) Enum() *PhzEnumProtoId {
	p := new(PhzEnumProtoId)
	*p = x
	return p
}
func (x PhzEnumProtoId) String() string {
	return proto.EnumName(PhzEnumProtoId_name, int32(x))
}
func (x *PhzEnumProtoId) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(PhzEnumProtoId_value, data, "PhzEnumProtoId")
	if err != nil {
		return err
	}
	*x = PhzEnumProtoId(value)
	return nil
}
func (PhzEnumProtoId) EnumDescriptor() ([]byte, []int) { return fileDescriptor48, []int{0} }

type PhzEnumRoomType int32

const (
	PhzEnumRoomType_PHZ_ROOMTYPE_FRIEND PhzEnumRoomType = 1
)

var PhzEnumRoomType_name = map[int32]string{
	1: "PHZ_ROOMTYPE_FRIEND",
}
var PhzEnumRoomType_value = map[string]int32{
	"PHZ_ROOMTYPE_FRIEND": 1,
}

func (x PhzEnumRoomType) Enum() *PhzEnumRoomType {
	p := new(PhzEnumRoomType)
	*p = x
	return p
}
func (x PhzEnumRoomType) String() string {
	return proto.EnumName(PhzEnumRoomType_name, int32(x))
}
func (x *PhzEnumRoomType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(PhzEnumRoomType_value, data, "PhzEnumRoomType")
	if err != nil {
		return err
	}
	*x = PhzEnumRoomType(value)
	return nil
}
func (PhzEnumRoomType) EnumDescriptor() ([]byte, []int) { return fileDescriptor48, []int{1} }

type PhzEnumTiType int32

const (
	PhzEnumTiType_PHZ_TITYPE_HAVE_THREE PhzEnumTiType = 1
	PhzEnumTiType_PHZ_TITYPE_HAVE_FOUR  PhzEnumTiType = 2
	PhzEnumTiType_PHZ_TITYPE_HAVE_WEI   PhzEnumTiType = 3
)

var PhzEnumTiType_name = map[int32]string{
	1: "PHZ_TITYPE_HAVE_THREE",
	2: "PHZ_TITYPE_HAVE_FOUR",
	3: "PHZ_TITYPE_HAVE_WEI",
}
var PhzEnumTiType_value = map[string]int32{
	"PHZ_TITYPE_HAVE_THREE": 1,
	"PHZ_TITYPE_HAVE_FOUR":  2,
	"PHZ_TITYPE_HAVE_WEI":   3,
}

func (x PhzEnumTiType) Enum() *PhzEnumTiType {
	p := new(PhzEnumTiType)
	*p = x
	return p
}
func (x PhzEnumTiType) String() string {
	return proto.EnumName(PhzEnumTiType_name, int32(x))
}
func (x *PhzEnumTiType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(PhzEnumTiType_value, data, "PhzEnumTiType")
	if err != nil {
		return err
	}
	*x = PhzEnumTiType(value)
	return nil
}
func (PhzEnumTiType) EnumDescriptor() ([]byte, []int) { return fileDescriptor48, []int{2} }

type PhzEnumPengType int32

const (
	PhzEnumPengType_PHZ_PENGTYPE_WEI        PhzEnumPengType = 1
	PhzEnumPengType_PHZ_PENGTYPE_CHOUWEI    PhzEnumPengType = 2
	PhzEnumPengType_PHZ_PENGTYPE_NORMALPENG PhzEnumPengType = 3
)

var PhzEnumPengType_name = map[int32]string{
	1: "PHZ_PENGTYPE_WEI",
	2: "PHZ_PENGTYPE_CHOUWEI",
	3: "PHZ_PENGTYPE_NORMALPENG",
}
var PhzEnumPengType_value = map[string]int32{
	"PHZ_PENGTYPE_WEI":        1,
	"PHZ_PENGTYPE_CHOUWEI":    2,
	"PHZ_PENGTYPE_NORMALPENG": 3,
}

func (x PhzEnumPengType) Enum() *PhzEnumPengType {
	p := new(PhzEnumPengType)
	*p = x
	return p
}
func (x PhzEnumPengType) String() string {
	return proto.EnumName(PhzEnumPengType_name, int32(x))
}
func (x *PhzEnumPengType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(PhzEnumPengType_value, data, "PhzEnumPengType")
	if err != nil {
		return err
	}
	*x = PhzEnumPengType(value)
	return nil
}
func (PhzEnumPengType) EnumDescriptor() ([]byte, []int) { return fileDescriptor48, []int{3} }

type PhzEnumPaoType int32

const (
	PhzEnumPaoType_PHZ_PAOTYPE_THREE_ONE PhzEnumPaoType = 1
	PhzEnumPaoType_PHZ_PAOTYPE_WEI_ONE   PhzEnumPaoType = 2
	PhzEnumPaoType_PHZ_PAOTYPE_PENG_ONE  PhzEnumPaoType = 3
)

var PhzEnumPaoType_name = map[int32]string{
	1: "PHZ_PAOTYPE_THREE_ONE",
	2: "PHZ_PAOTYPE_WEI_ONE",
	3: "PHZ_PAOTYPE_PENG_ONE",
}
var PhzEnumPaoType_value = map[string]int32{
	"PHZ_PAOTYPE_THREE_ONE": 1,
	"PHZ_PAOTYPE_WEI_ONE":   2,
	"PHZ_PAOTYPE_PENG_ONE":  3,
}

func (x PhzEnumPaoType) Enum() *PhzEnumPaoType {
	p := new(PhzEnumPaoType)
	*p = x
	return p
}
func (x PhzEnumPaoType) String() string {
	return proto.EnumName(PhzEnumPaoType_name, int32(x))
}
func (x *PhzEnumPaoType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(PhzEnumPaoType_value, data, "PhzEnumPaoType")
	if err != nil {
		return err
	}
	*x = PhzEnumPaoType(value)
	return nil
}
func (PhzEnumPaoType) EnumDescriptor() ([]byte, []int) { return fileDescriptor48, []int{4} }

// Desk的游戏状态
type PhzEnumDeskStatus int32

const (
	PhzEnumDeskStatus_PHZDESK_DESKSTATUS_INIT    PhzEnumDeskStatus = 1
	PhzEnumDeskStatus_PHZDESK_DESKSTATUS_READY   PhzEnumDeskStatus = 2
	PhzEnumDeskStatus_PHZDESK_DESKSTATUS_OPENING PhzEnumDeskStatus = 3
	PhzEnumDeskStatus_PHZDESK_DESKSTATUS_PLAY    PhzEnumDeskStatus = 4
	PhzEnumDeskStatus_PHZDESK_DESKSTATUS_LOTTERY PhzEnumDeskStatus = 5
	PhzEnumDeskStatus_PHZDESK_DESKSTATUS_FINISH  PhzEnumDeskStatus = 6
)

var PhzEnumDeskStatus_name = map[int32]string{
	1: "PHZDESK_DESKSTATUS_INIT",
	2: "PHZDESK_DESKSTATUS_READY",
	3: "PHZDESK_DESKSTATUS_OPENING",
	4: "PHZDESK_DESKSTATUS_PLAY",
	5: "PHZDESK_DESKSTATUS_LOTTERY",
	6: "PHZDESK_DESKSTATUS_FINISH",
}
var PhzEnumDeskStatus_value = map[string]int32{
	"PHZDESK_DESKSTATUS_INIT":    1,
	"PHZDESK_DESKSTATUS_READY":   2,
	"PHZDESK_DESKSTATUS_OPENING": 3,
	"PHZDESK_DESKSTATUS_PLAY":    4,
	"PHZDESK_DESKSTATUS_LOTTERY": 5,
	"PHZDESK_DESKSTATUS_FINISH":  6,
}

func (x PhzEnumDeskStatus) Enum() *PhzEnumDeskStatus {
	p := new(PhzEnumDeskStatus)
	*p = x
	return p
}
func (x PhzEnumDeskStatus) String() string {
	return proto.EnumName(PhzEnumDeskStatus_name, int32(x))
}
func (x *PhzEnumDeskStatus) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(PhzEnumDeskStatus_value, data, "PhzEnumDeskStatus")
	if err != nil {
		return err
	}
	*x = PhzEnumDeskStatus(value)
	return nil
}
func (PhzEnumDeskStatus) EnumDescriptor() ([]byte, []int) { return fileDescriptor48, []int{5} }

type PhzEnumUserStatus int32

const (
	// *******User的流程状态*********
	PhzEnumUserStatus_PHZUSER_STATUS_INIT     PhzEnumUserStatus = 1
	PhzEnumUserStatus_PHZUSER_STATUS_READY    PhzEnumUserStatus = 2
	PhzEnumUserStatus_PHZUSER_STATUS_BEGIN    PhzEnumUserStatus = 3
	PhzEnumUserStatus_PHZUSER_STATUS_MOPAI    PhzEnumUserStatus = 4
	PhzEnumUserStatus_PHZUSER_STATUS_CANCHI   PhzEnumUserStatus = 5
	PhzEnumUserStatus_PHZUSER_STATUS_CANPENG  PhzEnumUserStatus = 6
	PhzEnumUserStatus_PHZUSER_STATUS_CANHU    PhzEnumUserStatus = 7
	PhzEnumUserStatus_PHZUSER_STATUS_CHECKING PhzEnumUserStatus = 8
	PhzEnumUserStatus_PHZUSER_STATUS_OUTCARD  PhzEnumUserStatus = 9
	PhzEnumUserStatus_PHZUSER_STATUS_LOTTERY  PhzEnumUserStatus = 10
	PhzEnumUserStatus_PHZUSER_STATUS_FINISH   PhzEnumUserStatus = 11
)

var PhzEnumUserStatus_name = map[int32]string{
	1:  "PHZUSER_STATUS_INIT",
	2:  "PHZUSER_STATUS_READY",
	3:  "PHZUSER_STATUS_BEGIN",
	4:  "PHZUSER_STATUS_MOPAI",
	5:  "PHZUSER_STATUS_CANCHI",
	6:  "PHZUSER_STATUS_CANPENG",
	7:  "PHZUSER_STATUS_CANHU",
	8:  "PHZUSER_STATUS_CHECKING",
	9:  "PHZUSER_STATUS_OUTCARD",
	10: "PHZUSER_STATUS_LOTTERY",
	11: "PHZUSER_STATUS_FINISH",
}
var PhzEnumUserStatus_value = map[string]int32{
	"PHZUSER_STATUS_INIT":     1,
	"PHZUSER_STATUS_READY":    2,
	"PHZUSER_STATUS_BEGIN":    3,
	"PHZUSER_STATUS_MOPAI":    4,
	"PHZUSER_STATUS_CANCHI":   5,
	"PHZUSER_STATUS_CANPENG":  6,
	"PHZUSER_STATUS_CANHU":    7,
	"PHZUSER_STATUS_CHECKING": 8,
	"PHZUSER_STATUS_OUTCARD":  9,
	"PHZUSER_STATUS_LOTTERY":  10,
	"PHZUSER_STATUS_FINISH":   11,
}

func (x PhzEnumUserStatus) Enum() *PhzEnumUserStatus {
	p := new(PhzEnumUserStatus)
	*p = x
	return p
}
func (x PhzEnumUserStatus) String() string {
	return proto.EnumName(PhzEnumUserStatus_name, int32(x))
}
func (x *PhzEnumUserStatus) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(PhzEnumUserStatus_value, data, "PhzEnumUserStatus")
	if err != nil {
		return err
	}
	*x = PhzEnumUserStatus(value)
	return nil
}
func (PhzEnumUserStatus) EnumDescriptor() ([]byte, []int) { return fileDescriptor48, []int{6} }

type PhzPlayPaiIds struct {
	PokerId          []int32 `protobuf:"varint,1,rep,name=pokerId" json:"pokerId,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *PhzPlayPaiIds) Reset()                    { *m = PhzPlayPaiIds{} }
func (m *PhzPlayPaiIds) String() string            { return proto.CompactTextString(m) }
func (*PhzPlayPaiIds) ProtoMessage()               {}
func (*PhzPlayPaiIds) Descriptor() ([]byte, []int) { return fileDescriptor48, []int{0} }

func (m *PhzPlayPaiIds) GetPokerId() []int32 {
	if m != nil {
		return m.PokerId
	}
	return nil
}

type PhzBaseCreateOption struct {
	UserCount        *int32           `protobuf:"varint,1,opt,name=userCount" json:"userCount,omitempty"`
	RoomType         *PhzEnumRoomType `protobuf:"varint,2,opt,name=roomType,enum=ddproto.PhzEnumRoomType" json:"roomType,omitempty"`
	BoardsCout       *int32           `protobuf:"varint,3,opt,name=boardsCout" json:"boardsCout,omitempty"`
	HuXi             *int32           `protobuf:"varint,4,opt,name=huXi" json:"huXi,omitempty"`
	IsDaiKai         *bool            `protobuf:"varint,5,opt,name=isDaiKai" json:"isDaiKai,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *PhzBaseCreateOption) Reset()                    { *m = PhzBaseCreateOption{} }
func (m *PhzBaseCreateOption) String() string            { return proto.CompactTextString(m) }
func (*PhzBaseCreateOption) ProtoMessage()               {}
func (*PhzBaseCreateOption) Descriptor() ([]byte, []int) { return fileDescriptor48, []int{1} }

func (m *PhzBaseCreateOption) GetUserCount() int32 {
	if m != nil && m.UserCount != nil {
		return *m.UserCount
	}
	return 0
}

func (m *PhzBaseCreateOption) GetRoomType() PhzEnumRoomType {
	if m != nil && m.RoomType != nil {
		return *m.RoomType
	}
	return PhzEnumRoomType_PHZ_ROOMTYPE_FRIEND
}

func (m *PhzBaseCreateOption) GetBoardsCout() int32 {
	if m != nil && m.BoardsCout != nil {
		return *m.BoardsCout
	}
	return 0
}

func (m *PhzBaseCreateOption) GetHuXi() int32 {
	if m != nil && m.HuXi != nil {
		return *m.HuXi
	}
	return 0
}

func (m *PhzBaseCreateOption) GetIsDaiKai() bool {
	if m != nil && m.IsDaiKai != nil {
		return *m.IsDaiKai
	}
	return false
}

type PhzBaseRoomInfo struct {
	OwnerId          *uint32          `protobuf:"varint,1,opt,name=ownerId" json:"ownerId,omitempty"`
	Password         *string          `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
	UserCount        *int32           `protobuf:"varint,3,opt,name=userCount" json:"userCount,omitempty"`
	BoardsCout       *int32           `protobuf:"varint,4,opt,name=boardsCout" json:"boardsCout,omitempty"`
	CurrentRound     *int32           `protobuf:"varint,5,opt,name=currentRound" json:"currentRound,omitempty"`
	RoomType         *PhzEnumRoomType `protobuf:"varint,6,opt,name=roomType,enum=ddproto.PhzEnumRoomType" json:"roomType,omitempty"`
	Huxi             *int32           `protobuf:"varint,7,opt,name=huxi" json:"huxi,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *PhzBaseRoomInfo) Reset()                    { *m = PhzBaseRoomInfo{} }
func (m *PhzBaseRoomInfo) String() string            { return proto.CompactTextString(m) }
func (*PhzBaseRoomInfo) ProtoMessage()               {}
func (*PhzBaseRoomInfo) Descriptor() ([]byte, []int) { return fileDescriptor48, []int{2} }

func (m *PhzBaseRoomInfo) GetOwnerId() uint32 {
	if m != nil && m.OwnerId != nil {
		return *m.OwnerId
	}
	return 0
}

func (m *PhzBaseRoomInfo) GetPassword() string {
	if m != nil && m.Password != nil {
		return *m.Password
	}
	return ""
}

func (m *PhzBaseRoomInfo) GetUserCount() int32 {
	if m != nil && m.UserCount != nil {
		return *m.UserCount
	}
	return 0
}

func (m *PhzBaseRoomInfo) GetBoardsCout() int32 {
	if m != nil && m.BoardsCout != nil {
		return *m.BoardsCout
	}
	return 0
}

func (m *PhzBaseRoomInfo) GetCurrentRound() int32 {
	if m != nil && m.CurrentRound != nil {
		return *m.CurrentRound
	}
	return 0
}

func (m *PhzBaseRoomInfo) GetRoomType() PhzEnumRoomType {
	if m != nil && m.RoomType != nil {
		return *m.RoomType
	}
	return PhzEnumRoomType_PHZ_ROOMTYPE_FRIEND
}

func (m *PhzBaseRoomInfo) GetHuxi() int32 {
	if m != nil && m.Huxi != nil {
		return *m.Huxi
	}
	return 0
}

type PhzBaseDeskInfo struct {
	RoomInfo         *PhzBaseRoomInfo `protobuf:"bytes,1,opt,name=roomInfo" json:"roomInfo,omitempty"`
	GameStatus       *int32           `protobuf:"varint,2,opt,name=gameStatus" json:"gameStatus,omitempty"`
	RemainPaiCount   *int32           `protobuf:"varint,3,opt,name=remainPaiCount" json:"remainPaiCount,omitempty"`
	LiangZhang       *int32           `protobuf:"varint,4,opt,name=liangZhang" json:"liangZhang,omitempty"`
	ActiveUser       *uint32          `protobuf:"varint,5,opt,name=activeUser" json:"activeUser,omitempty"`
	IsOutCard        *bool            `protobuf:"varint,6,opt,name=isOutCard" json:"isOutCard,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *PhzBaseDeskInfo) Reset()                    { *m = PhzBaseDeskInfo{} }
func (m *PhzBaseDeskInfo) String() string            { return proto.CompactTextString(m) }
func (*PhzBaseDeskInfo) ProtoMessage()               {}
func (*PhzBaseDeskInfo) Descriptor() ([]byte, []int) { return fileDescriptor48, []int{3} }

func (m *PhzBaseDeskInfo) GetRoomInfo() *PhzBaseRoomInfo {
	if m != nil {
		return m.RoomInfo
	}
	return nil
}

func (m *PhzBaseDeskInfo) GetGameStatus() int32 {
	if m != nil && m.GameStatus != nil {
		return *m.GameStatus
	}
	return 0
}

func (m *PhzBaseDeskInfo) GetRemainPaiCount() int32 {
	if m != nil && m.RemainPaiCount != nil {
		return *m.RemainPaiCount
	}
	return 0
}

func (m *PhzBaseDeskInfo) GetLiangZhang() int32 {
	if m != nil && m.LiangZhang != nil {
		return *m.LiangZhang
	}
	return 0
}

func (m *PhzBaseDeskInfo) GetActiveUser() uint32 {
	if m != nil && m.ActiveUser != nil {
		return *m.ActiveUser
	}
	return 0
}

func (m *PhzBaseDeskInfo) GetIsOutCard() bool {
	if m != nil && m.IsOutCard != nil {
		return *m.IsOutCard
	}
	return false
}

type PhzBasePlayerInfo struct {
	UserId           *uint32            `protobuf:"varint,1,opt,name=userId" json:"userId,omitempty"`
	IsOwner          *bool              `protobuf:"varint,2,opt,name=isOwner" json:"isOwner,omitempty"`
	IsBanker         *bool              `protobuf:"varint,3,opt,name=isBanker" json:"isBanker,omitempty"`
	IsReady          *bool              `protobuf:"varint,4,opt,name=isReady" json:"isReady,omitempty"`
	IsLeave          *bool              `protobuf:"varint,5,opt,name=isLeave" json:"isLeave,omitempty"`
	HuXi             *int32             `protobuf:"varint,6,opt,name=huXi" json:"huXi,omitempty"`
	Score            *int64             `protobuf:"varint,7,opt,name=score" json:"score,omitempty"`
	WxInfo           *WeixinInfo        `protobuf:"bytes,8,opt,name=wxInfo" json:"wxInfo,omitempty"`
	HandPokers       []int32            `protobuf:"varint,9,rep,name=handPokers" json:"handPokers,omitempty"`
	TiPais           []*PhzPlayPaiIds   `protobuf:"bytes,10,rep,name=tiPais" json:"tiPais,omitempty"`
	PaoPais          []*PhzPlayPaiIds   `protobuf:"bytes,11,rep,name=paoPais" json:"paoPais,omitempty"`
	WeiPais          []*PhzPlayPaiIds   `protobuf:"bytes,12,rep,name=weiPais" json:"weiPais,omitempty"`
	PengPais         []*PhzPlayPaiIds   `protobuf:"bytes,13,rep,name=pengPais" json:"pengPais,omitempty"`
	ChiPais          []*PhzPlayPaiIds   `protobuf:"bytes,14,rep,name=chiPais" json:"chiPais,omitempty"`
	OutCards         []int32            `protobuf:"varint,15,rep,name=outCards" json:"outCards,omitempty"`
	PlayerStatus     *PhzEnumUserStatus `protobuf:"varint,16,opt,name=playerStatus,enum=ddproto.PhzEnumUserStatus" json:"playerStatus,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *PhzBasePlayerInfo) Reset()                    { *m = PhzBasePlayerInfo{} }
func (m *PhzBasePlayerInfo) String() string            { return proto.CompactTextString(m) }
func (*PhzBasePlayerInfo) ProtoMessage()               {}
func (*PhzBasePlayerInfo) Descriptor() ([]byte, []int) { return fileDescriptor48, []int{4} }

func (m *PhzBasePlayerInfo) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *PhzBasePlayerInfo) GetIsOwner() bool {
	if m != nil && m.IsOwner != nil {
		return *m.IsOwner
	}
	return false
}

func (m *PhzBasePlayerInfo) GetIsBanker() bool {
	if m != nil && m.IsBanker != nil {
		return *m.IsBanker
	}
	return false
}

func (m *PhzBasePlayerInfo) GetIsReady() bool {
	if m != nil && m.IsReady != nil {
		return *m.IsReady
	}
	return false
}

func (m *PhzBasePlayerInfo) GetIsLeave() bool {
	if m != nil && m.IsLeave != nil {
		return *m.IsLeave
	}
	return false
}

func (m *PhzBasePlayerInfo) GetHuXi() int32 {
	if m != nil && m.HuXi != nil {
		return *m.HuXi
	}
	return 0
}

func (m *PhzBasePlayerInfo) GetScore() int64 {
	if m != nil && m.Score != nil {
		return *m.Score
	}
	return 0
}

func (m *PhzBasePlayerInfo) GetWxInfo() *WeixinInfo {
	if m != nil {
		return m.WxInfo
	}
	return nil
}

func (m *PhzBasePlayerInfo) GetHandPokers() []int32 {
	if m != nil {
		return m.HandPokers
	}
	return nil
}

func (m *PhzBasePlayerInfo) GetTiPais() []*PhzPlayPaiIds {
	if m != nil {
		return m.TiPais
	}
	return nil
}

func (m *PhzBasePlayerInfo) GetPaoPais() []*PhzPlayPaiIds {
	if m != nil {
		return m.PaoPais
	}
	return nil
}

func (m *PhzBasePlayerInfo) GetWeiPais() []*PhzPlayPaiIds {
	if m != nil {
		return m.WeiPais
	}
	return nil
}

func (m *PhzBasePlayerInfo) GetPengPais() []*PhzPlayPaiIds {
	if m != nil {
		return m.PengPais
	}
	return nil
}

func (m *PhzBasePlayerInfo) GetChiPais() []*PhzPlayPaiIds {
	if m != nil {
		return m.ChiPais
	}
	return nil
}

func (m *PhzBasePlayerInfo) GetOutCards() []int32 {
	if m != nil {
		return m.OutCards
	}
	return nil
}

func (m *PhzBasePlayerInfo) GetPlayerStatus() PhzEnumUserStatus {
	if m != nil && m.PlayerStatus != nil {
		return *m.PlayerStatus
	}
	return PhzEnumUserStatus_PHZUSER_STATUS_INIT
}

func init() {
	proto.RegisterType((*PhzPlayPaiIds)(nil), "ddproto.phz_play_paiIds")
	proto.RegisterType((*PhzBaseCreateOption)(nil), "ddproto.phz_base_createOption")
	proto.RegisterType((*PhzBaseRoomInfo)(nil), "ddproto.phz_base_roomInfo")
	proto.RegisterType((*PhzBaseDeskInfo)(nil), "ddproto.phz_base_deskInfo")
	proto.RegisterType((*PhzBasePlayerInfo)(nil), "ddproto.phz_base_playerInfo")
	proto.RegisterEnum("ddproto.PhzEnumProtoId", PhzEnumProtoId_name, PhzEnumProtoId_value)
	proto.RegisterEnum("ddproto.PhzEnumRoomType", PhzEnumRoomType_name, PhzEnumRoomType_value)
	proto.RegisterEnum("ddproto.PhzEnumTiType", PhzEnumTiType_name, PhzEnumTiType_value)
	proto.RegisterEnum("ddproto.PhzEnumPengType", PhzEnumPengType_name, PhzEnumPengType_value)
	proto.RegisterEnum("ddproto.PhzEnumPaoType", PhzEnumPaoType_name, PhzEnumPaoType_value)
	proto.RegisterEnum("ddproto.PhzEnumDeskStatus", PhzEnumDeskStatus_name, PhzEnumDeskStatus_value)
	proto.RegisterEnum("ddproto.PhzEnumUserStatus", PhzEnumUserStatus_name, PhzEnumUserStatus_value)
}

func init() { proto.RegisterFile("phz_base.proto", fileDescriptor48) }

var fileDescriptor48 = []byte{
	// 1335 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x56, 0xdb, 0x72, 0xdb, 0x44,
	0x18, 0xae, 0x8f, 0x71, 0x36, 0xa7, 0xcd, 0xe6, 0x50, 0x35, 0x4d, 0x4b, 0x6a, 0xa0, 0x18, 0xb7,
	0x74, 0x98, 0x0e, 0xc3, 0x35, 0xb2, 0xbc, 0x89, 0x35, 0x76, 0x24, 0xb1, 0x92, 0x1b, 0xd2, 0x1b,
	0xa1, 0xda, 0x22, 0xd1, 0xb4, 0x91, 0x3c, 0x92, 0xdd, 0xa4, 0xbc, 0x5b, 0xdf, 0x80, 0x47, 0xe1,
	0x02, 0xae, 0xb8, 0x65, 0xf6, 0x97, 0x56, 0x59, 0xd9, 0x0e, 0x19, 0x6e, 0x3c, 0xde, 0xef, 0xfb,
	0xf6, 0x3f, 0xef, 0xae, 0xd0, 0xe6, 0xe4, 0xf2, 0x77, 0xf7, 0x9d, 0x97, 0xf8, 0xaf, 0x26, 0x71,
	0x34, 0x8d, 0xc8, 0xca, 0x78, 0x0c, 0x7f, 0x0e, 0x76, 0x46, 0xd1, 0xd5, 0x55, 0x14, 0xba, 0xa3,
	0x0f, 0x81, 0x1f, 0x4e, 0x53, 0xb6, 0xf9, 0x02, 0x6d, 0x71, 0xfd, 0xe4, 0x83, 0xf7, 0xc9, 0x9d,
	0x78, 0x81, 0x3e, 0x4e, 0x88, 0x82, 0x56, 0x26, 0xd1, 0x7b, 0x3f, 0xd6, 0xc7, 0x4a, 0xe9, 0xa8,
	0xd2, 0xaa, 0x31, 0xb1, 0x6c, 0x7e, 0x2e, 0xa1, 0x3d, 0x61, 0xdd, 0x1d, 0xc5, 0xbe, 0x37, 0xf5,
	0xcd, 0xc9, 0x34, 0x88, 0x42, 0x72, 0x88, 0x56, 0x67, 0x89, 0x1f, 0x6b, 0xd1, 0x2c, 0x9c, 0x2a,
	0xa5, 0xa3, 0x52, 0xab, 0xc6, 0x6e, 0x01, 0xf2, 0x23, 0x6a, 0xc4, 0x51, 0x74, 0xe5, 0x7c, 0x9a,
	0xf8, 0x4a, 0xf9, 0xa8, 0xd4, 0xda, 0x7c, 0x7d, 0xf0, 0x2a, 0x8b, 0xea, 0x15, 0xb7, 0xe7, 0x87,
	0xb3, 0x2b, 0x57, 0x28, 0x58, 0xae, 0x25, 0x4f, 0x11, 0x7a, 0x17, 0x79, 0xf1, 0x38, 0xd1, 0xa2,
	0xd9, 0x54, 0xa9, 0x80, 0x59, 0x09, 0x21, 0x04, 0x55, 0x2f, 0x67, 0xbf, 0x04, 0x4a, 0x15, 0x18,
	0xf8, 0x4f, 0x0e, 0x50, 0x23, 0x48, 0xba, 0x5e, 0xd0, 0xf7, 0x02, 0xa5, 0x76, 0x54, 0x6a, 0x35,
	0x58, 0xbe, 0x6e, 0xfe, 0x53, 0x42, 0xdb, 0x79, 0xfc, 0xdc, 0x8b, 0x1e, 0xfe, 0x16, 0xf1, 0x7c,
	0xa3, 0xeb, 0x30, 0xcb, 0xb7, 0xd4, 0xda, 0x60, 0x62, 0xc9, 0x6d, 0x4d, 0xbc, 0x24, 0xb9, 0x8e,
	0xe2, 0x31, 0xc4, 0xbd, 0xca, 0xf2, 0x75, 0x31, 0xe3, 0xca, 0x7c, 0xc6, 0xc5, 0xc8, 0xab, 0x0b,
	0x91, 0x37, 0xd1, 0xfa, 0x68, 0x16, 0xc7, 0x7e, 0x38, 0x65, 0xd1, 0x2c, 0x1c, 0x43, 0xa4, 0x35,
	0x56, 0xc0, 0x0a, 0x55, 0xab, 0xff, 0x8f, 0xaa, 0x41, 0x55, 0x6e, 0x02, 0x65, 0x45, 0x54, 0xe5,
	0x26, 0x68, 0xfe, 0x2d, 0x67, 0x3e, 0xf6, 0x93, 0xf7, 0x90, 0x79, 0xe6, 0x81, 0xff, 0x87, 0xd4,
	0xd7, 0xe6, 0x3c, 0x14, 0xea, 0xc4, 0x72, 0x2d, 0xcf, 0xee, 0xc2, 0xbb, 0xf2, 0xed, 0xa9, 0x37,
	0x9d, 0x25, 0x50, 0x99, 0x1a, 0x93, 0x10, 0xf2, 0x1c, 0x6d, 0xc6, 0xfe, 0x95, 0x17, 0x84, 0x96,
	0x17, 0xc8, 0x05, 0x9a, 0x43, 0xb9, 0x9d, 0x0f, 0x81, 0x17, 0x5e, 0xbc, 0xbd, 0xf4, 0xc2, 0x0b,
	0x51, 0xa5, 0x5b, 0x84, 0xf3, 0xde, 0x68, 0x1a, 0x7c, 0xf4, 0x87, 0x89, 0x1f, 0x43, 0x8d, 0x36,
	0x98, 0x84, 0xf0, 0x1e, 0x04, 0x89, 0x39, 0x9b, 0x6a, 0x5e, 0x3c, 0x86, 0x12, 0x35, 0xd8, 0x2d,
	0xd0, 0xfc, 0xab, 0x8a, 0x76, 0xf2, 0x2c, 0xf8, 0x80, 0xfb, 0x31, 0x44, 0xbf, 0x8f, 0xea, 0xbc,
	0x51, 0x79, 0xbb, 0xb3, 0x15, 0x9f, 0x83, 0x20, 0x31, 0x79, 0xeb, 0x21, 0xa5, 0x06, 0x13, 0xcb,
	0x74, 0xa6, 0x3a, 0x5e, 0xf8, 0xde, 0x8f, 0x21, 0x13, 0x98, 0xa9, 0x74, 0x9d, 0xee, 0x62, 0xbe,
	0x37, 0xfe, 0x04, 0x09, 0xc0, 0x2e, 0x58, 0xa6, 0xcc, 0xc0, 0xf7, 0x3e, 0xfa, 0xd9, 0x20, 0x8a,
	0x65, 0x3e, 0xb7, 0x75, 0x69, 0x6e, 0x77, 0x51, 0x2d, 0x19, 0x45, 0xb1, 0x0f, 0x6d, 0xab, 0xb0,
	0x74, 0x41, 0x5e, 0xa0, 0xfa, 0xf5, 0x0d, 0xf4, 0xa7, 0x01, 0xfd, 0xd9, 0xc9, 0xfb, 0x73, 0xe6,
	0x07, 0x37, 0x41, 0x08, 0x8d, 0xc9, 0x24, 0xbc, 0x5c, 0x97, 0x5e, 0x38, 0xb6, 0xf8, 0x69, 0x4d,
	0x94, 0x55, 0x38, 0xbb, 0x12, 0x42, 0xbe, 0x47, 0xf5, 0x69, 0x60, 0x79, 0x41, 0xa2, 0xa0, 0xa3,
	0x4a, 0x6b, 0xed, 0xb5, 0x52, 0x68, 0xb6, 0x74, 0x05, 0xb0, 0x4c, 0x47, 0x5e, 0xa3, 0x95, 0x89,
	0x17, 0xc1, 0x96, 0xb5, 0x7b, 0xb6, 0x08, 0x21, 0xdf, 0x73, 0xed, 0xa7, 0x6e, 0xd6, 0xef, 0xdb,
	0x93, 0x09, 0xc9, 0x0f, 0xa8, 0x31, 0xf1, 0xc3, 0x0b, 0xd8, 0xb4, 0x71, 0xcf, 0xa6, 0x5c, 0xc9,
	0x3d, 0x8d, 0x2e, 0x53, 0x4f, 0x9b, 0xf7, 0x79, 0xca, 0x84, 0xbc, 0x95, 0x51, 0x3a, 0x1f, 0x89,
	0xb2, 0x05, 0x15, 0xca, 0xd7, 0xe4, 0x27, 0xb4, 0x9e, 0x8e, 0x49, 0x36, 0xd8, 0x18, 0x0e, 0xdd,
	0xe1, 0xe2, 0xa1, 0xe3, 0x03, 0x93, 0x6a, 0x58, 0x61, 0x47, 0xfb, 0xcf, 0x06, 0xc2, 0xb9, 0x0a,
	0x36, 0xe9, 0x63, 0xb2, 0x87, 0xb6, 0xad, 0xde, 0x5b, 0xd7, 0xd2, 0xbb, 0x6e, 0x8f, 0xaa, 0xcc,
	0xe9, 0x50, 0xd5, 0xc1, 0x0f, 0xc8, 0x01, 0xda, 0x17, 0xb0, 0xc6, 0xa8, 0xea, 0xd0, 0x2e, 0xb5,
	0xfb, 0x2e, 0xa3, 0x3f, 0xe3, 0x12, 0x79, 0x84, 0xf6, 0x04, 0x47, 0x0d, 0x87, 0xb2, 0x9c, 0x2a,
	0x93, 0x5d, 0x84, 0x05, 0x05, 0xa8, 0xaa, 0xf5, 0x71, 0x85, 0x28, 0x68, 0x57, 0xa0, 0x27, 0xea,
	0x29, 0xd5, 0x8d, 0x63, 0x13, 0xf4, 0x55, 0x59, 0x2f, 0x18, 0x5c, 0x93, 0x63, 0x62, 0x54, 0xed,
	0x9e, 0x83, 0xb8, 0xbe, 0x08, 0x73, 0xeb, 0x2b, 0x64, 0x07, 0x6d, 0x09, 0xd8, 0xb4, 0xa8, 0xa1,
	0x1b, 0x27, 0xb8, 0x21, 0x6b, 0x6d, 0x6a, 0x74, 0x35, 0x95, 0x75, 0x6d, 0xbc, 0x4a, 0x1e, 0xa2,
	0x9d, 0x5c, 0x3b, 0x74, 0x38, 0x0a, 0xb6, 0xd1, 0x32, 0x82, 0x5b, 0x5f, 0x23, 0xdb, 0x68, 0x43,
	0x10, 0xa7, 0xa6, 0xa5, 0xea, 0x78, 0x5d, 0x0e, 0xda, 0x7c, 0x43, 0x99, 0x33, 0x64, 0x06, 0xde,
	0x90, 0xc3, 0xd0, 0x54, 0xc3, 0xa2, 0xc6, 0x09, 0xde, 0x94, 0xa5, 0x1c, 0x01, 0x67, 0x5b, 0x0b,
	0x28, 0xf7, 0x84, 0x0b, 0x06, 0x7a, 0x3a, 0x48, 0xb7, 0x09, 0x41, 0x9b, 0x92, 0x55, 0xad, 0xa7,
	0x63, 0x32, 0x2f, 0xe4, 0xbb, 0x77, 0xe4, 0x84, 0x3b, 0xba, 0xa5, 0xa6, 0xfb, 0x77, 0x17, 0x61,
	0xae, 0xde, 0x93, 0xb3, 0x72, 0x38, 0x8c, 0xf7, 0x65, 0x4f, 0x67, 0x14, 0xb0, 0x87, 0x32, 0x66,
	0xa9, 0x90, 0xbd, 0x52, 0x18, 0x98, 0xa1, 0x70, 0xf4, 0x48, 0xce, 0x49, 0x53, 0x0d, 0x60, 0xf0,
	0xc1, 0xa2, 0x98, 0xbb, 0x7f, 0x5c, 0x28, 0x80, 0x6a, 0xdb, 0x60, 0xe2, 0x70, 0x01, 0xe5, 0xda,
	0x27, 0xf2, 0xb4, 0x69, 0x43, 0xc6, 0xa8, 0xe1, 0x30, 0x6a, 0x0f, 0x07, 0x0e, 0x7e, 0x2a, 0x5b,
	0xa7, 0x46, 0x37, 0x83, 0xbf, 0x20, 0x4f, 0xd0, 0x23, 0x01, 0xab, 0x96, 0x35, 0x38, 0xef, 0xea,
	0xb6, 0x6d, 0x0e, 0xde, 0x50, 0x70, 0x73, 0x74, 0x37, 0xcd, 0xfd, 0x3d, 0x23, 0xcf, 0xd0, 0x93,
	0xa5, 0x74, 0x47, 0xd5, 0xd2, 0x29, 0x6f, 0xfe, 0xb7, 0x84, 0x5b, 0xf9, 0x92, 0x1c, 0x22, 0x25,
	0x3f, 0x08, 0x19, 0x9b, 0x1f, 0x93, 0xaf, 0xee, 0x64, 0xf9, 0xde, 0xaf, 0xe5, 0xb4, 0x86, 0x36,
	0x65, 0x1d, 0x46, 0xd5, 0x3e, 0x7e, 0x2e, 0x9f, 0x22, 0x80, 0x1c, 0xfd, 0x94, 0x9a, 0x43, 0x07,
	0x7f, 0x23, 0x97, 0x68, 0x40, 0x55, 0xc9, 0x53, 0x6b, 0x39, 0xc5, 0xdd, 0x7c, 0x2b, 0x97, 0xbb,
	0xaf, 0x6b, 0x7d, 0xee, 0x0a, 0xb7, 0xe5, 0x83, 0x20, 0x50, 0xb7, 0xa3, 0xe1, 0x17, 0x32, 0x71,
	0x4a, 0x6d, 0x5b, 0x3d, 0x49, 0xeb, 0xf9, 0x92, 0xec, 0x23, 0x32, 0x4f, 0x74, 0x34, 0xfc, 0x5d,
	0xfb, 0x65, 0xfa, 0xa8, 0x17, 0x3e, 0x04, 0x84, 0x15, 0x66, 0x9a, 0xa7, 0xce, 0xb9, 0x45, 0xdd,
	0x63, 0xa6, 0x53, 0xa3, 0x8b, 0x4b, 0x6d, 0x37, 0xfd, 0xd4, 0x03, 0xf5, 0x34, 0x00, 0x6d, 0x16,
	0xbb, 0xa3, 0x83, 0xb2, 0xa7, 0xbe, 0xa1, 0xae, 0xd3, 0x63, 0x94, 0xe2, 0x92, 0xa8, 0x85, 0x4c,
	0x1d, 0x9b, 0x43, 0x86, 0xcb, 0xc2, 0x81, 0xcc, 0x9c, 0x51, 0x1d, 0x57, 0xda, 0xbf, 0x4a, 0xe1,
	0xf0, 0x4b, 0x1a, 0x5c, 0x88, 0x1a, 0x50, 0xe3, 0x04, 0xf4, 0x5c, 0x9a, 0x5b, 0xcf, 0x51, 0xad,
	0x67, 0x0e, 0x39, 0x53, 0x26, 0x8f, 0xd1, 0xc3, 0x02, 0x63, 0x98, 0xec, 0x54, 0x1d, 0xc0, 0x61,
	0xe7, 0x1e, 0xa4, 0xeb, 0xd5, 0x8b, 0xe4, 0x1c, 0x2c, 0xd5, 0x04, 0x3d, 0xc4, 0xef, 0x9a, 0x06,
	0xcf, 0x41, 0x14, 0x34, 0xa3, 0xce, 0xa8, 0x0e, 0x44, 0x39, 0x77, 0x9f, 0x11, 0x70, 0x45, 0x70,
	0xa6, 0xd2, 0xfe, 0xa3, 0x94, 0x7e, 0x34, 0x80, 0x0b, 0xfe, 0xa1, 0x94, 0x7d, 0xd2, 0xa4, 0x61,
	0x41, 0x6f, 0xf9, 0x8f, 0xed, 0xa8, 0xce, 0xd0, 0x76, 0x75, 0x43, 0x77, 0x70, 0x29, 0x1b, 0xb6,
	0x79, 0x12, 0x6e, 0x50, 0x5c, 0x26, 0x4f, 0xd1, 0xc1, 0x12, 0x56, 0x5c, 0xa4, 0x95, 0x3b, 0x4c,
	0x5b, 0x03, 0xf5, 0x1c, 0x57, 0xef, 0xd8, 0x3c, 0x30, 0x1d, 0x87, 0xb2, 0x73, 0x5c, 0xcb, 0x8e,
	0xda, 0x3c, 0x7f, 0xac, 0x1b, 0xba, 0xdd, 0xc3, 0xf5, 0xf6, 0xe7, 0xb2, 0x94, 0xce, 0xed, 0xb3,
	0x95, 0x55, 0x06, 0x46, 0xaf, 0x98, 0x4a, 0x5a, 0x19, 0x99, 0x10, 0x69, 0x2c, 0x32, 0x1d, 0x7a,
	0xa2, 0x1b, 0xf9, 0xe3, 0x23, 0x33, 0xe9, 0x3d, 0x5e, 0xcd, 0x7a, 0x23, 0x33, 0xd9, 0x15, 0x5b,
	0xcb, 0x9e, 0xbf, 0x39, 0x0a, 0xda, 0x5c, 0x5f, 0x62, 0x10, 0x2e, 0x3c, 0xbc, 0x92, 0xd5, 0xaa,
	0xc0, 0xf4, 0xa8, 0xd6, 0x4f, 0x5f, 0xa4, 0x45, 0x93, 0xd9, 0x43, 0x83, 0x57, 0x97, 0x70, 0xa2,
	0x86, 0x68, 0x49, 0x94, 0x59, 0xfd, 0xd6, 0xac, 0x07, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0x7b,
	0x12, 0x2f, 0x3e, 0x4f, 0x0d, 0x00, 0x00,
}
