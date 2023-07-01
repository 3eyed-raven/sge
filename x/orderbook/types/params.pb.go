// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sge/orderbook/params.proto

package types

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/gogo/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = proto.Marshal
	_ = fmt.Errorf
	_ = math.Inf
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Params defines the parameters for the orderbook module.
type Params struct {
	// max_order_book_participations is the maximum number of participations per
	// book.
	MaxOrderBookParticipations uint64 `protobuf:"varint,1,opt,name=max_order_book_participations,json=maxOrderBookParticipations,proto3" json:"max_order_book_participations,omitempty" yaml:"max_order_book_participations"`
	// batch_settlement_count is the batch settlement deposit count.
	BatchSettlementCount uint64 `protobuf:"varint,2,opt,name=batch_settlement_count,json=batchSettlementCount,proto3" json:"batch_settlement_count,omitempty" yaml:"batch_settlement_count"`
	// requeue_threshold is the threshold at which a participation is re-queued in
	// orderbook.
	RequeueThreshold uint64 `protobuf:"varint,3,opt,name=requeue_threshold,json=requeueThreshold,proto3" json:"requeue_threshold,omitempty" yaml:"requeue_threshold"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_3c3d2056747df0cc, []int{0}
}

func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}

func (m *Params) XXX_Size() int {
	return m.Size()
}

func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetMaxOrderBookParticipations() uint64 {
	if m != nil {
		return m.MaxOrderBookParticipations
	}
	return 0
}

func (m *Params) GetBatchSettlementCount() uint64 {
	if m != nil {
		return m.BatchSettlementCount
	}
	return 0
}

func (m *Params) GetRequeueThreshold() uint64 {
	if m != nil {
		return m.RequeueThreshold
	}
	return 0
}

func init() {
	proto.RegisterType((*Params)(nil), "sgenetwork.sge.orderbook.Params")
}

func init() { proto.RegisterFile("sge/orderbook/params.proto", fileDescriptor_3c3d2056747df0cc) }

var fileDescriptor_3c3d2056747df0cc = []byte{
	// 322 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0x31, 0x4b, 0xfb, 0x50,
	0x14, 0xc5, 0x93, 0xfe, 0xff, 0x14, 0xc9, 0xa4, 0xa1, 0x48, 0x28, 0x36, 0xd1, 0xe0, 0xd0, 0x41,
	0x93, 0xc1, 0xad, 0x63, 0x05, 0xc1, 0xc9, 0x52, 0x05, 0xc1, 0x25, 0xbc, 0xa4, 0x97, 0x97, 0x90,
	0x26, 0x37, 0xbe, 0x77, 0x83, 0xed, 0xb7, 0x70, 0x74, 0x2c, 0xf8, 0x65, 0x1c, 0x3b, 0x3a, 0x15,
	0x69, 0x17, 0xe7, 0x7e, 0x02, 0xc9, 0x4b, 0x2d, 0x8a, 0xe2, 0x76, 0x39, 0xe7, 0x77, 0xcf, 0x85,
	0x73, 0x8d, 0xb6, 0xe4, 0xe0, 0xa3, 0x18, 0x81, 0x08, 0x11, 0x53, 0xbf, 0x60, 0x82, 0x65, 0xd2,
	0x2b, 0x04, 0x12, 0x9a, 0x96, 0xe4, 0x90, 0x03, 0x3d, 0xa0, 0x48, 0x3d, 0xc9, 0xc1, 0xdb, 0x62,
	0xed, 0x16, 0x47, 0x8e, 0x0a, 0xf2, 0xab, 0xa9, 0xe6, 0xdd, 0xe7, 0x86, 0xd1, 0x1c, 0xa8, 0x00,
	0x33, 0x35, 0x3a, 0x19, 0x9b, 0x04, 0x6a, 0x23, 0xa8, 0x56, 0x82, 0x82, 0x09, 0x4a, 0xa2, 0xa4,
	0x60, 0x94, 0x60, 0x2e, 0x2d, 0xfd, 0x50, 0xef, 0xfe, 0xef, 0x77, 0xd7, 0x0b, 0xe7, 0x78, 0xca,
	0xb2, 0x71, 0xcf, 0xfd, 0x13, 0x77, 0x87, 0xed, 0x8c, 0x4d, 0xae, 0x2a, 0xbb, 0x8f, 0x98, 0x0e,
	0xbe, 0x99, 0xe6, 0xad, 0xb1, 0x1f, 0x32, 0x8a, 0xe2, 0x40, 0x02, 0xd1, 0x18, 0x32, 0xc8, 0x29,
	0x88, 0xb0, 0xcc, 0xc9, 0x6a, 0xa8, 0x2b, 0x47, 0xeb, 0x85, 0xd3, 0xa9, 0xaf, 0xfc, 0xce, 0xb9,
	0xc3, 0x96, 0x32, 0xae, 0xb7, 0xfa, 0x79, 0x25, 0x9b, 0x97, 0xc6, 0x9e, 0x80, 0xfb, 0x12, 0x4a,
	0x08, 0x28, 0x16, 0x20, 0x63, 0x1c, 0x8f, 0xac, 0x7f, 0x2a, 0xf3, 0x60, 0xbd, 0x70, 0xac, 0x3a,
	0xf3, 0x07, 0xe2, 0x0e, 0x77, 0x37, 0xda, 0xcd, 0xa7, 0xd4, 0xdb, 0x79, 0x9a, 0x39, 0xda, 0xfb,
	0xcc, 0xd1, 0xfb, 0x17, 0x2f, 0x4b, 0x5b, 0x9f, 0x2f, 0x6d, 0xfd, 0x6d, 0x69, 0xeb, 0x8f, 0x2b,
	0x5b, 0x9b, 0xaf, 0x6c, 0xed, 0x75, 0x65, 0x6b, 0x77, 0x27, 0x3c, 0xa1, 0xb8, 0x0c, 0xbd, 0x08,
	0x33, 0x5f, 0x72, 0x38, 0xdd, 0x74, 0x5f, 0xcd, 0xfe, 0xe4, 0xcb, 0x93, 0x68, 0x5a, 0x80, 0x0c,
	0x9b, 0xaa, 0xf4, 0xb3, 0x8f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x5c, 0x56, 0xfa, 0x6d, 0xc2, 0x01,
	0x00, 0x00,
}

func (this *Params) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Params)
	if !ok {
		that2, ok := that.(Params)
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
	if this.MaxOrderBookParticipations != that1.MaxOrderBookParticipations {
		return false
	}
	if this.BatchSettlementCount != that1.BatchSettlementCount {
		return false
	}
	if this.RequeueThreshold != that1.RequeueThreshold {
		return false
	}
	return true
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.RequeueThreshold != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.RequeueThreshold))
		i--
		dAtA[i] = 0x18
	}
	if m.BatchSettlementCount != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.BatchSettlementCount))
		i--
		dAtA[i] = 0x10
	}
	if m.MaxOrderBookParticipations != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.MaxOrderBookParticipations))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}

func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MaxOrderBookParticipations != 0 {
		n += 1 + sovParams(uint64(m.MaxOrderBookParticipations))
	}
	if m.BatchSettlementCount != 0 {
		n += 1 + sovParams(uint64(m.BatchSettlementCount))
	}
	if m.RequeueThreshold != 0 {
		n += 1 + sovParams(uint64(m.RequeueThreshold))
	}
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}

func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}

func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxOrderBookParticipations", wireType)
			}
			m.MaxOrderBookParticipations = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxOrderBookParticipations |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BatchSettlementCount", wireType)
			}
			m.BatchSettlementCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BatchSettlementCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequeueThreshold", wireType)
			}
			m.RequeueThreshold = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RequeueThreshold |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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

func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)
