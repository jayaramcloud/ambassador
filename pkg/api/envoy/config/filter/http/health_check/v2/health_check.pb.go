// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/config/filter/http/health_check/v2/health_check.proto

package envoy_config_filter_http_health_check_v2

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	route "github.com/datawire/ambassador/pkg/api/envoy/api/v2/route"
	_type "github.com/datawire/ambassador/pkg/api/envoy/type"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
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

// [#next-free-field: 6]
type HealthCheck struct {
	// Specifies whether the filter operates in pass through mode or not.
	PassThroughMode *types.BoolValue `protobuf:"bytes,1,opt,name=pass_through_mode,json=passThroughMode,proto3" json:"pass_through_mode,omitempty"`
	// If operating in pass through mode, the amount of time in milliseconds
	// that the filter should cache the upstream response.
	CacheTime *types.Duration `protobuf:"bytes,3,opt,name=cache_time,json=cacheTime,proto3" json:"cache_time,omitempty"`
	// If operating in non-pass-through mode, specifies a set of upstream cluster
	// names and the minimum percentage of servers in each of those clusters that
	// must be healthy or degraded in order for the filter to return a 200.
	ClusterMinHealthyPercentages map[string]*_type.Percent `protobuf:"bytes,4,rep,name=cluster_min_healthy_percentages,json=clusterMinHealthyPercentages,proto3" json:"cluster_min_healthy_percentages,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Specifies a set of health check request headers to match on. The health check filter will
	// check a request’s headers against all the specified headers. To specify the health check
	// endpoint, set the ``:path`` header to match on.
	Headers              []*route.HeaderMatcher `protobuf:"bytes,5,rep,name=headers,proto3" json:"headers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *HealthCheck) Reset()         { *m = HealthCheck{} }
func (m *HealthCheck) String() string { return proto.CompactTextString(m) }
func (*HealthCheck) ProtoMessage()    {}
func (*HealthCheck) Descriptor() ([]byte, []int) {
	return fileDescriptor_75439d7b4d98e201, []int{0}
}
func (m *HealthCheck) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HealthCheck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HealthCheck.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *HealthCheck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheck.Merge(m, src)
}
func (m *HealthCheck) XXX_Size() int {
	return m.Size()
}
func (m *HealthCheck) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheck.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheck proto.InternalMessageInfo

func (m *HealthCheck) GetPassThroughMode() *types.BoolValue {
	if m != nil {
		return m.PassThroughMode
	}
	return nil
}

func (m *HealthCheck) GetCacheTime() *types.Duration {
	if m != nil {
		return m.CacheTime
	}
	return nil
}

func (m *HealthCheck) GetClusterMinHealthyPercentages() map[string]*_type.Percent {
	if m != nil {
		return m.ClusterMinHealthyPercentages
	}
	return nil
}

func (m *HealthCheck) GetHeaders() []*route.HeaderMatcher {
	if m != nil {
		return m.Headers
	}
	return nil
}

func init() {
	proto.RegisterType((*HealthCheck)(nil), "envoy.config.filter.http.health_check.v2.HealthCheck")
	proto.RegisterMapType((map[string]*_type.Percent)(nil), "envoy.config.filter.http.health_check.v2.HealthCheck.ClusterMinHealthyPercentagesEntry")
}

func init() {
	proto.RegisterFile("envoy/config/filter/http/health_check/v2/health_check.proto", fileDescriptor_75439d7b4d98e201)
}

var fileDescriptor_75439d7b4d98e201 = []byte{
	// 511 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0x4d, 0x6b, 0x14, 0x3f,
	0x18, 0x27, 0xbb, 0xdd, 0xff, 0xbf, 0xcd, 0x1e, 0x5c, 0xc7, 0x83, 0xe3, 0x22, 0x6b, 0xeb, 0x69,
	0x7b, 0x30, 0x81, 0x2d, 0x48, 0xb1, 0xb7, 0xad, 0x82, 0x08, 0x0b, 0xcb, 0x50, 0xf4, 0x38, 0xa4,
	0x99, 0x67, 0x67, 0x42, 0x67, 0x92, 0x90, 0xc9, 0xac, 0x9d, 0xaf, 0x20, 0x88, 0x57, 0x4f, 0x7e,
	0x1e, 0x8f, 0x7e, 0x04, 0x59, 0xbf, 0x81, 0x47, 0x0f, 0x22, 0x49, 0x66, 0xb1, 0xc5, 0xb7, 0x5e,
	0x86, 0x64, 0x9e, 0xdf, 0xcb, 0xf3, 0xfc, 0x9e, 0xe0, 0x13, 0x90, 0x6b, 0xd5, 0x52, 0xae, 0xe4,
	0x4a, 0xe4, 0x74, 0x25, 0x4a, 0x0b, 0x86, 0x16, 0xd6, 0x6a, 0x5a, 0x00, 0x2b, 0x6d, 0x91, 0xf2,
	0x02, 0xf8, 0x05, 0x5d, 0xcf, 0xae, 0xdd, 0x89, 0x36, 0xca, 0xaa, 0x68, 0xea, 0xc9, 0x24, 0x90,
	0x49, 0x20, 0x13, 0x47, 0x26, 0xd7, 0xc0, 0xeb, 0xd9, 0xf8, 0x30, 0xd8, 0x30, 0x2d, 0x9c, 0x94,
	0x51, 0x8d, 0x85, 0xf0, 0x4d, 0xb9, 0xaa, 0xb4, 0x92, 0x20, 0x6d, 0x1d, 0x44, 0xc7, 0x71, 0x80,
	0xda, 0x56, 0x03, 0xd5, 0x60, 0x38, 0x48, 0xdb, 0x55, 0x26, 0xb9, 0x52, 0x79, 0x09, 0xd4, 0xdf,
	0xce, 0x9b, 0x15, 0xcd, 0x1a, 0xc3, 0xac, 0x50, 0xf2, 0x4f, 0xf5, 0xd7, 0x86, 0x69, 0x0d, 0x66,
	0xab, 0x3c, 0x69, 0x32, 0xcd, 0x28, 0x93, 0x52, 0x59, 0x4f, 0xab, 0x69, 0x25, 0x72, 0xc3, 0x2c,
	0x74, 0xf5, 0xbb, 0x6b, 0x56, 0x8a, 0x8c, 0x59, 0xa0, 0xdb, 0x43, 0x28, 0x3c, 0xfc, 0xd2, 0xc7,
	0xc3, 0xe7, 0x7e, 0xa2, 0x53, 0x37, 0x50, 0xb4, 0xc4, 0xb7, 0x35, 0xab, 0xeb, 0xd4, 0x16, 0x46,
	0x35, 0x79, 0x91, 0x56, 0x2a, 0x83, 0x18, 0xed, 0xa3, 0xe9, 0x70, 0x36, 0x26, 0xa1, 0x09, 0xb2,
	0x6d, 0x82, 0xcc, 0x95, 0x2a, 0x5f, 0xb2, 0xb2, 0x81, 0xf9, 0xee, 0xb7, 0xf9, 0xe0, 0x0d, 0xea,
	0x8d, 0x50, 0x72, 0xcb, 0xd1, 0xcf, 0x02, 0x7b, 0xa1, 0x32, 0x88, 0x8e, 0x31, 0xe6, 0x8c, 0x17,
	0x90, 0x5a, 0x51, 0x41, 0xdc, 0xf7, 0x52, 0xf7, 0x7e, 0x91, 0x7a, 0xda, 0xcd, 0x9b, 0xec, 0x79,
	0xf0, 0x99, 0xa8, 0x20, 0xfa, 0x80, 0xf0, 0x03, 0x5e, 0x36, 0xb5, 0x05, 0x93, 0x56, 0x42, 0xa6,
	0x21, 0xf9, 0x36, 0xed, 0xa2, 0x63, 0x39, 0xd4, 0xf1, 0xce, 0x7e, 0x7f, 0x3a, 0x9c, 0xbd, 0x22,
	0x37, 0x5d, 0x17, 0xb9, 0x32, 0x2c, 0x39, 0x0d, 0xe2, 0x0b, 0x21, 0xc3, 0xdf, 0x76, 0xf9, 0x53,
	0xf9, 0x99, 0xb4, 0xa6, 0x4d, 0xee, 0xf3, 0xbf, 0x40, 0xa2, 0x13, 0xfc, 0x7f, 0x01, 0x2c, 0x03,
	0x53, 0xc7, 0x03, 0xdf, 0xc7, 0x41, 0xd7, 0x07, 0xd3, 0xc2, 0x79, 0xf9, 0x67, 0xe0, 0x1c, 0x33,
	0x30, 0x0b, 0x66, 0x79, 0x01, 0x26, 0xd9, 0x32, 0xc6, 0x19, 0x3e, 0xf8, 0xa7, 0x7f, 0x34, 0xc2,
	0xfd, 0x0b, 0x68, 0xfd, 0x02, 0xf6, 0x12, 0x77, 0x8c, 0x0e, 0xf1, 0x60, 0xed, 0x22, 0x8f, 0x7b,
	0x3e, 0xc9, 0x3b, 0x9d, 0xa3, 0x7b, 0x53, 0xa4, 0xa3, 0x27, 0x01, 0xf1, 0xa4, 0x77, 0x8c, 0x5e,
	0xec, 0xec, 0xf6, 0x46, 0xfd, 0xf9, 0x5b, 0xf4, 0x71, 0x33, 0x41, 0x9f, 0x36, 0x13, 0xf4, 0x79,
	0x33, 0x41, 0x5f, 0xdf, 0x7f, 0x7f, 0x37, 0xa0, 0xd1, 0xa3, 0xc0, 0x84, 0x4b, 0x0b, 0xb2, 0x76,
	0x6f, 0xa6, 0xcb, 0xad, 0xfe, 0x5d, 0x70, 0x47, 0xf8, 0xb1, 0x50, 0xc1, 0x4b, 0x1b, 0x75, 0xd9,
	0xde, 0x38, 0xf0, 0xf9, 0xe8, 0x4a, 0xe2, 0x4b, 0xb7, 0xee, 0x25, 0x3a, 0xff, 0xcf, 0xef, 0xfd,
	0xe8, 0x47, 0x00, 0x00, 0x00, 0xff, 0xff, 0x87, 0xd3, 0x71, 0xd4, 0xa3, 0x03, 0x00, 0x00,
}

func (m *HealthCheck) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HealthCheck) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *HealthCheck) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Headers) > 0 {
		for iNdEx := len(m.Headers) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Headers[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintHealthCheck(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.ClusterMinHealthyPercentages) > 0 {
		for k := range m.ClusterMinHealthyPercentages {
			v := m.ClusterMinHealthyPercentages[k]
			baseI := i
			if v != nil {
				{
					size, err := v.MarshalToSizedBuffer(dAtA[:i])
					if err != nil {
						return 0, err
					}
					i -= size
					i = encodeVarintHealthCheck(dAtA, i, uint64(size))
				}
				i--
				dAtA[i] = 0x12
			}
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintHealthCheck(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintHealthCheck(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x22
		}
	}
	if m.CacheTime != nil {
		{
			size, err := m.CacheTime.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintHealthCheck(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.PassThroughMode != nil {
		{
			size, err := m.PassThroughMode.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintHealthCheck(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintHealthCheck(dAtA []byte, offset int, v uint64) int {
	offset -= sovHealthCheck(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *HealthCheck) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PassThroughMode != nil {
		l = m.PassThroughMode.Size()
		n += 1 + l + sovHealthCheck(uint64(l))
	}
	if m.CacheTime != nil {
		l = m.CacheTime.Size()
		n += 1 + l + sovHealthCheck(uint64(l))
	}
	if len(m.ClusterMinHealthyPercentages) > 0 {
		for k, v := range m.ClusterMinHealthyPercentages {
			_ = k
			_ = v
			l = 0
			if v != nil {
				l = v.Size()
				l += 1 + sovHealthCheck(uint64(l))
			}
			mapEntrySize := 1 + len(k) + sovHealthCheck(uint64(len(k))) + l
			n += mapEntrySize + 1 + sovHealthCheck(uint64(mapEntrySize))
		}
	}
	if len(m.Headers) > 0 {
		for _, e := range m.Headers {
			l = e.Size()
			n += 1 + l + sovHealthCheck(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovHealthCheck(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozHealthCheck(x uint64) (n int) {
	return sovHealthCheck(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *HealthCheck) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHealthCheck
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
			return fmt.Errorf("proto: HealthCheck: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HealthCheck: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PassThroughMode", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHealthCheck
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
				return ErrInvalidLengthHealthCheck
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthHealthCheck
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.PassThroughMode == nil {
				m.PassThroughMode = &types.BoolValue{}
			}
			if err := m.PassThroughMode.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CacheTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHealthCheck
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
				return ErrInvalidLengthHealthCheck
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthHealthCheck
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.CacheTime == nil {
				m.CacheTime = &types.Duration{}
			}
			if err := m.CacheTime.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClusterMinHealthyPercentages", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHealthCheck
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
				return ErrInvalidLengthHealthCheck
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthHealthCheck
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ClusterMinHealthyPercentages == nil {
				m.ClusterMinHealthyPercentages = make(map[string]*_type.Percent)
			}
			var mapkey string
			var mapvalue *_type.Percent
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowHealthCheck
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
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowHealthCheck
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthHealthCheck
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthHealthCheck
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var mapmsglen int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowHealthCheck
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapmsglen |= int(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					if mapmsglen < 0 {
						return ErrInvalidLengthHealthCheck
					}
					postmsgIndex := iNdEx + mapmsglen
					if postmsgIndex < 0 {
						return ErrInvalidLengthHealthCheck
					}
					if postmsgIndex > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = &_type.Percent{}
					if err := mapvalue.Unmarshal(dAtA[iNdEx:postmsgIndex]); err != nil {
						return err
					}
					iNdEx = postmsgIndex
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipHealthCheck(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthHealthCheck
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.ClusterMinHealthyPercentages[mapkey] = mapvalue
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Headers", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHealthCheck
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
				return ErrInvalidLengthHealthCheck
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthHealthCheck
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Headers = append(m.Headers, &route.HeaderMatcher{})
			if err := m.Headers[len(m.Headers)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHealthCheck(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHealthCheck
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthHealthCheck
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
func skipHealthCheck(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowHealthCheck
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
					return 0, ErrIntOverflowHealthCheck
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
					return 0, ErrIntOverflowHealthCheck
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
				return 0, ErrInvalidLengthHealthCheck
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthHealthCheck
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowHealthCheck
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
				next, err := skipHealthCheck(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthHealthCheck
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
	ErrInvalidLengthHealthCheck = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowHealthCheck   = fmt.Errorf("proto: integer overflow")
)
