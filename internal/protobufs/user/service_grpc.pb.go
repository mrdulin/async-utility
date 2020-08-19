// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package user

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	GetUserByLoginname(ctx context.Context, in *GetUserByLoginnameRequest, opts ...grpc.CallOption) (*UserDetail, error)
	ValidateAccessToken(ctx context.Context, in *ValidateAccessTokenRequest, opts ...grpc.CallOption) (*UserEntity, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) GetUserByLoginname(ctx context.Context, in *GetUserByLoginnameRequest, opts ...grpc.CallOption) (*UserDetail, error) {
	out := new(UserDetail)
	err := c.cc.Invoke(ctx, "/user.UserService/GetUserByLoginname", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ValidateAccessToken(ctx context.Context, in *ValidateAccessTokenRequest, opts ...grpc.CallOption) (*UserEntity, error) {
	out := new(UserEntity)
	err := c.cc.Invoke(ctx, "/user.UserService/ValidateAccessToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	GetUserByLoginname(context.Context, *GetUserByLoginnameRequest) (*UserDetail, error)
	ValidateAccessToken(context.Context, *ValidateAccessTokenRequest) (*UserEntity, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (*UnimplementedUserServiceServer) GetUserByLoginname(context.Context, *GetUserByLoginnameRequest) (*UserDetail, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserByLoginname not implemented")
}
func (*UnimplementedUserServiceServer) ValidateAccessToken(context.Context, *ValidateAccessTokenRequest) (*UserEntity, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateAccessToken not implemented")
}
func (*UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_GetUserByLoginname_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserByLoginnameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserByLoginname(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/GetUserByLoginname",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserByLoginname(ctx, req.(*GetUserByLoginnameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ValidateAccessToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateAccessTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ValidateAccessToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/ValidateAccessToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ValidateAccessToken(ctx, req.(*ValidateAccessTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "user.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserByLoginname",
			Handler:    _UserService_GetUserByLoginname_Handler,
		},
		{
			MethodName: "ValidateAccessToken",
			Handler:    _UserService_ValidateAccessToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/protobufs/user/service.proto",
}