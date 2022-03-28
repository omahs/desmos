// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: desmos/profiles/v2/models_dtag_requests.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/codec/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/regen-network/cosmos-proto"
	_ "google.golang.org/protobuf/types/known/timestamppb"
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

// DTagTransferRequest represent a DTag transfer request between two users
type DTagTransferRequest struct {
	// DTagToTrade contains the value of the DTag that should be transferred from
	// the receiver of the request to the sender
	DTagToTrade string `protobuf:"bytes,1,opt,name=dtag_to_trade,json=dtagToTrade,proto3" json:"dtag_to_trade,omitempty" yaml:"dtag_to_trade"`
	// Sender represents the address of the account that sent the request
	Sender string `protobuf:"bytes,2,opt,name=sender,proto3" json:"sender,omitempty" yaml:"sender"`
	// Receiver represents the receiver of the request that, if accepted, will
	// give to the sender their DTag
	Receiver string `protobuf:"bytes,3,opt,name=receiver,proto3" json:"receiver,omitempty" yaml:"receiver"`
}

func (m *DTagTransferRequest) Reset()         { *m = DTagTransferRequest{} }
func (m *DTagTransferRequest) String() string { return proto.CompactTextString(m) }
func (*DTagTransferRequest) ProtoMessage()    {}
func (*DTagTransferRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3722cfac854d7654, []int{0}
}
func (m *DTagTransferRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DTagTransferRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DTagTransferRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DTagTransferRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DTagTransferRequest.Merge(m, src)
}
func (m *DTagTransferRequest) XXX_Size() int {
	return m.Size()
}
func (m *DTagTransferRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DTagTransferRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DTagTransferRequest proto.InternalMessageInfo

func init() {
	proto.RegisterType((*DTagTransferRequest)(nil), "desmos.profiles.v2.DTagTransferRequest")
}

func init() {
	proto.RegisterFile("desmos/profiles/v2/models_dtag_requests.proto", fileDescriptor_3722cfac854d7654)
}

var fileDescriptor_3722cfac854d7654 = []byte{
	// 355 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0xb1, 0x4e, 0xe3, 0x40,
	0x10, 0x86, 0xbd, 0x77, 0x52, 0x74, 0xe7, 0x5c, 0x74, 0xc2, 0x49, 0x11, 0x52, 0x78, 0x91, 0x1b,
	0x40, 0x28, 0x5e, 0x91, 0x74, 0x29, 0x23, 0xba, 0x74, 0x56, 0x2a, 0x1a, 0x6b, 0x6d, 0x4f, 0x1c,
	0x0b, 0xdb, 0x6b, 0x76, 0x37, 0x16, 0x79, 0x03, 0x4a, 0x4a, 0xca, 0x3c, 0x0e, 0x05, 0x45, 0x4a,
	0x2a, 0x0b, 0x39, 0x0d, 0x75, 0x9e, 0x00, 0xd9, 0xeb, 0x80, 0x40, 0xa2, 0x9b, 0x7f, 0xfe, 0x6f,
	0x66, 0x77, 0x7e, 0x7d, 0x18, 0x80, 0x48, 0x98, 0x20, 0x19, 0x67, 0x8b, 0x28, 0x06, 0x41, 0xf2,
	0x11, 0x49, 0x58, 0x00, 0xb1, 0x70, 0x03, 0x49, 0x43, 0x97, 0xc3, 0xed, 0x0a, 0x84, 0x14, 0x76,
	0xc6, 0x99, 0x64, 0x86, 0xa1, 0x70, 0xfb, 0x80, 0xdb, 0xf9, 0x68, 0xd0, 0x0b, 0x59, 0xc8, 0x6a,
	0x9b, 0x54, 0x95, 0x22, 0x07, 0xc7, 0x21, 0x63, 0x61, 0x0c, 0xa4, 0x56, 0xde, 0x6a, 0x41, 0x68,
	0xba, 0x6e, 0x2c, 0xfc, 0xdd, 0x92, 0x51, 0x02, 0x42, 0xd2, 0x24, 0x3b, 0xcc, 0xfa, 0xac, 0x7a,
	0xc5, 0x55, 0x4b, 0x95, 0x68, 0xac, 0x8b, 0x9f, 0xff, 0xeb, 0x2f, 0x69, 0x94, 0xba, 0x71, 0x94,
	0xde, 0x34, 0xb0, 0xf5, 0x8c, 0xf4, 0xee, 0xd5, 0x9c, 0x86, 0x73, 0x4e, 0x53, 0xb1, 0x00, 0xee,
	0xa8, 0x63, 0x8c, 0x99, 0xde, 0xa9, 0x8f, 0x93, 0xcc, 0x95, 0x9c, 0x06, 0xd0, 0x47, 0x27, 0xe8,
	0xec, 0xef, 0xf4, 0xb4, 0x2c, 0x70, 0xbb, 0xe6, 0xd9, 0xbc, 0x6a, 0xef, 0x0b, 0xdc, 0x5b, 0xd3,
	0x24, 0x9e, 0x58, 0x5f, 0x68, 0xcb, 0x69, 0x57, 0xba, 0x81, 0x8c, 0x73, 0xbd, 0x25, 0x20, 0x0d,
	0x80, 0xf7, 0x7f, 0xd5, 0x5b, 0x8e, 0xf6, 0x05, 0xee, 0xa8, 0x31, 0xd5, 0xb7, 0x9c, 0x06, 0x30,
	0x88, 0xfe, 0x87, 0x83, 0x0f, 0x51, 0x0e, 0xbc, 0xff, 0xbb, 0x86, 0xbb, 0xfb, 0x02, 0xff, 0x57,
	0xf0, 0xc1, 0xb1, 0x9c, 0x0f, 0x68, 0xf2, 0xef, 0x7e, 0x83, 0xb5, 0xc7, 0x0d, 0x46, 0x6f, 0x1b,
	0x8c, 0xa6, 0xb3, 0xa7, 0xd2, 0x44, 0xdb, 0xd2, 0x44, 0xaf, 0xa5, 0x89, 0x1e, 0x76, 0xa6, 0xb6,
	0xdd, 0x99, 0xda, 0xcb, 0xce, 0xd4, 0xae, 0x2f, 0xc3, 0x48, 0x2e, 0x57, 0x9e, 0xed, 0xb3, 0x84,
	0xa8, 0x80, 0x86, 0x31, 0xf5, 0x44, 0x53, 0x93, 0x7c, 0x4c, 0xee, 0x3e, 0x13, 0x93, 0xeb, 0x0c,
	0x84, 0xd7, 0xaa, 0x23, 0x1a, 0xbf, 0x07, 0x00, 0x00, 0xff, 0xff, 0x03, 0xd4, 0x84, 0x98, 0x01,
	0x02, 0x00, 0x00,
}

func (this *DTagTransferRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*DTagTransferRequest)
	if !ok {
		that2, ok := that.(DTagTransferRequest)
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
	if this.DTagToTrade != that1.DTagToTrade {
		return false
	}
	if this.Sender != that1.Sender {
		return false
	}
	if this.Receiver != that1.Receiver {
		return false
	}
	return true
}
func (m *DTagTransferRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DTagTransferRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DTagTransferRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Receiver) > 0 {
		i -= len(m.Receiver)
		copy(dAtA[i:], m.Receiver)
		i = encodeVarintModelsDtagRequests(dAtA, i, uint64(len(m.Receiver)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintModelsDtagRequests(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.DTagToTrade) > 0 {
		i -= len(m.DTagToTrade)
		copy(dAtA[i:], m.DTagToTrade)
		i = encodeVarintModelsDtagRequests(dAtA, i, uint64(len(m.DTagToTrade)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintModelsDtagRequests(dAtA []byte, offset int, v uint64) int {
	offset -= sovModelsDtagRequests(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *DTagTransferRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.DTagToTrade)
	if l > 0 {
		n += 1 + l + sovModelsDtagRequests(uint64(l))
	}
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovModelsDtagRequests(uint64(l))
	}
	l = len(m.Receiver)
	if l > 0 {
		n += 1 + l + sovModelsDtagRequests(uint64(l))
	}
	return n
}

func sovModelsDtagRequests(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozModelsDtagRequests(x uint64) (n int) {
	return sovModelsDtagRequests(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DTagTransferRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowModelsDtagRequests
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
			return fmt.Errorf("proto: DTagTransferRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DTagTransferRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DTagToTrade", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModelsDtagRequests
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
				return ErrInvalidLengthModelsDtagRequests
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthModelsDtagRequests
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DTagToTrade = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModelsDtagRequests
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
				return ErrInvalidLengthModelsDtagRequests
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthModelsDtagRequests
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Receiver", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModelsDtagRequests
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
				return ErrInvalidLengthModelsDtagRequests
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthModelsDtagRequests
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Receiver = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipModelsDtagRequests(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthModelsDtagRequests
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
func skipModelsDtagRequests(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowModelsDtagRequests
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
					return 0, ErrIntOverflowModelsDtagRequests
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
					return 0, ErrIntOverflowModelsDtagRequests
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
				return 0, ErrInvalidLengthModelsDtagRequests
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupModelsDtagRequests
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthModelsDtagRequests
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthModelsDtagRequests        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowModelsDtagRequests          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupModelsDtagRequests = fmt.Errorf("proto: unexpected end of group")
)