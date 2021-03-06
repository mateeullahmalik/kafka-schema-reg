// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user_money_withdrawn.proto

package v1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	math "math"
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

type UserMoneyWithdrawn struct {
	Uid                  string               `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	WithdrawAtTs         *timestamp.Timestamp `protobuf:"bytes,2,opt,name=withdraw_at_ts,json=withdrawAtTs,proto3" json:"withdraw_at_ts,omitempty"`
	Amount               float32              `protobuf:"fixed32,3,opt,name=amount,proto3" json:"amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *UserMoneyWithdrawn) Reset()         { *m = UserMoneyWithdrawn{} }
func (m *UserMoneyWithdrawn) String() string { return proto.CompactTextString(m) }
func (*UserMoneyWithdrawn) ProtoMessage()    {}
func (*UserMoneyWithdrawn) Descriptor() ([]byte, []int) {
	return fileDescriptor_c0299ec1fe53c25d, []int{0}
}

func (m *UserMoneyWithdrawn) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserMoneyWithdrawn.Unmarshal(m, b)
}
func (m *UserMoneyWithdrawn) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserMoneyWithdrawn.Marshal(b, m, deterministic)
}
func (m *UserMoneyWithdrawn) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserMoneyWithdrawn.Merge(m, src)
}
func (m *UserMoneyWithdrawn) XXX_Size() int {
	return xxx_messageInfo_UserMoneyWithdrawn.Size(m)
}
func (m *UserMoneyWithdrawn) XXX_DiscardUnknown() {
	xxx_messageInfo_UserMoneyWithdrawn.DiscardUnknown(m)
}

var xxx_messageInfo_UserMoneyWithdrawn proto.InternalMessageInfo

func (m *UserMoneyWithdrawn) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *UserMoneyWithdrawn) GetWithdrawAtTs() *timestamp.Timestamp {
	if m != nil {
		return m.WithdrawAtTs
	}
	return nil
}

func (m *UserMoneyWithdrawn) GetAmount() float32 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func init() {
	proto.RegisterType((*UserMoneyWithdrawn)(nil), "v1.UserMoneyWithdrawn")
}

func init() {
	proto.RegisterFile("user_money_withdrawn.proto", fileDescriptor_c0299ec1fe53c25d)
}

var fileDescriptor_c0299ec1fe53c25d = []byte{
	// 172 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x92, 0x2a, 0x2d, 0x4e, 0x2d,
	0x8a, 0xcf, 0xcd, 0xcf, 0x4b, 0xad, 0x8c, 0x2f, 0xcf, 0x2c, 0xc9, 0x48, 0x29, 0x4a, 0x2c, 0xcf,
	0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x33, 0x94, 0x92, 0x4f, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0x07, 0x8b, 0x24, 0x95, 0xa6, 0xe9, 0x97, 0x64, 0xe6, 0xa6, 0x16, 0x97, 0x24,
	0xe6, 0x16, 0x40, 0x14, 0x29, 0x35, 0x30, 0x72, 0x09, 0x85, 0x02, 0xcd, 0xf0, 0x05, 0x19, 0x11,
	0x0e, 0x33, 0x41, 0x48, 0x80, 0x8b, 0xb9, 0x34, 0x33, 0x45, 0x82, 0x51, 0x81, 0x51, 0x83, 0x33,
	0x08, 0xc4, 0x14, 0x72, 0xe0, 0xe2, 0x83, 0x59, 0x10, 0x9f, 0x58, 0x12, 0x5f, 0x52, 0x2c, 0xc1,
	0x04, 0x94, 0xe4, 0x36, 0x92, 0xd2, 0x83, 0x58, 0xa1, 0x07, 0xb3, 0x42, 0x2f, 0x04, 0x66, 0x45,
	0x10, 0x0f, 0x4c, 0x87, 0x63, 0x49, 0x48, 0xb1, 0x90, 0x18, 0x17, 0x5b, 0x62, 0x6e, 0x7e, 0x69,
	0x5e, 0x89, 0x04, 0x33, 0x50, 0x27, 0x53, 0x10, 0x94, 0x97, 0xc4, 0x06, 0xd6, 0x69, 0x0c, 0x08,
	0x00, 0x00, 0xff, 0xff, 0x4a, 0xef, 0xa9, 0x6a, 0xcc, 0x00, 0x00, 0x00,
}
