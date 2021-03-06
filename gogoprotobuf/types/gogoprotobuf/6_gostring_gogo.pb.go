// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: 6_gostring_gogo.proto

package gogoprotobuf

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	math "math"
	reflect "reflect"
	strings "strings"
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

type Test6 struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Test6) Reset()         { *m = Test6{} }
func (m *Test6) String() string { return proto.CompactTextString(m) }
func (*Test6) ProtoMessage()    {}
func (*Test6) Descriptor() ([]byte, []int) {
	return fileDescriptor_d109d982c8f34d6d, []int{0}
}
func (m *Test6) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Test6.Unmarshal(m, b)
}
func (m *Test6) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Test6.Marshal(b, m, deterministic)
}
func (m *Test6) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Test6.Merge(m, src)
}
func (m *Test6) XXX_Size() int {
	return xxx_messageInfo_Test6.Size(m)
}
func (m *Test6) XXX_DiscardUnknown() {
	xxx_messageInfo_Test6.DiscardUnknown(m)
}

var xxx_messageInfo_Test6 proto.InternalMessageInfo

func (m *Test6) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*Test6)(nil), "types.Test6")
}

func init() { proto.RegisterFile("6_gostring_gogo.proto", fileDescriptor_d109d982c8f34d6d) }

var fileDescriptor_d109d982c8f34d6d = []byte{
	// 128 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x35, 0x8b, 0x4f, 0xcf,
	0x2f, 0x2e, 0x29, 0xca, 0xcc, 0x4b, 0x8f, 0x4f, 0xcf, 0x4f, 0xcf, 0xd7, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0x62, 0x2d, 0xa9, 0x2c, 0x48, 0x2d, 0x96, 0xd2, 0x4d, 0xcf, 0x2c, 0xc9, 0x28, 0x4d,
	0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x07, 0xc9, 0xea, 0x83, 0x65, 0x93, 0x4a, 0xd3, 0xc0, 0x3c, 0x30,
	0x47, 0x1f, 0xa1, 0x4b, 0x49, 0x9a, 0x8b, 0x35, 0x24, 0xb5, 0xb8, 0xc4, 0x4c, 0x48, 0x88, 0x8b,
	0x25, 0x2f, 0x31, 0x37, 0x55, 0x82, 0x51, 0x81, 0x51, 0x83, 0x33, 0x08, 0xcc, 0x76, 0x12, 0xfd,
	0xf0, 0x50, 0x8e, 0x31, 0x8a, 0x5f, 0x4f, 0xdf, 0x1a, 0xae, 0x37, 0xa9, 0x34, 0x2d, 0x89, 0x0d,
	0xcc, 0x32, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xff, 0x55, 0xe6, 0x9a, 0x89, 0x00, 0x00, 0x00,
}

func (this *Test6) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&gogoprotobuf.Test6{")
	s = append(s, "Name: "+fmt.Sprintf("%#v", this.Name)+",\n")
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoString6GostringGogo(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
