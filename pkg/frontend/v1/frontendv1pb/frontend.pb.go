// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: frontend.proto

// Protobuf package should not be changed when moving around go packages
// in order to not break backward compatibility.

package frontendv1pb

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	httpgrpc "github.com/grafana/dskit/httpgrpc"
	stats "github.com/grafana/mimir/pkg/querier/stats"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strconv "strconv"
	strings "strings"
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

type Type int32

const (
	HTTP_REQUEST Type = 0
	GET_ID       Type = 1
)

var Type_name = map[int32]string{
	0: "HTTP_REQUEST",
	1: "GET_ID",
}

var Type_value = map[string]int32{
	"HTTP_REQUEST": 0,
	"GET_ID":       1,
}

func (Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_eca3873955a29cfe, []int{0}
}

type FrontendToClient struct {
	HttpRequest *httpgrpc.HTTPRequest `protobuf:"bytes,1,opt,name=httpRequest,proto3" json:"httpRequest,omitempty"`
	Type        Type                  `protobuf:"varint,2,opt,name=type,proto3,enum=frontend.Type" json:"type,omitempty"`
	// Whether query statistics tracking should be enabled. The response will include
	// statistics only when this option is enabled.
	StatsEnabled bool `protobuf:"varint,3,opt,name=statsEnabled,proto3" json:"statsEnabled,omitempty"`
}

func (m *FrontendToClient) Reset()      { *m = FrontendToClient{} }
func (*FrontendToClient) ProtoMessage() {}
func (*FrontendToClient) Descriptor() ([]byte, []int) {
	return fileDescriptor_eca3873955a29cfe, []int{0}
}
func (m *FrontendToClient) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FrontendToClient) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FrontendToClient.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FrontendToClient) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FrontendToClient.Merge(m, src)
}
func (m *FrontendToClient) XXX_Size() int {
	return m.Size()
}
func (m *FrontendToClient) XXX_DiscardUnknown() {
	xxx_messageInfo_FrontendToClient.DiscardUnknown(m)
}

var xxx_messageInfo_FrontendToClient proto.InternalMessageInfo

func (m *FrontendToClient) GetHttpRequest() *httpgrpc.HTTPRequest {
	if m != nil {
		return m.HttpRequest
	}
	return nil
}

func (m *FrontendToClient) GetType() Type {
	if m != nil {
		return m.Type
	}
	return HTTP_REQUEST
}

func (m *FrontendToClient) GetStatsEnabled() bool {
	if m != nil {
		return m.StatsEnabled
	}
	return false
}

type ClientToFrontend struct {
	HttpResponse *httpgrpc.HTTPResponse `protobuf:"bytes,1,opt,name=httpResponse,proto3" json:"httpResponse,omitempty"`
	ClientID     string                 `protobuf:"bytes,2,opt,name=clientID,proto3" json:"clientID,omitempty"`
	Stats        *stats.Stats           `protobuf:"bytes,3,opt,name=stats,proto3" json:"stats,omitempty"`
}

func (m *ClientToFrontend) Reset()      { *m = ClientToFrontend{} }
func (*ClientToFrontend) ProtoMessage() {}
func (*ClientToFrontend) Descriptor() ([]byte, []int) {
	return fileDescriptor_eca3873955a29cfe, []int{1}
}
func (m *ClientToFrontend) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ClientToFrontend) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ClientToFrontend.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ClientToFrontend) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientToFrontend.Merge(m, src)
}
func (m *ClientToFrontend) XXX_Size() int {
	return m.Size()
}
func (m *ClientToFrontend) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientToFrontend.DiscardUnknown(m)
}

var xxx_messageInfo_ClientToFrontend proto.InternalMessageInfo

func (m *ClientToFrontend) GetHttpResponse() *httpgrpc.HTTPResponse {
	if m != nil {
		return m.HttpResponse
	}
	return nil
}

func (m *ClientToFrontend) GetClientID() string {
	if m != nil {
		return m.ClientID
	}
	return ""
}

func (m *ClientToFrontend) GetStats() *stats.Stats {
	if m != nil {
		return m.Stats
	}
	return nil
}

type NotifyClientShutdownRequest struct {
	ClientID string `protobuf:"bytes,1,opt,name=clientID,proto3" json:"clientID,omitempty"`
}

func (m *NotifyClientShutdownRequest) Reset()      { *m = NotifyClientShutdownRequest{} }
func (*NotifyClientShutdownRequest) ProtoMessage() {}
func (*NotifyClientShutdownRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_eca3873955a29cfe, []int{2}
}
func (m *NotifyClientShutdownRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NotifyClientShutdownRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NotifyClientShutdownRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NotifyClientShutdownRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NotifyClientShutdownRequest.Merge(m, src)
}
func (m *NotifyClientShutdownRequest) XXX_Size() int {
	return m.Size()
}
func (m *NotifyClientShutdownRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NotifyClientShutdownRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NotifyClientShutdownRequest proto.InternalMessageInfo

func (m *NotifyClientShutdownRequest) GetClientID() string {
	if m != nil {
		return m.ClientID
	}
	return ""
}

type NotifyClientShutdownResponse struct {
}

func (m *NotifyClientShutdownResponse) Reset()      { *m = NotifyClientShutdownResponse{} }
func (*NotifyClientShutdownResponse) ProtoMessage() {}
func (*NotifyClientShutdownResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_eca3873955a29cfe, []int{3}
}
func (m *NotifyClientShutdownResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NotifyClientShutdownResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NotifyClientShutdownResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NotifyClientShutdownResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NotifyClientShutdownResponse.Merge(m, src)
}
func (m *NotifyClientShutdownResponse) XXX_Size() int {
	return m.Size()
}
func (m *NotifyClientShutdownResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NotifyClientShutdownResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NotifyClientShutdownResponse proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("frontend.Type", Type_name, Type_value)
	proto.RegisterType((*FrontendToClient)(nil), "frontend.FrontendToClient")
	proto.RegisterType((*ClientToFrontend)(nil), "frontend.ClientToFrontend")
	proto.RegisterType((*NotifyClientShutdownRequest)(nil), "frontend.NotifyClientShutdownRequest")
	proto.RegisterType((*NotifyClientShutdownResponse)(nil), "frontend.NotifyClientShutdownResponse")
}

func init() { proto.RegisterFile("frontend.proto", fileDescriptor_eca3873955a29cfe) }

var fileDescriptor_eca3873955a29cfe = []byte{
	// 477 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x41, 0x6e, 0xd3, 0x40,
	0x14, 0x86, 0xfd, 0xa0, 0x94, 0xf0, 0x62, 0x45, 0xd6, 0xa8, 0xa0, 0xc8, 0xa0, 0x51, 0x64, 0x01,
	0x8a, 0x58, 0xd8, 0x25, 0x48, 0x20, 0x58, 0x96, 0x86, 0xd2, 0x0d, 0x2a, 0x8e, 0xd9, 0xb0, 0xa9,
	0x9c, 0x64, 0xe2, 0x58, 0x6d, 0x3c, 0xae, 0x67, 0x02, 0xca, 0x8e, 0x13, 0x20, 0x24, 0x2e, 0xc1,
	0x19, 0x38, 0x01, 0xcb, 0x2c, 0xbb, 0x24, 0xce, 0x86, 0x65, 0x8f, 0x80, 0x3c, 0x93, 0xb8, 0x49,
	0x14, 0xc1, 0x66, 0x34, 0x33, 0xff, 0xbc, 0xf7, 0xbe, 0xff, 0xb7, 0xb1, 0x36, 0xc8, 0x78, 0x22,
	0x59, 0xd2, 0x77, 0xd3, 0x8c, 0x4b, 0x4e, 0x2a, 0xcb, 0xb3, 0xbd, 0x17, 0xf1, 0x88, 0xab, 0x4b,
	0xaf, 0xd8, 0x69, 0xdd, 0xde, 0x8f, 0x62, 0x39, 0x1c, 0x77, 0xdd, 0x1e, 0x1f, 0x79, 0x51, 0x16,
	0x0e, 0xc2, 0x24, 0xf4, 0xfa, 0xe2, 0x2c, 0x96, 0xde, 0x50, 0xca, 0x34, 0xca, 0xd2, 0x5e, 0xb9,
	0x59, 0x54, 0x3c, 0xdf, 0x52, 0x31, 0x8a, 0x47, 0x71, 0xe6, 0xa5, 0x67, 0x91, 0x77, 0x31, 0x66,
	0x59, 0xcc, 0x32, 0x4f, 0xc8, 0x50, 0x0a, 0xbd, 0xea, 0x3a, 0xe7, 0x3b, 0xa0, 0xf5, 0x66, 0x01,
	0x13, 0xf0, 0xd7, 0xe7, 0x31, 0x4b, 0x24, 0x79, 0x81, 0xd5, 0xa2, 0xbd, 0xcf, 0x2e, 0xc6, 0x4c,
	0xc8, 0x3a, 0x34, 0xa0, 0x59, 0x6d, 0xdd, 0x75, 0xcb, 0x91, 0x6f, 0x83, 0xe0, 0x64, 0x21, 0xfa,
	0xab, 0x2f, 0x89, 0x83, 0x3b, 0x72, 0x92, 0xb2, 0xfa, 0x8d, 0x06, 0x34, 0x6b, 0xad, 0x9a, 0x5b,
	0xda, 0x0e, 0x26, 0x29, 0xf3, 0x95, 0x46, 0x1c, 0x34, 0x15, 0x40, 0x3b, 0x09, 0xbb, 0xe7, 0xac,
	0x5f, 0xbf, 0xd9, 0x80, 0x66, 0xc5, 0x5f, 0xbb, 0x73, 0xbe, 0x02, 0x5a, 0x9a, 0x25, 0xe0, 0x4b,
	0x3a, 0xf2, 0x0a, 0x4d, 0x3d, 0x4b, 0xa4, 0x3c, 0x11, 0x6c, 0x81, 0x75, 0x6f, 0x13, 0x4b, 0xab,
	0xfe, 0xda, 0x5b, 0x62, 0x63, 0xa5, 0xa7, 0xfa, 0x1d, 0x1f, 0x2a, 0xb8, 0x3b, 0x7e, 0x79, 0x26,
	0x0e, 0xde, 0x52, 0xc3, 0x15, 0x49, 0xb5, 0x65, 0xba, 0x3a, 0x9f, 0x4e, 0xb1, 0xfa, 0x5a, 0x72,
	0x5e, 0xe2, 0xfd, 0x77, 0x5c, 0xc6, 0x83, 0x89, 0xa6, 0xea, 0x0c, 0xc7, 0xb2, 0xcf, 0x3f, 0x27,
	0x4b, 0xdf, 0xab, 0xed, 0x61, 0xbd, 0xbd, 0x43, 0xf1, 0xc1, 0xf6, 0x52, 0x8d, 0xf6, 0xe4, 0x21,
	0xee, 0x14, 0xe9, 0x10, 0x0b, 0xcd, 0xc2, 0xc0, 0xa9, 0xdf, 0x7e, 0xff, 0xa1, 0xdd, 0x09, 0x2c,
	0x83, 0x20, 0xee, 0x1e, 0xb5, 0x83, 0xd3, 0xe3, 0x43, 0x0b, 0x5a, 0x3f, 0x01, 0x2b, 0x65, 0x12,
	0x47, 0x78, 0xfb, 0x24, 0xe3, 0x3d, 0x26, 0x04, 0xb1, 0xaf, 0x33, 0xde, 0x0c, 0xcc, 0x5e, 0xd1,
	0x36, 0x3f, 0xb1, 0x63, 0x34, 0x61, 0x1f, 0x08, 0xc3, 0xbd, 0x6d, 0x6c, 0xe4, 0xd1, 0x75, 0xe5,
	0x3f, 0x6c, 0xdb, 0x8f, 0xff, 0xf7, 0x4c, 0x5b, 0x3c, 0x38, 0x98, 0xce, 0xa8, 0x71, 0x39, 0xa3,
	0xc6, 0xd5, 0x8c, 0xc2, 0x97, 0x9c, 0xc2, 0x8f, 0x9c, 0xc2, 0xaf, 0x9c, 0xc2, 0x34, 0xa7, 0xf0,
	0x3b, 0xa7, 0xf0, 0x27, 0xa7, 0xc6, 0x55, 0x4e, 0xe1, 0xdb, 0x9c, 0x1a, 0xd3, 0x39, 0x35, 0x2e,
	0xe7, 0xd4, 0xf8, 0x68, 0x2e, 0x9b, 0x7f, 0x7a, 0x9a, 0x76, 0xbb, 0xbb, 0xea, 0x7f, 0x7d, 0xf6,
	0x37, 0x00, 0x00, 0xff, 0xff, 0x72, 0xa4, 0x89, 0x79, 0x4b, 0x03, 0x00, 0x00,
}

func (x Type) String() string {
	s, ok := Type_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (this *FrontendToClient) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*FrontendToClient)
	if !ok {
		that2, ok := that.(FrontendToClient)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.HttpRequest.Equal(that1.HttpRequest) {
		return false
	}
	if this.Type != that1.Type {
		return false
	}
	if this.StatsEnabled != that1.StatsEnabled {
		return false
	}
	return true
}
func (this *ClientToFrontend) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ClientToFrontend)
	if !ok {
		that2, ok := that.(ClientToFrontend)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.HttpResponse.Equal(that1.HttpResponse) {
		return false
	}
	if this.ClientID != that1.ClientID {
		return false
	}
	if !this.Stats.Equal(that1.Stats) {
		return false
	}
	return true
}
func (this *NotifyClientShutdownRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*NotifyClientShutdownRequest)
	if !ok {
		that2, ok := that.(NotifyClientShutdownRequest)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.ClientID != that1.ClientID {
		return false
	}
	return true
}
func (this *NotifyClientShutdownResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*NotifyClientShutdownResponse)
	if !ok {
		that2, ok := that.(NotifyClientShutdownResponse)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	return true
}
func (this *FrontendToClient) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&frontendv1pb.FrontendToClient{")
	if this.HttpRequest != nil {
		s = append(s, "HttpRequest: "+fmt.Sprintf("%#v", this.HttpRequest)+",\n")
	}
	s = append(s, "Type: "+fmt.Sprintf("%#v", this.Type)+",\n")
	s = append(s, "StatsEnabled: "+fmt.Sprintf("%#v", this.StatsEnabled)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *ClientToFrontend) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&frontendv1pb.ClientToFrontend{")
	if this.HttpResponse != nil {
		s = append(s, "HttpResponse: "+fmt.Sprintf("%#v", this.HttpResponse)+",\n")
	}
	s = append(s, "ClientID: "+fmt.Sprintf("%#v", this.ClientID)+",\n")
	if this.Stats != nil {
		s = append(s, "Stats: "+fmt.Sprintf("%#v", this.Stats)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *NotifyClientShutdownRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&frontendv1pb.NotifyClientShutdownRequest{")
	s = append(s, "ClientID: "+fmt.Sprintf("%#v", this.ClientID)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *NotifyClientShutdownResponse) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 4)
	s = append(s, "&frontendv1pb.NotifyClientShutdownResponse{")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringFrontend(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FrontendClient is the client API for Frontend service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FrontendClient interface {
	// After calling this method, client enters a loop, in which it waits for
	// a "FrontendToClient" message and replies with single "ClientToFrontend" message.
	Process(ctx context.Context, opts ...grpc.CallOption) (Frontend_ProcessClient, error)
	// The client notifies the query-frontend that it started a graceful shutdown.
	NotifyClientShutdown(ctx context.Context, in *NotifyClientShutdownRequest, opts ...grpc.CallOption) (*NotifyClientShutdownResponse, error)
}

type frontendClient struct {
	cc *grpc.ClientConn
}

func NewFrontendClient(cc *grpc.ClientConn) FrontendClient {
	return &frontendClient{cc}
}

func (c *frontendClient) Process(ctx context.Context, opts ...grpc.CallOption) (Frontend_ProcessClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Frontend_serviceDesc.Streams[0], "/frontend.Frontend/Process", opts...)
	if err != nil {
		return nil, err
	}
	x := &frontendProcessClient{stream}
	return x, nil
}

type Frontend_ProcessClient interface {
	Send(*ClientToFrontend) error
	Recv() (*FrontendToClient, error)
	grpc.ClientStream
}

type frontendProcessClient struct {
	grpc.ClientStream
}

func (x *frontendProcessClient) Send(m *ClientToFrontend) error {
	return x.ClientStream.SendMsg(m)
}

func (x *frontendProcessClient) Recv() (*FrontendToClient, error) {
	m := new(FrontendToClient)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *frontendClient) NotifyClientShutdown(ctx context.Context, in *NotifyClientShutdownRequest, opts ...grpc.CallOption) (*NotifyClientShutdownResponse, error) {
	out := new(NotifyClientShutdownResponse)
	err := c.cc.Invoke(ctx, "/frontend.Frontend/NotifyClientShutdown", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FrontendServer is the server API for Frontend service.
type FrontendServer interface {
	// After calling this method, client enters a loop, in which it waits for
	// a "FrontendToClient" message and replies with single "ClientToFrontend" message.
	Process(Frontend_ProcessServer) error
	// The client notifies the query-frontend that it started a graceful shutdown.
	NotifyClientShutdown(context.Context, *NotifyClientShutdownRequest) (*NotifyClientShutdownResponse, error)
}

// UnimplementedFrontendServer can be embedded to have forward compatible implementations.
type UnimplementedFrontendServer struct {
}

func (*UnimplementedFrontendServer) Process(srv Frontend_ProcessServer) error {
	return status.Errorf(codes.Unimplemented, "method Process not implemented")
}
func (*UnimplementedFrontendServer) NotifyClientShutdown(ctx context.Context, req *NotifyClientShutdownRequest) (*NotifyClientShutdownResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NotifyClientShutdown not implemented")
}

func RegisterFrontendServer(s *grpc.Server, srv FrontendServer) {
	s.RegisterService(&_Frontend_serviceDesc, srv)
}

func _Frontend_Process_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(FrontendServer).Process(&frontendProcessServer{stream})
}

type Frontend_ProcessServer interface {
	Send(*FrontendToClient) error
	Recv() (*ClientToFrontend, error)
	grpc.ServerStream
}

type frontendProcessServer struct {
	grpc.ServerStream
}

func (x *frontendProcessServer) Send(m *FrontendToClient) error {
	return x.ServerStream.SendMsg(m)
}

func (x *frontendProcessServer) Recv() (*ClientToFrontend, error) {
	m := new(ClientToFrontend)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Frontend_NotifyClientShutdown_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotifyClientShutdownRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FrontendServer).NotifyClientShutdown(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/frontend.Frontend/NotifyClientShutdown",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FrontendServer).NotifyClientShutdown(ctx, req.(*NotifyClientShutdownRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Frontend_serviceDesc = grpc.ServiceDesc{
	ServiceName: "frontend.Frontend",
	HandlerType: (*FrontendServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NotifyClientShutdown",
			Handler:    _Frontend_NotifyClientShutdown_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Process",
			Handler:       _Frontend_Process_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "frontend.proto",
}

func (m *FrontendToClient) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FrontendToClient) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FrontendToClient) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.StatsEnabled {
		i--
		if m.StatsEnabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if m.Type != 0 {
		i = encodeVarintFrontend(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x10
	}
	if m.HttpRequest != nil {
		{
			size, err := m.HttpRequest.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintFrontend(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ClientToFrontend) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClientToFrontend) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ClientToFrontend) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Stats != nil {
		{
			size, err := m.Stats.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintFrontend(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ClientID) > 0 {
		i -= len(m.ClientID)
		copy(dAtA[i:], m.ClientID)
		i = encodeVarintFrontend(dAtA, i, uint64(len(m.ClientID)))
		i--
		dAtA[i] = 0x12
	}
	if m.HttpResponse != nil {
		{
			size, err := m.HttpResponse.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintFrontend(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *NotifyClientShutdownRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NotifyClientShutdownRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NotifyClientShutdownRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ClientID) > 0 {
		i -= len(m.ClientID)
		copy(dAtA[i:], m.ClientID)
		i = encodeVarintFrontend(dAtA, i, uint64(len(m.ClientID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *NotifyClientShutdownResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NotifyClientShutdownResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NotifyClientShutdownResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintFrontend(dAtA []byte, offset int, v uint64) int {
	offset -= sovFrontend(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *FrontendToClient) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.HttpRequest != nil {
		l = m.HttpRequest.Size()
		n += 1 + l + sovFrontend(uint64(l))
	}
	if m.Type != 0 {
		n += 1 + sovFrontend(uint64(m.Type))
	}
	if m.StatsEnabled {
		n += 2
	}
	return n
}

func (m *ClientToFrontend) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.HttpResponse != nil {
		l = m.HttpResponse.Size()
		n += 1 + l + sovFrontend(uint64(l))
	}
	l = len(m.ClientID)
	if l > 0 {
		n += 1 + l + sovFrontend(uint64(l))
	}
	if m.Stats != nil {
		l = m.Stats.Size()
		n += 1 + l + sovFrontend(uint64(l))
	}
	return n
}

func (m *NotifyClientShutdownRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ClientID)
	if l > 0 {
		n += 1 + l + sovFrontend(uint64(l))
	}
	return n
}

func (m *NotifyClientShutdownResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovFrontend(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozFrontend(x uint64) (n int) {
	return sovFrontend(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *FrontendToClient) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&FrontendToClient{`,
		`HttpRequest:` + strings.Replace(fmt.Sprintf("%v", this.HttpRequest), "HTTPRequest", "httpgrpc.HTTPRequest", 1) + `,`,
		`Type:` + fmt.Sprintf("%v", this.Type) + `,`,
		`StatsEnabled:` + fmt.Sprintf("%v", this.StatsEnabled) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ClientToFrontend) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ClientToFrontend{`,
		`HttpResponse:` + strings.Replace(fmt.Sprintf("%v", this.HttpResponse), "HTTPResponse", "httpgrpc.HTTPResponse", 1) + `,`,
		`ClientID:` + fmt.Sprintf("%v", this.ClientID) + `,`,
		`Stats:` + strings.Replace(fmt.Sprintf("%v", this.Stats), "Stats", "stats.Stats", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *NotifyClientShutdownRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&NotifyClientShutdownRequest{`,
		`ClientID:` + fmt.Sprintf("%v", this.ClientID) + `,`,
		`}`,
	}, "")
	return s
}
func (this *NotifyClientShutdownResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&NotifyClientShutdownResponse{`,
		`}`,
	}, "")
	return s
}
func valueToStringFrontend(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *FrontendToClient) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFrontend
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: FrontendToClient: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FrontendToClient: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HttpRequest", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFrontend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthFrontend
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFrontend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.HttpRequest == nil {
				m.HttpRequest = &httpgrpc.HTTPRequest{}
			}
			if err := m.HttpRequest.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFrontend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= Type(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StatsEnabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFrontend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.StatsEnabled = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipFrontend(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFrontend
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthFrontend
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ClientToFrontend) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFrontend
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ClientToFrontend: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClientToFrontend: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HttpResponse", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFrontend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthFrontend
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFrontend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.HttpResponse == nil {
				m.HttpResponse = &httpgrpc.HTTPResponse{}
			}
			if err := m.HttpResponse.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFrontend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthFrontend
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFrontend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClientID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Stats", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFrontend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthFrontend
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFrontend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Stats == nil {
				m.Stats = &stats.Stats{}
			}
			if err := m.Stats.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFrontend(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFrontend
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthFrontend
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *NotifyClientShutdownRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFrontend
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: NotifyClientShutdownRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NotifyClientShutdownRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFrontend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthFrontend
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFrontend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClientID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFrontend(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFrontend
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthFrontend
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *NotifyClientShutdownResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFrontend
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: NotifyClientShutdownResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NotifyClientShutdownResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipFrontend(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFrontend
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthFrontend
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipFrontend(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowFrontend
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowFrontend
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowFrontend
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthFrontend
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthFrontend
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowFrontend
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipFrontend(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthFrontend
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthFrontend = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowFrontend   = fmt.Errorf("proto: integer overflow")
)
