/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: service.proto

package main

import (
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type VersionRequest struct {
	// Version of the KMS plugin API.
	Version              string   `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VersionRequest) Reset()         { *m = VersionRequest{} }
func (m *VersionRequest) String() string { return proto.CompactTextString(m) }
func (*VersionRequest) ProtoMessage()    {}
func (*VersionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{0}
}
func (m *VersionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VersionRequest.Unmarshal(m, b)
}
func (m *VersionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VersionRequest.Marshal(b, m, deterministic)
}
func (m *VersionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VersionRequest.Merge(m, src)
}
func (m *VersionRequest) XXX_Size() int {
	return xxx_messageInfo_VersionRequest.Size(m)
}
func (m *VersionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_VersionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_VersionRequest proto.InternalMessageInfo

func (m *VersionRequest) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

type VersionResponse struct {
	// Version of the KMS plugin API.
	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	// Name of the KMS provider.
	RuntimeName string `protobuf:"bytes,2,opt,name=runtime_name,json=runtimeName,proto3" json:"runtime_name,omitempty"`
	// Version of the KMS provider. The string must be semver-compatible.
	RuntimeVersion       string   `protobuf:"bytes,3,opt,name=runtime_version,json=runtimeVersion,proto3" json:"runtime_version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VersionResponse) Reset()         { *m = VersionResponse{} }
func (m *VersionResponse) String() string { return proto.CompactTextString(m) }
func (*VersionResponse) ProtoMessage()    {}
func (*VersionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{1}
}
func (m *VersionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VersionResponse.Unmarshal(m, b)
}
func (m *VersionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VersionResponse.Marshal(b, m, deterministic)
}
func (m *VersionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VersionResponse.Merge(m, src)
}
func (m *VersionResponse) XXX_Size() int {
	return xxx_messageInfo_VersionResponse.Size(m)
}
func (m *VersionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_VersionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_VersionResponse proto.InternalMessageInfo

func (m *VersionResponse) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *VersionResponse) GetRuntimeName() string {
	if m != nil {
		return m.RuntimeName
	}
	return ""
}

func (m *VersionResponse) GetRuntimeVersion() string {
	if m != nil {
		return m.RuntimeVersion
	}
	return ""
}

type DecryptRequest struct {
	// Version of the KMS plugin API.
	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	// The data to be decrypted.
	Cipher               []byte   `protobuf:"bytes,2,opt,name=cipher,proto3" json:"cipher,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DecryptRequest) Reset()         { *m = DecryptRequest{} }
func (m *DecryptRequest) String() string { return proto.CompactTextString(m) }
func (*DecryptRequest) ProtoMessage()    {}
func (*DecryptRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{2}
}
func (m *DecryptRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DecryptRequest.Unmarshal(m, b)
}
func (m *DecryptRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DecryptRequest.Marshal(b, m, deterministic)
}
func (m *DecryptRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DecryptRequest.Merge(m, src)
}
func (m *DecryptRequest) XXX_Size() int {
	return xxx_messageInfo_DecryptRequest.Size(m)
}
func (m *DecryptRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DecryptRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DecryptRequest proto.InternalMessageInfo

func (m *DecryptRequest) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *DecryptRequest) GetCipher() []byte {
	if m != nil {
		return m.Cipher
	}
	return nil
}

type DecryptResponse struct {
	// The decrypted data.
	Plain                []byte   `protobuf:"bytes,1,opt,name=plain,proto3" json:"plain,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DecryptResponse) Reset()         { *m = DecryptResponse{} }
func (m *DecryptResponse) String() string { return proto.CompactTextString(m) }
func (*DecryptResponse) ProtoMessage()    {}
func (*DecryptResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{3}
}
func (m *DecryptResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DecryptResponse.Unmarshal(m, b)
}
func (m *DecryptResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DecryptResponse.Marshal(b, m, deterministic)
}
func (m *DecryptResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DecryptResponse.Merge(m, src)
}
func (m *DecryptResponse) XXX_Size() int {
	return xxx_messageInfo_DecryptResponse.Size(m)
}
func (m *DecryptResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DecryptResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DecryptResponse proto.InternalMessageInfo

func (m *DecryptResponse) GetPlain() []byte {
	if m != nil {
		return m.Plain
	}
	return nil
}

type EncryptRequest struct {
	// Version of the KMS plugin API.
	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	// The data to be encrypted.
	Plain                []byte   `protobuf:"bytes,2,opt,name=plain,proto3" json:"plain,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EncryptRequest) Reset()         { *m = EncryptRequest{} }
func (m *EncryptRequest) String() string { return proto.CompactTextString(m) }
func (*EncryptRequest) ProtoMessage()    {}
func (*EncryptRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{4}
}
func (m *EncryptRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EncryptRequest.Unmarshal(m, b)
}
func (m *EncryptRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EncryptRequest.Marshal(b, m, deterministic)
}
func (m *EncryptRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EncryptRequest.Merge(m, src)
}
func (m *EncryptRequest) XXX_Size() int {
	return xxx_messageInfo_EncryptRequest.Size(m)
}
func (m *EncryptRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EncryptRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EncryptRequest proto.InternalMessageInfo

func (m *EncryptRequest) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *EncryptRequest) GetPlain() []byte {
	if m != nil {
		return m.Plain
	}
	return nil
}

type EncryptResponse struct {
	// The encrypted data.
	Cipher               []byte   `protobuf:"bytes,1,opt,name=cipher,proto3" json:"cipher,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EncryptResponse) Reset()         { *m = EncryptResponse{} }
func (m *EncryptResponse) String() string { return proto.CompactTextString(m) }
func (*EncryptResponse) ProtoMessage()    {}
func (*EncryptResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{5}
}
func (m *EncryptResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EncryptResponse.Unmarshal(m, b)
}
func (m *EncryptResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EncryptResponse.Marshal(b, m, deterministic)
}
func (m *EncryptResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EncryptResponse.Merge(m, src)
}
func (m *EncryptResponse) XXX_Size() int {
	return xxx_messageInfo_EncryptResponse.Size(m)
}
func (m *EncryptResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EncryptResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EncryptResponse proto.InternalMessageInfo

func (m *EncryptResponse) GetCipher() []byte {
	if m != nil {
		return m.Cipher
	}
	return nil
}

func init() {
	proto.RegisterType((*VersionRequest)(nil), "v1beta1.VersionRequest")
	proto.RegisterType((*VersionResponse)(nil), "v1beta1.VersionResponse")
	proto.RegisterType((*DecryptRequest)(nil), "v1beta1.DecryptRequest")
	proto.RegisterType((*DecryptResponse)(nil), "v1beta1.DecryptResponse")
	proto.RegisterType((*EncryptRequest)(nil), "v1beta1.EncryptRequest")
	proto.RegisterType((*EncryptResponse)(nil), "v1beta1.EncryptResponse")
}

func init() { proto.RegisterFile("service.proto", fileDescriptor_a0b84a42fa06f626) }

var fileDescriptor_a0b84a42fa06f626 = []byte{
	// 287 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xcd, 0x4a, 0xc4, 0x30,
	0x10, 0xde, 0xae, 0xb8, 0xc5, 0xb1, 0xb6, 0x10, 0x16, 0x2d, 0x9e, 0x34, 0x97, 0x55, 0x0f, 0x85,
	0xd5, 0xbb, 0x88, 0xe8, 0x49, 0xf4, 0x50, 0xc1, 0xab, 0x64, 0xcb, 0xa0, 0x05, 0x9b, 0xc6, 0x24,
	0x5b, 0xd9, 0x17, 0xf5, 0x79, 0xc4, 0x66, 0x5a, 0xd3, 0x15, 0x71, 0x8f, 0x33, 0x99, 0xef, 0x6f,
	0x26, 0xb0, 0x67, 0x50, 0x37, 0x65, 0x81, 0x99, 0xd2, 0xb5, 0xad, 0x59, 0xd8, 0xcc, 0x17, 0x68,
	0xc5, 0x9c, 0x9f, 0x41, 0xfc, 0x84, 0xda, 0x94, 0xb5, 0xcc, 0xf1, 0x7d, 0x89, 0xc6, 0xb2, 0x14,
	0xc2, 0xc6, 0x75, 0xd2, 0xe0, 0x28, 0x38, 0xd9, 0xc9, 0xbb, 0x92, 0x7f, 0x40, 0xd2, 0xcf, 0x1a,
	0x55, 0x4b, 0x83, 0x7f, 0x0f, 0xb3, 0x63, 0x88, 0xf4, 0x52, 0xda, 0xb2, 0xc2, 0x67, 0x29, 0x2a,
	0x4c, 0xc7, 0xed, 0xf3, 0x2e, 0xf5, 0x1e, 0x44, 0x85, 0x6c, 0x06, 0x49, 0x37, 0xd2, 0x91, 0x6c,
	0xb5, 0x53, 0x31, 0xb5, 0x49, 0x8d, 0x5f, 0x43, 0x7c, 0x83, 0x85, 0x5e, 0x29, 0xfb, 0xaf, 0x49,
	0xb6, 0x0f, 0x93, 0xa2, 0x54, 0xaf, 0xa8, 0x5b, 0xc5, 0x28, 0xa7, 0x8a, 0xcf, 0x20, 0xe9, 0x39,
	0xc8, 0xfc, 0x14, 0xb6, 0xd5, 0x9b, 0x28, 0x1d, 0x45, 0x94, 0xbb, 0x82, 0x5f, 0x41, 0x7c, 0x2b,
	0x37, 0x14, 0xeb, 0x19, 0xc6, 0x3e, 0xc3, 0x29, 0x24, 0x3d, 0x03, 0x49, 0xfd, 0xb8, 0x0a, 0x7c,
	0x57, 0xe7, 0x9f, 0x01, 0x4c, 0xef, 0x70, 0x75, 0x2f, 0xa4, 0x78, 0xc1, 0x0a, 0xa5, 0x7d, 0x74,
	0x67, 0x62, 0x97, 0x10, 0x52, 0x7a, 0x76, 0x90, 0xd1, 0xb1, 0xb2, 0xe1, 0xa5, 0x0e, 0xd3, 0xdf,
	0x0f, 0x4e, 0x8e, 0x8f, 0xbe, 0xf1, 0x14, 0xd7, 0xc3, 0x0f, 0x97, 0xe8, 0xe1, 0xd7, 0x36, 0xe3,
	0xf0, 0x94, 0xc1, 0xc3, 0x0f, 0xf7, 0xe2, 0xe1, 0xd7, 0xe2, 0xf2, 0xd1, 0x62, 0xd2, 0xfe, 0xb3,
	0x8b, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x33, 0x8d, 0x09, 0xe1, 0x78, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// KeyManagementServiceClient is the client API for KeyManagementService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type KeyManagementServiceClient interface {
	// Version returns the runtime name and runtime version of the KMS provider.
	Version(ctx context.Context, in *VersionRequest, opts ...grpc.CallOption) (*VersionResponse, error)
	// Execute decryption operation in KMS provider.
	Decrypt(ctx context.Context, in *DecryptRequest, opts ...grpc.CallOption) (*DecryptResponse, error)
	// Execute encryption operation in KMS provider.
	Encrypt(ctx context.Context, in *EncryptRequest, opts ...grpc.CallOption) (*EncryptResponse, error)
}

type keyManagementServiceClient struct {
	cc *grpc.ClientConn
}

func NewKeyManagementServiceClient(cc *grpc.ClientConn) KeyManagementServiceClient {
	return &keyManagementServiceClient{cc}
}

func (c *keyManagementServiceClient) Version(ctx context.Context, in *VersionRequest, opts ...grpc.CallOption) (*VersionResponse, error) {
	out := new(VersionResponse)
	err := c.cc.Invoke(ctx, "/v1beta1.KeyManagementService/Version", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyManagementServiceClient) Decrypt(ctx context.Context, in *DecryptRequest, opts ...grpc.CallOption) (*DecryptResponse, error) {
	out := new(DecryptResponse)
	err := c.cc.Invoke(ctx, "/v1beta1.KeyManagementService/Decrypt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyManagementServiceClient) Encrypt(ctx context.Context, in *EncryptRequest, opts ...grpc.CallOption) (*EncryptResponse, error) {
	out := new(EncryptResponse)
	err := c.cc.Invoke(ctx, "/v1beta1.KeyManagementService/Encrypt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KeyManagementServiceServer is the server API for KeyManagementService service.
type KeyManagementServiceServer interface {
	// Version returns the runtime name and runtime version of the KMS provider.
	Version(context.Context, *VersionRequest) (*VersionResponse, error)
	// Execute decryption operation in KMS provider.
	Decrypt(context.Context, *DecryptRequest) (*DecryptResponse, error)
	// Execute encryption operation in KMS provider.
	Encrypt(context.Context, *EncryptRequest) (*EncryptResponse, error)
}

// UnimplementedKeyManagementServiceServer can be embedded to have forward compatible implementations.
type UnimplementedKeyManagementServiceServer struct {
}

func (*UnimplementedKeyManagementServiceServer) Version(ctx context.Context, req *VersionRequest) (*VersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Version not implemented")
}
func (*UnimplementedKeyManagementServiceServer) Decrypt(ctx context.Context, req *DecryptRequest) (*DecryptResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Decrypt not implemented")
}
func (*UnimplementedKeyManagementServiceServer) Encrypt(ctx context.Context, req *EncryptRequest) (*EncryptResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Encrypt not implemented")
}

func RegisterKeyManagementServiceServer(s *grpc.Server, srv KeyManagementServiceServer) {
	s.RegisterService(&_KeyManagementService_serviceDesc, srv)
}

func _KeyManagementService_Version_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyManagementServiceServer).Version(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1beta1.KeyManagementService/Version",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyManagementServiceServer).Version(ctx, req.(*VersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyManagementService_Decrypt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DecryptRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyManagementServiceServer).Decrypt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1beta1.KeyManagementService/Decrypt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyManagementServiceServer).Decrypt(ctx, req.(*DecryptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyManagementService_Encrypt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EncryptRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyManagementServiceServer).Encrypt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1beta1.KeyManagementService/Encrypt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyManagementServiceServer).Encrypt(ctx, req.(*EncryptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _KeyManagementService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1beta1.KeyManagementService",
	HandlerType: (*KeyManagementServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Version",
			Handler:    _KeyManagementService_Version_Handler,
		},
		{
			MethodName: "Decrypt",
			Handler:    _KeyManagementService_Decrypt_Handler,
		},
		{
			MethodName: "Encrypt",
			Handler:    _KeyManagementService_Encrypt_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
