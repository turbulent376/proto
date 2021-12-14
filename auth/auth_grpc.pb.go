// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package auth

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

// AuthServiceClient is the client API for AuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthServiceClient interface {
	// AuthUserByEmail authorizations user by login/pass
	AuthUserByEmail(ctx context.Context, in *AuthRequest, opts ...grpc.CallOption) (*AuthResponse, error)
	// AuthUserByFirebase authorizations user by firebase token
	AuthUserByFirebase(ctx context.Context, in *OAuthRequest, opts ...grpc.CallOption) (*AuthResponse, error)
	// RefreshToken resfreshes user jwt token
	RefreshToken(ctx context.Context, in *RefreshTokenRequest, opts ...grpc.CallOption) (*TokenPairResponse, error)
	// CreateUser creates a new user
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*User, error)
	// DeleteUser deletes user
	DeleteUser(ctx context.Context, in *UserIdRequest, opts ...grpc.CallOption) (*EmptyResponse, error)
	// UpdateUser updates user info
	UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*User, error)
	// GetUserById gets user by id
	GetUserById(ctx context.Context, in *UserIdRequest, opts ...grpc.CallOption) (*User, error)
	// GetUserByToken gets user by jwt token
	GetSessionByToken(ctx context.Context, in *TokenRequest, opts ...grpc.CallOption) (*Session, error)
	// SaveUserFCMToken saves new fcm token
	SaveUserFCMToken(ctx context.Context, in *FCMTokenRequest, opts ...grpc.CallOption) (*EmptyResponse, error)
	// GetUserFCMTokens get user's fcm tokens
	GetUserFCMTokens(ctx context.Context, in *UserIdRequest, opts ...grpc.CallOption) (*FCMTokensResponse, error)
	// GetUserSessions get active user's sessions
	GetUserSessions(ctx context.Context, in *UserIdRequest, opts ...grpc.CallOption) (*UserSessionsResponse, error)
	// CloseSession kill user session by device
	CloseSession(ctx context.Context, in *CloseSessionRequest, opts ...grpc.CallOption) (*EmptyResponse, error)
	// SetSessionInfo sets info about session - device name and client version
	SetSessionInfo(ctx context.Context, in *SetSessionInfoRequest, opts ...grpc.CallOption) (*EmptyResponse, error)
}

type authServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthServiceClient(cc grpc.ClientConnInterface) AuthServiceClient {
	return &authServiceClient{cc}
}

func (c *authServiceClient) AuthUserByEmail(ctx context.Context, in *AuthRequest, opts ...grpc.CallOption) (*AuthResponse, error) {
	out := new(AuthResponse)
	err := c.cc.Invoke(ctx, "/auth.AuthService/AuthUserByEmail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) AuthUserByFirebase(ctx context.Context, in *OAuthRequest, opts ...grpc.CallOption) (*AuthResponse, error) {
	out := new(AuthResponse)
	err := c.cc.Invoke(ctx, "/auth.AuthService/AuthUserByFirebase", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) RefreshToken(ctx context.Context, in *RefreshTokenRequest, opts ...grpc.CallOption) (*TokenPairResponse, error) {
	out := new(TokenPairResponse)
	err := c.cc.Invoke(ctx, "/auth.AuthService/RefreshToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/auth.AuthService/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) DeleteUser(ctx context.Context, in *UserIdRequest, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, "/auth.AuthService/DeleteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/auth.AuthService/UpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) GetUserById(ctx context.Context, in *UserIdRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/auth.AuthService/GetUserById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) GetSessionByToken(ctx context.Context, in *TokenRequest, opts ...grpc.CallOption) (*Session, error) {
	out := new(Session)
	err := c.cc.Invoke(ctx, "/auth.AuthService/GetSessionByToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) SaveUserFCMToken(ctx context.Context, in *FCMTokenRequest, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, "/auth.AuthService/SaveUserFCMToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) GetUserFCMTokens(ctx context.Context, in *UserIdRequest, opts ...grpc.CallOption) (*FCMTokensResponse, error) {
	out := new(FCMTokensResponse)
	err := c.cc.Invoke(ctx, "/auth.AuthService/GetUserFCMTokens", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) GetUserSessions(ctx context.Context, in *UserIdRequest, opts ...grpc.CallOption) (*UserSessionsResponse, error) {
	out := new(UserSessionsResponse)
	err := c.cc.Invoke(ctx, "/auth.AuthService/GetUserSessions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) CloseSession(ctx context.Context, in *CloseSessionRequest, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, "/auth.AuthService/CloseSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) SetSessionInfo(ctx context.Context, in *SetSessionInfoRequest, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, "/auth.AuthService/SetSessionInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServiceServer is the server API for AuthService service.
// All implementations must embed UnimplementedAuthServiceServer
// for forward compatibility
type AuthServiceServer interface {
	// AuthUserByEmail authorizations user by login/pass
	AuthUserByEmail(context.Context, *AuthRequest) (*AuthResponse, error)
	// AuthUserByFirebase authorizations user by firebase token
	AuthUserByFirebase(context.Context, *OAuthRequest) (*AuthResponse, error)
	// RefreshToken resfreshes user jwt token
	RefreshToken(context.Context, *RefreshTokenRequest) (*TokenPairResponse, error)
	// CreateUser creates a new user
	CreateUser(context.Context, *CreateUserRequest) (*User, error)
	// DeleteUser deletes user
	DeleteUser(context.Context, *UserIdRequest) (*EmptyResponse, error)
	// UpdateUser updates user info
	UpdateUser(context.Context, *UpdateUserRequest) (*User, error)
	// GetUserById gets user by id
	GetUserById(context.Context, *UserIdRequest) (*User, error)
	// GetUserByToken gets user by jwt token
	GetSessionByToken(context.Context, *TokenRequest) (*Session, error)
	// SaveUserFCMToken saves new fcm token
	SaveUserFCMToken(context.Context, *FCMTokenRequest) (*EmptyResponse, error)
	// GetUserFCMTokens get user's fcm tokens
	GetUserFCMTokens(context.Context, *UserIdRequest) (*FCMTokensResponse, error)
	// GetUserSessions get active user's sessions
	GetUserSessions(context.Context, *UserIdRequest) (*UserSessionsResponse, error)
	// CloseSession kill user session by device
	CloseSession(context.Context, *CloseSessionRequest) (*EmptyResponse, error)
	// SetSessionInfo sets info about session - device name and client version
	SetSessionInfo(context.Context, *SetSessionInfoRequest) (*EmptyResponse, error)
	mustEmbedUnimplementedAuthServiceServer()
}

// UnimplementedAuthServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAuthServiceServer struct {
}

func (UnimplementedAuthServiceServer) AuthUserByEmail(context.Context, *AuthRequest) (*AuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthUserByEmail not implemented")
}
func (UnimplementedAuthServiceServer) AuthUserByFirebase(context.Context, *OAuthRequest) (*AuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthUserByFirebase not implemented")
}
func (UnimplementedAuthServiceServer) RefreshToken(context.Context, *RefreshTokenRequest) (*TokenPairResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshToken not implemented")
}
func (UnimplementedAuthServiceServer) CreateUser(context.Context, *CreateUserRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedAuthServiceServer) DeleteUser(context.Context, *UserIdRequest) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedAuthServiceServer) UpdateUser(context.Context, *UpdateUserRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedAuthServiceServer) GetUserById(context.Context, *UserIdRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserById not implemented")
}
func (UnimplementedAuthServiceServer) GetSessionByToken(context.Context, *TokenRequest) (*Session, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSessionByToken not implemented")
}
func (UnimplementedAuthServiceServer) SaveUserFCMToken(context.Context, *FCMTokenRequest) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveUserFCMToken not implemented")
}
func (UnimplementedAuthServiceServer) GetUserFCMTokens(context.Context, *UserIdRequest) (*FCMTokensResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserFCMTokens not implemented")
}
func (UnimplementedAuthServiceServer) GetUserSessions(context.Context, *UserIdRequest) (*UserSessionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserSessions not implemented")
}
func (UnimplementedAuthServiceServer) CloseSession(context.Context, *CloseSessionRequest) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CloseSession not implemented")
}
func (UnimplementedAuthServiceServer) SetSessionInfo(context.Context, *SetSessionInfoRequest) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetSessionInfo not implemented")
}
func (UnimplementedAuthServiceServer) mustEmbedUnimplementedAuthServiceServer() {}

// UnsafeAuthServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthServiceServer will
// result in compilation errors.
type UnsafeAuthServiceServer interface {
	mustEmbedUnimplementedAuthServiceServer()
}

func RegisterAuthServiceServer(s grpc.ServiceRegistrar, srv AuthServiceServer) {
	s.RegisterService(&AuthService_ServiceDesc, srv)
}

func _AuthService_AuthUserByEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).AuthUserByEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthService/AuthUserByEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).AuthUserByEmail(ctx, req.(*AuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_AuthUserByFirebase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OAuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).AuthUserByFirebase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthService/AuthUserByFirebase",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).AuthUserByFirebase(ctx, req.(*OAuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_RefreshToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefreshTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).RefreshToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthService/RefreshToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).RefreshToken(ctx, req.(*RefreshTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthService/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthService/DeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).DeleteUser(ctx, req.(*UserIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthService/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).UpdateUser(ctx, req.(*UpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_GetUserById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).GetUserById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthService/GetUserById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).GetUserById(ctx, req.(*UserIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_GetSessionByToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).GetSessionByToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthService/GetSessionByToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).GetSessionByToken(ctx, req.(*TokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_SaveUserFCMToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FCMTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).SaveUserFCMToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthService/SaveUserFCMToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).SaveUserFCMToken(ctx, req.(*FCMTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_GetUserFCMTokens_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).GetUserFCMTokens(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthService/GetUserFCMTokens",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).GetUserFCMTokens(ctx, req.(*UserIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_GetUserSessions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).GetUserSessions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthService/GetUserSessions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).GetUserSessions(ctx, req.(*UserIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_CloseSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloseSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).CloseSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthService/CloseSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).CloseSession(ctx, req.(*CloseSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_SetSessionInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetSessionInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).SetSessionInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthService/SetSessionInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).SetSessionInfo(ctx, req.(*SetSessionInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthService_ServiceDesc is the grpc.ServiceDesc for AuthService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth.AuthService",
	HandlerType: (*AuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AuthUserByEmail",
			Handler:    _AuthService_AuthUserByEmail_Handler,
		},
		{
			MethodName: "AuthUserByFirebase",
			Handler:    _AuthService_AuthUserByFirebase_Handler,
		},
		{
			MethodName: "RefreshToken",
			Handler:    _AuthService_RefreshToken_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _AuthService_CreateUser_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _AuthService_DeleteUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _AuthService_UpdateUser_Handler,
		},
		{
			MethodName: "GetUserById",
			Handler:    _AuthService_GetUserById_Handler,
		},
		{
			MethodName: "GetSessionByToken",
			Handler:    _AuthService_GetSessionByToken_Handler,
		},
		{
			MethodName: "SaveUserFCMToken",
			Handler:    _AuthService_SaveUserFCMToken_Handler,
		},
		{
			MethodName: "GetUserFCMTokens",
			Handler:    _AuthService_GetUserFCMTokens_Handler,
		},
		{
			MethodName: "GetUserSessions",
			Handler:    _AuthService_GetUserSessions_Handler,
		},
		{
			MethodName: "CloseSession",
			Handler:    _AuthService_CloseSession_Handler,
		},
		{
			MethodName: "SetSessionInfo",
			Handler:    _AuthService_SetSessionInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth/auth.proto",
}