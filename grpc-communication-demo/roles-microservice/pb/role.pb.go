// Code generated by protoc-gen-go. DO NOT EDIT.
// source: role.proto

package roles

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Requests
type EmptyRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EmptyRequest) Reset()         { *m = EmptyRequest{} }
func (m *EmptyRequest) String() string { return proto.CompactTextString(m) }
func (*EmptyRequest) ProtoMessage()    {}
func (*EmptyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_role_99df254f4f6c51ad, []int{0}
}
func (m *EmptyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmptyRequest.Unmarshal(m, b)
}
func (m *EmptyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmptyRequest.Marshal(b, m, deterministic)
}
func (dst *EmptyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmptyRequest.Merge(dst, src)
}
func (m *EmptyRequest) XXX_Size() int {
	return xxx_messageInfo_EmptyRequest.Size(m)
}
func (m *EmptyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EmptyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EmptyRequest proto.InternalMessageInfo

type GetUserRoleRequest struct {
	UserId               int32    `protobuf:"varint,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserRoleRequest) Reset()         { *m = GetUserRoleRequest{} }
func (m *GetUserRoleRequest) String() string { return proto.CompactTextString(m) }
func (*GetUserRoleRequest) ProtoMessage()    {}
func (*GetUserRoleRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_role_99df254f4f6c51ad, []int{1}
}
func (m *GetUserRoleRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserRoleRequest.Unmarshal(m, b)
}
func (m *GetUserRoleRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserRoleRequest.Marshal(b, m, deterministic)
}
func (dst *GetUserRoleRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserRoleRequest.Merge(dst, src)
}
func (m *GetUserRoleRequest) XXX_Size() int {
	return xxx_messageInfo_GetUserRoleRequest.Size(m)
}
func (m *GetUserRoleRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserRoleRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserRoleRequest proto.InternalMessageInfo

func (m *GetUserRoleRequest) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

// Replys
type RolesReply struct {
	Roles                []*Role  `protobuf:"bytes,1,rep,name=roles" json:"roles,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RolesReply) Reset()         { *m = RolesReply{} }
func (m *RolesReply) String() string { return proto.CompactTextString(m) }
func (*RolesReply) ProtoMessage()    {}
func (*RolesReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_role_99df254f4f6c51ad, []int{2}
}
func (m *RolesReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RolesReply.Unmarshal(m, b)
}
func (m *RolesReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RolesReply.Marshal(b, m, deterministic)
}
func (dst *RolesReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RolesReply.Merge(dst, src)
}
func (m *RolesReply) XXX_Size() int {
	return xxx_messageInfo_RolesReply.Size(m)
}
func (m *RolesReply) XXX_DiscardUnknown() {
	xxx_messageInfo_RolesReply.DiscardUnknown(m)
}

var xxx_messageInfo_RolesReply proto.InternalMessageInfo

func (m *RolesReply) GetRoles() []*Role {
	if m != nil {
		return m.Roles
	}
	return nil
}

type Role struct {
	Id                   int32    `protobuf:"varint,2,opt,name=id" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Role) Reset()         { *m = Role{} }
func (m *Role) String() string { return proto.CompactTextString(m) }
func (*Role) ProtoMessage()    {}
func (*Role) Descriptor() ([]byte, []int) {
	return fileDescriptor_role_99df254f4f6c51ad, []int{3}
}
func (m *Role) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Role.Unmarshal(m, b)
}
func (m *Role) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Role.Marshal(b, m, deterministic)
}
func (dst *Role) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Role.Merge(dst, src)
}
func (m *Role) XXX_Size() int {
	return xxx_messageInfo_Role.Size(m)
}
func (m *Role) XXX_DiscardUnknown() {
	xxx_messageInfo_Role.DiscardUnknown(m)
}

var xxx_messageInfo_Role proto.InternalMessageInfo

func (m *Role) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Role) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type UserRoleReply struct {
	UserId               int32    `protobuf:"varint,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	Roles                []*Role  `protobuf:"bytes,2,rep,name=roles" json:"roles,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserRoleReply) Reset()         { *m = UserRoleReply{} }
func (m *UserRoleReply) String() string { return proto.CompactTextString(m) }
func (*UserRoleReply) ProtoMessage()    {}
func (*UserRoleReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_role_99df254f4f6c51ad, []int{4}
}
func (m *UserRoleReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserRoleReply.Unmarshal(m, b)
}
func (m *UserRoleReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserRoleReply.Marshal(b, m, deterministic)
}
func (dst *UserRoleReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserRoleReply.Merge(dst, src)
}
func (m *UserRoleReply) XXX_Size() int {
	return xxx_messageInfo_UserRoleReply.Size(m)
}
func (m *UserRoleReply) XXX_DiscardUnknown() {
	xxx_messageInfo_UserRoleReply.DiscardUnknown(m)
}

var xxx_messageInfo_UserRoleReply proto.InternalMessageInfo

func (m *UserRoleReply) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *UserRoleReply) GetRoles() []*Role {
	if m != nil {
		return m.Roles
	}
	return nil
}

func init() {
	proto.RegisterType((*EmptyRequest)(nil), "roles.EmptyRequest")
	proto.RegisterType((*GetUserRoleRequest)(nil), "roles.GetUserRoleRequest")
	proto.RegisterType((*RolesReply)(nil), "roles.RolesReply")
	proto.RegisterType((*Role)(nil), "roles.Role")
	proto.RegisterType((*UserRoleReply)(nil), "roles.UserRoleReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RolesClient is the client API for Roles service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RolesClient interface {
	GetRoles(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*RolesReply, error)
	GetUserRole(ctx context.Context, in *GetUserRoleRequest, opts ...grpc.CallOption) (*UserRoleReply, error)
}

type rolesClient struct {
	cc *grpc.ClientConn
}

func NewRolesClient(cc *grpc.ClientConn) RolesClient {
	return &rolesClient{cc}
}

func (c *rolesClient) GetRoles(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*RolesReply, error) {
	out := new(RolesReply)
	err := c.cc.Invoke(ctx, "/roles.Roles/GetRoles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rolesClient) GetUserRole(ctx context.Context, in *GetUserRoleRequest, opts ...grpc.CallOption) (*UserRoleReply, error) {
	out := new(UserRoleReply)
	err := c.cc.Invoke(ctx, "/roles.Roles/GetUserRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RolesServer is the server API for Roles service.
type RolesServer interface {
	GetRoles(context.Context, *EmptyRequest) (*RolesReply, error)
	GetUserRole(context.Context, *GetUserRoleRequest) (*UserRoleReply, error)
}

func RegisterRolesServer(s *grpc.Server, srv RolesServer) {
	s.RegisterService(&_Roles_serviceDesc, srv)
}

func _Roles_GetRoles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RolesServer).GetRoles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/roles.Roles/GetRoles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RolesServer).GetRoles(ctx, req.(*EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Roles_GetUserRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RolesServer).GetUserRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/roles.Roles/GetUserRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RolesServer).GetUserRole(ctx, req.(*GetUserRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Roles_serviceDesc = grpc.ServiceDesc{
	ServiceName: "roles.Roles",
	HandlerType: (*RolesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRoles",
			Handler:    _Roles_GetRoles_Handler,
		},
		{
			MethodName: "GetUserRole",
			Handler:    _Roles_GetUserRole_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "role.proto",
}

func init() { proto.RegisterFile("role.proto", fileDescriptor_role_99df254f4f6c51ad) }

var fileDescriptor_role_99df254f4f6c51ad = []byte{
	// 231 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x50, 0xc1, 0x4a, 0xc4, 0x30,
	0x14, 0xdc, 0xc6, 0xed, 0xaa, 0xaf, 0xba, 0xe0, 0x53, 0xb0, 0xf6, 0x54, 0x73, 0x2a, 0x82, 0x15,
	0xaa, 0x1f, 0xe0, 0x45, 0x8a, 0x78, 0x0b, 0x78, 0x16, 0x25, 0xef, 0x50, 0x68, 0x4d, 0x4d, 0xd2,
	0x43, 0x4f, 0xfe, 0xba, 0x24, 0x0d, 0xb4, 0xb2, 0xf4, 0x36, 0x93, 0x4c, 0x66, 0x26, 0x03, 0xa0,
	0x55, 0x4b, 0x65, 0xaf, 0x95, 0x55, 0x18, 0x3b, 0x6c, 0xf8, 0x1e, 0xce, 0x5e, 0xba, 0xde, 0x8e,
	0x82, 0x7e, 0x06, 0x32, 0x96, 0xdf, 0x03, 0xd6, 0x64, 0xdf, 0x0d, 0x69, 0xa1, 0x5a, 0x0a, 0xa7,
	0x78, 0x0d, 0xc7, 0x83, 0x21, 0xfd, 0xd1, 0xc8, 0x34, 0xca, 0xa3, 0x22, 0x16, 0x3b, 0x47, 0x5f,
	0x25, 0x7f, 0x00, 0x70, 0x3a, 0x23, 0xa8, 0x6f, 0x47, 0xbc, 0x85, 0xc9, 0x35, 0x8d, 0xf2, 0xa3,
	0x22, 0xa9, 0x92, 0xd2, 0xb3, 0xd2, 0x3b, 0x85, 0xbc, 0x3b, 0xd8, 0x3a, 0x8a, 0x7b, 0x60, 0x8d,
	0x4c, 0x99, 0x37, 0x63, 0x8d, 0x44, 0x84, 0xed, 0xf7, 0x67, 0x47, 0xde, 0xfe, 0x54, 0x78, 0xcc,
	0xdf, 0xe0, 0x7c, 0x2e, 0xe2, 0xfc, 0xd7, 0x6a, 0xcc, 0xc1, 0x6c, 0x2d, 0xb8, 0xfa, 0x85, 0xd8,
	0x37, 0xc5, 0x27, 0x38, 0xa9, 0xc9, 0x4e, 0xf8, 0x32, 0x08, 0x97, 0x13, 0x64, 0x17, 0x8b, 0xd7,
	0xd3, 0xc7, 0xf8, 0x06, 0x9f, 0x21, 0x59, 0xec, 0x82, 0x37, 0x41, 0x73, 0xb8, 0x55, 0x76, 0x15,
	0xae, 0xfe, 0x55, 0xe7, 0x9b, 0xaf, 0x9d, 0xdf, 0xfd, 0xf1, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x95,
	0x61, 0x8d, 0x7b, 0x85, 0x01, 0x00, 0x00,
}
