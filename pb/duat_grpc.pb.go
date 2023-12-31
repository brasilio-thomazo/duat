// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.3
// source: duat.proto

package pb

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

// DepartmentServiceClient is the client API for DepartmentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DepartmentServiceClient interface {
	FindByID(ctx context.Context, in *IDRequest, opts ...grpc.CallOption) (*DepartmentReply, error)
	All(ctx context.Context, in *PageRequest, opts ...grpc.CallOption) (*DepartmentListReply, error)
}

type departmentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDepartmentServiceClient(cc grpc.ClientConnInterface) DepartmentServiceClient {
	return &departmentServiceClient{cc}
}

func (c *departmentServiceClient) FindByID(ctx context.Context, in *IDRequest, opts ...grpc.CallOption) (*DepartmentReply, error) {
	out := new(DepartmentReply)
	err := c.cc.Invoke(ctx, "/duat.DepartmentService/FindByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *departmentServiceClient) All(ctx context.Context, in *PageRequest, opts ...grpc.CallOption) (*DepartmentListReply, error) {
	out := new(DepartmentListReply)
	err := c.cc.Invoke(ctx, "/duat.DepartmentService/All", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DepartmentServiceServer is the server API for DepartmentService service.
// All implementations must embed UnimplementedDepartmentServiceServer
// for forward compatibility
type DepartmentServiceServer interface {
	FindByID(context.Context, *IDRequest) (*DepartmentReply, error)
	All(context.Context, *PageRequest) (*DepartmentListReply, error)
	mustEmbedUnimplementedDepartmentServiceServer()
}

// UnimplementedDepartmentServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDepartmentServiceServer struct {
}

func (UnimplementedDepartmentServiceServer) FindByID(context.Context, *IDRequest) (*DepartmentReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindByID not implemented")
}
func (UnimplementedDepartmentServiceServer) All(context.Context, *PageRequest) (*DepartmentListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method All not implemented")
}
func (UnimplementedDepartmentServiceServer) mustEmbedUnimplementedDepartmentServiceServer() {}

// UnsafeDepartmentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DepartmentServiceServer will
// result in compilation errors.
type UnsafeDepartmentServiceServer interface {
	mustEmbedUnimplementedDepartmentServiceServer()
}

func RegisterDepartmentServiceServer(s grpc.ServiceRegistrar, srv DepartmentServiceServer) {
	s.RegisterService(&DepartmentService_ServiceDesc, srv)
}

func _DepartmentService_FindByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DepartmentServiceServer).FindByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/duat.DepartmentService/FindByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DepartmentServiceServer).FindByID(ctx, req.(*IDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DepartmentService_All_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DepartmentServiceServer).All(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/duat.DepartmentService/All",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DepartmentServiceServer).All(ctx, req.(*PageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DepartmentService_ServiceDesc is the grpc.ServiceDesc for DepartmentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DepartmentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "duat.DepartmentService",
	HandlerType: (*DepartmentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindByID",
			Handler:    _DepartmentService_FindByID_Handler,
		},
		{
			MethodName: "All",
			Handler:    _DepartmentService_All_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "duat.proto",
}

// DocumentServiceClient is the client API for DocumentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DocumentServiceClient interface {
	All(ctx context.Context, in *PageRequest, opts ...grpc.CallOption) (*DocumentListReply, error)
	FindByID(ctx context.Context, in *UUIDRequest, opts ...grpc.CallOption) (*DocumentReply, error)
	Store(ctx context.Context, in *DocumentRequest, opts ...grpc.CallOption) (*DocumentReply, error)
}

type documentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDocumentServiceClient(cc grpc.ClientConnInterface) DocumentServiceClient {
	return &documentServiceClient{cc}
}

func (c *documentServiceClient) All(ctx context.Context, in *PageRequest, opts ...grpc.CallOption) (*DocumentListReply, error) {
	out := new(DocumentListReply)
	err := c.cc.Invoke(ctx, "/duat.DocumentService/All", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *documentServiceClient) FindByID(ctx context.Context, in *UUIDRequest, opts ...grpc.CallOption) (*DocumentReply, error) {
	out := new(DocumentReply)
	err := c.cc.Invoke(ctx, "/duat.DocumentService/FindByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *documentServiceClient) Store(ctx context.Context, in *DocumentRequest, opts ...grpc.CallOption) (*DocumentReply, error) {
	out := new(DocumentReply)
	err := c.cc.Invoke(ctx, "/duat.DocumentService/Store", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DocumentServiceServer is the server API for DocumentService service.
// All implementations must embed UnimplementedDocumentServiceServer
// for forward compatibility
type DocumentServiceServer interface {
	All(context.Context, *PageRequest) (*DocumentListReply, error)
	FindByID(context.Context, *UUIDRequest) (*DocumentReply, error)
	Store(context.Context, *DocumentRequest) (*DocumentReply, error)
	mustEmbedUnimplementedDocumentServiceServer()
}

// UnimplementedDocumentServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDocumentServiceServer struct {
}

func (UnimplementedDocumentServiceServer) All(context.Context, *PageRequest) (*DocumentListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method All not implemented")
}
func (UnimplementedDocumentServiceServer) FindByID(context.Context, *UUIDRequest) (*DocumentReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindByID not implemented")
}
func (UnimplementedDocumentServiceServer) Store(context.Context, *DocumentRequest) (*DocumentReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Store not implemented")
}
func (UnimplementedDocumentServiceServer) mustEmbedUnimplementedDocumentServiceServer() {}

// UnsafeDocumentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DocumentServiceServer will
// result in compilation errors.
type UnsafeDocumentServiceServer interface {
	mustEmbedUnimplementedDocumentServiceServer()
}

func RegisterDocumentServiceServer(s grpc.ServiceRegistrar, srv DocumentServiceServer) {
	s.RegisterService(&DocumentService_ServiceDesc, srv)
}

func _DocumentService_All_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DocumentServiceServer).All(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/duat.DocumentService/All",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DocumentServiceServer).All(ctx, req.(*PageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DocumentService_FindByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UUIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DocumentServiceServer).FindByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/duat.DocumentService/FindByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DocumentServiceServer).FindByID(ctx, req.(*UUIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DocumentService_Store_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DocumentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DocumentServiceServer).Store(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/duat.DocumentService/Store",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DocumentServiceServer).Store(ctx, req.(*DocumentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DocumentService_ServiceDesc is the grpc.ServiceDesc for DocumentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DocumentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "duat.DocumentService",
	HandlerType: (*DocumentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "All",
			Handler:    _DocumentService_All_Handler,
		},
		{
			MethodName: "FindByID",
			Handler:    _DocumentService_FindByID_Handler,
		},
		{
			MethodName: "Store",
			Handler:    _DocumentService_Store_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "duat.proto",
}

// ImageServiceClient is the client API for ImageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ImageServiceClient interface {
	Store(ctx context.Context, opts ...grpc.CallOption) (ImageService_StoreClient, error)
	Upload(ctx context.Context, opts ...grpc.CallOption) (ImageService_UploadClient, error)
}

type imageServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewImageServiceClient(cc grpc.ClientConnInterface) ImageServiceClient {
	return &imageServiceClient{cc}
}

func (c *imageServiceClient) Store(ctx context.Context, opts ...grpc.CallOption) (ImageService_StoreClient, error) {
	stream, err := c.cc.NewStream(ctx, &ImageService_ServiceDesc.Streams[0], "/duat.ImageService/Store", opts...)
	if err != nil {
		return nil, err
	}
	x := &imageServiceStoreClient{stream}
	return x, nil
}

type ImageService_StoreClient interface {
	Send(*DocumentImageRequest) error
	CloseAndRecv() (*DocumentImageReply, error)
	grpc.ClientStream
}

type imageServiceStoreClient struct {
	grpc.ClientStream
}

func (x *imageServiceStoreClient) Send(m *DocumentImageRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *imageServiceStoreClient) CloseAndRecv() (*DocumentImageReply, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(DocumentImageReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *imageServiceClient) Upload(ctx context.Context, opts ...grpc.CallOption) (ImageService_UploadClient, error) {
	stream, err := c.cc.NewStream(ctx, &ImageService_ServiceDesc.Streams[1], "/duat.ImageService/Upload", opts...)
	if err != nil {
		return nil, err
	}
	x := &imageServiceUploadClient{stream}
	return x, nil
}

type ImageService_UploadClient interface {
	Send(*UploadRequest) error
	CloseAndRecv() (*UploadReply, error)
	grpc.ClientStream
}

type imageServiceUploadClient struct {
	grpc.ClientStream
}

func (x *imageServiceUploadClient) Send(m *UploadRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *imageServiceUploadClient) CloseAndRecv() (*UploadReply, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(UploadReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ImageServiceServer is the server API for ImageService service.
// All implementations must embed UnimplementedImageServiceServer
// for forward compatibility
type ImageServiceServer interface {
	Store(ImageService_StoreServer) error
	Upload(ImageService_UploadServer) error
	mustEmbedUnimplementedImageServiceServer()
}

// UnimplementedImageServiceServer must be embedded to have forward compatible implementations.
type UnimplementedImageServiceServer struct {
}

func (UnimplementedImageServiceServer) Store(ImageService_StoreServer) error {
	return status.Errorf(codes.Unimplemented, "method Store not implemented")
}
func (UnimplementedImageServiceServer) Upload(ImageService_UploadServer) error {
	return status.Errorf(codes.Unimplemented, "method Upload not implemented")
}
func (UnimplementedImageServiceServer) mustEmbedUnimplementedImageServiceServer() {}

// UnsafeImageServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ImageServiceServer will
// result in compilation errors.
type UnsafeImageServiceServer interface {
	mustEmbedUnimplementedImageServiceServer()
}

func RegisterImageServiceServer(s grpc.ServiceRegistrar, srv ImageServiceServer) {
	s.RegisterService(&ImageService_ServiceDesc, srv)
}

func _ImageService_Store_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ImageServiceServer).Store(&imageServiceStoreServer{stream})
}

type ImageService_StoreServer interface {
	SendAndClose(*DocumentImageReply) error
	Recv() (*DocumentImageRequest, error)
	grpc.ServerStream
}

type imageServiceStoreServer struct {
	grpc.ServerStream
}

func (x *imageServiceStoreServer) SendAndClose(m *DocumentImageReply) error {
	return x.ServerStream.SendMsg(m)
}

func (x *imageServiceStoreServer) Recv() (*DocumentImageRequest, error) {
	m := new(DocumentImageRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ImageService_Upload_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ImageServiceServer).Upload(&imageServiceUploadServer{stream})
}

type ImageService_UploadServer interface {
	SendAndClose(*UploadReply) error
	Recv() (*UploadRequest, error)
	grpc.ServerStream
}

type imageServiceUploadServer struct {
	grpc.ServerStream
}

func (x *imageServiceUploadServer) SendAndClose(m *UploadReply) error {
	return x.ServerStream.SendMsg(m)
}

func (x *imageServiceUploadServer) Recv() (*UploadRequest, error) {
	m := new(UploadRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ImageService_ServiceDesc is the grpc.ServiceDesc for ImageService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ImageService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "duat.ImageService",
	HandlerType: (*ImageServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Store",
			Handler:       _ImageService_Store_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Upload",
			Handler:       _ImageService_Upload_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "duat.proto",
}
