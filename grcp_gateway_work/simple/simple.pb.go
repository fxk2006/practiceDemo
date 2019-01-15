// Code generated by protoc-gen-go. DO NOT EDIT.
// source: simple.proto

package simple

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Package struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Package) Reset()         { *m = Package{} }
func (m *Package) String() string { return proto.CompactTextString(m) }
func (*Package) ProtoMessage()    {}
func (*Package) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ffd045dd4d042c1, []int{0}
}

func (m *Package) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Package.Unmarshal(m, b)
}
func (m *Package) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Package.Marshal(b, m, deterministic)
}
func (m *Package) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Package.Merge(m, src)
}
func (m *Package) XXX_Size() int {
	return xxx_messageInfo_Package.Size(m)
}
func (m *Package) XXX_DiscardUnknown() {
	xxx_messageInfo_Package.DiscardUnknown(m)
}

var xxx_messageInfo_Package proto.InternalMessageInfo

type Package_Request struct {
	Cmd                  uint32   `protobuf:"varint,1,opt,name=Cmd,proto3" json:"Cmd,omitempty"`
	Reserved             uint32   `protobuf:"varint,2,opt,name=reserved,proto3" json:"reserved,omitempty"`
	Body                 []byte   `protobuf:"bytes,3,opt,name=Body,proto3" json:"Body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Package_Request) Reset()         { *m = Package_Request{} }
func (m *Package_Request) String() string { return proto.CompactTextString(m) }
func (*Package_Request) ProtoMessage()    {}
func (*Package_Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ffd045dd4d042c1, []int{0, 0}
}

func (m *Package_Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Package_Request.Unmarshal(m, b)
}
func (m *Package_Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Package_Request.Marshal(b, m, deterministic)
}
func (m *Package_Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Package_Request.Merge(m, src)
}
func (m *Package_Request) XXX_Size() int {
	return xxx_messageInfo_Package_Request.Size(m)
}
func (m *Package_Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Package_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Package_Request proto.InternalMessageInfo

func (m *Package_Request) GetCmd() uint32 {
	if m != nil {
		return m.Cmd
	}
	return 0
}

func (m *Package_Request) GetReserved() uint32 {
	if m != nil {
		return m.Reserved
	}
	return 0
}

func (m *Package_Request) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

type Package_Response struct {
	Code                 uint32   `protobuf:"varint,1,opt,name=Code,proto3" json:"Code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data                 []byte   `protobuf:"bytes,3,opt,name=Data,proto3" json:"Data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Package_Response) Reset()         { *m = Package_Response{} }
func (m *Package_Response) String() string { return proto.CompactTextString(m) }
func (*Package_Response) ProtoMessage()    {}
func (*Package_Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ffd045dd4d042c1, []int{0, 1}
}

func (m *Package_Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Package_Response.Unmarshal(m, b)
}
func (m *Package_Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Package_Response.Marshal(b, m, deterministic)
}
func (m *Package_Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Package_Response.Merge(m, src)
}
func (m *Package_Response) XXX_Size() int {
	return xxx_messageInfo_Package_Response.Size(m)
}
func (m *Package_Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Package_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Package_Response proto.InternalMessageInfo

func (m *Package_Response) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Package_Response) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Package_Response) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*Package)(nil), "simple.Package")
	proto.RegisterType((*Package_Request)(nil), "simple.Package.Request")
	proto.RegisterType((*Package_Response)(nil), "simple.Package.Response")
}

func init() { proto.RegisterFile("simple.proto", fileDescriptor_5ffd045dd4d042c1) }

var fileDescriptor_5ffd045dd4d042c1 = []byte{
	// 253 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0x31, 0x4e, 0xc3, 0x30,
	0x14, 0x86, 0x95, 0x16, 0x25, 0xe1, 0x29, 0x95, 0x90, 0x19, 0xb0, 0x2c, 0x86, 0x2a, 0x53, 0xc5,
	0x90, 0x08, 0xd8, 0x18, 0x29, 0x1b, 0x0c, 0x95, 0x73, 0x01, 0x1e, 0xcd, 0x53, 0x88, 0x68, 0xf2,
	0x82, 0x6d, 0x2a, 0xb1, 0x72, 0x05, 0xae, 0xc0, 0x8d, 0xb8, 0x02, 0x07, 0x41, 0x76, 0x52, 0x16,
	0xb6, 0xef, 0xf7, 0x7b, 0xfa, 0xed, 0xcf, 0x90, 0xd9, 0xb6, 0x1b, 0x76, 0x54, 0x0c, 0x86, 0x1d,
	0x8b, 0x78, 0x4c, 0xea, 0xbc, 0x61, 0x6e, 0x76, 0x54, 0xe2, 0xd0, 0x96, 0xd8, 0xf7, 0xec, 0xd0,
	0xb5, 0xdc, 0xdb, 0x71, 0x2b, 0xff, 0x8a, 0x20, 0xd9, 0xe0, 0xf6, 0x05, 0x1b, 0x52, 0xf7, 0x90,
	0x68, 0x7a, 0x7d, 0x23, 0xeb, 0xc4, 0x09, 0xcc, 0xd7, 0x5d, 0x2d, 0xa3, 0x65, 0xb4, 0x5a, 0x68,
	0x8f, 0x42, 0x41, 0x6a, 0xc8, 0x92, 0xd9, 0x53, 0x2d, 0x67, 0xe1, 0xf8, 0x2f, 0x0b, 0x01, 0x47,
	0xb7, 0x5c, 0xbf, 0xcb, 0xf9, 0x32, 0x5a, 0x65, 0x3a, 0xb0, 0x7a, 0x80, 0x54, 0x93, 0x1d, 0xb8,
	0xb7, 0xe4, 0xe7, 0x6b, 0xae, 0x69, 0xaa, 0x0b, 0x2c, 0x24, 0x24, 0x1d, 0x59, 0x8b, 0x0d, 0x85,
	0xba, 0x63, 0x7d, 0x88, 0x7e, 0xfb, 0x0e, 0x1d, 0x1e, 0xda, 0x3c, 0x5f, 0x3d, 0x42, 0x56, 0x05,
	0x9d, 0xca, 0xdf, 0x68, 0xc4, 0x06, 0xe2, 0xca, 0x19, 0xc2, 0x4e, 0x9c, 0x15, 0x93, 0xf5, 0x64,
	0x51, 0x4c, 0x0a, 0x4a, 0xfe, 0x1f, 0x8c, 0xcf, 0xc9, 0x4f, 0x3f, 0xbe, 0x7f, 0x3e, 0x67, 0x8b,
	0x3c, 0x2d, 0xf7, 0x97, 0x25, 0x6d, 0x9f, 0xf9, 0x26, 0xba, 0x78, 0x8a, 0xc3, 0x7f, 0x5c, 0xff,
	0x06, 0x00, 0x00, 0xff, 0xff, 0x06, 0xd2, 0x16, 0x47, 0x45, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SimpleServerClient is the client API for SimpleServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SimpleServerClient interface {
	Stream(ctx context.Context, in *Package_Request, opts ...grpc.CallOption) (*Package_Response, error)
}

type simpleServerClient struct {
	cc *grpc.ClientConn
}

func NewSimpleServerClient(cc *grpc.ClientConn) SimpleServerClient {
	return &simpleServerClient{cc}
}

func (c *simpleServerClient) Stream(ctx context.Context, in *Package_Request, opts ...grpc.CallOption) (*Package_Response, error) {
	out := new(Package_Response)
	err := c.cc.Invoke(ctx, "/simple.SimpleServer/Stream", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SimpleServerServer is the server API for SimpleServer service.
type SimpleServerServer interface {
	Stream(context.Context, *Package_Request) (*Package_Response, error)
}

func RegisterSimpleServerServer(s *grpc.Server, srv SimpleServerServer) {
	s.RegisterService(&_SimpleServer_serviceDesc, srv)
}

func _SimpleServer_Stream_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Package_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SimpleServerServer).Stream(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/simple.SimpleServer/Stream",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SimpleServerServer).Stream(ctx, req.(*Package_Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _SimpleServer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "simple.SimpleServer",
	HandlerType: (*SimpleServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Stream",
			Handler:    _SimpleServer_Stream_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "simple.proto",
}
