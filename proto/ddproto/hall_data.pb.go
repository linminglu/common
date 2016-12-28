// Code generated by protoc-gen-go.
// source: hall_data.proto
// DO NOT EDIT!

package ddproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type HallEnumProtoId int32

const (
	HallEnumProtoId_HALL_PID_HEARTBEAT                   HallEnumProtoId = 0
	HallEnumProtoId_HALL_PID_QUICK_CONN                  HallEnumProtoId = 1
	HallEnumProtoId_HALL_PID_QUICK_CONN_ACK              HallEnumProtoId = 2
	HallEnumProtoId_HALL_PID_GAME_LOGIN                  HallEnumProtoId = 3
	HallEnumProtoId_HALL_PID_GAME_LOGIN_ACK              HallEnumProtoId = 4
	HallEnumProtoId_HALL_PID_USER_DATA                   HallEnumProtoId = 9
	HallEnumProtoId_HALL_PID_USER_DATA_ACK               HallEnumProtoId = 10
	HallEnumProtoId_HALL_PID_DRAW_LOTTERY                HallEnumProtoId = 11
	HallEnumProtoId_HALL_PID_DRAW_LOTTERY_ACK            HallEnumProtoId = 12
	HallEnumProtoId_HALL_PID_DS_LOTTERY_ITEMS_ACK        HallEnumProtoId = 13
	HallEnumProtoId_HALL_PID_SIGN_LOTTERY                HallEnumProtoId = 14
	HallEnumProtoId_HALL_PID_SIGN_LOTTERY_ACK            HallEnumProtoId = 15
	HallEnumProtoId_HALL_PID_ONLINEWARD_REQ              HallEnumProtoId = 16
	HallEnumProtoId_HALL_PID_ONLINEWARD_ACK              HallEnumProtoId = 17
	HallEnumProtoId_HALL_PID_EVENT_REQ                   HallEnumProtoId = 18
	HallEnumProtoId_HALL_PID_EVENT_ACK                   HallEnumProtoId = 19
	HallEnumProtoId_HALL_PID_GOODS_LIST_REQ              HallEnumProtoId = 20
	HallEnumProtoId_HALL_PID_GOODS_LIST_ACK              HallEnumProtoId = 21
	HallEnumProtoId_HALL_PID_HOTUPDATEREQVERSIONINFO     HallEnumProtoId = 22
	HallEnumProtoId_HALL_PID_HOTUPDATEACKVERSIONINFO     HallEnumProtoId = 23
	HallEnumProtoId_HALL_PID_BAG_ITEMS_REQ               HallEnumProtoId = 24
	HallEnumProtoId_HALL_PID_BAG_ITEMS_ACK               HallEnumProtoId = 25
	HallEnumProtoId_HALL_PID_HOTUPDATEGAMEASSETSINFO_REQ HallEnumProtoId = 26
	HallEnumProtoId_HALL_PID_HOTUPDATEGAMEASSETSINFO_ACK HallEnumProtoId = 27
	HallEnumProtoId_HALL_PID_APPLE_PAY_CB_REQ            HallEnumProtoId = 28
	HallEnumProtoId_HALL_PID_HOTUPDATEASSETSINFO_REQ     HallEnumProtoId = 29
	HallEnumProtoId_HALL_PID_HOTUPDATEASSETSINFO_ACK     HallEnumProtoId = 30
)

var HallEnumProtoId_name = map[int32]string{
	0:  "HALL_PID_HEARTBEAT",
	1:  "HALL_PID_QUICK_CONN",
	2:  "HALL_PID_QUICK_CONN_ACK",
	3:  "HALL_PID_GAME_LOGIN",
	4:  "HALL_PID_GAME_LOGIN_ACK",
	9:  "HALL_PID_USER_DATA",
	10: "HALL_PID_USER_DATA_ACK",
	11: "HALL_PID_DRAW_LOTTERY",
	12: "HALL_PID_DRAW_LOTTERY_ACK",
	13: "HALL_PID_DS_LOTTERY_ITEMS_ACK",
	14: "HALL_PID_SIGN_LOTTERY",
	15: "HALL_PID_SIGN_LOTTERY_ACK",
	16: "HALL_PID_ONLINEWARD_REQ",
	17: "HALL_PID_ONLINEWARD_ACK",
	18: "HALL_PID_EVENT_REQ",
	19: "HALL_PID_EVENT_ACK",
	20: "HALL_PID_GOODS_LIST_REQ",
	21: "HALL_PID_GOODS_LIST_ACK",
	22: "HALL_PID_HOTUPDATEREQVERSIONINFO",
	23: "HALL_PID_HOTUPDATEACKVERSIONINFO",
	24: "HALL_PID_BAG_ITEMS_REQ",
	25: "HALL_PID_BAG_ITEMS_ACK",
	26: "HALL_PID_HOTUPDATEGAMEASSETSINFO_REQ",
	27: "HALL_PID_HOTUPDATEGAMEASSETSINFO_ACK",
	28: "HALL_PID_APPLE_PAY_CB_REQ",
	29: "HALL_PID_HOTUPDATEASSETSINFO_REQ",
	30: "HALL_PID_HOTUPDATEASSETSINFO_ACK",
}
var HallEnumProtoId_value = map[string]int32{
	"HALL_PID_HEARTBEAT":                   0,
	"HALL_PID_QUICK_CONN":                  1,
	"HALL_PID_QUICK_CONN_ACK":              2,
	"HALL_PID_GAME_LOGIN":                  3,
	"HALL_PID_GAME_LOGIN_ACK":              4,
	"HALL_PID_USER_DATA":                   9,
	"HALL_PID_USER_DATA_ACK":               10,
	"HALL_PID_DRAW_LOTTERY":                11,
	"HALL_PID_DRAW_LOTTERY_ACK":            12,
	"HALL_PID_DS_LOTTERY_ITEMS_ACK":        13,
	"HALL_PID_SIGN_LOTTERY":                14,
	"HALL_PID_SIGN_LOTTERY_ACK":            15,
	"HALL_PID_ONLINEWARD_REQ":              16,
	"HALL_PID_ONLINEWARD_ACK":              17,
	"HALL_PID_EVENT_REQ":                   18,
	"HALL_PID_EVENT_ACK":                   19,
	"HALL_PID_GOODS_LIST_REQ":              20,
	"HALL_PID_GOODS_LIST_ACK":              21,
	"HALL_PID_HOTUPDATEREQVERSIONINFO":     22,
	"HALL_PID_HOTUPDATEACKVERSIONINFO":     23,
	"HALL_PID_BAG_ITEMS_REQ":               24,
	"HALL_PID_BAG_ITEMS_ACK":               25,
	"HALL_PID_HOTUPDATEGAMEASSETSINFO_REQ": 26,
	"HALL_PID_HOTUPDATEGAMEASSETSINFO_ACK": 27,
	"HALL_PID_APPLE_PAY_CB_REQ":            28,
	"HALL_PID_HOTUPDATEASSETSINFO_REQ":     29,
	"HALL_PID_HOTUPDATEASSETSINFO_ACK":     30,
}

func (x HallEnumProtoId) Enum() *HallEnumProtoId {
	p := new(HallEnumProtoId)
	*p = x
	return p
}
func (x HallEnumProtoId) String() string {
	return proto.EnumName(HallEnumProtoId_name, int32(x))
}
func (x *HallEnumProtoId) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(HallEnumProtoId_value, data, "HallEnumProtoId")
	if err != nil {
		return err
	}
	*x = HallEnumProtoId(value)
	return nil
}
func (HallEnumProtoId) EnumDescriptor() ([]byte, []int) { return fileDescriptor14, []int{0} }

// 活动类型
type HallEnumEvent int32

const (
	HallEnumEvent_TYPE_TIME HallEnumEvent = 1
	HallEnumEvent_TYPE_NEW  HallEnumEvent = 2
	HallEnumEvent_TYPE_NULL HallEnumEvent = 3
)

var HallEnumEvent_name = map[int32]string{
	1: "TYPE_TIME",
	2: "TYPE_NEW",
	3: "TYPE_NULL",
}
var HallEnumEvent_value = map[string]int32{
	"TYPE_TIME": 1,
	"TYPE_NEW":  2,
	"TYPE_NULL": 3,
}

func (x HallEnumEvent) Enum() *HallEnumEvent {
	p := new(HallEnumEvent)
	*p = x
	return p
}
func (x HallEnumEvent) String() string {
	return proto.EnumName(HallEnumEvent_name, int32(x))
}
func (x *HallEnumEvent) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(HallEnumEvent_value, data, "HallEnumEvent")
	if err != nil {
		return err
	}
	*x = HallEnumEvent(value)
	return nil
}
func (HallEnumEvent) EnumDescriptor() ([]byte, []int) { return fileDescriptor14, []int{1} }

// 活动奖品
type HallEnum_Reward int32

const (
	HallEnum_Reward_RE_EXP  HallEnum_Reward = 1
	HallEnum_Reward_RE_GIFT HallEnum_Reward = 2
)

var HallEnum_Reward_name = map[int32]string{
	1: "RE_EXP",
	2: "RE_GIFT",
}
var HallEnum_Reward_value = map[string]int32{
	"RE_EXP":  1,
	"RE_GIFT": 2,
}

func (x HallEnum_Reward) Enum() *HallEnum_Reward {
	p := new(HallEnum_Reward)
	*p = x
	return p
}
func (x HallEnum_Reward) String() string {
	return proto.EnumName(HallEnum_Reward_name, int32(x))
}
func (x *HallEnum_Reward) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(HallEnum_Reward_value, data, "HallEnum_Reward")
	if err != nil {
		return err
	}
	*x = HallEnum_Reward(value)
	return nil
}
func (HallEnum_Reward) EnumDescriptor() ([]byte, []int) { return fileDescriptor14, []int{2} }

// 邮件类型
type HallEnumMailType int32

const (
	HallEnumMailType_SYSTEM     HallEnumMailType = 1
	HallEnumMailType_FRIEND_ADD HallEnumMailType = 2
)

var HallEnumMailType_name = map[int32]string{
	1: "SYSTEM",
	2: "FRIEND_ADD",
}
var HallEnumMailType_value = map[string]int32{
	"SYSTEM":     1,
	"FRIEND_ADD": 2,
}

func (x HallEnumMailType) Enum() *HallEnumMailType {
	p := new(HallEnumMailType)
	*p = x
	return p
}
func (x HallEnumMailType) String() string {
	return proto.EnumName(HallEnumMailType_name, int32(x))
}
func (x *HallEnumMailType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(HallEnumMailType_value, data, "HallEnumMailType")
	if err != nil {
		return err
	}
	*x = HallEnumMailType(value)
	return nil
}
func (HallEnumMailType) EnumDescriptor() ([]byte, []int) { return fileDescriptor14, []int{3} }

// 道具类型
type HallEnumPropsType int32

const (
	HallEnumPropsType_TYPE_LABA     HallEnumPropsType = 1
	HallEnumPropsType_TYPE_JIPAIQI  HallEnumPropsType = 2
	HallEnumPropsType_TYPE_FANGKA   HallEnumPropsType = 3
	HallEnumPropsType_TYPE_EXP_3000 HallEnumPropsType = 4
)

var HallEnumPropsType_name = map[int32]string{
	1: "TYPE_LABA",
	2: "TYPE_JIPAIQI",
	3: "TYPE_FANGKA",
	4: "TYPE_EXP_3000",
}
var HallEnumPropsType_value = map[string]int32{
	"TYPE_LABA":     1,
	"TYPE_JIPAIQI":  2,
	"TYPE_FANGKA":   3,
	"TYPE_EXP_3000": 4,
}

func (x HallEnumPropsType) Enum() *HallEnumPropsType {
	p := new(HallEnumPropsType)
	*p = x
	return p
}
func (x HallEnumPropsType) String() string {
	return proto.EnumName(HallEnumPropsType_name, int32(x))
}
func (x *HallEnumPropsType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(HallEnumPropsType_value, data, "HallEnumPropsType")
	if err != nil {
		return err
	}
	*x = HallEnumPropsType(value)
	return nil
}
func (HallEnumPropsType) EnumDescriptor() ([]byte, []int) { return fileDescriptor14, []int{4} }

// 任务类型
type HallEnumTaskType int32

const (
	HallEnumTaskType_TYPE_MJ  HallEnumTaskType = 1
	HallEnumTaskType_TYPE_DDZ HallEnumTaskType = 2
	HallEnumTaskType_TYPE_ZJH HallEnumTaskType = 3
)

var HallEnumTaskType_name = map[int32]string{
	1: "TYPE_MJ",
	2: "TYPE_DDZ",
	3: "TYPE_ZJH",
}
var HallEnumTaskType_value = map[string]int32{
	"TYPE_MJ":  1,
	"TYPE_DDZ": 2,
	"TYPE_ZJH": 3,
}

func (x HallEnumTaskType) Enum() *HallEnumTaskType {
	p := new(HallEnumTaskType)
	*p = x
	return p
}
func (x HallEnumTaskType) String() string {
	return proto.EnumName(HallEnumTaskType_name, int32(x))
}
func (x *HallEnumTaskType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(HallEnumTaskType_value, data, "HallEnumTaskType")
	if err != nil {
		return err
	}
	*x = HallEnumTaskType(value)
	return nil
}
func (HallEnumTaskType) EnumDescriptor() ([]byte, []int) { return fileDescriptor14, []int{5} }

// vip等级
type HallUser_VIP int32

const (
	HallUser_VIP_LV_1 HallUser_VIP = 1
	HallUser_VIP_LV_2 HallUser_VIP = 2
	HallUser_VIP_LV_3 HallUser_VIP = 3
	HallUser_VIP_LV_4 HallUser_VIP = 4
	HallUser_VIP_LV_5 HallUser_VIP = 5
	HallUser_VIP_LV_6 HallUser_VIP = 6
)

var HallUser_VIP_name = map[int32]string{
	1: "LV_1",
	2: "LV_2",
	3: "LV_3",
	4: "LV_4",
	5: "LV_5",
	6: "LV_6",
}
var HallUser_VIP_value = map[string]int32{
	"LV_1": 1,
	"LV_2": 2,
	"LV_3": 3,
	"LV_4": 4,
	"LV_5": 5,
	"LV_6": 6,
}

func (x HallUser_VIP) Enum() *HallUser_VIP {
	p := new(HallUser_VIP)
	*p = x
	return p
}
func (x HallUser_VIP) String() string {
	return proto.EnumName(HallUser_VIP_name, int32(x))
}
func (x *HallUser_VIP) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(HallUser_VIP_value, data, "HallUser_VIP")
	if err != nil {
		return err
	}
	*x = HallUser_VIP(value)
	return nil
}
func (HallUser_VIP) EnumDescriptor() ([]byte, []int) { return fileDescriptor14, []int{6} }

// 奖励物品信息
type HallLotteryItem struct {
	Id               *int32  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name             *string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *HallLotteryItem) Reset()                    { *m = HallLotteryItem{} }
func (m *HallLotteryItem) String() string            { return proto.CompactTextString(m) }
func (*HallLotteryItem) ProtoMessage()               {}
func (*HallLotteryItem) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{0} }

func (m *HallLotteryItem) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *HallLotteryItem) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

// 签到信息与奖励物品
type HallSignLotteryInfo struct {
	SignItems        []*HallLotteryItem `protobuf:"bytes,1,rep,name=signItems" json:"signItems,omitempty"`
	TotalDays        *int32             `protobuf:"varint,2,opt,name=totalDays" json:"totalDays,omitempty"`
	ContinuousDays   *int32             `protobuf:"varint,3,opt,name=continuousDays" json:"continuousDays,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *HallSignLotteryInfo) Reset()                    { *m = HallSignLotteryInfo{} }
func (m *HallSignLotteryInfo) String() string            { return proto.CompactTextString(m) }
func (*HallSignLotteryInfo) ProtoMessage()               {}
func (*HallSignLotteryInfo) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{1} }

func (m *HallSignLotteryInfo) GetSignItems() []*HallLotteryItem {
	if m != nil {
		return m.SignItems
	}
	return nil
}

func (m *HallSignLotteryInfo) GetTotalDays() int32 {
	if m != nil && m.TotalDays != nil {
		return *m.TotalDays
	}
	return 0
}

func (m *HallSignLotteryInfo) GetContinuousDays() int32 {
	if m != nil && m.ContinuousDays != nil {
		return *m.ContinuousDays
	}
	return 0
}

// 单个活动
type HallItemEvent struct {
	Id               *int32           `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Type             *HallEnumEvent   `protobuf:"varint,2,opt,name=type,enum=ddproto.HallEnumEvent" json:"type,omitempty"`
	Reward           *HallEnum_Reward `protobuf:"varint,3,opt,name=reward,enum=ddproto.HallEnum_Reward" json:"reward,omitempty"`
	RichText         []string         `protobuf:"bytes,5,rep,name=richText" json:"richText,omitempty"`
	Title            *string          `protobuf:"bytes,6,opt,name=title" json:"title,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *HallItemEvent) Reset()                    { *m = HallItemEvent{} }
func (m *HallItemEvent) String() string            { return proto.CompactTextString(m) }
func (*HallItemEvent) ProtoMessage()               {}
func (*HallItemEvent) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{2} }

func (m *HallItemEvent) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *HallItemEvent) GetType() HallEnumEvent {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return HallEnumEvent_TYPE_TIME
}

func (m *HallItemEvent) GetReward() HallEnum_Reward {
	if m != nil && m.Reward != nil {
		return *m.Reward
	}
	return HallEnum_Reward_RE_EXP
}

func (m *HallItemEvent) GetRichText() []string {
	if m != nil {
		return m.RichText
	}
	return nil
}

func (m *HallItemEvent) GetTitle() string {
	if m != nil && m.Title != nil {
		return *m.Title
	}
	return ""
}

// 单个邮件
type HallMailItem struct {
	Id               *int32            `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Type             *HallEnumMailType `protobuf:"varint,2,opt,name=type,enum=ddproto.HallEnumMailType" json:"type,omitempty"`
	Title            *string           `protobuf:"bytes,3,opt,name=title" json:"title,omitempty"`
	Content          *string           `protobuf:"bytes,4,opt,name=content" json:"content,omitempty"`
	IsWatch          *bool             `protobuf:"varint,5,opt,name=isWatch" json:"isWatch,omitempty"`
	IsCheck          *bool             `protobuf:"varint,6,opt,name=isCheck" json:"isCheck,omitempty"`
	Attach           []*HallBagItem    `protobuf:"bytes,7,rep,name=attach" json:"attach,omitempty"`
	Date             *int32            `protobuf:"varint,8,opt,name=date" json:"date,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *HallMailItem) Reset()                    { *m = HallMailItem{} }
func (m *HallMailItem) String() string            { return proto.CompactTextString(m) }
func (*HallMailItem) ProtoMessage()               {}
func (*HallMailItem) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{3} }

func (m *HallMailItem) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *HallMailItem) GetType() HallEnumMailType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return HallEnumMailType_SYSTEM
}

func (m *HallMailItem) GetTitle() string {
	if m != nil && m.Title != nil {
		return *m.Title
	}
	return ""
}

func (m *HallMailItem) GetContent() string {
	if m != nil && m.Content != nil {
		return *m.Content
	}
	return ""
}

func (m *HallMailItem) GetIsWatch() bool {
	if m != nil && m.IsWatch != nil {
		return *m.IsWatch
	}
	return false
}

func (m *HallMailItem) GetIsCheck() bool {
	if m != nil && m.IsCheck != nil {
		return *m.IsCheck
	}
	return false
}

func (m *HallMailItem) GetAttach() []*HallBagItem {
	if m != nil {
		return m.Attach
	}
	return nil
}

func (m *HallMailItem) GetDate() int32 {
	if m != nil && m.Date != nil {
		return *m.Date
	}
	return 0
}

// 背包单个道具
type HallBagItem struct {
	Type             *HallEnumPropsType `protobuf:"varint,1,opt,name=type,enum=ddproto.HallEnumPropsType" json:"type,omitempty"`
	Amount           *int32             `protobuf:"varint,2,opt,name=amount" json:"amount,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *HallBagItem) Reset()                    { *m = HallBagItem{} }
func (m *HallBagItem) String() string            { return proto.CompactTextString(m) }
func (*HallBagItem) ProtoMessage()               {}
func (*HallBagItem) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{4} }

func (m *HallBagItem) GetType() HallEnumPropsType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return HallEnumPropsType_TYPE_LABA
}

func (m *HallBagItem) GetAmount() int32 {
	if m != nil && m.Amount != nil {
		return *m.Amount
	}
	return 0
}

// 单个任务
type HallItemTask struct {
	Id               *int32            `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Type             *HallEnumTaskType `protobuf:"varint,2,opt,name=type,enum=ddproto.HallEnumTaskType" json:"type,omitempty"`
	Title            *string           `protobuf:"bytes,3,opt,name=title" json:"title,omitempty"`
	Reward           *string           `protobuf:"bytes,4,opt,name=reward" json:"reward,omitempty"`
	Num              *int32            `protobuf:"varint,5,opt,name=Num" json:"Num,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *HallItemTask) Reset()                    { *m = HallItemTask{} }
func (m *HallItemTask) String() string            { return proto.CompactTextString(m) }
func (*HallItemTask) ProtoMessage()               {}
func (*HallItemTask) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{5} }

func (m *HallItemTask) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *HallItemTask) GetType() HallEnumTaskType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return HallEnumTaskType_TYPE_MJ
}

func (m *HallItemTask) GetTitle() string {
	if m != nil && m.Title != nil {
		return *m.Title
	}
	return ""
}

func (m *HallItemTask) GetReward() string {
	if m != nil && m.Reward != nil {
		return *m.Reward
	}
	return ""
}

func (m *HallItemTask) GetNum() int32 {
	if m != nil && m.Num != nil {
		return *m.Num
	}
	return 0
}

// 个人信息
type HallUserData struct {
	UserName          *string       `protobuf:"bytes,1,opt,name=userName" json:"userName,omitempty"`
	UserId            *int32        `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	NiceValue         *int64        `protobuf:"varint,3,opt,name=niceValue" json:"niceValue,omitempty"`
	EvilValue         *int64        `protobuf:"varint,4,opt,name=evilValue" json:"evilValue,omitempty"`
	UserLevel         *int32        `protobuf:"varint,5,opt,name=userLevel" json:"userLevel,omitempty"`
	UserVIP           *bool         `protobuf:"varint,6,opt,name=userVIP" json:"userVIP,omitempty"`
	UserVIPLv         *HallUser_VIP `protobuf:"varint,7,opt,name=userVIPLv,enum=ddproto.HallUser_VIP" json:"userVIPLv,omitempty"`
	UserMoney         *int64        `protobuf:"varint,8,opt,name=userMoney" json:"userMoney,omitempty"`
	UserDiamond       *int64        `protobuf:"varint,9,opt,name=userDiamond" json:"userDiamond,omitempty"`
	UserRedBag        *string       `protobuf:"bytes,10,opt,name=userRedBag" json:"userRedBag,omitempty"`
	UserLotteryTicket *int64        `protobuf:"varint,11,opt,name=userLotteryTicket" json:"userLotteryTicket,omitempty"`
	Sex               *bool         `protobuf:"varint,12,opt,name=sex" json:"sex,omitempty"`
	XXX_unrecognized  []byte        `json:"-"`
}

func (m *HallUserData) Reset()                    { *m = HallUserData{} }
func (m *HallUserData) String() string            { return proto.CompactTextString(m) }
func (*HallUserData) ProtoMessage()               {}
func (*HallUserData) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{6} }

func (m *HallUserData) GetUserName() string {
	if m != nil && m.UserName != nil {
		return *m.UserName
	}
	return ""
}

func (m *HallUserData) GetUserId() int32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *HallUserData) GetNiceValue() int64 {
	if m != nil && m.NiceValue != nil {
		return *m.NiceValue
	}
	return 0
}

func (m *HallUserData) GetEvilValue() int64 {
	if m != nil && m.EvilValue != nil {
		return *m.EvilValue
	}
	return 0
}

func (m *HallUserData) GetUserLevel() int32 {
	if m != nil && m.UserLevel != nil {
		return *m.UserLevel
	}
	return 0
}

func (m *HallUserData) GetUserVIP() bool {
	if m != nil && m.UserVIP != nil {
		return *m.UserVIP
	}
	return false
}

func (m *HallUserData) GetUserVIPLv() HallUser_VIP {
	if m != nil && m.UserVIPLv != nil {
		return *m.UserVIPLv
	}
	return HallUser_VIP_LV_1
}

func (m *HallUserData) GetUserMoney() int64 {
	if m != nil && m.UserMoney != nil {
		return *m.UserMoney
	}
	return 0
}

func (m *HallUserData) GetUserDiamond() int64 {
	if m != nil && m.UserDiamond != nil {
		return *m.UserDiamond
	}
	return 0
}

func (m *HallUserData) GetUserRedBag() string {
	if m != nil && m.UserRedBag != nil {
		return *m.UserRedBag
	}
	return ""
}

func (m *HallUserData) GetUserLotteryTicket() int64 {
	if m != nil && m.UserLotteryTicket != nil {
		return *m.UserLotteryTicket
	}
	return 0
}

func (m *HallUserData) GetSex() bool {
	if m != nil && m.Sex != nil {
		return *m.Sex
	}
	return false
}

// 排行信息
type HallItemRank struct {
	Placing          *int32  `protobuf:"varint,1,opt,name=placing" json:"placing,omitempty"`
	UserId           *int32  `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	RankInfo         *string `protobuf:"bytes,3,opt,name=rankInfo" json:"rankInfo,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *HallItemRank) Reset()                    { *m = HallItemRank{} }
func (m *HallItemRank) String() string            { return proto.CompactTextString(m) }
func (*HallItemRank) ProtoMessage()               {}
func (*HallItemRank) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{7} }

func (m *HallItemRank) GetPlacing() int32 {
	if m != nil && m.Placing != nil {
		return *m.Placing
	}
	return 0
}

func (m *HallItemRank) GetUserId() int32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *HallItemRank) GetRankInfo() string {
	if m != nil && m.RankInfo != nil {
		return *m.RankInfo
	}
	return ""
}

// 金币专区
type CoinZone struct {
	Pay              []*GoodsItem `protobuf:"bytes,1,rep,name=pay" json:"pay,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *CoinZone) Reset()                    { *m = CoinZone{} }
func (m *CoinZone) String() string            { return proto.CompactTextString(m) }
func (*CoinZone) ProtoMessage()               {}
func (*CoinZone) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{8} }

func (m *CoinZone) GetPay() []*GoodsItem {
	if m != nil {
		return m.Pay
	}
	return nil
}

// 钻石专区
type DiamondZone struct {
	Pay              []*GoodsItem `protobuf:"bytes,1,rep,name=pay" json:"pay,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *DiamondZone) Reset()                    { *m = DiamondZone{} }
func (m *DiamondZone) String() string            { return proto.CompactTextString(m) }
func (*DiamondZone) ProtoMessage()               {}
func (*DiamondZone) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{9} }

func (m *DiamondZone) GetPay() []*GoodsItem {
	if m != nil {
		return m.Pay
	}
	return nil
}

// 兑换专区
type ExchangeZone struct {
	Money            []*GoodsItem `protobuf:"bytes,1,rep,name=money" json:"money,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *ExchangeZone) Reset()                    { *m = ExchangeZone{} }
func (m *ExchangeZone) String() string            { return proto.CompactTextString(m) }
func (*ExchangeZone) ProtoMessage()               {}
func (*ExchangeZone) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{10} }

func (m *ExchangeZone) GetMoney() []*GoodsItem {
	if m != nil {
		return m.Money
	}
	return nil
}

// 购买专区
type BuyZone struct {
	Pay              []*GoodsItem `protobuf:"bytes,1,rep,name=pay" json:"pay,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *BuyZone) Reset()                    { *m = BuyZone{} }
func (m *BuyZone) String() string            { return proto.CompactTextString(m) }
func (*BuyZone) ProtoMessage()               {}
func (*BuyZone) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{11} }

func (m *BuyZone) GetPay() []*GoodsItem {
	if m != nil {
		return m.Pay
	}
	return nil
}

// 商品类型
type GoodsItem struct {
	Id               *int32  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Money            *int32  `protobuf:"varint,2,opt,name=money" json:"money,omitempty"`
	Img              *string `protobuf:"bytes,3,opt,name=img" json:"img,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *GoodsItem) Reset()                    { *m = GoodsItem{} }
func (m *GoodsItem) String() string            { return proto.CompactTextString(m) }
func (*GoodsItem) ProtoMessage()               {}
func (*GoodsItem) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{12} }

func (m *GoodsItem) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *GoodsItem) GetMoney() int32 {
	if m != nil && m.Money != nil {
		return *m.Money
	}
	return 0
}

func (m *GoodsItem) GetImg() string {
	if m != nil && m.Img != nil {
		return *m.Img
	}
	return ""
}

func init() {
	proto.RegisterType((*HallLotteryItem)(nil), "ddproto.hall_lottery_item")
	proto.RegisterType((*HallSignLotteryInfo)(nil), "ddproto.hall_sign_lottery_info")
	proto.RegisterType((*HallItemEvent)(nil), "ddproto.hall_item_event")
	proto.RegisterType((*HallMailItem)(nil), "ddproto.hall_mail_item")
	proto.RegisterType((*HallBagItem)(nil), "ddproto.hall_bag_item")
	proto.RegisterType((*HallItemTask)(nil), "ddproto.hall_item_task")
	proto.RegisterType((*HallUserData)(nil), "ddproto.hall_userData")
	proto.RegisterType((*HallItemRank)(nil), "ddproto.hall_item_rank")
	proto.RegisterType((*CoinZone)(nil), "ddproto.CoinZone")
	proto.RegisterType((*DiamondZone)(nil), "ddproto.DiamondZone")
	proto.RegisterType((*ExchangeZone)(nil), "ddproto.ExchangeZone")
	proto.RegisterType((*BuyZone)(nil), "ddproto.BuyZone")
	proto.RegisterType((*GoodsItem)(nil), "ddproto.GoodsItem")
	proto.RegisterEnum("ddproto.HallEnumProtoId", HallEnumProtoId_name, HallEnumProtoId_value)
	proto.RegisterEnum("ddproto.HallEnumEvent", HallEnumEvent_name, HallEnumEvent_value)
	proto.RegisterEnum("ddproto.HallEnum_Reward", HallEnum_Reward_name, HallEnum_Reward_value)
	proto.RegisterEnum("ddproto.HallEnumMailType", HallEnumMailType_name, HallEnumMailType_value)
	proto.RegisterEnum("ddproto.HallEnumPropsType", HallEnumPropsType_name, HallEnumPropsType_value)
	proto.RegisterEnum("ddproto.HallEnumTaskType", HallEnumTaskType_name, HallEnumTaskType_value)
	proto.RegisterEnum("ddproto.HallUser_VIP", HallUser_VIP_name, HallUser_VIP_value)
}

var fileDescriptor14 = []byte{
	// 1239 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x56, 0xdd, 0x72, 0xda, 0xc6,
	0x17, 0xff, 0x0b, 0x01, 0x36, 0x07, 0x9b, 0xac, 0x37, 0x0e, 0x91, 0x9d, 0x38, 0xc3, 0x9f, 0xc9,
	0x74, 0x18, 0x9a, 0x71, 0x6c, 0x27, 0x6d, 0x73, 0x93, 0x0b, 0x19, 0xd6, 0x44, 0x31, 0x16, 0x78,
	0x91, 0x71, 0xec, 0xce, 0x54, 0xb3, 0x45, 0x8a, 0xad, 0x31, 0x48, 0x1e, 0x10, 0x6e, 0xfc, 0x10,
	0xbd, 0xe8, 0x5d, 0x1f, 0xa4, 0x2f, 0xd1, 0x57, 0xea, 0x55, 0x67, 0x57, 0x8b, 0x40, 0x06, 0x26,
	0xcd, 0xdd, 0x39, 0xe7, 0x77, 0x3e, 0x7f, 0xe7, 0xe8, 0x03, 0x1e, 0x5d, 0xb3, 0x7e, 0xdf, 0x76,
	0x58, 0xc8, 0x76, 0x6f, 0x87, 0x41, 0x18, 0xe0, 0x15, 0xc7, 0x11, 0x42, 0xf9, 0x27, 0xd8, 0x10,
	0x58, 0x3f, 0x08, 0x43, 0x77, 0x78, 0x6f, 0x7b, 0xa1, 0x3b, 0xc0, 0x05, 0x48, 0x79, 0x8e, 0xa6,
	0x94, 0x94, 0x4a, 0x86, 0xa6, 0x3c, 0x07, 0x63, 0x48, 0xfb, 0x6c, 0xe0, 0x6a, 0xa9, 0x92, 0x52,
	0xc9, 0x51, 0x21, 0x97, 0xff, 0x54, 0xa0, 0x28, 0x22, 0x47, 0xde, 0x95, 0x3f, 0x0d, 0xf7, 0x3f,
	0x07, 0xf8, 0x1d, 0xe4, 0xb8, 0xd1, 0x08, 0xdd, 0xc1, 0x48, 0x53, 0x4a, 0x6a, 0x25, 0x7f, 0xb0,
	0xbd, 0x2b, 0x0b, 0xee, 0xce, 0x55, 0xa3, 0x53, 0x67, 0xfc, 0x1c, 0x72, 0x61, 0x10, 0xb2, 0x7e,
	0x9d, 0xdd, 0x8f, 0x44, 0xb5, 0x0c, 0x9d, 0x1a, 0xf0, 0x77, 0x50, 0xe8, 0x05, 0x7e, 0xe8, 0xf9,
	0xe3, 0x60, 0x3c, 0x12, 0x2e, 0xaa, 0x70, 0x79, 0x60, 0x2d, 0xff, 0xa5, 0xc8, 0x81, 0x79, 0x7a,
	0xdb, 0xbd, 0x73, 0xfd, 0x70, 0x6e, 0xa4, 0x57, 0x90, 0x0e, 0xef, 0x6f, 0xa3, 0x91, 0x0a, 0x07,
	0x5a, 0xb2, 0x3d, 0xd7, 0x1f, 0xcb, 0x38, 0x2a, 0xbc, 0xf0, 0x3e, 0x64, 0x87, 0xee, 0x6f, 0x6c,
	0xe8, 0x88, 0x8a, 0x85, 0x83, 0xad, 0x05, 0xfe, 0x54, 0x38, 0x50, 0xe9, 0x88, 0xb7, 0x61, 0x75,
	0xe8, 0xf5, 0xae, 0x2d, 0xf7, 0x4b, 0xa8, 0x65, 0x4a, 0x6a, 0x25, 0x47, 0x63, 0x1d, 0x6f, 0x42,
	0x26, 0xf4, 0xc2, 0xbe, 0xab, 0x65, 0x05, 0xa1, 0x91, 0x52, 0xfe, 0x47, 0x81, 0x82, 0x48, 0x37,
	0x60, 0x5e, 0x7f, 0xf1, 0x22, 0xf6, 0x12, 0x5d, 0x3f, 0x5f, 0xd0, 0x85, 0x88, 0xe5, 0x3e, 0xb2,
	0xf3, 0xb8, 0x94, 0x3a, 0x53, 0x0a, 0x6b, 0xb0, 0xc2, 0x39, 0x73, 0xfd, 0x50, 0x4b, 0x0b, 0xfb,
	0x44, 0xe5, 0x88, 0x37, 0x3a, 0x67, 0x61, 0xef, 0x5a, 0xcb, 0x94, 0x94, 0xca, 0x2a, 0x9d, 0xa8,
	0x11, 0x52, 0xbb, 0x76, 0x7b, 0x37, 0xa2, 0x6d, 0x81, 0x08, 0x15, 0xef, 0x42, 0x96, 0x85, 0x21,
	0xeb, 0x5d, 0x6b, 0x2b, 0x62, 0xd9, 0xc5, 0x64, 0x5f, 0xbf, 0xb2, 0xab, 0x68, 0xd1, 0xd2, 0x8b,
	0x9f, 0x93, 0xc3, 0x42, 0x57, 0x5b, 0x15, 0x73, 0x09, 0xb9, 0x7c, 0x09, 0xeb, 0x09, 0x67, 0xbc,
	0x2f, 0x47, 0x55, 0xc4, 0xa8, 0x3b, 0x0b, 0x46, 0xbd, 0x1d, 0x06, 0xb7, 0xa3, 0xd9, 0x59, 0x8b,
	0x90, 0x65, 0x83, 0x60, 0xec, 0x87, 0xf2, 0x74, 0xa4, 0x56, 0xfe, 0x63, 0x42, 0xac, 0xb8, 0x87,
	0x90, 0x8d, 0x6e, 0xe6, 0x88, 0x7d, 0x9d, 0x20, 0xf6, 0xd9, 0x82, 0x6a, 0x3c, 0xcc, 0xfa, 0x1a,
	0xaf, 0xc5, 0xf8, 0x4e, 0x22, 0x5a, 0x27, 0xc7, 0x80, 0x40, 0x35, 0xc7, 0x03, 0xc1, 0x68, 0x86,
	0x72, 0xb1, 0xfc, 0xbb, 0x2a, 0x07, 0x1e, 0x8f, 0xdc, 0x61, 0x9d, 0x85, 0x8c, 0x1f, 0x0c, 0x97,
	0x4d, 0xfe, 0xa0, 0x29, 0x22, 0x3a, 0xd6, 0x79, 0x5e, 0x2e, 0x1b, 0xce, 0x64, 0xb2, 0x48, 0xe3,
	0xcf, 0x8b, 0xef, 0xf5, 0xdc, 0x2e, 0xeb, 0x8f, 0xa3, 0x4e, 0x54, 0x3a, 0x35, 0x70, 0xd4, 0xbd,
	0xf3, 0xfa, 0x11, 0x9a, 0x8e, 0xd0, 0xd8, 0xc0, 0x51, 0x9e, 0xa5, 0xe9, 0xde, 0xb9, 0x7d, 0xd9,
	0xd9, 0xd4, 0xc0, 0xb7, 0xcd, 0x95, 0xae, 0xd1, 0x9e, 0x6c, 0x5b, 0xaa, 0xf8, 0x6d, 0x14, 0xd7,
	0x35, 0xda, 0xcd, 0x3b, 0x6d, 0x45, 0xf0, 0xf5, 0x60, 0xe1, 0x1c, 0xb6, 0xbb, 0x46, 0x9b, 0x4e,
	0x1d, 0x27, 0xd5, 0x4e, 0x02, 0xdf, 0xbd, 0x17, 0x8b, 0x57, 0xe9, 0xd4, 0x80, 0x4b, 0x90, 0x17,
	0x3c, 0x78, 0x6c, 0x10, 0xf8, 0x8e, 0x96, 0x13, 0xf8, 0xac, 0x09, 0xbf, 0x00, 0xe0, 0x2a, 0x75,
	0x9d, 0x43, 0x76, 0xa5, 0x81, 0xe0, 0x67, 0xc6, 0x82, 0x5f, 0xc1, 0x86, 0x68, 0x3e, 0x7a, 0xb1,
	0x58, 0x5e, 0xef, 0xc6, 0x0d, 0xb5, 0xbc, 0xc8, 0x33, 0x0f, 0xf0, 0x7d, 0x8c, 0xdc, 0x2f, 0xda,
	0x9a, 0x98, 0x8c, 0x8b, 0xe5, 0x5f, 0x66, 0x4f, 0x64, 0xc8, 0xfc, 0x1b, 0xce, 0xc0, 0x6d, 0x9f,
	0xf5, 0x3c, 0xff, 0x4a, 0xde, 0xc9, 0x44, 0x5d, 0xba, 0x0d, 0xfe, 0xc8, 0x33, 0xff, 0xc6, 0xf0,
	0x3f, 0x07, 0xf2, 0x2c, 0x62, 0xbd, 0xbc, 0x07, 0xab, 0xb5, 0xc0, 0xf3, 0x2f, 0x03, 0xdf, 0xc5,
	0x2f, 0x41, 0xbd, 0x65, 0xf7, 0xf2, 0xcd, 0x88, 0x63, 0xee, 0x1a, 0x41, 0xe0, 0x8c, 0xf8, 0x7b,
	0x90, 0x72, 0xb8, 0xfc, 0x06, 0xf2, 0x72, 0xf8, 0x6f, 0x08, 0x7a, 0x07, 0x6b, 0xe4, 0x4b, 0xef,
	0x9a, 0xf9, 0x57, 0xae, 0x88, 0xaa, 0x40, 0x66, 0x20, 0x28, 0x5f, 0x1e, 0x17, 0x39, 0x94, 0x5f,
	0xc3, 0xca, 0xe1, 0xf8, 0xfe, 0x1b, 0x4a, 0xd5, 0x20, 0x17, 0x5b, 0xe6, 0x9e, 0xa7, 0xcd, 0x49,
	0xdd, 0x88, 0xa1, 0x48, 0xe1, 0xb4, 0x7b, 0x83, 0x2b, 0xc9, 0x0d, 0x17, 0xab, 0x7f, 0x67, 0xe5,
	0xf7, 0x67, 0xf2, 0x44, 0x87, 0x81, 0xe1, 0xe0, 0x22, 0xe0, 0x0f, 0x7a, 0xb3, 0x69, 0xb7, 0x8d,
	0xba, 0xfd, 0x81, 0xe8, 0xd4, 0x3a, 0x24, 0xba, 0x85, 0xfe, 0x87, 0x9f, 0xc2, 0xe3, 0xd8, 0x7e,
	0x7a, 0x66, 0xd4, 0x8e, 0xed, 0x5a, 0xcb, 0x34, 0x91, 0x82, 0x9f, 0xc1, 0xd3, 0x05, 0x80, 0xad,
	0xd7, 0x8e, 0x51, 0x2a, 0x11, 0xd5, 0xd0, 0x4f, 0x88, 0xdd, 0x6c, 0x35, 0x0c, 0x13, 0xa9, 0x89,
	0xa8, 0x29, 0x20, 0xa2, 0xd2, 0x89, 0x1e, 0xce, 0x3a, 0x84, 0xda, 0x75, 0xdd, 0xd2, 0x51, 0x0e,
	0x6f, 0x43, 0x71, 0xde, 0x2e, 0x62, 0x00, 0x6f, 0xc1, 0x93, 0x18, 0xab, 0x53, 0xfd, 0xdc, 0x6e,
	0xb6, 0x2c, 0x8b, 0xd0, 0x0b, 0x94, 0xc7, 0x3b, 0xb0, 0xb5, 0x10, 0x12, 0x91, 0x6b, 0xf8, 0xff,
	0xb0, 0x33, 0x85, 0x3b, 0x31, 0x68, 0x58, 0xe4, 0xa4, 0x23, 0x5c, 0xd6, 0x13, 0xc9, 0x3b, 0x46,
	0xc3, 0x8c, 0x93, 0x17, 0x12, 0xc9, 0x67, 0x21, 0x11, 0xf9, 0x28, 0x31, 0x67, 0xcb, 0x6c, 0x1a,
	0x26, 0x39, 0xd7, 0x69, 0xdd, 0xa6, 0xe4, 0x14, 0xa1, 0x65, 0x20, 0x8f, 0xdc, 0x48, 0x90, 0x40,
	0xba, 0xc4, 0xb4, 0x44, 0x10, 0x5e, 0x60, 0xe7, 0xfe, 0x8f, 0x93, 0x8c, 0xb6, 0x5a, 0x7c, 0x12,
	0xa3, 0x13, 0x05, 0x6d, 0x2e, 0x03, 0x79, 0xe4, 0x13, 0xfc, 0x12, 0x4a, 0xd3, 0x95, 0xb7, 0xac,
	0xb3, 0x76, 0x5d, 0xb7, 0x08, 0x25, 0xa7, 0x5d, 0x42, 0x3b, 0x46, 0xcb, 0x34, 0xcc, 0xa3, 0x16,
	0x2a, 0x2e, 0xf6, 0xd2, 0x6b, 0xc7, 0xb3, 0x5e, 0x4f, 0x13, 0x2b, 0x3a, 0xd4, 0x1b, 0x92, 0x45,
	0xde, 0x84, 0xb6, 0x04, 0xe3, 0x3d, 0x6c, 0xe1, 0x0a, 0xbc, 0x9c, 0xcf, 0xce, 0x0f, 0x43, 0xef,
	0x74, 0x88, 0xd5, 0xe1, 0xd9, 0x45, 0x96, 0xed, 0xff, 0xe4, 0xc9, 0x73, 0x3e, 0x4b, 0xac, 0x46,
	0x6f, 0xb7, 0x9b, 0xc4, 0x6e, 0xeb, 0x17, 0x76, 0xed, 0x50, 0x24, 0x7a, 0xbe, 0x64, 0xa0, 0x64,
	0xb9, 0x9d, 0xaf, 0x7a, 0xf1, 0x52, 0x2f, 0xaa, 0xef, 0xe5, 0x5f, 0xcf, 0xf4, 0xef, 0x05, 0xaf,
	0x43, 0xce, 0xba, 0x68, 0x13, 0xdb, 0x32, 0x4e, 0x08, 0x52, 0xf0, 0x1a, 0xac, 0x0a, 0xd5, 0x24,
	0xe7, 0x28, 0x15, 0x83, 0xe6, 0x59, 0xb3, 0x89, 0xd4, 0xea, 0xf7, 0x80, 0x1e, 0xfe, 0xcc, 0x60,
	0x80, 0x2c, 0x25, 0x36, 0xf9, 0xd4, 0x46, 0x0a, 0xce, 0xc3, 0x0a, 0x25, 0x76, 0xc3, 0x38, 0xb2,
	0x50, 0xaa, 0xba, 0x0f, 0x8f, 0x17, 0xfc, 0x73, 0x70, 0xff, 0xce, 0x45, 0xc7, 0x22, 0x27, 0x48,
	0xc1, 0x05, 0x80, 0x23, 0x6a, 0x10, 0xb3, 0x6e, 0xeb, 0xf5, 0x3a, 0x4a, 0x55, 0x7f, 0x86, 0xcd,
	0x45, 0xdf, 0xee, 0xb8, 0x8d, 0xa6, 0x7e, 0xa8, 0x23, 0x05, 0x23, 0x58, 0x13, 0xea, 0x47, 0xa3,
	0xad, 0x1b, 0xa7, 0x06, 0x4a, 0xe1, 0x47, 0x90, 0x17, 0x96, 0x23, 0xdd, 0x6c, 0x1c, 0xeb, 0x48,
	0xc5, 0x1b, 0xb0, 0x2e, 0x0c, 0xe4, 0x53, 0xdb, 0x7e, 0xb3, 0xb7, 0xb7, 0x87, 0xd2, 0xd5, 0xf7,
	0x80, 0xe7, 0x3f, 0xd5, 0xbc, 0x65, 0xe1, 0x78, 0xf2, 0x71, 0x66, 0xf8, 0x7a, 0xfd, 0x12, 0xa5,
	0x62, 0xed, 0xf2, 0xe3, 0x07, 0xa4, 0x56, 0x8f, 0x67, 0x3e, 0xc6, 0xfc, 0xcb, 0x85, 0x57, 0x21,
	0xdd, 0xec, 0xda, 0xfb, 0x48, 0x91, 0xd2, 0x01, 0x4a, 0x49, 0xe9, 0x0d, 0x52, 0xa5, 0xf4, 0x16,
	0xa5, 0xa5, 0xf4, 0x03, 0xca, 0x48, 0xe9, 0x47, 0x94, 0xfd, 0x37, 0x00, 0x00, 0xff, 0xff, 0x0f,
	0x1b, 0x1e, 0xe0, 0x6d, 0x0b, 0x00, 0x00,
}
