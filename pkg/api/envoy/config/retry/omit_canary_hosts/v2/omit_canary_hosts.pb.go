// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/config/retry/omit_canary_hosts/v2/omit_canary_hosts.proto

package envoy_config_retry_omit_canary_hosts_v2

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

type OmitCanaryHostsPredicate struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OmitCanaryHostsPredicate) Reset()         { *m = OmitCanaryHostsPredicate{} }
func (m *OmitCanaryHostsPredicate) String() string { return proto.CompactTextString(m) }
func (*OmitCanaryHostsPredicate) ProtoMessage()    {}
func (*OmitCanaryHostsPredicate) Descriptor() ([]byte, []int) {
	return fileDescriptor_08501cb61947b43e, []int{0}
}
func (m *OmitCanaryHostsPredicate) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OmitCanaryHostsPredicate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OmitCanaryHostsPredicate.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OmitCanaryHostsPredicate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OmitCanaryHostsPredicate.Merge(m, src)
}
func (m *OmitCanaryHostsPredicate) XXX_Size() int {
	return m.Size()
}
func (m *OmitCanaryHostsPredicate) XXX_DiscardUnknown() {
	xxx_messageInfo_OmitCanaryHostsPredicate.DiscardUnknown(m)
}

var xxx_messageInfo_OmitCanaryHostsPredicate proto.InternalMessageInfo

func init() {
	proto.RegisterType((*OmitCanaryHostsPredicate)(nil), "envoy.config.retry.omit_canary_hosts.v2.OmitCanaryHostsPredicate")
}

func init() {
	proto.RegisterFile("envoy/config/retry/omit_canary_hosts/v2/omit_canary_hosts.proto", fileDescriptor_08501cb61947b43e)
}

var fileDescriptor_08501cb61947b43e = []byte{
	// 156 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0x4f, 0xcd, 0x2b, 0xcb,
	0xaf, 0xd4, 0x4f, 0xce, 0xcf, 0x4b, 0xcb, 0x4c, 0xd7, 0x2f, 0x4a, 0x2d, 0x29, 0xaa, 0xd4, 0xcf,
	0xcf, 0xcd, 0x2c, 0x89, 0x4f, 0x4e, 0xcc, 0x4b, 0x2c, 0xaa, 0x8c, 0xcf, 0xc8, 0x2f, 0x2e, 0x29,
	0xd6, 0x2f, 0x33, 0xc2, 0x14, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x52, 0x07, 0x1b, 0xa0,
	0x07, 0x31, 0x40, 0x0f, 0x6c, 0x80, 0x1e, 0xa6, 0xda, 0x32, 0x23, 0x25, 0x29, 0x2e, 0x09, 0xff,
	0xdc, 0xcc, 0x12, 0x67, 0xb0, 0xb0, 0x07, 0x48, 0x34, 0xa0, 0x28, 0x35, 0x25, 0x33, 0x39, 0xb1,
	0x24, 0xd5, 0x29, 0xfa, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63,
	0xe4, 0x32, 0xcd, 0xcc, 0xd7, 0x03, 0x9b, 0x5a, 0x50, 0x94, 0x5f, 0x51, 0xa9, 0x47, 0xa4, 0x05,
	0x4e, 0x22, 0x18, 0xc6, 0xe7, 0x97, 0xe4, 0x07, 0x30, 0x26, 0xb1, 0x81, 0x1d, 0x6a, 0x0c, 0x08,
	0x00, 0x00, 0xff, 0xff, 0x3f, 0x9c, 0x84, 0xee, 0xeb, 0x00, 0x00, 0x00,
}

func (m *OmitCanaryHostsPredicate) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OmitCanaryHostsPredicate) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OmitCanaryHostsPredicate) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	return len(dAtA) - i, nil
}

func encodeVarintOmitCanaryHosts(dAtA []byte, offset int, v uint64) int {
	offset -= sovOmitCanaryHosts(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *OmitCanaryHostsPredicate) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovOmitCanaryHosts(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozOmitCanaryHosts(x uint64) (n int) {
	return sovOmitCanaryHosts(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *OmitCanaryHostsPredicate) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOmitCanaryHosts
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
			return fmt.Errorf("proto: OmitCanaryHostsPredicate: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OmitCanaryHostsPredicate: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipOmitCanaryHosts(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthOmitCanaryHosts
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthOmitCanaryHosts
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
func skipOmitCanaryHosts(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowOmitCanaryHosts
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
					return 0, ErrIntOverflowOmitCanaryHosts
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowOmitCanaryHosts
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
				return 0, ErrInvalidLengthOmitCanaryHosts
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthOmitCanaryHosts
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowOmitCanaryHosts
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipOmitCanaryHosts(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthOmitCanaryHosts
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthOmitCanaryHosts = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowOmitCanaryHosts   = fmt.Errorf("proto: integer overflow")
)
