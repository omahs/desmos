// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: desmos/profiles/v3/query_params.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/codec/types"
	_ "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/regen-network/cosmos-proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// QueryParamsRequest is the request type for the Query/Params RPC endpoint
type QueryParamsRequest struct {
}

func (m *QueryParamsRequest) Reset()         { *m = QueryParamsRequest{} }
func (m *QueryParamsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryParamsRequest) ProtoMessage()    {}
func (*QueryParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9afa8e27b0a00e24, []int{0}
}
func (m *QueryParamsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsRequest.Merge(m, src)
}
func (m *QueryParamsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsRequest proto.InternalMessageInfo

// QueryParamsResponse is the response type for the Query/Params RPC method.
type QueryParamsResponse struct {
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (m *QueryParamsResponse) Reset()         { *m = QueryParamsResponse{} }
func (m *QueryParamsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryParamsResponse) ProtoMessage()    {}
func (*QueryParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9afa8e27b0a00e24, []int{1}
}
func (m *QueryParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsResponse.Merge(m, src)
}
func (m *QueryParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsResponse proto.InternalMessageInfo

func (m *QueryParamsResponse) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func init() {
	proto.RegisterType((*QueryParamsRequest)(nil), "desmos.profiles.v3.QueryParamsRequest")
	proto.RegisterType((*QueryParamsResponse)(nil), "desmos.profiles.v3.QueryParamsResponse")
}

func init() {
	proto.RegisterFile("desmos/profiles/v3/query_params.proto", fileDescriptor_9afa8e27b0a00e24)
}

var fileDescriptor_9afa8e27b0a00e24 = []byte{
	// 294 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xb1, 0x4e, 0xf3, 0x30,
	0x14, 0x85, 0x13, 0xe9, 0x57, 0x87, 0xfc, 0x5b, 0xe8, 0x40, 0x23, 0x64, 0x50, 0x25, 0x10, 0x42,
	0xc2, 0x57, 0x25, 0x0c, 0xcc, 0x5d, 0x19, 0x80, 0x8e, 0x2c, 0x95, 0xdd, 0xba, 0x26, 0x52, 0xe2,
	0xeb, 0xc6, 0x4e, 0x44, 0xde, 0x82, 0xc7, 0xea, 0xd8, 0x91, 0x09, 0xa1, 0xe4, 0x45, 0x50, 0x62,
	0x23, 0x54, 0xc1, 0x76, 0x8f, 0xcf, 0xe7, 0xe3, 0xe3, 0x1b, 0x9d, 0xaf, 0x85, 0x29, 0xd0, 0x80,
	0x2e, 0x71, 0x93, 0xe5, 0xc2, 0x40, 0x9d, 0xc2, 0xb6, 0x12, 0x65, 0xb3, 0xd4, 0xac, 0x64, 0x85,
	0xa1, 0xba, 0x44, 0x8b, 0x71, 0xec, 0x30, 0xfa, 0x8d, 0xd1, 0x3a, 0x4d, 0xc6, 0x12, 0x25, 0x0e,
	0x36, 0xf4, 0x93, 0x23, 0x93, 0x13, 0x89, 0x28, 0x73, 0x01, 0x4c, 0x67, 0xc0, 0x94, 0x42, 0xcb,
	0x6c, 0x86, 0xca, 0xe7, 0x24, 0x13, 0xef, 0x0e, 0x8a, 0x57, 0x1b, 0x60, 0xaa, 0xf1, 0xd6, 0xc5,
	0x1f, 0x4d, 0x0a, 0x5c, 0x8b, 0xdc, 0x1c, 0x54, 0x49, 0x26, 0x2b, 0xec, 0xb9, 0xa5, 0x7b, 0xd9,
	0x09, 0x6f, 0x5d, 0x39, 0x05, 0x9c, 0x19, 0xe1, 0x7e, 0x01, 0xf5, 0x8c, 0x0b, 0xcb, 0x66, 0xa0,
	0x99, 0xcc, 0xd4, 0x50, 0xc5, 0xb1, 0xd3, 0x71, 0x14, 0x3f, 0xf5, 0xc4, 0xe3, 0x90, 0xbd, 0x10,
	0xdb, 0x4a, 0x18, 0x3b, 0x7d, 0x88, 0x8e, 0x0e, 0x4e, 0x8d, 0x46, 0x65, 0x44, 0x7c, 0x17, 0x8d,
	0x5c, 0x87, 0xe3, 0xf0, 0x2c, 0xbc, 0xfc, 0x7f, 0x93, 0xd0, 0xdf, 0xfb, 0xa0, 0xee, 0xce, 0xfc,
	0xdf, 0xee, 0xe3, 0x34, 0x58, 0x78, 0x7e, 0x7e, 0xbf, 0x6b, 0x49, 0xb8, 0x6f, 0x49, 0xf8, 0xd9,
	0x92, 0xf0, 0xad, 0x23, 0xc1, 0xbe, 0x23, 0xc1, 0x7b, 0x47, 0x82, 0xe7, 0x99, 0xcc, 0xec, 0x4b,
	0xc5, 0xe9, 0x0a, 0x0b, 0x70, 0x69, 0xd7, 0x39, 0xe3, 0xc6, 0xcf, 0x50, 0xdf, 0xc2, 0xeb, 0xcf,
	0x2e, 0x6c, 0xa3, 0x85, 0xe1, 0xa3, 0xa1, 0x7a, 0xfa, 0x15, 0x00, 0x00, 0xff, 0xff, 0xc4, 0x40,
	0x1c, 0xd7, 0xb5, 0x01, 0x00, 0x00,
}

func (m *QueryParamsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQueryParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintQueryParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovQueryParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryParamsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovQueryParams(uint64(l))
	return n
}

func sovQueryParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQueryParams(x uint64) (n int) {
	return sovQueryParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryParamsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQueryParams
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
			return fmt.Errorf("proto: QueryParamsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQueryParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQueryParams
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
func (m *QueryParamsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQueryParams
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
			return fmt.Errorf("proto: QueryParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryParams
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
				return ErrInvalidLengthQueryParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQueryParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQueryParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQueryParams
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
func skipQueryParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQueryParams
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
					return 0, ErrIntOverflowQueryParams
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
					return 0, ErrIntOverflowQueryParams
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
				return 0, ErrInvalidLengthQueryParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQueryParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQueryParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQueryParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQueryParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQueryParams = fmt.Errorf("proto: unexpected end of group")
)
