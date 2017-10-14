// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common_client_pay.proto

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

// 订单的状态
type PayEnumTradeStatus int32

const (
	PayEnumTradeStatus_PAY_S_UNIFIEDORDER PayEnumTradeStatus = 1
	PayEnumTradeStatus_PAY_S_WATINGPAY    PayEnumTradeStatus = 2
	PayEnumTradeStatus_PAY_S_SUCC         PayEnumTradeStatus = 3
)

var PayEnumTradeStatus_name = map[int32]string{
	1: "PAY_S_UNIFIEDORDER",
	2: "PAY_S_WATINGPAY",
	3: "PAY_S_SUCC",
}
var PayEnumTradeStatus_value = map[string]int32{
	"PAY_S_UNIFIEDORDER": 1,
	"PAY_S_WATINGPAY":    2,
	"PAY_S_SUCC":         3,
}

func (x PayEnumTradeStatus) Enum() *PayEnumTradeStatus {
	p := new(PayEnumTradeStatus)
	*p = x
	return p
}
func (x PayEnumTradeStatus) String() string {
	return proto.EnumName(PayEnumTradeStatus_name, int32(x))
}
func (x *PayEnumTradeStatus) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(PayEnumTradeStatus_value, data, "PayEnumTradeStatus")
	if err != nil {
		return err
	}
	*x = PayEnumTradeStatus(value)
	return nil
}
func (PayEnumTradeStatus) EnumDescriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

// app收到同步回调之后，请求服务器检测接收异步回调，得到充值之后的信息
type PayReturn int32

const (
	PayReturn_ERR_ONLY_APPLE PayReturn = -250
)

var PayReturn_name = map[int32]string{
	-250: "ERR_ONLY_APPLE",
}
var PayReturn_value = map[string]int32{
	"ERR_ONLY_APPLE": -250,
}

func (x PayReturn) Enum() *PayReturn {
	p := new(PayReturn)
	*p = x
	return p
}
func (x PayReturn) String() string {
	return proto.EnumName(PayReturn_name, int32(x))
}
func (x *PayReturn) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(PayReturn_value, data, "PayReturn")
	if err != nil {
		return err
	}
	*x = PayReturn(value)
	return nil
}
func (PayReturn) EnumDescriptor() ([]byte, []int) { return fileDescriptor5, []int{1} }

// 充值套餐
type PayBaseProduct struct {
	Id               *int32  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Money            *int64  `protobuf:"varint,2,opt,name=money" json:"money,omitempty"`
	Diamond          *int64  `protobuf:"varint,3,opt,name=diamond" json:"diamond,omitempty"`
	Memo             *string `protobuf:"bytes,4,opt,name=memo" json:"memo,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *PayBaseProduct) Reset()                    { *m = PayBaseProduct{} }
func (m *PayBaseProduct) String() string            { return proto.CompactTextString(m) }
func (*PayBaseProduct) ProtoMessage()               {}
func (*PayBaseProduct) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

func (m *PayBaseProduct) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *PayBaseProduct) GetMoney() int64 {
	if m != nil && m.Money != nil {
		return *m.Money
	}
	return 0
}

func (m *PayBaseProduct) GetDiamond() int64 {
	if m != nil && m.Diamond != nil {
		return *m.Diamond
	}
	return 0
}

func (m *PayBaseProduct) GetMemo() string {
	if m != nil && m.Memo != nil {
		return *m.Memo
	}
	return ""
}

// 付款的模式,制定对应的商户号，appid,appkey
type PayBasePaymodel struct {
	Id               *int32  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	AppId            *string `protobuf:"bytes,2,opt,name=appId" json:"appId,omitempty"`
	MchId            *string `protobuf:"bytes,3,opt,name=mchId" json:"mchId,omitempty"`
	AppKey           *string `protobuf:"bytes,4,opt,name=appKey" json:"appKey,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *PayBasePaymodel) Reset()                    { *m = PayBasePaymodel{} }
func (m *PayBasePaymodel) String() string            { return proto.CompactTextString(m) }
func (*PayBasePaymodel) ProtoMessage()               {}
func (*PayBasePaymodel) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{1} }

func (m *PayBasePaymodel) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *PayBasePaymodel) GetAppId() string {
	if m != nil && m.AppId != nil {
		return *m.AppId
	}
	return ""
}

func (m *PayBasePaymodel) GetMchId() string {
	if m != nil && m.MchId != nil {
		return *m.MchId
	}
	return ""
}

func (m *PayBasePaymodel) GetAppKey() string {
	if m != nil && m.AppKey != nil {
		return *m.AppKey
	}
	return ""
}

// 支付明细
type PayBaseDetails struct {
	Id               *int32              `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	UserId           *uint32             `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	PayModelId       *int32              `protobuf:"varint,3,opt,name=payModelId" json:"payModelId,omitempty"`
	ProductId        *int32              `protobuf:"varint,4,opt,name=productId" json:"productId,omitempty"`
	TradeNo          *string             `protobuf:"bytes,5,opt,name=tradeNo" json:"tradeNo,omitempty"`
	Status           *PayEnumTradeStatus `protobuf:"varint,6,opt,name=status,enum=ddproto.PayEnumTradeStatus" json:"status,omitempty"`
	Diamond          *int64              `protobuf:"varint,7,opt,name=Diamond" json:"Diamond,omitempty"`
	ChangeDiamond    *int64              `protobuf:"varint,8,opt,name=ChangeDiamond" json:"ChangeDiamond,omitempty"`
	Coin             *int64              `protobuf:"varint,9,opt,name=Coin" json:"Coin,omitempty"`
	ChangeCoin       *int64              `protobuf:"varint,10,opt,name=ChangeCoin" json:"ChangeCoin,omitempty"`
	XXX_unrecognized []byte              `json:"-"`
}

func (m *PayBaseDetails) Reset()                    { *m = PayBaseDetails{} }
func (m *PayBaseDetails) String() string            { return proto.CompactTextString(m) }
func (*PayBaseDetails) ProtoMessage()               {}
func (*PayBaseDetails) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{2} }

func (m *PayBaseDetails) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *PayBaseDetails) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *PayBaseDetails) GetPayModelId() int32 {
	if m != nil && m.PayModelId != nil {
		return *m.PayModelId
	}
	return 0
}

func (m *PayBaseDetails) GetProductId() int32 {
	if m != nil && m.ProductId != nil {
		return *m.ProductId
	}
	return 0
}

func (m *PayBaseDetails) GetTradeNo() string {
	if m != nil && m.TradeNo != nil {
		return *m.TradeNo
	}
	return ""
}

func (m *PayBaseDetails) GetStatus() PayEnumTradeStatus {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return PayEnumTradeStatus_PAY_S_UNIFIEDORDER
}

func (m *PayBaseDetails) GetDiamond() int64 {
	if m != nil && m.Diamond != nil {
		return *m.Diamond
	}
	return 0
}

func (m *PayBaseDetails) GetChangeDiamond() int64 {
	if m != nil && m.ChangeDiamond != nil {
		return *m.ChangeDiamond
	}
	return 0
}

func (m *PayBaseDetails) GetCoin() int64 {
	if m != nil && m.Coin != nil {
		return *m.Coin
	}
	return 0
}

func (m *PayBaseDetails) GetChangeCoin() int64 {
	if m != nil && m.ChangeCoin != nil {
		return *m.ChangeCoin
	}
	return 0
}

// 统一下单 proto
//    总金额	total_fee	是	Int	888	订单总金额，单位为分，详见支付金额
type WxpayReqUnifiedorder struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	ProductId        *int32       `protobuf:"varint,2,opt,name=productId" json:"productId,omitempty"`
	PayModelId       *int32       `protobuf:"varint,3,opt,name=payModelId" json:"payModelId,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *WxpayReqUnifiedorder) Reset()                    { *m = WxpayReqUnifiedorder{} }
func (m *WxpayReqUnifiedorder) String() string            { return proto.CompactTextString(m) }
func (*WxpayReqUnifiedorder) ProtoMessage()               {}
func (*WxpayReqUnifiedorder) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{3} }

func (m *WxpayReqUnifiedorder) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *WxpayReqUnifiedorder) GetProductId() int32 {
	if m != nil && m.ProductId != nil {
		return *m.ProductId
	}
	return 0
}

func (m *WxpayReqUnifiedorder) GetPayModelId() int32 {
	if m != nil && m.PayModelId != nil {
		return *m.PayModelId
	}
	return 0
}

// app请求统一下单之后，返回的加密字符串
type WxpayAckUnifiedorder struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	PartnerId        *string      `protobuf:"bytes,2,opt,name=partnerId" json:"partnerId,omitempty"`
	PrepayId         *string      `protobuf:"bytes,3,opt,name=prepay_id,json=prepayId" json:"prepay_id,omitempty"`
	NonceStr         *string      `protobuf:"bytes,4,opt,name=nonce_str,json=nonceStr" json:"nonce_str,omitempty"`
	TimeStamp        *string      `protobuf:"bytes,5,opt,name=timeStamp" json:"timeStamp,omitempty"`
	Package          *string      `protobuf:"bytes,6,opt,name=package" json:"package,omitempty"`
	Sign             *string      `protobuf:"bytes,7,opt,name=sign" json:"sign,omitempty"`
	TradeNo          *string      `protobuf:"bytes,8,opt,name=tradeNo" json:"tradeNo,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *WxpayAckUnifiedorder) Reset()                    { *m = WxpayAckUnifiedorder{} }
func (m *WxpayAckUnifiedorder) String() string            { return proto.CompactTextString(m) }
func (*WxpayAckUnifiedorder) ProtoMessage()               {}
func (*WxpayAckUnifiedorder) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{4} }

func (m *WxpayAckUnifiedorder) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *WxpayAckUnifiedorder) GetPartnerId() string {
	if m != nil && m.PartnerId != nil {
		return *m.PartnerId
	}
	return ""
}

func (m *WxpayAckUnifiedorder) GetPrepayId() string {
	if m != nil && m.PrepayId != nil {
		return *m.PrepayId
	}
	return ""
}

func (m *WxpayAckUnifiedorder) GetNonceStr() string {
	if m != nil && m.NonceStr != nil {
		return *m.NonceStr
	}
	return ""
}

func (m *WxpayAckUnifiedorder) GetTimeStamp() string {
	if m != nil && m.TimeStamp != nil {
		return *m.TimeStamp
	}
	return ""
}

func (m *WxpayAckUnifiedorder) GetPackage() string {
	if m != nil && m.Package != nil {
		return *m.Package
	}
	return ""
}

func (m *WxpayAckUnifiedorder) GetSign() string {
	if m != nil && m.Sign != nil {
		return *m.Sign
	}
	return ""
}

func (m *WxpayAckUnifiedorder) GetTradeNo() string {
	if m != nil && m.TradeNo != nil {
		return *m.TradeNo
	}
	return ""
}

type WxpayReqSyncChecker struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	TradeNo          *string      `protobuf:"bytes,2,opt,name=tradeNo" json:"tradeNo,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *WxpayReqSyncChecker) Reset()                    { *m = WxpayReqSyncChecker{} }
func (m *WxpayReqSyncChecker) String() string            { return proto.CompactTextString(m) }
func (*WxpayReqSyncChecker) ProtoMessage()               {}
func (*WxpayReqSyncChecker) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{5} }

func (m *WxpayReqSyncChecker) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *WxpayReqSyncChecker) GetTradeNo() string {
	if m != nil && m.TradeNo != nil {
		return *m.TradeNo
	}
	return ""
}

// 服务器收到微信充值的异步回调之后，返回给app当前的余额信息
type WxpayAckSyncChecker struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	WxpayId          *int32       `protobuf:"varint,2,opt,name=wxpayId" json:"wxpayId,omitempty"`
	Diamond          *int64       `protobuf:"varint,3,opt,name=diamond" json:"diamond,omitempty"`
	ChangeDiamond    *int64       `protobuf:"varint,4,opt,name=changeDiamond" json:"changeDiamond,omitempty"`
	Coin             *int64       `protobuf:"varint,5,opt,name=coin" json:"coin,omitempty"`
	ChangeCoin       *int64       `protobuf:"varint,6,opt,name=changeCoin" json:"changeCoin,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *WxpayAckSyncChecker) Reset()                    { *m = WxpayAckSyncChecker{} }
func (m *WxpayAckSyncChecker) String() string            { return proto.CompactTextString(m) }
func (*WxpayAckSyncChecker) ProtoMessage()               {}
func (*WxpayAckSyncChecker) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{6} }

func (m *WxpayAckSyncChecker) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *WxpayAckSyncChecker) GetWxpayId() int32 {
	if m != nil && m.WxpayId != nil {
		return *m.WxpayId
	}
	return 0
}

func (m *WxpayAckSyncChecker) GetDiamond() int64 {
	if m != nil && m.Diamond != nil {
		return *m.Diamond
	}
	return 0
}

func (m *WxpayAckSyncChecker) GetChangeDiamond() int64 {
	if m != nil && m.ChangeDiamond != nil {
		return *m.ChangeDiamond
	}
	return 0
}

func (m *WxpayAckSyncChecker) GetCoin() int64 {
	if m != nil && m.Coin != nil {
		return *m.Coin
	}
	return 0
}

func (m *WxpayAckSyncChecker) GetChangeCoin() int64 {
	if m != nil && m.ChangeCoin != nil {
		return *m.ChangeCoin
	}
	return 0
}

// 苹果支付充值回调
type ApplepayReqRechargecb struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	ProductId        *int32       `protobuf:"varint,2,opt,name=productId" json:"productId,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *ApplepayReqRechargecb) Reset()                    { *m = ApplepayReqRechargecb{} }
func (m *ApplepayReqRechargecb) String() string            { return proto.CompactTextString(m) }
func (*ApplepayReqRechargecb) ProtoMessage()               {}
func (*ApplepayReqRechargecb) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{7} }

func (m *ApplepayReqRechargecb) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *ApplepayReqRechargecb) GetProductId() int32 {
	if m != nil && m.ProductId != nil {
		return *m.ProductId
	}
	return 0
}

type ApplepayAcksRechargecb struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	RechargeDiamond  *int64       `protobuf:"varint,2,opt,name=rechargeDiamond" json:"rechargeDiamond,omitempty"`
	Diamond          *int64       `protobuf:"varint,3,opt,name=diamond" json:"diamond,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *ApplepayAcksRechargecb) Reset()                    { *m = ApplepayAcksRechargecb{} }
func (m *ApplepayAcksRechargecb) String() string            { return proto.CompactTextString(m) }
func (*ApplepayAcksRechargecb) ProtoMessage()               {}
func (*ApplepayAcksRechargecb) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{8} }

func (m *ApplepayAcksRechargecb) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *ApplepayAcksRechargecb) GetRechargeDiamond() int64 {
	if m != nil && m.RechargeDiamond != nil {
		return *m.RechargeDiamond
	}
	return 0
}

func (m *ApplepayAcksRechargecb) GetDiamond() int64 {
	if m != nil && m.Diamond != nil {
		return *m.Diamond
	}
	return 0
}

func init() {
	proto.RegisterType((*PayBaseProduct)(nil), "ddproto.pay_base_product")
	proto.RegisterType((*PayBasePaymodel)(nil), "ddproto.pay_base_paymodel")
	proto.RegisterType((*PayBaseDetails)(nil), "ddproto.pay_base_details")
	proto.RegisterType((*WxpayReqUnifiedorder)(nil), "ddproto.wxpay_req_unifiedorder")
	proto.RegisterType((*WxpayAckUnifiedorder)(nil), "ddproto.wxpay_ack_unifiedorder")
	proto.RegisterType((*WxpayReqSyncChecker)(nil), "ddproto.wxpay_req_syncChecker")
	proto.RegisterType((*WxpayAckSyncChecker)(nil), "ddproto.wxpay_ack_syncChecker")
	proto.RegisterType((*ApplepayReqRechargecb)(nil), "ddproto.applepay_req_rechargecb")
	proto.RegisterType((*ApplepayAcksRechargecb)(nil), "ddproto.applepay_acks_rechargecb")
	proto.RegisterEnum("ddproto.PayEnumTradeStatus", PayEnumTradeStatus_name, PayEnumTradeStatus_value)
	proto.RegisterEnum("ddproto.PayReturn", PayReturn_name, PayReturn_value)
}

func init() { proto.RegisterFile("common_client_pay.proto", fileDescriptor5) }

var fileDescriptor5 = []byte{
	// 681 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x54, 0x4d, 0x6f, 0xd3, 0x4a,
	0x14, 0x7d, 0x76, 0x9b, 0x34, 0xbe, 0x4f, 0x4d, 0xf3, 0xa6, 0x7d, 0xa9, 0xf5, 0xfa, 0x40, 0x91,
	0xd5, 0x45, 0xa8, 0x50, 0x17, 0x95, 0xf8, 0x01, 0x51, 0x1a, 0x20, 0xa2, 0xa4, 0xd1, 0x84, 0x0a,
	0x75, 0x65, 0x4d, 0x67, 0xa6, 0x89, 0x95, 0x78, 0x6c, 0xc6, 0x8e, 0x20, 0xfb, 0x8a, 0x15, 0x7f,
	0x89, 0x5f, 0xc1, 0xff, 0x01, 0x34, 0x1f, 0x8e, 0x93, 0x52, 0x40, 0x50, 0xb2, 0x88, 0xe6, 0x9e,
	0x7b, 0xe3, 0x73, 0x7d, 0xce, 0xc9, 0xc0, 0x3e, 0x4d, 0xe2, 0x38, 0x11, 0x21, 0x9d, 0x45, 0x5c,
	0xe4, 0x61, 0x4a, 0x16, 0xc7, 0xa9, 0x4c, 0xf2, 0x04, 0x6d, 0x31, 0xa6, 0x0f, 0xff, 0xed, 0xae,
	0x4d, 0x98, 0x6e, 0x70, 0x0d, 0x8d, 0x94, 0x2c, 0xc2, 0x2b, 0x92, 0xf1, 0x30, 0x95, 0x09, 0x9b,
	0xd3, 0x1c, 0xd5, 0xc1, 0x8d, 0x98, 0xef, 0xb4, 0x9c, 0x76, 0x05, 0xbb, 0x11, 0x43, 0x7b, 0x50,
	0x89, 0x13, 0xc1, 0x17, 0xbe, 0xdb, 0x72, 0xda, 0x1b, 0xd8, 0x14, 0xc8, 0x87, 0x2d, 0x16, 0x91,
	0x38, 0x11, 0xcc, 0xdf, 0xd0, 0x78, 0x51, 0x22, 0x04, 0x9b, 0x31, 0x8f, 0x13, 0x7f, 0xb3, 0xe5,
	0xb4, 0x3d, 0xac, 0xcf, 0xc1, 0x18, 0xfe, 0x29, 0x79, 0xc8, 0x22, 0x4e, 0x18, 0x9f, 0xdd, 0x45,
	0x44, 0xd2, 0xb4, 0xcf, 0x34, 0x91, 0x87, 0x4d, 0xa1, 0xe9, 0xe9, 0xa4, 0x6f, 0x68, 0x3c, 0x6c,
	0x0a, 0xd4, 0x84, 0x2a, 0x49, 0xd3, 0x17, 0x7c, 0x61, 0x69, 0x6c, 0x15, 0x7c, 0x74, 0x57, 0xde,
	0x88, 0xf1, 0x9c, 0x44, 0xb3, 0xec, 0x1b, 0xa2, 0x26, 0x54, 0xe7, 0x19, 0x97, 0x96, 0x69, 0x1b,
	0xdb, 0x0a, 0x3d, 0x04, 0x48, 0xc9, 0xe2, 0xa5, 0x5a, 0xce, 0xf2, 0x55, 0xf0, 0x0a, 0x82, 0xfe,
	0x07, 0xcf, 0x8a, 0xd4, 0x67, 0x9a, 0xb7, 0x82, 0x4b, 0x40, 0x29, 0x92, 0x4b, 0xc2, 0xf8, 0x20,
	0xf1, 0x2b, 0x7a, 0xa7, 0xa2, 0x44, 0x4f, 0xa0, 0x9a, 0xe5, 0x24, 0x9f, 0x67, 0x7e, 0xb5, 0xe5,
	0xb4, 0xeb, 0x27, 0x0f, 0x8e, 0xad, 0x29, 0xc7, 0x6a, 0x55, 0x2e, 0xe6, 0x71, 0xa8, 0x47, 0x47,
	0x7a, 0x08, 0xdb, 0x61, 0xf5, 0xc0, 0x53, 0x2b, 0xf1, 0x96, 0x91, 0xd8, 0x96, 0xe8, 0x10, 0xb6,
	0xbb, 0x13, 0x22, 0xc6, 0xbc, 0xe8, 0xd7, 0x74, 0x7f, 0x1d, 0x54, 0x46, 0x74, 0x93, 0x48, 0xf8,
	0x9e, 0x6e, 0xea, 0xb3, 0x7a, 0x45, 0x33, 0xa4, 0x3b, 0xa0, 0x3b, 0x2b, 0x48, 0x70, 0xe3, 0x40,
	0xf3, 0xed, 0x3b, 0xb5, 0x96, 0xe4, 0x6f, 0xc2, 0xb9, 0x88, 0xae, 0x23, 0xce, 0x12, 0xc9, 0xb8,
	0x44, 0x8f, 0xa1, 0x3a, 0xe1, 0x84, 0x71, 0xa9, 0x95, 0xfc, 0xfb, 0x64, 0x6f, 0xf9, 0x16, 0x43,
	0xf5, 0xfd, 0x5c, 0xf7, 0xb0, 0x9d, 0x59, 0xd7, 0xca, 0xbd, 0xad, 0xd5, 0x4f, 0x94, 0x0e, 0x6e,
	0xdc, 0x62, 0x0d, 0x42, 0xa7, 0xf7, 0x5d, 0x83, 0xc8, 0x5c, 0x2c, 0xdd, 0xf6, 0x70, 0x09, 0xa0,
	0x03, 0xb5, 0x24, 0x57, 0x34, 0x51, 0x91, 0xaf, 0x9a, 0x01, 0x4c, 0x53, 0x24, 0x82, 0xf2, 0x30,
	0xcb, 0xa5, 0x4d, 0x59, 0x4d, 0x03, 0xa3, 0x5c, 0x3f, 0x37, 0x8f, 0x62, 0xe5, 0x58, 0x9c, 0x5a,
	0xbb, 0x4b, 0x40, 0x39, 0x97, 0x12, 0x3a, 0x25, 0x63, 0xae, 0x1d, 0xf7, 0x70, 0x51, 0x2a, 0x4f,
	0xb2, 0x68, 0x2c, 0xb4, 0xa1, 0x1e, 0xd6, 0xe7, 0xd5, 0xe0, 0xd4, 0xd6, 0x82, 0x13, 0x84, 0xf0,
	0x6f, 0x69, 0x46, 0xb6, 0x10, 0xb4, 0x3b, 0xe1, 0x74, 0xfa, 0xcb, 0x22, 0xac, 0x10, 0xb8, 0xeb,
	0x04, 0x9f, 0x9c, 0x82, 0x41, 0xe9, 0x7c, 0x2f, 0x06, 0xfd, 0x98, 0xa5, 0xd7, 0x45, 0xf9, 0x83,
	0x7b, 0xe2, 0x10, 0xb6, 0xe9, 0x5a, 0x88, 0x37, 0x4d, 0x88, 0xe9, 0xed, 0x10, 0x53, 0x15, 0xd5,
	0x8a, 0x09, 0x31, 0xb5, 0x21, 0xa6, 0x65, 0x88, 0xab, 0x26, 0xc4, 0x25, 0x12, 0x70, 0xd8, 0x27,
	0x69, 0x3a, 0xe3, 0x85, 0x72, 0x92, 0xd3, 0x09, 0x91, 0x63, 0x4e, 0xaf, 0xfe, 0x64, 0x88, 0x83,
	0x0f, 0x0e, 0xf8, 0x4b, 0x1e, 0x42, 0xa7, 0xd9, 0xef, 0x13, 0xb5, 0x61, 0xa7, 0xf8, 0x6d, 0xa1,
	0x86, 0xb9, 0x6d, 0x6f, 0xc3, 0xdf, 0xd7, 0xf3, 0x68, 0x04, 0x7b, 0x77, 0x5d, 0x27, 0xa8, 0x09,
	0x68, 0xd8, 0xb9, 0x0c, 0x47, 0xe1, 0xc5, 0xa0, 0xff, 0xb4, 0xdf, 0x3b, 0x3d, 0xc7, 0xa7, 0x3d,
	0xdc, 0x70, 0xd0, 0x2e, 0xec, 0x18, 0xfc, 0x75, 0xe7, 0x55, 0x7f, 0xf0, 0x6c, 0xd8, 0xb9, 0x6c,
	0xb8, 0xa8, 0x0e, 0x60, 0xc0, 0xd1, 0x45, 0xb7, 0xdb, 0xd8, 0x38, 0x7a, 0xa4, 0xff, 0xa8, 0xa1,
	0xe4, 0xf9, 0x5c, 0x0a, 0x74, 0x00, 0xf5, 0x1e, 0xc6, 0xe1, 0xf9, 0xe0, 0xec, 0x32, 0xec, 0x0c,
	0x87, 0x67, 0xbd, 0xc6, 0xfb, 0xcf, 0x5f, 0xcc, 0xc7, 0x19, 0xfe, 0xf5, 0x35, 0x00, 0x00, 0xff,
	0xff, 0x2b, 0x2b, 0xff, 0x0f, 0x85, 0x06, 0x00, 0x00,
}
