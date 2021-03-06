// Code generated by protoc-gen-go. DO NOT EDIT.
// source: security.proto

package proto

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

//*
// A wrapper message to logically group other messages
type TLSConfiguration struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TLSConfiguration) Reset()         { *m = TLSConfiguration{} }
func (m *TLSConfiguration) String() string { return proto.CompactTextString(m) }
func (*TLSConfiguration) ProtoMessage()    {}
func (*TLSConfiguration) Descriptor() ([]byte, []int) {
	return fileDescriptor_55a487c716a8b59c, []int{0}
}

func (m *TLSConfiguration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TLSConfiguration.Unmarshal(m, b)
}
func (m *TLSConfiguration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TLSConfiguration.Marshal(b, m, deterministic)
}
func (m *TLSConfiguration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TLSConfiguration.Merge(m, src)
}
func (m *TLSConfiguration) XXX_Size() int {
	return xxx_messageInfo_TLSConfiguration.Size(m)
}
func (m *TLSConfiguration) XXX_DiscardUnknown() {
	xxx_messageInfo_TLSConfiguration.DiscardUnknown(m)
}

var xxx_messageInfo_TLSConfiguration proto.InternalMessageInfo

// It's an empty Request received by RPC service
type TLSConfiguration_Request struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TLSConfiguration_Request) Reset()         { *m = TLSConfiguration_Request{} }
func (m *TLSConfiguration_Request) String() string { return proto.CompactTextString(m) }
func (*TLSConfiguration_Request) ProtoMessage()    {}
func (*TLSConfiguration_Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_55a487c716a8b59c, []int{0, 0}
}

func (m *TLSConfiguration_Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TLSConfiguration_Request.Unmarshal(m, b)
}
func (m *TLSConfiguration_Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TLSConfiguration_Request.Marshal(b, m, deterministic)
}
func (m *TLSConfiguration_Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TLSConfiguration_Request.Merge(m, src)
}
func (m *TLSConfiguration_Request) XXX_Size() int {
	return xxx_messageInfo_TLSConfiguration_Request.Size(m)
}
func (m *TLSConfiguration_Request) XXX_DiscardUnknown() {
	xxx_messageInfo_TLSConfiguration_Request.DiscardUnknown(m)
}

var xxx_messageInfo_TLSConfiguration_Request proto.InternalMessageInfo

// Response from RPC service
type TLSConfiguration_Response struct {
	Data                 *TLSConfiguration_Data `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *TLSConfiguration_Response) Reset()         { *m = TLSConfiguration_Response{} }
func (m *TLSConfiguration_Response) String() string { return proto.CompactTextString(m) }
func (*TLSConfiguration_Response) ProtoMessage()    {}
func (*TLSConfiguration_Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_55a487c716a8b59c, []int{0, 1}
}

func (m *TLSConfiguration_Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TLSConfiguration_Response.Unmarshal(m, b)
}
func (m *TLSConfiguration_Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TLSConfiguration_Response.Marshal(b, m, deterministic)
}
func (m *TLSConfiguration_Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TLSConfiguration_Response.Merge(m, src)
}
func (m *TLSConfiguration_Response) XXX_Size() int {
	return xxx_messageInfo_TLSConfiguration_Response.Size(m)
}
func (m *TLSConfiguration_Response) XXX_DiscardUnknown() {
	xxx_messageInfo_TLSConfiguration_Response.DiscardUnknown(m)
}

var xxx_messageInfo_TLSConfiguration_Response proto.InternalMessageInfo

func (m *TLSConfiguration_Response) GetData() *TLSConfiguration_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

// TLS configuration data for `geth`
type TLSConfiguration_Data struct {
	// Private key in PEM format
	KeyPem []byte `protobuf:"bytes,1,opt,name=keyPem,proto3" json:"keyPem,omitempty"`
	// Certificate in PEM format
	CertPem []byte `protobuf:"bytes,2,opt,name=certPem,proto3" json:"certPem,omitempty"`
	// List of cipher suites constants being supported by the server
	CipherSuites         []uint32 `protobuf:"varint,3,rep,packed,name=cipherSuites,proto3" json:"cipherSuites,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TLSConfiguration_Data) Reset()         { *m = TLSConfiguration_Data{} }
func (m *TLSConfiguration_Data) String() string { return proto.CompactTextString(m) }
func (*TLSConfiguration_Data) ProtoMessage()    {}
func (*TLSConfiguration_Data) Descriptor() ([]byte, []int) {
	return fileDescriptor_55a487c716a8b59c, []int{0, 2}
}

func (m *TLSConfiguration_Data) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TLSConfiguration_Data.Unmarshal(m, b)
}
func (m *TLSConfiguration_Data) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TLSConfiguration_Data.Marshal(b, m, deterministic)
}
func (m *TLSConfiguration_Data) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TLSConfiguration_Data.Merge(m, src)
}
func (m *TLSConfiguration_Data) XXX_Size() int {
	return xxx_messageInfo_TLSConfiguration_Data.Size(m)
}
func (m *TLSConfiguration_Data) XXX_DiscardUnknown() {
	xxx_messageInfo_TLSConfiguration_Data.DiscardUnknown(m)
}

var xxx_messageInfo_TLSConfiguration_Data proto.InternalMessageInfo

func (m *TLSConfiguration_Data) GetKeyPem() []byte {
	if m != nil {
		return m.KeyPem
	}
	return nil
}

func (m *TLSConfiguration_Data) GetCertPem() []byte {
	if m != nil {
		return m.CertPem
	}
	return nil
}

func (m *TLSConfiguration_Data) GetCipherSuites() []uint32 {
	if m != nil {
		return m.CipherSuites
	}
	return nil
}

//
// Representing a permission being extracted from access token by the plugin implementation.
// This permission is then stored in security context of a request and
// used internally to decide if the access is granted/denied
type GrantedAuthority struct {
	// `geth` RPC API namespace. E.g.: rpc, eth, admin, debug, ...
	Service string `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	// `geth` RPC API function. E.g.: nodeInfo, blockNumber, ...
	Method string `protobuf:"bytes,2,opt,name=method,proto3" json:"method,omitempty"`
	// raw string of the the granted authority value. This gives plugin implementation freedom to interpret the value
	Raw                  string   `protobuf:"bytes,3,opt,name=raw,proto3" json:"raw,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GrantedAuthority) Reset()         { *m = GrantedAuthority{} }
func (m *GrantedAuthority) String() string { return proto.CompactTextString(m) }
func (*GrantedAuthority) ProtoMessage()    {}
func (*GrantedAuthority) Descriptor() ([]byte, []int) {
	return fileDescriptor_55a487c716a8b59c, []int{1}
}

func (m *GrantedAuthority) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GrantedAuthority.Unmarshal(m, b)
}
func (m *GrantedAuthority) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GrantedAuthority.Marshal(b, m, deterministic)
}
func (m *GrantedAuthority) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GrantedAuthority.Merge(m, src)
}
func (m *GrantedAuthority) XXX_Size() int {
	return xxx_messageInfo_GrantedAuthority.Size(m)
}
func (m *GrantedAuthority) XXX_DiscardUnknown() {
	xxx_messageInfo_GrantedAuthority.DiscardUnknown(m)
}

var xxx_messageInfo_GrantedAuthority proto.InternalMessageInfo

func (m *GrantedAuthority) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

func (m *GrantedAuthority) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *GrantedAuthority) GetRaw() string {
	if m != nil {
		return m.Raw
	}
	return ""
}

//
// Representing the access token for an authentication request
type AuthenticationToken struct {
	RawToken             []byte   `protobuf:"bytes,1,opt,name=rawToken,proto3" json:"rawToken,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthenticationToken) Reset()         { *m = AuthenticationToken{} }
func (m *AuthenticationToken) String() string { return proto.CompactTextString(m) }
func (*AuthenticationToken) ProtoMessage()    {}
func (*AuthenticationToken) Descriptor() ([]byte, []int) {
	return fileDescriptor_55a487c716a8b59c, []int{2}
}

func (m *AuthenticationToken) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthenticationToken.Unmarshal(m, b)
}
func (m *AuthenticationToken) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthenticationToken.Marshal(b, m, deterministic)
}
func (m *AuthenticationToken) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthenticationToken.Merge(m, src)
}
func (m *AuthenticationToken) XXX_Size() int {
	return xxx_messageInfo_AuthenticationToken.Size(m)
}
func (m *AuthenticationToken) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthenticationToken.DiscardUnknown(m)
}

var xxx_messageInfo_AuthenticationToken proto.InternalMessageInfo

func (m *AuthenticationToken) GetRawToken() []byte {
	if m != nil {
		return m.RawToken
	}
	return nil
}

//
// Representing an authenticated principal after `AuthenticationToken` has been processed
type PreAuthenticatedAuthenticationToken struct {
	RawToken             []byte               `protobuf:"bytes,1,opt,name=rawToken,proto3" json:"rawToken,omitempty"`
	ExpiredAt            *timestamp.Timestamp `protobuf:"bytes,2,opt,name=expiredAt,proto3" json:"expiredAt,omitempty"`
	Authorities          []*GrantedAuthority  `protobuf:"bytes,3,rep,name=authorities,proto3" json:"authorities,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *PreAuthenticatedAuthenticationToken) Reset()         { *m = PreAuthenticatedAuthenticationToken{} }
func (m *PreAuthenticatedAuthenticationToken) String() string { return proto.CompactTextString(m) }
func (*PreAuthenticatedAuthenticationToken) ProtoMessage()    {}
func (*PreAuthenticatedAuthenticationToken) Descriptor() ([]byte, []int) {
	return fileDescriptor_55a487c716a8b59c, []int{3}
}

func (m *PreAuthenticatedAuthenticationToken) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PreAuthenticatedAuthenticationToken.Unmarshal(m, b)
}
func (m *PreAuthenticatedAuthenticationToken) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PreAuthenticatedAuthenticationToken.Marshal(b, m, deterministic)
}
func (m *PreAuthenticatedAuthenticationToken) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PreAuthenticatedAuthenticationToken.Merge(m, src)
}
func (m *PreAuthenticatedAuthenticationToken) XXX_Size() int {
	return xxx_messageInfo_PreAuthenticatedAuthenticationToken.Size(m)
}
func (m *PreAuthenticatedAuthenticationToken) XXX_DiscardUnknown() {
	xxx_messageInfo_PreAuthenticatedAuthenticationToken.DiscardUnknown(m)
}

var xxx_messageInfo_PreAuthenticatedAuthenticationToken proto.InternalMessageInfo

func (m *PreAuthenticatedAuthenticationToken) GetRawToken() []byte {
	if m != nil {
		return m.RawToken
	}
	return nil
}

func (m *PreAuthenticatedAuthenticationToken) GetExpiredAt() *timestamp.Timestamp {
	if m != nil {
		return m.ExpiredAt
	}
	return nil
}

func (m *PreAuthenticatedAuthenticationToken) GetAuthorities() []*GrantedAuthority {
	if m != nil {
		return m.Authorities
	}
	return nil
}

func init() {
	proto.RegisterType((*TLSConfiguration)(nil), "proto.TLSConfiguration")
	proto.RegisterType((*TLSConfiguration_Request)(nil), "proto.TLSConfiguration.Request")
	proto.RegisterType((*TLSConfiguration_Response)(nil), "proto.TLSConfiguration.Response")
	proto.RegisterType((*TLSConfiguration_Data)(nil), "proto.TLSConfiguration.Data")
	proto.RegisterType((*GrantedAuthority)(nil), "proto.GrantedAuthority")
	proto.RegisterType((*AuthenticationToken)(nil), "proto.AuthenticationToken")
	proto.RegisterType((*PreAuthenticatedAuthenticationToken)(nil), "proto.PreAuthenticatedAuthenticationToken")
}

func init() { proto.RegisterFile("security.proto", fileDescriptor_55a487c716a8b59c) }

var fileDescriptor_55a487c716a8b59c = []byte{
	// 452 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0xc1, 0x6e, 0x13, 0x31,
	0x10, 0x55, 0xd8, 0xd2, 0x26, 0x93, 0x80, 0x22, 0x23, 0xda, 0x68, 0x85, 0x44, 0xb4, 0x5c, 0xaa,
	0x1e, 0x1c, 0x58, 0x2e, 0x20, 0xb8, 0xb4, 0x20, 0x95, 0x03, 0x48, 0x95, 0x13, 0xf5, 0x80, 0xb8,
	0x38, 0x9b, 0xe9, 0xc6, 0x6a, 0xd6, 0xde, 0x7a, 0xc7, 0x94, 0xfc, 0x18, 0x12, 0x7f, 0x87, 0xd6,
	0xeb, 0xa5, 0x69, 0x44, 0x24, 0x38, 0x79, 0x9e, 0x67, 0xe6, 0x79, 0xfc, 0x9e, 0x0d, 0x8f, 0x2b,
	0xcc, 0x9c, 0x55, 0xb4, 0xe6, 0xa5, 0x35, 0x64, 0xd8, 0x43, 0xbf, 0xc4, 0xef, 0x72, 0x45, 0x4b,
	0x37, 0xe7, 0x99, 0x29, 0x26, 0xb9, 0x59, 0x49, 0x9d, 0x4f, 0x7c, 0x62, 0xee, 0xae, 0x26, 0x25,
	0xad, 0x4b, 0xac, 0x26, 0xa4, 0x0a, 0xac, 0x48, 0x16, 0xe5, 0x5d, 0xd4, 0x70, 0x24, 0xbf, 0x3a,
	0x30, 0x9c, 0x7d, 0x9e, 0x7e, 0x30, 0xfa, 0x4a, 0xe5, 0xce, 0x4a, 0x52, 0x46, 0xc7, 0x3d, 0x38,
	0x10, 0x78, 0xe3, 0xb0, 0xa2, 0xf8, 0x3d, 0x74, 0x05, 0x56, 0xa5, 0xd1, 0x15, 0xb2, 0x97, 0xb0,
	0xb7, 0x90, 0x24, 0x47, 0x9d, 0x71, 0xe7, 0xb8, 0x9f, 0x3e, 0x6b, 0x18, 0xf8, 0x76, 0x37, 0xff,
	0x28, 0x49, 0x0a, 0x5f, 0x19, 0x7f, 0x83, 0xbd, 0x1a, 0xb1, 0x43, 0xd8, 0xbf, 0xc6, 0xf5, 0x05,
	0x16, 0xbe, 0x77, 0x20, 0x02, 0x62, 0x23, 0x38, 0xc8, 0xd0, 0x52, 0x9d, 0x78, 0xe0, 0x13, 0x2d,
	0x64, 0x09, 0x0c, 0x32, 0x55, 0x2e, 0xd1, 0x4e, 0x9d, 0x22, 0xac, 0x46, 0xd1, 0x38, 0x3a, 0x7e,
	0x24, 0xee, 0xed, 0x25, 0x97, 0x30, 0x3c, 0xb7, 0x52, 0x13, 0x2e, 0x4e, 0x1d, 0x2d, 0x4d, 0xad,
	0x4c, 0xcd, 0x58, 0xa1, 0xfd, 0xae, 0x32, 0xf4, 0x47, 0xf5, 0x44, 0x0b, 0xeb, 0x19, 0x0a, 0xa4,
	0xa5, 0x59, 0xf8, 0xa3, 0x7a, 0x22, 0x20, 0x36, 0x84, 0xc8, 0xca, 0xdb, 0x51, 0xe4, 0x37, 0xeb,
	0x30, 0x79, 0x05, 0x4f, 0x6a, 0x42, 0xd4, 0xa4, 0x32, 0x7f, 0xa5, 0x99, 0xb9, 0x46, 0xcd, 0x62,
	0xe8, 0x5a, 0x79, 0xeb, 0xe3, 0x70, 0x8d, 0x3f, 0x38, 0xf9, 0xd9, 0x81, 0x17, 0x17, 0x16, 0x37,
	0xda, 0x9a, 0xa1, 0xfe, 0x83, 0x83, 0xbd, 0x81, 0x1e, 0xfe, 0x28, 0x95, 0xc5, 0xc5, 0x29, 0xf9,
	0x19, 0xfb, 0x69, 0xcc, 0x73, 0x63, 0xf2, 0x15, 0xf2, 0xd6, 0x50, 0x3e, 0x6b, 0xfd, 0x13, 0x77,
	0xc5, 0xec, 0x2d, 0xf4, 0x65, 0x50, 0x40, 0x05, 0xad, 0xfa, 0xe9, 0x51, 0xf0, 0x67, 0x5b, 0x22,
	0xb1, 0x59, 0x9b, 0xce, 0xe1, 0x70, 0xdb, 0xc0, 0xa9, 0x71, 0x36, 0x43, 0xf6, 0x09, 0xa2, 0x73,
	0x24, 0xf6, 0x7c, 0x97, 0xcd, 0xed, 0x0b, 0x19, 0xef, 0x2e, 0x68, 0xde, 0x4d, 0x6a, 0xe0, 0xe9,
	0x7d, 0x2d, 0xbe, 0x48, 0x2d, 0x73, 0xb4, 0xec, 0x12, 0x06, 0x9b, 0x8a, 0xb1, 0x38, 0x50, 0xfd,
	0x45, 0xb9, 0xf8, 0x24, 0xe4, 0xfe, 0x41, 0xe5, 0xb3, 0x13, 0x38, 0xca, 0x4c, 0xc1, 0x6f, 0x9c,
	0xb1, 0xae, 0xe0, 0xe5, 0xca, 0xe5, 0x4a, 0x37, 0xed, 0x67, 0xdd, 0x69, 0xf8, 0x43, 0x5f, 0x9b,
	0xdf, 0x33, 0xdf, 0xf7, 0xcb, 0xeb, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xcc, 0x7f, 0xc4, 0xe9,
	0x5d, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TLSConfigurationSourceClient is the client API for TLSConfigurationSource service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TLSConfigurationSourceClient interface {
	Get(ctx context.Context, in *TLSConfiguration_Request, opts ...grpc.CallOption) (*TLSConfiguration_Response, error)
}

type tLSConfigurationSourceClient struct {
	cc *grpc.ClientConn
}

func NewTLSConfigurationSourceClient(cc *grpc.ClientConn) TLSConfigurationSourceClient {
	return &tLSConfigurationSourceClient{cc}
}

func (c *tLSConfigurationSourceClient) Get(ctx context.Context, in *TLSConfiguration_Request, opts ...grpc.CallOption) (*TLSConfiguration_Response, error) {
	out := new(TLSConfiguration_Response)
	err := c.cc.Invoke(ctx, "/proto.TLSConfigurationSource/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TLSConfigurationSourceServer is the server API for TLSConfigurationSource service.
type TLSConfigurationSourceServer interface {
	Get(context.Context, *TLSConfiguration_Request) (*TLSConfiguration_Response, error)
}

// UnimplementedTLSConfigurationSourceServer can be embedded to have forward compatible implementations.
type UnimplementedTLSConfigurationSourceServer struct {
}

func (*UnimplementedTLSConfigurationSourceServer) Get(ctx context.Context, req *TLSConfiguration_Request) (*TLSConfiguration_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}

func RegisterTLSConfigurationSourceServer(s *grpc.Server, srv TLSConfigurationSourceServer) {
	s.RegisterService(&_TLSConfigurationSource_serviceDesc, srv)
}

func _TLSConfigurationSource_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TLSConfiguration_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TLSConfigurationSourceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.TLSConfigurationSource/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TLSConfigurationSourceServer).Get(ctx, req.(*TLSConfiguration_Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _TLSConfigurationSource_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.TLSConfigurationSource",
	HandlerType: (*TLSConfigurationSourceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _TLSConfigurationSource_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "security.proto",
}

// AuthenticationManagerClient is the client API for AuthenticationManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthenticationManagerClient interface {
	//
	// Perform authentication of the token. Return a token that contains expiry date and granted authorities
	Authenticate(ctx context.Context, in *AuthenticationToken, opts ...grpc.CallOption) (*PreAuthenticatedAuthenticationToken, error)
}

type authenticationManagerClient struct {
	cc *grpc.ClientConn
}

func NewAuthenticationManagerClient(cc *grpc.ClientConn) AuthenticationManagerClient {
	return &authenticationManagerClient{cc}
}

func (c *authenticationManagerClient) Authenticate(ctx context.Context, in *AuthenticationToken, opts ...grpc.CallOption) (*PreAuthenticatedAuthenticationToken, error) {
	out := new(PreAuthenticatedAuthenticationToken)
	err := c.cc.Invoke(ctx, "/proto.AuthenticationManager/Authenticate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthenticationManagerServer is the server API for AuthenticationManager service.
type AuthenticationManagerServer interface {
	//
	// Perform authentication of the token. Return a token that contains expiry date and granted authorities
	Authenticate(context.Context, *AuthenticationToken) (*PreAuthenticatedAuthenticationToken, error)
}

// UnimplementedAuthenticationManagerServer can be embedded to have forward compatible implementations.
type UnimplementedAuthenticationManagerServer struct {
}

func (*UnimplementedAuthenticationManagerServer) Authenticate(ctx context.Context, req *AuthenticationToken) (*PreAuthenticatedAuthenticationToken, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Authenticate not implemented")
}

func RegisterAuthenticationManagerServer(s *grpc.Server, srv AuthenticationManagerServer) {
	s.RegisterService(&_AuthenticationManager_serviceDesc, srv)
}

func _AuthenticationManager_Authenticate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthenticationToken)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationManagerServer).Authenticate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.AuthenticationManager/Authenticate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationManagerServer).Authenticate(ctx, req.(*AuthenticationToken))
	}
	return interceptor(ctx, in, info, handler)
}

var _AuthenticationManager_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.AuthenticationManager",
	HandlerType: (*AuthenticationManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Authenticate",
			Handler:    _AuthenticationManager_Authenticate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "security.proto",
}
