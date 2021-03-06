// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: 2_face.proto

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

type Test2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string  `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	Num float64 `protobuf:"fixed64,2,opt,name=num,proto3" json:"num,omitempty"`
}

func (x *Test2) Reset() {
	*x = Test2{}
	if protoimpl.UnsafeEnabled {
		mi := &file__2_face_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Test2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Test2) ProtoMessage() {}

func (x *Test2) ProtoReflect() protoreflect.Message {
	mi := &file__2_face_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Test2.ProtoReflect.Descriptor instead.
func (*Test2) Descriptor() ([]byte, []int) {
	return file__2_face_proto_rawDescGZIP(), []int{0}
}

func (x *Test2) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *Test2) GetNum() float64 {
	if x != nil {
		return x.Num
	}
	return 0
}

var File__2_face_proto protoreflect.FileDescriptor

var file__2_face_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x32, 0x5f, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x22, 0x2b, 0x0a, 0x05, 0x54, 0x65, 0x73, 0x74, 0x32, 0x12, 0x10,
	0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67,
	0x12, 0x10, 0x0a, 0x03, 0x6e, 0x75, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x6e,
	0x75, 0x6d, 0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file__2_face_proto_rawDescOnce sync.Once
	file__2_face_proto_rawDescData = file__2_face_proto_rawDesc
)

func file__2_face_proto_rawDescGZIP() []byte {
	file__2_face_proto_rawDescOnce.Do(func() {
		file__2_face_proto_rawDescData = protoimpl.X.CompressGZIP(file__2_face_proto_rawDescData)
	})
	return file__2_face_proto_rawDescData
}

var file__2_face_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file__2_face_proto_goTypes = []interface{}{
	(*Test2)(nil), // 0: types.Test2
}
var file__2_face_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file__2_face_proto_init() }
func file__2_face_proto_init() {
	if File__2_face_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file__2_face_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Test2); i {
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
			RawDescriptor: file__2_face_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file__2_face_proto_goTypes,
		DependencyIndexes: file__2_face_proto_depIdxs,
		MessageInfos:      file__2_face_proto_msgTypes,
	}.Build()
	File__2_face_proto = out.File
	file__2_face_proto_rawDesc = nil
	file__2_face_proto_goTypes = nil
	file__2_face_proto_depIdxs = nil
}
