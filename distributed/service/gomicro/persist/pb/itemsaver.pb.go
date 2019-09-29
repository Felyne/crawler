// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/itemsaver.proto

package pb

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

type ServiceName int32

const (
	ServiceName_ITEMSAVER_SERVICE ServiceName = 0
)

var ServiceName_name = map[int32]string{
	0: "ITEMSAVER_SERVICE",
}

var ServiceName_value = map[string]int32{
	"ITEMSAVER_SERVICE": 0,
}

func (x ServiceName) String() string {
	return proto.EnumName(ServiceName_name, int32(x))
}

func (ServiceName) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0e543f287d81f699, []int{0}
}

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
	return fileDescriptor_0e543f287d81f699, []int{0}
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

type Resp struct {
	Result               string   `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Resp) Reset()         { *m = Resp{} }
func (m *Resp) String() string { return proto.CompactTextString(m) }
func (*Resp) ProtoMessage()    {}
func (*Resp) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e543f287d81f699, []int{1}
}

func (m *Resp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Resp.Unmarshal(m, b)
}
func (m *Resp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Resp.Marshal(b, m, deterministic)
}
func (m *Resp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Resp.Merge(m, src)
}
func (m *Resp) XXX_Size() int {
	return xxx_messageInfo_Resp.Size(m)
}
func (m *Resp) XXX_DiscardUnknown() {
	xxx_messageInfo_Resp.DiscardUnknown(m)
}

var xxx_messageInfo_Resp proto.InternalMessageInfo

func (m *Resp) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

func init() {
	proto.RegisterEnum("pb.ServiceName", ServiceName_name, ServiceName_value)
	proto.RegisterType((*Item)(nil), "pb.Item")
	proto.RegisterType((*Resp)(nil), "pb.Resp")
}

func init() { proto.RegisterFile("pb/itemsaver.proto", fileDescriptor_0e543f287d81f699) }

var fileDescriptor_0e543f287d81f699 = []byte{
	// 211 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x34, 0x8f, 0x41, 0x4b, 0xc4, 0x30,
	0x10, 0x85, 0xb7, 0xd9, 0xb0, 0xba, 0xa3, 0xc8, 0x3a, 0xa0, 0x04, 0x11, 0x59, 0x8a, 0x87, 0xd5,
	0x43, 0x05, 0xfd, 0x05, 0x22, 0x39, 0xe4, 0xa0, 0x48, 0xb2, 0xf6, 0x2a, 0x89, 0x9d, 0x43, 0xa1,
	0xa5, 0x21, 0x8d, 0x85, 0xfe, 0x7b, 0x49, 0xa8, 0xb7, 0xef, 0x7d, 0x87, 0x99, 0xf7, 0x00, 0xbd,
	0x7b, 0x6a, 0x23, 0xf5, 0xa3, 0x9d, 0x28, 0x54, 0x3e, 0x0c, 0x71, 0x40, 0xe6, 0x5d, 0x59, 0x03,
	0x57, 0x91, 0x7a, 0xdc, 0xc1, 0xfa, 0x2b, 0x74, 0xa2, 0xd8, 0x17, 0x87, 0xad, 0x4e, 0x88, 0x08,
	0xfc, 0x38, 0x7b, 0x12, 0x2c, 0xab, 0xcc, 0x78, 0x01, 0x4c, 0x35, 0x62, 0x9d, 0x0d, 0x53, 0x0d,
	0x0a, 0x38, 0xf9, 0xb4, 0x73, 0x37, 0xd8, 0x46, 0xf0, 0x7d, 0x71, 0x38, 0xd7, 0xff, 0xb1, 0xbc,
	0x03, 0xae, 0x69, 0xf4, 0x78, 0x0d, 0x9b, 0x40, 0xe3, 0x6f, 0x17, 0x97, 0xd3, 0x4b, 0x7a, 0xbc,
	0x87, 0x33, 0x43, 0x61, 0x6a, 0x7f, 0xe8, 0xc3, 0xf6, 0x84, 0x57, 0x70, 0xa9, 0x8e, 0xf2, 0xdd,
	0xbc, 0xd6, 0x52, 0x7f, 0x1b, 0xa9, 0x6b, 0xf5, 0x26, 0x77, 0xab, 0xe7, 0x07, 0xd8, 0xa6, 0x76,
	0x26, 0x95, 0xc6, 0x5b, 0xe0, 0x09, 0xf0, 0xb4, 0xf2, 0xae, 0x4a, 0xfa, 0x26, 0x53, 0x7a, 0x53,
	0xae, 0xdc, 0x26, 0x6f, 0x7a, 0xf9, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x41, 0x70, 0xc6, 0xbf, 0xe9,
	0x00, 0x00, 0x00,
}
