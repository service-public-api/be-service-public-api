// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.23.0--rc1
// source: public-api/repository/grpc/proto/product.proto

package product

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

// ProductServiceClient is the client API for ProductService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProductServiceClient interface {
	GetListKeyProduct(ctx context.Context, in *ListKeyProductServiceRequest, opts ...grpc.CallOption) (*ListKeyBulkResponse, error)
	UpdateListKeyStatusProduct(ctx context.Context, in *UpdateListKeyStatusProductServiceRequest, opts ...grpc.CallOption) (*UpdateListKeyStatusProductServiceResponse, error)
	GetListKeyProductByProductIDAndLimit(ctx context.Context, in *ListKeyProductByProductIDAndLimitServiceRequest, opts ...grpc.CallOption) (*ListKeyBulkResponse, error)
	GetProductByID(ctx context.Context, in *ProductIDRequestServiceRequest, opts ...grpc.CallOption) (*DetailProductServiceResponse, error)
	GetAllProduct(ctx context.Context, in *RequestAdditionalData, opts ...grpc.CallOption) (*GetAllProductResponse, error)
}

type productServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProductServiceClient(cc grpc.ClientConnInterface) ProductServiceClient {
	return &productServiceClient{cc}
}

func (c *productServiceClient) GetListKeyProduct(ctx context.Context, in *ListKeyProductServiceRequest, opts ...grpc.CallOption) (*ListKeyBulkResponse, error) {
	out := new(ListKeyBulkResponse)
	err := c.cc.Invoke(ctx, "/product.ProductService/GetListKeyProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) UpdateListKeyStatusProduct(ctx context.Context, in *UpdateListKeyStatusProductServiceRequest, opts ...grpc.CallOption) (*UpdateListKeyStatusProductServiceResponse, error) {
	out := new(UpdateListKeyStatusProductServiceResponse)
	err := c.cc.Invoke(ctx, "/product.ProductService/UpdateListKeyStatusProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) GetListKeyProductByProductIDAndLimit(ctx context.Context, in *ListKeyProductByProductIDAndLimitServiceRequest, opts ...grpc.CallOption) (*ListKeyBulkResponse, error) {
	out := new(ListKeyBulkResponse)
	err := c.cc.Invoke(ctx, "/product.ProductService/GetListKeyProductByProductIDAndLimit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) GetProductByID(ctx context.Context, in *ProductIDRequestServiceRequest, opts ...grpc.CallOption) (*DetailProductServiceResponse, error) {
	out := new(DetailProductServiceResponse)
	err := c.cc.Invoke(ctx, "/product.ProductService/GetProductByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) GetAllProduct(ctx context.Context, in *RequestAdditionalData, opts ...grpc.CallOption) (*GetAllProductResponse, error) {
	out := new(GetAllProductResponse)
	err := c.cc.Invoke(ctx, "/product.ProductService/GetAllProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductServiceServer is the server API for ProductService service.
// All implementations must embed UnimplementedProductServiceServer
// for forward compatibility
type ProductServiceServer interface {
	GetListKeyProduct(context.Context, *ListKeyProductServiceRequest) (*ListKeyBulkResponse, error)
	UpdateListKeyStatusProduct(context.Context, *UpdateListKeyStatusProductServiceRequest) (*UpdateListKeyStatusProductServiceResponse, error)
	GetListKeyProductByProductIDAndLimit(context.Context, *ListKeyProductByProductIDAndLimitServiceRequest) (*ListKeyBulkResponse, error)
	GetProductByID(context.Context, *ProductIDRequestServiceRequest) (*DetailProductServiceResponse, error)
	GetAllProduct(context.Context, *RequestAdditionalData) (*GetAllProductResponse, error)
	mustEmbedUnimplementedProductServiceServer()
}

// UnimplementedProductServiceServer must be embedded to have forward compatible implementations.
type UnimplementedProductServiceServer struct {
}

func (UnimplementedProductServiceServer) GetListKeyProduct(context.Context, *ListKeyProductServiceRequest) (*ListKeyBulkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetListKeyProduct not implemented")
}
func (UnimplementedProductServiceServer) UpdateListKeyStatusProduct(context.Context, *UpdateListKeyStatusProductServiceRequest) (*UpdateListKeyStatusProductServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateListKeyStatusProduct not implemented")
}
func (UnimplementedProductServiceServer) GetListKeyProductByProductIDAndLimit(context.Context, *ListKeyProductByProductIDAndLimitServiceRequest) (*ListKeyBulkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetListKeyProductByProductIDAndLimit not implemented")
}
func (UnimplementedProductServiceServer) GetProductByID(context.Context, *ProductIDRequestServiceRequest) (*DetailProductServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProductByID not implemented")
}
func (UnimplementedProductServiceServer) GetAllProduct(context.Context, *RequestAdditionalData) (*GetAllProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllProduct not implemented")
}
func (UnimplementedProductServiceServer) mustEmbedUnimplementedProductServiceServer() {}

// UnsafeProductServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProductServiceServer will
// result in compilation errors.
type UnsafeProductServiceServer interface {
	mustEmbedUnimplementedProductServiceServer()
}

func RegisterProductServiceServer(s grpc.ServiceRegistrar, srv ProductServiceServer) {
	s.RegisterService(&ProductService_ServiceDesc, srv)
}

func _ProductService_GetListKeyProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListKeyProductServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).GetListKeyProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product.ProductService/GetListKeyProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).GetListKeyProduct(ctx, req.(*ListKeyProductServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_UpdateListKeyStatusProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateListKeyStatusProductServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).UpdateListKeyStatusProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product.ProductService/UpdateListKeyStatusProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).UpdateListKeyStatusProduct(ctx, req.(*UpdateListKeyStatusProductServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_GetListKeyProductByProductIDAndLimit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListKeyProductByProductIDAndLimitServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).GetListKeyProductByProductIDAndLimit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product.ProductService/GetListKeyProductByProductIDAndLimit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).GetListKeyProductByProductIDAndLimit(ctx, req.(*ListKeyProductByProductIDAndLimitServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_GetProductByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductIDRequestServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).GetProductByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product.ProductService/GetProductByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).GetProductByID(ctx, req.(*ProductIDRequestServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_GetAllProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestAdditionalData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).GetAllProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product.ProductService/GetAllProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).GetAllProduct(ctx, req.(*RequestAdditionalData))
	}
	return interceptor(ctx, in, info, handler)
}

// ProductService_ServiceDesc is the grpc.ServiceDesc for ProductService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProductService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "product.ProductService",
	HandlerType: (*ProductServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetListKeyProduct",
			Handler:    _ProductService_GetListKeyProduct_Handler,
		},
		{
			MethodName: "UpdateListKeyStatusProduct",
			Handler:    _ProductService_UpdateListKeyStatusProduct_Handler,
		},
		{
			MethodName: "GetListKeyProductByProductIDAndLimit",
			Handler:    _ProductService_GetListKeyProductByProductIDAndLimit_Handler,
		},
		{
			MethodName: "GetProductByID",
			Handler:    _ProductService_GetProductByID_Handler,
		},
		{
			MethodName: "GetAllProduct",
			Handler:    _ProductService_GetAllProduct_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "public-api/repository/grpc/proto/product.proto",
}