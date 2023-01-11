// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lottery/lottery/lottery_info.proto

package types

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

type LotteryInfo struct {
	NextId    int64 `protobuf:"varint,1,opt,name=nextId,proto3" json:"nextId,omitempty"`
	NextOrder int64 `protobuf:"varint,2,opt,name=nextOrder,proto3" json:"nextOrder,omitempty"`
	TotalAmt  int64 `protobuf:"varint,3,opt,name=totalAmt,proto3" json:"totalAmt,omitempty"`
}

func (m *LotteryInfo) Reset()         { *m = LotteryInfo{} }
func (m *LotteryInfo) String() string { return proto.CompactTextString(m) }
func (*LotteryInfo) ProtoMessage()    {}
func (*LotteryInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8cadedeb4e07f20, []int{0}
}
func (m *LotteryInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LotteryInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LotteryInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LotteryInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LotteryInfo.Merge(m, src)
}
func (m *LotteryInfo) XXX_Size() int {
	return m.Size()
}
func (m *LotteryInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_LotteryInfo.DiscardUnknown(m)
}

var xxx_messageInfo_LotteryInfo proto.InternalMessageInfo

func (m *LotteryInfo) GetNextId() int64 {
	if m != nil {
		return m.NextId
	}
	return 0
}

func (m *LotteryInfo) GetNextOrder() int64 {
	if m != nil {
		return m.NextOrder
	}
	return 0
}

func (m *LotteryInfo) GetTotalAmt() int64 {
	if m != nil {
		return m.TotalAmt
	}
	return 0
}

func init() {
	proto.RegisterType((*LotteryInfo)(nil), "naruto0913.lottery.lottery.LotteryInfo")
}

func init() {
	proto.RegisterFile("lottery/lottery/lottery_info.proto", fileDescriptor_a8cadedeb4e07f20)
}

var fileDescriptor_a8cadedeb4e07f20 = []byte{
	// 185 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0xca, 0xc9, 0x2f, 0x29,
	0x49, 0x2d, 0xaa, 0xd4, 0x47, 0xa3, 0xe3, 0x33, 0xf3, 0xd2, 0xf2, 0xf5, 0x0a, 0x8a, 0xf2, 0x4b,
	0xf2, 0x85, 0xa4, 0xf2, 0x12, 0x8b, 0x4a, 0x4b, 0xf2, 0x0d, 0x2c, 0x0d, 0x8d, 0xf5, 0xa0, 0xd2,
	0x30, 0x5a, 0x29, 0x9e, 0x8b, 0xdb, 0x07, 0xc2, 0xf4, 0xcc, 0x4b, 0xcb, 0x17, 0x12, 0xe3, 0x62,
	0xcb, 0x4b, 0xad, 0x28, 0xf1, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x0e, 0x82, 0xf2, 0x84,
	0x64, 0xb8, 0x38, 0x41, 0x2c, 0xff, 0xa2, 0x94, 0xd4, 0x22, 0x09, 0x26, 0xb0, 0x14, 0x42, 0x40,
	0x48, 0x8a, 0x8b, 0xa3, 0x24, 0xbf, 0x24, 0x31, 0xc7, 0x31, 0xb7, 0x44, 0x82, 0x19, 0x2c, 0x09,
	0xe7, 0x3b, 0xb9, 0x9f, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c,
	0x13, 0x1e, 0xcb, 0x31, 0x5c, 0x78, 0x2c, 0xc7, 0x70, 0xe3, 0xb1, 0x1c, 0x43, 0x94, 0x6e, 0x7a,
	0x66, 0x49, 0x46, 0x69, 0x92, 0x5e, 0x72, 0x7e, 0xae, 0x3e, 0xc2, 0x85, 0x70, 0x8f, 0x54, 0xc0,
	0x59, 0x25, 0x95, 0x05, 0xa9, 0xc5, 0x49, 0x6c, 0x60, 0xcf, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff,
	0xff, 0xae, 0xd4, 0x11, 0x15, 0xf2, 0x00, 0x00, 0x00,
}

func (m *LotteryInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LotteryInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LotteryInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.TotalAmt != 0 {
		i = encodeVarintLotteryInfo(dAtA, i, uint64(m.TotalAmt))
		i--
		dAtA[i] = 0x18
	}
	if m.NextOrder != 0 {
		i = encodeVarintLotteryInfo(dAtA, i, uint64(m.NextOrder))
		i--
		dAtA[i] = 0x10
	}
	if m.NextId != 0 {
		i = encodeVarintLotteryInfo(dAtA, i, uint64(m.NextId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintLotteryInfo(dAtA []byte, offset int, v uint64) int {
	offset -= sovLotteryInfo(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *LotteryInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.NextId != 0 {
		n += 1 + sovLotteryInfo(uint64(m.NextId))
	}
	if m.NextOrder != 0 {
		n += 1 + sovLotteryInfo(uint64(m.NextOrder))
	}
	if m.TotalAmt != 0 {
		n += 1 + sovLotteryInfo(uint64(m.TotalAmt))
	}
	return n
}

func sovLotteryInfo(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozLotteryInfo(x uint64) (n int) {
	return sovLotteryInfo(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *LotteryInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLotteryInfo
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
			return fmt.Errorf("proto: LotteryInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LotteryInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NextId", wireType)
			}
			m.NextId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLotteryInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NextId |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NextOrder", wireType)
			}
			m.NextOrder = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLotteryInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NextOrder |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalAmt", wireType)
			}
			m.TotalAmt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLotteryInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TotalAmt |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipLotteryInfo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLotteryInfo
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
func skipLotteryInfo(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLotteryInfo
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
					return 0, ErrIntOverflowLotteryInfo
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
					return 0, ErrIntOverflowLotteryInfo
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
				return 0, ErrInvalidLengthLotteryInfo
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupLotteryInfo
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthLotteryInfo
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthLotteryInfo        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLotteryInfo          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupLotteryInfo = fmt.Errorf("proto: unexpected end of group")
)
