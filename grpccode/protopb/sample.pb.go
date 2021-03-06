// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sample.proto

package person

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

type Person struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Person) Reset()         { *m = Person{} }
func (m *Person) String() string { return proto.CompactTextString(m) }
func (*Person) ProtoMessage()    {}
func (*Person) Descriptor() ([]byte, []int) {
	return fileDescriptor_2141552de9bf11d0, []int{0}
}

func (m *Person) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Person.Unmarshal(m, b)
}
func (m *Person) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Person.Marshal(b, m, deterministic)
}
func (m *Person) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Person.Merge(m, src)
}
func (m *Person) XXX_Size() int {
	return xxx_messageInfo_Person.Size(m)
}
func (m *Person) XXX_DiscardUnknown() {
	xxx_messageInfo_Person.DiscardUnknown(m)
}

var xxx_messageInfo_Person proto.InternalMessageInfo

func (m *Person) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Person) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type PersonRequest struct {
	Person               *Person  `protobuf:"bytes,1,opt,name=person,proto3" json:"person,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PersonRequest) Reset()         { *m = PersonRequest{} }
func (m *PersonRequest) String() string { return proto.CompactTextString(m) }
func (*PersonRequest) ProtoMessage()    {}
func (*PersonRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2141552de9bf11d0, []int{1}
}

func (m *PersonRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PersonRequest.Unmarshal(m, b)
}
func (m *PersonRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PersonRequest.Marshal(b, m, deterministic)
}
func (m *PersonRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PersonRequest.Merge(m, src)
}
func (m *PersonRequest) XXX_Size() int {
	return xxx_messageInfo_PersonRequest.Size(m)
}
func (m *PersonRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PersonRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PersonRequest proto.InternalMessageInfo

func (m *PersonRequest) GetPerson() *Person {
	if m != nil {
		return m.Person
	}
	return nil
}

type PersonResponse struct {
	Result               string   `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PersonResponse) Reset()         { *m = PersonResponse{} }
func (m *PersonResponse) String() string { return proto.CompactTextString(m) }
func (*PersonResponse) ProtoMessage()    {}
func (*PersonResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2141552de9bf11d0, []int{2}
}

func (m *PersonResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PersonResponse.Unmarshal(m, b)
}
func (m *PersonResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PersonResponse.Marshal(b, m, deterministic)
}
func (m *PersonResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PersonResponse.Merge(m, src)
}
func (m *PersonResponse) XXX_Size() int {
	return xxx_messageInfo_PersonResponse.Size(m)
}
func (m *PersonResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PersonResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PersonResponse proto.InternalMessageInfo

func (m *PersonResponse) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

type PersonStreamingRequest struct {
	Person               *Person  `protobuf:"bytes,1,opt,name=person,proto3" json:"person,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PersonStreamingRequest) Reset()         { *m = PersonStreamingRequest{} }
func (m *PersonStreamingRequest) String() string { return proto.CompactTextString(m) }
func (*PersonStreamingRequest) ProtoMessage()    {}
func (*PersonStreamingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2141552de9bf11d0, []int{3}
}

func (m *PersonStreamingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PersonStreamingRequest.Unmarshal(m, b)
}
func (m *PersonStreamingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PersonStreamingRequest.Marshal(b, m, deterministic)
}
func (m *PersonStreamingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PersonStreamingRequest.Merge(m, src)
}
func (m *PersonStreamingRequest) XXX_Size() int {
	return xxx_messageInfo_PersonStreamingRequest.Size(m)
}
func (m *PersonStreamingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PersonStreamingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PersonStreamingRequest proto.InternalMessageInfo

func (m *PersonStreamingRequest) GetPerson() *Person {
	if m != nil {
		return m.Person
	}
	return nil
}

type PersonStreamingResponse struct {
	Result               string   `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PersonStreamingResponse) Reset()         { *m = PersonStreamingResponse{} }
func (m *PersonStreamingResponse) String() string { return proto.CompactTextString(m) }
func (*PersonStreamingResponse) ProtoMessage()    {}
func (*PersonStreamingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2141552de9bf11d0, []int{4}
}

func (m *PersonStreamingResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PersonStreamingResponse.Unmarshal(m, b)
}
func (m *PersonStreamingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PersonStreamingResponse.Marshal(b, m, deterministic)
}
func (m *PersonStreamingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PersonStreamingResponse.Merge(m, src)
}
func (m *PersonStreamingResponse) XXX_Size() int {
	return xxx_messageInfo_PersonStreamingResponse.Size(m)
}
func (m *PersonStreamingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PersonStreamingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PersonStreamingResponse proto.InternalMessageInfo

func (m *PersonStreamingResponse) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

type PersonClientStreamingRequest struct {
	Person               *Person  `protobuf:"bytes,1,opt,name=person,proto3" json:"person,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PersonClientStreamingRequest) Reset()         { *m = PersonClientStreamingRequest{} }
func (m *PersonClientStreamingRequest) String() string { return proto.CompactTextString(m) }
func (*PersonClientStreamingRequest) ProtoMessage()    {}
func (*PersonClientStreamingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2141552de9bf11d0, []int{5}
}

func (m *PersonClientStreamingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PersonClientStreamingRequest.Unmarshal(m, b)
}
func (m *PersonClientStreamingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PersonClientStreamingRequest.Marshal(b, m, deterministic)
}
func (m *PersonClientStreamingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PersonClientStreamingRequest.Merge(m, src)
}
func (m *PersonClientStreamingRequest) XXX_Size() int {
	return xxx_messageInfo_PersonClientStreamingRequest.Size(m)
}
func (m *PersonClientStreamingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PersonClientStreamingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PersonClientStreamingRequest proto.InternalMessageInfo

func (m *PersonClientStreamingRequest) GetPerson() *Person {
	if m != nil {
		return m.Person
	}
	return nil
}

type PersonClientStreamingResponse struct {
	Result               string   `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PersonClientStreamingResponse) Reset()         { *m = PersonClientStreamingResponse{} }
func (m *PersonClientStreamingResponse) String() string { return proto.CompactTextString(m) }
func (*PersonClientStreamingResponse) ProtoMessage()    {}
func (*PersonClientStreamingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2141552de9bf11d0, []int{6}
}

func (m *PersonClientStreamingResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PersonClientStreamingResponse.Unmarshal(m, b)
}
func (m *PersonClientStreamingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PersonClientStreamingResponse.Marshal(b, m, deterministic)
}
func (m *PersonClientStreamingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PersonClientStreamingResponse.Merge(m, src)
}
func (m *PersonClientStreamingResponse) XXX_Size() int {
	return xxx_messageInfo_PersonClientStreamingResponse.Size(m)
}
func (m *PersonClientStreamingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PersonClientStreamingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PersonClientStreamingResponse proto.InternalMessageInfo

func (m *PersonClientStreamingResponse) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

type PersonClientserverStreamingRequest struct {
	Person               *Person  `protobuf:"bytes,1,opt,name=person,proto3" json:"person,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PersonClientserverStreamingRequest) Reset()         { *m = PersonClientserverStreamingRequest{} }
func (m *PersonClientserverStreamingRequest) String() string { return proto.CompactTextString(m) }
func (*PersonClientserverStreamingRequest) ProtoMessage()    {}
func (*PersonClientserverStreamingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2141552de9bf11d0, []int{7}
}

func (m *PersonClientserverStreamingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PersonClientserverStreamingRequest.Unmarshal(m, b)
}
func (m *PersonClientserverStreamingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PersonClientserverStreamingRequest.Marshal(b, m, deterministic)
}
func (m *PersonClientserverStreamingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PersonClientserverStreamingRequest.Merge(m, src)
}
func (m *PersonClientserverStreamingRequest) XXX_Size() int {
	return xxx_messageInfo_PersonClientserverStreamingRequest.Size(m)
}
func (m *PersonClientserverStreamingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PersonClientserverStreamingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PersonClientserverStreamingRequest proto.InternalMessageInfo

func (m *PersonClientserverStreamingRequest) GetPerson() *Person {
	if m != nil {
		return m.Person
	}
	return nil
}

type PersonClientServerStreamingResponse struct {
	Result               string   `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PersonClientServerStreamingResponse) Reset()         { *m = PersonClientServerStreamingResponse{} }
func (m *PersonClientServerStreamingResponse) String() string { return proto.CompactTextString(m) }
func (*PersonClientServerStreamingResponse) ProtoMessage()    {}
func (*PersonClientServerStreamingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2141552de9bf11d0, []int{8}
}

func (m *PersonClientServerStreamingResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PersonClientServerStreamingResponse.Unmarshal(m, b)
}
func (m *PersonClientServerStreamingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PersonClientServerStreamingResponse.Marshal(b, m, deterministic)
}
func (m *PersonClientServerStreamingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PersonClientServerStreamingResponse.Merge(m, src)
}
func (m *PersonClientServerStreamingResponse) XXX_Size() int {
	return xxx_messageInfo_PersonClientServerStreamingResponse.Size(m)
}
func (m *PersonClientServerStreamingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PersonClientServerStreamingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PersonClientServerStreamingResponse proto.InternalMessageInfo

func (m *PersonClientServerStreamingResponse) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

func init() {
	proto.RegisterType((*Person)(nil), "person.person")
	proto.RegisterType((*PersonRequest)(nil), "person.PersonRequest")
	proto.RegisterType((*PersonResponse)(nil), "person.PersonResponse")
	proto.RegisterType((*PersonStreamingRequest)(nil), "person.PersonStreamingRequest")
	proto.RegisterType((*PersonStreamingResponse)(nil), "person.PersonStreamingResponse")
	proto.RegisterType((*PersonClientStreamingRequest)(nil), "person.PersonClientStreamingRequest")
	proto.RegisterType((*PersonClientStreamingResponse)(nil), "person.PersonClientStreamingResponse")
	proto.RegisterType((*PersonClientserverStreamingRequest)(nil), "person.PersonClientserverStreamingRequest")
	proto.RegisterType((*PersonClientServerStreamingResponse)(nil), "person.PersonClientServerStreamingResponse")
}

func init() { proto.RegisterFile("sample.proto", fileDescriptor_2141552de9bf11d0) }

var fileDescriptor_2141552de9bf11d0 = []byte{
	// 326 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0xcb, 0x4b, 0xf3, 0x40,
	0x14, 0xc5, 0xbf, 0x09, 0x9f, 0xc5, 0x5e, 0x6b, 0x17, 0x17, 0x8c, 0x25, 0xf8, 0x28, 0xf1, 0x41,
	0x50, 0x29, 0xb5, 0x2e, 0xba, 0x12, 0xd4, 0x82, 0x6e, 0x5c, 0x94, 0xb8, 0x73, 0x23, 0xd1, 0x5e,
	0x64, 0x20, 0x99, 0xc4, 0x99, 0xa9, 0x7f, 0xbc, 0x2b, 0xe9, 0x24, 0x03, 0x26, 0x4d, 0x0d, 0x76,
	0x95, 0x99, 0xc9, 0x3d, 0xe7, 0x77, 0x92, 0x93, 0x40, 0x47, 0x45, 0x49, 0x16, 0xd3, 0x20, 0x93,
	0xa9, 0x4e, 0xb1, 0x95, 0x91, 0x54, 0xa9, 0xf0, 0x2f, 0xa0, 0x58, 0x61, 0x17, 0x1c, 0x3e, 0xeb,
	0xb1, 0x3e, 0x0b, 0x36, 0x42, 0x87, 0xcf, 0x10, 0xe1, 0xbf, 0x88, 0x12, 0xea, 0x39, 0x7d, 0x16,
	0xb4, 0x43, 0xb3, 0xf6, 0xc7, 0xb0, 0x3d, 0x35, 0xd3, 0x21, 0x7d, 0xcc, 0x49, 0x69, 0x3c, 0xb5,
	0x72, 0x23, 0xdc, 0x1a, 0x75, 0x07, 0xf9, 0xb6, 0xb8, 0x84, 0x16, 0x13, 0x40, 0xd7, 0x0a, 0x55,
	0x96, 0x0a, 0x45, 0xe8, 0x42, 0x4b, 0x92, 0x9a, 0xc7, 0xda, 0x28, 0xdb, 0x61, 0xb1, 0xf3, 0x6f,
	0xc0, 0xcd, 0x27, 0x9f, 0xb4, 0xa4, 0x28, 0xe1, 0xe2, 0xfd, 0xaf, 0xac, 0x4b, 0xd8, 0x5d, 0x72,
	0x68, 0x80, 0xde, 0xc3, 0x5e, 0x2e, 0x99, 0xc4, 0x9c, 0x84, 0x5e, 0x1b, 0x3d, 0x86, 0xfd, 0x15,
	0x3e, 0x0d, 0x01, 0x1e, 0xc1, 0xff, 0x29, 0x54, 0x24, 0x3f, 0x49, 0xae, 0x1d, 0xe3, 0x1a, 0x8e,
	0x4a, 0x31, 0xaa, 0x6e, 0xbf, 0x87, 0x19, 0x7d, 0x39, 0xb6, 0xe6, 0x85, 0x92, 0xbf, 0x11, 0xde,
	0xda, 0xfa, 0x26, 0x92, 0x22, 0xcd, 0x53, 0x81, 0x3b, 0x16, 0x5d, 0xfa, 0x1e, 0x3c, 0xb7, 0x7a,
	0x9c, 0xa3, 0xfc, 0x7f, 0x38, 0x85, 0xf6, 0x03, 0xe9, 0xfc, 0x18, 0x0f, 0xca, 0x63, 0xd5, 0x07,
	0xf5, 0x0e, 0x57, 0xde, 0xb7, 0x7e, 0x43, 0x86, 0x2f, 0xd0, 0x31, 0x71, 0xa8, 0x30, 0x3d, 0x2e,
	0x8b, 0xea, 0xab, 0xf4, 0x4e, 0x1a, 0xa6, 0x2c, 0x20, 0x60, 0xa8, 0x00, 0x97, 0x5f, 0x20, 0x9e,
	0xd5, 0x19, 0xd4, 0x17, 0xe6, 0x9d, 0xd7, 0xc2, 0xea, 0xeb, 0x58, 0x20, 0x87, 0xec, 0x6e, 0xf3,
	0xb9, 0x68, 0xf1, 0xb5, 0x65, 0xfe, 0xd4, 0xab, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc5, 0x5a,
	0xb4, 0x03, 0xb9, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// PersonServiceClient is the client API for PersonService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PersonServiceClient interface {
	PersonCreation(ctx context.Context, in *PersonRequest, opts ...grpc.CallOption) (*PersonResponse, error)
	GetPerson(ctx context.Context, in *PersonStreamingRequest, opts ...grpc.CallOption) (PersonService_GetPersonClient, error)
	CreatePerson(ctx context.Context, opts ...grpc.CallOption) (PersonService_CreatePersonClient, error)
	ClientServerStream(ctx context.Context, opts ...grpc.CallOption) (PersonService_ClientServerStreamClient, error)
}

type personServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPersonServiceClient(cc grpc.ClientConnInterface) PersonServiceClient {
	return &personServiceClient{cc}
}

func (c *personServiceClient) PersonCreation(ctx context.Context, in *PersonRequest, opts ...grpc.CallOption) (*PersonResponse, error) {
	out := new(PersonResponse)
	err := c.cc.Invoke(ctx, "/person.PersonService/PersonCreation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *personServiceClient) GetPerson(ctx context.Context, in *PersonStreamingRequest, opts ...grpc.CallOption) (PersonService_GetPersonClient, error) {
	stream, err := c.cc.NewStream(ctx, &_PersonService_serviceDesc.Streams[0], "/person.PersonService/GetPerson", opts...)
	if err != nil {
		return nil, err
	}
	x := &personServiceGetPersonClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PersonService_GetPersonClient interface {
	Recv() (*PersonStreamingResponse, error)
	grpc.ClientStream
}

type personServiceGetPersonClient struct {
	grpc.ClientStream
}

func (x *personServiceGetPersonClient) Recv() (*PersonStreamingResponse, error) {
	m := new(PersonStreamingResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *personServiceClient) CreatePerson(ctx context.Context, opts ...grpc.CallOption) (PersonService_CreatePersonClient, error) {
	stream, err := c.cc.NewStream(ctx, &_PersonService_serviceDesc.Streams[1], "/person.PersonService/CreatePerson", opts...)
	if err != nil {
		return nil, err
	}
	x := &personServiceCreatePersonClient{stream}
	return x, nil
}

type PersonService_CreatePersonClient interface {
	Send(*PersonClientStreamingRequest) error
	CloseAndRecv() (*PersonClientStreamingResponse, error)
	grpc.ClientStream
}

type personServiceCreatePersonClient struct {
	grpc.ClientStream
}

func (x *personServiceCreatePersonClient) Send(m *PersonClientStreamingRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *personServiceCreatePersonClient) CloseAndRecv() (*PersonClientStreamingResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(PersonClientStreamingResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *personServiceClient) ClientServerStream(ctx context.Context, opts ...grpc.CallOption) (PersonService_ClientServerStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_PersonService_serviceDesc.Streams[2], "/person.PersonService/ClientServerStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &personServiceClientServerStreamClient{stream}
	return x, nil
}

type PersonService_ClientServerStreamClient interface {
	Send(*PersonClientserverStreamingRequest) error
	Recv() (*PersonClientServerStreamingResponse, error)
	grpc.ClientStream
}

type personServiceClientServerStreamClient struct {
	grpc.ClientStream
}

func (x *personServiceClientServerStreamClient) Send(m *PersonClientserverStreamingRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *personServiceClientServerStreamClient) Recv() (*PersonClientServerStreamingResponse, error) {
	m := new(PersonClientServerStreamingResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PersonServiceServer is the server API for PersonService service.
type PersonServiceServer interface {
	PersonCreation(context.Context, *PersonRequest) (*PersonResponse, error)
	GetPerson(*PersonStreamingRequest, PersonService_GetPersonServer) error
	CreatePerson(PersonService_CreatePersonServer) error
	ClientServerStream(PersonService_ClientServerStreamServer) error
}

// UnimplementedPersonServiceServer can be embedded to have forward compatible implementations.
type UnimplementedPersonServiceServer struct {
}

func (*UnimplementedPersonServiceServer) PersonCreation(ctx context.Context, req *PersonRequest) (*PersonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PersonCreation not implemented")
}
func (*UnimplementedPersonServiceServer) GetPerson(req *PersonStreamingRequest, srv PersonService_GetPersonServer) error {
	return status.Errorf(codes.Unimplemented, "method GetPerson not implemented")
}
func (*UnimplementedPersonServiceServer) CreatePerson(srv PersonService_CreatePersonServer) error {
	return status.Errorf(codes.Unimplemented, "method CreatePerson not implemented")
}
func (*UnimplementedPersonServiceServer) ClientServerStream(srv PersonService_ClientServerStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method ClientServerStream not implemented")
}

func RegisterPersonServiceServer(s *grpc.Server, srv PersonServiceServer) {
	s.RegisterService(&_PersonService_serviceDesc, srv)
}

func _PersonService_PersonCreation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PersonRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PersonServiceServer).PersonCreation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/person.PersonService/PersonCreation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PersonServiceServer).PersonCreation(ctx, req.(*PersonRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PersonService_GetPerson_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PersonStreamingRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PersonServiceServer).GetPerson(m, &personServiceGetPersonServer{stream})
}

type PersonService_GetPersonServer interface {
	Send(*PersonStreamingResponse) error
	grpc.ServerStream
}

type personServiceGetPersonServer struct {
	grpc.ServerStream
}

func (x *personServiceGetPersonServer) Send(m *PersonStreamingResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _PersonService_CreatePerson_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PersonServiceServer).CreatePerson(&personServiceCreatePersonServer{stream})
}

type PersonService_CreatePersonServer interface {
	SendAndClose(*PersonClientStreamingResponse) error
	Recv() (*PersonClientStreamingRequest, error)
	grpc.ServerStream
}

type personServiceCreatePersonServer struct {
	grpc.ServerStream
}

func (x *personServiceCreatePersonServer) SendAndClose(m *PersonClientStreamingResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *personServiceCreatePersonServer) Recv() (*PersonClientStreamingRequest, error) {
	m := new(PersonClientStreamingRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _PersonService_ClientServerStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PersonServiceServer).ClientServerStream(&personServiceClientServerStreamServer{stream})
}

type PersonService_ClientServerStreamServer interface {
	Send(*PersonClientServerStreamingResponse) error
	Recv() (*PersonClientserverStreamingRequest, error)
	grpc.ServerStream
}

type personServiceClientServerStreamServer struct {
	grpc.ServerStream
}

func (x *personServiceClientServerStreamServer) Send(m *PersonClientServerStreamingResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *personServiceClientServerStreamServer) Recv() (*PersonClientserverStreamingRequest, error) {
	m := new(PersonClientserverStreamingRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _PersonService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "person.PersonService",
	HandlerType: (*PersonServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PersonCreation",
			Handler:    _PersonService_PersonCreation_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetPerson",
			Handler:       _PersonService_GetPerson_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "CreatePerson",
			Handler:       _PersonService_CreatePerson_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "ClientServerStream",
			Handler:       _PersonService_ClientServerStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "sample.proto",
}
