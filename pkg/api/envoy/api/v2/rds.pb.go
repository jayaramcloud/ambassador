// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/api/v2/rds.proto

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
type RdsDummy struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RdsDummy) Reset()         { *m = RdsDummy{} }
func (m *RdsDummy) String() string { return proto.CompactTextString(m) }
func (*RdsDummy) ProtoMessage()    {}
func (*RdsDummy) Descriptor() ([]byte, []int) {
	return fileDescriptor_78812f46dcff924a, []int{0}
}
func (m *RdsDummy) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RdsDummy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RdsDummy.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RdsDummy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RdsDummy.Merge(m, src)
}
func (m *RdsDummy) XXX_Size() int {
	return m.Size()
}
func (m *RdsDummy) XXX_DiscardUnknown() {
	xxx_messageInfo_RdsDummy.DiscardUnknown(m)
}

var xxx_messageInfo_RdsDummy proto.InternalMessageInfo

func init() {
	proto.RegisterType((*RdsDummy)(nil), "envoy.api.v2.RdsDummy")
}

func init() { proto.RegisterFile("envoy/api/v2/rds.proto", fileDescriptor_78812f46dcff924a) }

var fileDescriptor_78812f46dcff924a = []byte{
	// 426 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0xbd, 0x8e, 0xd3, 0x40,
	0x14, 0x85, 0x99, 0x20, 0x02, 0x9a, 0xa4, 0xc1, 0x82, 0x10, 0x99, 0xc8, 0x09, 0xe1, 0x57, 0x14,
	0x36, 0x72, 0xba, 0x94, 0x21, 0x42, 0x94, 0x51, 0xa2, 0x50, 0x22, 0x4d, 0xec, 0x89, 0x19, 0xc9,
	0xf6, 0x0c, 0xf3, 0x63, 0x70, 0x4b, 0x85, 0x68, 0x28, 0x90, 0x10, 0x0f, 0x40, 0xc7, 0x8b, 0x50,
	0x22, 0xf1, 0x02, 0xab, 0x68, 0xeb, 0x2d, 0x76, 0xfb, 0xd5, 0xca, 0x63, 0x7b, 0x37, 0xb3, 0xd1,
	0x6e, 0x95, 0xce, 0xf6, 0x77, 0xee, 0xf1, 0xb9, 0x77, 0xee, 0xc0, 0x0e, 0x4e, 0x33, 0x9a, 0x7b,
	0x88, 0x11, 0x2f, 0xf3, 0x3d, 0x1e, 0x0a, 0x97, 0x71, 0x2a, 0xa9, 0xd5, 0xd6, 0xdf, 0x5d, 0xc4,
	0x88, 0x9b, 0xf9, 0x76, 0xcf, 0x50, 0x85, 0x44, 0x04, 0x34, 0xc3, 0x3c, 0x2f, 0xb5, 0x76, 0x2f,
	0xa2, 0x34, 0x8a, 0xb1, 0xc6, 0x28, 0x4d, 0xa9, 0x44, 0x92, 0xd0, 0xb4, 0x72, 0xb2, 0x9d, 0x8a,
	0xea, 0xb7, 0x95, 0x5a, 0x7b, 0x9f, 0x38, 0x62, 0x0c, 0xf3, 0x9a, 0x0f, 0x2a, 0xef, 0x8b, 0x42,
	0x8f, 0x63, 0x41, 0x15, 0x0f, 0x70, 0xed, 0xa0, 0x42, 0x86, 0x0c, 0x41, 0x42, 0x22, 0x8e, 0x64,
	0xcd, 0x1f, 0x64, 0x28, 0x26, 0x21, 0x92, 0xd8, 0xab, 0x1f, 0x2a, 0xd0, 0x35, 0x9b, 0xa3, 0xaa,
	0x26, 0x43, 0x08, 0xef, 0xcc, 0x43, 0x31, 0x55, 0x49, 0x92, 0xfb, 0x47, 0x0d, 0x78, 0x7f, 0x5e,
	0xb0, 0x69, 0xdd, 0xd7, 0x02, 0xf3, 0x8c, 0x04, 0xd8, 0x5a, 0xc2, 0xf6, 0x42, 0x72, 0x8c, 0x12,
	0x8d, 0x85, 0xe5, 0xb8, 0xdb, 0x53, 0x71, 0xcf, 0xf5, 0x73, 0xfc, 0x51, 0x61, 0x21, 0xed, 0xfe,
	0x95, 0x5c, 0x30, 0x9a, 0x0a, 0x3c, 0xbc, 0xf1, 0x02, 0xbc, 0x02, 0xd6, 0x7b, 0xd8, 0x9a, 0xe2,
	0x58, 0xa2, 0xca, 0xf5, 0xf1, 0xa5, 0xaa, 0x02, 0xed, 0x58, 0x3f, 0xb9, 0x5e, 0x64, 0xf8, 0x2b,
	0xd8, 0x7a, 0x83, 0x65, 0xf0, 0x61, 0x5f, 0xa9, 0x9f, 0x7e, 0xf9, 0x7f, 0xf8, 0xa3, 0xd1, 0x19,
	0xde, 0x33, 0x0e, 0x7f, 0xac, 0xe7, 0x29, 0x34, 0xbb, 0x39, 0x06, 0x2f, 0xed, 0xe7, 0xdf, 0x7e,
	0xff, 0x3c, 0xb9, 0xfd, 0x08, 0xf6, 0x0d, 0x3b, 0x1d, 0xe2, 0x35, 0x4d, 0xd7, 0x24, 0x52, 0x5c,
	0x1f, 0x9f, 0xff, 0x07, 0xc0, 0x87, 0xef, 0x08, 0x97, 0x0a, 0xc5, 0x6f, 0xa9, 0x90, 0x3b, 0x63,
	0x0f, 0xe1, 0x5d, 0xdd, 0xdf, 0x96, 0x66, 0xff, 0x53, 0xb2, 0x9f, 0xe9, 0xb8, 0x03, 0xe8, 0x18,
	0x25, 0xe5, 0x92, 0x6c, 0xfd, 0x79, 0xb2, 0xfc, 0xbb, 0x71, 0xc0, 0xbf, 0x8d, 0x03, 0x0e, 0x36,
	0x0e, 0x38, 0xfe, 0x75, 0xfa, 0xfd, 0x56, 0xd7, 0x2a, 0x6f, 0x8d, 0x2b, 0xca, 0xc0, 0x55, 0x51,
	0x36, 0x82, 0x36, 0xa1, 0x65, 0x02, 0xc6, 0xe9, 0xe7, 0xdc, 0x08, 0x33, 0x29, 0x16, 0x6e, 0x56,
	0x2c, 0xdf, 0x0c, 0x7c, 0x05, 0x60, 0xd6, 0x5c, 0x35, 0xf5, 0x2a, 0x8e, 0xce, 0x02, 0x00, 0x00,
	0xff, 0xff, 0xa9, 0xb7, 0xdf, 0x8c, 0x83, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RouteDiscoveryServiceClient is the client API for RouteDiscoveryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RouteDiscoveryServiceClient interface {
	StreamRoutes(ctx context.Context, opts ...grpc.CallOption) (RouteDiscoveryService_StreamRoutesClient, error)
	DeltaRoutes(ctx context.Context, opts ...grpc.CallOption) (RouteDiscoveryService_DeltaRoutesClient, error)
	FetchRoutes(ctx context.Context, in *DiscoveryRequest, opts ...grpc.CallOption) (*DiscoveryResponse, error)
}

type routeDiscoveryServiceClient struct {
	cc *grpc.ClientConn
}

func NewRouteDiscoveryServiceClient(cc *grpc.ClientConn) RouteDiscoveryServiceClient {
	return &routeDiscoveryServiceClient{cc}
}

func (c *routeDiscoveryServiceClient) StreamRoutes(ctx context.Context, opts ...grpc.CallOption) (RouteDiscoveryService_StreamRoutesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_RouteDiscoveryService_serviceDesc.Streams[0], "/envoy.api.v2.RouteDiscoveryService/StreamRoutes", opts...)
	if err != nil {
		return nil, err
	}
	x := &routeDiscoveryServiceStreamRoutesClient{stream}
	return x, nil
}

type RouteDiscoveryService_StreamRoutesClient interface {
	Send(*DiscoveryRequest) error
	Recv() (*DiscoveryResponse, error)
	grpc.ClientStream
}

type routeDiscoveryServiceStreamRoutesClient struct {
	grpc.ClientStream
}

func (x *routeDiscoveryServiceStreamRoutesClient) Send(m *DiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *routeDiscoveryServiceStreamRoutesClient) Recv() (*DiscoveryResponse, error) {
	m := new(DiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *routeDiscoveryServiceClient) DeltaRoutes(ctx context.Context, opts ...grpc.CallOption) (RouteDiscoveryService_DeltaRoutesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_RouteDiscoveryService_serviceDesc.Streams[1], "/envoy.api.v2.RouteDiscoveryService/DeltaRoutes", opts...)
	if err != nil {
		return nil, err
	}
	x := &routeDiscoveryServiceDeltaRoutesClient{stream}
	return x, nil
}

type RouteDiscoveryService_DeltaRoutesClient interface {
	Send(*DeltaDiscoveryRequest) error
	Recv() (*DeltaDiscoveryResponse, error)
	grpc.ClientStream
}

type routeDiscoveryServiceDeltaRoutesClient struct {
	grpc.ClientStream
}

func (x *routeDiscoveryServiceDeltaRoutesClient) Send(m *DeltaDiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *routeDiscoveryServiceDeltaRoutesClient) Recv() (*DeltaDiscoveryResponse, error) {
	m := new(DeltaDiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *routeDiscoveryServiceClient) FetchRoutes(ctx context.Context, in *DiscoveryRequest, opts ...grpc.CallOption) (*DiscoveryResponse, error) {
	out := new(DiscoveryResponse)
	err := c.cc.Invoke(ctx, "/envoy.api.v2.RouteDiscoveryService/FetchRoutes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RouteDiscoveryServiceServer is the server API for RouteDiscoveryService service.
type RouteDiscoveryServiceServer interface {
	StreamRoutes(RouteDiscoveryService_StreamRoutesServer) error
	DeltaRoutes(RouteDiscoveryService_DeltaRoutesServer) error
	FetchRoutes(context.Context, *DiscoveryRequest) (*DiscoveryResponse, error)
}

// UnimplementedRouteDiscoveryServiceServer can be embedded to have forward compatible implementations.
type UnimplementedRouteDiscoveryServiceServer struct {
}

func (*UnimplementedRouteDiscoveryServiceServer) StreamRoutes(srv RouteDiscoveryService_StreamRoutesServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamRoutes not implemented")
}
func (*UnimplementedRouteDiscoveryServiceServer) DeltaRoutes(srv RouteDiscoveryService_DeltaRoutesServer) error {
	return status.Errorf(codes.Unimplemented, "method DeltaRoutes not implemented")
}
func (*UnimplementedRouteDiscoveryServiceServer) FetchRoutes(ctx context.Context, req *DiscoveryRequest) (*DiscoveryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchRoutes not implemented")
}

func RegisterRouteDiscoveryServiceServer(s *grpc.Server, srv RouteDiscoveryServiceServer) {
	s.RegisterService(&_RouteDiscoveryService_serviceDesc, srv)
}

func _RouteDiscoveryService_StreamRoutes_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RouteDiscoveryServiceServer).StreamRoutes(&routeDiscoveryServiceStreamRoutesServer{stream})
}

type RouteDiscoveryService_StreamRoutesServer interface {
	Send(*DiscoveryResponse) error
	Recv() (*DiscoveryRequest, error)
	grpc.ServerStream
}

type routeDiscoveryServiceStreamRoutesServer struct {
	grpc.ServerStream
}

func (x *routeDiscoveryServiceStreamRoutesServer) Send(m *DiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *routeDiscoveryServiceStreamRoutesServer) Recv() (*DiscoveryRequest, error) {
	m := new(DiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _RouteDiscoveryService_DeltaRoutes_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RouteDiscoveryServiceServer).DeltaRoutes(&routeDiscoveryServiceDeltaRoutesServer{stream})
}

type RouteDiscoveryService_DeltaRoutesServer interface {
	Send(*DeltaDiscoveryResponse) error
	Recv() (*DeltaDiscoveryRequest, error)
	grpc.ServerStream
}

type routeDiscoveryServiceDeltaRoutesServer struct {
	grpc.ServerStream
}

func (x *routeDiscoveryServiceDeltaRoutesServer) Send(m *DeltaDiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *routeDiscoveryServiceDeltaRoutesServer) Recv() (*DeltaDiscoveryRequest, error) {
	m := new(DeltaDiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _RouteDiscoveryService_FetchRoutes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DiscoveryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouteDiscoveryServiceServer).FetchRoutes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/envoy.api.v2.RouteDiscoveryService/FetchRoutes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouteDiscoveryServiceServer).FetchRoutes(ctx, req.(*DiscoveryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RouteDiscoveryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.api.v2.RouteDiscoveryService",
	HandlerType: (*RouteDiscoveryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchRoutes",
			Handler:    _RouteDiscoveryService_FetchRoutes_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamRoutes",
			Handler:       _RouteDiscoveryService_StreamRoutes_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "DeltaRoutes",
			Handler:       _RouteDiscoveryService_DeltaRoutes_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "envoy/api/v2/rds.proto",
}

// VirtualHostDiscoveryServiceClient is the client API for VirtualHostDiscoveryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type VirtualHostDiscoveryServiceClient interface {
	DeltaVirtualHosts(ctx context.Context, opts ...grpc.CallOption) (VirtualHostDiscoveryService_DeltaVirtualHostsClient, error)
}

type virtualHostDiscoveryServiceClient struct {
	cc *grpc.ClientConn
}

func NewVirtualHostDiscoveryServiceClient(cc *grpc.ClientConn) VirtualHostDiscoveryServiceClient {
	return &virtualHostDiscoveryServiceClient{cc}
}

func (c *virtualHostDiscoveryServiceClient) DeltaVirtualHosts(ctx context.Context, opts ...grpc.CallOption) (VirtualHostDiscoveryService_DeltaVirtualHostsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_VirtualHostDiscoveryService_serviceDesc.Streams[0], "/envoy.api.v2.VirtualHostDiscoveryService/DeltaVirtualHosts", opts...)
	if err != nil {
		return nil, err
	}
	x := &virtualHostDiscoveryServiceDeltaVirtualHostsClient{stream}
	return x, nil
}

type VirtualHostDiscoveryService_DeltaVirtualHostsClient interface {
	Send(*DeltaDiscoveryRequest) error
	Recv() (*DeltaDiscoveryResponse, error)
	grpc.ClientStream
}

type virtualHostDiscoveryServiceDeltaVirtualHostsClient struct {
	grpc.ClientStream
}

func (x *virtualHostDiscoveryServiceDeltaVirtualHostsClient) Send(m *DeltaDiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *virtualHostDiscoveryServiceDeltaVirtualHostsClient) Recv() (*DeltaDiscoveryResponse, error) {
	m := new(DeltaDiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// VirtualHostDiscoveryServiceServer is the server API for VirtualHostDiscoveryService service.
type VirtualHostDiscoveryServiceServer interface {
	DeltaVirtualHosts(VirtualHostDiscoveryService_DeltaVirtualHostsServer) error
}

// UnimplementedVirtualHostDiscoveryServiceServer can be embedded to have forward compatible implementations.
type UnimplementedVirtualHostDiscoveryServiceServer struct {
}

func (*UnimplementedVirtualHostDiscoveryServiceServer) DeltaVirtualHosts(srv VirtualHostDiscoveryService_DeltaVirtualHostsServer) error {
	return status.Errorf(codes.Unimplemented, "method DeltaVirtualHosts not implemented")
}

func RegisterVirtualHostDiscoveryServiceServer(s *grpc.Server, srv VirtualHostDiscoveryServiceServer) {
	s.RegisterService(&_VirtualHostDiscoveryService_serviceDesc, srv)
}

func _VirtualHostDiscoveryService_DeltaVirtualHosts_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(VirtualHostDiscoveryServiceServer).DeltaVirtualHosts(&virtualHostDiscoveryServiceDeltaVirtualHostsServer{stream})
}

type VirtualHostDiscoveryService_DeltaVirtualHostsServer interface {
	Send(*DeltaDiscoveryResponse) error
	Recv() (*DeltaDiscoveryRequest, error)
	grpc.ServerStream
}

type virtualHostDiscoveryServiceDeltaVirtualHostsServer struct {
	grpc.ServerStream
}

func (x *virtualHostDiscoveryServiceDeltaVirtualHostsServer) Send(m *DeltaDiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *virtualHostDiscoveryServiceDeltaVirtualHostsServer) Recv() (*DeltaDiscoveryRequest, error) {
	m := new(DeltaDiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _VirtualHostDiscoveryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.api.v2.VirtualHostDiscoveryService",
	HandlerType: (*VirtualHostDiscoveryServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "DeltaVirtualHosts",
			Handler:       _VirtualHostDiscoveryService_DeltaVirtualHosts_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "envoy/api/v2/rds.proto",
}

func (m *RdsDummy) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RdsDummy) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RdsDummy) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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

func encodeVarintRds(dAtA []byte, offset int, v uint64) int {
	offset -= sovRds(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *RdsDummy) Size() (n int) {
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

func sovRds(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozRds(x uint64) (n int) {
	return sovRds(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RdsDummy) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRds
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
			return fmt.Errorf("proto: RdsDummy: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RdsDummy: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipRds(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRds
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthRds
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
func skipRds(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRds
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
					return 0, ErrIntOverflowRds
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
					return 0, ErrIntOverflowRds
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
				return 0, ErrInvalidLengthRds
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthRds
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowRds
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
				next, err := skipRds(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthRds
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
	ErrInvalidLengthRds = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRds   = fmt.Errorf("proto: integer overflow")
)
