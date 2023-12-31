// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
// source: analytics.proto

package genproto

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

const (
	AnalyticsService_GetTotalSales_FullMethodName     = "/GoSalesStream.analytics.v1.AnalyticsService/GetTotalSales"
	AnalyticsService_GetSalesByProduct_FullMethodName = "/GoSalesStream.analytics.v1.AnalyticsService/GetSalesByProduct"
	AnalyticsService_GetTop5Customers_FullMethodName  = "/GoSalesStream.analytics.v1.AnalyticsService/GetTop5Customers"
)

// AnalyticsServiceClient is the client API for AnalyticsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AnalyticsServiceClient interface {
	GetTotalSales(ctx context.Context, in *TotalSalesRequest, opts ...grpc.CallOption) (*TotalSales, error)
	GetSalesByProduct(ctx context.Context, in *SalesByProductRequest, opts ...grpc.CallOption) (*SalesByProductResponse, error)
	GetTop5Customers(ctx context.Context, in *Top5CustomersRequest, opts ...grpc.CallOption) (*Top5CustomersResponse, error)
}

type analyticsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAnalyticsServiceClient(cc grpc.ClientConnInterface) AnalyticsServiceClient {
	return &analyticsServiceClient{cc}
}

func (c *analyticsServiceClient) GetTotalSales(ctx context.Context, in *TotalSalesRequest, opts ...grpc.CallOption) (*TotalSales, error) {
	out := new(TotalSales)
	err := c.cc.Invoke(ctx, AnalyticsService_GetTotalSales_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *analyticsServiceClient) GetSalesByProduct(ctx context.Context, in *SalesByProductRequest, opts ...grpc.CallOption) (*SalesByProductResponse, error) {
	out := new(SalesByProductResponse)
	err := c.cc.Invoke(ctx, AnalyticsService_GetSalesByProduct_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *analyticsServiceClient) GetTop5Customers(ctx context.Context, in *Top5CustomersRequest, opts ...grpc.CallOption) (*Top5CustomersResponse, error) {
	out := new(Top5CustomersResponse)
	err := c.cc.Invoke(ctx, AnalyticsService_GetTop5Customers_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AnalyticsServiceServer is the server API for AnalyticsService service.
// All implementations must embed UnimplementedAnalyticsServiceServer
// for forward compatibility
type AnalyticsServiceServer interface {
	GetTotalSales(context.Context, *TotalSalesRequest) (*TotalSales, error)
	GetSalesByProduct(context.Context, *SalesByProductRequest) (*SalesByProductResponse, error)
	GetTop5Customers(context.Context, *Top5CustomersRequest) (*Top5CustomersResponse, error)
	mustEmbedUnimplementedAnalyticsServiceServer()
}

// UnimplementedAnalyticsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAnalyticsServiceServer struct {
}

func (UnimplementedAnalyticsServiceServer) GetTotalSales(context.Context, *TotalSalesRequest) (*TotalSales, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTotalSales not implemented")
}
func (UnimplementedAnalyticsServiceServer) GetSalesByProduct(context.Context, *SalesByProductRequest) (*SalesByProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSalesByProduct not implemented")
}
func (UnimplementedAnalyticsServiceServer) GetTop5Customers(context.Context, *Top5CustomersRequest) (*Top5CustomersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTop5Customers not implemented")
}
func (UnimplementedAnalyticsServiceServer) mustEmbedUnimplementedAnalyticsServiceServer() {}

// UnsafeAnalyticsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AnalyticsServiceServer will
// result in compilation errors.
type UnsafeAnalyticsServiceServer interface {
	mustEmbedUnimplementedAnalyticsServiceServer()
}

func RegisterAnalyticsServiceServer(s grpc.ServiceRegistrar, srv AnalyticsServiceServer) {
	s.RegisterService(&AnalyticsService_ServiceDesc, srv)
}

func _AnalyticsService_GetTotalSales_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TotalSalesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnalyticsServiceServer).GetTotalSales(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AnalyticsService_GetTotalSales_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnalyticsServiceServer).GetTotalSales(ctx, req.(*TotalSalesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AnalyticsService_GetSalesByProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SalesByProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnalyticsServiceServer).GetSalesByProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AnalyticsService_GetSalesByProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnalyticsServiceServer).GetSalesByProduct(ctx, req.(*SalesByProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AnalyticsService_GetTop5Customers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Top5CustomersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnalyticsServiceServer).GetTop5Customers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AnalyticsService_GetTop5Customers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnalyticsServiceServer).GetTop5Customers(ctx, req.(*Top5CustomersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AnalyticsService_ServiceDesc is the grpc.ServiceDesc for AnalyticsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AnalyticsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "GoSalesStream.analytics.v1.AnalyticsService",
	HandlerType: (*AnalyticsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTotalSales",
			Handler:    _AnalyticsService_GetTotalSales_Handler,
		},
		{
			MethodName: "GetSalesByProduct",
			Handler:    _AnalyticsService_GetSalesByProduct_Handler,
		},
		{
			MethodName: "GetTop5Customers",
			Handler:    _AnalyticsService_GetTop5Customers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "analytics.proto",
}
