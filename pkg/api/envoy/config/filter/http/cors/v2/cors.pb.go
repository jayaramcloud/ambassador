// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/config/filter/http/cors/v2/cors.proto

package envoy_config_filter_http_cors_v2

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
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

// Cors filter config.
type Cors struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Cors) Reset()         { *m = Cors{} }
func (m *Cors) String() string { return proto.CompactTextString(m) }
func (*Cors) ProtoMessage()    {}
func (*Cors) Descriptor() ([]byte, []int) {
	return fileDescriptor_8ba80e325986c17c, []int{0}
}
func (m *Cors) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Cors) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Cors.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Cors) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Cors.Merge(m, src)
}
func (m *Cors) XXX_Size() int {
	return m.Size()
}
func (m *Cors) XXX_DiscardUnknown() {
	xxx_messageInfo_Cors.DiscardUnknown(m)
}

var xxx_messageInfo_Cors proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Cors)(nil), "envoy.config.filter.http.cors.v2.Cors")
}

func init() {
	proto.RegisterFile("envoy/config/filter/http/cors/v2/cors.proto", fileDescriptor_8ba80e325986c17c)
}

var fileDescriptor_8ba80e325986c17c = []byte{
	// 189 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x4e, 0xcd, 0x2b, 0xcb,
	0xaf, 0xd4, 0x4f, 0xce, 0xcf, 0x4b, 0xcb, 0x4c, 0xd7, 0x4f, 0xcb, 0xcc, 0x29, 0x49, 0x2d, 0xd2,
	0xcf, 0x28, 0x29, 0x29, 0xd0, 0x4f, 0xce, 0x2f, 0x2a, 0xd6, 0x2f, 0x33, 0x02, 0xd3, 0x7a, 0x05,
	0x45, 0xf9, 0x25, 0xf9, 0x42, 0x0a, 0x60, 0xc5, 0x7a, 0x10, 0xc5, 0x7a, 0x10, 0xc5, 0x7a, 0x20,
	0xc5, 0x7a, 0x60, 0x45, 0x65, 0x46, 0x52, 0x72, 0xa5, 0x29, 0x05, 0x89, 0xfa, 0x89, 0x79, 0x79,
	0xf9, 0x25, 0x89, 0x25, 0x99, 0xf9, 0x79, 0xc5, 0xfa, 0xb9, 0x99, 0xe9, 0x45, 0x89, 0x25, 0xa9,
	0x10, 0x13, 0x94, 0xd8, 0xb8, 0x58, 0x9c, 0xf3, 0x8b, 0x8a, 0x9d, 0xca, 0x4e, 0x3c, 0x92, 0x63,
	0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x4f, 0x33, 0xfe, 0xf5, 0xb3, 0xaa, 0x0b,
	0xa9, 0x42, 0x4c, 0x4f, 0xad, 0x28, 0x49, 0xcd, 0x2b, 0x06, 0xe9, 0x85, 0xda, 0x50, 0x8c, 0x6c,
	0x85, 0x31, 0x97, 0x5e, 0x66, 0xbe, 0x1e, 0x58, 0x65, 0x41, 0x51, 0x7e, 0x45, 0xa5, 0x1e, 0x21,
	0x27, 0x39, 0x71, 0x82, 0x2c, 0x0c, 0x00, 0xd9, 0x1e, 0xc0, 0x98, 0xc4, 0x06, 0x76, 0x86, 0x31,
	0x20, 0x00, 0x00, 0xff, 0xff, 0xb2, 0xaf, 0x32, 0xc9, 0xf7, 0x00, 0x00, 0x00,
}

func (m *Cors) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Cors) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Cors) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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

func encodeVarintCors(dAtA []byte, offset int, v uint64) int {
	offset -= sovCors(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Cors) Size() (n int) {
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

func sovCors(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCors(x uint64) (n int) {
	return sovCors(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Cors) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCors
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
			return fmt.Errorf("proto: Cors: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Cors: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipCors(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCors
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCors
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
func skipCors(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCors
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
					return 0, ErrIntOverflowCors
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
					return 0, ErrIntOverflowCors
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
				return 0, ErrInvalidLengthCors
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthCors
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowCors
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
				next, err := skipCors(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthCors
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
	ErrInvalidLengthCors = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCors   = fmt.Errorf("proto: integer overflow")
)
