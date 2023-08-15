// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.23.4
// source: protos/payment_guide.proto

package __

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// PaymentGuideClient is the client API for PaymentGuide service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PaymentGuideClient interface {
	CreatePayment(ctx context.Context, in *PaymentIntent, opts ...grpc.CallOption) (*Payment, error)
}

type paymentGuideClient struct {
	cc grpc.ClientConnInterface
}

func NewPaymentGuideClient(cc grpc.ClientConnInterface) PaymentGuideClient {
	return &paymentGuideClient{cc}
}

func (c *paymentGuideClient) CreatePayment(ctx context.Context, in *PaymentIntent, opts ...grpc.CallOption) (*Payment, error) {
	out := new(Payment)
	err := c.cc.Invoke(ctx, "/protos.PaymentGuide/CreatePayment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PaymentGuideServer is the server API for PaymentGuide service.
// All implementations must embed UnimplementedPaymentGuideServer
// for forward compatibility
type PaymentGuideServer interface {
	CreatePayment(context.Context, *PaymentIntent) (*Payment, error)
	mustEmbedUnimplementedPaymentGuideServer()
}

// UnimplementedPaymentGuideServer must be embedded to have forward compatible implementations.
type UnimplementedPaymentGuideServer struct {
}

func (UnimplementedPaymentGuideServer) CreatePayment(context.Context, *PaymentIntent) (*Payment, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePayment not implemented")
}
func (UnimplementedPaymentGuideServer) mustEmbedUnimplementedPaymentGuideServer() {}

// UnsafePaymentGuideServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PaymentGuideServer will
// result in compilation errors.
type UnsafePaymentGuideServer interface {
	mustEmbedUnimplementedPaymentGuideServer()
}

func RegisterPaymentGuideServer(s grpc.ServiceRegistrar, srv PaymentGuideServer) {
	s.RegisterService(&PaymentGuide_ServiceDesc, srv)
}

func _PaymentGuide_CreatePayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PaymentIntent)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentGuideServer).CreatePayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.PaymentGuide/CreatePayment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentGuideServer).CreatePayment(ctx, req.(*PaymentIntent))
	}
	return interceptor(ctx, in, info, handler)
}

// PaymentGuide_ServiceDesc is the grpc.ServiceDesc for PaymentGuide service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PaymentGuide_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protos.PaymentGuide",
	HandlerType: (*PaymentGuideServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePayment",
			Handler:    _PaymentGuide_CreatePayment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/payment_guide.proto",
}
