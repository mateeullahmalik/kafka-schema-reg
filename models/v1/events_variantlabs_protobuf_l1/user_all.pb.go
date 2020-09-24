// Code generated by protoc-gen-go. DO NOT EDIT.
// source: events/variantlabs/protobuf/l1/user_all.proto

package events_variantlabs_protobuf_l1

import (
	fmt "fmt"
	math "math"

	l0 "github.com/mateeullahmalik/kafka-schema-reg/models/v1/events_variantlabs_protobuf_l0"

	proto "github.com/golang/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type UserAll struct {
	// Types that are valid to be assigned to OneofType:
	//	*UserAll_Registration
	//	*UserAll_DocUpload
	//	*UserAll_Login
	//	*UserAll_MoneyDeposit
	//	*UserAll_TableStart
	//	*UserAll_HandPlayed
	//	*UserAll_Bet
	//	*UserAll_MoneyWithdraw
	//	*UserAll_BonusReceived
	OneofType            isUserAll_OneofType `protobuf_oneof:"oneof_type"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *UserAll) Reset()         { *m = UserAll{} }
func (m *UserAll) String() string { return proto.CompactTextString(m) }
func (*UserAll) ProtoMessage()    {}
func (*UserAll) Descriptor() ([]byte, []int) {
	return fileDescriptor_b2f4de4a35c4b1dc, []int{0}
}

func (m *UserAll) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserAll.Unmarshal(m, b)
}
func (m *UserAll) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserAll.Marshal(b, m, deterministic)
}
func (m *UserAll) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserAll.Merge(m, src)
}
func (m *UserAll) XXX_Size() int {
	return xxx_messageInfo_UserAll.Size(m)
}
func (m *UserAll) XXX_DiscardUnknown() {
	xxx_messageInfo_UserAll.DiscardUnknown(m)
}

var xxx_messageInfo_UserAll proto.InternalMessageInfo

type isUserAll_OneofType interface {
	isUserAll_OneofType()
}

type UserAll_Registration struct {
	Registration *l0.UserRegistered `protobuf:"bytes,1,opt,name=registration,proto3,oneof"`
}

type UserAll_DocUpload struct {
	DocUpload *l0.UserDocumentUploaded `protobuf:"bytes,2,opt,name=doc_upload,json=docUpload,proto3,oneof"`
}

type UserAll_Login struct {
	Login *l0.UserLoggedIn `protobuf:"bytes,3,opt,name=login,proto3,oneof"`
}

type UserAll_MoneyDeposit struct {
	MoneyDeposit *l0.UserMoneyDeposited `protobuf:"bytes,4,opt,name=money_deposit,json=moneyDeposit,proto3,oneof"`
}

type UserAll_TableStart struct {
	TableStart *l0.UserTableStarted `protobuf:"bytes,5,opt,name=table_start,json=tableStart,proto3,oneof"`
}

type UserAll_HandPlayed struct {
	HandPlayed *l0.UserHandPlayed `protobuf:"bytes,6,opt,name=hand_played,json=handPlayed,proto3,oneof"`
}

type UserAll_Bet struct {
	Bet *l0.UserBetMade `protobuf:"bytes,7,opt,name=bet,proto3,oneof"`
}

type UserAll_MoneyWithdraw struct {
	MoneyWithdraw *l0.UserMoneyWithdrawn `protobuf:"bytes,8,opt,name=money_withdraw,json=moneyWithdraw,proto3,oneof"`
}

type UserAll_BonusReceived struct {
	BonusReceived *l0.UserBonusReceived `protobuf:"bytes,9,opt,name=bonus_received,json=bonusReceived,proto3,oneof"`
}

func (*UserAll_Registration) isUserAll_OneofType() {}

func (*UserAll_DocUpload) isUserAll_OneofType() {}

func (*UserAll_Login) isUserAll_OneofType() {}

func (*UserAll_MoneyDeposit) isUserAll_OneofType() {}

func (*UserAll_TableStart) isUserAll_OneofType() {}

func (*UserAll_HandPlayed) isUserAll_OneofType() {}

func (*UserAll_Bet) isUserAll_OneofType() {}

func (*UserAll_MoneyWithdraw) isUserAll_OneofType() {}

func (*UserAll_BonusReceived) isUserAll_OneofType() {}

func (m *UserAll) GetOneofType() isUserAll_OneofType {
	if m != nil {
		return m.OneofType
	}
	return nil
}

func (m *UserAll) GetRegistration() *l0.UserRegistered {
	if x, ok := m.GetOneofType().(*UserAll_Registration); ok {
		return x.Registration
	}
	return nil
}

func (m *UserAll) GetDocUpload() *l0.UserDocumentUploaded {
	if x, ok := m.GetOneofType().(*UserAll_DocUpload); ok {
		return x.DocUpload
	}
	return nil
}

func (m *UserAll) GetLogin() *l0.UserLoggedIn {
	if x, ok := m.GetOneofType().(*UserAll_Login); ok {
		return x.Login
	}
	return nil
}

func (m *UserAll) GetMoneyDeposit() *l0.UserMoneyDeposited {
	if x, ok := m.GetOneofType().(*UserAll_MoneyDeposit); ok {
		return x.MoneyDeposit
	}
	return nil
}

func (m *UserAll) GetTableStart() *l0.UserTableStarted {
	if x, ok := m.GetOneofType().(*UserAll_TableStart); ok {
		return x.TableStart
	}
	return nil
}

func (m *UserAll) GetHandPlayed() *l0.UserHandPlayed {
	if x, ok := m.GetOneofType().(*UserAll_HandPlayed); ok {
		return x.HandPlayed
	}
	return nil
}

func (m *UserAll) GetBet() *l0.UserBetMade {
	if x, ok := m.GetOneofType().(*UserAll_Bet); ok {
		return x.Bet
	}
	return nil
}

func (m *UserAll) GetMoneyWithdraw() *l0.UserMoneyWithdrawn {
	if x, ok := m.GetOneofType().(*UserAll_MoneyWithdraw); ok {
		return x.MoneyWithdraw
	}
	return nil
}

func (m *UserAll) GetBonusReceived() *l0.UserBonusReceived {
	if x, ok := m.GetOneofType().(*UserAll_BonusReceived); ok {
		return x.BonusReceived
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*UserAll) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*UserAll_Registration)(nil),
		(*UserAll_DocUpload)(nil),
		(*UserAll_Login)(nil),
		(*UserAll_MoneyDeposit)(nil),
		(*UserAll_TableStart)(nil),
		(*UserAll_HandPlayed)(nil),
		(*UserAll_Bet)(nil),
		(*UserAll_MoneyWithdraw)(nil),
		(*UserAll_BonusReceived)(nil),
	}
}

func init() {
	proto.RegisterType((*UserAll)(nil), "events.variantlabs.protobuf.l1.UserAll")
}

func init() {
	proto.RegisterFile("events/variantlabs/protobuf/l1/user_all.proto", fileDescriptor_b2f4de4a35c4b1dc)
}

var fileDescriptor_b2f4de4a35c4b1dc = []byte{
	// 456 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x94, 0xcf, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0x19, 0x63, 0x1d, 0x73, 0xc7, 0x0e, 0x3e, 0x59, 0x3b, 0x20, 0xc4, 0x09, 0x09, 0x48,
	0xda, 0x6e, 0x08, 0x10, 0x07, 0xc4, 0xd4, 0xc3, 0x90, 0x98, 0x04, 0xdd, 0x2a, 0x04, 0x1c, 0x2c,
	0x67, 0x79, 0x6b, 0x2d, 0x39, 0x76, 0xe4, 0x38, 0x9d, 0xfa, 0x17, 0xf0, 0x6f, 0xcf, 0x3f, 0x92,
	0x28, 0xb9, 0x4c, 0xee, 0xd1, 0xaf, 0xef, 0xf3, 0x71, 0xfd, 0x7d, 0xaf, 0x45, 0xef, 0x61, 0x03,
	0xd2, 0x54, 0xe9, 0x86, 0x69, 0xce, 0xa4, 0x11, 0x2c, 0xab, 0xd2, 0x52, 0x2b, 0xa3, 0xb2, 0xfa,
	0x2e, 0x15, 0xd3, 0xb4, 0xae, 0x40, 0x53, 0x26, 0x44, 0xe2, 0x8b, 0xf8, 0x65, 0x68, 0x4f, 0x7a,
	0xed, 0x49, 0xdb, 0x9e, 0x88, 0xe9, 0xe9, 0xa7, 0x47, 0x75, 0x93, 0xa0, 0xcb, 0x94, 0xac, 0x2b,
	0xaa, 0xe1, 0x16, 0xf8, 0x06, 0xf2, 0xc0, 0x9f, 0x7e, 0x88, 0x21, 0xd7, 0x4c, 0xe6, 0xb4, 0x14,
	0x6c, 0xdb, 0x61, 0x9f, 0x63, 0xb0, 0x42, 0x49, 0xd8, 0xd2, 0x1c, 0x4a, 0x55, 0x71, 0xd3, 0xa1,
	0xe7, 0x31, 0xa8, 0x86, 0x15, 0xaf, 0x0c, 0xe8, 0x8e, 0x9a, 0x45, 0xbd, 0x10, 0x0c, 0x2d, 0x58,
	0x0e, 0x0d, 0xf3, 0x25, 0x86, 0xc9, 0xd5, 0x6d, 0x5d, 0xd8, 0x36, 0x5a, 0x97, 0x42, 0x59, 0xb6,
	0xbd, 0xf0, 0x2c, 0x06, 0x16, 0x6a, 0xb5, 0x82, 0x9c, 0x72, 0xb9, 0x7b, 0x2c, 0xf7, 0xdc, 0xac,
	0x73, 0xcd, 0xee, 0x5b, 0xf4, 0x63, 0x0c, 0x6a, 0x58, 0x26, 0x80, 0x56, 0x86, 0xe9, 0x2e, 0xcf,
	0xd7, 0xff, 0x47, 0xe8, 0x70, 0x69, 0x3f, 0xfc, 0x26, 0x04, 0xbe, 0x41, 0xc7, 0x21, 0x39, 0xcd,
	0x0c, 0x57, 0x92, 0xec, 0xbd, 0xda, 0x7b, 0x33, 0x9e, 0x25, 0xc9, 0xa3, 0xeb, 0x33, 0x49, 0x1c,
	0xbe, 0xe8, 0x12, 0xbf, 0x7c, 0xb2, 0x18, 0x58, 0xf0, 0x12, 0x21, 0x9b, 0x52, 0x13, 0x10, 0x79,
	0xea, 0x9d, 0xe7, 0x31, 0xce, 0x79, 0x93, 0xed, 0xb2, 0x89, 0xd6, 0x9a, 0x8f, 0xac, 0x29, 0x1c,
	0xf1, 0x1c, 0x1d, 0xd8, 0xfc, 0xb8, 0x24, 0xfb, 0xde, 0xf8, 0x2e, 0xc6, 0xf8, 0xc3, 0x07, 0xfe,
	0x5d, 0x5a, 0x53, 0x80, 0xf1, 0x1f, 0xf4, 0x62, 0xb0, 0x67, 0xe4, 0x99, 0xb7, 0xcd, 0x62, 0x6c,
	0x57, 0x0e, 0x9c, 0xb7, 0xfb, 0xe9, 0xde, 0x5d, 0xf4, 0x2a, 0xf8, 0x1a, 0x8d, 0x7b, 0x81, 0x93,
	0x03, 0x2f, 0x9e, 0xc4, 0x88, 0x6f, 0x1c, 0x76, 0x1d, 0xc6, 0x64, 0xb5, 0xc8, 0x74, 0x67, 0xfc,
	0x0b, 0x8d, 0x7b, 0x3f, 0x27, 0x32, 0x8a, 0x9f, 0xd0, 0xa5, 0xc5, 0x7e, 0x7a, 0xca, 0x29, 0xd7,
	0xdd, 0x09, 0x7f, 0x45, 0xfb, 0x76, 0xf3, 0xc9, 0xa1, 0x57, 0xbd, 0x8d, 0x51, 0x5d, 0x80, 0xb9,
	0xb2, 0x03, 0xb1, 0x1e, 0x47, 0xe2, 0x7f, 0xe8, 0x64, 0xb8, 0x94, 0xe4, 0xf9, 0x8e, 0x21, 0xfe,
	0x6e, 0xb7, 0xd9, 0x2a, 0xc3, 0x3c, 0xda, 0x0a, 0xfe, 0x8b, 0x4e, 0x86, 0xff, 0x3c, 0xe4, 0xc8,
	0xcb, 0xa7, 0x51, 0x5f, 0xd4, 0x91, 0x8b, 0x06, 0x74, 0xee, 0xac, 0x5f, 0xb8, 0x38, 0x46, 0xc8,
	0xde, 0xa5, 0xee, 0xa8, 0xd9, 0x96, 0x90, 0x8d, 0x3c, 0x7d, 0xf6, 0x10, 0x00, 0x00, 0xff, 0xff,
	0x2c, 0xd7, 0x55, 0x6d, 0x5d, 0x05, 0x00, 0x00,
}
