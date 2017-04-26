// Code generated by protoc-gen-go.
// source: rpc_hall.proto
// DO NOT EDIT!

package ddproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// 用于更新hall的配置文件
// 更新配置文件
type RpcHallUpdateConfig struct {
	ConfigId         *int32 `protobuf:"varint,1,opt,name=configId" json:"configId,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *RpcHallUpdateConfig) Reset()                    { *m = RpcHallUpdateConfig{} }
func (m *RpcHallUpdateConfig) String() string            { return proto.CompactTextString(m) }
func (*RpcHallUpdateConfig) ProtoMessage()               {}
func (*RpcHallUpdateConfig) Descriptor() ([]byte, []int) { return fileDescriptor39, []int{0} }

func (m *RpcHallUpdateConfig) GetConfigId() int32 {
	if m != nil && m.ConfigId != nil {
		return *m.ConfigId
	}
	return 0
}

func init() {
	proto.RegisterType((*RpcHallUpdateConfig)(nil), "ddproto.rpc_hall_update_config")
}

var fileDescriptor39 = []byte{
	// 79 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x2b, 0x2a, 0x48, 0x8e,
	0xcf, 0x48, 0xcc, 0xc9, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4f, 0x49, 0x01, 0x33,
	0x94, 0xb4, 0xb8, 0xc4, 0x60, 0x52, 0xf1, 0xa5, 0x05, 0x29, 0x89, 0x25, 0xa9, 0xf1, 0xc9, 0xf9,
	0x79, 0x69, 0x99, 0xe9, 0x42, 0x02, 0x5c, 0x1c, 0x10, 0x96, 0x67, 0x8a, 0x04, 0xa3, 0x02, 0xa3,
	0x06, 0x2b, 0x20, 0x00, 0x00, 0xff, 0xff, 0xc4, 0x17, 0xee, 0xce, 0x45, 0x00, 0x00, 0x00,
}