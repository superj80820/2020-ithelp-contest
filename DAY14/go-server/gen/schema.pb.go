// Code generated by protoc-gen-go. DO NOT EDIT.
// source: schema.proto

package digimon

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type CreateRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c5fb4d8cc22d66a, []int{0}
}

func (m *CreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRequest.Unmarshal(m, b)
}
func (m *CreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRequest.Marshal(b, m, deterministic)
}
func (m *CreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequest.Merge(m, src)
}
func (m *CreateRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRequest.Size(m)
}
func (m *CreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRequest proto.InternalMessageInfo

func (m *CreateRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type QueryRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryRequest) Reset()         { *m = QueryRequest{} }
func (m *QueryRequest) String() string { return proto.CompactTextString(m) }
func (*QueryRequest) ProtoMessage()    {}
func (*QueryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c5fb4d8cc22d66a, []int{1}
}

func (m *QueryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryRequest.Unmarshal(m, b)
}
func (m *QueryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryRequest.Marshal(b, m, deterministic)
}
func (m *QueryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryRequest.Merge(m, src)
}
func (m *QueryRequest) XXX_Size() int {
	return xxx_messageInfo_QueryRequest.Size(m)
}
func (m *QueryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryRequest proto.InternalMessageInfo

func (m *QueryRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type FosterRequest struct {
	Id                   string              `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Food                 *FosterRequest_Food `protobuf:"bytes,2,opt,name=food,proto3" json:"food,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *FosterRequest) Reset()         { *m = FosterRequest{} }
func (m *FosterRequest) String() string { return proto.CompactTextString(m) }
func (*FosterRequest) ProtoMessage()    {}
func (*FosterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c5fb4d8cc22d66a, []int{2}
}

func (m *FosterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FosterRequest.Unmarshal(m, b)
}
func (m *FosterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FosterRequest.Marshal(b, m, deterministic)
}
func (m *FosterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FosterRequest.Merge(m, src)
}
func (m *FosterRequest) XXX_Size() int {
	return xxx_messageInfo_FosterRequest.Size(m)
}
func (m *FosterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FosterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FosterRequest proto.InternalMessageInfo

func (m *FosterRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *FosterRequest) GetFood() *FosterRequest_Food {
	if m != nil {
		return m.Food
	}
	return nil
}

type FosterRequest_Food struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FosterRequest_Food) Reset()         { *m = FosterRequest_Food{} }
func (m *FosterRequest_Food) String() string { return proto.CompactTextString(m) }
func (*FosterRequest_Food) ProtoMessage()    {}
func (*FosterRequest_Food) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c5fb4d8cc22d66a, []int{2, 0}
}

func (m *FosterRequest_Food) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FosterRequest_Food.Unmarshal(m, b)
}
func (m *FosterRequest_Food) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FosterRequest_Food.Marshal(b, m, deterministic)
}
func (m *FosterRequest_Food) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FosterRequest_Food.Merge(m, src)
}
func (m *FosterRequest_Food) XXX_Size() int {
	return xxx_messageInfo_FosterRequest_Food.Size(m)
}
func (m *FosterRequest_Food) XXX_DiscardUnknown() {
	xxx_messageInfo_FosterRequest_Food.DiscardUnknown(m)
}

var xxx_messageInfo_FosterRequest_Food proto.InternalMessageInfo

func (m *FosterRequest_Food) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type CreateResponse struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Status               string   `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateResponse) Reset()         { *m = CreateResponse{} }
func (m *CreateResponse) String() string { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()    {}
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c5fb4d8cc22d66a, []int{3}
}

func (m *CreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateResponse.Unmarshal(m, b)
}
func (m *CreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateResponse.Marshal(b, m, deterministic)
}
func (m *CreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateResponse.Merge(m, src)
}
func (m *CreateResponse) XXX_Size() int {
	return xxx_messageInfo_CreateResponse.Size(m)
}
func (m *CreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateResponse proto.InternalMessageInfo

func (m *CreateResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *CreateResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type QueryResponse struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Status               string   `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	Location             string   `protobuf:"bytes,4,opt,name=location,proto3" json:"location,omitempty"`
	Weather              string   `protobuf:"bytes,5,opt,name=weather,proto3" json:"weather,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryResponse) Reset()         { *m = QueryResponse{} }
func (m *QueryResponse) String() string { return proto.CompactTextString(m) }
func (*QueryResponse) ProtoMessage()    {}
func (*QueryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c5fb4d8cc22d66a, []int{4}
}

func (m *QueryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryResponse.Unmarshal(m, b)
}
func (m *QueryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryResponse.Marshal(b, m, deterministic)
}
func (m *QueryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryResponse.Merge(m, src)
}
func (m *QueryResponse) XXX_Size() int {
	return xxx_messageInfo_QueryResponse.Size(m)
}
func (m *QueryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryResponse proto.InternalMessageInfo

func (m *QueryResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *QueryResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *QueryResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *QueryResponse) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

func (m *QueryResponse) GetWeather() string {
	if m != nil {
		return m.Weather
	}
	return ""
}

type FosterResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FosterResponse) Reset()         { *m = FosterResponse{} }
func (m *FosterResponse) String() string { return proto.CompactTextString(m) }
func (*FosterResponse) ProtoMessage()    {}
func (*FosterResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c5fb4d8cc22d66a, []int{5}
}

func (m *FosterResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FosterResponse.Unmarshal(m, b)
}
func (m *FosterResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FosterResponse.Marshal(b, m, deterministic)
}
func (m *FosterResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FosterResponse.Merge(m, src)
}
func (m *FosterResponse) XXX_Size() int {
	return xxx_messageInfo_FosterResponse.Size(m)
}
func (m *FosterResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FosterResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FosterResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*CreateRequest)(nil), "digimon.CreateRequest")
	proto.RegisterType((*QueryRequest)(nil), "digimon.QueryRequest")
	proto.RegisterType((*FosterRequest)(nil), "digimon.FosterRequest")
	proto.RegisterType((*FosterRequest_Food)(nil), "digimon.FosterRequest.Food")
	proto.RegisterType((*CreateResponse)(nil), "digimon.CreateResponse")
	proto.RegisterType((*QueryResponse)(nil), "digimon.QueryResponse")
	proto.RegisterType((*FosterResponse)(nil), "digimon.FosterResponse")
}

func init() { proto.RegisterFile("schema.proto", fileDescriptor_1c5fb4d8cc22d66a) }

var fileDescriptor_1c5fb4d8cc22d66a = []byte{
	// 306 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x52, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0xad, 0x43, 0x68, 0xe1, 0xda, 0x44, 0xc8, 0x12, 0xc1, 0x0a, 0x12, 0xaa, 0xcc, 0xd2, 0x29,
	0xa0, 0x32, 0xb2, 0x20, 0x81, 0x3a, 0xb1, 0x10, 0xbe, 0xc0, 0x34, 0x07, 0x8d, 0xd4, 0xc4, 0xc5,
	0x76, 0x84, 0x18, 0xf9, 0x38, 0xfe, 0x0b, 0xd5, 0x89, 0x43, 0x42, 0x61, 0x62, 0xcb, 0x9d, 0xdf,
	0xbb, 0xf7, 0xde, 0x5d, 0x60, 0xa2, 0x97, 0x2b, 0x2c, 0x44, 0xb2, 0x51, 0xd2, 0x48, 0x3a, 0xca,
	0xf2, 0x97, 0xbc, 0x90, 0x25, 0x3f, 0x87, 0xe0, 0x56, 0xa1, 0x30, 0x98, 0xe2, 0x6b, 0x85, 0xda,
	0x50, 0x0a, 0x7e, 0x29, 0x0a, 0x64, 0x64, 0x4a, 0x66, 0x87, 0xa9, 0xfd, 0xe6, 0x67, 0x30, 0x79,
	0xa8, 0x50, 0xbd, 0x3b, 0x4c, 0x08, 0x5e, 0x9e, 0x35, 0x08, 0x2f, 0xcf, 0xf8, 0x1a, 0x82, 0x85,
	0xd4, 0x06, 0xd5, 0x1f, 0x00, 0x7a, 0x01, 0xfe, 0xb3, 0x94, 0x19, 0xf3, 0xa6, 0x64, 0x36, 0x9e,
	0x9f, 0x26, 0x8d, 0x7a, 0xd2, 0x63, 0x25, 0x0b, 0x29, 0xb3, 0xd4, 0x02, 0xe3, 0x18, 0xfc, 0x6d,
	0xf5, 0xab, 0x9b, 0x7b, 0x08, 0x9d, 0x65, 0xbd, 0x91, 0xa5, 0xc6, 0x1d, 0x39, 0xc7, 0xf2, 0xbe,
	0x59, 0x34, 0x82, 0xa1, 0x36, 0xc2, 0x54, 0x9a, 0xed, 0xd9, 0x6e, 0x53, 0xf1, 0x0f, 0x02, 0x41,
	0x13, 0xee, 0xff, 0xd3, 0x68, 0x0c, 0x07, 0x6b, 0xb9, 0x14, 0x26, 0x97, 0x25, 0xf3, 0xed, 0x4b,
	0x5b, 0x53, 0x06, 0xa3, 0x37, 0x14, 0x66, 0x85, 0x8a, 0xed, 0xdb, 0x27, 0x57, 0xf2, 0x23, 0x08,
	0xdd, 0x26, 0x6a, 0x0f, 0xf3, 0x4f, 0x02, 0xa3, 0xbb, 0x7a, 0x49, 0xf4, 0x1a, 0x86, 0x75, 0x5e,
	0x1a, 0xb5, 0x8b, 0xeb, 0xdd, 0x2c, 0x3e, 0xd9, 0xe9, 0xd7, 0x63, 0xf8, 0x80, 0xde, 0xc0, 0xd8,
	0xa6, 0x7b, 0x34, 0x0a, 0x45, 0x41, 0x8f, 0x5b, 0x64, 0xf7, 0xa0, 0x71, 0xf4, 0xb3, 0xed, 0xf8,
	0x97, 0x64, 0x2b, 0x5f, 0x9b, 0xeb, 0xc8, 0xf7, 0xee, 0xd6, 0x91, 0xef, 0xa7, 0xe0, 0x83, 0xa7,
	0xa1, 0xfd, 0xdd, 0xae, 0xbe, 0x02, 0x00, 0x00, 0xff, 0xff, 0xd4, 0x5e, 0x27, 0x99, 0x7e, 0x02,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// DigimonClient is the client API for Digimon service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DigimonClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	QueryStream(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (Digimon_QueryStreamClient, error)
	Foster(ctx context.Context, in *FosterRequest, opts ...grpc.CallOption) (*FosterResponse, error)
}

type digimonClient struct {
	cc grpc.ClientConnInterface
}

func NewDigimonClient(cc grpc.ClientConnInterface) DigimonClient {
	return &digimonClient{cc}
}

func (c *digimonClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/digimon.Digimon/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *digimonClient) QueryStream(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (Digimon_QueryStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Digimon_serviceDesc.Streams[0], "/digimon.Digimon/QueryStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &digimonQueryStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Digimon_QueryStreamClient interface {
	Recv() (*QueryResponse, error)
	grpc.ClientStream
}

type digimonQueryStreamClient struct {
	grpc.ClientStream
}

func (x *digimonQueryStreamClient) Recv() (*QueryResponse, error) {
	m := new(QueryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *digimonClient) Foster(ctx context.Context, in *FosterRequest, opts ...grpc.CallOption) (*FosterResponse, error) {
	out := new(FosterResponse)
	err := c.cc.Invoke(ctx, "/digimon.Digimon/Foster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DigimonServer is the server API for Digimon service.
type DigimonServer interface {
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	QueryStream(*QueryRequest, Digimon_QueryStreamServer) error
	Foster(context.Context, *FosterRequest) (*FosterResponse, error)
}

// UnimplementedDigimonServer can be embedded to have forward compatible implementations.
type UnimplementedDigimonServer struct {
}

func (*UnimplementedDigimonServer) Create(ctx context.Context, req *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedDigimonServer) QueryStream(req *QueryRequest, srv Digimon_QueryStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method QueryStream not implemented")
}
func (*UnimplementedDigimonServer) Foster(ctx context.Context, req *FosterRequest) (*FosterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Foster not implemented")
}

func RegisterDigimonServer(s *grpc.Server, srv DigimonServer) {
	s.RegisterService(&_Digimon_serviceDesc, srv)
}

func _Digimon_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DigimonServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/digimon.Digimon/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DigimonServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Digimon_QueryStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(QueryRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DigimonServer).QueryStream(m, &digimonQueryStreamServer{stream})
}

type Digimon_QueryStreamServer interface {
	Send(*QueryResponse) error
	grpc.ServerStream
}

type digimonQueryStreamServer struct {
	grpc.ServerStream
}

func (x *digimonQueryStreamServer) Send(m *QueryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Digimon_Foster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FosterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DigimonServer).Foster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/digimon.Digimon/Foster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DigimonServer).Foster(ctx, req.(*FosterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Digimon_serviceDesc = grpc.ServiceDesc{
	ServiceName: "digimon.Digimon",
	HandlerType: (*DigimonServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Digimon_Create_Handler,
		},
		{
			MethodName: "Foster",
			Handler:    _Digimon_Foster_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "QueryStream",
			Handler:       _Digimon_QueryStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "schema.proto",
}
