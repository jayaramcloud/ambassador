// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/api/v2/ratelimit/ratelimit.proto

package envoy_api_v2_ratelimit

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

// A RateLimitDescriptor is a list of hierarchical entries that are used by the service to
// determine the final rate limit key and overall allowed limit. Here are some examples of how
// they might be used for the domain "envoy".
//
// .. code-block:: cpp
//
//   ["authenticated": "false"], ["remote_address": "10.0.0.1"]
//
// What it does: Limits all unauthenticated traffic for the IP address 10.0.0.1. The
// configuration supplies a default limit for the *remote_address* key. If there is a desire to
// raise the limit for 10.0.0.1 or block it entirely it can be specified directly in the
// configuration.
//
// .. code-block:: cpp
//
//   ["authenticated": "false"], ["path": "/foo/bar"]
//
// What it does: Limits all unauthenticated traffic globally for a specific path (or prefix if
// configured that way in the service).
//
// .. code-block:: cpp
//
//   ["authenticated": "false"], ["path": "/foo/bar"], ["remote_address": "10.0.0.1"]
//
// What it does: Limits unauthenticated traffic to a specific path for a specific IP address.
// Like (1) we can raise/block specific IP addresses if we want with an override configuration.
//
// .. code-block:: cpp
//
//   ["authenticated": "true"], ["client_id": "foo"]
//
// What it does: Limits all traffic for an authenticated client "foo"
//
// .. code-block:: cpp
//
//   ["authenticated": "true"], ["client_id": "foo"], ["path": "/foo/bar"]
//
// What it does: Limits traffic to a specific path for an authenticated client "foo"
//
// The idea behind the API is that (1)/(2)/(3) and (4)/(5) can be sent in 1 request if desired.
// This enables building complex application scenarios with a generic backend.
type RateLimitDescriptor struct {
	// Descriptor entries.
	Entries              []*RateLimitDescriptor_Entry `protobuf:"bytes,1,rep,name=entries,proto3" json:"entries,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *RateLimitDescriptor) Reset()         { *m = RateLimitDescriptor{} }
func (m *RateLimitDescriptor) String() string { return proto.CompactTextString(m) }
func (*RateLimitDescriptor) ProtoMessage()    {}
func (*RateLimitDescriptor) Descriptor() ([]byte, []int) {
	return fileDescriptor_5684844e04543b8d, []int{0}
}
func (m *RateLimitDescriptor) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RateLimitDescriptor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RateLimitDescriptor.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RateLimitDescriptor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RateLimitDescriptor.Merge(m, src)
}
func (m *RateLimitDescriptor) XXX_Size() int {
	return m.Size()
}
func (m *RateLimitDescriptor) XXX_DiscardUnknown() {
	xxx_messageInfo_RateLimitDescriptor.DiscardUnknown(m)
}

var xxx_messageInfo_RateLimitDescriptor proto.InternalMessageInfo

func (m *RateLimitDescriptor) GetEntries() []*RateLimitDescriptor_Entry {
	if m != nil {
		return m.Entries
	}
	return nil
}

type RateLimitDescriptor_Entry struct {
	// Descriptor key.
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// Descriptor value.
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RateLimitDescriptor_Entry) Reset()         { *m = RateLimitDescriptor_Entry{} }
func (m *RateLimitDescriptor_Entry) String() string { return proto.CompactTextString(m) }
func (*RateLimitDescriptor_Entry) ProtoMessage()    {}
func (*RateLimitDescriptor_Entry) Descriptor() ([]byte, []int) {
	return fileDescriptor_5684844e04543b8d, []int{0, 0}
}
func (m *RateLimitDescriptor_Entry) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RateLimitDescriptor_Entry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RateLimitDescriptor_Entry.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RateLimitDescriptor_Entry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RateLimitDescriptor_Entry.Merge(m, src)
}
func (m *RateLimitDescriptor_Entry) XXX_Size() int {
	return m.Size()
}
func (m *RateLimitDescriptor_Entry) XXX_DiscardUnknown() {
	xxx_messageInfo_RateLimitDescriptor_Entry.DiscardUnknown(m)
}

var xxx_messageInfo_RateLimitDescriptor_Entry proto.InternalMessageInfo

func (m *RateLimitDescriptor_Entry) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *RateLimitDescriptor_Entry) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*RateLimitDescriptor)(nil), "envoy.api.v2.ratelimit.RateLimitDescriptor")
	proto.RegisterType((*RateLimitDescriptor_Entry)(nil), "envoy.api.v2.ratelimit.RateLimitDescriptor.Entry")
}

func init() {
	proto.RegisterFile("envoy/api/v2/ratelimit/ratelimit.proto", fileDescriptor_5684844e04543b8d)
}

var fileDescriptor_5684844e04543b8d = []byte{
	// 296 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0x4f, 0x4a, 0xc3, 0x40,
	0x14, 0xc6, 0x99, 0xd4, 0xd8, 0x3a, 0x82, 0x8b, 0x08, 0x1a, 0x03, 0x86, 0x20, 0xa5, 0x74, 0x35,
	0x83, 0xe9, 0x09, 0x1c, 0x74, 0xe7, 0xa2, 0x04, 0x3c, 0xc0, 0xd8, 0x0e, 0x32, 0x98, 0xcc, 0x0c,
	0x93, 0xd7, 0xa1, 0x39, 0x85, 0xe0, 0xca, 0x5b, 0x78, 0x0d, 0x97, 0x1e, 0x41, 0x72, 0x04, 0x97,
	0x5d, 0x88, 0x24, 0x21, 0xfe, 0x81, 0x76, 0xf7, 0xf1, 0x7e, 0xf3, 0xde, 0x7c, 0xdf, 0x87, 0x27,
	0x42, 0x39, 0x5d, 0x51, 0x6e, 0x24, 0x75, 0x29, 0xb5, 0x1c, 0x44, 0x2e, 0x0b, 0x09, 0xbf, 0x8a,
	0x18, 0xab, 0x41, 0x07, 0x27, 0xed, 0x3b, 0xc2, 0x8d, 0x24, 0x2e, 0x25, 0x3f, 0x34, 0x8a, 0x57,
	0x4b, 0xc3, 0x29, 0x57, 0x4a, 0x03, 0x07, 0xa9, 0x55, 0x49, 0x0b, 0xf9, 0xd0, 0xe0, 0x6e, 0x2f,
	0x3a, 0x75, 0x3c, 0x97, 0x4b, 0x0e, 0x82, 0xf6, 0xa2, 0x03, 0x17, 0xaf, 0x08, 0x1f, 0x67, 0x1c,
	0xc4, 0x6d, 0x73, 0xe6, 0x5a, 0x94, 0x0b, 0x2b, 0x0d, 0x68, 0x1b, 0xdc, 0xe1, 0xa1, 0x50, 0x60,
	0xa5, 0x28, 0x43, 0x94, 0x0c, 0xa6, 0x87, 0xe9, 0x25, 0xd9, 0xfe, 0x35, 0xd9, 0xb2, 0x4d, 0x6e,
	0x14, 0xd8, 0x8a, 0x8d, 0x36, 0xcc, 0x7f, 0x46, 0xde, 0x08, 0x65, 0xfd, 0xad, 0xe8, 0x0a, 0xfb,
	0x2d, 0x0b, 0xce, 0xf0, 0xe0, 0x51, 0x54, 0x21, 0x4a, 0xd0, 0xf4, 0x80, 0x0d, 0x37, 0x6c, 0xcf,
	0x7a, 0x09, 0xca, 0x9a, 0x59, 0x70, 0x8e, 0x7d, 0xc7, 0xf3, 0x95, 0x08, 0xbd, 0xff, 0xb0, 0x9b,
	0x32, 0xf3, 0x56, 0xc7, 0xe8, 0xbd, 0x8e, 0xd1, 0x47, 0x1d, 0xa3, 0xcf, 0x97, 0xaf, 0x27, 0x7f,
	0x12, 0x8c, 0x3b, 0x6f, 0x62, 0x0d, 0x42, 0x95, 0x4d, 0x7c, 0xb2, 0xd0, 0x45, 0xa1, 0xd5, 0x1f,
	0x9b, 0x6e, 0x86, 0xc7, 0x52, 0x77, 0x21, 0x8c, 0xd5, 0xeb, 0x6a, 0x47, 0x1e, 0x76, 0x94, 0xf5,
	0x72, 0xde, 0x34, 0x34, 0x47, 0xf7, 0xfb, 0x6d, 0x55, 0xb3, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x2c, 0xb8, 0x15, 0x86, 0xa5, 0x01, 0x00, 0x00,
}

func (m *RateLimitDescriptor) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RateLimitDescriptor) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RateLimitDescriptor) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Entries) > 0 {
		for iNdEx := len(m.Entries) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Entries[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintRatelimit(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *RateLimitDescriptor_Entry) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RateLimitDescriptor_Entry) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RateLimitDescriptor_Entry) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintRatelimit(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i = encodeVarintRatelimit(dAtA, i, uint64(len(m.Key)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintRatelimit(dAtA []byte, offset int, v uint64) int {
	offset -= sovRatelimit(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *RateLimitDescriptor) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Entries) > 0 {
		for _, e := range m.Entries {
			l = e.Size()
			n += 1 + l + sovRatelimit(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *RateLimitDescriptor_Entry) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovRatelimit(uint64(l))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovRatelimit(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovRatelimit(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozRatelimit(x uint64) (n int) {
	return sovRatelimit(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RateLimitDescriptor) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRatelimit
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
			return fmt.Errorf("proto: RateLimitDescriptor: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RateLimitDescriptor: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Entries", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRatelimit
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
				return ErrInvalidLengthRatelimit
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRatelimit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Entries = append(m.Entries, &RateLimitDescriptor_Entry{})
			if err := m.Entries[len(m.Entries)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRatelimit(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRatelimit
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthRatelimit
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
func (m *RateLimitDescriptor_Entry) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRatelimit
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
			return fmt.Errorf("proto: Entry: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Entry: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRatelimit
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
				return ErrInvalidLengthRatelimit
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRatelimit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRatelimit
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
				return ErrInvalidLengthRatelimit
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRatelimit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRatelimit(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRatelimit
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthRatelimit
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
func skipRatelimit(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRatelimit
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
					return 0, ErrIntOverflowRatelimit
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
					return 0, ErrIntOverflowRatelimit
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
				return 0, ErrInvalidLengthRatelimit
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthRatelimit
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowRatelimit
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
				next, err := skipRatelimit(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthRatelimit
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
	ErrInvalidLengthRatelimit = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRatelimit   = fmt.Errorf("proto: integer overflow")
)
