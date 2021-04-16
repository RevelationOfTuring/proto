// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: 0_enum_prefix.proto

package protobuf

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type MyEnum int32

const (
	MyEnum_Black  MyEnum = 0
	MyEnum_White  MyEnum = 1
	MyEnum_Red    MyEnum = 2
	MyEnum_Yellow MyEnum = 3
)

// Enum value maps for MyEnum.
var (
	MyEnum_name = map[int32]string{
		0: "Black",
		1: "White",
		2: "Red",
		3: "Yellow",
	}
	MyEnum_value = map[string]int32{
		"Black":  0,
		"White":  1,
		"Red":    2,
		"Yellow": 3,
	}
)

func (x MyEnum) Enum() *MyEnum {
	p := new(MyEnum)
	*p = x
	return p
}

func (x MyEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MyEnum) Descriptor() protoreflect.EnumDescriptor {
	return file__0_enum_prefix_proto_enumTypes[0].Descriptor()
}

func (MyEnum) Type() protoreflect.EnumType {
	return &file__0_enum_prefix_proto_enumTypes[0]
}

func (x MyEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MyEnum.Descriptor instead.
func (MyEnum) EnumDescriptor() ([]byte, []int) {
	return file__0_enum_prefix_proto_rawDescGZIP(), []int{0}
}

var File__0_enum_prefix_proto protoreflect.FileDescriptor

var file__0_enum_prefix_proto_rawDesc = []byte{
	0x0a, 0x13, 0x30, 0x5f, 0x65, 0x6e, 0x75, 0x6d, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2a, 0x33, 0x0a, 0x06,
	0x4d, 0x79, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x09, 0x0a, 0x05, 0x42, 0x6c, 0x61, 0x63, 0x6b, 0x10,
	0x00, 0x12, 0x09, 0x0a, 0x05, 0x57, 0x68, 0x69, 0x74, 0x65, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03,
	0x52, 0x65, 0x64, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x59, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x10,
	0x03, 0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file__0_enum_prefix_proto_rawDescOnce sync.Once
	file__0_enum_prefix_proto_rawDescData = file__0_enum_prefix_proto_rawDesc
)

func file__0_enum_prefix_proto_rawDescGZIP() []byte {
	file__0_enum_prefix_proto_rawDescOnce.Do(func() {
		file__0_enum_prefix_proto_rawDescData = protoimpl.X.CompressGZIP(file__0_enum_prefix_proto_rawDescData)
	})
	return file__0_enum_prefix_proto_rawDescData
}

var file__0_enum_prefix_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file__0_enum_prefix_proto_goTypes = []interface{}{
	(MyEnum)(0), // 0: types.MyEnum
}
var file__0_enum_prefix_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file__0_enum_prefix_proto_init() }
func file__0_enum_prefix_proto_init() {
	if File__0_enum_prefix_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file__0_enum_prefix_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file__0_enum_prefix_proto_goTypes,
		DependencyIndexes: file__0_enum_prefix_proto_depIdxs,
		EnumInfos:         file__0_enum_prefix_proto_enumTypes,
	}.Build()
	File__0_enum_prefix_proto = out.File
	file__0_enum_prefix_proto_rawDesc = nil
	file__0_enum_prefix_proto_goTypes = nil
	file__0_enum_prefix_proto_depIdxs = nil
}
