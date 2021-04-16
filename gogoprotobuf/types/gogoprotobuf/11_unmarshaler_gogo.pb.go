// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: 11_unmarshaler_gogo.proto

package gogoprotobuf

import (
	encoding_binary "encoding/binary"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
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

type Test11 struct {
	T                    *Test11_1 `protobuf:"bytes,1,opt,name=t,proto3" json:"t,omitempty"`
	Str                  string    `protobuf:"bytes,2,opt,name=str,proto3" json:"str,omitempty"`
	Num                  float64   `protobuf:"fixed64,3,opt,name=num,proto3" json:"num,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Test11) Reset()         { *m = Test11{} }
func (m *Test11) String() string { return proto.CompactTextString(m) }
func (*Test11) ProtoMessage()    {}
func (*Test11) Descriptor() ([]byte, []int) {
	return fileDescriptor_480601f53b5cdcf4, []int{0}
}
func (m *Test11) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Test11) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Test11.Marshal(b, m, deterministic)
}
func (m *Test11) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Test11.Merge(m, src)
}
func (m *Test11) XXX_Size() int {
	return xxx_messageInfo_Test11.Size(m)
}
func (m *Test11) XXX_DiscardUnknown() {
	xxx_messageInfo_Test11.DiscardUnknown(m)
}

var xxx_messageInfo_Test11 proto.InternalMessageInfo

func (m *Test11) GetT() *Test11_1 {
	if m != nil {
		return m.T
	}
	return nil
}

func (m *Test11) GetStr() string {
	if m != nil {
		return m.Str
	}
	return ""
}

func (m *Test11) GetNum() float64 {
	if m != nil {
		return m.Num
	}
	return 0
}

type Test11_1 struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Test11_1) Reset()         { *m = Test11_1{} }
func (m *Test11_1) String() string { return proto.CompactTextString(m) }
func (*Test11_1) ProtoMessage()    {}
func (*Test11_1) Descriptor() ([]byte, []int) {
	return fileDescriptor_480601f53b5cdcf4, []int{1}
}
func (m *Test11_1) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Test11_1) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Test11_1.Marshal(b, m, deterministic)
}
func (m *Test11_1) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Test11_1.Merge(m, src)
}
func (m *Test11_1) XXX_Size() int {
	return xxx_messageInfo_Test11_1.Size(m)
}
func (m *Test11_1) XXX_DiscardUnknown() {
	xxx_messageInfo_Test11_1.DiscardUnknown(m)
}

var xxx_messageInfo_Test11_1 proto.InternalMessageInfo

func (m *Test11_1) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*Test11)(nil), "types.Test11")
	proto.RegisterType((*Test11_1)(nil), "types.Test11_1")
}

func init() { proto.RegisterFile("11_unmarshaler_gogo.proto", fileDescriptor_480601f53b5cdcf4) }

var fileDescriptor_480601f53b5cdcf4 = []byte{
	// 187 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x34, 0x34, 0x8c, 0x2f,
	0xcd, 0xcb, 0x4d, 0x2c, 0x2a, 0xce, 0x48, 0xcc, 0x49, 0x2d, 0x8a, 0x4f, 0xcf, 0x4f, 0xcf, 0xd7,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2d, 0xa9, 0x2c, 0x48, 0x2d, 0x96, 0xd2, 0x4d, 0xcf,
	0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x07, 0xc9, 0xea, 0x83, 0x65, 0x93, 0x4a,
	0xd3, 0xc0, 0x3c, 0x30, 0x47, 0x1f, 0xa1, 0x4b, 0xc9, 0x9b, 0x8b, 0x2d, 0x24, 0xb5, 0xb8, 0xc4,
	0xd0, 0x50, 0x48, 0x96, 0x8b, 0xb1, 0x44, 0x82, 0x51, 0x81, 0x51, 0x83, 0xdb, 0x88, 0x5f, 0x0f,
	0x6c, 0x96, 0x1e, 0x44, 0x26, 0xde, 0x30, 0x88, 0xb1, 0x44, 0x48, 0x80, 0x8b, 0xb9, 0xb8, 0xa4,
	0x48, 0x82, 0x49, 0x81, 0x51, 0x83, 0x33, 0x08, 0xc4, 0x04, 0x89, 0xe4, 0x95, 0xe6, 0x4a, 0x30,
	0x2b, 0x30, 0x6a, 0x30, 0x06, 0x81, 0x98, 0x4a, 0x32, 0x5c, 0x1c, 0x30, 0x2d, 0x20, 0xd9, 0xdc,
	0xe2, 0x74, 0xb0, 0x81, 0x9c, 0x41, 0x20, 0xa6, 0x93, 0xe8, 0x85, 0x47, 0x72, 0x8c, 0x51, 0xfc,
	0x7a, 0xfa, 0xd6, 0x70, 0x97, 0x24, 0x95, 0xa6, 0x25, 0xb1, 0x81, 0x59, 0xc6, 0x80, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xbd, 0xd8, 0x93, 0x00, 0xdb, 0x00, 0x00, 0x00,
}

func (m *Test11) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflow11UnmarshalerGogo
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Test11: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Test11: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field T", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow11UnmarshalerGogo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLength11UnmarshalerGogo
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLength11UnmarshalerGogo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.T == nil {
				m.T = &Test11_1{}
			}
			if err := m.T.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Str", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow11UnmarshalerGogo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLength11UnmarshalerGogo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLength11UnmarshalerGogo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Str = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Num", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.Num = float64(math.Float64frombits(v))
		default:
			iNdEx = preIndex
			skippy, err := skip11UnmarshalerGogo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLength11UnmarshalerGogo
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Test11_1) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflow11UnmarshalerGogo
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Test11_1: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Test11_1: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Msg", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow11UnmarshalerGogo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLength11UnmarshalerGogo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLength11UnmarshalerGogo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Msg = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skip11UnmarshalerGogo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLength11UnmarshalerGogo
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skip11UnmarshalerGogo(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflow11UnmarshalerGogo
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflow11UnmarshalerGogo
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflow11UnmarshalerGogo
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLength11UnmarshalerGogo
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroup11UnmarshalerGogo
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLength11UnmarshalerGogo
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLength11UnmarshalerGogo        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflow11UnmarshalerGogo          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroup11UnmarshalerGogo = fmt.Errorf("proto: unexpected end of group")
)