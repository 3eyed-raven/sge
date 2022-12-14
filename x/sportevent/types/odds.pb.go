// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sge/sportevent/odds.proto

package types

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

type Odds struct {
	UID     string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid"`
	Details string `protobuf:"bytes,2,opt,name=details,proto3" json:"details,omitempty"`
}

func (m *Odds) Reset()         { *m = Odds{} }
func (m *Odds) String() string { return proto.CompactTextString(m) }
func (*Odds) ProtoMessage()    {}
func (*Odds) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4eaa1643bcc33f0, []int{0}
}
func (m *Odds) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Odds) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Odds.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Odds) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Odds.Merge(m, src)
}
func (m *Odds) XXX_Size() int {
	return m.Size()
}
func (m *Odds) XXX_DiscardUnknown() {
	xxx_messageInfo_Odds.DiscardUnknown(m)
}

var xxx_messageInfo_Odds proto.InternalMessageInfo

func (m *Odds) GetUID() string {
	if m != nil {
		return m.UID
	}
	return ""
}

func (m *Odds) GetDetails() string {
	if m != nil {
		return m.Details
	}
	return ""
}

func init() {
	proto.RegisterType((*Odds)(nil), "sgenetwork.sge.sportevent.Odds")
}

func init() { proto.RegisterFile("sge/sportevent/odds.proto", fileDescriptor_c4eaa1643bcc33f0) }

var fileDescriptor_c4eaa1643bcc33f0 = []byte{
	// 200 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2c, 0x4e, 0x4f, 0xd5,
	0x2f, 0x2e, 0xc8, 0x2f, 0x2a, 0x49, 0x2d, 0x4b, 0xcd, 0x2b, 0xd1, 0xcf, 0x4f, 0x49, 0x29, 0xd6,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x02, 0x49, 0xe5, 0xa5, 0x96, 0x94, 0xe7, 0x17, 0x65, 0xeb,
	0x15, 0xa7, 0xa7, 0xea, 0x21, 0x54, 0x49, 0x89, 0xa4, 0xe7, 0xa7, 0xe7, 0x83, 0x55, 0xe9, 0x83,
	0x58, 0x10, 0x0d, 0x4a, 0x4e, 0x5c, 0x2c, 0xfe, 0x29, 0x29, 0xc5, 0x42, 0x0a, 0x5c, 0xcc, 0xa5,
	0x99, 0x29, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x4e, 0x7c, 0x8f, 0xee, 0xc9, 0x33, 0x87, 0x7a,
	0xba, 0xbc, 0xba, 0x27, 0x0f, 0x12, 0x0d, 0x02, 0x11, 0x42, 0x12, 0x5c, 0xec, 0x29, 0xa9, 0x25,
	0x89, 0x99, 0x39, 0xc5, 0x12, 0x4c, 0x20, 0x55, 0x41, 0x30, 0xae, 0x93, 0xfb, 0x89, 0x47, 0x72,
	0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0xc3, 0x85, 0xc7,
	0x72, 0x0c, 0x37, 0x1e, 0xcb, 0x31, 0x44, 0xe9, 0xa6, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25,
	0xe7, 0xe7, 0xea, 0x17, 0xa7, 0xa7, 0xea, 0x42, 0x9d, 0x06, 0x62, 0xeb, 0x57, 0x20, 0x7b, 0xa1,
	0xa4, 0xb2, 0x20, 0xb5, 0x38, 0x89, 0x0d, 0xec, 0x26, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x7b, 0x4b, 0xa6, 0x8a, 0xe1, 0x00, 0x00, 0x00,
}

func (m *Odds) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Odds) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Odds) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Details) > 0 {
		i -= len(m.Details)
		copy(dAtA[i:], m.Details)
		i = encodeVarintOdds(dAtA, i, uint64(len(m.Details)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.UID) > 0 {
		i -= len(m.UID)
		copy(dAtA[i:], m.UID)
		i = encodeVarintOdds(dAtA, i, uint64(len(m.UID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintOdds(dAtA []byte, offset int, v uint64) int {
	offset -= sovOdds(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Odds) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.UID)
	if l > 0 {
		n += 1 + l + sovOdds(uint64(l))
	}
	l = len(m.Details)
	if l > 0 {
		n += 1 + l + sovOdds(uint64(l))
	}
	return n
}

func sovOdds(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozOdds(x uint64) (n int) {
	return sovOdds(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Odds) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOdds
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
			return fmt.Errorf("proto: Odds: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Odds: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOdds
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
				return ErrInvalidLengthOdds
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOdds
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Details", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOdds
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
				return ErrInvalidLengthOdds
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOdds
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Details = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipOdds(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOdds
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
func skipOdds(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowOdds
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
					return 0, ErrIntOverflowOdds
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
					return 0, ErrIntOverflowOdds
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
				return 0, ErrInvalidLengthOdds
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupOdds
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthOdds
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthOdds        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowOdds          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupOdds = fmt.Errorf("proto: unexpected end of group")
)
