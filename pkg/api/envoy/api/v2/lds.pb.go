// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/api/v2/lds.proto

package envoy_api_v2

import (
	context "context"
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/datawire/ambassador/pkg/api/envoy/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	_ "istio.io/gogo-genproto/googleapis/google/api"
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

// [#not-implemented-hide:] Not configuration. Workaround c++ protobuf issue with importing
// services: https://github.com/google/protobuf/issues/4221 and protoxform to upgrade the file.
type LdsDummy struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LdsDummy) Reset()         { *m = LdsDummy{} }
func (m *LdsDummy) String() string { return proto.CompactTextString(m) }
func (*LdsDummy) ProtoMessage()    {}
func (*LdsDummy) Descriptor() ([]byte, []int) {
	return fileDescriptor_34e2cd84a105bcd1, []int{0}
}
func (m *LdsDummy) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LdsDummy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LdsDummy.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LdsDummy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LdsDummy.Merge(m, src)
}
func (m *LdsDummy) XXX_Size() int {
	return m.Size()
}
func (m *LdsDummy) XXX_DiscardUnknown() {
	xxx_messageInfo_LdsDummy.DiscardUnknown(m)
}

var xxx_messageInfo_LdsDummy proto.InternalMessageInfo

func init() {
	proto.RegisterType((*LdsDummy)(nil), "envoy.api.v2.LdsDummy")
}

func init() { proto.RegisterFile("envoy/api/v2/lds.proto", fileDescriptor_34e2cd84a105bcd1) }

var fileDescriptor_34e2cd84a105bcd1 = []byte{
	// 389 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0x4f, 0x4b, 0xe3, 0x40,
	0x18, 0xc6, 0x77, 0xba, 0xec, 0x76, 0x19, 0x96, 0x2e, 0x04, 0x76, 0xbb, 0x4d, 0xbb, 0x59, 0xa9,
	0x1e, 0x8a, 0x87, 0x44, 0xda, 0x5b, 0x8f, 0xa5, 0x78, 0xea, 0xa1, 0xd8, 0x4b, 0xaf, 0xd3, 0xe4,
	0xb5, 0x0e, 0x24, 0x99, 0x71, 0x66, 0x12, 0x0d, 0xde, 0x3c, 0x89, 0x17, 0x0f, 0x82, 0xf8, 0x01,
	0xfc, 0x30, 0x1e, 0x05, 0xbf, 0x80, 0x14, 0x3f, 0x81, 0xbd, 0x8b, 0x34, 0xc9, 0xb4, 0xa6, 0x50,
	0x4f, 0xde, 0x92, 0xfc, 0x9e, 0xf7, 0x79, 0xff, 0x3c, 0xc1, 0x7f, 0x20, 0x8c, 0x59, 0xe2, 0x10,
	0x4e, 0x9d, 0xb8, 0xed, 0xf8, 0x9e, 0xb4, 0xb9, 0x60, 0x8a, 0x19, 0x3f, 0xd3, 0xef, 0x36, 0xe1,
	0xd4, 0x8e, 0xdb, 0x66, 0xa3, 0xa0, 0xf2, 0xa8, 0x74, 0x59, 0x0c, 0x22, 0xc9, 0xb4, 0x66, 0x63,
	0xca, 0xd8, 0xd4, 0x87, 0x14, 0x93, 0x30, 0x64, 0x8a, 0x28, 0xca, 0xc2, 0xdc, 0xc9, 0xb4, 0x72,
	0x9a, 0xbe, 0x4d, 0xa2, 0x43, 0xc7, 0x8b, 0x44, 0x2a, 0xd8, 0xc4, 0x4f, 0x04, 0xe1, 0x1c, 0x84,
	0xae, 0xdf, 0xca, 0x7b, 0xaf, 0x8c, 0x1d, 0x01, 0x92, 0x45, 0xc2, 0x05, 0xed, 0x10, 0x79, 0x9c,
	0x14, 0x04, 0x01, 0x9d, 0x0a, 0xa2, 0x34, 0xaf, 0xc6, 0xc4, 0xa7, 0x1e, 0x51, 0xe0, 0xe8, 0x87,
	0x1c, 0xd4, 0x8b, 0xcb, 0x53, 0xa9, 0x20, 0x04, 0x91, 0xc1, 0x26, 0xc6, 0x3f, 0x06, 0x9e, 0xec,
	0x47, 0x41, 0x90, 0xb4, 0xe7, 0x25, 0xfc, 0x77, 0x90, 0xe3, 0xbe, 0xde, 0x7e, 0x04, 0x22, 0xa6,
	0x2e, 0x18, 0x04, 0x57, 0xfa, 0xe0, 0x2b, 0xa2, 0x05, 0xd2, 0xd8, 0xb6, 0xdf, 0x5f, 0xcf, 0x4e,
	0xe9, 0xb2, 0xec, 0x00, 0x8e, 0x23, 0x90, 0xca, 0xdc, 0xf9, 0x58, 0x24, 0x39, 0x0b, 0x25, 0x34,
	0xbf, 0xb4, 0xd0, 0x1e, 0x32, 0xc6, 0xf8, 0xd7, 0x48, 0x09, 0x20, 0xc1, 0xaa, 0x87, 0xb5, 0x56,
	0xbe, 0x6e, 0xff, 0x7f, 0x23, 0x2f, 0x38, 0x9f, 0xe1, 0xca, 0x3e, 0x28, 0xf7, 0xe8, 0x13, 0x8d,
	0x5b, 0xe7, 0x8f, 0xcf, 0xd7, 0xa5, 0x5a, 0xb3, 0x5a, 0xf8, 0x57, 0xba, 0xfa, 0xbc, 0x32, 0xc5,
	0x5f, 0xbb, 0x68, 0xd7, 0xfc, 0x77, 0x79, 0x77, 0x33, 0x2f, 0x57, 0xf1, 0xef, 0x82, 0xa3, 0x1e,
	0xa5, 0x37, 0xbe, 0x9f, 0x59, 0xe8, 0x61, 0x66, 0xa1, 0xa7, 0x99, 0x85, 0x5e, 0x6e, 0x5f, 0xaf,
	0xbe, 0xd5, 0x8d, 0x5a, 0x26, 0x95, 0xd9, 0xed, 0xed, 0x65, 0x66, 0x71, 0x07, 0x9b, 0x94, 0x65,
	0xa3, 0x71, 0xc1, 0x4e, 0x93, 0xc2, 0x94, 0xbd, 0x45, 0x94, 0xc3, 0x45, 0xac, 0x43, 0x74, 0x81,
	0xd0, 0xb0, 0x3c, 0xf9, 0x9e, 0x86, 0xdc, 0x79, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x81, 0x1b, 0xd6,
	0x1d, 0x00, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ListenerDiscoveryServiceClient is the client API for ListenerDiscoveryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ListenerDiscoveryServiceClient interface {
	DeltaListeners(ctx context.Context, opts ...grpc.CallOption) (ListenerDiscoveryService_DeltaListenersClient, error)
	StreamListeners(ctx context.Context, opts ...grpc.CallOption) (ListenerDiscoveryService_StreamListenersClient, error)
	FetchListeners(ctx context.Context, in *DiscoveryRequest, opts ...grpc.CallOption) (*DiscoveryResponse, error)
}

type listenerDiscoveryServiceClient struct {
	cc *grpc.ClientConn
}

func NewListenerDiscoveryServiceClient(cc *grpc.ClientConn) ListenerDiscoveryServiceClient {
	return &listenerDiscoveryServiceClient{cc}
}

func (c *listenerDiscoveryServiceClient) DeltaListeners(ctx context.Context, opts ...grpc.CallOption) (ListenerDiscoveryService_DeltaListenersClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ListenerDiscoveryService_serviceDesc.Streams[0], "/envoy.api.v2.ListenerDiscoveryService/DeltaListeners", opts...)
	if err != nil {
		return nil, err
	}
	x := &listenerDiscoveryServiceDeltaListenersClient{stream}
	return x, nil
}

type ListenerDiscoveryService_DeltaListenersClient interface {
	Send(*DeltaDiscoveryRequest) error
	Recv() (*DeltaDiscoveryResponse, error)
	grpc.ClientStream
}

type listenerDiscoveryServiceDeltaListenersClient struct {
	grpc.ClientStream
}

func (x *listenerDiscoveryServiceDeltaListenersClient) Send(m *DeltaDiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *listenerDiscoveryServiceDeltaListenersClient) Recv() (*DeltaDiscoveryResponse, error) {
	m := new(DeltaDiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *listenerDiscoveryServiceClient) StreamListeners(ctx context.Context, opts ...grpc.CallOption) (ListenerDiscoveryService_StreamListenersClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ListenerDiscoveryService_serviceDesc.Streams[1], "/envoy.api.v2.ListenerDiscoveryService/StreamListeners", opts...)
	if err != nil {
		return nil, err
	}
	x := &listenerDiscoveryServiceStreamListenersClient{stream}
	return x, nil
}

type ListenerDiscoveryService_StreamListenersClient interface {
	Send(*DiscoveryRequest) error
	Recv() (*DiscoveryResponse, error)
	grpc.ClientStream
}

type listenerDiscoveryServiceStreamListenersClient struct {
	grpc.ClientStream
}

func (x *listenerDiscoveryServiceStreamListenersClient) Send(m *DiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *listenerDiscoveryServiceStreamListenersClient) Recv() (*DiscoveryResponse, error) {
	m := new(DiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *listenerDiscoveryServiceClient) FetchListeners(ctx context.Context, in *DiscoveryRequest, opts ...grpc.CallOption) (*DiscoveryResponse, error) {
	out := new(DiscoveryResponse)
	err := c.cc.Invoke(ctx, "/envoy.api.v2.ListenerDiscoveryService/FetchListeners", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ListenerDiscoveryServiceServer is the server API for ListenerDiscoveryService service.
type ListenerDiscoveryServiceServer interface {
	DeltaListeners(ListenerDiscoveryService_DeltaListenersServer) error
	StreamListeners(ListenerDiscoveryService_StreamListenersServer) error
	FetchListeners(context.Context, *DiscoveryRequest) (*DiscoveryResponse, error)
}

// UnimplementedListenerDiscoveryServiceServer can be embedded to have forward compatible implementations.
type UnimplementedListenerDiscoveryServiceServer struct {
}

func (*UnimplementedListenerDiscoveryServiceServer) DeltaListeners(srv ListenerDiscoveryService_DeltaListenersServer) error {
	return status.Errorf(codes.Unimplemented, "method DeltaListeners not implemented")
}
func (*UnimplementedListenerDiscoveryServiceServer) StreamListeners(srv ListenerDiscoveryService_StreamListenersServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamListeners not implemented")
}
func (*UnimplementedListenerDiscoveryServiceServer) FetchListeners(ctx context.Context, req *DiscoveryRequest) (*DiscoveryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchListeners not implemented")
}

func RegisterListenerDiscoveryServiceServer(s *grpc.Server, srv ListenerDiscoveryServiceServer) {
	s.RegisterService(&_ListenerDiscoveryService_serviceDesc, srv)
}

func _ListenerDiscoveryService_DeltaListeners_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ListenerDiscoveryServiceServer).DeltaListeners(&listenerDiscoveryServiceDeltaListenersServer{stream})
}

type ListenerDiscoveryService_DeltaListenersServer interface {
	Send(*DeltaDiscoveryResponse) error
	Recv() (*DeltaDiscoveryRequest, error)
	grpc.ServerStream
}

type listenerDiscoveryServiceDeltaListenersServer struct {
	grpc.ServerStream
}

func (x *listenerDiscoveryServiceDeltaListenersServer) Send(m *DeltaDiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *listenerDiscoveryServiceDeltaListenersServer) Recv() (*DeltaDiscoveryRequest, error) {
	m := new(DeltaDiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ListenerDiscoveryService_StreamListeners_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ListenerDiscoveryServiceServer).StreamListeners(&listenerDiscoveryServiceStreamListenersServer{stream})
}

type ListenerDiscoveryService_StreamListenersServer interface {
	Send(*DiscoveryResponse) error
	Recv() (*DiscoveryRequest, error)
	grpc.ServerStream
}

type listenerDiscoveryServiceStreamListenersServer struct {
	grpc.ServerStream
}

func (x *listenerDiscoveryServiceStreamListenersServer) Send(m *DiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *listenerDiscoveryServiceStreamListenersServer) Recv() (*DiscoveryRequest, error) {
	m := new(DiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ListenerDiscoveryService_FetchListeners_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DiscoveryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ListenerDiscoveryServiceServer).FetchListeners(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/envoy.api.v2.ListenerDiscoveryService/FetchListeners",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ListenerDiscoveryServiceServer).FetchListeners(ctx, req.(*DiscoveryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ListenerDiscoveryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.api.v2.ListenerDiscoveryService",
	HandlerType: (*ListenerDiscoveryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchListeners",
			Handler:    _ListenerDiscoveryService_FetchListeners_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "DeltaListeners",
			Handler:       _ListenerDiscoveryService_DeltaListeners_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "StreamListeners",
			Handler:       _ListenerDiscoveryService_StreamListeners_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "envoy/api/v2/lds.proto",
}

func (m *LdsDummy) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LdsDummy) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LdsDummy) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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

func encodeVarintLds(dAtA []byte, offset int, v uint64) int {
	offset -= sovLds(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *LdsDummy) Size() (n int) {
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

func sovLds(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozLds(x uint64) (n int) {
	return sovLds(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *LdsDummy) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLds
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
			return fmt.Errorf("proto: LdsDummy: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LdsDummy: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipLds(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLds
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthLds
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
func skipLds(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLds
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
					return 0, ErrIntOverflowLds
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
					return 0, ErrIntOverflowLds
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
				return 0, ErrInvalidLengthLds
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthLds
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowLds
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
				next, err := skipLds(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthLds
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
	ErrInvalidLengthLds = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLds   = fmt.Errorf("proto: integer overflow")
)
