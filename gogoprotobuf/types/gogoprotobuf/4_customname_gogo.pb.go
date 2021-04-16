// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: 4_customname_gogo.proto

package gogoprotobuf

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Test4 struct {
	MyName               string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	MyNum                float64  `protobuf:"fixed64,2,opt,name=num,proto3" json:"num,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Test4) Reset()         { *m = Test4{} }
func (m *Test4) String() string { return proto.CompactTextString(m) }
func (*Test4) ProtoMessage()    {}
func (*Test4) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc4a553b09c21d97, []int{0}
}
func (m *Test4) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Test4.Unmarshal(m, b)
}
func (m *Test4) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Test4.Marshal(b, m, deterministic)
}
func (m *Test4) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Test4.Merge(m, src)
}
func (m *Test4) XXX_Size() int {
	return xxx_messageInfo_Test4.Size(m)
}
func (m *Test4) XXX_DiscardUnknown() {
	xxx_messageInfo_Test4.DiscardUnknown(m)
}

var xxx_messageInfo_Test4 proto.InternalMessageInfo

func (m *Test4) GetMyName() string {
	if m != nil {
		return m.MyName
	}
	return ""
}

func (m *Test4) GetMyNum() float64 {
	if m != nil {
		return m.MyNum
	}
	return 0
}

func init() {
	proto.RegisterType((*Test4)(nil), "types.Test4")
}

func init() { proto.RegisterFile("4_customname_gogo.proto", fileDescriptor_fc4a553b09c21d97) }

var fileDescriptor_fc4a553b09c21d97 = []byte{
	// 164 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x37, 0x89, 0x4f, 0x2e,
	0x2d, 0x2e, 0xc9, 0xcf, 0xcd, 0x4b, 0xcc, 0x4d, 0x8d, 0x4f, 0xcf, 0x4f, 0xcf, 0xd7, 0x2b, 0x28,
	0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2d, 0xa9, 0x2c, 0x48, 0x2d, 0x96, 0xd2, 0x4d, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x07, 0xc9, 0xea, 0x83, 0x65, 0x93, 0x4a, 0xd3, 0xc0,
	0x3c, 0x30, 0x47, 0x1f, 0xa1, 0x4b, 0xc9, 0x85, 0x8b, 0x35, 0x24, 0xb5, 0xb8, 0xc4, 0x44, 0x48,
	0x8e, 0x8b, 0x05, 0x64, 0xa2, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0xa7, 0x13, 0xd7, 0xa3, 0x7b, 0xf2,
	0x6c, 0xbe, 0x95, 0x7e, 0x89, 0xb9, 0xa9, 0x41, 0x60, 0x71, 0x21, 0x69, 0x2e, 0xe6, 0xbc, 0xd2,
	0x5c, 0x09, 0x26, 0x05, 0x46, 0x0d, 0x46, 0x27, 0xce, 0x47, 0xf7, 0xe4, 0x59, 0x7d, 0x2b, 0xfd,
	0x4a, 0x73, 0x83, 0x40, 0xa2, 0x4e, 0x82, 0x51, 0xfc, 0x7a, 0xfa, 0xd6, 0x70, 0x0b, 0x92, 0x4a,
	0xd3, 0x92, 0xd8, 0xc0, 0x2c, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x1c, 0x0b, 0x7a, 0xbc,
	0xb0, 0x00, 0x00, 0x00,
}
