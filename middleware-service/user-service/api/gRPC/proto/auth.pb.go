// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: auth.proto

package proto

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type AuthToken struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token   string               `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Success bool                 `protobuf:"varint,2,opt,name=success,proto3" json:"success,omitempty"`
	Expires *timestamp.Timestamp `protobuf:"bytes,3,opt,name=expires,proto3" json:"expires,omitempty"`
}

func (x *AuthToken) Reset() {
	*x = AuthToken{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthToken) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthToken) ProtoMessage() {}

func (x *AuthToken) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthToken.ProtoReflect.Descriptor instead.
func (*AuthToken) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{0}
}

func (x *AuthToken) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *AuthToken) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *AuthToken) GetExpires() *timestamp.Timestamp {
	if x != nil {
		return x.Expires
	}
	return nil
}

var File_auth_proto protoreflect.FileDescriptor

var file_auth_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x71, 0x0a,
	0x09, 0x41, 0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x34, 0x0a, 0x07, 0x65, 0x78,
	0x70, 0x69, 0x72, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73,
	0x32, 0xc7, 0x02, 0x0a, 0x0b, 0x41, 0x75, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x27, 0x0a, 0x06, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x12, 0x0b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x41, 0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x26, 0x0a, 0x05, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x12, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x1a,
	0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x12, 0x27, 0x0a, 0x06, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x12, 0x0b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x41, 0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x31, 0x0a, 0x05, 0x47, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x10, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x2e, 0x0a,
	0x0d, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x0b,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x10, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x2d, 0x0a,
	0x07, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x12, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x41, 0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x1a, 0x10, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x2c, 0x0a, 0x06,
	0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x12, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41,
	0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x1a, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x41, 0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x11, 0x5a, 0x07, 0x2e, 0x3b,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x05, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_auth_proto_rawDescOnce sync.Once
	file_auth_proto_rawDescData = file_auth_proto_rawDesc
)

func file_auth_proto_rawDescGZIP() []byte {
	file_auth_proto_rawDescOnce.Do(func() {
		file_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_auth_proto_rawDescData)
	})
	return file_auth_proto_rawDescData
}

var file_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_auth_proto_goTypes = []interface{}{
	(*AuthToken)(nil),           // 0: proto.AuthToken
	(*timestamp.Timestamp)(nil), // 1: google.protobuf.Timestamp
	(*User)(nil),                // 2: proto.User
	(*empty.Empty)(nil),         // 3: google.protobuf.Empty
}
var file_auth_proto_depIdxs = []int32{
	1, // 0: proto.AuthToken.expires:type_name -> google.protobuf.Timestamp
	2, // 1: proto.AuthService.SignUp:input_type -> proto.User
	2, // 2: proto.AuthService.Login:input_type -> proto.User
	2, // 3: proto.AuthService.Remote:input_type -> proto.User
	3, // 4: proto.AuthService.Guest:input_type -> google.protobuf.Empty
	2, // 5: proto.AuthService.GenerateToken:input_type -> proto.User
	0, // 6: proto.AuthService.Refresh:input_type -> proto.AuthToken
	0, // 7: proto.AuthService.Logout:input_type -> proto.AuthToken
	0, // 8: proto.AuthService.SignUp:output_type -> proto.AuthToken
	0, // 9: proto.AuthService.Login:output_type -> proto.AuthToken
	0, // 10: proto.AuthService.Remote:output_type -> proto.AuthToken
	0, // 11: proto.AuthService.Guest:output_type -> proto.AuthToken
	0, // 12: proto.AuthService.GenerateToken:output_type -> proto.AuthToken
	0, // 13: proto.AuthService.Refresh:output_type -> proto.AuthToken
	0, // 14: proto.AuthService.Logout:output_type -> proto.AuthToken
	8, // [8:15] is the sub-list for method output_type
	1, // [1:8] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_auth_proto_init() }
func file_auth_proto_init() {
	if File_auth_proto != nil {
		return
	}
	file_user_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_auth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthToken); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_auth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_auth_proto_goTypes,
		DependencyIndexes: file_auth_proto_depIdxs,
		MessageInfos:      file_auth_proto_msgTypes,
	}.Build()
	File_auth_proto = out.File
	file_auth_proto_rawDesc = nil
	file_auth_proto_goTypes = nil
	file_auth_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AuthServiceClient is the client API for AuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthServiceClient interface {
	SignUp(ctx context.Context, in *User, opts ...grpc.CallOption) (*AuthToken, error)
	Login(ctx context.Context, in *User, opts ...grpc.CallOption) (*AuthToken, error)
	Remote(ctx context.Context, in *User, opts ...grpc.CallOption) (*AuthToken, error)
	Guest(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*AuthToken, error)
	GenerateToken(ctx context.Context, in *User, opts ...grpc.CallOption) (*AuthToken, error)
	Refresh(ctx context.Context, in *AuthToken, opts ...grpc.CallOption) (*AuthToken, error)
	Logout(ctx context.Context, in *AuthToken, opts ...grpc.CallOption) (*AuthToken, error)
}

type authServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthServiceClient(cc grpc.ClientConnInterface) AuthServiceClient {
	return &authServiceClient{cc}
}

func (c *authServiceClient) SignUp(ctx context.Context, in *User, opts ...grpc.CallOption) (*AuthToken, error) {
	out := new(AuthToken)
	err := c.cc.Invoke(ctx, "/proto.AuthService/SignUp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) Login(ctx context.Context, in *User, opts ...grpc.CallOption) (*AuthToken, error) {
	out := new(AuthToken)
	err := c.cc.Invoke(ctx, "/proto.AuthService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) Remote(ctx context.Context, in *User, opts ...grpc.CallOption) (*AuthToken, error) {
	out := new(AuthToken)
	err := c.cc.Invoke(ctx, "/proto.AuthService/Remote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) Guest(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*AuthToken, error) {
	out := new(AuthToken)
	err := c.cc.Invoke(ctx, "/proto.AuthService/Guest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) GenerateToken(ctx context.Context, in *User, opts ...grpc.CallOption) (*AuthToken, error) {
	out := new(AuthToken)
	err := c.cc.Invoke(ctx, "/proto.AuthService/GenerateToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) Refresh(ctx context.Context, in *AuthToken, opts ...grpc.CallOption) (*AuthToken, error) {
	out := new(AuthToken)
	err := c.cc.Invoke(ctx, "/proto.AuthService/Refresh", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) Logout(ctx context.Context, in *AuthToken, opts ...grpc.CallOption) (*AuthToken, error) {
	out := new(AuthToken)
	err := c.cc.Invoke(ctx, "/proto.AuthService/Logout", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServiceServer is the server API for AuthService service.
type AuthServiceServer interface {
	SignUp(context.Context, *User) (*AuthToken, error)
	Login(context.Context, *User) (*AuthToken, error)
	Remote(context.Context, *User) (*AuthToken, error)
	Guest(context.Context, *empty.Empty) (*AuthToken, error)
	GenerateToken(context.Context, *User) (*AuthToken, error)
	Refresh(context.Context, *AuthToken) (*AuthToken, error)
	Logout(context.Context, *AuthToken) (*AuthToken, error)
}

// UnimplementedAuthServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAuthServiceServer struct {
}

func (*UnimplementedAuthServiceServer) SignUp(context.Context, *User) (*AuthToken, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignUp not implemented")
}
func (*UnimplementedAuthServiceServer) Login(context.Context, *User) (*AuthToken, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (*UnimplementedAuthServiceServer) Remote(context.Context, *User) (*AuthToken, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Remote not implemented")
}
func (*UnimplementedAuthServiceServer) Guest(context.Context, *empty.Empty) (*AuthToken, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Guest not implemented")
}
func (*UnimplementedAuthServiceServer) GenerateToken(context.Context, *User) (*AuthToken, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateToken not implemented")
}
func (*UnimplementedAuthServiceServer) Refresh(context.Context, *AuthToken) (*AuthToken, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Refresh not implemented")
}
func (*UnimplementedAuthServiceServer) Logout(context.Context, *AuthToken) (*AuthToken, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Logout not implemented")
}

func RegisterAuthServiceServer(s *grpc.Server, srv AuthServiceServer) {
	s.RegisterService(&_AuthService_serviceDesc, srv)
}

func _AuthService_SignUp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).SignUp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.AuthService/SignUp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).SignUp(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.AuthService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Login(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_Remote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Remote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.AuthService/Remote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Remote(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_Guest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Guest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.AuthService/Guest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Guest(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_GenerateToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).GenerateToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.AuthService/GenerateToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).GenerateToken(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_Refresh_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthToken)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Refresh(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.AuthService/Refresh",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Refresh(ctx, req.(*AuthToken))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthToken)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.AuthService/Logout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Logout(ctx, req.(*AuthToken))
	}
	return interceptor(ctx, in, info, handler)
}

var _AuthService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.AuthService",
	HandlerType: (*AuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SignUp",
			Handler:    _AuthService_SignUp_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _AuthService_Login_Handler,
		},
		{
			MethodName: "Remote",
			Handler:    _AuthService_Remote_Handler,
		},
		{
			MethodName: "Guest",
			Handler:    _AuthService_Guest_Handler,
		},
		{
			MethodName: "GenerateToken",
			Handler:    _AuthService_GenerateToken_Handler,
		},
		{
			MethodName: "Refresh",
			Handler:    _AuthService_Refresh_Handler,
		},
		{
			MethodName: "Logout",
			Handler:    _AuthService_Logout_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth.proto",
}
