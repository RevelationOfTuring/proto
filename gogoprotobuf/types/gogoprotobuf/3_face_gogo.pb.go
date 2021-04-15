// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: 3_face_gogo.proto

package gogoprotobuf

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	github_com_gogo_protobuf_proto "github.com/gogo/protobuf/proto"
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

type Test struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	Num                  float64  `protobuf:"fixed64,2,opt,name=num,proto3" json:"num,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Test) Reset()         { *m = Test{} }
func (m *Test) String() string { return proto.CompactTextString(m) }
func (*Test) ProtoMessage()    {}
func (*Test) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc17d42709de042b, []int{0}
}
func (m *Test) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Test.Unmarshal(m, b)
}
func (m *Test) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Test.Marshal(b, m, deterministic)
}
func (m *Test) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Test.Merge(m, src)
}
func (m *Test) XXX_Size() int {
	return xxx_messageInfo_Test.Size(m)
}
func (m *Test) XXX_DiscardUnknown() {
	xxx_messageInfo_Test.DiscardUnknown(m)
}

var xxx_messageInfo_Test proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Test)(nil), "types.Test")
}

func init() { proto.RegisterFile("3_face_gogo.proto", fileDescriptor_fc17d42709de042b) }

var fileDescriptor_fc17d42709de042b = []byte{
	// 142 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x34, 0x8e, 0x4f, 0x4b,
	0x4c, 0x4e, 0x8d, 0x4f, 0xcf, 0x4f, 0xcf, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2d,
	0xa9, 0x2c, 0x48, 0x2d, 0x96, 0xd2, 0x4d, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf,
	0xd5, 0x07, 0xc9, 0xea, 0x83, 0x65, 0x93, 0x4a, 0xd3, 0xc0, 0x3c, 0x30, 0x47, 0x1f, 0xa1, 0x4b,
	0xc9, 0x84, 0x8b, 0x25, 0x24, 0xb5, 0xb8, 0x44, 0x48, 0x80, 0x8b, 0x39, 0xb7, 0x38, 0x5d, 0x82,
	0x51, 0x81, 0x51, 0x83, 0x33, 0x08, 0xc4, 0x04, 0x89, 0xe4, 0x95, 0xe6, 0x4a, 0x30, 0x29, 0x30,
	0x6a, 0x30, 0x06, 0x81, 0x98, 0x56, 0x1c, 0x1d, 0x0b, 0xe4, 0x19, 0x56, 0x2c, 0x90, 0x67, 0x74,
	0x12, 0x8c, 0xe2, 0xd7, 0xd3, 0xb7, 0x86, 0x1b, 0x98, 0x54, 0x9a, 0x96, 0xc4, 0x06, 0x66, 0x19,
	0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x59, 0x98, 0x7f, 0x24, 0x9a, 0x00, 0x00, 0x00,
}

type TestFace interface {
	Proto() github_com_gogo_protobuf_proto.Message
	GetMsg() string
	GetNum() float64
}

func (this *Test) Proto() github_com_gogo_protobuf_proto.Message {
	return this
}

func (this *Test) TestProto() github_com_gogo_protobuf_proto.Message {
	return NewTestFromFace(this)
}

func (this *Test) GetMsg() string {
	return this.Msg
}

func (this *Test) GetNum() float64 {
	return this.Num
}

func NewTestFromFace(that TestFace) *Test {
	this := &Test{}
	this.Msg = that.GetMsg()
	this.Num = that.GetNum()
	return this
}