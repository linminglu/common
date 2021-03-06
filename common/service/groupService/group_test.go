package groupService

import (
	"casino_common/proto/ddproto"
	"casino_common/test"
	"github.com/golang/protobuf/proto"
	"testing"
)

func init() {
	test.TestInit()
}

func TestGroupInfo_AddMember(t *testing.T) {
	group := GetGroupInfoById(123)
	t.Log(group)
}

func TestGroupInfo_SaveToMongo(t *testing.T) {
	group := GroupInfo{
		GroupInfo: ddproto.GroupInfo{
			Id:    proto.Int32(123),
			Name:  proto.String("麻将测试群"),
			Info:  proto.String("测试群组开房。。"),
			Owner: proto.Uint32(10010),
			Members: []*ddproto.GroupMemberInfo{
				&ddproto.GroupMemberInfo{
					Uid:      proto.Uint32(10010),
					NickName: proto.String("群主大人"),
					Remark:   proto.String(""),
					HeadImg:  proto.String("http://baidu.com"),
				},
			},
			GameOpts: []*ddproto.GameOption{
				&ddproto.GameOption{
					Id:     proto.Int32(0),
					GameId: proto.Int32(17),
					Remark: proto.String("经典斗牛"),
					Option: proto.String("经典牛牛 3人 10场"),
				},
			},
			SyncId: proto.Int32(1),
		},
	}

	group.SaveToMongo()
}

//获取群列表
func TestGetUserGroups(t *testing.T) {
	list, err := GetUserGroups(10010)
	t.Log(list, err)
}
