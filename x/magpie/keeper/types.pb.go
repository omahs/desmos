// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: desmos/magpie/v1beta1/keeper/types.proto

package keeper

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
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

// WrappedSessionLength wraps the session length
type WrappedSessionLength struct {
	Length uint64 `protobuf:"varint,1,opt,name=length,proto3" json:"length,omitempty"`
}

func (m *WrappedSessionLength) Reset()         { *m = WrappedSessionLength{} }
func (m *WrappedSessionLength) String() string { return proto.CompactTextString(m) }
func (*WrappedSessionLength) ProtoMessage()    {}
func (*WrappedSessionLength) Descriptor() ([]byte, []int) {
	return fileDescriptor_785a50363639ec77, []int{0}
}
func (m *WrappedSessionLength) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *WrappedSessionLength) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_WrappedSessionLength.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *WrappedSessionLength) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WrappedSessionLength.Merge(m, src)
}
func (m *WrappedSessionLength) XXX_Size() int {
	return m.Size()
}
func (m *WrappedSessionLength) XXX_DiscardUnknown() {
	xxx_messageInfo_WrappedSessionLength.DiscardUnknown(m)
}

var xxx_messageInfo_WrappedSessionLength proto.InternalMessageInfo

func (m *WrappedSessionLength) GetLength() uint64 {
	if m != nil {
		return m.Length
	}
	return 0
}

func init() {
	proto.RegisterType((*WrappedSessionLength)(nil), "desmos.magpie.v1beta1.keeper.WrappedSessionLength")
}

func init() {
	proto.RegisterFile("desmos/magpie/v1beta1/keeper/types.proto", fileDescriptor_785a50363639ec77)
}

var fileDescriptor_785a50363639ec77 = []byte{
	// 203 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x48, 0x49, 0x2d, 0xce,
	0xcd, 0x2f, 0xd6, 0xcf, 0x4d, 0x4c, 0x2f, 0xc8, 0x4c, 0xd5, 0x2f, 0x33, 0x4c, 0x4a, 0x2d, 0x49,
	0x34, 0xd4, 0xcf, 0x4e, 0x4d, 0x2d, 0x48, 0x2d, 0xd2, 0x2f, 0xa9, 0x2c, 0x48, 0x2d, 0xd6, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x81, 0xa8, 0xd4, 0x83, 0xa8, 0xd4, 0x83, 0xaa, 0xd4, 0x83,
	0xa8, 0x94, 0x12, 0x49, 0xcf, 0x4f, 0xcf, 0x07, 0x2b, 0xd4, 0x07, 0xb1, 0x20, 0x7a, 0x94, 0x2c,
	0xb8, 0x44, 0xc2, 0x8b, 0x12, 0x0b, 0x0a, 0x52, 0x53, 0x82, 0x53, 0x8b, 0x8b, 0x33, 0xf3, 0xf3,
	0x7c, 0x52, 0xf3, 0xd2, 0x4b, 0x32, 0x84, 0xc4, 0xb8, 0xd8, 0x72, 0xc0, 0x2c, 0x09, 0x46, 0x05,
	0x46, 0x0d, 0x96, 0x20, 0x28, 0xcf, 0x8a, 0x63, 0xc6, 0x02, 0x79, 0xc6, 0x17, 0x0b, 0xe4, 0x19,
	0x9d, 0xdc, 0x4f, 0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x09,
	0x8f, 0xe5, 0x18, 0x2e, 0x3c, 0x96, 0x63, 0xb8, 0xf1, 0x58, 0x8e, 0x21, 0x4a, 0x37, 0x3d, 0xb3,
	0x24, 0xa3, 0x34, 0x49, 0x2f, 0x39, 0x3f, 0x57, 0x1f, 0xe2, 0x24, 0xdd, 0x9c, 0xc4, 0xa4, 0x62,
	0x28, 0x5b, 0xbf, 0x02, 0xe6, 0x15, 0x88, 0xc3, 0x92, 0xd8, 0xc0, 0x2e, 0x31, 0x06, 0x04, 0x00,
	0x00, 0xff, 0xff, 0xc4, 0xcc, 0x4e, 0x24, 0xe9, 0x00, 0x00, 0x00,
}

func (this *WrappedSessionLength) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*WrappedSessionLength)
	if !ok {
		that2, ok := that.(WrappedSessionLength)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Length != that1.Length {
		return false
	}
	return true
}
func (m *WrappedSessionLength) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *WrappedSessionLength) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *WrappedSessionLength) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Length != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Length))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintTypes(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypes(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *WrappedSessionLength) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Length != 0 {
		n += 1 + sovTypes(uint64(m.Length))
	}
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *WrappedSessionLength) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: WrappedSessionLength: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WrappedSessionLength: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Length", wireType)
			}
			m.Length = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Length |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
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
func skipTypes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
				return 0, ErrInvalidLengthTypes
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypes
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypes
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypes        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypes          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypes = fmt.Errorf("proto: unexpected end of group")
)