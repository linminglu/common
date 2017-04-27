// Code generated by protoc-gen-go.
// source: niuniu_base.proto
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

type NiuniuEnumProtoid int32

const (
	// //////////////////////////////////
	NiuniuEnumProtoid_NIU_PID_HEARTBEAT          NiuniuEnumProtoid = 0
	NiuniuEnumProtoid_NIU_PID_QUICK_CONN         NiuniuEnumProtoid = 1
	NiuniuEnumProtoid_NIU_PID_QUICK_CONN_ACK     NiuniuEnumProtoid = 2
	NiuniuEnumProtoid_NIU_PID_GAME_LOGIN         NiuniuEnumProtoid = 3
	NiuniuEnumProtoid_NIU_PID_GAME_LOGIN_ACK     NiuniuEnumProtoid = 4
	NiuniuEnumProtoid_NIU_PID_CREATE_DESK_REQ    NiuniuEnumProtoid = 5
	NiuniuEnumProtoid_NIU_PID_ENTER_DESK_REQ     NiuniuEnumProtoid = 6
	NiuniuEnumProtoid_NIU_PID_ENTER_DESK_ACK     NiuniuEnumProtoid = 7
	NiuniuEnumProtoid_NIU_PID_ENTER_DESK_BC      NiuniuEnumProtoid = 8
	NiuniuEnumProtoid_NIU_PID_READY_REQ          NiuniuEnumProtoid = 9
	NiuniuEnumProtoid_NIU_PID_READY_ACK          NiuniuEnumProtoid = 10
	NiuniuEnumProtoid_NIU_PID_READY_BC           NiuniuEnumProtoid = 11
	NiuniuEnumProtoid_NIU_PID_START_GAME_OT      NiuniuEnumProtoid = 12
	NiuniuEnumProtoid_NIU_PID_QIANGZHUANG_OT     NiuniuEnumProtoid = 13
	NiuniuEnumProtoid_NIU_PID_QIANGZHUANG_REQ    NiuniuEnumProtoid = 14
	NiuniuEnumProtoid_NIU_PID_QIANGZHUANG_ACK    NiuniuEnumProtoid = 15
	NiuniuEnumProtoid_NIU_PID_QIANGZHUANG_BC     NiuniuEnumProtoid = 16
	NiuniuEnumProtoid_NIU_PID_JIABEI_OT          NiuniuEnumProtoid = 17
	NiuniuEnumProtoid_NIU_PID_JIABEI_REQ         NiuniuEnumProtoid = 18
	NiuniuEnumProtoid_NIU_PID_JIABEI_ACK         NiuniuEnumProtoid = 19
	NiuniuEnumProtoid_NIU_PID_JIABEI_BC          NiuniuEnumProtoid = 20
	NiuniuEnumProtoid_NIU_PID_BIPAI_RESULT_BC    NiuniuEnumProtoid = 21
	NiuniuEnumProtoid_NIU_PID_GAME_END_BC        NiuniuEnumProtoid = 22
	NiuniuEnumProtoid_NIU_PID_APPLY_DISSOLVE_REQ NiuniuEnumProtoid = 23
	NiuniuEnumProtoid_NIU_PID_APPLY_DISSOLVE_ACK NiuniuEnumProtoid = 24
	NiuniuEnumProtoid_NIU_PID_DISSOLVE_BACK_REQ  NiuniuEnumProtoid = 25
	NiuniuEnumProtoid_NIU_PID_DISSOLVE_BACK_ACK  NiuniuEnumProtoid = 26
	NiuniuEnumProtoid_NIU_PID_DISSOLVE_DONE_BC   NiuniuEnumProtoid = 27
	NiuniuEnumProtoid_NIU_PID_SEND_MESSAGE_REQ   NiuniuEnumProtoid = 31
	NiuniuEnumProtoid_NIU_PID_MESSAGE_BC         NiuniuEnumProtoid = 32
	NiuniuEnumProtoid_NIU_PID_LEAVE_DESK_REQ     NiuniuEnumProtoid = 33
	NiuniuEnumProtoid_NIU_PID_LEAVE_DESK_ACK     NiuniuEnumProtoid = 34
	NiuniuEnumProtoid_NIU_PID_OWNER_DISSOLVE_REQ NiuniuEnumProtoid = 35
	NiuniuEnumProtoid_NIU_PID_OWNER_DISSOLVE_ACK NiuniuEnumProtoid = 36
)

var NiuniuEnumProtoid_name = map[int32]string{
	0:  "NIU_PID_HEARTBEAT",
	1:  "NIU_PID_QUICK_CONN",
	2:  "NIU_PID_QUICK_CONN_ACK",
	3:  "NIU_PID_GAME_LOGIN",
	4:  "NIU_PID_GAME_LOGIN_ACK",
	5:  "NIU_PID_CREATE_DESK_REQ",
	6:  "NIU_PID_ENTER_DESK_REQ",
	7:  "NIU_PID_ENTER_DESK_ACK",
	8:  "NIU_PID_ENTER_DESK_BC",
	9:  "NIU_PID_READY_REQ",
	10: "NIU_PID_READY_ACK",
	11: "NIU_PID_READY_BC",
	12: "NIU_PID_START_GAME_OT",
	13: "NIU_PID_QIANGZHUANG_OT",
	14: "NIU_PID_QIANGZHUANG_REQ",
	15: "NIU_PID_QIANGZHUANG_ACK",
	16: "NIU_PID_QIANGZHUANG_BC",
	17: "NIU_PID_JIABEI_OT",
	18: "NIU_PID_JIABEI_REQ",
	19: "NIU_PID_JIABEI_ACK",
	20: "NIU_PID_JIABEI_BC",
	21: "NIU_PID_BIPAI_RESULT_BC",
	22: "NIU_PID_GAME_END_BC",
	23: "NIU_PID_APPLY_DISSOLVE_REQ",
	24: "NIU_PID_APPLY_DISSOLVE_ACK",
	25: "NIU_PID_DISSOLVE_BACK_REQ",
	26: "NIU_PID_DISSOLVE_BACK_ACK",
	27: "NIU_PID_DISSOLVE_DONE_BC",
	31: "NIU_PID_SEND_MESSAGE_REQ",
	32: "NIU_PID_MESSAGE_BC",
	33: "NIU_PID_LEAVE_DESK_REQ",
	34: "NIU_PID_LEAVE_DESK_ACK",
	35: "NIU_PID_OWNER_DISSOLVE_REQ",
	36: "NIU_PID_OWNER_DISSOLVE_ACK",
}
var NiuniuEnumProtoid_value = map[string]int32{
	"NIU_PID_HEARTBEAT":          0,
	"NIU_PID_QUICK_CONN":         1,
	"NIU_PID_QUICK_CONN_ACK":     2,
	"NIU_PID_GAME_LOGIN":         3,
	"NIU_PID_GAME_LOGIN_ACK":     4,
	"NIU_PID_CREATE_DESK_REQ":    5,
	"NIU_PID_ENTER_DESK_REQ":     6,
	"NIU_PID_ENTER_DESK_ACK":     7,
	"NIU_PID_ENTER_DESK_BC":      8,
	"NIU_PID_READY_REQ":          9,
	"NIU_PID_READY_ACK":          10,
	"NIU_PID_READY_BC":           11,
	"NIU_PID_START_GAME_OT":      12,
	"NIU_PID_QIANGZHUANG_OT":     13,
	"NIU_PID_QIANGZHUANG_REQ":    14,
	"NIU_PID_QIANGZHUANG_ACK":    15,
	"NIU_PID_QIANGZHUANG_BC":     16,
	"NIU_PID_JIABEI_OT":          17,
	"NIU_PID_JIABEI_REQ":         18,
	"NIU_PID_JIABEI_ACK":         19,
	"NIU_PID_JIABEI_BC":          20,
	"NIU_PID_BIPAI_RESULT_BC":    21,
	"NIU_PID_GAME_END_BC":        22,
	"NIU_PID_APPLY_DISSOLVE_REQ": 23,
	"NIU_PID_APPLY_DISSOLVE_ACK": 24,
	"NIU_PID_DISSOLVE_BACK_REQ":  25,
	"NIU_PID_DISSOLVE_BACK_ACK":  26,
	"NIU_PID_DISSOLVE_DONE_BC":   27,
	"NIU_PID_SEND_MESSAGE_REQ":   31,
	"NIU_PID_MESSAGE_BC":         32,
	"NIU_PID_LEAVE_DESK_REQ":     33,
	"NIU_PID_LEAVE_DESK_ACK":     34,
	"NIU_PID_OWNER_DISSOLVE_REQ": 35,
	"NIU_PID_OWNER_DISSOLVE_ACK": 36,
}

func (x NiuniuEnumProtoid) Enum() *NiuniuEnumProtoid {
	p := new(NiuniuEnumProtoid)
	*p = x
	return p
}
func (x NiuniuEnumProtoid) String() string {
	return proto.EnumName(NiuniuEnumProtoid_name, int32(x))
}
func (x *NiuniuEnumProtoid) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(NiuniuEnumProtoid_value, data, "NiuniuEnumProtoid")
	if err != nil {
		return err
	}
	*x = NiuniuEnumProtoid(value)
	return nil
}
func (NiuniuEnumProtoid) EnumDescriptor() ([]byte, []int) { return fileDescriptor27, []int{0} }

// =================================公共================================
// 牛牛牌的类型
type NiuniuEnum_PokerType int32

const (
	NiuniuEnum_PokerType_NO_NIU           NiuniuEnum_PokerType = 1
	NiuniuEnum_PokerType_NIU_1            NiuniuEnum_PokerType = 2
	NiuniuEnum_PokerType_NIU_2            NiuniuEnum_PokerType = 3
	NiuniuEnum_PokerType_NIU_3            NiuniuEnum_PokerType = 4
	NiuniuEnum_PokerType_NIU_4            NiuniuEnum_PokerType = 5
	NiuniuEnum_PokerType_NIU_5            NiuniuEnum_PokerType = 6
	NiuniuEnum_PokerType_NIU_6            NiuniuEnum_PokerType = 7
	NiuniuEnum_PokerType_NIU_7            NiuniuEnum_PokerType = 8
	NiuniuEnum_PokerType_NIU_8            NiuniuEnum_PokerType = 9
	NiuniuEnum_PokerType_NIU_9            NiuniuEnum_PokerType = 10
	NiuniuEnum_PokerType_NIU_NIU          NiuniuEnum_PokerType = 11
	NiuniuEnum_PokerType_YIN_NIU          NiuniuEnum_PokerType = 12
	NiuniuEnum_PokerType_JIN_NIU          NiuniuEnum_PokerType = 13
	NiuniuEnum_PokerType_WU_XIAO_NIU      NiuniuEnum_PokerType = 14
	NiuniuEnum_PokerType_NIU_ZHA_DAN      NiuniuEnum_PokerType = 15
	NiuniuEnum_PokerType_NIU_YI_TIAO_LONG NiuniuEnum_PokerType = 16
)

var NiuniuEnum_PokerType_name = map[int32]string{
	1:  "NO_NIU",
	2:  "NIU_1",
	3:  "NIU_2",
	4:  "NIU_3",
	5:  "NIU_4",
	6:  "NIU_5",
	7:  "NIU_6",
	8:  "NIU_7",
	9:  "NIU_8",
	10: "NIU_9",
	11: "NIU_NIU",
	12: "YIN_NIU",
	13: "JIN_NIU",
	14: "WU_XIAO_NIU",
	15: "NIU_ZHA_DAN",
	16: "NIU_YI_TIAO_LONG",
}
var NiuniuEnum_PokerType_value = map[string]int32{
	"NO_NIU":           1,
	"NIU_1":            2,
	"NIU_2":            3,
	"NIU_3":            4,
	"NIU_4":            5,
	"NIU_5":            6,
	"NIU_6":            7,
	"NIU_7":            8,
	"NIU_8":            9,
	"NIU_9":            10,
	"NIU_NIU":          11,
	"YIN_NIU":          12,
	"JIN_NIU":          13,
	"WU_XIAO_NIU":      14,
	"NIU_ZHA_DAN":      15,
	"NIU_YI_TIAO_LONG": 16,
}

func (x NiuniuEnum_PokerType) Enum() *NiuniuEnum_PokerType {
	p := new(NiuniuEnum_PokerType)
	*p = x
	return p
}
func (x NiuniuEnum_PokerType) String() string {
	return proto.EnumName(NiuniuEnum_PokerType_name, int32(x))
}
func (x *NiuniuEnum_PokerType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(NiuniuEnum_PokerType_value, data, "NiuniuEnum_PokerType")
	if err != nil {
		return err
	}
	*x = NiuniuEnum_PokerType(value)
	return nil
}
func (NiuniuEnum_PokerType) EnumDescriptor() ([]byte, []int) { return fileDescriptor27, []int{1} }

// 房间状态
type NiuniuEnumDeskState int32

const (
	NiuniuEnumDeskState_NIU_DESK_STATUS_WAIT_ENTER       NiuniuEnumDeskState = 1
	NiuniuEnumDeskState_NIU_DESK_STATUS_WAIT_READY       NiuniuEnumDeskState = 2
	NiuniuEnumDeskState_NIU_DESK_STATUS_WAIT_START       NiuniuEnumDeskState = 3
	NiuniuEnumDeskState_NIU_DESK_STATUS_WAIT_QIANGZHUANG NiuniuEnumDeskState = 4
	NiuniuEnumDeskState_NIU_DESK_STATUS_WAIT_JIABEI      NiuniuEnumDeskState = 5
	NiuniuEnumDeskState_NIU_DESK_STATUS_WAIT_BIPAI       NiuniuEnumDeskState = 6
	NiuniuEnumDeskState_NIU_DESK_STATUS_WAIT_RESULT      NiuniuEnumDeskState = 7
)

var NiuniuEnumDeskState_name = map[int32]string{
	1: "NIU_DESK_STATUS_WAIT_ENTER",
	2: "NIU_DESK_STATUS_WAIT_READY",
	3: "NIU_DESK_STATUS_WAIT_START",
	4: "NIU_DESK_STATUS_WAIT_QIANGZHUANG",
	5: "NIU_DESK_STATUS_WAIT_JIABEI",
	6: "NIU_DESK_STATUS_WAIT_BIPAI",
	7: "NIU_DESK_STATUS_WAIT_RESULT",
}
var NiuniuEnumDeskState_value = map[string]int32{
	"NIU_DESK_STATUS_WAIT_ENTER":       1,
	"NIU_DESK_STATUS_WAIT_READY":       2,
	"NIU_DESK_STATUS_WAIT_START":       3,
	"NIU_DESK_STATUS_WAIT_QIANGZHUANG": 4,
	"NIU_DESK_STATUS_WAIT_JIABEI":      5,
	"NIU_DESK_STATUS_WAIT_BIPAI":       6,
	"NIU_DESK_STATUS_WAIT_RESULT":      7,
}

func (x NiuniuEnumDeskState) Enum() *NiuniuEnumDeskState {
	p := new(NiuniuEnumDeskState)
	*p = x
	return p
}
func (x NiuniuEnumDeskState) String() string {
	return proto.EnumName(NiuniuEnumDeskState_name, int32(x))
}
func (x *NiuniuEnumDeskState) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(NiuniuEnumDeskState_value, data, "NiuniuEnumDeskState")
	if err != nil {
		return err
	}
	*x = NiuniuEnumDeskState(value)
	return nil
}
func (NiuniuEnumDeskState) EnumDescriptor() ([]byte, []int) { return fileDescriptor27, []int{2} }

// 坐庄规则
type NiuniuEnumBankerRule int32

const (
	NiuniuEnumBankerRule_DING_ZHUANG       NiuniuEnumBankerRule = 1
	NiuniuEnumBankerRule_SUI_JI_ZUO_ZHUANG NiuniuEnumBankerRule = 2
	NiuniuEnumBankerRule_QIANG_ZHUANG      NiuniuEnumBankerRule = 3
)

var NiuniuEnumBankerRule_name = map[int32]string{
	1: "DING_ZHUANG",
	2: "SUI_JI_ZUO_ZHUANG",
	3: "QIANG_ZHUANG",
}
var NiuniuEnumBankerRule_value = map[string]int32{
	"DING_ZHUANG":       1,
	"SUI_JI_ZUO_ZHUANG": 2,
	"QIANG_ZHUANG":      3,
}

func (x NiuniuEnumBankerRule) Enum() *NiuniuEnumBankerRule {
	p := new(NiuniuEnumBankerRule)
	*p = x
	return p
}
func (x NiuniuEnumBankerRule) String() string {
	return proto.EnumName(NiuniuEnumBankerRule_name, int32(x))
}
func (x *NiuniuEnumBankerRule) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(NiuniuEnumBankerRule_value, data, "NiuniuEnumBankerRule")
	if err != nil {
		return err
	}
	*x = NiuniuEnumBankerRule(value)
	return nil
}
func (NiuniuEnumBankerRule) EnumDescriptor() ([]byte, []int) { return fileDescriptor27, []int{3} }

// 打出去的牌
type NiuniuClientPoker struct {
	Pais             []*ClientBasePoker    `protobuf:"bytes,2,rep,name=pais" json:"pais,omitempty"`
	Type             *NiuniuEnum_PokerType `protobuf:"varint,3,opt,name=type,enum=ddproto.NiuniuEnum_PokerType" json:"type,omitempty"`
	XXX_unrecognized []byte                `json:"-"`
}

func (m *NiuniuClientPoker) Reset()                    { *m = NiuniuClientPoker{} }
func (m *NiuniuClientPoker) String() string            { return proto.CompactTextString(m) }
func (*NiuniuClientPoker) ProtoMessage()               {}
func (*NiuniuClientPoker) Descriptor() ([]byte, []int) { return fileDescriptor27, []int{0} }

func (m *NiuniuClientPoker) GetPais() []*ClientBasePoker {
	if m != nil {
		return m.Pais
	}
	return nil
}

func (m *NiuniuClientPoker) GetType() NiuniuEnum_PokerType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return NiuniuEnum_PokerType_NO_NIU
}

// 对局账单信息
type NiuniuUserBill struct {
	Score            *int32  `protobuf:"varint,1,opt,name=score" json:"score,omitempty"`
	CountHasNiu      *int32  `protobuf:"varint,2,opt,name=count_has_niu,json=countHasNiu" json:"count_has_niu,omitempty"`
	CountNoNiu       *int32  `protobuf:"varint,3,opt,name=count_no_niu,json=countNoNiu" json:"count_no_niu,omitempty"`
	CountWin         *int32  `protobuf:"varint,4,opt,name=count_win,json=countWin" json:"count_win,omitempty"`
	CountLost        *int32  `protobuf:"varint,5,opt,name=count_lost,json=countLost" json:"count_lost,omitempty"`
	UserId           *uint32 `protobuf:"varint,6,opt,name=userId" json:"userId,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *NiuniuUserBill) Reset()                    { *m = NiuniuUserBill{} }
func (m *NiuniuUserBill) String() string            { return proto.CompactTextString(m) }
func (*NiuniuUserBill) ProtoMessage()               {}
func (*NiuniuUserBill) Descriptor() ([]byte, []int) { return fileDescriptor27, []int{1} }

func (m *NiuniuUserBill) GetScore() int32 {
	if m != nil && m.Score != nil {
		return *m.Score
	}
	return 0
}

func (m *NiuniuUserBill) GetCountHasNiu() int32 {
	if m != nil && m.CountHasNiu != nil {
		return *m.CountHasNiu
	}
	return 0
}

func (m *NiuniuUserBill) GetCountNoNiu() int32 {
	if m != nil && m.CountNoNiu != nil {
		return *m.CountNoNiu
	}
	return 0
}

func (m *NiuniuUserBill) GetCountWin() int32 {
	if m != nil && m.CountWin != nil {
		return *m.CountWin
	}
	return 0
}

func (m *NiuniuUserBill) GetCountLost() int32 {
	if m != nil && m.CountLost != nil {
		return *m.CountLost
	}
	return 0
}

func (m *NiuniuUserBill) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

// desk 配置选项
type NiuniuDeskOption struct {
	MinUser          *int32                `protobuf:"varint,1,opt,name=minUser" json:"minUser,omitempty"`
	MaxUser          *int32                `protobuf:"varint,2,opt,name=maxUser" json:"maxUser,omitempty"`
	MaxCircle        *int32                `protobuf:"varint,3,opt,name=maxCircle" json:"maxCircle,omitempty"`
	HasFlower        *bool                 `protobuf:"varint,4,opt,name=hasFlower" json:"hasFlower,omitempty"`
	BankRule         *NiuniuEnumBankerRule `protobuf:"varint,5,opt,name=bankRule,enum=ddproto.NiuniuEnumBankerRule" json:"bankRule,omitempty"`
	IsFlowerPlay     *bool                 `protobuf:"varint,6,opt,name=isFlowerPlay" json:"isFlowerPlay,omitempty"`
	IsJiaoFenJiaBei  *bool                 `protobuf:"varint,7,opt,name=isJiaoFenJiaBei" json:"isJiaoFenJiaBei,omitempty"`
	XXX_unrecognized []byte                `json:"-"`
}

func (m *NiuniuDeskOption) Reset()                    { *m = NiuniuDeskOption{} }
func (m *NiuniuDeskOption) String() string            { return proto.CompactTextString(m) }
func (*NiuniuDeskOption) ProtoMessage()               {}
func (*NiuniuDeskOption) Descriptor() ([]byte, []int) { return fileDescriptor27, []int{2} }

func (m *NiuniuDeskOption) GetMinUser() int32 {
	if m != nil && m.MinUser != nil {
		return *m.MinUser
	}
	return 0
}

func (m *NiuniuDeskOption) GetMaxUser() int32 {
	if m != nil && m.MaxUser != nil {
		return *m.MaxUser
	}
	return 0
}

func (m *NiuniuDeskOption) GetMaxCircle() int32 {
	if m != nil && m.MaxCircle != nil {
		return *m.MaxCircle
	}
	return 0
}

func (m *NiuniuDeskOption) GetHasFlower() bool {
	if m != nil && m.HasFlower != nil {
		return *m.HasFlower
	}
	return false
}

func (m *NiuniuDeskOption) GetBankRule() NiuniuEnumBankerRule {
	if m != nil && m.BankRule != nil {
		return *m.BankRule
	}
	return NiuniuEnumBankerRule_DING_ZHUANG
}

func (m *NiuniuDeskOption) GetIsFlowerPlay() bool {
	if m != nil && m.IsFlowerPlay != nil {
		return *m.IsFlowerPlay
	}
	return false
}

func (m *NiuniuDeskOption) GetIsJiaoFenJiaBei() bool {
	if m != nil && m.IsJiaoFenJiaBei != nil {
		return *m.IsJiaoFenJiaBei
	}
	return false
}

func init() {
	proto.RegisterType((*NiuniuClientPoker)(nil), "ddproto.niuniu_client_poker")
	proto.RegisterType((*NiuniuUserBill)(nil), "ddproto.niuniu_user_bill")
	proto.RegisterType((*NiuniuDeskOption)(nil), "ddproto.niuniu_desk_option")
	proto.RegisterEnum("ddproto.NiuniuEnumProtoid", NiuniuEnumProtoid_name, NiuniuEnumProtoid_value)
	proto.RegisterEnum("ddproto.NiuniuEnum_PokerType", NiuniuEnum_PokerType_name, NiuniuEnum_PokerType_value)
	proto.RegisterEnum("ddproto.NiuniuEnumDeskState", NiuniuEnumDeskState_name, NiuniuEnumDeskState_value)
	proto.RegisterEnum("ddproto.NiuniuEnumBankerRule", NiuniuEnumBankerRule_name, NiuniuEnumBankerRule_value)
}

var fileDescriptor27 = []byte{
	// 956 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x55, 0xdb, 0x72, 0xe2, 0x46,
	0x10, 0x5d, 0x71, 0xa7, 0xc1, 0xf6, 0x78, 0x6c, 0x63, 0x19, 0xef, 0x85, 0x90, 0x7d, 0xa0, 0xfc,
	0xe0, 0xaa, 0x90, 0x7b, 0x55, 0x5e, 0x46, 0x42, 0x8b, 0x85, 0x59, 0xc1, 0xea, 0xb2, 0x8e, 0xfd,
	0x32, 0x25, 0x83, 0xaa, 0x56, 0x65, 0x90, 0x28, 0x84, 0x6a, 0xd7, 0x6f, 0xf9, 0x84, 0xfc, 0x4b,
	0xfe, 0x20, 0xff, 0x92, 0xff, 0x48, 0xcd, 0x48, 0xb2, 0xc7, 0x6b, 0xb4, 0x6f, 0xdd, 0xe7, 0x9c,
	0xee, 0x39, 0x3d, 0xa3, 0x06, 0xd8, 0x0f, 0xfc, 0x38, 0xf0, 0x63, 0x7a, 0xeb, 0x46, 0xde, 0xf9,
	0x6a, 0x1d, 0x6e, 0x42, 0x5c, 0x9d, 0xcf, 0x79, 0xd0, 0x3e, 0x98, 0x85, 0xcb, 0x65, 0x18, 0xd0,
	0xd9, 0xc2, 0xf7, 0x82, 0x4d, 0xc2, 0x76, 0xef, 0xe1, 0x20, 0x2d, 0x49, 0x60, 0xba, 0x0a, 0xef,
	0xbc, 0x35, 0x3e, 0x87, 0xd2, 0xca, 0xf5, 0x23, 0xb9, 0xd0, 0x29, 0xf6, 0x1a, 0xfd, 0xf6, 0x79,
	0xda, 0xe3, 0x3c, 0x15, 0xb1, 0xf6, 0x89, 0xd2, 0xe4, 0x3a, 0xdc, 0x87, 0xd2, 0xe6, 0x7e, 0xe5,
	0xc9, 0xc5, 0x8e, 0xd4, 0xdb, 0xed, 0xbf, 0x7e, 0xd0, 0xa7, 0xbd, 0xbd, 0x20, 0x5e, 0xd2, 0x29,
	0xd3, 0xdb, 0xf7, 0x2b, 0xcf, 0xe4, 0xda, 0xee, 0xbf, 0x12, 0xa0, 0x94, 0x8f, 0x23, 0x6f, 0x4d,
	0x6f, 0xfd, 0xc5, 0x02, 0x1f, 0x42, 0x39, 0x9a, 0x85, 0x6b, 0x4f, 0x96, 0x3a, 0x52, 0xaf, 0x6c,
	0x26, 0x09, 0xee, 0xc2, 0xce, 0x2c, 0x8c, 0x83, 0x0d, 0xfd, 0xe4, 0x46, 0x34, 0xf0, 0x63, 0xb9,
	0xc0, 0xd9, 0x06, 0x07, 0x2f, 0xdc, 0xc8, 0xf0, 0x63, 0xdc, 0x81, 0x66, 0xa2, 0x09, 0x42, 0x2e,
	0x29, 0x72, 0x09, 0x70, 0xcc, 0x08, 0x99, 0xe2, 0x14, 0xea, 0x89, 0xe2, 0xb3, 0x1f, 0xc8, 0x25,
	0x4e, 0xd7, 0x38, 0x70, 0xe5, 0x07, 0xf8, 0x15, 0x24, 0x52, 0xba, 0x08, 0xa3, 0x8d, 0x5c, 0xe6,
	0x6c, 0x22, 0x1f, 0x87, 0xd1, 0x06, 0xb7, 0xa0, 0xc2, 0x4c, 0xea, 0x73, 0xb9, 0xd2, 0x91, 0x7a,
	0x3b, 0x66, 0x9a, 0x75, 0xff, 0x2e, 0x00, 0x4e, 0x87, 0x98, 0x7b, 0xd1, 0x1d, 0x0d, 0x57, 0x1b,
	0x3f, 0x0c, 0xb0, 0x0c, 0xd5, 0xa5, 0x1f, 0x38, 0x91, 0xb7, 0x4e, 0x07, 0xc9, 0x52, 0xce, 0xb8,
	0x5f, 0x38, 0x53, 0x48, 0x99, 0x24, 0xc5, 0x2f, 0xa1, 0xbe, 0x74, 0xbf, 0xa8, 0xfe, 0x7a, 0xb6,
	0xf0, 0x52, 0xf7, 0x8f, 0x00, 0x63, 0x3f, 0xb9, 0xd1, 0xbb, 0x45, 0xf8, 0xd9, 0x5b, 0x73, 0xf3,
	0x35, 0xf3, 0x11, 0xc0, 0x7f, 0x40, 0xed, 0xd6, 0x0d, 0xee, 0xcc, 0x78, 0xe1, 0x71, 0xef, 0xbb,
	0xfd, 0xce, 0xd6, 0x37, 0x60, 0x22, 0x6f, 0x4d, 0xd7, 0xf1, 0xc2, 0x33, 0x1f, 0x2a, 0x70, 0x17,
	0x9a, 0x7e, 0xda, 0x69, 0xba, 0x70, 0xef, 0xf9, 0x88, 0x35, 0xf3, 0x09, 0x86, 0x7b, 0xb0, 0xe7,
	0x47, 0x23, 0xdf, 0x0d, 0xdf, 0x79, 0xc1, 0xc8, 0x77, 0x15, 0xcf, 0x97, 0xab, 0x5c, 0xf6, 0x35,
	0x7c, 0xf6, 0x4f, 0xf5, 0xe1, 0x9b, 0xe2, 0x67, 0x72, 0x17, 0xfe, 0x1c, 0x1f, 0xc1, 0xbe, 0xa1,
	0x3b, 0x74, 0xaa, 0x0f, 0xe8, 0x85, 0x46, 0x4c, 0x5b, 0xd1, 0x88, 0x8d, 0x5e, 0xe0, 0x16, 0xe0,
	0x0c, 0xfe, 0xe0, 0xe8, 0xea, 0x25, 0x55, 0x27, 0x86, 0x81, 0x24, 0xdc, 0x86, 0xd6, 0x73, 0x9c,
	0x12, 0xf5, 0x12, 0x15, 0xc4, 0x9a, 0x21, 0x79, 0xaf, 0xd1, 0xf1, 0x64, 0xa8, 0x1b, 0xa8, 0x28,
	0xd6, 0x3c, 0xe2, 0xbc, 0xa6, 0x84, 0x4f, 0xe1, 0x38, 0xe3, 0x54, 0x53, 0x23, 0xb6, 0x46, 0x07,
	0x9a, 0x75, 0x49, 0x4d, 0xed, 0x03, 0x2a, 0x8b, 0x85, 0x9a, 0x61, 0x6b, 0xe6, 0x23, 0x57, 0xc9,
	0xe1, 0x58, 0xd3, 0x2a, 0x3e, 0x81, 0xa3, 0x2d, 0x9c, 0xa2, 0xa2, 0x9a, 0x38, 0xae, 0xa9, 0x91,
	0xc1, 0x35, 0xef, 0x56, 0x7f, 0x0e, 0xb3, 0x46, 0x80, 0x0f, 0x01, 0x3d, 0x85, 0x15, 0x15, 0x35,
	0xc4, 0xf6, 0x96, 0x4d, 0x4c, 0x3b, 0x99, 0x6a, 0x62, 0xa3, 0xe6, 0x93, 0xeb, 0xd1, 0x89, 0x31,
	0xbc, 0xb9, 0x70, 0x88, 0x31, 0x64, 0xdc, 0x8e, 0x38, 0xaa, 0xc8, 0x31, 0x03, 0xbb, 0x79, 0x24,
	0xb3, 0xb1, 0x97, 0xd7, 0x55, 0x51, 0x11, 0x12, 0x9d, 0x8f, 0x74, 0xa2, 0x68, 0x3a, 0x3b, 0x6c,
	0x5f, 0x7c, 0x8b, 0x14, 0x66, 0xe7, 0xe0, 0x2d, 0x38, 0x3b, 0xe2, 0x60, 0x4b, 0x1b, 0x45, 0x45,
	0x87, 0xa2, 0x2d, 0x45, 0x9f, 0x12, 0xd6, 0xc5, 0x72, 0xc6, 0x36, 0x23, 0x8f, 0xf0, 0x31, 0x1c,
	0x3c, 0x79, 0x57, 0xcd, 0x18, 0x30, 0xa2, 0x85, 0x5f, 0x43, 0x3b, 0x23, 0xc8, 0x74, 0x3a, 0xbe,
	0xa6, 0x03, 0xdd, 0xb2, 0x26, 0xe3, 0x8f, 0x1a, 0x37, 0x71, 0xfc, 0x0d, 0x9e, 0x99, 0x91, 0xf1,
	0x2b, 0x38, 0xc9, 0xf8, 0x07, 0x46, 0x21, 0x6a, 0xf2, 0xf4, 0x27, 0xf9, 0x34, 0xab, 0x6e, 0xe3,
	0x97, 0x20, 0x3f, 0xa3, 0x07, 0x13, 0x43, 0x63, 0xde, 0x4e, 0x45, 0xd6, 0x62, 0x86, 0xdf, 0x6b,
	0x96, 0x45, 0x86, 0x89, 0xb3, 0x37, 0xe2, 0xf5, 0x64, 0x84, 0xa2, 0xa2, 0x8e, 0xf8, 0x02, 0x63,
	0x8d, 0x7c, 0x14, 0xbe, 0xd2, 0xef, 0x72, 0x38, 0xe6, 0xa5, 0x2b, 0x4e, 0x3a, 0xb9, 0x32, 0xd8,
	0x97, 0x28, 0xde, 0xc4, 0xf7, 0xdf, 0xe0, 0x59, 0xfd, 0xdb, 0xb3, 0xff, 0x24, 0x38, 0xda, 0xfa,
	0x6b, 0x8d, 0x01, 0x2a, 0xc6, 0x84, 0x1a, 0xba, 0x83, 0x24, 0x5c, 0x87, 0x32, 0xeb, 0xf2, 0x03,
	0x2a, 0x64, 0x61, 0x1f, 0x15, 0xb3, 0xf0, 0x47, 0x54, 0xca, 0xc2, 0x9f, 0x50, 0x39, 0x0b, 0x7f,
	0x46, 0x95, 0x2c, 0xfc, 0x05, 0x55, 0xb3, 0xf0, 0x57, 0x54, 0xcb, 0xc2, 0xdf, 0x50, 0x3d, 0x0b,
	0x7f, 0x47, 0x80, 0x1b, 0x50, 0x65, 0x21, 0x3b, 0xaf, 0xc1, 0x92, 0x6b, 0xdd, 0xe0, 0x49, 0x93,
	0x25, 0xa3, 0x34, 0xd9, 0xc1, 0x7b, 0xd0, 0xb8, 0x72, 0xe8, 0x9f, 0x3a, 0x49, 0xac, 0xed, 0x32,
	0x80, 0xd5, 0xdd, 0x5c, 0x10, 0x3a, 0x20, 0x06, 0xda, 0xcb, 0x56, 0xea, 0x5a, 0xa7, 0x36, 0x53,
	0x8d, 0x27, 0xc6, 0x10, 0xa1, 0xb3, 0xbf, 0x0a, 0xd0, 0x12, 0xe7, 0xe4, 0xbf, 0xda, 0xd1, 0xc6,
	0xdd, 0x78, 0xd9, 0x15, 0xf1, 0x4b, 0xb5, 0x6c, 0x62, 0x3b, 0x16, 0xbd, 0x22, 0xba, 0x9d, 0x6c,
	0x36, 0x92, 0x72, 0x79, 0xbe, 0xb0, 0xa8, 0x90, 0xcb, 0xf3, 0xd5, 0x45, 0x45, 0xfc, 0x16, 0x3a,
	0x5b, 0x79, 0x61, 0xd3, 0x50, 0x09, 0xbf, 0x81, 0xd3, 0xad, 0xaa, 0x64, 0x59, 0x50, 0x39, 0xf7,
	0x18, 0xbe, 0x36, 0xa8, 0x92, 0xdb, 0x20, 0x59, 0x28, 0x54, 0x3d, 0xb3, 0xe0, 0x38, 0xe7, 0x3f,
	0x81, 0x5d, 0xe2, 0x40, 0x37, 0x86, 0x34, 0x75, 0x23, 0xb1, 0x6d, 0xb5, 0x1c, 0x9d, 0x8e, 0x74,
	0x7a, 0xe3, 0x4c, 0x32, 0xb8, 0x80, 0x11, 0x34, 0xb9, 0xeb, 0x0c, 0x29, 0x4e, 0x5f, 0xfc, 0x1f,
	0x00, 0x00, 0xff, 0xff, 0xed, 0x14, 0xc4, 0x05, 0x7c, 0x08, 0x00, 0x00,
}
