// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: 3_nullable.proto

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

type Test3 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	T   *Test3_1 `protobuf:"bytes,1,opt,name=t,proto3" json:"t,omitempty"`
	Num float64  `protobuf:"fixed64,2,opt,name=num,proto3" json:"num,omitempty"`
}

func (x *Test3) Reset() {
	*x = Test3{}
	if protoimpl.UnsafeEnabled {
		mi := &file__3_nullable_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Test3) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Test3) ProtoMessage() {}

func (x *Test3) ProtoReflect() protoreflect.Message {
	mi := &file__3_nullable_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Test3.ProtoReflect.Descriptor instead.
func (*Test3) Descriptor() ([]byte, []int) {
	return file__3_nullable_proto_rawDescGZIP(), []int{0}
}

func (x *Test3) GetT() *Test3_1 {
	if x != nil {
		return x.T
	}
	return nil
}

func (x *Test3) GetNum() float64 {
	if x != nil {
		return x.Num
	}
	return 0
}

type Test3_1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *Test3_1) Reset() {
	*x = Test3_1{}
	if protoimpl.UnsafeEnabled {
		mi := &file__3_nullable_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Test3_1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Test3_1) ProtoMessage() {}

func (x *Test3_1) ProtoReflect() protoreflect.Message {
	mi := &file__3_nullable_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Test3_1.ProtoReflect.Descriptor instead.
func (*Test3_1) Descriptor() ([]byte, []int) {
	return file__3_nullable_proto_rawDescGZIP(), []int{1}
}

func (x *Test3_1) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File__3_nullable_proto protoreflect.FileDescriptor

var file__3_nullable_proto_rawDesc = []byte{
	0x0a, 0x10, 0x33, 0x5f, 0x6e, 0x75, 0x6c, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x05, 0x74, 0x79, 0x70, 0x65, 0x73, 0x22, 0x37, 0x0a, 0x05, 0x54, 0x65, 0x73,
	0x74, 0x33, 0x12, 0x1c, 0x0a, 0x01, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x33, 0x5f, 0x31, 0x52, 0x01, 0x74,
	0x12, 0x10, 0x0a, 0x03, 0x6e, 0x75, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x6e,
	0x75, 0x6d, 0x22, 0x1b, 0x0a, 0x07, 0x54, 0x65, 0x73, 0x74, 0x33, 0x5f, 0x31, 0x12, 0x10, 0x0a,
	0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x42,
	0x0d, 0x5a, 0x0b, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file__3_nullable_proto_rawDescOnce sync.Once
	file__3_nullable_proto_rawDescData = file__3_nullable_proto_rawDesc
)

func file__3_nullable_proto_rawDescGZIP() []byte {
	file__3_nullable_proto_rawDescOnce.Do(func() {
		file__3_nullable_proto_rawDescData = protoimpl.X.CompressGZIP(file__3_nullable_proto_rawDescData)
	})
	return file__3_nullable_proto_rawDescData
}

var file__3_nullable_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file__3_nullable_proto_goTypes = []interface{}{
	(*Test3)(nil),   // 0: types.Test3
	(*Test3_1)(nil), // 1: types.Test3_1
}
var file__3_nullable_proto_depIdxs = []int32{
	1, // 0: types.Test3.t:type_name -> types.Test3_1
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file__3_nullable_proto_init() }
func file__3_nullable_proto_init() {
	if File__3_nullable_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file__3_nullable_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Test3); i {
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
		file__3_nullable_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Test3_1); i {
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
			RawDescriptor: file__3_nullable_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file__3_nullable_proto_goTypes,
		DependencyIndexes: file__3_nullable_proto_depIdxs,
		MessageInfos:      file__3_nullable_proto_msgTypes,
	}.Build()
	File__3_nullable_proto = out.File
	file__3_nullable_proto_rawDesc = nil
	file__3_nullable_proto_goTypes = nil
	file__3_nullable_proto_depIdxs = nil
}
