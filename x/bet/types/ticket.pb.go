// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sge/bet/ticket.proto

package types

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/sge-network/sge/types"
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

// BetPlacementTicketPayload indicates data of bet placement ticket.
type BetPlacementTicketPayload struct {
	// selected_odds is the user-selected odds to place bet.
	SelectedOdds *BetOdds `protobuf:"bytes,1,opt,name=selected_odds,json=selectedOdds,proto3" json:"selected_odds,omitempty"`
	// kyc_data contains the details of user kyc.
	KycData types.KycDataPayload `protobuf:"bytes,2,opt,name=kyc_data,json=kycData,proto3" json:"kyc_data"`
	// odds_type is the type of odds that are going to be placed
	// such as decimal, fraction, moneyline.
	OddsType OddsType `protobuf:"varint,3,opt,name=odds_type,json=oddsType,proto3,enum=sgenetwork.sge.bet.OddsType" json:"odds_type,omitempty"`
}

func (m *BetPlacementTicketPayload) Reset()         { *m = BetPlacementTicketPayload{} }
func (m *BetPlacementTicketPayload) String() string { return proto.CompactTextString(m) }
func (*BetPlacementTicketPayload) ProtoMessage()    {}
func (*BetPlacementTicketPayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_cf6959e7db451613, []int{0}
}

func (m *BetPlacementTicketPayload) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *BetPlacementTicketPayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BetPlacementTicketPayload.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *BetPlacementTicketPayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BetPlacementTicketPayload.Merge(m, src)
}

func (m *BetPlacementTicketPayload) XXX_Size() int {
	return m.Size()
}

func (m *BetPlacementTicketPayload) XXX_DiscardUnknown() {
	xxx_messageInfo_BetPlacementTicketPayload.DiscardUnknown(m)
}

var xxx_messageInfo_BetPlacementTicketPayload proto.InternalMessageInfo

func (m *BetPlacementTicketPayload) GetSelectedOdds() *BetOdds {
	if m != nil {
		return m.SelectedOdds
	}
	return nil
}

func (m *BetPlacementTicketPayload) GetKycData() types.KycDataPayload {
	if m != nil {
		return m.KycData
	}
	return types.KycDataPayload{}
}

func (m *BetPlacementTicketPayload) GetOddsType() OddsType {
	if m != nil {
		return m.OddsType
	}
	return OddsType_ODDS_TYPE_UNSPECIFIED
}

func init() {
	proto.RegisterType((*BetPlacementTicketPayload)(nil), "sgenetwork.sge.bet.BetPlacementTicketPayload")
}

func init() { proto.RegisterFile("sge/bet/ticket.proto", fileDescriptor_cf6959e7db451613) }

var fileDescriptor_cf6959e7db451613 = []byte{
	// 307 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xcf, 0x4e, 0x32, 0x31,
	0x14, 0xc5, 0xa7, 0xdf, 0x67, 0x14, 0xc7, 0x3f, 0x8b, 0x09, 0x51, 0x44, 0x53, 0x89, 0x26, 0x86,
	0x8d, 0x6d, 0x82, 0x2b, 0x77, 0x64, 0xc2, 0xce, 0x85, 0x84, 0xb0, 0x72, 0x43, 0x3a, 0xed, 0x4d,
	0x25, 0x03, 0x74, 0x42, 0xaf, 0xd1, 0xbe, 0x85, 0x8f, 0xc5, 0x92, 0xa5, 0x2b, 0x62, 0xe0, 0x45,
	0x4c, 0x3b, 0x83, 0x0b, 0x75, 0x77, 0xe6, 0xdc, 0x7b, 0xce, 0xfd, 0x65, 0x1a, 0xd7, 0xad, 0x06,
	0x9e, 0x01, 0x72, 0x1c, 0xcb, 0x1c, 0x90, 0x15, 0x73, 0x83, 0x26, 0x49, 0xac, 0x86, 0x19, 0xe0,
	0xab, 0x99, 0xe7, 0xcc, 0x6a, 0x60, 0x19, 0x60, 0xb3, 0xae, 0x8d, 0x36, 0x61, 0xcc, 0xbd, 0x2a,
	0x37, 0x9b, 0x7e, 0x93, 0xa3, 0x2b, 0x80, 0xe7, 0x4e, 0x56, 0xde, 0xc9, 0xb6, 0x33, 0x03, 0x1c,
	0x19, 0xa5, 0x6c, 0xe5, 0x9f, 0x6e, 0x7d, 0xef, 0x8d, 0x7c, 0xa8, 0x1c, 0x5c, 0xad, 0x48, 0x7c,
	0x96, 0x02, 0xf6, 0x27, 0x42, 0xc2, 0x14, 0x66, 0x38, 0x0c, 0x2c, 0x7d, 0xe1, 0x26, 0x46, 0xa8,
	0xa4, 0x1b, 0x1f, 0x59, 0x98, 0x80, 0x44, 0x50, 0xa1, 0xad, 0x41, 0x5a, 0xa4, 0x7d, 0xd0, 0x39,
	0x67, 0xbf, 0x21, 0x59, 0x0a, 0xf8, 0xa8, 0x94, 0x1d, 0x1c, 0x6e, 0x13, 0xfe, 0x2b, 0xe9, 0xc5,
	0xb5, 0xdc, 0xc9, 0x91, 0x12, 0x28, 0x1a, 0xff, 0x42, 0xf8, 0xfa, 0x67, 0x38, 0xd0, 0x3c, 0x38,
	0xd9, 0x13, 0x28, 0xaa, 0xc3, 0xe9, 0xce, 0x62, 0x75, 0x19, 0x0d, 0xf6, 0xf2, 0xd2, 0x4d, 0xee,
	0xe3, 0xfd, 0x6f, 0xf0, 0xc6, 0xff, 0x16, 0x69, 0x1f, 0x77, 0x2e, 0xfe, 0x62, 0xf0, 0x27, 0x87,
	0xae, 0x80, 0x41, 0xcd, 0x54, 0x2a, 0xed, 0x2e, 0xd6, 0x94, 0x2c, 0xd7, 0x94, 0x7c, 0xae, 0x29,
	0x79, 0xdf, 0xd0, 0x68, 0xb9, 0xa1, 0xd1, 0xc7, 0x86, 0x46, 0x4f, 0x37, 0x7a, 0x8c, 0xcf, 0x2f,
	0x19, 0x93, 0x66, 0xca, 0xad, 0x86, 0xdb, 0xaa, 0xcc, 0x6b, 0xfe, 0x56, 0x3e, 0x8c, 0x2b, 0xc0,
	0x66, 0xbb, 0xe1, 0x4f, 0xdd, 0x7d, 0x05, 0x00, 0x00, 0xff, 0xff, 0xe5, 0x3a, 0xe0, 0x84, 0xb0,
	0x01, 0x00, 0x00,
}

func (m *BetPlacementTicketPayload) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BetPlacementTicketPayload) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BetPlacementTicketPayload) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.OddsType != 0 {
		i = encodeVarintTicket(dAtA, i, uint64(m.OddsType))
		i--
		dAtA[i] = 0x18
	}
	{
		size, err := m.KycData.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTicket(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.SelectedOdds != nil {
		{
			size, err := m.SelectedOdds.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTicket(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTicket(dAtA []byte, offset int, v uint64) int {
	offset -= sovTicket(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}

func (m *BetPlacementTicketPayload) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SelectedOdds != nil {
		l = m.SelectedOdds.Size()
		n += 1 + l + sovTicket(uint64(l))
	}
	l = m.KycData.Size()
	n += 1 + l + sovTicket(uint64(l))
	if m.OddsType != 0 {
		n += 1 + sovTicket(uint64(m.OddsType))
	}
	return n
}

func sovTicket(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}

func sozTicket(x uint64) (n int) {
	return sovTicket(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}

func (m *BetPlacementTicketPayload) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTicket
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
			return fmt.Errorf("proto: BetPlacementTicketPayload: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BetPlacementTicketPayload: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SelectedOdds", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicket
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
				return ErrInvalidLengthTicket
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTicket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.SelectedOdds == nil {
				m.SelectedOdds = &BetOdds{}
			}
			if err := m.SelectedOdds.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field KycData", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicket
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
				return ErrInvalidLengthTicket
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTicket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.KycData.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OddsType", wireType)
			}
			m.OddsType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.OddsType |= OddsType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTicket(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTicket
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

func skipTicket(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTicket
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
					return 0, ErrIntOverflowTicket
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
					return 0, ErrIntOverflowTicket
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
				return 0, ErrInvalidLengthTicket
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTicket
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTicket
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTicket        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTicket          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTicket = fmt.Errorf("proto: unexpected end of group")
)
