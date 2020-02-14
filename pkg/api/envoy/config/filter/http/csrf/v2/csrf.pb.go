// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/config/filter/http/csrf/v2/csrf.proto

package envoy_config_filter_http_csrf_v2

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	core "github.com/datawire/ambassador/pkg/api/envoy/api/v2/core"
	matcher "github.com/datawire/ambassador/pkg/api/envoy/type/matcher"
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

// CSRF filter config.
type CsrfPolicy struct {
	// Specifies the % of requests for which the CSRF filter is enabled.
	//
	// If :ref:`runtime_key <envoy_api_field_core.RuntimeFractionalPercent.runtime_key>` is specified,
	// Envoy will lookup the runtime key to get the percentage of requests to filter.
	//
	// .. note::
	//
	//   This field defaults to 100/:ref:`HUNDRED
	//   <envoy_api_enum_type.FractionalPercent.DenominatorType>`.
	FilterEnabled *core.RuntimeFractionalPercent `protobuf:"bytes,1,opt,name=filter_enabled,json=filterEnabled,proto3" json:"filter_enabled,omitempty"`
	// Specifies that CSRF policies will be evaluated and tracked, but not enforced.
	//
	// This is intended to be used when ``filter_enabled`` is off and will be ignored otherwise.
	//
	// If :ref:`runtime_key <envoy_api_field_core.RuntimeFractionalPercent.runtime_key>` is specified,
	// Envoy will lookup the runtime key to get the percentage of requests for which it will evaluate
	// and track the request's *Origin* and *Destination* to determine if it's valid, but will not
	// enforce any policies.
	ShadowEnabled *core.RuntimeFractionalPercent `protobuf:"bytes,2,opt,name=shadow_enabled,json=shadowEnabled,proto3" json:"shadow_enabled,omitempty"`
	// Specifies additional source origins that will be allowed in addition to
	// the destination origin.
	//
	// More information on how this can be configured via runtime can be found
	// :ref:`here <csrf-configuration>`.
	AdditionalOrigins    []*matcher.StringMatcher `protobuf:"bytes,3,rep,name=additional_origins,json=additionalOrigins,proto3" json:"additional_origins,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *CsrfPolicy) Reset()         { *m = CsrfPolicy{} }
func (m *CsrfPolicy) String() string { return proto.CompactTextString(m) }
func (*CsrfPolicy) ProtoMessage()    {}
func (*CsrfPolicy) Descriptor() ([]byte, []int) {
	return fileDescriptor_a9146cdf92353980, []int{0}
}
func (m *CsrfPolicy) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CsrfPolicy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CsrfPolicy.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CsrfPolicy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CsrfPolicy.Merge(m, src)
}
func (m *CsrfPolicy) XXX_Size() int {
	return m.Size()
}
func (m *CsrfPolicy) XXX_DiscardUnknown() {
	xxx_messageInfo_CsrfPolicy.DiscardUnknown(m)
}

var xxx_messageInfo_CsrfPolicy proto.InternalMessageInfo

func (m *CsrfPolicy) GetFilterEnabled() *core.RuntimeFractionalPercent {
	if m != nil {
		return m.FilterEnabled
	}
	return nil
}

func (m *CsrfPolicy) GetShadowEnabled() *core.RuntimeFractionalPercent {
	if m != nil {
		return m.ShadowEnabled
	}
	return nil
}

func (m *CsrfPolicy) GetAdditionalOrigins() []*matcher.StringMatcher {
	if m != nil {
		return m.AdditionalOrigins
	}
	return nil
}

func init() {
	proto.RegisterType((*CsrfPolicy)(nil), "envoy.config.filter.http.csrf.v2.CsrfPolicy")
}

func init() {
	proto.RegisterFile("envoy/config/filter/http/csrf/v2/csrf.proto", fileDescriptor_a9146cdf92353980)
}

var fileDescriptor_a9146cdf92353980 = []byte{
	// 374 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0xc1, 0x4a, 0xeb, 0x40,
	0x14, 0x86, 0x99, 0x96, 0x5e, 0xee, 0x4d, 0x69, 0xb9, 0x37, 0x9b, 0x5b, 0xca, 0x25, 0xb7, 0x0a,
	0x62, 0xa1, 0x30, 0x03, 0xe9, 0x1b, 0x44, 0x74, 0x27, 0x86, 0xb8, 0x15, 0xca, 0x69, 0x32, 0x49,
	0x07, 0xd2, 0x99, 0x30, 0x33, 0x8d, 0xcd, 0x2b, 0xb8, 0xd1, 0xa5, 0x8f, 0xe4, 0xd2, 0x47, 0x90,
	0x3e, 0x82, 0x4b, 0x17, 0x22, 0xc9, 0x49, 0xd5, 0x9d, 0xb8, 0x4a, 0xc2, 0xf9, 0xff, 0xef, 0x9c,
	0xfc, 0xbf, 0x33, 0xe3, 0xb2, 0x54, 0x15, 0x8b, 0x95, 0x4c, 0x45, 0xc6, 0x52, 0x91, 0x5b, 0xae,
	0xd9, 0xca, 0xda, 0x82, 0xc5, 0x46, 0xa7, 0xac, 0xf4, 0x9b, 0x27, 0x2d, 0xb4, 0xb2, 0xca, 0x9d,
	0x34, 0x62, 0x8a, 0x62, 0x8a, 0x62, 0x5a, 0x8b, 0x69, 0x23, 0x2a, 0xfd, 0xf1, 0x3f, 0xc4, 0x41,
	0x21, 0x1a, 0xab, 0xd2, 0x9c, 0x2d, 0xc1, 0x70, 0xf4, 0x8f, 0xff, 0xe3, 0xd4, 0x56, 0x05, 0x67,
	0x6b, 0xb0, 0xf1, 0x8a, 0x6b, 0x66, 0xac, 0x16, 0x32, 0x6b, 0x05, 0xde, 0x26, 0x29, 0x80, 0x81,
	0x94, 0xca, 0x82, 0x15, 0x4a, 0x1a, 0xb6, 0x16, 0x99, 0x06, 0xbb, 0x07, 0xfc, 0x2d, 0x21, 0x17,
	0x09, 0x58, 0xce, 0xf6, 0x2f, 0x38, 0x38, 0xbc, 0xeb, 0x38, 0xce, 0x89, 0xd1, 0x69, 0xa8, 0x72,
	0x11, 0x57, 0xee, 0x95, 0x33, 0xc4, 0xeb, 0x16, 0x5c, 0xc2, 0x32, 0xe7, 0xc9, 0x88, 0x4c, 0xc8,
	0xb4, 0xef, 0xcf, 0x28, 0xfe, 0x01, 0x14, 0x82, 0x96, 0x3e, 0xad, 0xef, 0xa3, 0xd1, 0x46, 0x5a,
	0xb1, 0xe6, 0x67, 0x1a, 0xe2, 0x7a, 0x25, 0xe4, 0x21, 0xd7, 0x31, 0x97, 0x36, 0xf8, 0xf9, 0x12,
	0xf4, 0x6e, 0x48, 0xe7, 0x37, 0x89, 0x06, 0x08, 0x3b, 0x45, 0x96, 0x1b, 0x39, 0x43, 0xb3, 0x82,
	0x44, 0x5d, 0xbf, 0xd3, 0x3b, 0xdf, 0xa6, 0x47, 0x03, 0x44, 0xec, 0x99, 0xa1, 0xe3, 0x42, 0x92,
	0x08, 0xd4, 0x2c, 0x94, 0x16, 0x99, 0x90, 0x66, 0xd4, 0x9d, 0x74, 0xa7, 0x7d, 0xff, 0xa0, 0xe5,
	0xd6, 0xb9, 0xd1, 0x36, 0x37, 0x7a, 0xd9, 0xe4, 0x76, 0x8e, 0x5f, 0xd1, 0x9f, 0x0f, 0xf3, 0x05,
	0x7a, 0x83, 0xf2, 0x61, 0xe7, 0x91, 0xc7, 0x9d, 0x47, 0x9e, 0x76, 0x1e, 0x79, 0xbe, 0x7f, 0xbd,
	0xed, 0x1d, 0xbb, 0x47, 0x08, 0xe2, 0x5b, 0xcb, 0xa5, 0xa9, 0xf3, 0x6d, 0x4b, 0x34, 0x9f, 0x5b,
	0x9c, 0x3b, 0x54, 0x28, 0x5c, 0x59, 0x68, 0xb5, 0xad, 0xe8, 0x57, 0xad, 0x07, 0xbf, 0x9a, 0xe8,
	0xeb, 0x22, 0x42, 0xb2, 0xfc, 0xd1, 0x34, 0x32, 0x7f, 0x0b, 0x00, 0x00, 0xff, 0xff, 0xd4, 0xb4,
	0xaa, 0x78, 0x5a, 0x02, 0x00, 0x00,
}

func (m *CsrfPolicy) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CsrfPolicy) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CsrfPolicy) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.AdditionalOrigins) > 0 {
		for iNdEx := len(m.AdditionalOrigins) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.AdditionalOrigins[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintCsrf(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.ShadowEnabled != nil {
		{
			size, err := m.ShadowEnabled.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCsrf(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.FilterEnabled != nil {
		{
			size, err := m.FilterEnabled.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCsrf(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintCsrf(dAtA []byte, offset int, v uint64) int {
	offset -= sovCsrf(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *CsrfPolicy) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.FilterEnabled != nil {
		l = m.FilterEnabled.Size()
		n += 1 + l + sovCsrf(uint64(l))
	}
	if m.ShadowEnabled != nil {
		l = m.ShadowEnabled.Size()
		n += 1 + l + sovCsrf(uint64(l))
	}
	if len(m.AdditionalOrigins) > 0 {
		for _, e := range m.AdditionalOrigins {
			l = e.Size()
			n += 1 + l + sovCsrf(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovCsrf(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCsrf(x uint64) (n int) {
	return sovCsrf(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CsrfPolicy) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCsrf
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
			return fmt.Errorf("proto: CsrfPolicy: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CsrfPolicy: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FilterEnabled", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCsrf
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
				return ErrInvalidLengthCsrf
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCsrf
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.FilterEnabled == nil {
				m.FilterEnabled = &core.RuntimeFractionalPercent{}
			}
			if err := m.FilterEnabled.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ShadowEnabled", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCsrf
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
				return ErrInvalidLengthCsrf
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCsrf
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ShadowEnabled == nil {
				m.ShadowEnabled = &core.RuntimeFractionalPercent{}
			}
			if err := m.ShadowEnabled.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AdditionalOrigins", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCsrf
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
				return ErrInvalidLengthCsrf
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCsrf
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AdditionalOrigins = append(m.AdditionalOrigins, &matcher.StringMatcher{})
			if err := m.AdditionalOrigins[len(m.AdditionalOrigins)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCsrf(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCsrf
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCsrf
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
func skipCsrf(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCsrf
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
					return 0, ErrIntOverflowCsrf
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
					return 0, ErrIntOverflowCsrf
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
				return 0, ErrInvalidLengthCsrf
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthCsrf
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowCsrf
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
				next, err := skipCsrf(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthCsrf
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
	ErrInvalidLengthCsrf = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCsrf   = fmt.Errorf("proto: integer overflow")
)
