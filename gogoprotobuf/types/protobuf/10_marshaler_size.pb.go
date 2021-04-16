// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: 10_marshaler_size.proto

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

type Test10 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	T   *Test10_1 `protobuf:"bytes,1,opt,name=t,proto3" json:"t,omitempty"`
	Str string    `protobuf:"bytes,2,opt,name=str,proto3" json:"str,omitempty"`
	Num float64   `protobuf:"fixed64,3,opt,name=num,proto3" json:"num,omitempty"`
}

func (x *Test10) Reset() {
	*x = Test10{}
	if protoimpl.UnsafeEnabled {
		mi := &file__10_marshaler_size_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Test10) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Test10) ProtoMessage() {}

func (x *Test10) ProtoReflect() protoreflect.Message {
	mi := &file__10_marshaler_size_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Test10.ProtoReflect.Descriptor instead.
func (*Test10) Descriptor() ([]byte, []int) {
	return file__10_marshaler_size_proto_rawDescGZIP(), []int{0}
}

func (x *Test10) GetT() *Test10_1 {
	if x != nil {
		return x.T
	}
	return nil
}

func (x *Test10) GetStr() string {
	if x != nil {
		return x.Str
	}
	return ""
}

func (x *Test10) GetNum() float64 {
	if x != nil {
		return x.Num
	}
	return 0
}

type Test10_1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *Test10_1) Reset() {
	*x = Test10_1{}
	if protoimpl.UnsafeEnabled {
		mi := &file__10_marshaler_size_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Test10_1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Test10_1) ProtoMessage() {}

func (x *Test10_1) ProtoReflect() protoreflect.Message {
	mi := &file__10_marshaler_size_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Test10_1.ProtoReflect.Descriptor instead.
func (*Test10_1) Descriptor() ([]byte, []int) {
	return file__10_marshaler_size_proto_rawDescGZIP(), []int{1}
}

func (x *Test10_1) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File__10_marshaler_size_proto protoreflect.FileDescriptor

var file__10_marshaler_size_proto_rawDesc = []byte{
	0x0a, 0x17, 0x31, 0x30, 0x5f, 0x6d, 0x61, 0x72, 0x73, 0x68, 0x61, 0x6c, 0x65, 0x72, 0x5f, 0x73,
	0x69, 0x7a, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x22, 0x4b, 0x0a, 0x06, 0x54, 0x65, 0x73, 0x74, 0x31, 0x30, 0x12, 0x1d, 0x0a, 0x01, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x54, 0x65,
	0x73, 0x74, 0x31, 0x30, 0x5f, 0x31, 0x52, 0x01, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x74, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x74, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x6e,
	0x75, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x6e, 0x75, 0x6d, 0x22, 0x1c, 0x0a,
	0x08, 0x54, 0x65, 0x73, 0x74, 0x31, 0x30, 0x5f, 0x31, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x42, 0x0d, 0x5a, 0x0b, 0x2e,
	0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file__10_marshaler_size_proto_rawDescOnce sync.Once
	file__10_marshaler_size_proto_rawDescData = file__10_marshaler_size_proto_rawDesc
)

func file__10_marshaler_size_proto_rawDescGZIP() []byte {
	file__10_marshaler_size_proto_rawDescOnce.Do(func() {
		file__10_marshaler_size_proto_rawDescData = protoimpl.X.CompressGZIP(file__10_marshaler_size_proto_rawDescData)
	})
	return file__10_marshaler_size_proto_rawDescData
}

var file__10_marshaler_size_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file__10_marshaler_size_proto_goTypes = []interface{}{
	(*Test10)(nil),   // 0: types.Test10
	(*Test10_1)(nil), // 1: types.Test10_1
}
var file__10_marshaler_size_proto_depIdxs = []int32{
	1, // 0: types.Test10.t:type_name -> types.Test10_1
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file__10_marshaler_size_proto_init() }
func file__10_marshaler_size_proto_init() {
	if File__10_marshaler_size_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file__10_marshaler_size_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Test10); i {
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
		file__10_marshaler_size_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Test10_1); i {
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
			RawDescriptor: file__10_marshaler_size_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file__10_marshaler_size_proto_goTypes,
		DependencyIndexes: file__10_marshaler_size_proto_depIdxs,
		MessageInfos:      file__10_marshaler_size_proto_msgTypes,
	}.Build()
	File__10_marshaler_size_proto = out.File
	file__10_marshaler_size_proto_rawDesc = nil
	file__10_marshaler_size_proto_goTypes = nil
	file__10_marshaler_size_proto_depIdxs = nil
}
