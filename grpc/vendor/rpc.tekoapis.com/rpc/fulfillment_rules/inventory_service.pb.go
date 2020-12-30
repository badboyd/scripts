// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/fulfillment_rules/v1/inventory_service.proto

package fulfillment_rules

import (
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// GetInventoryRequest
type GetInventoryRequest struct {
	SellerId             uint32   `protobuf:"varint,1,opt,name=seller_id,proto3" json:"seller_id,omitempty"`
	WardId               string   `protobuf:"bytes,2,opt,name=ward_id,proto3" json:"ward_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetInventoryRequest) Reset()         { *m = GetInventoryRequest{} }
func (m *GetInventoryRequest) String() string { return proto.CompactTextString(m) }
func (*GetInventoryRequest) ProtoMessage()    {}
func (*GetInventoryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_67643d0ca5430dc5, []int{0}
}
func (m *GetInventoryRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetInventoryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetInventoryRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetInventoryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetInventoryRequest.Merge(m, src)
}
func (m *GetInventoryRequest) XXX_Size() int {
	return m.Size()
}
func (m *GetInventoryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetInventoryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetInventoryRequest proto.InternalMessageInfo

func (m *GetInventoryRequest) GetSellerId() uint32 {
	if m != nil {
		return m.SellerId
	}
	return 0
}

func (m *GetInventoryRequest) GetWardId() string {
	if m != nil {
		return m.WardId
	}
	return ""
}

// GetInventoryResponse
type GetInventoryResponse struct {
	SiteId               string   `protobuf:"bytes,1,opt,name=site_id,proto3" json:"site_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetInventoryResponse) Reset()         { *m = GetInventoryResponse{} }
func (m *GetInventoryResponse) String() string { return proto.CompactTextString(m) }
func (*GetInventoryResponse) ProtoMessage()    {}
func (*GetInventoryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_67643d0ca5430dc5, []int{1}
}
func (m *GetInventoryResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetInventoryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetInventoryResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetInventoryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetInventoryResponse.Merge(m, src)
}
func (m *GetInventoryResponse) XXX_Size() int {
	return m.Size()
}
func (m *GetInventoryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetInventoryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetInventoryResponse proto.InternalMessageInfo

func (m *GetInventoryResponse) GetSiteId() string {
	if m != nil {
		return m.SiteId
	}
	return ""
}

func init() {
	proto.RegisterType((*GetInventoryRequest)(nil), "fulfillment_rules.v1.GetInventoryRequest")
	proto.RegisterType((*GetInventoryResponse)(nil), "fulfillment_rules.v1.GetInventoryResponse")
}

func init() {
	proto.RegisterFile("proto/fulfillment_rules/v1/inventory_service.proto", fileDescriptor_67643d0ca5430dc5)
}

var fileDescriptor_67643d0ca5430dc5 = []byte{
	// 304 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x31, 0x4b, 0x03, 0x31,
	0x1c, 0xc5, 0x9b, 0x22, 0x96, 0x86, 0x0a, 0x12, 0x0b, 0x96, 0x22, 0x47, 0x29, 0x28, 0xb5, 0x43,
	0x62, 0xeb, 0xe2, 0x7c, 0x8b, 0xb8, 0x9e, 0xe0, 0xe0, 0x72, 0xc4, 0xbb, 0x7f, 0x4b, 0x30, 0x4d,
	0x62, 0x92, 0x3b, 0xe9, 0xea, 0xee, 0xe4, 0xe2, 0x47, 0x72, 0x14, 0xfc, 0x02, 0x72, 0xfa, 0x29,
	0x9c, 0xa4, 0x3d, 0xce, 0x2a, 0xed, 0xe0, 0x96, 0x3f, 0xff, 0xdf, 0x7b, 0xc9, 0x7b, 0xc1, 0x63,
	0x63, 0xb5, 0xd7, 0x6c, 0x92, 0xc9, 0x89, 0x90, 0x72, 0x06, 0xca, 0xc7, 0x36, 0x93, 0xe0, 0x58,
	0x3e, 0x62, 0x42, 0xe5, 0xa0, 0xbc, 0xb6, 0xf3, 0xd8, 0x81, 0xcd, 0x45, 0x02, 0x74, 0x09, 0x93,
	0xf6, 0x1a, 0x4d, 0xf3, 0x51, 0xf7, 0x60, 0xaa, 0xf5, 0x54, 0x02, 0xe3, 0x46, 0x30, 0xae, 0x94,
	0xf6, 0xdc, 0x0b, 0xad, 0x5c, 0xa9, 0xe9, 0xee, 0xe7, 0x5c, 0x8a, 0x94, 0x7b, 0x60, 0xd5, 0xa1,
	0x5c, 0xf4, 0xaf, 0xf0, 0xde, 0x39, 0xf8, 0x8b, 0xea, 0xaa, 0x08, 0xee, 0x32, 0x70, 0x9e, 0x1c,
	0xe2, 0xa6, 0x03, 0x29, 0xc1, 0xc6, 0x22, 0xed, 0xa0, 0x1e, 0x1a, 0xec, 0x84, 0x8d, 0xaf, 0x70,
	0x6b, 0x58, 0xef, 0xd5, 0xa2, 0xd5, 0x86, 0x74, 0x70, 0xe3, 0x9e, 0xdb, 0x74, 0x01, 0xd5, 0x7b,
	0x68, 0xd0, 0x8c, 0xaa, 0xb1, 0x7f, 0x82, 0xdb, 0x7f, 0x7d, 0x9d, 0xd1, 0xca, 0xc1, 0x42, 0xe1,
	0x84, 0x87, 0xca, 0xb6, 0x19, 0x55, 0xe3, 0xf8, 0x11, 0xe1, 0xdd, 0x1f, 0xfe, 0xb2, 0x4c, 0x4c,
	0xe6, 0xb8, 0xf5, 0xdb, 0x86, 0x1c, 0xd3, 0x4d, 0xe1, 0xe9, 0x86, 0x08, 0xdd, 0xe1, 0x7f, 0xd0,
	0xf2, 0x55, 0x7d, 0xf2, 0xf0, 0xf6, 0xf9, 0x54, 0x6f, 0x11, 0xbc, 0x2a, 0x3d, 0x3c, 0x7b, 0x29,
	0x02, 0xf4, 0x5a, 0x04, 0xe8, 0xbd, 0x08, 0xd0, 0xf3, 0x47, 0x50, 0xbb, 0x3e, 0xb2, 0x26, 0xa1,
	0x1e, 0x6e, 0x35, 0x37, 0xc2, 0xd1, 0x44, 0xcf, 0x98, 0x35, 0xc9, 0xfa, 0xdf, 0xdd, 0x6c, 0x2f,
	0xab, 0x3d, 0xfd, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x55, 0x7a, 0xa5, 0x5e, 0xdd, 0x01, 0x00, 0x00,
}

func (m *GetInventoryRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetInventoryRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetInventoryRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.WardId) > 0 {
		i -= len(m.WardId)
		copy(dAtA[i:], m.WardId)
		i = encodeVarintInventoryService(dAtA, i, uint64(len(m.WardId)))
		i--
		dAtA[i] = 0x12
	}
	if m.SellerId != 0 {
		i = encodeVarintInventoryService(dAtA, i, uint64(m.SellerId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *GetInventoryResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetInventoryResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetInventoryResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.SiteId) > 0 {
		i -= len(m.SiteId)
		copy(dAtA[i:], m.SiteId)
		i = encodeVarintInventoryService(dAtA, i, uint64(len(m.SiteId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintInventoryService(dAtA []byte, offset int, v uint64) int {
	offset -= sovInventoryService(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GetInventoryRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SellerId != 0 {
		n += 1 + sovInventoryService(uint64(m.SellerId))
	}
	l = len(m.WardId)
	if l > 0 {
		n += 1 + l + sovInventoryService(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *GetInventoryResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.SiteId)
	if l > 0 {
		n += 1 + l + sovInventoryService(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovInventoryService(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozInventoryService(x uint64) (n int) {
	return sovInventoryService(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GetInventoryRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInventoryService
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
			return fmt.Errorf("proto: GetInventoryRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetInventoryRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SellerId", wireType)
			}
			m.SellerId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInventoryService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SellerId |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WardId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInventoryService
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
				return ErrInvalidLengthInventoryService
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInventoryService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.WardId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipInventoryService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthInventoryService
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthInventoryService
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
func (m *GetInventoryResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInventoryService
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
			return fmt.Errorf("proto: GetInventoryResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetInventoryResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SiteId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInventoryService
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
				return ErrInvalidLengthInventoryService
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInventoryService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SiteId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipInventoryService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthInventoryService
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthInventoryService
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
func skipInventoryService(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowInventoryService
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
					return 0, ErrIntOverflowInventoryService
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
					return 0, ErrIntOverflowInventoryService
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
				return 0, ErrInvalidLengthInventoryService
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupInventoryService
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthInventoryService
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthInventoryService        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowInventoryService          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupInventoryService = fmt.Errorf("proto: unexpected end of group")
)