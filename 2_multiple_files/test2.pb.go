// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: test2.proto

package test2

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
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

type M2 struct {
	Other *Other `protobuf:"bytes,1,opt,name=other,proto3" json:"other,omitempty"`
}

func (m *M2) Reset()         { *m = M2{} }
func (m *M2) String() string { return proto.CompactTextString(m) }
func (*M2) ProtoMessage()    {}
func (*M2) Descriptor() ([]byte, []int) {
	return fileDescriptor_61e5535eff96cd18, []int{0}
}
func (m *M2) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *M2) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_M2.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *M2) XXX_Merge(src proto.Message) {
	xxx_messageInfo_M2.Merge(m, src)
}
func (m *M2) XXX_Size() int {
	return m.Size()
}
func (m *M2) XXX_DiscardUnknown() {
	xxx_messageInfo_M2.DiscardUnknown(m)
}

var xxx_messageInfo_M2 proto.InternalMessageInfo

func (m *M2) GetOther() *Other {
	if m != nil {
		return m.Other
	}
	return nil
}

func init() {
	proto.RegisterType((*M2)(nil), "test2.M2")
}

func init() { proto.RegisterFile("test2.proto", fileDescriptor_61e5535eff96cd18) }

var fileDescriptor_61e5535eff96cd18 = []byte{
	// 111 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2e, 0x49, 0x2d, 0x2e,
	0x31, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x73, 0xa4, 0xb8, 0xf3, 0x4b, 0x32,
	0x52, 0x8b, 0x20, 0x62, 0x4a, 0x1a, 0x5c, 0x4c, 0xbe, 0x46, 0x42, 0x4a, 0x5c, 0xac, 0x60, 0x41,
	0x09, 0x46, 0x05, 0x46, 0x0d, 0x6e, 0x23, 0x1e, 0x3d, 0x88, 0x36, 0x7f, 0x90, 0x58, 0x10, 0x44,
	0xca, 0x49, 0xe2, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c,
	0xf0, 0x58, 0x8e, 0xe1, 0xc2, 0x63, 0x39, 0x86, 0x1b, 0x8f, 0xe5, 0x18, 0x92, 0xd8, 0xc0, 0x46,
	0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x1d, 0xfe, 0x12, 0x0a, 0x6d, 0x00, 0x00, 0x00,
}

func (m *M2) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *M2) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *M2) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Other != nil {
		{
			size, err := m.Other.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTest2(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTest2(dAtA []byte, offset int, v uint64) int {
	offset -= sovTest2(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *M2) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Other != nil {
		l = m.Other.Size()
		n += 1 + l + sovTest2(uint64(l))
	}
	return n
}

func sovTest2(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTest2(x uint64) (n int) {
	return sovTest2(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *M2) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTest2
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
			return fmt.Errorf("proto: M2: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: M2: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Other", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTest2
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
				return ErrInvalidLengthTest2
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTest2
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Other == nil {
				m.Other = &Other{}
			}
			if err := m.Other.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTest2(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTest2
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTest2
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTest2(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTest2
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
					return 0, ErrIntOverflowTest2
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
					return 0, ErrIntOverflowTest2
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
				return 0, ErrInvalidLengthTest2
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTest2
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTest2
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTest2        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTest2          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTest2 = fmt.Errorf("proto: unexpected end of group")
)
