// Code generated by protoc-gen-go.
// source: hall_common_server_model.proto
// DO NOT EDIT!

package yjprotogo

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// 用户的资
type User struct {
	RoomCard            *int64   `protobuf:"varint,1,opt,name=RoomCard" json:"RoomCard,omitempty"`
	Pwd                 *string  `protobuf:"bytes,2,opt,name=pwd" json:"pwd,omitempty"`
	Coin                *int64   `protobuf:"varint,3,opt,name=coin" json:"coin,omitempty"`
	Id                  *uint32  `protobuf:"varint,4,opt,name=id" json:"id,omitempty"`
	NickName            *string  `protobuf:"bytes,5,opt,name=nickName" json:"nickName,omitempty"`
	Scores              *int32   `protobuf:"varint,6,opt,name=scores" json:"scores,omitempty"`
	LastDrawLotteryTime *string  `protobuf:"bytes,7,opt,name=lastDrawLotteryTime" json:"lastDrawLotteryTime,omitempty"`
	LastSignTime        *string  `protobuf:"bytes,8,opt,name=lastSignTime" json:"lastSignTime,omitempty"`
	SignTotalDays       *int32   `protobuf:"varint,9,opt,name=signTotalDays" json:"signTotalDays,omitempty"`
	SignContinuousDays  *int32   `protobuf:"varint,10,opt,name=signContinuousDays" json:"signContinuousDays,omitempty"`
	Diamond             *int64   `protobuf:"varint,11,opt,name=Diamond" json:"Diamond,omitempty"`
	Diamond2            *int64   `protobuf:"varint,12,opt,name=Diamond2" json:"Diamond2,omitempty"`
	OpenId              *string  `protobuf:"bytes,13,opt,name=openId" json:"openId,omitempty"`
	UnionId             *string  `protobuf:"bytes,14,opt,name=UnionId" json:"UnionId,omitempty"`
	HeadUrl             *string  `protobuf:"bytes,15,opt,name=headUrl" json:"headUrl,omitempty"`
	City                *string  `protobuf:"bytes,16,opt,name=city" json:"city,omitempty"`
	Sex                 *int32   `protobuf:"varint,17,opt,name=sex" json:"sex,omitempty"`
	RobotType           *int32   `protobuf:"varint,18,opt,name=robotType" json:"robotType,omitempty"`
	Ticket              *int32   `protobuf:"varint,19,opt,name=ticket" json:"ticket,omitempty"`
	Bonus               *float64 `protobuf:"fixed64,20,opt,name=bonus" json:"bonus,omitempty"`
	RegTime             *int64   `protobuf:"varint,21,opt,name=regTime" json:"regTime,omitempty"`
	RegChannel          *string  `protobuf:"bytes,22,opt,name=regChannel" json:"regChannel,omitempty"`
	AgentId             *uint32  `protobuf:"varint,23,opt,name=agentId" json:"agentId,omitempty"`
	LastIp              *string  `protobuf:"bytes,24,opt,name=lastIp" json:"lastIp,omitempty"`
	LastTime            *int64   `protobuf:"varint,25,opt,name=lastTime" json:"lastTime,omitempty"`
	// 用户兑换信息
	RealName         *string  `protobuf:"bytes,26,opt,name=realName" json:"realName,omitempty"`
	PhoneNumber      *string  `protobuf:"bytes,27,opt,name=phoneNumber" json:"phoneNumber,omitempty"`
	WxNumber         *string  `protobuf:"bytes,28,opt,name=wxNumber" json:"wxNumber,omitempty"`
	RealAddress      *string  `protobuf:"bytes,29,opt,name=realAddress" json:"realAddress,omitempty"`
	ChannelId        *int32   `protobuf:"varint,30,opt,name=channelId" json:"channelId,omitempty"`
	NewUserAward     *bool    `protobuf:"varint,31,opt,name=newUserAward" json:"newUserAward,omitempty"`
	RegIp            *string  `protobuf:"bytes,32,opt,name=regIp" json:"regIp,omitempty"`
	Longitude        *float32 `protobuf:"fixed32,33,opt,name=longitude" json:"longitude,omitempty"`
	Latitude         *float32 `protobuf:"fixed32,34,opt,name=latitude" json:"latitude,omitempty"`
	Location         *string  `protobuf:"bytes,35,opt,name=location" json:"location,omitempty"`
	BagPassword      *string  `protobuf:"bytes,36,opt,name=bagPassword" json:"bagPassword,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

func (m *User) GetRoomCard() int64 {
	if m != nil && m.RoomCard != nil {
		return *m.RoomCard
	}
	return 0
}

func (m *User) GetPwd() string {
	if m != nil && m.Pwd != nil {
		return *m.Pwd
	}
	return ""
}

func (m *User) GetCoin() int64 {
	if m != nil && m.Coin != nil {
		return *m.Coin
	}
	return 0
}

func (m *User) GetId() uint32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *User) GetNickName() string {
	if m != nil && m.NickName != nil {
		return *m.NickName
	}
	return ""
}

func (m *User) GetScores() int32 {
	if m != nil && m.Scores != nil {
		return *m.Scores
	}
	return 0
}

func (m *User) GetLastDrawLotteryTime() string {
	if m != nil && m.LastDrawLotteryTime != nil {
		return *m.LastDrawLotteryTime
	}
	return ""
}

func (m *User) GetLastSignTime() string {
	if m != nil && m.LastSignTime != nil {
		return *m.LastSignTime
	}
	return ""
}

func (m *User) GetSignTotalDays() int32 {
	if m != nil && m.SignTotalDays != nil {
		return *m.SignTotalDays
	}
	return 0
}

func (m *User) GetSignContinuousDays() int32 {
	if m != nil && m.SignContinuousDays != nil {
		return *m.SignContinuousDays
	}
	return 0
}

func (m *User) GetDiamond() int64 {
	if m != nil && m.Diamond != nil {
		return *m.Diamond
	}
	return 0
}

func (m *User) GetDiamond2() int64 {
	if m != nil && m.Diamond2 != nil {
		return *m.Diamond2
	}
	return 0
}

func (m *User) GetOpenId() string {
	if m != nil && m.OpenId != nil {
		return *m.OpenId
	}
	return ""
}

func (m *User) GetUnionId() string {
	if m != nil && m.UnionId != nil {
		return *m.UnionId
	}
	return ""
}

func (m *User) GetHeadUrl() string {
	if m != nil && m.HeadUrl != nil {
		return *m.HeadUrl
	}
	return ""
}

func (m *User) GetCity() string {
	if m != nil && m.City != nil {
		return *m.City
	}
	return ""
}

func (m *User) GetSex() int32 {
	if m != nil && m.Sex != nil {
		return *m.Sex
	}
	return 0
}

func (m *User) GetRobotType() int32 {
	if m != nil && m.RobotType != nil {
		return *m.RobotType
	}
	return 0
}

func (m *User) GetTicket() int32 {
	if m != nil && m.Ticket != nil {
		return *m.Ticket
	}
	return 0
}

func (m *User) GetBonus() float64 {
	if m != nil && m.Bonus != nil {
		return *m.Bonus
	}
	return 0
}

func (m *User) GetRegTime() int64 {
	if m != nil && m.RegTime != nil {
		return *m.RegTime
	}
	return 0
}

func (m *User) GetRegChannel() string {
	if m != nil && m.RegChannel != nil {
		return *m.RegChannel
	}
	return ""
}

func (m *User) GetAgentId() uint32 {
	if m != nil && m.AgentId != nil {
		return *m.AgentId
	}
	return 0
}

func (m *User) GetLastIp() string {
	if m != nil && m.LastIp != nil {
		return *m.LastIp
	}
	return ""
}

func (m *User) GetLastTime() int64 {
	if m != nil && m.LastTime != nil {
		return *m.LastTime
	}
	return 0
}

func (m *User) GetRealName() string {
	if m != nil && m.RealName != nil {
		return *m.RealName
	}
	return ""
}

func (m *User) GetPhoneNumber() string {
	if m != nil && m.PhoneNumber != nil {
		return *m.PhoneNumber
	}
	return ""
}

func (m *User) GetWxNumber() string {
	if m != nil && m.WxNumber != nil {
		return *m.WxNumber
	}
	return ""
}

func (m *User) GetRealAddress() string {
	if m != nil && m.RealAddress != nil {
		return *m.RealAddress
	}
	return ""
}

func (m *User) GetChannelId() int32 {
	if m != nil && m.ChannelId != nil {
		return *m.ChannelId
	}
	return 0
}

func (m *User) GetNewUserAward() bool {
	if m != nil && m.NewUserAward != nil {
		return *m.NewUserAward
	}
	return false
}

func (m *User) GetRegIp() string {
	if m != nil && m.RegIp != nil {
		return *m.RegIp
	}
	return ""
}

func (m *User) GetLongitude() float32 {
	if m != nil && m.Longitude != nil {
		return *m.Longitude
	}
	return 0
}

func (m *User) GetLatitude() float32 {
	if m != nil && m.Latitude != nil {
		return *m.Latitude
	}
	return 0
}

func (m *User) GetLocation() string {
	if m != nil && m.Location != nil {
		return *m.Location
	}
	return ""
}

func (m *User) GetBagPassword() string {
	if m != nil && m.BagPassword != nil {
		return *m.BagPassword
	}
	return ""
}

// 通知公告
type TNotice struct {
	Id               *int32   `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	NoticeType       *int32   `protobuf:"varint,2,opt,name=noticeType" json:"noticeType,omitempty"`
	NoticeTitle      *string  `protobuf:"bytes,3,opt,name=noticeTitle" json:"noticeTitle,omitempty"`
	NoticeContent    *string  `protobuf:"bytes,4,opt,name=noticeContent" json:"noticeContent,omitempty"`
	NoticeMemo       *string  `protobuf:"bytes,5,opt,name=noticeMemo" json:"noticeMemo,omitempty"`
	Noticefileds     []string `protobuf:"bytes,6,rep,name=noticefileds" json:"noticefileds,omitempty"`
	ChannelId        *string  `protobuf:"bytes,7,opt,name=channelId" json:"channelId,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *TNotice) Reset()                    { *m = TNotice{} }
func (m *TNotice) String() string            { return proto.CompactTextString(m) }
func (*TNotice) ProtoMessage()               {}
func (*TNotice) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{1} }

func (m *TNotice) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *TNotice) GetNoticeType() int32 {
	if m != nil && m.NoticeType != nil {
		return *m.NoticeType
	}
	return 0
}

func (m *TNotice) GetNoticeTitle() string {
	if m != nil && m.NoticeTitle != nil {
		return *m.NoticeTitle
	}
	return ""
}

func (m *TNotice) GetNoticeContent() string {
	if m != nil && m.NoticeContent != nil {
		return *m.NoticeContent
	}
	return ""
}

func (m *TNotice) GetNoticeMemo() string {
	if m != nil && m.NoticeMemo != nil {
		return *m.NoticeMemo
	}
	return ""
}

func (m *TNotice) GetNoticefileds() []string {
	if m != nil {
		return m.Noticefileds
	}
	return nil
}

func (m *TNotice) GetChannelId() string {
	if m != nil && m.ChannelId != nil {
		return *m.ChannelId
	}
	return ""
}

// 玩家游戏场次统计
type TGameCounts struct {
	Id               *int32  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	UserId           *uint32 `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	DdzCount         *int32  `protobuf:"varint,3,opt,name=ddzCount" json:"ddzCount,omitempty"`
	DdzWinCount      *int32  `protobuf:"varint,4,opt,name=ddzWinCount" json:"ddzWinCount,omitempty"`
	MjCount          *int32  `protobuf:"varint,5,opt,name=mjCount" json:"mjCount,omitempty"`
	MjWinCount       *int32  `protobuf:"varint,6,opt,name=mjWinCount" json:"mjWinCount,omitempty"`
	ThCount          *int32  `protobuf:"varint,7,opt,name=thCount" json:"thCount,omitempty"`
	ThWinWinCount    *int32  `protobuf:"varint,8,opt,name=thWinWinCount" json:"thWinWinCount,omitempty"`
	ZjhCount         *int32  `protobuf:"varint,9,opt,name=zjhCount" json:"zjhCount,omitempty"`
	ZjhWinCount      *int32  `protobuf:"varint,10,opt,name=zjhWinCount" json:"zjhWinCount,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *TGameCounts) Reset()                    { *m = TGameCounts{} }
func (m *TGameCounts) String() string            { return proto.CompactTextString(m) }
func (*TGameCounts) ProtoMessage()               {}
func (*TGameCounts) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{2} }

func (m *TGameCounts) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *TGameCounts) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *TGameCounts) GetDdzCount() int32 {
	if m != nil && m.DdzCount != nil {
		return *m.DdzCount
	}
	return 0
}

func (m *TGameCounts) GetDdzWinCount() int32 {
	if m != nil && m.DdzWinCount != nil {
		return *m.DdzWinCount
	}
	return 0
}

func (m *TGameCounts) GetMjCount() int32 {
	if m != nil && m.MjCount != nil {
		return *m.MjCount
	}
	return 0
}

func (m *TGameCounts) GetMjWinCount() int32 {
	if m != nil && m.MjWinCount != nil {
		return *m.MjWinCount
	}
	return 0
}

func (m *TGameCounts) GetThCount() int32 {
	if m != nil && m.ThCount != nil {
		return *m.ThCount
	}
	return 0
}

func (m *TGameCounts) GetThWinWinCount() int32 {
	if m != nil && m.ThWinWinCount != nil {
		return *m.ThWinWinCount
	}
	return 0
}

func (m *TGameCounts) GetZjhCount() int32 {
	if m != nil && m.ZjhCount != nil {
		return *m.ZjhCount
	}
	return 0
}

func (m *TGameCounts) GetZjhWinCount() int32 {
	if m != nil && m.ZjhWinCount != nil {
		return *m.ZjhWinCount
	}
	return 0
}

// 玩家的任务
type TUserTask struct {
	Id               *int32  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	UserId           *uint32 `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	TaskId           *int32  `protobuf:"varint,3,opt,name=taskId" json:"taskId,omitempty"`
	AwardId          *int32  `protobuf:"varint,4,opt,name=awardId" json:"awardId,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *TUserTask) Reset()                    { *m = TUserTask{} }
func (m *TUserTask) String() string            { return proto.CompactTextString(m) }
func (*TUserTask) ProtoMessage()               {}
func (*TUserTask) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{3} }

func (m *TUserTask) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *TUserTask) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *TUserTask) GetTaskId() int32 {
	if m != nil && m.TaskId != nil {
		return *m.TaskId
	}
	return 0
}

func (m *TUserTask) GetAwardId() int32 {
	if m != nil && m.AwardId != nil {
		return *m.AwardId
	}
	return 0
}

// 推送协议
type Push struct {
	Id               *uint32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Data             []byte  `protobuf:"bytes,2,opt,name=data" json:"data,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Push) Reset()                    { *m = Push{} }
func (m *Push) String() string            { return proto.CompactTextString(m) }
func (*Push) ProtoMessage()               {}
func (*Push) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{4} }

func (m *Push) GetId() uint32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *Push) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*User)(nil), "yjprotogo.User")
	proto.RegisterType((*TNotice)(nil), "yjprotogo.TNotice")
	proto.RegisterType((*TGameCounts)(nil), "yjprotogo.TGameCounts")
	proto.RegisterType((*TUserTask)(nil), "yjprotogo.TUserTask")
	proto.RegisterType((*Push)(nil), "yjprotogo.Push")
}

var fileDescriptor5 = []byte{
	// 649 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x93, 0xcb, 0x6e, 0x13, 0x31,
	0x14, 0x86, 0x95, 0x5b, 0x93, 0x71, 0x92, 0x5e, 0xa6, 0x2d, 0x98, 0x16, 0x4a, 0x09, 0x2c, 0x58,
	0xb1, 0xe0, 0x0d, 0xaa, 0x54, 0x82, 0x48, 0x50, 0x55, 0x90, 0x8a, 0x65, 0xe5, 0xc4, 0x26, 0xe3,
	0x74, 0xc6, 0x27, 0xb2, 0x3d, 0xa4, 0xe9, 0x63, 0xc0, 0x2b, 0xf1, 0x60, 0x1c, 0x9f, 0x19, 0x87,
	0xcb, 0x82, 0xdd, 0x9c, 0x6f, 0x7e, 0x1f, 0xff, 0xe7, 0x62, 0x76, 0x96, 0x89, 0x3c, 0xbf, 0x9d,
	0x43, 0x51, 0x80, 0xb9, 0x75, 0xca, 0x7e, 0x53, 0xf6, 0xb6, 0x00, 0xa9, 0xf2, 0x37, 0x2b, 0x0b,
	0x1e, 0xd2, 0x64, 0xb3, 0xa4, 0x8f, 0x05, 0x8c, 0x7e, 0x74, 0x58, 0xfb, 0x06, 0x25, 0xe9, 0x3e,
	0xeb, 0x7d, 0x02, 0x28, 0xc6, 0xc2, 0x4a, 0xde, 0x38, 0x6f, 0xbc, 0x6e, 0xa5, 0x7d, 0xd6, 0x5a,
	0xad, 0x25, 0x6f, 0x62, 0x90, 0xa4, 0x03, 0xd6, 0x9e, 0x83, 0x36, 0xbc, 0x45, 0xbf, 0x18, 0x6b,
	0x6a, 0xc9, 0xdb, 0xf8, 0x3d, 0x0c, 0x07, 0x8d, 0x9e, 0xdf, 0x5d, 0x89, 0x42, 0xf1, 0x0e, 0x69,
	0x77, 0xd9, 0x8e, 0x9b, 0x83, 0x55, 0x8e, 0xef, 0x60, 0xdc, 0x49, 0x4f, 0xd9, 0x61, 0x2e, 0x9c,
	0xbf, 0xb4, 0x62, 0xfd, 0x01, 0xbc, 0x57, 0x76, 0x33, 0xd5, 0x28, 0xee, 0x92, 0xf8, 0x88, 0x0d,
	0xc2, 0xcf, 0xcf, 0x7a, 0x61, 0x88, 0xf6, 0x88, 0x1e, 0xb3, 0xa1, 0x0b, 0x04, 0xbc, 0xc8, 0x2f,
	0xc5, 0xc6, 0xf1, 0x84, 0x32, 0x9d, 0xb0, 0x34, 0xe0, 0x31, 0x18, 0xaf, 0x4d, 0x09, 0xa5, 0xa3,
	0x7f, 0x8c, 0xfe, 0xed, 0xb1, 0xee, 0xa5, 0x16, 0x58, 0xb1, 0xe4, 0x7d, 0x32, 0x89, 0xc6, 0x6a,
	0xf0, 0x96, 0x0f, 0x88, 0xa0, 0x31, 0x58, 0x29, 0x33, 0x91, 0x7c, 0x48, 0xb7, 0xe0, 0x91, 0x1b,
	0xa3, 0x21, 0x80, 0xdd, 0x08, 0x32, 0x25, 0xe4, 0x8d, 0xcd, 0xf9, 0xde, 0xb6, 0x6c, 0xed, 0x37,
	0x7c, 0x9f, 0x22, 0xec, 0x88, 0x53, 0xf7, 0xfc, 0x80, 0xee, 0x3b, 0x60, 0x89, 0x85, 0x19, 0xf8,
	0xe9, 0x66, 0xa5, 0x78, 0x4a, 0x08, 0xf3, 0x7b, 0x6c, 0x85, 0xf2, 0xfc, 0x90, 0xe2, 0x21, 0xeb,
	0xcc, 0xc0, 0x94, 0x8e, 0x1f, 0x61, 0xd8, 0x08, 0xd9, 0xad, 0x5a, 0x50, 0x95, 0xc7, 0xe4, 0x07,
	0xfb, 0x88, 0x60, 0x9c, 0x09, 0x63, 0x54, 0xce, 0x1f, 0x45, 0x0b, 0x62, 0xa1, 0x8c, 0x47, 0x4f,
	0x8f, 0xa9, 0xbf, 0x98, 0x34, 0x34, 0x68, 0xb2, 0xe2, 0x9c, 0x04, 0x58, 0x56, 0x88, 0x29, 0xcd,
	0x93, 0x58, 0xa8, 0x55, 0x22, 0xa7, 0x09, 0x9c, 0x90, 0xe6, 0x90, 0xf5, 0x57, 0x19, 0x18, 0x75,
	0x55, 0x16, 0x33, 0x65, 0xf9, 0x69, 0x3c, 0xb8, 0xbe, 0xaf, 0xc9, 0xd3, 0x28, 0x0b, 0x07, 0x2f,
	0xa4, 0xc4, 0x61, 0x39, 0xfe, 0x8c, 0x20, 0xd6, 0x35, 0xaf, 0x1c, 0xa1, 0x85, 0x33, 0xaa, 0x03,
	0x67, 0x64, 0xd4, 0x3a, 0xac, 0xc9, 0xc5, 0x3a, 0xec, 0xc7, 0x73, 0xa4, 0xbd, 0x50, 0x1d, 0xba,
	0x47, 0x5f, 0xe7, 0xf1, 0x5c, 0x0e, 0x66, 0xa1, 0x7d, 0x29, 0x15, 0x7f, 0x81, 0xa8, 0x59, 0x59,
	0xf5, 0x15, 0x19, 0x6d, 0x09, 0xcc, 0x91, 0x81, 0xe1, 0x2f, 0xa3, 0x87, 0x99, 0x58, 0x5c, 0x0b,
	0xe7, 0xd6, 0x80, 0xa9, 0x5f, 0x05, 0x38, 0xfa, 0xde, 0x60, 0xdd, 0xe9, 0x15, 0x60, 0x33, 0x55,
	0xbd, 0x6b, 0x0d, 0x32, 0x82, 0x81, 0x21, 0x4a, 0x4d, 0x6f, 0x12, 0xc3, 0x04, 0x35, 0xd3, 0x3e,
	0x57, 0xb4, 0xa0, 0xb4, 0x3f, 0x15, 0x0c, 0xab, 0x82, 0xdd, 0xa4, 0x5d, 0x4d, 0x7e, 0x9f, 0xff,
	0xa8, 0x0a, 0xa8, 0xb7, 0x35, 0x14, 0x47, 0xec, 0xab, 0xce, 0x95, 0x0c, 0x3b, 0xdb, 0xfa, 0xb7,
	0x0b, 0xb4, 0xa9, 0xa3, 0x9f, 0x0d, 0xd6, 0x9f, 0xbe, 0xc3, 0x26, 0x8f, 0xa1, 0x34, 0xde, 0xfd,
	0x65, 0x0c, 0x87, 0x54, 0x62, 0x7b, 0x26, 0xd5, 0x73, 0xa1, 0x47, 0x21, 0xe5, 0x03, 0x09, 0xc9,
	0x11, 0xd9, 0x44, 0xf2, 0x45, 0x9b, 0x0a, 0xb6, 0xe3, 0xce, 0x16, 0xcb, 0x0a, 0x74, 0x62, 0x81,
	0xc5, 0x72, 0x2b, 0xda, 0x89, 0x22, 0x9f, 0x55, 0xa0, 0x4b, 0x00, 0x8b, 0xf3, 0x19, 0x8a, 0xb6,
	0xba, 0x1e, 0x61, 0xbc, 0xf3, 0x61, 0x59, 0x0b, 0x93, 0x78, 0x27, 0x92, 0xad, 0x8c, 0xde, 0xc9,
	0xe8, 0x3d, 0x4b, 0xa6, 0x61, 0x94, 0x53, 0xe1, 0xee, 0xfe, 0x5b, 0x43, 0xd8, 0x66, 0xd4, 0x60,
	0xdc, 0x8a, 0x3e, 0x44, 0x18, 0xff, 0xa4, 0x7a, 0xf9, 0x9d, 0xd1, 0x39, 0x6b, 0x5f, 0x97, 0x2e,
	0xfb, 0x23, 0xc9, 0x30, 0x3c, 0x18, 0x29, 0xbc, 0xa0, 0x14, 0x83, 0x5f, 0x01, 0x00, 0x00, 0xff,
	0xff, 0xf0, 0x07, 0x9a, 0xeb, 0x89, 0x04, 0x00, 0x00,
}