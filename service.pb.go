// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

package totp

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"

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

type IssueKeyRequest struct {
	Issuer               string   `protobuf:"bytes,1,opt,name=issuer" json:"issuer,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username" json:"username,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IssueKeyRequest) Reset()         { *m = IssueKeyRequest{} }
func (m *IssueKeyRequest) String() string { return proto.CompactTextString(m) }
func (*IssueKeyRequest) ProtoMessage()    {}
func (*IssueKeyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_e284110ba5c8e031, []int{0}
}
func (m *IssueKeyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IssueKeyRequest.Unmarshal(m, b)
}
func (m *IssueKeyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IssueKeyRequest.Marshal(b, m, deterministic)
}
func (dst *IssueKeyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IssueKeyRequest.Merge(dst, src)
}
func (m *IssueKeyRequest) XXX_Size() int {
	return xxx_messageInfo_IssueKeyRequest.Size(m)
}
func (m *IssueKeyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_IssueKeyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_IssueKeyRequest proto.InternalMessageInfo

func (m *IssueKeyRequest) GetIssuer() string {
	if m != nil {
		return m.Issuer
	}
	return ""
}

func (m *IssueKeyRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type IssueKeyReply struct {
	Url                  string   `protobuf:"bytes,1,opt,name=url" json:"url,omitempty"`
	Recover              string   `protobuf:"bytes,2,opt,name=recover" json:"recover,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IssueKeyReply) Reset()         { *m = IssueKeyReply{} }
func (m *IssueKeyReply) String() string { return proto.CompactTextString(m) }
func (*IssueKeyReply) ProtoMessage()    {}
func (*IssueKeyReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_e284110ba5c8e031, []int{1}
}
func (m *IssueKeyReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IssueKeyReply.Unmarshal(m, b)
}
func (m *IssueKeyReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IssueKeyReply.Marshal(b, m, deterministic)
}
func (dst *IssueKeyReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IssueKeyReply.Merge(dst, src)
}
func (m *IssueKeyReply) XXX_Size() int {
	return xxx_messageInfo_IssueKeyReply.Size(m)
}
func (m *IssueKeyReply) XXX_DiscardUnknown() {
	xxx_messageInfo_IssueKeyReply.DiscardUnknown(m)
}

var xxx_messageInfo_IssueKeyReply proto.InternalMessageInfo

func (m *IssueKeyReply) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *IssueKeyReply) GetRecover() string {
	if m != nil {
		return m.Recover
	}
	return ""
}

type ValidateRequest struct {
	Issuer               string               `protobuf:"bytes,1,opt,name=issuer" json:"issuer,omitempty"`
	Username             string               `protobuf:"bytes,2,opt,name=username" json:"username,omitempty"`
	Passcode             string               `protobuf:"bytes,3,opt,name=passcode" json:"passcode,omitempty"`
	Time                 *timestamp.Timestamp `protobuf:"bytes,4,opt,name=time" json:"time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ValidateRequest) Reset()         { *m = ValidateRequest{} }
func (m *ValidateRequest) String() string { return proto.CompactTextString(m) }
func (*ValidateRequest) ProtoMessage()    {}
func (*ValidateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_e284110ba5c8e031, []int{2}
}
func (m *ValidateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ValidateRequest.Unmarshal(m, b)
}
func (m *ValidateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ValidateRequest.Marshal(b, m, deterministic)
}
func (dst *ValidateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidateRequest.Merge(dst, src)
}
func (m *ValidateRequest) XXX_Size() int {
	return xxx_messageInfo_ValidateRequest.Size(m)
}
func (m *ValidateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ValidateRequest proto.InternalMessageInfo

func (m *ValidateRequest) GetIssuer() string {
	if m != nil {
		return m.Issuer
	}
	return ""
}

func (m *ValidateRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *ValidateRequest) GetPasscode() string {
	if m != nil {
		return m.Passcode
	}
	return ""
}

func (m *ValidateRequest) GetTime() *timestamp.Timestamp {
	if m != nil {
		return m.Time
	}
	return nil
}

type ValidateReply struct {
	Validate             bool     `protobuf:"varint,1,opt,name=validate" json:"validate,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ValidateReply) Reset()         { *m = ValidateReply{} }
func (m *ValidateReply) String() string { return proto.CompactTextString(m) }
func (*ValidateReply) ProtoMessage()    {}
func (*ValidateReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_e284110ba5c8e031, []int{3}
}
func (m *ValidateReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ValidateReply.Unmarshal(m, b)
}
func (m *ValidateReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ValidateReply.Marshal(b, m, deterministic)
}
func (dst *ValidateReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidateReply.Merge(dst, src)
}
func (m *ValidateReply) XXX_Size() int {
	return xxx_messageInfo_ValidateReply.Size(m)
}
func (m *ValidateReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidateReply.DiscardUnknown(m)
}

var xxx_messageInfo_ValidateReply proto.InternalMessageInfo

func (m *ValidateReply) GetValidate() bool {
	if m != nil {
		return m.Validate
	}
	return false
}

type RecoverRequest struct {
	Issuer               string   `protobuf:"bytes,1,opt,name=issuer" json:"issuer,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username" json:"username,omitempty"`
	Recover              string   `protobuf:"bytes,3,opt,name=recover" json:"recover,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RecoverRequest) Reset()         { *m = RecoverRequest{} }
func (m *RecoverRequest) String() string { return proto.CompactTextString(m) }
func (*RecoverRequest) ProtoMessage()    {}
func (*RecoverRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_e284110ba5c8e031, []int{4}
}
func (m *RecoverRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RecoverRequest.Unmarshal(m, b)
}
func (m *RecoverRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RecoverRequest.Marshal(b, m, deterministic)
}
func (dst *RecoverRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RecoverRequest.Merge(dst, src)
}
func (m *RecoverRequest) XXX_Size() int {
	return xxx_messageInfo_RecoverRequest.Size(m)
}
func (m *RecoverRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RecoverRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RecoverRequest proto.InternalMessageInfo

func (m *RecoverRequest) GetIssuer() string {
	if m != nil {
		return m.Issuer
	}
	return ""
}

func (m *RecoverRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *RecoverRequest) GetRecover() string {
	if m != nil {
		return m.Recover
	}
	return ""
}

type RecoverReply struct {
	Validate             bool     `protobuf:"varint,1,opt,name=validate" json:"validate,omitempty"`
	Url                  string   `protobuf:"bytes,2,opt,name=url" json:"url,omitempty"`
	Recover              string   `protobuf:"bytes,3,opt,name=recover" json:"recover,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RecoverReply) Reset()         { *m = RecoverReply{} }
func (m *RecoverReply) String() string { return proto.CompactTextString(m) }
func (*RecoverReply) ProtoMessage()    {}
func (*RecoverReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_e284110ba5c8e031, []int{5}
}
func (m *RecoverReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RecoverReply.Unmarshal(m, b)
}
func (m *RecoverReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RecoverReply.Marshal(b, m, deterministic)
}
func (dst *RecoverReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RecoverReply.Merge(dst, src)
}
func (m *RecoverReply) XXX_Size() int {
	return xxx_messageInfo_RecoverReply.Size(m)
}
func (m *RecoverReply) XXX_DiscardUnknown() {
	xxx_messageInfo_RecoverReply.DiscardUnknown(m)
}

var xxx_messageInfo_RecoverReply proto.InternalMessageInfo

func (m *RecoverReply) GetValidate() bool {
	if m != nil {
		return m.Validate
	}
	return false
}

func (m *RecoverReply) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *RecoverReply) GetRecover() string {
	if m != nil {
		return m.Recover
	}
	return ""
}

type RemoveKeyRequest struct {
	Issuer               string   `protobuf:"bytes,1,opt,name=issuer" json:"issuer,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username" json:"username,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoveKeyRequest) Reset()         { *m = RemoveKeyRequest{} }
func (m *RemoveKeyRequest) String() string { return proto.CompactTextString(m) }
func (*RemoveKeyRequest) ProtoMessage()    {}
func (*RemoveKeyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_e284110ba5c8e031, []int{6}
}
func (m *RemoveKeyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveKeyRequest.Unmarshal(m, b)
}
func (m *RemoveKeyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveKeyRequest.Marshal(b, m, deterministic)
}
func (dst *RemoveKeyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveKeyRequest.Merge(dst, src)
}
func (m *RemoveKeyRequest) XXX_Size() int {
	return xxx_messageInfo_RemoveKeyRequest.Size(m)
}
func (m *RemoveKeyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveKeyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveKeyRequest proto.InternalMessageInfo

func (m *RemoveKeyRequest) GetIssuer() string {
	if m != nil {
		return m.Issuer
	}
	return ""
}

func (m *RemoveKeyRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type RemoveKeyReply struct {
	Removed              bool     `protobuf:"varint,1,opt,name=removed" json:"removed,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoveKeyReply) Reset()         { *m = RemoveKeyReply{} }
func (m *RemoveKeyReply) String() string { return proto.CompactTextString(m) }
func (*RemoveKeyReply) ProtoMessage()    {}
func (*RemoveKeyReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_e284110ba5c8e031, []int{7}
}
func (m *RemoveKeyReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveKeyReply.Unmarshal(m, b)
}
func (m *RemoveKeyReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveKeyReply.Marshal(b, m, deterministic)
}
func (dst *RemoveKeyReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveKeyReply.Merge(dst, src)
}
func (m *RemoveKeyReply) XXX_Size() int {
	return xxx_messageInfo_RemoveKeyReply.Size(m)
}
func (m *RemoveKeyReply) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveKeyReply.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveKeyReply proto.InternalMessageInfo

func (m *RemoveKeyReply) GetRemoved() bool {
	if m != nil {
		return m.Removed
	}
	return false
}

func init() {
	proto.RegisterType((*IssueKeyRequest)(nil), "totp.IssueKeyRequest")
	proto.RegisterType((*IssueKeyReply)(nil), "totp.IssueKeyReply")
	proto.RegisterType((*ValidateRequest)(nil), "totp.ValidateRequest")
	proto.RegisterType((*ValidateReply)(nil), "totp.ValidateReply")
	proto.RegisterType((*RecoverRequest)(nil), "totp.RecoverRequest")
	proto.RegisterType((*RecoverReply)(nil), "totp.RecoverReply")
	proto.RegisterType((*RemoveKeyRequest)(nil), "totp.RemoveKeyRequest")
	proto.RegisterType((*RemoveKeyReply)(nil), "totp.RemoveKeyReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Totp service

type TotpClient interface {
	IssueKey(ctx context.Context, in *IssueKeyRequest, opts ...grpc.CallOption) (*IssueKeyReply, error)
	Validate(ctx context.Context, in *ValidateRequest, opts ...grpc.CallOption) (*ValidateReply, error)
	RecoverKey(ctx context.Context, in *RecoverRequest, opts ...grpc.CallOption) (*RecoverReply, error)
	RemoveKey(ctx context.Context, in *RemoveKeyRequest, opts ...grpc.CallOption) (*RemoveKeyReply, error)
}

type totpClient struct {
	cc *grpc.ClientConn
}

func NewTotpClient(cc *grpc.ClientConn) TotpClient {
	return &totpClient{cc}
}

func (c *totpClient) IssueKey(ctx context.Context, in *IssueKeyRequest, opts ...grpc.CallOption) (*IssueKeyReply, error) {
	out := new(IssueKeyReply)
	err := c.cc.Invoke(ctx, "/totp.Totp/issueKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *totpClient) Validate(ctx context.Context, in *ValidateRequest, opts ...grpc.CallOption) (*ValidateReply, error) {
	out := new(ValidateReply)
	err := c.cc.Invoke(ctx, "/totp.Totp/validate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *totpClient) RecoverKey(ctx context.Context, in *RecoverRequest, opts ...grpc.CallOption) (*RecoverReply, error) {
	out := new(RecoverReply)
	err := c.cc.Invoke(ctx, "/totp.Totp/recoverKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *totpClient) RemoveKey(ctx context.Context, in *RemoveKeyRequest, opts ...grpc.CallOption) (*RemoveKeyReply, error) {
	out := new(RemoveKeyReply)
	err := c.cc.Invoke(ctx, "/totp.Totp/removeKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Totp service

type TotpServer interface {
	IssueKey(context.Context, *IssueKeyRequest) (*IssueKeyReply, error)
	Validate(context.Context, *ValidateRequest) (*ValidateReply, error)
	RecoverKey(context.Context, *RecoverRequest) (*RecoverReply, error)
	RemoveKey(context.Context, *RemoveKeyRequest) (*RemoveKeyReply, error)
}

func RegisterTotpServer(s *grpc.Server, srv TotpServer) {
	s.RegisterService(&_Totp_serviceDesc, srv)
}

func _Totp_IssueKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IssueKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TotpServer).IssueKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/totp.Totp/IssueKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TotpServer).IssueKey(ctx, req.(*IssueKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Totp_Validate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TotpServer).Validate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/totp.Totp/Validate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TotpServer).Validate(ctx, req.(*ValidateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Totp_RecoverKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RecoverRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TotpServer).RecoverKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/totp.Totp/RecoverKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TotpServer).RecoverKey(ctx, req.(*RecoverRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Totp_RemoveKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TotpServer).RemoveKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/totp.Totp/RemoveKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TotpServer).RemoveKey(ctx, req.(*RemoveKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Totp_serviceDesc = grpc.ServiceDesc{
	ServiceName: "totp.Totp",
	HandlerType: (*TotpServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "issueKey",
			Handler:    _Totp_IssueKey_Handler,
		},
		{
			MethodName: "validate",
			Handler:    _Totp_Validate_Handler,
		},
		{
			MethodName: "recoverKey",
			Handler:    _Totp_RecoverKey_Handler,
		},
		{
			MethodName: "removeKey",
			Handler:    _Totp_RemoveKey_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}

func init() { proto.RegisterFile("service.proto", fileDescriptor_service_e284110ba5c8e031) }

var fileDescriptor_service_e284110ba5c8e031 = []byte{
	// 400 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x52, 0xc1, 0x6a, 0xdb, 0x40,
	0x10, 0xad, 0x6c, 0xe1, 0xca, 0xd3, 0xba, 0x36, 0x5b, 0xdb, 0x08, 0x5d, 0x6a, 0x44, 0x0f, 0xa6,
	0x85, 0x35, 0xb8, 0x97, 0x82, 0x6f, 0x86, 0x04, 0x42, 0x2e, 0x46, 0x18, 0x1f, 0x03, 0xb2, 0xb4,
	0x71, 0x04, 0x52, 0x56, 0xd9, 0x5d, 0x09, 0xf4, 0x19, 0xf9, 0xd4, 0xfc, 0x41, 0xd8, 0xd5, 0xae,
	0x1c, 0x29, 0x81, 0x1c, 0x7c, 0xd3, 0xbc, 0x9d, 0x37, 0x7a, 0xef, 0xcd, 0xc0, 0x88, 0x13, 0x56,
	0x26, 0x11, 0xc1, 0x39, 0xa3, 0x82, 0x22, 0x5b, 0x50, 0x91, 0x7b, 0xbf, 0x4e, 0x94, 0x9e, 0x52,
	0xb2, 0x52, 0xd8, 0xb1, 0xb8, 0x5f, 0x89, 0x24, 0x23, 0x5c, 0x84, 0x59, 0x5e, 0xb7, 0xf9, 0x57,
	0x30, 0xbe, 0xe1, 0xbc, 0x20, 0xb7, 0xa4, 0x0a, 0xc8, 0x53, 0x41, 0xb8, 0x40, 0x73, 0x18, 0x24,
	0x12, 0x62, 0xae, 0xb5, 0xb0, 0x96, 0xc3, 0x40, 0x57, 0xc8, 0x03, 0xa7, 0xe0, 0x84, 0x3d, 0x86,
	0x19, 0x71, 0x7b, 0xea, 0xa5, 0xa9, 0xfd, 0x0d, 0x8c, 0xce, 0x63, 0xf2, 0xb4, 0x42, 0x13, 0xe8,
	0x17, 0x2c, 0xd5, 0x13, 0xe4, 0x27, 0x72, 0xe1, 0x2b, 0x23, 0x11, 0x2d, 0x09, 0xd3, 0x6c, 0x53,
	0xfa, 0xcf, 0x16, 0x8c, 0x0f, 0x61, 0x9a, 0xc4, 0xa1, 0x20, 0x17, 0x88, 0x90, 0x6f, 0x79, 0xc8,
	0x79, 0x44, 0x63, 0xe2, 0xf6, 0xeb, 0x37, 0x53, 0x23, 0x0c, 0xb6, 0xb4, 0xee, 0xda, 0x0b, 0x6b,
	0xf9, 0x6d, 0xed, 0xe1, 0x3a, 0x17, 0x6c, 0x72, 0xc1, 0x7b, 0x93, 0x4b, 0xa0, 0xfa, 0xfc, 0xbf,
	0x30, 0x3a, 0x4b, 0x92, 0x86, 0x3c, 0x70, 0x4a, 0x0d, 0x28, 0x49, 0x4e, 0xd0, 0xd4, 0xfe, 0x1d,
	0xfc, 0x08, 0x6a, 0x2f, 0x97, 0xc8, 0x7f, 0x13, 0x50, 0xbf, 0x1d, 0xd0, 0x01, 0xbe, 0x37, 0xf3,
	0x3f, 0xd1, 0x62, 0x82, 0xef, 0x7d, 0x18, 0x7c, 0x67, 0xee, 0x35, 0x4c, 0x02, 0x92, 0xd1, 0xf2,
	0xd2, 0xed, 0xff, 0x91, 0xfe, 0x9b, 0x39, 0x52, 0xa1, 0xfa, 0xa7, 0x44, 0x62, 0x2d, 0xd0, 0x94,
	0xeb, 0x17, 0x0b, 0xec, 0x3d, 0x15, 0x39, 0xfa, 0x0f, 0x4e, 0xa2, 0x4f, 0x06, 0xcd, 0xb0, 0xbc,
	0x56, 0xdc, 0xb9, 0x44, 0xef, 0x67, 0x17, 0xce, 0xd3, 0xca, 0xff, 0x22, 0x99, 0x8d, 0x5d, 0xcd,
	0xec, 0x9c, 0x8f, 0x61, 0xb6, 0x56, 0xa8, 0x98, 0xa0, 0xbd, 0xcb, 0xbf, 0x4e, 0xeb, 0xa6, 0xf6,
	0xea, 0x3c, 0xd4, 0x41, 0x6b, 0xe6, 0x06, 0x86, 0xcc, 0x58, 0x44, 0x73, 0xd3, 0xd2, 0xce, 0xce,
	0x9b, 0xbe, 0xc3, 0x15, 0x79, 0xfb, 0x1b, 0x66, 0x51, 0x4a, 0x8b, 0x18, 0x57, 0x3c, 0x11, 0x31,
	0x0e, 0x0b, 0xf1, 0xa0, 0x3a, 0xb7, 0x43, 0x99, 0xc4, 0x4e, 0xde, 0xe0, 0xce, 0x3a, 0x0e, 0xd4,
	0x31, 0xfe, 0x7b, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xbb, 0xf5, 0x17, 0x3e, 0xc9, 0x03, 0x00, 0x00,
}
