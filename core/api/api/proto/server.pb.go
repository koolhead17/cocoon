// Code generated by protoc-gen-go.
// source: server.proto
// DO NOT EDIT!

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	server.proto

It has these top-level messages:
	LoginRequest
	AddCocoonToIdentityRequest
	CreateCocoonRequest
	GetCocoonRequest
	GetIdentityRequest
	CreateReleaseRequest
	DeployRequest
	CreateIdentityRequest
	Response
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type LoginRequest struct {
	Email    string `protobuf:"bytes,1,opt,name=email" json:"email,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
}

func (m *LoginRequest) Reset()                    { *m = LoginRequest{} }
func (m *LoginRequest) String() string            { return proto1.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()               {}
func (*LoginRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *LoginRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *LoginRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type AddCocoonToIdentityRequest struct {
	Email    string `protobuf:"bytes,1,opt,name=email" json:"email,omitempty"`
	CocoonId string `protobuf:"bytes,2,opt,name=cocoonId" json:"cocoonId,omitempty"`
}

func (m *AddCocoonToIdentityRequest) Reset()                    { *m = AddCocoonToIdentityRequest{} }
func (m *AddCocoonToIdentityRequest) String() string            { return proto1.CompactTextString(m) }
func (*AddCocoonToIdentityRequest) ProtoMessage()               {}
func (*AddCocoonToIdentityRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *AddCocoonToIdentityRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *AddCocoonToIdentityRequest) GetCocoonId() string {
	if m != nil {
		return m.CocoonId
	}
	return ""
}

type CreateCocoonRequest struct {
	ID             string   `protobuf:"bytes,1,opt,name=ID" json:"ID,omitempty"`
	URL            string   `protobuf:"bytes,2,opt,name=URL" json:"URL,omitempty"`
	Language       string   `protobuf:"bytes,3,opt,name=language" json:"language,omitempty"`
	ReleaseTag     string   `protobuf:"bytes,4,opt,name=releaseTag" json:"releaseTag,omitempty"`
	BuildParam     string   `protobuf:"bytes,5,opt,name=buildParam" json:"buildParam,omitempty"`
	Memory         string   `protobuf:"bytes,6,opt,name=memory" json:"memory,omitempty"`
	CPUShare       string   `protobuf:"bytes,7,opt,name=CPUShare" json:"CPUShare,omitempty"`
	Releases       []string `protobuf:"bytes,8,rep,name=releases" json:"releases,omitempty"`
	Instances      int32    `protobuf:"varint,9,opt,name=instances" json:"instances,omitempty"`
	NumSignatories int32    `protobuf:"varint,10,opt,name=numSignatories" json:"numSignatories,omitempty"`
	SigThreshold   int32    `protobuf:"varint,11,opt,name=sigThreshold" json:"sigThreshold,omitempty"`
	Signatories    []string `protobuf:"bytes,12,rep,name=signatories" json:"signatories,omitempty"`
}

func (m *CreateCocoonRequest) Reset()                    { *m = CreateCocoonRequest{} }
func (m *CreateCocoonRequest) String() string            { return proto1.CompactTextString(m) }
func (*CreateCocoonRequest) ProtoMessage()               {}
func (*CreateCocoonRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *CreateCocoonRequest) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *CreateCocoonRequest) GetURL() string {
	if m != nil {
		return m.URL
	}
	return ""
}

func (m *CreateCocoonRequest) GetLanguage() string {
	if m != nil {
		return m.Language
	}
	return ""
}

func (m *CreateCocoonRequest) GetReleaseTag() string {
	if m != nil {
		return m.ReleaseTag
	}
	return ""
}

func (m *CreateCocoonRequest) GetBuildParam() string {
	if m != nil {
		return m.BuildParam
	}
	return ""
}

func (m *CreateCocoonRequest) GetMemory() string {
	if m != nil {
		return m.Memory
	}
	return ""
}

func (m *CreateCocoonRequest) GetCPUShare() string {
	if m != nil {
		return m.CPUShare
	}
	return ""
}

func (m *CreateCocoonRequest) GetReleases() []string {
	if m != nil {
		return m.Releases
	}
	return nil
}

func (m *CreateCocoonRequest) GetInstances() int32 {
	if m != nil {
		return m.Instances
	}
	return 0
}

func (m *CreateCocoonRequest) GetNumSignatories() int32 {
	if m != nil {
		return m.NumSignatories
	}
	return 0
}

func (m *CreateCocoonRequest) GetSigThreshold() int32 {
	if m != nil {
		return m.SigThreshold
	}
	return 0
}

func (m *CreateCocoonRequest) GetSignatories() []string {
	if m != nil {
		return m.Signatories
	}
	return nil
}

type GetCocoonRequest struct {
	ID       string `protobuf:"bytes,1,opt,name=ID" json:"ID,omitempty"`
	Identity string `protobuf:"bytes,2,opt,name=Identity" json:"Identity,omitempty"`
}

func (m *GetCocoonRequest) Reset()                    { *m = GetCocoonRequest{} }
func (m *GetCocoonRequest) String() string            { return proto1.CompactTextString(m) }
func (*GetCocoonRequest) ProtoMessage()               {}
func (*GetCocoonRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *GetCocoonRequest) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *GetCocoonRequest) GetIdentity() string {
	if m != nil {
		return m.Identity
	}
	return ""
}

type GetIdentityRequest struct {
	Email string `protobuf:"bytes,1,opt,name=email" json:"email,omitempty"`
}

func (m *GetIdentityRequest) Reset()                    { *m = GetIdentityRequest{} }
func (m *GetIdentityRequest) String() string            { return proto1.CompactTextString(m) }
func (*GetIdentityRequest) ProtoMessage()               {}
func (*GetIdentityRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *GetIdentityRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

type CreateReleaseRequest struct {
	ID          string `protobuf:"bytes,1,opt,name=ID" json:"ID,omitempty"`
	CocoonID    string `protobuf:"bytes,2,opt,name=cocoonID" json:"cocoonID,omitempty"`
	URL         string `protobuf:"bytes,3,opt,name=URL" json:"URL,omitempty"`
	Language    string `protobuf:"bytes,4,opt,name=language" json:"language,omitempty"`
	ReleaseTag  string `protobuf:"bytes,5,opt,name=releaseTag" json:"releaseTag,omitempty"`
	BuildParam  string `protobuf:"bytes,6,opt,name=buildParam" json:"buildParam,omitempty"`
	SigApproved int32  `protobuf:"varint,7,opt,name=sigApproved" json:"sigApproved,omitempty"`
	SigDenied   int32  `protobuf:"varint,8,opt,name=sigDenied" json:"sigDenied,omitempty"`
}

func (m *CreateReleaseRequest) Reset()                    { *m = CreateReleaseRequest{} }
func (m *CreateReleaseRequest) String() string            { return proto1.CompactTextString(m) }
func (*CreateReleaseRequest) ProtoMessage()               {}
func (*CreateReleaseRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *CreateReleaseRequest) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *CreateReleaseRequest) GetCocoonID() string {
	if m != nil {
		return m.CocoonID
	}
	return ""
}

func (m *CreateReleaseRequest) GetURL() string {
	if m != nil {
		return m.URL
	}
	return ""
}

func (m *CreateReleaseRequest) GetLanguage() string {
	if m != nil {
		return m.Language
	}
	return ""
}

func (m *CreateReleaseRequest) GetReleaseTag() string {
	if m != nil {
		return m.ReleaseTag
	}
	return ""
}

func (m *CreateReleaseRequest) GetBuildParam() string {
	if m != nil {
		return m.BuildParam
	}
	return ""
}

func (m *CreateReleaseRequest) GetSigApproved() int32 {
	if m != nil {
		return m.SigApproved
	}
	return 0
}

func (m *CreateReleaseRequest) GetSigDenied() int32 {
	if m != nil {
		return m.SigDenied
	}
	return 0
}

type DeployRequest struct {
	ID         string `protobuf:"bytes,1,opt,name=ID" json:"ID,omitempty"`
	URL        string `protobuf:"bytes,2,opt,name=URL" json:"URL,omitempty"`
	Language   string `protobuf:"bytes,3,opt,name=language" json:"language,omitempty"`
	ReleaseTag string `protobuf:"bytes,4,opt,name=releaseTag" json:"releaseTag,omitempty"`
	BuildParam []byte `protobuf:"bytes,5,opt,name=buildParam,proto3" json:"buildParam,omitempty"`
	Memory     string `protobuf:"bytes,6,opt,name=memory" json:"memory,omitempty"`
	CPUShare   string `protobuf:"bytes,7,opt,name=CPUShare" json:"CPUShare,omitempty"`
}

func (m *DeployRequest) Reset()                    { *m = DeployRequest{} }
func (m *DeployRequest) String() string            { return proto1.CompactTextString(m) }
func (*DeployRequest) ProtoMessage()               {}
func (*DeployRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *DeployRequest) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *DeployRequest) GetURL() string {
	if m != nil {
		return m.URL
	}
	return ""
}

func (m *DeployRequest) GetLanguage() string {
	if m != nil {
		return m.Language
	}
	return ""
}

func (m *DeployRequest) GetReleaseTag() string {
	if m != nil {
		return m.ReleaseTag
	}
	return ""
}

func (m *DeployRequest) GetBuildParam() []byte {
	if m != nil {
		return m.BuildParam
	}
	return nil
}

func (m *DeployRequest) GetMemory() string {
	if m != nil {
		return m.Memory
	}
	return ""
}

func (m *DeployRequest) GetCPUShare() string {
	if m != nil {
		return m.CPUShare
	}
	return ""
}

type CreateIdentityRequest struct {
	Email    string `protobuf:"bytes,1,opt,name=email" json:"email,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
}

func (m *CreateIdentityRequest) Reset()                    { *m = CreateIdentityRequest{} }
func (m *CreateIdentityRequest) String() string            { return proto1.CompactTextString(m) }
func (*CreateIdentityRequest) ProtoMessage()               {}
func (*CreateIdentityRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *CreateIdentityRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *CreateIdentityRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type Response struct {
	ID     string `protobuf:"bytes,1,opt,name=ID" json:"ID,omitempty"`
	Status int32  `protobuf:"varint,2,opt,name=status" json:"status,omitempty"`
	Body   []byte `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto1.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *Response) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Response) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *Response) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

func init() {
	proto1.RegisterType((*LoginRequest)(nil), "proto.LoginRequest")
	proto1.RegisterType((*AddCocoonToIdentityRequest)(nil), "proto.AddCocoonToIdentityRequest")
	proto1.RegisterType((*CreateCocoonRequest)(nil), "proto.CreateCocoonRequest")
	proto1.RegisterType((*GetCocoonRequest)(nil), "proto.GetCocoonRequest")
	proto1.RegisterType((*GetIdentityRequest)(nil), "proto.GetIdentityRequest")
	proto1.RegisterType((*CreateReleaseRequest)(nil), "proto.CreateReleaseRequest")
	proto1.RegisterType((*DeployRequest)(nil), "proto.DeployRequest")
	proto1.RegisterType((*CreateIdentityRequest)(nil), "proto.CreateIdentityRequest")
	proto1.RegisterType((*Response)(nil), "proto.Response")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for API service

type APIClient interface {
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*Response, error)
	CreateCocoon(ctx context.Context, in *CreateCocoonRequest, opts ...grpc.CallOption) (*Response, error)
	CreateRelease(ctx context.Context, in *CreateReleaseRequest, opts ...grpc.CallOption) (*Response, error)
	CreateIdentity(ctx context.Context, in *CreateIdentityRequest, opts ...grpc.CallOption) (*Response, error)
	Deploy(ctx context.Context, in *DeployRequest, opts ...grpc.CallOption) (*Response, error)
	GetCocoon(ctx context.Context, in *GetCocoonRequest, opts ...grpc.CallOption) (*Response, error)
	GetIdentity(ctx context.Context, in *GetIdentityRequest, opts ...grpc.CallOption) (*Response, error)
	AddCocoonToIdentity(ctx context.Context, in *AddCocoonToIdentityRequest, opts ...grpc.CallOption) (*Response, error)
}

type aPIClient struct {
	cc *grpc.ClientConn
}

func NewAPIClient(cc *grpc.ClientConn) APIClient {
	return &aPIClient{cc}
}

func (c *aPIClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/proto.API/Login", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) CreateCocoon(ctx context.Context, in *CreateCocoonRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/proto.API/CreateCocoon", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) CreateRelease(ctx context.Context, in *CreateReleaseRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/proto.API/CreateRelease", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) CreateIdentity(ctx context.Context, in *CreateIdentityRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/proto.API/CreateIdentity", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) Deploy(ctx context.Context, in *DeployRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/proto.API/Deploy", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) GetCocoon(ctx context.Context, in *GetCocoonRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/proto.API/GetCocoon", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) GetIdentity(ctx context.Context, in *GetIdentityRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/proto.API/GetIdentity", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) AddCocoonToIdentity(ctx context.Context, in *AddCocoonToIdentityRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/proto.API/AddCocoonToIdentity", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for API service

type APIServer interface {
	Login(context.Context, *LoginRequest) (*Response, error)
	CreateCocoon(context.Context, *CreateCocoonRequest) (*Response, error)
	CreateRelease(context.Context, *CreateReleaseRequest) (*Response, error)
	CreateIdentity(context.Context, *CreateIdentityRequest) (*Response, error)
	Deploy(context.Context, *DeployRequest) (*Response, error)
	GetCocoon(context.Context, *GetCocoonRequest) (*Response, error)
	GetIdentity(context.Context, *GetIdentityRequest) (*Response, error)
	AddCocoonToIdentity(context.Context, *AddCocoonToIdentityRequest) (*Response, error)
}

func RegisterAPIServer(s *grpc.Server, srv APIServer) {
	s.RegisterService(&_API_serviceDesc, srv)
}

func _API_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.API/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_CreateCocoon_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCocoonRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).CreateCocoon(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.API/CreateCocoon",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).CreateCocoon(ctx, req.(*CreateCocoonRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_CreateRelease_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateReleaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).CreateRelease(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.API/CreateRelease",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).CreateRelease(ctx, req.(*CreateReleaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_CreateIdentity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateIdentityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).CreateIdentity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.API/CreateIdentity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).CreateIdentity(ctx, req.(*CreateIdentityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_Deploy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeployRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).Deploy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.API/Deploy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).Deploy(ctx, req.(*DeployRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_GetCocoon_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCocoonRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).GetCocoon(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.API/GetCocoon",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).GetCocoon(ctx, req.(*GetCocoonRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_GetIdentity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetIdentityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).GetIdentity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.API/GetIdentity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).GetIdentity(ctx, req.(*GetIdentityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_AddCocoonToIdentity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddCocoonToIdentityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).AddCocoonToIdentity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.API/AddCocoonToIdentity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).AddCocoonToIdentity(ctx, req.(*AddCocoonToIdentityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _API_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.API",
	HandlerType: (*APIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _API_Login_Handler,
		},
		{
			MethodName: "CreateCocoon",
			Handler:    _API_CreateCocoon_Handler,
		},
		{
			MethodName: "CreateRelease",
			Handler:    _API_CreateRelease_Handler,
		},
		{
			MethodName: "CreateIdentity",
			Handler:    _API_CreateIdentity_Handler,
		},
		{
			MethodName: "Deploy",
			Handler:    _API_Deploy_Handler,
		},
		{
			MethodName: "GetCocoon",
			Handler:    _API_GetCocoon_Handler,
		},
		{
			MethodName: "GetIdentity",
			Handler:    _API_GetIdentity_Handler,
		},
		{
			MethodName: "AddCocoonToIdentity",
			Handler:    _API_AddCocoonToIdentity_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "server.proto",
}

func init() { proto1.RegisterFile("server.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 596 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xc4, 0x55, 0x5d, 0x6e, 0xd3, 0x40,
	0x10, 0x56, 0xe2, 0xda, 0x24, 0x53, 0xb7, 0x54, 0xdb, 0x52, 0x96, 0x50, 0xa1, 0xe0, 0x07, 0x54,
	0x21, 0x51, 0x24, 0x10, 0x0f, 0x08, 0xf1, 0x53, 0x35, 0xa2, 0xb2, 0x54, 0xa1, 0xc8, 0x4d, 0x0f,
	0xb0, 0x89, 0x47, 0x8e, 0x25, 0xc7, 0x6b, 0x76, 0x9d, 0xa2, 0x1c, 0x8f, 0x93, 0x70, 0x04, 0x38,
	0x02, 0xf2, 0xae, 0xed, 0xd8, 0x71, 0x9a, 0x54, 0xe2, 0x81, 0x27, 0xfb, 0x9b, 0xd9, 0xef, 0x1b,
	0x6b, 0xbe, 0xd9, 0x31, 0xd8, 0x12, 0xc5, 0x2d, 0x8a, 0xb3, 0x44, 0xf0, 0x94, 0x13, 0x53, 0x3d,
	0x9c, 0x2f, 0x60, 0x5f, 0xf1, 0x20, 0x8c, 0x3d, 0xfc, 0x3e, 0x47, 0x99, 0x92, 0x23, 0x30, 0x71,
	0xc6, 0xc2, 0x88, 0xb6, 0xfa, 0xad, 0xd3, 0xae, 0xa7, 0x01, 0xe9, 0x41, 0x27, 0x61, 0x52, 0xfe,
	0xe0, 0xc2, 0xa7, 0x6d, 0x95, 0x28, 0xb1, 0xf3, 0x0d, 0x7a, 0xe7, 0xbe, 0x7f, 0xc1, 0x27, 0x9c,
	0xc7, 0x23, 0xee, 0xfa, 0x18, 0xa7, 0x61, 0xba, 0xd8, 0xaa, 0x37, 0x51, 0x04, 0xb7, 0xd4, 0x2b,
	0xb0, 0xf3, 0xbb, 0x0d, 0x87, 0x17, 0x02, 0x59, 0x8a, 0x5a, 0xb3, 0x50, 0xda, 0x87, 0xb6, 0x3b,
	0xc8, 0x65, 0xda, 0xee, 0x80, 0x1c, 0x80, 0x71, 0xe3, 0x5d, 0xe5, 0xf4, 0xec, 0x35, 0x53, 0x8d,
	0x58, 0x1c, 0xcc, 0x59, 0x80, 0xd4, 0xd0, 0xaa, 0x05, 0x26, 0xcf, 0x00, 0x04, 0x46, 0xc8, 0x24,
	0x8e, 0x58, 0x40, 0x77, 0x54, 0xb6, 0x12, 0xc9, 0xf2, 0xe3, 0x79, 0x18, 0xf9, 0x43, 0x26, 0xd8,
	0x8c, 0x9a, 0x3a, 0xbf, 0x8c, 0x90, 0x63, 0xb0, 0x66, 0x38, 0xe3, 0x62, 0x41, 0x2d, 0x95, 0xcb,
	0x51, 0x56, 0xf3, 0x62, 0x78, 0x73, 0x3d, 0x65, 0x02, 0xe9, 0x03, 0x5d, 0xb3, 0xc0, 0x59, 0x2e,
	0xaf, 0x20, 0x69, 0xa7, 0x6f, 0x64, 0xb9, 0x02, 0x93, 0x13, 0xe8, 0x86, 0xb1, 0x4c, 0x59, 0x3c,
	0x41, 0x49, 0xbb, 0xfd, 0xd6, 0xa9, 0xe9, 0x2d, 0x03, 0xe4, 0x05, 0xec, 0xc7, 0xf3, 0xd9, 0x75,
	0x18, 0xc4, 0x2c, 0xe5, 0x22, 0x44, 0x49, 0x41, 0x1d, 0x59, 0x89, 0x12, 0x07, 0x6c, 0x19, 0x06,
	0xa3, 0xa9, 0x40, 0x39, 0xe5, 0x91, 0x4f, 0x77, 0xd5, 0xa9, 0x5a, 0x8c, 0xf4, 0x61, 0x57, 0x56,
	0x84, 0x6c, 0xf5, 0x21, 0xd5, 0x90, 0xf3, 0x09, 0x0e, 0x2e, 0x31, 0xdd, 0xdc, 0xed, 0x1e, 0x74,
	0x0a, 0x6b, 0x0b, 0xc7, 0x0a, 0xec, 0xbc, 0x04, 0x72, 0x89, 0xe9, 0xbd, 0x9c, 0x77, 0xfe, 0xb4,
	0xe0, 0x48, 0xbb, 0xeb, 0xe9, 0x56, 0x6c, 0x28, 0x98, 0x8f, 0xc4, 0x60, 0x65, 0x44, 0x4a, 0xeb,
	0x8d, 0xf5, 0xd6, 0xef, 0x6c, 0xb4, 0xde, 0xdc, 0x62, 0xbd, 0xd5, 0xb0, 0x5e, 0x37, 0xf0, 0x3c,
	0x49, 0x04, 0xbf, 0x45, 0x5f, 0xb9, 0x6c, 0x7a, 0xd5, 0x50, 0x66, 0xa6, 0x0c, 0x83, 0x01, 0xc6,
	0x21, 0xfa, 0xb4, 0xa3, 0xcd, 0x2c, 0x03, 0xce, 0xcf, 0x16, 0xec, 0x0d, 0x30, 0x89, 0xf8, 0xe2,
	0x7f, 0x8d, 0xb2, 0xfd, 0xaf, 0xa3, 0xec, 0xb8, 0xf0, 0x48, 0xbb, 0x76, 0xef, 0xfb, 0x7d, 0xe7,
	0xbe, 0xf8, 0x0a, 0x1d, 0x0f, 0x65, 0xc2, 0x63, 0x89, 0x8d, 0x46, 0x1c, 0x83, 0x25, 0x53, 0x96,
	0xce, 0xa5, 0x62, 0x99, 0x5e, 0x8e, 0x08, 0x81, 0x9d, 0x31, 0xf7, 0x17, 0xaa, 0x15, 0xb6, 0xa7,
	0xde, 0xdf, 0xfc, 0x32, 0xc0, 0x38, 0x1f, 0xba, 0xe4, 0x15, 0x98, 0x6a, 0x83, 0x91, 0x43, 0xbd,
	0xd9, 0xce, 0xaa, 0xfb, 0xac, 0xf7, 0x30, 0x0f, 0x96, 0x25, 0x3f, 0x80, 0x5d, 0xdd, 0x2e, 0xa4,
	0x97, 0x1f, 0x58, 0xb3, 0x72, 0x9a, 0xe4, 0x8f, 0xb0, 0x57, 0x1b, 0x5e, 0xf2, 0xb4, 0xc6, 0xae,
	0x8f, 0x74, 0x93, 0xfe, 0x19, 0xf6, 0xeb, 0x5d, 0x24, 0x27, 0x35, 0xfe, 0x4a, 0x73, 0x9b, 0x02,
	0xaf, 0xc1, 0xd2, 0x93, 0x44, 0x8e, 0xf2, 0x54, 0x6d, 0xb0, 0x9a, 0x84, 0x77, 0xd0, 0x2d, 0xaf,
	0x36, 0x79, 0x9c, 0x67, 0x57, 0x2f, 0x7b, 0x93, 0xf6, 0x1e, 0x76, 0x2b, 0x37, 0x9a, 0x3c, 0x59,
	0x12, 0xb7, 0x7e, 0xa2, 0x0b, 0x87, 0x6b, 0x7e, 0x07, 0xe4, 0x79, 0x7e, 0xee, 0xee, 0x5f, 0x45,
	0x43, 0x6a, 0x6c, 0x29, 0xfc, 0xf6, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x7f, 0xc9, 0xf7, 0xf4,
	0xb9, 0x06, 0x00, 0x00,
}