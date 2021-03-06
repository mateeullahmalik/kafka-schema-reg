// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user_registered.proto

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

type UserAccountStatus int32

const (
	UserAccountStatus_ACTIVE   UserAccountStatus = 0
	UserAccountStatus_INACTIVE UserAccountStatus = 1
	UserAccountStatus_DISABLED UserAccountStatus = 2
	UserAccountStatus_BLOCKED  UserAccountStatus = 3
)

var UserAccountStatus_name = map[int32]string{
	0: "ACTIVE",
	1: "INACTIVE",
	2: "DISABLED",
	3: "BLOCKED",
}

var UserAccountStatus_value = map[string]int32{
	"ACTIVE":   0,
	"INACTIVE": 1,
	"DISABLED": 2,
	"BLOCKED":  3,
}

func (x UserAccountStatus) String() string {
	return proto.EnumName(UserAccountStatus_name, int32(x))
}

func (UserAccountStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ff42b8cde735cf43, []int{0}
}

type UserContactInfo struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Mobile               int32    `protobuf:"varint,2,opt,name=mobile,proto3" json:"mobile,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserContactInfo) Reset()         { *m = UserContactInfo{} }
func (m *UserContactInfo) String() string { return proto.CompactTextString(m) }
func (*UserContactInfo) ProtoMessage()    {}
func (*UserContactInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff42b8cde735cf43, []int{0}
}

func (m *UserContactInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserContactInfo.Unmarshal(m, b)
}
func (m *UserContactInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserContactInfo.Marshal(b, m, deterministic)
}
func (m *UserContactInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserContactInfo.Merge(m, src)
}
func (m *UserContactInfo) XXX_Size() int {
	return xxx_messageInfo_UserContactInfo.Size(m)
}
func (m *UserContactInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_UserContactInfo.DiscardUnknown(m)
}

var xxx_messageInfo_UserContactInfo proto.InternalMessageInfo

func (m *UserContactInfo) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *UserContactInfo) GetMobile() int32 {
	if m != nil {
		return m.Mobile
	}
	return 0
}

type UserAddressInfo struct {
	AddressString        string   `protobuf:"bytes,1,opt,name=address_string,json=addressString,proto3" json:"address_string,omitempty"`
	City                 string   `protobuf:"bytes,2,opt,name=city,proto3" json:"city,omitempty"`
	State                string   `protobuf:"bytes,3,opt,name=state,proto3" json:"state,omitempty"`
	Country              string   `protobuf:"bytes,4,opt,name=country,proto3" json:"country,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserAddressInfo) Reset()         { *m = UserAddressInfo{} }
func (m *UserAddressInfo) String() string { return proto.CompactTextString(m) }
func (*UserAddressInfo) ProtoMessage()    {}
func (*UserAddressInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff42b8cde735cf43, []int{1}
}

func (m *UserAddressInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserAddressInfo.Unmarshal(m, b)
}
func (m *UserAddressInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserAddressInfo.Marshal(b, m, deterministic)
}
func (m *UserAddressInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserAddressInfo.Merge(m, src)
}
func (m *UserAddressInfo) XXX_Size() int {
	return xxx_messageInfo_UserAddressInfo.Size(m)
}
func (m *UserAddressInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_UserAddressInfo.DiscardUnknown(m)
}

var xxx_messageInfo_UserAddressInfo proto.InternalMessageInfo

func (m *UserAddressInfo) GetAddressString() string {
	if m != nil {
		return m.AddressString
	}
	return ""
}

func (m *UserAddressInfo) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

func (m *UserAddressInfo) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *UserAddressInfo) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

type UserRegistered struct {
	Uid                  string               `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	FirstName            string               `protobuf:"bytes,2,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName             string               `protobuf:"bytes,3,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Username             string               `protobuf:"bytes,4,opt,name=username,proto3" json:"username,omitempty"`
	DateOfBirth          *timestamp.Timestamp `protobuf:"bytes,5,opt,name=date_of_birth,json=dateOfBirth,proto3" json:"date_of_birth,omitempty"`
	DateOfJoining        *timestamp.Timestamp `protobuf:"bytes,6,opt,name=date_of_joining,json=dateOfJoining,proto3" json:"date_of_joining,omitempty"`
	Contact              *UserContactInfo     `protobuf:"bytes,7,opt,name=contact,proto3" json:"contact,omitempty"`
	Address              *UserAddressInfo     `protobuf:"bytes,8,opt,name=address,proto3" json:"address,omitempty"`
	Language             string               `protobuf:"bytes,11,opt,name=language,proto3" json:"language,omitempty"`
	AccountStatus        UserAccountStatus    `protobuf:"varint,12,opt,name=account_status,json=accountStatus,proto3,enum=v1.UserAccountStatus" json:"account_status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *UserRegistered) Reset()         { *m = UserRegistered{} }
func (m *UserRegistered) String() string { return proto.CompactTextString(m) }
func (*UserRegistered) ProtoMessage()    {}
func (*UserRegistered) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff42b8cde735cf43, []int{2}
}

func (m *UserRegistered) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserRegistered.Unmarshal(m, b)
}
func (m *UserRegistered) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserRegistered.Marshal(b, m, deterministic)
}
func (m *UserRegistered) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserRegistered.Merge(m, src)
}
func (m *UserRegistered) XXX_Size() int {
	return xxx_messageInfo_UserRegistered.Size(m)
}
func (m *UserRegistered) XXX_DiscardUnknown() {
	xxx_messageInfo_UserRegistered.DiscardUnknown(m)
}

var xxx_messageInfo_UserRegistered proto.InternalMessageInfo

func (m *UserRegistered) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *UserRegistered) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *UserRegistered) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *UserRegistered) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *UserRegistered) GetDateOfBirth() *timestamp.Timestamp {
	if m != nil {
		return m.DateOfBirth
	}
	return nil
}

func (m *UserRegistered) GetDateOfJoining() *timestamp.Timestamp {
	if m != nil {
		return m.DateOfJoining
	}
	return nil
}

func (m *UserRegistered) GetContact() *UserContactInfo {
	if m != nil {
		return m.Contact
	}
	return nil
}

func (m *UserRegistered) GetAddress() *UserAddressInfo {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *UserRegistered) GetLanguage() string {
	if m != nil {
		return m.Language
	}
	return ""
}

func (m *UserRegistered) GetAccountStatus() UserAccountStatus {
	if m != nil {
		return m.AccountStatus
	}
	return UserAccountStatus_ACTIVE
}

func init() {
	proto.RegisterEnum("v1.UserAccountStatus", UserAccountStatus_name, UserAccountStatus_value)
	proto.RegisterType((*UserContactInfo)(nil), "v1.UserContactInfo")
	proto.RegisterType((*UserAddressInfo)(nil), "v1.UserAddressInfo")
	proto.RegisterType((*UserRegistered)(nil), "v1.UserRegistered")
}

func init() {
	proto.RegisterFile("user_registered.proto", fileDescriptor_ff42b8cde735cf43)
}

var fileDescriptor_ff42b8cde735cf43 = []byte{
	// 447 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x51, 0x5d, 0x6b, 0x13, 0x41,
	0x14, 0x35, 0x49, 0xf3, 0x75, 0xd3, 0xa4, 0xf1, 0x6a, 0x65, 0x88, 0x88, 0x12, 0x10, 0x44, 0x70,
	0x4b, 0xeb, 0xab, 0x28, 0x49, 0x5a, 0x30, 0x5a, 0x5a, 0xd8, 0x54, 0x5f, 0x97, 0xcd, 0xee, 0xec,
	0x3a, 0xb2, 0xbb, 0x53, 0x66, 0x66, 0x0b, 0x05, 0x7f, 0x97, 0xbf, 0xcf, 0xf9, 0xd8, 0x09, 0x51,
	0x1f, 0xfa, 0xb6, 0xe7, 0xdc, 0x73, 0xef, 0xd9, 0x39, 0x07, 0x8e, 0x6b, 0x49, 0x45, 0x24, 0x68,
	0xce, 0xa4, 0xa2, 0x82, 0xa6, 0xc1, 0xad, 0xe0, 0x8a, 0x63, 0xfb, 0xee, 0x74, 0xf6, 0x32, 0xe7,
	0x3c, 0x2f, 0xe8, 0x89, 0x65, 0xb6, 0x75, 0x76, 0xa2, 0x58, 0x49, 0xa5, 0x8a, 0xcb, 0x5b, 0x27,
	0x9a, 0x7f, 0x82, 0xa3, 0x6f, 0x7a, 0x7b, 0xc5, 0x2b, 0x15, 0x27, 0x6a, 0x5d, 0x65, 0x1c, 0x9f,
	0x42, 0x97, 0x96, 0x31, 0x2b, 0x48, 0xeb, 0x55, 0xeb, 0xcd, 0x30, 0x74, 0x00, 0x9f, 0x41, 0xaf,
	0xe4, 0x5b, 0x56, 0x50, 0xd2, 0xd6, 0x74, 0x37, 0x6c, 0xd0, 0xfc, 0x97, 0x3b, 0xb0, 0x48, 0x53,
	0x41, 0xa5, 0xb4, 0x07, 0x5e, 0xc3, 0x24, 0x76, 0x30, 0x92, 0x4a, 0xb0, 0x2a, 0x6f, 0x2e, 0x8d,
	0x1b, 0x76, 0x63, 0x49, 0x44, 0x38, 0x48, 0x98, 0xba, 0xb7, 0xf7, 0x86, 0xa1, 0xfd, 0x36, 0xde,
	0xfa, 0xef, 0x14, 0x25, 0x1d, 0xe7, 0x6d, 0x01, 0x12, 0xe8, 0x27, 0xbc, 0xae, 0x94, 0xb8, 0x27,
	0x07, 0x96, 0xf7, 0x70, 0xfe, 0xbb, 0x03, 0x13, 0x63, 0x1f, 0xee, 0x1e, 0x8f, 0x53, 0xe8, 0xd4,
	0x2c, 0x6d, 0x2c, 0xcd, 0x27, 0xbe, 0x00, 0xc8, 0x98, 0x90, 0x2a, 0xaa, 0xe2, 0x92, 0x36, 0x76,
	0x43, 0xcb, 0x5c, 0x69, 0x02, 0x9f, 0xc3, 0xb0, 0x88, 0xfd, 0xd4, 0xf9, 0x0e, 0x0c, 0x61, 0x87,
	0x33, 0x18, 0x98, 0x74, 0xed, 0xcc, 0x79, 0xef, 0x30, 0x7e, 0x84, 0x71, 0xaa, 0x7f, 0x2f, 0xe2,
	0x59, 0xb4, 0x65, 0x42, 0xfd, 0x20, 0x5d, 0x2d, 0x18, 0x9d, 0xcd, 0x02, 0x17, 0x7a, 0xe0, 0x43,
	0x0f, 0x6e, 0x7c, 0xe8, 0xe1, 0xc8, 0x2c, 0x5c, 0x67, 0x4b, 0x23, 0xc7, 0x25, 0x1c, 0xf9, 0xfd,
	0x9f, 0x9c, 0x55, 0x26, 0xa8, 0xde, 0x83, 0x17, 0xc6, 0xee, 0xc2, 0x17, 0xb7, 0x80, 0xef, 0x4c,
	0x34, 0xb6, 0x3b, 0xd2, 0xb7, 0xbb, 0x4f, 0x82, 0xbb, 0xd3, 0xe0, 0x9f, 0x4a, 0x43, 0xaf, 0x31,
	0xf2, 0xa6, 0x04, 0x32, 0xf8, 0x5b, 0xbe, 0x57, 0x60, 0xe8, 0x35, 0xe6, 0xf5, 0x45, 0x5c, 0xe5,
	0x75, 0x9c, 0x53, 0x32, 0xf2, 0xc9, 0x38, 0x8c, 0x1f, 0x74, 0xcb, 0x89, 0xed, 0x21, 0x32, 0x2d,
	0xd5, 0x92, 0x1c, 0x6a, 0xc5, 0xe4, 0xec, 0x78, 0x77, 0xd1, 0x4d, 0x37, 0x76, 0xa8, 0xcb, 0xdf,
	0x87, 0x6f, 0x3f, 0xc3, 0xe3, 0xff, 0x34, 0x08, 0xd0, 0x5b, 0xac, 0x6e, 0xd6, 0xdf, 0x2f, 0xa6,
	0x8f, 0xf0, 0x10, 0x06, 0xeb, 0xab, 0x06, 0xb5, 0x0c, 0x3a, 0x5f, 0x6f, 0x16, 0xcb, 0xcb, 0x8b,
	0xf3, 0x69, 0x1b, 0x47, 0xd0, 0x5f, 0x5e, 0x5e, 0xaf, 0xbe, 0x6a, 0xd0, 0xd9, 0xf6, 0x6c, 0x48,
	0xef, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0xf1, 0x7b, 0x0b, 0xb8, 0x06, 0x03, 0x00, 0x00,
}
