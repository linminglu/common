// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common_pdk.proto

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

// Ignoring public import of common_req_list_coin_info from common_client.proto

// Ignoring public import of common_ack_list_coin_info from common_client.proto

// Ignoring public import of CommonCoinLevelInfo from common_client.proto

// Ignoring public import of common_enum_reg from common_client.proto

// Ignoring public import of common_enum_os_type from common_client.proto

// Ignoring public import of common_enum_pokerColor from common_client.proto

// 牌型
type PdkEnumPaiType int32

const (
	PdkEnumPaiType_PDK_ERRORCARD          PdkEnumPaiType = 0
	PdkEnumPaiType_PDK_SINGLECARD         PdkEnumPaiType = 1
	PdkEnumPaiType_PDK_DOUBLECARD         PdkEnumPaiType = 2
	PdkEnumPaiType_PDK_THREECARD          PdkEnumPaiType = 3
	PdkEnumPaiType_PDK_BOMBCARD           PdkEnumPaiType = 4
	PdkEnumPaiType_PDK_THREEONECARD       PdkEnumPaiType = 5
	PdkEnumPaiType_PDK_THREETWOCARD       PdkEnumPaiType = 6
	PdkEnumPaiType_PDK_CONNECTCARD        PdkEnumPaiType = 9
	PdkEnumPaiType_PDK_COMPANYCARD        PdkEnumPaiType = 10
	PdkEnumPaiType_PDK_AIRCRAFTCARD       PdkEnumPaiType = 11
	PdkEnumPaiType_PDK_AIRCRAFTSINGLECARD PdkEnumPaiType = 12
	PdkEnumPaiType_PDK_AIRCRAFTFOURSIG    PdkEnumPaiType = 16
	PdkEnumPaiType_PDK_AIRCRAFTLASTCARD   PdkEnumPaiType = 17
	PdkEnumPaiType_PDK_AIRCRAFTLASTDOU    PdkEnumPaiType = 18
)

var PdkEnumPaiType_name = map[int32]string{
	0:  "PDK_ERRORCARD",
	1:  "PDK_SINGLECARD",
	2:  "PDK_DOUBLECARD",
	3:  "PDK_THREECARD",
	4:  "PDK_BOMBCARD",
	5:  "PDK_THREEONECARD",
	6:  "PDK_THREETWOCARD",
	9:  "PDK_CONNECTCARD",
	10: "PDK_COMPANYCARD",
	11: "PDK_AIRCRAFTCARD",
	12: "PDK_AIRCRAFTSINGLECARD",
	16: "PDK_AIRCRAFTFOURSIG",
	17: "PDK_AIRCRAFTLASTCARD",
	18: "PDK_AIRCRAFTLASTDOU",
}
var PdkEnumPaiType_value = map[string]int32{
	"PDK_ERRORCARD":          0,
	"PDK_SINGLECARD":         1,
	"PDK_DOUBLECARD":         2,
	"PDK_THREECARD":          3,
	"PDK_BOMBCARD":           4,
	"PDK_THREEONECARD":       5,
	"PDK_THREETWOCARD":       6,
	"PDK_CONNECTCARD":        9,
	"PDK_COMPANYCARD":        10,
	"PDK_AIRCRAFTCARD":       11,
	"PDK_AIRCRAFTSINGLECARD": 12,
	"PDK_AIRCRAFTFOURSIG":    16,
	"PDK_AIRCRAFTLASTCARD":   17,
	"PDK_AIRCRAFTLASTDOU":    18,
}

func (x PdkEnumPaiType) Enum() *PdkEnumPaiType {
	p := new(PdkEnumPaiType)
	*p = x
	return p
}
func (x PdkEnumPaiType) String() string {
	return proto.EnumName(PdkEnumPaiType_name, int32(x))
}
func (x *PdkEnumPaiType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(PdkEnumPaiType_value, data, "PdkEnumPaiType")
	if err != nil {
		return err
	}
	*x = PdkEnumPaiType(value)
	return nil
}
func (PdkEnumPaiType) EnumDescriptor() ([]byte, []int) { return fileDescriptor8, []int{0} }

// 玩家当前状态
type PdkEnumPlayerStatus int32

const (
	PdkEnumPlayerStatus_PDK_NORMAL_DEFAULT      PdkEnumPlayerStatus = 0
	PdkEnumPlayerStatus_PDK_NORMAL_ENTERDESK    PdkEnumPlayerStatus = 1
	PdkEnumPlayerStatus_PDK_NORMAL_READY        PdkEnumPlayerStatus = 2
	PdkEnumPlayerStatus_PDK_NORMAL_DAPAI        PdkEnumPlayerStatus = 3
	PdkEnumPlayerStatus_PDK_NORMAL_LOTTERY      PdkEnumPlayerStatus = 4
	PdkEnumPlayerStatus_PDK_NORMAL_END          PdkEnumPlayerStatus = 5
	PdkEnumPlayerStatus_PDK_HLJIAO_NOACT        PdkEnumPlayerStatus = 6
	PdkEnumPlayerStatus_PDK_HLJIAO_JIAO         PdkEnumPlayerStatus = 7
	PdkEnumPlayerStatus_PDK_HLJIAO_BUJIAO       PdkEnumPlayerStatus = 8
	PdkEnumPlayerStatus_PDK_HLQIANG_NOACT       PdkEnumPlayerStatus = 9
	PdkEnumPlayerStatus_PDK_HLQIANG_QIANG       PdkEnumPlayerStatus = 10
	PdkEnumPlayerStatus_PDK_HLQIANG_BUQIANG     PdkEnumPlayerStatus = 11
	PdkEnumPlayerStatus_PDK_HLJB_NOACT          PdkEnumPlayerStatus = 12
	PdkEnumPlayerStatus_PDK_HLJB_JIA            PdkEnumPlayerStatus = 13
	PdkEnumPlayerStatus_PDK_HLJB_BUJIA          PdkEnumPlayerStatus = 14
	PdkEnumPlayerStatus_PDK_SCMZ_NOACT          PdkEnumPlayerStatus = 15
	PdkEnumPlayerStatus_PDK_SCMZ_MENZHUA        PdkEnumPlayerStatus = 16
	PdkEnumPlayerStatus_PDK_SCMZ_SEECARDS       PdkEnumPlayerStatus = 17
	PdkEnumPlayerStatus_PDK_SCZP_NOACT          PdkEnumPlayerStatus = 18
	PdkEnumPlayerStatus_PDK_SCZP_ZHUA           PdkEnumPlayerStatus = 19
	PdkEnumPlayerStatus_PDK_SCZP_BUZHUA         PdkEnumPlayerStatus = 20
	PdkEnumPlayerStatus_PDK_SCLD_NOACT          PdkEnumPlayerStatus = 21
	PdkEnumPlayerStatus_PDK_SCLD_LA             PdkEnumPlayerStatus = 22
	PdkEnumPlayerStatus_PDK_SCLD_DAO            PdkEnumPlayerStatus = 23
	PdkEnumPlayerStatus_PDK_SCLD_BULD           PdkEnumPlayerStatus = 24
	PdkEnumPlayerStatus_PDK_JDJIAO_NOACT        PdkEnumPlayerStatus = 25
	PdkEnumPlayerStatus_PDK_JDJIAO_JIAO         PdkEnumPlayerStatus = 26
	PdkEnumPlayerStatus_PDK_JDJIAO_BUJIAO       PdkEnumPlayerStatus = 27
	PdkEnumPlayerStatus_PDK_SHOWPOKER_READY     PdkEnumPlayerStatus = 28
	PdkEnumPlayerStatus_PDK_SHOWPOKER_ENTERDESK PdkEnumPlayerStatus = 29
)

var PdkEnumPlayerStatus_name = map[int32]string{
	0:  "PDK_NORMAL_DEFAULT",
	1:  "PDK_NORMAL_ENTERDESK",
	2:  "PDK_NORMAL_READY",
	3:  "PDK_NORMAL_DAPAI",
	4:  "PDK_NORMAL_LOTTERY",
	5:  "PDK_NORMAL_END",
	6:  "PDK_HLJIAO_NOACT",
	7:  "PDK_HLJIAO_JIAO",
	8:  "PDK_HLJIAO_BUJIAO",
	9:  "PDK_HLQIANG_NOACT",
	10: "PDK_HLQIANG_QIANG",
	11: "PDK_HLQIANG_BUQIANG",
	12: "PDK_HLJB_NOACT",
	13: "PDK_HLJB_JIA",
	14: "PDK_HLJB_BUJIA",
	15: "PDK_SCMZ_NOACT",
	16: "PDK_SCMZ_MENZHUA",
	17: "PDK_SCMZ_SEECARDS",
	18: "PDK_SCZP_NOACT",
	19: "PDK_SCZP_ZHUA",
	20: "PDK_SCZP_BUZHUA",
	21: "PDK_SCLD_NOACT",
	22: "PDK_SCLD_LA",
	23: "PDK_SCLD_DAO",
	24: "PDK_SCLD_BULD",
	25: "PDK_JDJIAO_NOACT",
	26: "PDK_JDJIAO_JIAO",
	27: "PDK_JDJIAO_BUJIAO",
	28: "PDK_SHOWPOKER_READY",
	29: "PDK_SHOWPOKER_ENTERDESK",
}
var PdkEnumPlayerStatus_value = map[string]int32{
	"PDK_NORMAL_DEFAULT":      0,
	"PDK_NORMAL_ENTERDESK":    1,
	"PDK_NORMAL_READY":        2,
	"PDK_NORMAL_DAPAI":        3,
	"PDK_NORMAL_LOTTERY":      4,
	"PDK_NORMAL_END":          5,
	"PDK_HLJIAO_NOACT":        6,
	"PDK_HLJIAO_JIAO":         7,
	"PDK_HLJIAO_BUJIAO":       8,
	"PDK_HLQIANG_NOACT":       9,
	"PDK_HLQIANG_QIANG":       10,
	"PDK_HLQIANG_BUQIANG":     11,
	"PDK_HLJB_NOACT":          12,
	"PDK_HLJB_JIA":            13,
	"PDK_HLJB_BUJIA":          14,
	"PDK_SCMZ_NOACT":          15,
	"PDK_SCMZ_MENZHUA":        16,
	"PDK_SCMZ_SEECARDS":       17,
	"PDK_SCZP_NOACT":          18,
	"PDK_SCZP_ZHUA":           19,
	"PDK_SCZP_BUZHUA":         20,
	"PDK_SCLD_NOACT":          21,
	"PDK_SCLD_LA":             22,
	"PDK_SCLD_DAO":            23,
	"PDK_SCLD_BULD":           24,
	"PDK_JDJIAO_NOACT":        25,
	"PDK_JDJIAO_JIAO":         26,
	"PDK_JDJIAO_BUJIAO":       27,
	"PDK_SHOWPOKER_READY":     28,
	"PDK_SHOWPOKER_ENTERDESK": 29,
}

func (x PdkEnumPlayerStatus) Enum() *PdkEnumPlayerStatus {
	p := new(PdkEnumPlayerStatus)
	*p = x
	return p
}
func (x PdkEnumPlayerStatus) String() string {
	return proto.EnumName(PdkEnumPlayerStatus_name, int32(x))
}
func (x *PdkEnumPlayerStatus) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(PdkEnumPlayerStatus_value, data, "PdkEnumPlayerStatus")
	if err != nil {
		return err
	}
	*x = PdkEnumPlayerStatus(value)
	return nil
}
func (PdkEnumPlayerStatus) EnumDescriptor() ([]byte, []int) { return fileDescriptor8, []int{1} }

type PdkEnumRoomType int32

const (
	PdkEnumRoomType_PDK_T_NORMAL_PDK  PdkEnumRoomType = 1
	PdkEnumRoomType_PDK_T_FIFTEEN_PDK PdkEnumRoomType = 2
	PdkEnumRoomType_PDK_T_SIREN_PDK   PdkEnumRoomType = 3
)

var PdkEnumRoomType_name = map[int32]string{
	1: "PDK_T_NORMAL_PDK",
	2: "PDK_T_FIFTEEN_PDK",
	3: "PDK_T_SIREN_PDK",
}
var PdkEnumRoomType_value = map[string]int32{
	"PDK_T_NORMAL_PDK":  1,
	"PDK_T_FIFTEEN_PDK": 2,
	"PDK_T_SIREN_PDK":   3,
}

func (x PdkEnumRoomType) Enum() *PdkEnumRoomType {
	p := new(PdkEnumRoomType)
	*p = x
	return p
}
func (x PdkEnumRoomType) String() string {
	return proto.EnumName(PdkEnumRoomType_name, int32(x))
}
func (x *PdkEnumRoomType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(PdkEnumRoomType_value, data, "PdkEnumRoomType")
	if err != nil {
		return err
	}
	*x = PdkEnumRoomType(value)
	return nil
}
func (PdkEnumRoomType) EnumDescriptor() ([]byte, []int) { return fileDescriptor8, []int{2} }

type PdkEnumEnterType int32

const (
	PdkEnumEnterType_PDK_FRIEND PdkEnumEnterType = 1
	PdkEnumEnterType_PDK_COIN   PdkEnumEnterType = 2
)

var PdkEnumEnterType_name = map[int32]string{
	1: "PDK_FRIEND",
	2: "PDK_COIN",
}
var PdkEnumEnterType_value = map[string]int32{
	"PDK_FRIEND": 1,
	"PDK_COIN":   2,
}

func (x PdkEnumEnterType) Enum() *PdkEnumEnterType {
	p := new(PdkEnumEnterType)
	*p = x
	return p
}
func (x PdkEnumEnterType) String() string {
	return proto.EnumName(PdkEnumEnterType_name, int32(x))
}
func (x *PdkEnumEnterType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(PdkEnumEnterType_value, data, "PdkEnumEnterType")
	if err != nil {
		return err
	}
	*x = PdkEnumEnterType(value)
	return nil
}
func (PdkEnumEnterType) EnumDescriptor() ([]byte, []int) { return fileDescriptor8, []int{3} }

type PdkEnumCoinRoomLevel int32

const (
	PdkEnumCoinRoomLevel_PDK_LV1 PdkEnumCoinRoomLevel = 1
	PdkEnumCoinRoomLevel_PDK_LV2 PdkEnumCoinRoomLevel = 2
	PdkEnumCoinRoomLevel_PDK_LV3 PdkEnumCoinRoomLevel = 3
	PdkEnumCoinRoomLevel_PDK_LV4 PdkEnumCoinRoomLevel = 4
)

var PdkEnumCoinRoomLevel_name = map[int32]string{
	1: "PDK_LV1",
	2: "PDK_LV2",
	3: "PDK_LV3",
	4: "PDK_LV4",
}
var PdkEnumCoinRoomLevel_value = map[string]int32{
	"PDK_LV1": 1,
	"PDK_LV2": 2,
	"PDK_LV3": 3,
	"PDK_LV4": 4,
}

func (x PdkEnumCoinRoomLevel) Enum() *PdkEnumCoinRoomLevel {
	p := new(PdkEnumCoinRoomLevel)
	*p = x
	return p
}
func (x PdkEnumCoinRoomLevel) String() string {
	return proto.EnumName(PdkEnumCoinRoomLevel_name, int32(x))
}
func (x *PdkEnumCoinRoomLevel) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(PdkEnumCoinRoomLevel_value, data, "PdkEnumCoinRoomLevel")
	if err != nil {
		return err
	}
	*x = PdkEnumCoinRoomLevel(value)
	return nil
}
func (PdkEnumCoinRoomLevel) EnumDescriptor() ([]byte, []int) { return fileDescriptor8, []int{4} }

type PdkEnumDeskGameStatus int32

const (
	PdkEnumDeskGameStatus_PDK_INIT       PdkEnumDeskGameStatus = 0
	PdkEnumDeskGameStatus_PDK_DEAL_CARDS PdkEnumDeskGameStatus = 1
	PdkEnumDeskGameStatus_PDK_JIAO_DIZHU PdkEnumDeskGameStatus = 2
	PdkEnumDeskGameStatus_PDK_PLAYING    PdkEnumDeskGameStatus = 3
	PdkEnumDeskGameStatus_PDK_FINISH     PdkEnumDeskGameStatus = 4
)

var PdkEnumDeskGameStatus_name = map[int32]string{
	0: "PDK_INIT",
	1: "PDK_DEAL_CARDS",
	2: "PDK_JIAO_DIZHU",
	3: "PDK_PLAYING",
	4: "PDK_FINISH",
}
var PdkEnumDeskGameStatus_value = map[string]int32{
	"PDK_INIT":       0,
	"PDK_DEAL_CARDS": 1,
	"PDK_JIAO_DIZHU": 2,
	"PDK_PLAYING":    3,
	"PDK_FINISH":     4,
}

func (x PdkEnumDeskGameStatus) Enum() *PdkEnumDeskGameStatus {
	p := new(PdkEnumDeskGameStatus)
	*p = x
	return p
}
func (x PdkEnumDeskGameStatus) String() string {
	return proto.EnumName(PdkEnumDeskGameStatus_name, int32(x))
}
func (x *PdkEnumDeskGameStatus) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(PdkEnumDeskGameStatus_value, data, "PdkEnumDeskGameStatus")
	if err != nil {
		return err
	}
	*x = PdkEnumDeskGameStatus(value)
	return nil
}
func (PdkEnumDeskGameStatus) EnumDescriptor() ([]byte, []int) { return fileDescriptor8, []int{5} }

type PdkBaseRoomTypeInfo struct {
	RoomType         *PdkEnumRoomType                `protobuf:"varint,1,opt,name=roomType,enum=ddproto.PdkEnumRoomType" json:"roomType,omitempty"`
	BoardsCount      *int32                          `protobuf:"varint,2,opt,name=boardsCount" json:"boardsCount,omitempty"`
	CapMax           *int64                          `protobuf:"varint,3,opt,name=capMax" json:"capMax,omitempty"`
	BaseValue        *int64                          `protobuf:"varint,4,opt,name=baseValue" json:"baseValue,omitempty"`
	IsJiaoFen        *bool                           `protobuf:"varint,5,opt,name=isJiaoFen" json:"isJiaoFen,omitempty"`
	UserCountLimit   *int32                          `protobuf:"varint,6,opt,name=UserCountLimit" json:"UserCountLimit,omitempty"`
	UserMinCoin      *int64                          `protobuf:"varint,7,opt,name=userMinCoin" json:"userMinCoin,omitempty"`
	UserMaxCoin      *int64                          `protobuf:"varint,8,opt,name=userMaxCoin" json:"userMaxCoin,omitempty"`
	CoinTicket       *int64                          `protobuf:"varint,9,opt,name=coinTicket" json:"coinTicket,omitempty"`
	RoomLv           *PdkEnumCoinRoomLevel           `protobuf:"varint,10,opt,name=roomLv,enum=ddproto.PdkEnumCoinRoomLevel" json:"roomLv,omitempty"`
	IsDaikai         *bool                           `protobuf:"varint,11,opt,name=isDaikai" json:"isDaikai,omitempty"`
	IsShowCardsNum   *bool                           `protobuf:"varint,12,opt,name=isShowCardsNum" json:"isShowCardsNum,omitempty"`
	IsZhuaNiao       *bool                           `protobuf:"varint,13,opt,name=isZhuaNiao" json:"isZhuaNiao,omitempty"`
	IsSpadeThree     *bool                           `protobuf:"varint,14,opt,name=isSpadeThree" json:"isSpadeThree,omitempty"`
	ChanelId         *int32                          `protobuf:"varint,15,opt,name=chanelId" json:"chanelId,omitempty"`
	RoomCardBillType *COMMON_ENUM_ROOMCARD_BILL_TYPE `protobuf:"varint,16,opt,name=roomCardBillType,enum=ddproto.COMMON_ENUM_ROOMCARD_BILL_TYPE" json:"roomCardBillType,omitempty"`
	XXX_unrecognized []byte                          `json:"-"`
}

func (m *PdkBaseRoomTypeInfo) Reset()                    { *m = PdkBaseRoomTypeInfo{} }
func (m *PdkBaseRoomTypeInfo) String() string            { return proto.CompactTextString(m) }
func (*PdkBaseRoomTypeInfo) ProtoMessage()               {}
func (*PdkBaseRoomTypeInfo) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{0} }

func (m *PdkBaseRoomTypeInfo) GetRoomType() PdkEnumRoomType {
	if m != nil && m.RoomType != nil {
		return *m.RoomType
	}
	return PdkEnumRoomType_PDK_T_NORMAL_PDK
}

func (m *PdkBaseRoomTypeInfo) GetBoardsCount() int32 {
	if m != nil && m.BoardsCount != nil {
		return *m.BoardsCount
	}
	return 0
}

func (m *PdkBaseRoomTypeInfo) GetCapMax() int64 {
	if m != nil && m.CapMax != nil {
		return *m.CapMax
	}
	return 0
}

func (m *PdkBaseRoomTypeInfo) GetBaseValue() int64 {
	if m != nil && m.BaseValue != nil {
		return *m.BaseValue
	}
	return 0
}

func (m *PdkBaseRoomTypeInfo) GetIsJiaoFen() bool {
	if m != nil && m.IsJiaoFen != nil {
		return *m.IsJiaoFen
	}
	return false
}

func (m *PdkBaseRoomTypeInfo) GetUserCountLimit() int32 {
	if m != nil && m.UserCountLimit != nil {
		return *m.UserCountLimit
	}
	return 0
}

func (m *PdkBaseRoomTypeInfo) GetUserMinCoin() int64 {
	if m != nil && m.UserMinCoin != nil {
		return *m.UserMinCoin
	}
	return 0
}

func (m *PdkBaseRoomTypeInfo) GetUserMaxCoin() int64 {
	if m != nil && m.UserMaxCoin != nil {
		return *m.UserMaxCoin
	}
	return 0
}

func (m *PdkBaseRoomTypeInfo) GetCoinTicket() int64 {
	if m != nil && m.CoinTicket != nil {
		return *m.CoinTicket
	}
	return 0
}

func (m *PdkBaseRoomTypeInfo) GetRoomLv() PdkEnumCoinRoomLevel {
	if m != nil && m.RoomLv != nil {
		return *m.RoomLv
	}
	return PdkEnumCoinRoomLevel_PDK_LV1
}

func (m *PdkBaseRoomTypeInfo) GetIsDaikai() bool {
	if m != nil && m.IsDaikai != nil {
		return *m.IsDaikai
	}
	return false
}

func (m *PdkBaseRoomTypeInfo) GetIsShowCardsNum() bool {
	if m != nil && m.IsShowCardsNum != nil {
		return *m.IsShowCardsNum
	}
	return false
}

func (m *PdkBaseRoomTypeInfo) GetIsZhuaNiao() bool {
	if m != nil && m.IsZhuaNiao != nil {
		return *m.IsZhuaNiao
	}
	return false
}

func (m *PdkBaseRoomTypeInfo) GetIsSpadeThree() bool {
	if m != nil && m.IsSpadeThree != nil {
		return *m.IsSpadeThree
	}
	return false
}

func (m *PdkBaseRoomTypeInfo) GetChanelId() int32 {
	if m != nil && m.ChanelId != nil {
		return *m.ChanelId
	}
	return 0
}

func (m *PdkBaseRoomTypeInfo) GetRoomCardBillType() COMMON_ENUM_ROOMCARD_BILL_TYPE {
	if m != nil && m.RoomCardBillType != nil {
		return *m.RoomCardBillType
	}
	return COMMON_ENUM_ROOMCARD_BILL_TYPE_OWNER_PAY
}

// 定时信息
type PdkBaseTimerInfo struct {
	UserId           *uint32 `protobuf:"varint,1,opt,name=userId" json:"userId,omitempty"`
	RemainSec        *int32  `protobuf:"varint,2,opt,name=remainSec" json:"remainSec,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *PdkBaseTimerInfo) Reset()                    { *m = PdkBaseTimerInfo{} }
func (m *PdkBaseTimerInfo) String() string            { return proto.CompactTextString(m) }
func (*PdkBaseTimerInfo) ProtoMessage()               {}
func (*PdkBaseTimerInfo) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{1} }

func (m *PdkBaseTimerInfo) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *PdkBaseTimerInfo) GetRemainSec() int32 {
	if m != nil && m.RemainSec != nil {
		return *m.RemainSec
	}
	return 0
}

func init() {
	proto.RegisterType((*PdkBaseRoomTypeInfo)(nil), "ddproto.pdk_base_roomTypeInfo")
	proto.RegisterType((*PdkBaseTimerInfo)(nil), "ddproto.pdk_base_timerInfo")
	proto.RegisterEnum("ddproto.PdkEnumPaiType", PdkEnumPaiType_name, PdkEnumPaiType_value)
	proto.RegisterEnum("ddproto.PdkEnumPlayerStatus", PdkEnumPlayerStatus_name, PdkEnumPlayerStatus_value)
	proto.RegisterEnum("ddproto.PdkEnumRoomType", PdkEnumRoomType_name, PdkEnumRoomType_value)
	proto.RegisterEnum("ddproto.PdkEnumEnterType", PdkEnumEnterType_name, PdkEnumEnterType_value)
	proto.RegisterEnum("ddproto.PdkEnumCoinRoomLevel", PdkEnumCoinRoomLevel_name, PdkEnumCoinRoomLevel_value)
	proto.RegisterEnum("ddproto.PdkEnumDeskGameStatus", PdkEnumDeskGameStatus_name, PdkEnumDeskGameStatus_value)
}

func init() { proto.RegisterFile("common_pdk.proto", fileDescriptor8) }

var fileDescriptor8 = []byte{
	// 1010 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x95, 0xdd, 0x52, 0xe3, 0x36,
	0x14, 0xc7, 0x49, 0xc2, 0x47, 0x50, 0x42, 0x10, 0x02, 0x82, 0xcb, 0x6e, 0x5b, 0x86, 0x8b, 0x96,
	0xc9, 0x05, 0x33, 0xa5, 0x9d, 0xf6, 0xda, 0xb1, 0x1d, 0xa2, 0xe0, 0x8f, 0x54, 0xb6, 0x77, 0x07,
	0x6e, 0x3c, 0xde, 0x44, 0x1d, 0x34, 0x24, 0x76, 0xc6, 0x4e, 0xb6, 0xbb, 0xd7, 0x7d, 0x94, 0xbe,
	0x4b, 0x9f, 0xab, 0x23, 0x59, 0xb2, 0x4d, 0xd8, 0x9b, 0x0c, 0xfa, 0x9d, 0xff, 0x39, 0x3a, 0xe7,
	0xf8, 0x6f, 0x03, 0xe0, 0x2c, 0x5d, 0x2e, 0xd3, 0x24, 0x5a, 0xcd, 0x5f, 0x6e, 0x57, 0x59, 0xba,
	0x4e, 0xd1, 0xc1, 0x7c, 0x2e, 0xfe, 0xb8, 0x3c, 0x95, 0xa1, 0xd9, 0x82, 0xd1, 0x64, 0x5d, 0x44,
	0xaf, 0xff, 0xd9, 0x03, 0xe7, 0xab, 0xf9, 0x4b, 0xf4, 0x29, 0xce, 0x69, 0x94, 0xa5, 0xe9, 0x32,
	0xf8, 0xba, 0xa2, 0x38, 0xf9, 0x2b, 0x45, 0xbf, 0x83, 0xb6, 0x3a, 0x6b, 0x8d, 0xab, 0xc6, 0x4d,
	0xef, 0xee, 0xf2, 0x56, 0x96, 0xba, 0xe5, 0x19, 0x34, 0xd9, 0x2c, 0xcb, 0x0c, 0x52, 0x6a, 0xd1,
	0x15, 0xe8, 0x7c, 0x4a, 0xe3, 0x6c, 0x9e, 0x1b, 0xe9, 0x26, 0x59, 0x6b, 0xcd, 0xab, 0xc6, 0xcd,
	0x1e, 0xa9, 0x23, 0xd4, 0x07, 0xfb, 0xb3, 0x78, 0xe5, 0xc4, 0x5f, 0xb4, 0xd6, 0x55, 0xe3, 0xa6,
	0x45, 0xe4, 0x09, 0xbd, 0x07, 0x87, 0xbc, 0x8d, 0x0f, 0xf1, 0x62, 0x43, 0xb5, 0x5d, 0x11, 0xaa,
	0x00, 0x8f, 0xb2, 0x7c, 0xc2, 0xe2, 0x74, 0x44, 0x13, 0x6d, 0xef, 0xaa, 0x71, 0xd3, 0x26, 0x15,
	0x40, 0x3f, 0x81, 0x5e, 0x98, 0xd3, 0x4c, 0x5c, 0x60, 0xb3, 0x25, 0x5b, 0x6b, 0xfb, 0xe2, 0xe2,
	0x2d, 0xca, 0xbb, 0xdb, 0xe4, 0x34, 0x73, 0x58, 0x62, 0xa4, 0x2c, 0xd1, 0x0e, 0xc4, 0x2d, 0x75,
	0x54, 0x2a, 0xe2, 0x2f, 0x42, 0xd1, 0xae, 0x29, 0x0a, 0x84, 0x7e, 0x00, 0x60, 0x96, 0xb2, 0x24,
	0x60, 0xb3, 0x17, 0xba, 0xd6, 0x0e, 0x85, 0xa0, 0x46, 0xd0, 0x1f, 0x60, 0x9f, 0x6f, 0xc3, 0xfe,
	0xac, 0x01, 0xb1, 0xb7, 0x1f, 0xdf, 0xee, 0x8d, 0xab, 0x09, 0xd7, 0xd0, 0xcf, 0x74, 0x41, 0xa4,
	0x1c, 0x5d, 0x82, 0x36, 0xcb, 0xcd, 0x98, 0xbd, 0xc4, 0x4c, 0xeb, 0x88, 0x09, 0xcb, 0x33, 0x1f,
	0x90, 0xe5, 0xfe, 0x73, 0xfa, 0xb7, 0xc1, 0x17, 0xe9, 0x6e, 0x96, 0x5a, 0x57, 0x28, 0xb6, 0x28,
	0x6f, 0x8e, 0xe5, 0x4f, 0xcf, 0x9b, 0xd8, 0x65, 0x71, 0xaa, 0x1d, 0x09, 0x4d, 0x8d, 0xa0, 0x6b,
	0xd0, 0x65, 0xb9, 0xbf, 0x8a, 0xe7, 0x34, 0x78, 0xce, 0x28, 0xd5, 0x7a, 0x42, 0xf1, 0x8a, 0xf1,
	0x3e, 0x66, 0xcf, 0x71, 0x42, 0x17, 0x78, 0xae, 0x1d, 0x8b, 0x35, 0x96, 0x67, 0xe4, 0x03, 0xc8,
	0xbb, 0xe5, 0xf7, 0x0d, 0xd9, 0x62, 0x21, 0xec, 0x01, 0xc5, 0x98, 0x3f, 0x97, 0x63, 0x1a, 0x9e,
	0xe3, 0x78, 0x6e, 0x64, 0xb9, 0xa1, 0x13, 0x11, 0xcf, 0x73, 0x0c, 0x9d, 0x98, 0xd1, 0x10, 0xdb,
	0x76, 0x14, 0x3c, 0x4e, 0x2d, 0xf2, 0xa6, 0xc0, 0xf5, 0x04, 0xa0, 0xd2, 0x84, 0x6b, 0xb6, 0xa4,
	0x99, 0x70, 0x60, 0x1f, 0xec, 0xf3, 0xb5, 0xe3, 0xb9, 0xf0, 0xdf, 0x11, 0x91, 0x27, 0xee, 0x84,
	0x8c, 0x2e, 0x63, 0x96, 0xf8, 0x74, 0x26, 0xfd, 0x55, 0x81, 0xc1, 0x7f, 0x4d, 0x00, 0xcb, 0x3d,
	0xaf, 0x62, 0x26, 0x4c, 0x79, 0x02, 0x8e, 0xa6, 0xe6, 0x43, 0x64, 0x11, 0xe2, 0x11, 0xde, 0x0d,
	0xdc, 0x41, 0x08, 0xf4, 0x38, 0xf2, 0xb1, 0x7b, 0x6f, 0x5b, 0x82, 0x35, 0x14, 0x33, 0xbd, 0x70,
	0x28, 0x59, 0x53, 0xa5, 0x06, 0x63, 0x62, 0x15, 0xa8, 0x85, 0x20, 0xe8, 0x72, 0x34, 0xf4, 0x9c,
	0xa1, 0x20, 0xbb, 0xe8, 0x0c, 0xc0, 0x52, 0xe4, 0xb9, 0x85, 0x6e, 0xef, 0x15, 0x0d, 0x3e, 0x7a,
	0x82, 0xee, 0xa3, 0x53, 0x70, 0xcc, 0xa9, 0xe1, 0xb9, 0xae, 0x65, 0x04, 0x02, 0x1e, 0x56, 0xd0,
	0x99, 0xea, 0xee, 0xa3, 0x80, 0x40, 0xe5, 0xeb, 0x98, 0x18, 0x44, 0x1f, 0x15, 0xd2, 0x0e, 0xba,
	0x04, 0xfd, 0x3a, 0xad, 0x0d, 0xd0, 0x45, 0x17, 0xe0, 0xb4, 0x1e, 0x1b, 0x79, 0x21, 0xf1, 0xf1,
	0x3d, 0x84, 0x48, 0x03, 0x67, 0xf5, 0x80, 0xad, 0xfb, 0x45, 0xb9, 0x93, 0xed, 0x14, 0x1e, 0x31,
	0xbd, 0x10, 0xa2, 0xc1, 0xbf, 0xf2, 0xd3, 0x50, 0x2c, 0x72, 0x11, 0x7f, 0xa5, 0x99, 0xbf, 0x8e,
	0xd7, 0x9b, 0x1c, 0xf5, 0x01, 0xe2, 0x29, 0xae, 0x47, 0x1c, 0xdd, 0x8e, 0x4c, 0x6b, 0xa4, 0x87,
	0x76, 0x00, 0x77, 0xd4, 0x25, 0x92, 0x5b, 0x6e, 0x60, 0x11, 0xd3, 0xf2, 0x1f, 0x60, 0x43, 0x4d,
	0x22, 0x23, 0xc4, 0xd2, 0xcd, 0x47, 0xd8, 0xdc, 0xa2, 0xa6, 0x3e, 0xd5, 0x31, 0x6c, 0x6d, 0x55,
	0xb7, 0xbd, 0x20, 0xb0, 0xc8, 0x23, 0xdc, 0x55, 0x0f, 0xa7, 0xac, 0x5e, 0xdb, 0xf0, 0xd8, 0x9e,
	0x60, 0xdd, 0x8b, 0x5c, 0x4f, 0x37, 0x82, 0x6a, 0xc3, 0x92, 0xf2, 0x1f, 0x78, 0x80, 0xce, 0xc1,
	0x49, 0x0d, 0x0e, 0x43, 0x81, 0xdb, 0x15, 0xfe, 0x13, 0xeb, 0xee, 0xbd, 0x2c, 0x71, 0xb8, 0x8d,
	0xc5, 0x2f, 0x04, 0x6a, 0x59, 0x0a, 0x0f, 0xc3, 0x22, 0xd0, 0x51, 0xcd, 0x8d, 0xed, 0xc9, 0x50,
	0xd6, 0xe8, 0x2a, 0x9b, 0x08, 0x36, 0xc1, 0x3a, 0x3c, 0x7a, 0xa5, 0x12, 0x1d, 0xc0, 0x5e, 0xe9,
	0x43, 0xc3, 0x79, 0x92, 0x99, 0xc7, 0x6a, 0x2c, 0xc1, 0x1c, 0xcb, 0x7d, 0x1a, 0x87, 0x3a, 0x84,
	0xaa, 0x27, 0x41, 0xfd, 0xc2, 0x8c, 0x3e, 0x3c, 0xa9, 0x0a, 0x3c, 0x4d, 0x65, 0x01, 0xa4, 0x4c,
	0x2b, 0x98, 0xc8, 0x3e, 0x55, 0x4b, 0x11, 0x68, 0x18, 0x0a, 0x78, 0x56, 0xe5, 0xda, 0xa6, 0xcc,
	0x3d, 0x47, 0xc7, 0xa0, 0x53, 0x32, 0x5b, 0x87, 0x7d, 0x35, 0x87, 0x00, 0xa6, 0xee, 0xc1, 0x8b,
	0xaa, 0xbc, 0x6d, 0x46, 0xc3, 0xd0, 0x36, 0xa1, 0xa6, 0x5a, 0x9e, 0x98, 0xb5, 0x27, 0xf1, 0x9d,
	0xba, 0x54, 0x52, 0xb1, 0xf2, 0x4b, 0x35, 0x87, 0x84, 0xf2, 0x49, 0xbc, 0x53, 0xbb, 0xf5, 0xc7,
	0xde, 0xc7, 0xa9, 0xf7, 0x60, 0x11, 0x69, 0x93, 0xf7, 0xe8, 0x1d, 0xb8, 0x78, 0x1d, 0xa8, 0x9c,
	0xf5, 0xfd, 0x20, 0x04, 0x27, 0x6f, 0xfe, 0x1b, 0x95, 0x2f, 0x9e, 0x32, 0xcb, 0xd4, 0xe4, 0x26,
	0x94, 0xf7, 0x06, 0xd1, 0x08, 0x8f, 0x02, 0xcb, 0x72, 0x05, 0x6e, 0xaa, 0x1e, 0x83, 0xc8, 0xc7,
	0x44, 0xc2, 0xd6, 0xe0, 0xae, 0xf8, 0x22, 0x89, 0xb2, 0x34, 0x59, 0xd3, 0x4c, 0xd4, 0xed, 0x01,
	0xc0, 0xa5, 0x23, 0x82, 0xb9, 0xfd, 0x1a, 0xa8, 0x0b, 0xda, 0xc5, 0x5b, 0x8b, 0x5d, 0xd8, 0x1c,
	0xd8, 0xa0, 0xff, 0xed, 0x0f, 0x3c, 0xea, 0x80, 0x03, 0xae, 0xb3, 0x3f, 0xfc, 0x02, 0x1b, 0xd5,
	0xe1, 0x0e, 0x36, 0xab, 0xc3, 0xaf, 0xb0, 0x55, 0x1d, 0x7e, 0x83, 0xbb, 0x83, 0x15, 0xb8, 0x28,
	0xab, 0xcd, 0x69, 0xfe, 0x72, 0x1f, 0x2f, 0xa9, 0x7c, 0xff, 0xe4, 0xb5, 0xd8, 0xc5, 0x41, 0xf5,
	0x21, 0x33, 0x2d, 0xdd, 0x8e, 0x0a, 0x4f, 0x94, 0x1f, 0x32, 0xb1, 0x60, 0x13, 0x3f, 0x8d, 0x43,
	0xd8, 0x54, 0xcf, 0x75, 0x6a, 0xeb, 0x8f, 0xd8, 0xbd, 0x87, 0xad, 0x72, 0x1a, 0xec, 0x62, 0x7f,
	0x0c, 0x77, 0xa7, 0x3b, 0xff, 0x07, 0x00, 0x00, 0xff, 0xff, 0x07, 0x8d, 0x60, 0x82, 0x3e, 0x08,
	0x00, 0x00,
}
