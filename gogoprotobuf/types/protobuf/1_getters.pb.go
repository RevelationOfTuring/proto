// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: 1_getters.proto

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

type Struct struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Salary int32 `protobuf:"varint,1,opt,name=salary,proto3" json:"salary,omitempty"`
}

func (x *Struct) Reset() {
	*x = Struct{}
	if protoimpl.UnsafeEnabled {
		mi := &file__1_getters_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Struct) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Struct) ProtoMessage() {}

func (x *Struct) ProtoReflect() protoreflect.Message {
	mi := &file__1_getters_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Struct.ProtoReflect.Descriptor instead.
func (*Struct) Descriptor() ([]byte, []int) {
	return file__1_getters_proto_rawDescGZIP(), []int{0}
}

func (x *Struct) GetSalary() int32 {
	if x != nil {
		return x.Salary
	}
	return 0
}

var File__1_getters_proto protoreflect.FileDescriptor

var file__1_getters_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x31, 0x5f, 0x67, 0x65, 0x74, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x74, 0x79, 0x70, 0x65, 0x73, 0x22, 0x20, 0x0a, 0x06, 0x53, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x61, 0x6c, 0x61, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x73, 0x61, 0x6c, 0x61, 0x72, 0x79, 0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x2f,
	0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file__1_getters_proto_rawDescOnce sync.Once
	file__1_getters_proto_rawDescData = file__1_getters_proto_rawDesc
)

func file__1_getters_proto_rawDescGZIP() []byte {
	file__1_getters_proto_rawDescOnce.Do(func() {
		file__1_getters_proto_rawDescData = protoimpl.X.CompressGZIP(file__1_getters_proto_rawDescData)
	})
	return file__1_getters_proto_rawDescData
}

var file__1_getters_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file__1_getters_proto_goTypes = []interface{}{
	(*Struct)(nil), // 0: types.Struct
}
var file__1_getters_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file__1_getters_proto_init() }
func file__1_getters_proto_init() {
	if File__1_getters_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file__1_getters_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Struct); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file__1_getters_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file__1_getters_proto_goTypes,
		DependencyIndexes: file__1_getters_proto_depIdxs,
		MessageInfos:      file__1_getters_proto_msgTypes,
	}.Build()
	File__1_getters_proto = out.File
	file__1_getters_proto_rawDesc = nil
	file__1_getters_proto_goTypes = nil
	file__1_getters_proto_depIdxs = nil
}