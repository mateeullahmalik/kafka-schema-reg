// Code generated by protoc-gen-go. DO NOT EDIT.
// source: events/variantlabs/protobuf/l0/user_document_uploaded.proto

package events_variantlabs_protobuf_l0

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

type UserDocumentType int32

const (
	UserDocumentType_PASSPORT       UserDocumentType = 0
	UserDocumentType_LICENCE_DRIVER UserDocumentType = 1
)

var UserDocumentType_name = map[int32]string{
	0: "PASSPORT",
	1: "LICENCE_DRIVER",
}

var UserDocumentType_value = map[string]int32{
	"PASSPORT":       0,
	"LICENCE_DRIVER": 1,
}

func (x UserDocumentType) String() string {
	return proto.EnumName(UserDocumentType_name, int32(x))
}

func (UserDocumentType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d03a737b319c554b, []int{0}
}

type UserDocumentUploaded struct {
	Uid                  string                 `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Did                  string                 `protobuf:"bytes,2,opt,name=did,proto3" json:"did,omitempty"`
	Url                  string                 `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
	DocType              UserDocumentType       `protobuf:"varint,4,opt,name=doc_type,json=docType,proto3,enum=events.variantlabs.protobuf.l0.UserDocumentType" json:"doc_type,omitempty"`
	UploadedTs           *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=uploaded_ts,json=uploadedTs,proto3" json:"uploaded_ts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *UserDocumentUploaded) Reset()         { *m = UserDocumentUploaded{} }
func (m *UserDocumentUploaded) String() string { return proto.CompactTextString(m) }
func (*UserDocumentUploaded) ProtoMessage()    {}
func (*UserDocumentUploaded) Descriptor() ([]byte, []int) {
	return fileDescriptor_d03a737b319c554b, []int{0}
}

func (m *UserDocumentUploaded) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserDocumentUploaded.Unmarshal(m, b)
}
func (m *UserDocumentUploaded) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserDocumentUploaded.Marshal(b, m, deterministic)
}
func (m *UserDocumentUploaded) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserDocumentUploaded.Merge(m, src)
}
func (m *UserDocumentUploaded) XXX_Size() int {
	return xxx_messageInfo_UserDocumentUploaded.Size(m)
}
func (m *UserDocumentUploaded) XXX_DiscardUnknown() {
	xxx_messageInfo_UserDocumentUploaded.DiscardUnknown(m)
}

var xxx_messageInfo_UserDocumentUploaded proto.InternalMessageInfo

func (m *UserDocumentUploaded) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *UserDocumentUploaded) GetDid() string {
	if m != nil {
		return m.Did
	}
	return ""
}

func (m *UserDocumentUploaded) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *UserDocumentUploaded) GetDocType() UserDocumentType {
	if m != nil {
		return m.DocType
	}
	return UserDocumentType_PASSPORT
}

func (m *UserDocumentUploaded) GetUploadedTs() *timestamppb.Timestamp {
	if m != nil {
		return m.UploadedTs
	}
	return nil
}

func init() {
	proto.RegisterEnum("events.variantlabs.protobuf.l0.UserDocumentType", UserDocumentType_name, UserDocumentType_value)
	proto.RegisterType((*UserDocumentUploaded)(nil), "events.variantlabs.protobuf.l0.UserDocumentUploaded")
}

func init() {
	proto.RegisterFile("events/variantlabs/protobuf/l0/user_document_uploaded.proto", fileDescriptor_d03a737b319c554b)
}

var fileDescriptor_d03a737b319c554b = []byte{
	// 276 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x90, 0xcf, 0x4b, 0xc3, 0x30,
	0x1c, 0xc5, 0xad, 0xf3, 0xc7, 0xfc, 0x4e, 0x46, 0x09, 0x1e, 0xca, 0x0e, 0x22, 0x9e, 0xc4, 0x43,
	0x5a, 0xa7, 0xb7, 0x9d, 0x64, 0xeb, 0x61, 0x28, 0x3a, 0xb2, 0xce, 0x6b, 0x48, 0x4d, 0x1c, 0x85,
	0x74, 0x29, 0xf9, 0x31, 0xf0, 0x1f, 0xf5, 0xef, 0x31, 0x5d, 0x5b, 0x1c, 0x1e, 0x76, 0xfb, 0xf2,
	0x78, 0xef, 0xf3, 0x5e, 0x02, 0x13, 0xb1, 0x15, 0x1b, 0x6b, 0xe2, 0x2d, 0xd3, 0x05, 0xdb, 0x58,
	0xc9, 0x72, 0x13, 0x57, 0x5a, 0x59, 0x95, 0xbb, 0xaf, 0x58, 0x26, 0xb1, 0x33, 0x42, 0x53, 0xae,
	0x3e, 0x5d, 0xe9, 0x6d, 0xd4, 0x55, 0x52, 0x31, 0x2e, 0x38, 0xde, 0x59, 0xd0, 0x75, 0x13, 0xc6,
	0x7b, 0x61, 0xdc, 0x85, 0xb1, 0x4c, 0x46, 0x0f, 0x87, 0xe0, 0xce, 0x16, 0xd2, 0xc4, 0xb6, 0x28,
	0x85, 0xb1, 0xac, 0xac, 0x9a, 0xe0, 0xed, 0x4f, 0x00, 0x57, 0x2b, 0xdf, 0x39, 0x6b, 0x2b, 0x57,
	0x6d, 0x23, 0x0a, 0xa1, 0xe7, 0x0a, 0x1e, 0x05, 0x37, 0xc1, 0xdd, 0x05, 0xa9, 0xcf, 0x5a, 0xe1,
	0x5e, 0x39, 0x6e, 0x14, 0xde, 0x28, 0x4e, 0xcb, 0xa8, 0xd7, 0x7a, 0xb4, 0x44, 0x2f, 0xd0, 0xf7,
	0xe3, 0xa9, 0xfd, 0xae, 0x44, 0x74, 0xe2, 0xe5, 0xe1, 0x38, 0xc1, 0x87, 0x47, 0xe3, 0xfd, 0xf6,
	0xcc, 0xe7, 0xc8, 0xb9, 0x27, 0xd4, 0x07, 0x9a, 0xc0, 0xa0, 0xfb, 0x00, 0x6a, 0x4d, 0x74, 0xea,
	0x79, 0x83, 0xf1, 0x08, 0xaf, 0x95, 0x5a, 0x4b, 0xf1, 0xc7, 0xc8, 0xba, 0x27, 0x11, 0xe8, 0xec,
	0x99, 0xb9, 0x7f, 0x82, 0xf0, 0x3f, 0x19, 0x5d, 0x42, 0x7f, 0xf1, 0xbc, 0x5c, 0x2e, 0xde, 0x49,
	0x16, 0x1e, 0x21, 0x04, 0xc3, 0xd7, 0xf9, 0x34, 0x7d, 0x9b, 0xa6, 0x74, 0x46, 0xe6, 0x1f, 0x29,
	0x09, 0x83, 0xfc, 0x6c, 0x47, 0x7d, 0xfc, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xe2, 0xb9, 0x0f, 0x06,
	0xa7, 0x01, 0x00, 0x00,
}