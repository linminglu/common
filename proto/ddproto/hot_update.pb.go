// Code generated by protoc-gen-go.
// source: hot_update.proto
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

// Ignoring public import of common_ack_reg from common_client.proto

// Ignoring public import of common_req_gameLogin from common_client.proto

// Ignoring public import of common_ack_gameLogin from common_client.proto

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

// Ignoring public import of common_enum_reg from common_client.proto

// Ignoring public import of common_enum_os_type from common_client.proto

// Ignoring public import of common_enum_pokerColor from common_client.proto

// Ignoring public import of common_enum_game from common_enum.proto

// Ignoring public import of COMMON_ENUM_ROOMTYPE from common_enum.proto

// Ignoring public import of COMMON_ENUM_GAMESTATUS from common_enum.proto

// Ignoring public import of COMMON_ENUM_RELEASETAG from common_enum.proto

// Ignoring public import of COMMON_ENUM_KICKOUT from common_enum.proto

// Ignoring public import of COMMON_ENUM_APPLYDISSOLVE from common_enum.proto

// Ignoring public import of BTN_TYPE from common_enum.proto

// Ignoring public import of COMMON_ENUM_ERROR_TYPE_ENTERCOINROOM from common_enum.proto

type VersionInfo struct {
	FileId           *int32 `protobuf:"varint,1,opt,name=fileId" json:"fileId,omitempty"`
	FileVer          *int32 `protobuf:"varint,2,opt,name=fileVer" json:"fileVer,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *VersionInfo) Reset()                    { *m = VersionInfo{} }
func (m *VersionInfo) String() string            { return proto.CompactTextString(m) }
func (*VersionInfo) ProtoMessage()               {}
func (*VersionInfo) Descriptor() ([]byte, []int) { return fileDescriptor17, []int{0} }

func (m *VersionInfo) GetFileId() int32 {
	if m != nil && m.FileId != nil {
		return *m.FileId
	}
	return 0
}

func (m *VersionInfo) GetFileVer() int32 {
	if m != nil && m.FileVer != nil {
		return *m.FileVer
	}
	return 0
}

type AssetInfo struct {
	FileId           *int32          `protobuf:"varint,1,opt,name=fileId" json:"fileId,omitempty"`
	FileVer          *int32          `protobuf:"varint,2,opt,name=fileVer" json:"fileVer,omitempty"`
	FilePath         *string         `protobuf:"bytes,3,opt,name=filePath" json:"filePath,omitempty"`
	FileSize         *int32          `protobuf:"varint,4,opt,name=fileSize" json:"fileSize,omitempty"`
	Md5              *string         `protobuf:"bytes,5,opt,name=md5" json:"md5,omitempty"`
	IsCompress       *bool           `protobuf:"varint,6,opt,name=isCompress" json:"isCompress,omitempty"`
	IsCode           *bool           `protobuf:"varint,7,opt,name=isCode" json:"isCode,omitempty"`
	GameId           *CommonEnumGame `protobuf:"varint,8,opt,name=gameId,enum=ddproto.CommonEnumGame" json:"gameId,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *AssetInfo) Reset()                    { *m = AssetInfo{} }
func (m *AssetInfo) String() string            { return proto.CompactTextString(m) }
func (*AssetInfo) ProtoMessage()               {}
func (*AssetInfo) Descriptor() ([]byte, []int) { return fileDescriptor17, []int{1} }

func (m *AssetInfo) GetFileId() int32 {
	if m != nil && m.FileId != nil {
		return *m.FileId
	}
	return 0
}

func (m *AssetInfo) GetFileVer() int32 {
	if m != nil && m.FileVer != nil {
		return *m.FileVer
	}
	return 0
}

func (m *AssetInfo) GetFilePath() string {
	if m != nil && m.FilePath != nil {
		return *m.FilePath
	}
	return ""
}

func (m *AssetInfo) GetFileSize() int32 {
	if m != nil && m.FileSize != nil {
		return *m.FileSize
	}
	return 0
}

func (m *AssetInfo) GetMd5() string {
	if m != nil && m.Md5 != nil {
		return *m.Md5
	}
	return ""
}

func (m *AssetInfo) GetIsCompress() bool {
	if m != nil && m.IsCompress != nil {
		return *m.IsCompress
	}
	return false
}

func (m *AssetInfo) GetIsCode() bool {
	if m != nil && m.IsCode != nil {
		return *m.IsCode
	}
	return false
}

func (m *AssetInfo) GetGameId() CommonEnumGame {
	if m != nil && m.GameId != nil {
		return *m.GameId
	}
	return CommonEnumGame_GID_SRC
}

// 获取资源文件的最新版本信息
type HotupdateReqVersionInfo struct {
	Header              *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	ClientAppId         *int32       `protobuf:"varint,2,opt,name=clientAppId" json:"clientAppId,omitempty"`
	DownloadGameId      []int32      `protobuf:"varint,3,rep,name=downloadGameId" json:"downloadGameId,omitempty"`
	ClientAssetsVersion *int32       `protobuf:"varint,4,opt,name=clientAssetsVersion" json:"clientAssetsVersion,omitempty"`
	ClientResolution    *int32       `protobuf:"varint,5,opt,name=clientResolution" json:"clientResolution,omitempty"`
	AppVersion          *int32       `protobuf:"varint,6,opt,name=AppVersion" json:"AppVersion,omitempty"`
	NativeVersion       *int32       `protobuf:"varint,7,opt,name=nativeVersion" json:"nativeVersion,omitempty"`
	ClientOSType        *int32       `protobuf:"varint,8,opt,name=clientOSType" json:"clientOSType,omitempty"`
	XXX_unrecognized    []byte       `json:"-"`
}

func (m *HotupdateReqVersionInfo) Reset()                    { *m = HotupdateReqVersionInfo{} }
func (m *HotupdateReqVersionInfo) String() string            { return proto.CompactTextString(m) }
func (*HotupdateReqVersionInfo) ProtoMessage()               {}
func (*HotupdateReqVersionInfo) Descriptor() ([]byte, []int) { return fileDescriptor17, []int{2} }

func (m *HotupdateReqVersionInfo) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *HotupdateReqVersionInfo) GetClientAppId() int32 {
	if m != nil && m.ClientAppId != nil {
		return *m.ClientAppId
	}
	return 0
}

func (m *HotupdateReqVersionInfo) GetDownloadGameId() []int32 {
	if m != nil {
		return m.DownloadGameId
	}
	return nil
}

func (m *HotupdateReqVersionInfo) GetClientAssetsVersion() int32 {
	if m != nil && m.ClientAssetsVersion != nil {
		return *m.ClientAssetsVersion
	}
	return 0
}

func (m *HotupdateReqVersionInfo) GetClientResolution() int32 {
	if m != nil && m.ClientResolution != nil {
		return *m.ClientResolution
	}
	return 0
}

func (m *HotupdateReqVersionInfo) GetAppVersion() int32 {
	if m != nil && m.AppVersion != nil {
		return *m.AppVersion
	}
	return 0
}

func (m *HotupdateReqVersionInfo) GetNativeVersion() int32 {
	if m != nil && m.NativeVersion != nil {
		return *m.NativeVersion
	}
	return 0
}

func (m *HotupdateReqVersionInfo) GetClientOSType() int32 {
	if m != nil && m.ClientOSType != nil {
		return *m.ClientOSType
	}
	return 0
}

type HotupdateAckVersionInfo struct {
	Header               *ProtoHeader   `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Version              []*VersionInfo `protobuf:"bytes,2,rep,name=version" json:"version,omitempty"`
	LastestAssetsVersion *int32         `protobuf:"varint,3,opt,name=lastestAssetsVersion" json:"lastestAssetsVersion,omitempty"`
	ForceUpdate          *bool          `protobuf:"varint,4,opt,name=forceUpdate" json:"forceUpdate,omitempty"`
	ReleaseTag           *int32         `protobuf:"varint,5,opt,name=releaseTag" json:"releaseTag,omitempty"`
	AppDownloadUrl       *string        `protobuf:"bytes,6,opt,name=appDownloadUrl" json:"appDownloadUrl,omitempty"`
	XXX_unrecognized     []byte         `json:"-"`
}

func (m *HotupdateAckVersionInfo) Reset()                    { *m = HotupdateAckVersionInfo{} }
func (m *HotupdateAckVersionInfo) String() string            { return proto.CompactTextString(m) }
func (*HotupdateAckVersionInfo) ProtoMessage()               {}
func (*HotupdateAckVersionInfo) Descriptor() ([]byte, []int) { return fileDescriptor17, []int{3} }

func (m *HotupdateAckVersionInfo) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *HotupdateAckVersionInfo) GetVersion() []*VersionInfo {
	if m != nil {
		return m.Version
	}
	return nil
}

func (m *HotupdateAckVersionInfo) GetLastestAssetsVersion() int32 {
	if m != nil && m.LastestAssetsVersion != nil {
		return *m.LastestAssetsVersion
	}
	return 0
}

func (m *HotupdateAckVersionInfo) GetForceUpdate() bool {
	if m != nil && m.ForceUpdate != nil {
		return *m.ForceUpdate
	}
	return false
}

func (m *HotupdateAckVersionInfo) GetReleaseTag() int32 {
	if m != nil && m.ReleaseTag != nil {
		return *m.ReleaseTag
	}
	return 0
}

func (m *HotupdateAckVersionInfo) GetAppDownloadUrl() string {
	if m != nil && m.AppDownloadUrl != nil {
		return *m.AppDownloadUrl
	}
	return ""
}

// 获取资源文件的详细信息
type HotupdateReqAssetsInfo struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	ClientAppId      *int32       `protobuf:"varint,2,opt,name=clientAppId" json:"clientAppId,omitempty"`
	FileIds          []int32      `protobuf:"varint,3,rep,name=fileIds" json:"fileIds,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *HotupdateReqAssetsInfo) Reset()                    { *m = HotupdateReqAssetsInfo{} }
func (m *HotupdateReqAssetsInfo) String() string            { return proto.CompactTextString(m) }
func (*HotupdateReqAssetsInfo) ProtoMessage()               {}
func (*HotupdateReqAssetsInfo) Descriptor() ([]byte, []int) { return fileDescriptor17, []int{4} }

func (m *HotupdateReqAssetsInfo) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *HotupdateReqAssetsInfo) GetClientAppId() int32 {
	if m != nil && m.ClientAppId != nil {
		return *m.ClientAppId
	}
	return 0
}

func (m *HotupdateReqAssetsInfo) GetFileIds() []int32 {
	if m != nil {
		return m.FileIds
	}
	return nil
}

type HotupdateAckAssetsInfo struct {
	Header               *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	AssetHost            *string      `protobuf:"bytes,2,opt,name=assetHost" json:"assetHost,omitempty"`
	Assets               []*AssetInfo `protobuf:"bytes,3,rep,name=assets" json:"assets,omitempty"`
	LastestAssetsVersion *int32       `protobuf:"varint,4,opt,name=lastestAssetsVersion" json:"lastestAssetsVersion,omitempty"`
	ClientResolution     *int32       `protobuf:"varint,5,opt,name=clientResolution" json:"clientResolution,omitempty"`
	XXX_unrecognized     []byte       `json:"-"`
}

func (m *HotupdateAckAssetsInfo) Reset()                    { *m = HotupdateAckAssetsInfo{} }
func (m *HotupdateAckAssetsInfo) String() string            { return proto.CompactTextString(m) }
func (*HotupdateAckAssetsInfo) ProtoMessage()               {}
func (*HotupdateAckAssetsInfo) Descriptor() ([]byte, []int) { return fileDescriptor17, []int{5} }

func (m *HotupdateAckAssetsInfo) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *HotupdateAckAssetsInfo) GetAssetHost() string {
	if m != nil && m.AssetHost != nil {
		return *m.AssetHost
	}
	return ""
}

func (m *HotupdateAckAssetsInfo) GetAssets() []*AssetInfo {
	if m != nil {
		return m.Assets
	}
	return nil
}

func (m *HotupdateAckAssetsInfo) GetLastestAssetsVersion() int32 {
	if m != nil && m.LastestAssetsVersion != nil {
		return *m.LastestAssetsVersion
	}
	return 0
}

func (m *HotupdateAckAssetsInfo) GetClientResolution() int32 {
	if m != nil && m.ClientResolution != nil {
		return *m.ClientResolution
	}
	return 0
}

// 获取某个独立游戏的下载信息
type HotupdateReqGameAssetsInfo struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	ClientAppId      *int32       `protobuf:"varint,2,opt,name=clientAppId" json:"clientAppId,omitempty"`
	GameId           *int32       `protobuf:"varint,3,opt,name=gameId" json:"gameId,omitempty"`
	ClientResolution *int32       `protobuf:"varint,4,opt,name=clientResolution" json:"clientResolution,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *HotupdateReqGameAssetsInfo) Reset()                    { *m = HotupdateReqGameAssetsInfo{} }
func (m *HotupdateReqGameAssetsInfo) String() string            { return proto.CompactTextString(m) }
func (*HotupdateReqGameAssetsInfo) ProtoMessage()               {}
func (*HotupdateReqGameAssetsInfo) Descriptor() ([]byte, []int) { return fileDescriptor17, []int{6} }

func (m *HotupdateReqGameAssetsInfo) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *HotupdateReqGameAssetsInfo) GetClientAppId() int32 {
	if m != nil && m.ClientAppId != nil {
		return *m.ClientAppId
	}
	return 0
}

func (m *HotupdateReqGameAssetsInfo) GetGameId() int32 {
	if m != nil && m.GameId != nil {
		return *m.GameId
	}
	return 0
}

func (m *HotupdateReqGameAssetsInfo) GetClientResolution() int32 {
	if m != nil && m.ClientResolution != nil {
		return *m.ClientResolution
	}
	return 0
}

type HotupdateAckGameAssetsInfo struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	AssetHost        *string      `protobuf:"bytes,2,opt,name=assetHost" json:"assetHost,omitempty"`
	Assets           []*AssetInfo `protobuf:"bytes,3,rep,name=assets" json:"assets,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *HotupdateAckGameAssetsInfo) Reset()                    { *m = HotupdateAckGameAssetsInfo{} }
func (m *HotupdateAckGameAssetsInfo) String() string            { return proto.CompactTextString(m) }
func (*HotupdateAckGameAssetsInfo) ProtoMessage()               {}
func (*HotupdateAckGameAssetsInfo) Descriptor() ([]byte, []int) { return fileDescriptor17, []int{7} }

func (m *HotupdateAckGameAssetsInfo) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *HotupdateAckGameAssetsInfo) GetAssetHost() string {
	if m != nil && m.AssetHost != nil {
		return *m.AssetHost
	}
	return ""
}

func (m *HotupdateAckGameAssetsInfo) GetAssets() []*AssetInfo {
	if m != nil {
		return m.Assets
	}
	return nil
}

func init() {
	proto.RegisterType((*VersionInfo)(nil), "ddproto.VersionInfo")
	proto.RegisterType((*AssetInfo)(nil), "ddproto.AssetInfo")
	proto.RegisterType((*HotupdateReqVersionInfo)(nil), "ddproto.hotupdate_req_versionInfo")
	proto.RegisterType((*HotupdateAckVersionInfo)(nil), "ddproto.hotupdate_ack_versionInfo")
	proto.RegisterType((*HotupdateReqAssetsInfo)(nil), "ddproto.hotupdate_req_assetsInfo")
	proto.RegisterType((*HotupdateAckAssetsInfo)(nil), "ddproto.hotupdate_ack_assetsInfo")
	proto.RegisterType((*HotupdateReqGameAssetsInfo)(nil), "ddproto.hotupdate_req_gameAssetsInfo")
	proto.RegisterType((*HotupdateAckGameAssetsInfo)(nil), "ddproto.hotupdate_ack_gameAssetsInfo")
}

func init() { proto.RegisterFile("hot_update.proto", fileDescriptor17) }

var fileDescriptor17 = []byte{
	// 511 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x53, 0x4d, 0x8e, 0xd3, 0x30,
	0x14, 0xc6, 0x0d, 0x4d, 0xdb, 0x17, 0x28, 0x1d, 0xb7, 0x8c, 0x3c, 0xc3, 0x2c, 0xa2, 0x08, 0xa4,
	0xb2, 0xe9, 0xa2, 0x12, 0x07, 0x28, 0x20, 0xd1, 0xae, 0xa8, 0x98, 0x9f, 0x6d, 0x65, 0xd5, 0xaf,
	0xd3, 0x88, 0x24, 0x36, 0xb1, 0x3b, 0x08, 0x58, 0x72, 0x1d, 0x24, 0x8e, 0xc1, 0x4d, 0x38, 0x07,
	0xb2, 0xe3, 0x96, 0x14, 0x69, 0x04, 0x14, 0x36, 0x51, 0xfc, 0xbd, 0x67, 0x7f, 0x3f, 0x7e, 0x86,
	0xde, 0x5a, 0x9a, 0xc5, 0x46, 0x09, 0x6e, 0x70, 0xa4, 0x4a, 0x69, 0x24, 0x6d, 0x09, 0xe1, 0x7e,
	0x4e, 0xfb, 0x4b, 0x99, 0xe7, 0xb2, 0x58, 0x2c, 0xb3, 0x14, 0x0b, 0x53, 0x55, 0x4f, 0x8f, 0x3c,
	0x88, 0xc5, 0x26, 0xaf, 0xa0, 0x64, 0x04, 0xd1, 0x15, 0x96, 0x3a, 0x95, 0xc5, 0xac, 0x58, 0x49,
	0xda, 0x85, 0x70, 0x95, 0x66, 0x38, 0x13, 0x8c, 0xc4, 0x64, 0xd8, 0xa4, 0x0f, 0xa0, 0x65, 0xd7,
	0x57, 0x58, 0xb2, 0x86, 0x05, 0x92, 0x2f, 0x04, 0x3a, 0x13, 0xad, 0xd1, 0xfc, 0x51, 0x3b, 0xed,
	0x41, 0xdb, 0x02, 0x73, 0x6e, 0xd6, 0x2c, 0x88, 0xc9, 0xb0, 0xb3, 0x45, 0xce, 0xd3, 0x8f, 0xc8,
	0xee, 0xba, 0x9e, 0x08, 0x82, 0x5c, 0x3c, 0x63, 0x4d, 0x57, 0xa6, 0x00, 0xa9, 0x7e, 0x21, 0x73,
	0x55, 0xa2, 0xd6, 0x2c, 0x8c, 0xc9, 0xb0, 0x6d, 0x59, 0x2c, 0x26, 0x90, 0xb5, 0xdc, 0xfa, 0x29,
	0x84, 0xd7, 0x3c, 0xb7, 0xac, 0xed, 0x98, 0x0c, 0xbb, 0xe3, 0x93, 0x91, 0x77, 0x3d, 0xaa, 0xf9,
	0x5b, 0xd8, 0x96, 0xe4, 0x3b, 0x81, 0x93, 0xb5, 0x34, 0x55, 0x46, 0x8b, 0x12, 0xdf, 0x2d, 0x6e,
	0x6a, 0x6e, 0x1f, 0x43, 0xb8, 0x46, 0x2e, 0xb0, 0x74, 0xf2, 0xa3, 0xf1, 0x60, 0x77, 0xd0, 0xdc,
	0x7e, 0xa7, 0xae, 0x46, 0xfb, 0x10, 0x55, 0x29, 0x4e, 0x94, 0x9a, 0x09, 0x6f, 0xec, 0x18, 0xba,
	0x42, 0xbe, 0x2f, 0x32, 0xc9, 0xc5, 0xab, 0x4a, 0x4b, 0x10, 0x07, 0xc3, 0x26, 0x7d, 0x04, 0x7d,
	0xdf, 0x6c, 0x43, 0xd2, 0x3e, 0x5b, 0xef, 0x94, 0x41, 0xaf, 0x2a, 0xbe, 0x41, 0x2d, 0xb3, 0x8d,
	0xb1, 0x95, 0xa6, 0xab, 0x50, 0x80, 0x89, 0x52, 0xdb, 0xee, 0xd0, 0x61, 0x0f, 0xe1, 0x7e, 0xc1,
	0x4d, 0x7a, 0x83, 0x5b, 0xb8, 0xe5, 0xe0, 0x01, 0xdc, 0xab, 0x0e, 0x79, 0x7d, 0x7e, 0xf1, 0x41,
	0xa1, 0xcb, 0xa0, 0x99, 0x7c, 0xdb, 0x33, 0xca, 0x97, 0x6f, 0x0f, 0x30, 0xfa, 0x04, 0x5a, 0x7e,
	0x13, 0x6b, 0xc4, 0xc1, 0x5e, 0x5b, 0x7d, 0x46, 0xce, 0x60, 0x90, 0x71, 0x6d, 0x50, 0xff, 0xe2,
	0x31, 0x70, 0xf2, 0xfa, 0x10, 0xad, 0x64, 0xb9, 0xc4, 0x4b, 0xa7, 0xc4, 0x19, 0x6f, 0x5b, 0x7b,
	0x25, 0x66, 0xc8, 0x35, 0x5e, 0xf0, 0x6b, 0x6f, 0xf9, 0x18, 0xba, 0x5c, 0xa9, 0x97, 0x3e, 0xc4,
	0xcb, 0x32, 0x73, 0xb6, 0x3b, 0xc9, 0x0a, 0xd8, 0xfe, 0x8d, 0x71, 0xc7, 0xf2, 0xaf, 0x17, 0xe6,
	0x47, 0x73, 0x26, 0x74, 0x75, 0x53, 0xc9, 0x57, 0x52, 0x27, 0xb2, 0x89, 0xfd, 0x35, 0xd1, 0x11,
	0x74, 0xdc, 0x9e, 0xa9, 0xd4, 0xc6, 0xd1, 0x74, 0x68, 0x02, 0x61, 0x75, 0x8c, 0x63, 0x89, 0xc6,
	0x74, 0xb7, 0xf1, 0xe7, 0xab, 0xb9, 0x2d, 0xc0, 0xdf, 0x0c, 0x49, 0xf2, 0x99, 0xc0, 0xd9, 0x7e,
	0x34, 0x76, 0xc6, 0x27, 0xff, 0x25, 0x9e, 0xee, 0xee, 0x4d, 0x05, 0xb7, 0xaa, 0x70, 0xfa, 0x92,
	0x4f, 0x75, 0x11, 0x36, 0xb6, 0x83, 0x44, 0x1c, 0x16, 0xdd, 0xf3, 0xc6, 0x34, 0x98, 0xdf, 0x99,
	0x93, 0x1f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xae, 0x93, 0x2c, 0x57, 0xfa, 0x04, 0x00, 0x00,
}
