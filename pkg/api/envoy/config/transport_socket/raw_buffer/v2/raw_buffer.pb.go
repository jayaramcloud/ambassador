// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/config/transport_socket/raw_buffer/v2/raw_buffer.proto

package envoy_config_transport_socket_raw_buffer_v2

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

// Configuration for raw buffer transport socket.
type RawBuffer struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RawBuffer) Reset()         { *m = RawBuffer{} }
func (m *RawBuffer) String() string { return proto.CompactTextString(m) }
func (*RawBuffer) ProtoMessage()    {}
func (*RawBuffer) Descriptor() ([]byte, []int) {
	return fileDescriptor_7d9b219af9f95c92, []int{0}
}
func (m *RawBuffer) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RawBuffer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RawBuffer.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RawBuffer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RawBuffer.Merge(m, src)
}
func (m *RawBuffer) XXX_Size() int {
	return m.Size()
}
func (m *RawBuffer) XXX_DiscardUnknown() {
	xxx_messageInfo_RawBuffer.DiscardUnknown(m)
}

var xxx_messageInfo_RawBuffer proto.InternalMessageInfo

func init() {
	proto.RegisterType((*RawBuffer)(nil), "envoy.config.transport_socket.raw_buffer.v2.RawBuffer")
}

func init() {
	proto.RegisterFile("envoy/config/transport_socket/raw_buffer/v2/raw_buffer.proto", fileDescriptor_7d9b219af9f95c92)
}

var fileDescriptor_7d9b219af9f95c92 = []byte{
	// 202 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0x49, 0xcd, 0x2b, 0xcb,
	0xaf, 0xd4, 0x4f, 0xce, 0xcf, 0x4b, 0xcb, 0x4c, 0xd7, 0x2f, 0x29, 0x4a, 0xcc, 0x2b, 0x2e, 0xc8,
	0x2f, 0x2a, 0x89, 0x2f, 0xce, 0x4f, 0xce, 0x4e, 0x2d, 0xd1, 0x2f, 0x4a, 0x2c, 0x8f, 0x4f, 0x2a,
	0x4d, 0x4b, 0x4b, 0x2d, 0xd2, 0x2f, 0x33, 0x42, 0xe2, 0xe9, 0x15, 0x14, 0xe5, 0x97, 0xe4, 0x0b,
	0x69, 0x83, 0x75, 0xeb, 0x41, 0x74, 0xeb, 0xa1, 0xeb, 0xd6, 0x43, 0x52, 0x5f, 0x66, 0x24, 0x25,
	0x57, 0x9a, 0x52, 0x90, 0xa8, 0x9f, 0x98, 0x97, 0x97, 0x5f, 0x92, 0x58, 0x92, 0x99, 0x9f, 0x57,
	0xac, 0x9f, 0x9b, 0x99, 0x5e, 0x94, 0x58, 0x92, 0x0a, 0x31, 0x4c, 0x89, 0x9b, 0x8b, 0x33, 0x28,
	0xb1, 0xdc, 0x09, 0xac, 0xde, 0x69, 0x22, 0xe3, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31,
	0x3e, 0x78, 0x24, 0xc7, 0xf8, 0x69, 0xc6, 0xbf, 0x7e, 0x56, 0x23, 0x21, 0x03, 0x88, 0x75, 0xa9,
	0x15, 0x25, 0xa9, 0x79, 0xc5, 0x20, 0x13, 0x30, 0xac, 0x2c, 0x46, 0xb1, 0xd3, 0x98, 0xcb, 0x32,
	0x33, 0x5f, 0x0f, 0xac, 0xa9, 0xa0, 0x28, 0xbf, 0xa2, 0x52, 0x8f, 0x04, 0xe7, 0x3a, 0xf1, 0xc1,
	0x1d, 0x13, 0x00, 0x72, 0x5e, 0x00, 0x63, 0x12, 0x1b, 0xd8, 0x9d, 0xc6, 0x80, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xdc, 0xb9, 0x69, 0x80, 0x34, 0x01, 0x00, 0x00,
}

func (m *RawBuffer) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RawBuffer) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RawBuffer) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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

func encodeVarintRawBuffer(dAtA []byte, offset int, v uint64) int {
	offset -= sovRawBuffer(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *RawBuffer) Size() (n int) {
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

func sovRawBuffer(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozRawBuffer(x uint64) (n int) {
	return sovRawBuffer(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RawBuffer) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRawBuffer
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
			return fmt.Errorf("proto: RawBuffer: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RawBuffer: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipRawBuffer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRawBuffer
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthRawBuffer
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
func skipRawBuffer(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRawBuffer
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
					return 0, ErrIntOverflowRawBuffer
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
					return 0, ErrIntOverflowRawBuffer
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
				return 0, ErrInvalidLengthRawBuffer
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthRawBuffer
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowRawBuffer
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
				next, err := skipRawBuffer(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthRawBuffer
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
	ErrInvalidLengthRawBuffer = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRawBuffer   = fmt.Errorf("proto: integer overflow")
)
