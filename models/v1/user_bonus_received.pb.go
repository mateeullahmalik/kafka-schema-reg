// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user_bonus_received.proto

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

type UserBonusReceived struct {
	Uid                  string               `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	ReceivedAtTs         *timestamp.Timestamp `protobuf:"bytes,2,opt,name=received_at_ts,json=receivedAtTs,proto3" json:"received_at_ts,omitempty"`
	Amount               float32              `protobuf:"fixed32,3,opt,name=amount,proto3" json:"amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *UserBonusReceived) Reset()         { *m = UserBonusReceived{} }
func (m *UserBonusReceived) String() string { return proto.CompactTextString(m) }
func (*UserBonusReceived) ProtoMessage()    {}
func (*UserBonusReceived) Descriptor() ([]byte, []int) {
	return fileDescriptor_f3544bfddeec6abd, []int{0}
}

func (m *UserBonusReceived) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserBonusReceived.Unmarshal(m, b)
}
func (m *UserBonusReceived) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserBonusReceived.Marshal(b, m, deterministic)
}
func (m *UserBonusReceived) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserBonusReceived.Merge(m, src)
}
func (m *UserBonusReceived) XXX_Size() int {
	return xxx_messageInfo_UserBonusReceived.Size(m)
}
func (m *UserBonusReceived) XXX_DiscardUnknown() {
	xxx_messageInfo_UserBonusReceived.DiscardUnknown(m)
}

var xxx_messageInfo_UserBonusReceived proto.InternalMessageInfo

func (m *UserBonusReceived) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *UserBonusReceived) GetReceivedAtTs() *timestamp.Timestamp {
	if m != nil {
		return m.ReceivedAtTs
	}
	return nil
}

func (m *UserBonusReceived) GetAmount() float32 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func init() {
	proto.RegisterType((*UserBonusReceived)(nil), "v1.UserBonusReceived")
}

func init() {
	proto.RegisterFile("user_bonus_received.proto", fileDescriptor_f3544bfddeec6abd)
}

var fileDescriptor_f3544bfddeec6abd = []byte{
	// 170 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x92, 0x2c, 0x2d, 0x4e, 0x2d,
	0x8a, 0x4f, 0xca, 0xcf, 0x2b, 0x2d, 0x8e, 0x2f, 0x4a, 0x4d, 0x4e, 0xcd, 0x2c, 0x4b, 0x4d, 0xd1,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x33, 0x94, 0x92, 0x4f, 0xcf, 0xcf, 0x4f, 0xcf,
	0x49, 0xd5, 0x07, 0x8b, 0x24, 0x95, 0xa6, 0xe9, 0x97, 0x64, 0xe6, 0xa6, 0x16, 0x97, 0x24, 0xe6,
	0x16, 0x40, 0x14, 0x29, 0xd5, 0x73, 0x09, 0x86, 0x02, 0x4d, 0x70, 0x02, 0x19, 0x10, 0x04, 0xd5,
	0x2f, 0x24, 0xc0, 0xc5, 0x5c, 0x9a, 0x99, 0x22, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x19, 0x04, 0x62,
	0x0a, 0x39, 0x70, 0xf1, 0xc1, 0x4c, 0x8f, 0x4f, 0x2c, 0x89, 0x2f, 0x29, 0x96, 0x60, 0x02, 0x4a,
	0x72, 0x1b, 0x49, 0xe9, 0x41, 0x2c, 0xd0, 0x83, 0x59, 0xa0, 0x17, 0x02, 0xb3, 0x20, 0x88, 0x07,
	0xa6, 0xc3, 0xb1, 0x24, 0xa4, 0x58, 0x48, 0x8c, 0x8b, 0x2d, 0x31, 0x37, 0xbf, 0x34, 0xaf, 0x44,
	0x82, 0x19, 0xa8, 0x93, 0x29, 0x08, 0xca, 0x4b, 0x62, 0x03, 0xeb, 0x34, 0x06, 0x04, 0x00, 0x00,
	0xff, 0xff, 0x0a, 0xec, 0x1a, 0x7c, 0xc9, 0x00, 0x00, 0x00,
}
