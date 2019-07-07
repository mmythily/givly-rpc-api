// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rpc/merchant/merchant.proto

package merchant

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

type Merchant struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Storename            string               `protobuf:"bytes,2,opt,name=storename,proto3" json:"storename,omitempty"`
	Wallet               string               `protobuf:"bytes,3,opt,name=wallet,proto3" json:"wallet,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	LastModified         *timestamp.Timestamp `protobuf:"bytes,5,opt,name=last_modified,json=lastModified,proto3" json:"last_modified,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Merchant) Reset()         { *m = Merchant{} }
func (m *Merchant) String() string { return proto.CompactTextString(m) }
func (*Merchant) ProtoMessage()    {}
func (*Merchant) Descriptor() ([]byte, []int) {
	return fileDescriptor_58ecbaafc626f209, []int{0}
}

func (m *Merchant) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Merchant.Unmarshal(m, b)
}
func (m *Merchant) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Merchant.Marshal(b, m, deterministic)
}
func (m *Merchant) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Merchant.Merge(m, src)
}
func (m *Merchant) XXX_Size() int {
	return xxx_messageInfo_Merchant.Size(m)
}
func (m *Merchant) XXX_DiscardUnknown() {
	xxx_messageInfo_Merchant.DiscardUnknown(m)
}

var xxx_messageInfo_Merchant proto.InternalMessageInfo

func (m *Merchant) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Merchant) GetStorename() string {
	if m != nil {
		return m.Storename
	}
	return ""
}

func (m *Merchant) GetWallet() string {
	if m != nil {
		return m.Wallet
	}
	return ""
}

func (m *Merchant) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Merchant) GetLastModified() *timestamp.Timestamp {
	if m != nil {
		return m.LastModified
	}
	return nil
}

func init() {
	proto.RegisterType((*Merchant)(nil), "merchant.Merchant")
}

func init() { proto.RegisterFile("rpc/merchant/merchant.proto", fileDescriptor_58ecbaafc626f209) }

var fileDescriptor_58ecbaafc626f209 = []byte{
	// 242 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0xbf, 0x4f, 0xfb, 0x30,
	0x10, 0xc5, 0x95, 0x7c, 0xbf, 0x54, 0x8d, 0xf9, 0x31, 0x78, 0x40, 0x56, 0x41, 0xa2, 0x62, 0xea,
	0x52, 0x1b, 0xc1, 0xc4, 0x84, 0x60, 0xef, 0x52, 0x31, 0xb1, 0x54, 0x17, 0xfb, 0x9a, 0x9e, 0x64,
	0xd7, 0x96, 0x73, 0x01, 0xf1, 0x3f, 0xf2, 0x47, 0x21, 0xd2, 0x04, 0xd8, 0xd8, 0xee, 0xbd, 0xfb,
	0xbc, 0xd3, 0xd3, 0x89, 0x8b, 0x9c, 0xac, 0x09, 0x98, 0xed, 0x0e, 0xf6, 0xfc, 0x3d, 0xe8, 0x94,
	0x23, 0x47, 0x39, 0x1d, 0xf5, 0xec, 0xaa, 0x89, 0xb1, 0xf1, 0x68, 0x7a, 0xbf, 0xee, 0xb6, 0x86,
	0x29, 0x60, 0xcb, 0x10, 0xd2, 0x01, 0xbd, 0xfe, 0x28, 0xc4, 0x74, 0x35, 0xd0, 0xf2, 0x4c, 0x94,
	0xe4, 0x54, 0x31, 0x2f, 0x16, 0xd5, 0xba, 0x24, 0x27, 0x2f, 0x45, 0xd5, 0x72, 0xcc, 0xb8, 0x87,
	0x80, 0xaa, 0xec, 0xed, 0x1f, 0x43, 0x9e, 0x8b, 0xc9, 0x1b, 0x78, 0x8f, 0xac, 0xfe, 0xf5, 0xab,
	0x41, 0xc9, 0x7b, 0x21, 0x6c, 0x46, 0x60, 0x74, 0x1b, 0x60, 0xf5, 0x7f, 0x5e, 0x2c, 0x8e, 0x6f,
	0x67, 0xfa, 0x50, 0x44, 0x8f, 0x45, 0xf4, 0xf3, 0x58, 0x64, 0x5d, 0x0d, 0xf4, 0x23, 0xcb, 0x07,
	0x71, 0xea, 0xa1, 0xe5, 0x4d, 0x88, 0x8e, 0xb6, 0x84, 0x4e, 0x1d, 0xfd, 0x99, 0x3e, 0xf9, 0x0a,
	0xac, 0x06, 0xfe, 0xe9, 0xe6, 0x45, 0x37, 0xc4, 0xbb, 0xae, 0xd6, 0x36, 0x06, 0x93, 0xbb, 0xd0,
	0x66, 0x08, 0x64, 0x1a, 0x7a, 0xf5, 0xef, 0xcb, 0x9c, 0xec, 0x12, 0x12, 0x99, 0xdf, 0xaf, 0xab,
	0x27, 0xfd, 0xcd, 0xbb, 0xcf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x36, 0xbe, 0xac, 0x51, 0x51, 0x01,
	0x00, 0x00,
}
