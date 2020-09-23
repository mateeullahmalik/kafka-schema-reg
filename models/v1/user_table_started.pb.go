// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user_table_started.proto

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

type UserTableStarted struct {
	Uid                  string               `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Tid                  string               `protobuf:"bytes,2,opt,name=tid,proto3" json:"tid,omitempty"`
	StartedAtTs          *timestamp.Timestamp `protobuf:"bytes,3,opt,name=started_at_ts,json=startedAtTs,proto3" json:"started_at_ts,omitempty"`
	Players              []string             `protobuf:"bytes,4,rep,name=players,proto3" json:"players,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *UserTableStarted) Reset()         { *m = UserTableStarted{} }
func (m *UserTableStarted) String() string { return proto.CompactTextString(m) }
func (*UserTableStarted) ProtoMessage()    {}
func (*UserTableStarted) Descriptor() ([]byte, []int) {
	return fileDescriptor_b0ec2a1d282d3f4c, []int{0}
}

func (m *UserTableStarted) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserTableStarted.Unmarshal(m, b)
}
func (m *UserTableStarted) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserTableStarted.Marshal(b, m, deterministic)
}
func (m *UserTableStarted) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserTableStarted.Merge(m, src)
}
func (m *UserTableStarted) XXX_Size() int {
	return xxx_messageInfo_UserTableStarted.Size(m)
}
func (m *UserTableStarted) XXX_DiscardUnknown() {
	xxx_messageInfo_UserTableStarted.DiscardUnknown(m)
}

var xxx_messageInfo_UserTableStarted proto.InternalMessageInfo

func (m *UserTableStarted) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *UserTableStarted) GetTid() string {
	if m != nil {
		return m.Tid
	}
	return ""
}

func (m *UserTableStarted) GetStartedAtTs() *timestamp.Timestamp {
	if m != nil {
		return m.StartedAtTs
	}
	return nil
}

func (m *UserTableStarted) GetPlayers() []string {
	if m != nil {
		return m.Players
	}
	return nil
}

func init() {
	proto.RegisterType((*UserTableStarted)(nil), "v1.UserTableStarted")
}

func init() {
	proto.RegisterFile("user_table_started.proto", fileDescriptor_b0ec2a1d282d3f4c)
}

var fileDescriptor_b0ec2a1d282d3f4c = []byte{
	// 180 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x92, 0x28, 0x2d, 0x4e, 0x2d,
	0x8a, 0x2f, 0x49, 0x4c, 0xca, 0x49, 0x8d, 0x2f, 0x2e, 0x49, 0x2c, 0x2a, 0x49, 0x4d, 0xd1, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x33, 0x94, 0x92, 0x4f, 0xcf, 0xcf, 0x4f, 0xcf, 0x49,
	0xd5, 0x07, 0x8b, 0x24, 0x95, 0xa6, 0xe9, 0x97, 0x64, 0xe6, 0xa6, 0x02, 0x95, 0xe5, 0x16, 0x40,
	0x14, 0x29, 0x4d, 0x60, 0xe4, 0x12, 0x08, 0x05, 0x9a, 0x10, 0x02, 0x32, 0x20, 0x18, 0xa2, 0x5f,
	0x48, 0x80, 0x8b, 0xb9, 0x34, 0x33, 0x45, 0x82, 0x51, 0x81, 0x51, 0x83, 0x33, 0x08, 0xc4, 0x04,
	0x89, 0x94, 0x00, 0x45, 0x98, 0x20, 0x22, 0x40, 0xa6, 0x90, 0x1d, 0x17, 0x2f, 0xd4, 0xba, 0xf8,
	0xc4, 0x92, 0xf8, 0x92, 0x62, 0x09, 0x66, 0xa0, 0x1c, 0xb7, 0x91, 0x94, 0x1e, 0xc4, 0x46, 0x3d,
	0x98, 0x8d, 0x7a, 0x21, 0x30, 0x1b, 0x83, 0xb8, 0xa1, 0x1a, 0x1c, 0x4b, 0x42, 0x8a, 0x85, 0x24,
	0xb8, 0xd8, 0x0b, 0x72, 0x12, 0x2b, 0x53, 0x8b, 0x8a, 0x25, 0x58, 0x14, 0x98, 0x81, 0xa6, 0xc2,
	0xb8, 0x49, 0x6c, 0x60, 0xad, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x22, 0x36, 0x73, 0xdb,
	0xda, 0x00, 0x00, 0x00,
}
