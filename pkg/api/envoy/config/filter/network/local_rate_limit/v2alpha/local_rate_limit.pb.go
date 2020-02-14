// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/config/filter/network/local_rate_limit/v2alpha/local_rate_limit.proto

package envoy_config_filter_network_local_rate_limit_v2alpha

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	core "github.com/datawire/ambassador/pkg/api/envoy/api/v2/core"
	_type "github.com/datawire/ambassador/pkg/api/envoy/type"
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

type LocalRateLimit struct {
	// The prefix to use when emitting :ref:`statistics
	// <config_network_filters_local_rate_limit_stats>`.
	StatPrefix string `protobuf:"bytes,1,opt,name=stat_prefix,json=statPrefix,proto3" json:"stat_prefix,omitempty"`
	// The token bucket configuration to use for rate limiting connections that are processed by the
	// filter's filter chain. Each incoming connection processed by the filter consumes a single
	// token. If the token is available, the connection will be allowed. If no tokens are available,
	// the connection will be immediately closed.
	//
	// .. note::
	//   In the current implementation each filter and filter chain has an independent rate limit.
	//
	// .. note::
	//   In the current implementation the token bucket's :ref:`fill_interval
	//   <envoy_api_field_type.TokenBucket.fill_interval>` must be >= 50ms to avoid too aggressive
	//   refills.
	TokenBucket *_type.TokenBucket `protobuf:"bytes,2,opt,name=token_bucket,json=tokenBucket,proto3" json:"token_bucket,omitempty"`
	// Runtime flag that controls whether the filter is enabled or not. If not specified, defaults
	// to enabled.
	RuntimeEnabled       *core.RuntimeFeatureFlag `protobuf:"bytes,3,opt,name=runtime_enabled,json=runtimeEnabled,proto3" json:"runtime_enabled,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *LocalRateLimit) Reset()         { *m = LocalRateLimit{} }
func (m *LocalRateLimit) String() string { return proto.CompactTextString(m) }
func (*LocalRateLimit) ProtoMessage()    {}
func (*LocalRateLimit) Descriptor() ([]byte, []int) {
	return fileDescriptor_984b0e5e5a865836, []int{0}
}
func (m *LocalRateLimit) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LocalRateLimit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LocalRateLimit.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LocalRateLimit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LocalRateLimit.Merge(m, src)
}
func (m *LocalRateLimit) XXX_Size() int {
	return m.Size()
}
func (m *LocalRateLimit) XXX_DiscardUnknown() {
	xxx_messageInfo_LocalRateLimit.DiscardUnknown(m)
}

var xxx_messageInfo_LocalRateLimit proto.InternalMessageInfo

func (m *LocalRateLimit) GetStatPrefix() string {
	if m != nil {
		return m.StatPrefix
	}
	return ""
}

func (m *LocalRateLimit) GetTokenBucket() *_type.TokenBucket {
	if m != nil {
		return m.TokenBucket
	}
	return nil
}

func (m *LocalRateLimit) GetRuntimeEnabled() *core.RuntimeFeatureFlag {
	if m != nil {
		return m.RuntimeEnabled
	}
	return nil
}

func init() {
	proto.RegisterType((*LocalRateLimit)(nil), "envoy.config.filter.network.local_rate_limit.v2alpha.LocalRateLimit")
}

func init() {
	proto.RegisterFile("envoy/config/filter/network/local_rate_limit/v2alpha/local_rate_limit.proto", fileDescriptor_984b0e5e5a865836)
}

var fileDescriptor_984b0e5e5a865836 = []byte{
	// 391 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x91, 0xcf, 0xaa, 0xd3, 0x40,
	0x14, 0xc6, 0x99, 0xea, 0xf5, 0x4f, 0x2a, 0x57, 0x89, 0x8b, 0x5b, 0x8a, 0x96, 0x22, 0x08, 0x5d,
	0xcd, 0x40, 0xab, 0x2f, 0x30, 0xe8, 0xdd, 0x78, 0x91, 0x12, 0xdc, 0x87, 0x93, 0xf4, 0xb4, 0x0e,
	0x9d, 0xce, 0x0c, 0x93, 0x93, 0xd8, 0xbe, 0x82, 0x1b, 0xb7, 0xae, 0x7d, 0x1a, 0x97, 0xf7, 0x11,
	0xa4, 0x8f, 0xe0, 0xf2, 0x2e, 0x44, 0x26, 0x33, 0xa2, 0x52, 0x57, 0xee, 0x42, 0x7e, 0x39, 0xbf,
	0xf3, 0xe5, 0x3b, 0xd9, 0x1b, 0x34, 0x9d, 0x3d, 0x88, 0xda, 0x9a, 0xb5, 0xda, 0x88, 0xb5, 0xd2,
	0x84, 0x5e, 0x18, 0xa4, 0x0f, 0xd6, 0x6f, 0x85, 0xb6, 0x35, 0xe8, 0xd2, 0x03, 0x61, 0xa9, 0xd5,
	0x4e, 0x91, 0xe8, 0xe6, 0xa0, 0xdd, 0x7b, 0x38, 0x01, 0xdc, 0x79, 0x4b, 0x36, 0x7f, 0xd1, 0xcb,
	0x78, 0x94, 0xf1, 0x28, 0xe3, 0x49, 0xc6, 0x4f, 0x66, 0x92, 0x6c, 0xfc, 0x24, 0x46, 0x00, 0xa7,
	0x44, 0x37, 0x17, 0xb5, 0xf5, 0x28, 0x2a, 0x68, 0x30, 0x3a, 0xc7, 0x4f, 0x23, 0xa5, 0x83, 0x43,
	0x41, 0x76, 0x8b, 0xa6, 0xac, 0xda, 0x7a, 0x8b, 0x69, 0xe5, 0x78, 0xd2, 0xae, 0x1c, 0x08, 0x30,
	0xc6, 0x12, 0x90, 0xb2, 0xa6, 0x11, 0x3b, 0xb5, 0x09, 0x4b, 0x12, 0xbf, 0xe8, 0x40, 0xab, 0x15,
	0x10, 0x8a, 0x5f, 0x0f, 0x11, 0x3c, 0xbb, 0x66, 0xd9, 0xf9, 0x55, 0x88, 0x54, 0x00, 0xe1, 0x55,
	0x08, 0x94, 0xcf, 0xb2, 0x61, 0x43, 0x40, 0xa5, 0xf3, 0xb8, 0x56, 0xfb, 0x11, 0x9b, 0xb2, 0xd9,
	0x7d, 0x79, 0xf7, 0x46, 0xde, 0xf6, 0x83, 0x29, 0x2b, 0xb2, 0xc0, 0x96, 0x3d, 0xca, 0x5f, 0x65,
	0x0f, 0xfe, 0xcc, 0x32, 0x1a, 0x4c, 0xd9, 0x6c, 0x38, 0xbf, 0xe0, 0xf1, 0xff, 0x43, 0x56, 0xfe,
	0x2e, 0x70, 0xd9, 0x63, 0x79, 0xef, 0x46, 0x9e, 0x7d, 0x64, 0x83, 0x47, 0xac, 0x18, 0xd2, 0xef,
	0xd7, 0xf9, 0xdb, 0xec, 0xa1, 0x6f, 0x0d, 0xa9, 0x1d, 0x96, 0x68, 0xa0, 0xd2, 0xb8, 0x1a, 0xdd,
	0xea, 0x45, 0xcf, 0x93, 0x08, 0x9c, 0xe2, 0xdd, 0x9c, 0x87, 0x4a, 0x78, 0x11, 0xbf, 0xbc, 0x44,
	0xa0, 0xd6, 0xe3, 0xa5, 0x86, 0x4d, 0x71, 0x9e, 0xa6, 0x5f, 0xc7, 0x61, 0xf9, 0x85, 0x7d, 0x3d,
	0x4e, 0xd8, 0xf5, 0x71, 0xc2, 0xbe, 0x1d, 0x27, 0xec, 0xfb, 0xe7, 0x1f, 0x9f, 0xce, 0x5e, 0xe6,
	0x8b, 0xa8, 0xc2, 0x3d, 0xa1, 0x69, 0x42, 0x41, 0xe9, 0x2e, 0xcd, 0x3f, 0x0e, 0x93, 0xee, 0xb2,
	0xc8, 0xa4, 0xb2, 0x31, 0x82, 0xf3, 0x76, 0x7f, 0xe0, 0xff, 0x73, 0x56, 0xf9, 0xf8, 0xef, 0x76,
	0x97, 0xa1, 0xf5, 0x25, 0xab, 0xee, 0xf4, 0xf5, 0x2f, 0x7e, 0x06, 0x00, 0x00, 0xff, 0xff, 0xd5,
	0x54, 0xfb, 0xba, 0x79, 0x02, 0x00, 0x00,
}

func (m *LocalRateLimit) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LocalRateLimit) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LocalRateLimit) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.RuntimeEnabled != nil {
		{
			size, err := m.RuntimeEnabled.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintLocalRateLimit(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.TokenBucket != nil {
		{
			size, err := m.TokenBucket.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintLocalRateLimit(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.StatPrefix) > 0 {
		i -= len(m.StatPrefix)
		copy(dAtA[i:], m.StatPrefix)
		i = encodeVarintLocalRateLimit(dAtA, i, uint64(len(m.StatPrefix)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintLocalRateLimit(dAtA []byte, offset int, v uint64) int {
	offset -= sovLocalRateLimit(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *LocalRateLimit) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.StatPrefix)
	if l > 0 {
		n += 1 + l + sovLocalRateLimit(uint64(l))
	}
	if m.TokenBucket != nil {
		l = m.TokenBucket.Size()
		n += 1 + l + sovLocalRateLimit(uint64(l))
	}
	if m.RuntimeEnabled != nil {
		l = m.RuntimeEnabled.Size()
		n += 1 + l + sovLocalRateLimit(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovLocalRateLimit(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozLocalRateLimit(x uint64) (n int) {
	return sovLocalRateLimit(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *LocalRateLimit) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLocalRateLimit
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
			return fmt.Errorf("proto: LocalRateLimit: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LocalRateLimit: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StatPrefix", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLocalRateLimit
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
				return ErrInvalidLengthLocalRateLimit
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLocalRateLimit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StatPrefix = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenBucket", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLocalRateLimit
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
				return ErrInvalidLengthLocalRateLimit
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLocalRateLimit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.TokenBucket == nil {
				m.TokenBucket = &_type.TokenBucket{}
			}
			if err := m.TokenBucket.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RuntimeEnabled", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLocalRateLimit
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
				return ErrInvalidLengthLocalRateLimit
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLocalRateLimit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.RuntimeEnabled == nil {
				m.RuntimeEnabled = &core.RuntimeFeatureFlag{}
			}
			if err := m.RuntimeEnabled.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLocalRateLimit(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLocalRateLimit
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthLocalRateLimit
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
func skipLocalRateLimit(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLocalRateLimit
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
					return 0, ErrIntOverflowLocalRateLimit
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
					return 0, ErrIntOverflowLocalRateLimit
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
				return 0, ErrInvalidLengthLocalRateLimit
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthLocalRateLimit
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowLocalRateLimit
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
				next, err := skipLocalRateLimit(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthLocalRateLimit
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
	ErrInvalidLengthLocalRateLimit = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLocalRateLimit   = fmt.Errorf("proto: integer overflow")
)
