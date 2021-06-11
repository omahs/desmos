// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: desmos/profiles/v1beta1/query_dtag_requests.proto

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

// QueryDTagTransfersRequest is the request type for the Query/DTagTransfers RPC
// endpoint
type QueryDTagTransfersRequest struct {
	// Address or DTag of the user to query the transfer requests for
	User string `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (m *QueryDTagTransfersRequest) Reset()         { *m = QueryDTagTransfersRequest{} }
func (m *QueryDTagTransfersRequest) String() string { return proto.CompactTextString(m) }
func (*QueryDTagTransfersRequest) ProtoMessage()    {}
func (*QueryDTagTransfersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2be2e173228b210d, []int{0}
}
func (m *QueryDTagTransfersRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryDTagTransfersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryDTagTransfersRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryDTagTransfersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryDTagTransfersRequest.Merge(m, src)
}
func (m *QueryDTagTransfersRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryDTagTransfersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryDTagTransfersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryDTagTransfersRequest proto.InternalMessageInfo

// QueryDTagTransfersResponse is the response type for the Query/DTagTransfers
// RPC method.
type QueryDTagTransfersResponse struct {
	// relationships represent the list of all the blocks for the queried user
	Requests []DTagTransferRequest `protobuf:"bytes,1,rep,name=requests,proto3" json:"requests"`
}

func (m *QueryDTagTransfersResponse) Reset()         { *m = QueryDTagTransfersResponse{} }
func (m *QueryDTagTransfersResponse) String() string { return proto.CompactTextString(m) }
func (*QueryDTagTransfersResponse) ProtoMessage()    {}
func (*QueryDTagTransfersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2be2e173228b210d, []int{1}
}
func (m *QueryDTagTransfersResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryDTagTransfersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryDTagTransfersResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryDTagTransfersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryDTagTransfersResponse.Merge(m, src)
}
func (m *QueryDTagTransfersResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryDTagTransfersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryDTagTransfersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryDTagTransfersResponse proto.InternalMessageInfo

func (m *QueryDTagTransfersResponse) GetRequests() []DTagTransferRequest {
	if m != nil {
		return m.Requests
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryDTagTransfersRequest)(nil), "desmos.profiles.v1beta1.QueryDTagTransfersRequest")
	proto.RegisterType((*QueryDTagTransfersResponse)(nil), "desmos.profiles.v1beta1.QueryDTagTransfersResponse")
}

func init() {
	proto.RegisterFile("desmos/profiles/v1beta1/query_dtag_requests.proto", fileDescriptor_2be2e173228b210d)
}

var fileDescriptor_2be2e173228b210d = []byte{
	// 331 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x50, 0xbf, 0x4e, 0x32, 0x41,
	0x10, 0xbf, 0xcb, 0x47, 0xbe, 0xf0, 0xdd, 0xd7, 0x5d, 0x4c, 0x84, 0x8b, 0x39, 0x08, 0x15, 0x31,
	0xba, 0x1b, 0xb0, 0xd2, 0x92, 0x58, 0x58, 0x99, 0x48, 0xa8, 0x6c, 0xc8, 0x1e, 0x0c, 0xeb, 0x25,
	0xc7, 0xce, 0xb1, 0xb3, 0x67, 0xe4, 0x0d, 0x2c, 0x7d, 0x04, 0x1e, 0x87, 0x92, 0xd2, 0xca, 0x18,
	0x68, 0x7c, 0x0c, 0xc3, 0xed, 0x82, 0x16, 0xd0, 0xcd, 0xe4, 0xf7, 0x77, 0x26, 0xe8, 0x8c, 0x81,
	0xa6, 0x48, 0x3c, 0xd7, 0x38, 0x49, 0x33, 0x20, 0xfe, 0xdc, 0x49, 0xc0, 0x88, 0x0e, 0x9f, 0x15,
	0xa0, 0xe7, 0xc3, 0xb1, 0x11, 0x72, 0xa8, 0x61, 0x56, 0x00, 0x19, 0x62, 0xb9, 0x46, 0x83, 0xe1,
	0xa9, 0x95, 0xb0, 0x9d, 0x84, 0x39, 0x49, 0x74, 0x22, 0x51, 0x62, 0xc9, 0xe1, 0xdb, 0xc9, 0xd2,
	0xa3, 0x33, 0x89, 0x28, 0x33, 0xe0, 0x22, 0x4f, 0xb9, 0x50, 0x0a, 0x8d, 0x30, 0x29, 0x2a, 0x67,
	0x16, 0xd5, 0x1d, 0x5a, 0x6e, 0x49, 0x31, 0xe1, 0x42, 0xcd, 0x1d, 0xd4, 0x3d, 0x56, 0x6d, 0x8a,
	0x63, 0xc8, 0xe8, 0x50, 0xb7, 0xa8, 0x3e, 0xc2, 0xad, 0x66, 0x68, 0x5b, 0xd8, 0xc5, 0x41, 0xe7,
	0x76, 0xe3, 0x89, 0x20, 0xb0, 0xd7, 0xed, 0x0d, 0x73, 0x21, 0x53, 0x55, 0xd6, 0xb2, 0xdc, 0xd6,
	0x75, 0x50, 0x7f, 0xd8, 0x32, 0x6e, 0x07, 0x42, 0x0e, 0xb4, 0x50, 0x34, 0x01, 0x4d, 0x7d, 0x1b,
	0x15, 0x86, 0x41, 0xa5, 0x20, 0xd0, 0x35, 0xbf, 0xe9, 0xb7, 0xff, 0xf5, 0xcb, 0xf9, 0xa6, 0xfa,
	0xba, 0x68, 0x78, 0x5f, 0x8b, 0x86, 0xd7, 0xca, 0x82, 0xe8, 0x90, 0x94, 0x72, 0x54, 0x04, 0xe1,
	0x7d, 0x50, 0xdd, 0x35, 0xae, 0xf9, 0xcd, 0x3f, 0xed, 0xff, 0xdd, 0x0b, 0x76, 0xe4, 0x9d, 0xec,
	0xb7, 0x83, 0xcb, 0xee, 0x55, 0x96, 0x1f, 0x0d, 0xaf, 0xbf, 0xf7, 0xe8, 0xdd, 0x2d, 0xd7, 0xb1,
	0xbf, 0x5a, 0xc7, 0xfe, 0xe7, 0x3a, 0xf6, 0xdf, 0x36, 0xb1, 0xb7, 0xda, 0xc4, 0xde, 0xfb, 0x26,
	0xf6, 0x1e, 0x99, 0x4c, 0xcd, 0x53, 0x91, 0xb0, 0x11, 0x4e, 0xb9, 0x4d, 0xb8, 0xcc, 0x44, 0x42,
	0x6e, 0xe6, 0x2f, 0x3f, 0x6f, 0x35, 0xf3, 0x1c, 0x28, 0xf9, 0x5b, 0x5e, 0x7e, 0xf5, 0x1d, 0x00,
	0x00, 0xff, 0xff, 0x8f, 0x90, 0x6a, 0x69, 0x11, 0x02, 0x00, 0x00,
}

func (m *QueryDTagTransfersRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryDTagTransfersRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryDTagTransfersRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.User) > 0 {
		i -= len(m.User)
		copy(dAtA[i:], m.User)
		i = encodeVarintQueryDtagRequests(dAtA, i, uint64(len(m.User)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryDTagTransfersResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryDTagTransfersResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryDTagTransfersResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Requests) > 0 {
		for iNdEx := len(m.Requests) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Requests[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQueryDtagRequests(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintQueryDtagRequests(dAtA []byte, offset int, v uint64) int {
	offset -= sovQueryDtagRequests(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryDTagTransfersRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.User)
	if l > 0 {
		n += 1 + l + sovQueryDtagRequests(uint64(l))
	}
	return n
}

func (m *QueryDTagTransfersResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Requests) > 0 {
		for _, e := range m.Requests {
			l = e.Size()
			n += 1 + l + sovQueryDtagRequests(uint64(l))
		}
	}
	return n
}

func sovQueryDtagRequests(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQueryDtagRequests(x uint64) (n int) {
	return sovQueryDtagRequests(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryDTagTransfersRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQueryDtagRequests
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
			return fmt.Errorf("proto: QueryDTagTransfersRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryDTagTransfersRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field User", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryDtagRequests
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
				return ErrInvalidLengthQueryDtagRequests
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQueryDtagRequests
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.User = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQueryDtagRequests(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQueryDtagRequests
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
func (m *QueryDTagTransfersResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQueryDtagRequests
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
			return fmt.Errorf("proto: QueryDTagTransfersResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryDTagTransfersResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Requests", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryDtagRequests
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
				return ErrInvalidLengthQueryDtagRequests
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQueryDtagRequests
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Requests = append(m.Requests, DTagTransferRequest{})
			if err := m.Requests[len(m.Requests)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQueryDtagRequests(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQueryDtagRequests
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
func skipQueryDtagRequests(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQueryDtagRequests
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
					return 0, ErrIntOverflowQueryDtagRequests
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
					return 0, ErrIntOverflowQueryDtagRequests
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
				return 0, ErrInvalidLengthQueryDtagRequests
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQueryDtagRequests
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQueryDtagRequests
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQueryDtagRequests        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQueryDtagRequests          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQueryDtagRequests = fmt.Errorf("proto: unexpected end of group")
)