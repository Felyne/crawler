// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common.proto

package common

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

//存储的一条记录
type Item struct {
	Url                  string   `protobuf:"bytes,1,opt,name=Url,proto3" json:"Url,omitempty"`
	Type                 string   `protobuf:"bytes,2,opt,name=Type,proto3" json:"Type,omitempty"`
	Id                   string   `protobuf:"bytes,3,opt,name=Id,proto3" json:"Id,omitempty"`
	Payload              []byte   `protobuf:"bytes,4,opt,name=Payload,proto3" json:"Payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Item) Reset()         { *m = Item{} }
func (m *Item) String() string { return proto.CompactTextString(m) }
func (*Item) ProtoMessage()    {}
func (*Item) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{0}
}

func (m *Item) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Item.Unmarshal(m, b)
}
func (m *Item) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Item.Marshal(b, m, deterministic)
}
func (m *Item) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Item.Merge(m, src)
}
func (m *Item) XXX_Size() int {
	return xxx_messageInfo_Item.Size(m)
}
func (m *Item) XXX_DiscardUnknown() {
	xxx_messageInfo_Item.DiscardUnknown(m)
}

var xxx_messageInfo_Item proto.InternalMessageInfo

func (m *Item) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *Item) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Item) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Item) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func init() {
	proto.RegisterType((*Item)(nil), "common.Item")
}

func init() { proto.RegisterFile("common.proto", fileDescriptor_555bd8c177793206) }

var fileDescriptor_555bd8c177793206 = []byte{
	// 175 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x34, 0x8e, 0xbf, 0xcb, 0xc2, 0x30,
	0x14, 0x00, 0xe9, 0x0f, 0xfa, 0xf1, 0x85, 0x22, 0x92, 0x29, 0x63, 0x71, 0xea, 0xd4, 0x0c, 0xae,
	0xe2, 0xe0, 0x20, 0x74, 0x93, 0xa2, 0x0e, 0x6e, 0x6d, 0xf2, 0xa8, 0x81, 0xa4, 0xaf, 0xbc, 0xa6,
	0x4a, 0xff, 0x7b, 0x31, 0xc5, 0xed, 0xee, 0xa6, 0x63, 0xb9, 0x42, 0xe7, 0x70, 0xa8, 0x46, 0x42,
	0x8f, 0x3c, 0x5b, 0x6d, 0x77, 0x67, 0x69, 0xed, 0xc1, 0xf1, 0x2d, 0x4b, 0x6e, 0x64, 0x45, 0x54,
	0x44, 0xe5, 0x7f, 0xf3, 0x45, 0xce, 0x59, 0x7a, 0x5d, 0x46, 0x10, 0x71, 0x48, 0x81, 0xf9, 0x86,
	0xc5, 0xb5, 0x16, 0x49, 0x28, 0x71, 0xad, 0xb9, 0x60, 0x7f, 0x97, 0x76, 0xb1, 0xd8, 0x6a, 0x91,
	0x16, 0x51, 0x99, 0x37, 0x3f, 0x3d, 0x1d, 0x1f, 0x87, 0xde, 0xf8, 0xe7, 0xdc, 0x55, 0x0a, 0x9d,
	0x3c, 0x83, 0x5d, 0x06, 0x90, 0x8a, 0xda, 0xb7, 0x05, 0x92, 0xda, 0x4c, 0x9e, 0x4c, 0x37, 0x7b,
	0xd0, 0x72, 0x02, 0x7a, 0x19, 0x05, 0xb2, 0x47, 0x67, 0x14, 0xa1, 0x5c, 0xbf, 0xba, 0x2c, 0x6c,
	0xee, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xcb, 0xf3, 0x9b, 0xbc, 0xb6, 0x00, 0x00, 0x00,
}
