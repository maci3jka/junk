// Code generated by protoc-gen-go.
// source: api.proto
// DO NOT EDIT!

/*
Package types is a generated protocol buffer package.

It is generated from these files:
	api.proto

It has these top-level messages:
	StartJobRequest
	StartJobResponse
	DeleteJobRequest
	DeleteJobResponse
	ListJobsRequest
	Job
	ListJobsResponse
	StateRequest
	StateResponse
	LogsRequest
	Log
*/
package types

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
const _ = proto.ProtoPackageIsVersion1

type StartJobRequest struct {
	Name      string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Args      []string `protobuf:"bytes,2,rep,name=args" json:"args,omitempty"`
	Artifacts string   `protobuf:"bytes,3,opt,name=artifacts" json:"artifacts,omitempty"`
}

func (m *StartJobRequest) Reset()                    { *m = StartJobRequest{} }
func (m *StartJobRequest) String() string            { return proto.CompactTextString(m) }
func (*StartJobRequest) ProtoMessage()               {}
func (*StartJobRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type StartJobResponse struct {
	Id uint32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *StartJobResponse) Reset()                    { *m = StartJobResponse{} }
func (m *StartJobResponse) String() string            { return proto.CompactTextString(m) }
func (*StartJobResponse) ProtoMessage()               {}
func (*StartJobResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type DeleteJobRequest struct {
	Id uint32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *DeleteJobRequest) Reset()                    { *m = DeleteJobRequest{} }
func (m *DeleteJobRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteJobRequest) ProtoMessage()               {}
func (*DeleteJobRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type DeleteJobResponse struct {
}

func (m *DeleteJobResponse) Reset()                    { *m = DeleteJobResponse{} }
func (m *DeleteJobResponse) String() string            { return proto.CompactTextString(m) }
func (*DeleteJobResponse) ProtoMessage()               {}
func (*DeleteJobResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type ListJobsRequest struct {
}

func (m *ListJobsRequest) Reset()                    { *m = ListJobsRequest{} }
func (m *ListJobsRequest) String() string            { return proto.CompactTextString(m) }
func (*ListJobsRequest) ProtoMessage()               {}
func (*ListJobsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type Job struct {
	Id        uint32   `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name      string   `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Args      []string `protobuf:"bytes,3,rep,name=args" json:"args,omitempty"`
	Artifacts string   `protobuf:"bytes,4,opt,name=artifacts" json:"artifacts,omitempty"`
	Status    string   `protobuf:"bytes,5,opt,name=status" json:"status,omitempty"`
}

func (m *Job) Reset()                    { *m = Job{} }
func (m *Job) String() string            { return proto.CompactTextString(m) }
func (*Job) ProtoMessage()               {}
func (*Job) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type ListJobsResponse struct {
	Jobs []*Job `protobuf:"bytes,1,rep,name=jobs" json:"jobs,omitempty"`
}

func (m *ListJobsResponse) Reset()                    { *m = ListJobsResponse{} }
func (m *ListJobsResponse) String() string            { return proto.CompactTextString(m) }
func (*ListJobsResponse) ProtoMessage()               {}
func (*ListJobsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ListJobsResponse) GetJobs() []*Job {
	if m != nil {
		return m.Jobs
	}
	return nil
}

type StateRequest struct {
	Id uint32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *StateRequest) Reset()                    { *m = StateRequest{} }
func (m *StateRequest) String() string            { return proto.CompactTextString(m) }
func (*StateRequest) ProtoMessage()               {}
func (*StateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

type StateResponse struct {
	Status string `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
}

func (m *StateResponse) Reset()                    { *m = StateResponse{} }
func (m *StateResponse) String() string            { return proto.CompactTextString(m) }
func (*StateResponse) ProtoMessage()               {}
func (*StateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

type LogsRequest struct {
	Id     uint32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Follow bool   `protobuf:"varint,2,opt,name=follow" json:"follow,omitempty"`
}

func (m *LogsRequest) Reset()                    { *m = LogsRequest{} }
func (m *LogsRequest) String() string            { return proto.CompactTextString(m) }
func (*LogsRequest) ProtoMessage()               {}
func (*LogsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

type Log struct {
	Log string `protobuf:"bytes,1,opt,name=log" json:"log,omitempty"`
}

func (m *Log) Reset()                    { *m = Log{} }
func (m *Log) String() string            { return proto.CompactTextString(m) }
func (*Log) ProtoMessage()               {}
func (*Log) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func init() {
	proto.RegisterType((*StartJobRequest)(nil), "types.StartJobRequest")
	proto.RegisterType((*StartJobResponse)(nil), "types.StartJobResponse")
	proto.RegisterType((*DeleteJobRequest)(nil), "types.DeleteJobRequest")
	proto.RegisterType((*DeleteJobResponse)(nil), "types.DeleteJobResponse")
	proto.RegisterType((*ListJobsRequest)(nil), "types.ListJobsRequest")
	proto.RegisterType((*Job)(nil), "types.Job")
	proto.RegisterType((*ListJobsResponse)(nil), "types.ListJobsResponse")
	proto.RegisterType((*StateRequest)(nil), "types.StateRequest")
	proto.RegisterType((*StateResponse)(nil), "types.StateResponse")
	proto.RegisterType((*LogsRequest)(nil), "types.LogsRequest")
	proto.RegisterType((*Log)(nil), "types.Log")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for API service

type APIClient interface {
	StartJob(ctx context.Context, in *StartJobRequest, opts ...grpc.CallOption) (*StartJobResponse, error)
	DeleteJob(ctx context.Context, in *DeleteJobRequest, opts ...grpc.CallOption) (*DeleteJobResponse, error)
	ListJobs(ctx context.Context, in *ListJobsRequest, opts ...grpc.CallOption) (*ListJobsResponse, error)
	State(ctx context.Context, in *StateRequest, opts ...grpc.CallOption) (*StateResponse, error)
	Logs(ctx context.Context, in *LogsRequest, opts ...grpc.CallOption) (API_LogsClient, error)
}

type aPIClient struct {
	cc *grpc.ClientConn
}

func NewAPIClient(cc *grpc.ClientConn) APIClient {
	return &aPIClient{cc}
}

func (c *aPIClient) StartJob(ctx context.Context, in *StartJobRequest, opts ...grpc.CallOption) (*StartJobResponse, error) {
	out := new(StartJobResponse)
	err := grpc.Invoke(ctx, "/types.API/StartJob", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) DeleteJob(ctx context.Context, in *DeleteJobRequest, opts ...grpc.CallOption) (*DeleteJobResponse, error) {
	out := new(DeleteJobResponse)
	err := grpc.Invoke(ctx, "/types.API/DeleteJob", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) ListJobs(ctx context.Context, in *ListJobsRequest, opts ...grpc.CallOption) (*ListJobsResponse, error) {
	out := new(ListJobsResponse)
	err := grpc.Invoke(ctx, "/types.API/ListJobs", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) State(ctx context.Context, in *StateRequest, opts ...grpc.CallOption) (*StateResponse, error) {
	out := new(StateResponse)
	err := grpc.Invoke(ctx, "/types.API/State", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) Logs(ctx context.Context, in *LogsRequest, opts ...grpc.CallOption) (API_LogsClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_API_serviceDesc.Streams[0], c.cc, "/types.API/Logs", opts...)
	if err != nil {
		return nil, err
	}
	x := &aPILogsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type API_LogsClient interface {
	Recv() (*Log, error)
	grpc.ClientStream
}

type aPILogsClient struct {
	grpc.ClientStream
}

func (x *aPILogsClient) Recv() (*Log, error) {
	m := new(Log)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for API service

type APIServer interface {
	StartJob(context.Context, *StartJobRequest) (*StartJobResponse, error)
	DeleteJob(context.Context, *DeleteJobRequest) (*DeleteJobResponse, error)
	ListJobs(context.Context, *ListJobsRequest) (*ListJobsResponse, error)
	State(context.Context, *StateRequest) (*StateResponse, error)
	Logs(*LogsRequest, API_LogsServer) error
}

func RegisterAPIServer(s *grpc.Server, srv APIServer) {
	s.RegisterService(&_API_serviceDesc, srv)
}

func _API_StartJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(StartJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(APIServer).StartJob(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _API_DeleteJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(DeleteJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(APIServer).DeleteJob(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _API_ListJobs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(ListJobsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(APIServer).ListJobs(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _API_State_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(StateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(APIServer).State(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _API_Logs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(LogsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(APIServer).Logs(m, &aPILogsServer{stream})
}

type API_LogsServer interface {
	Send(*Log) error
	grpc.ServerStream
}

type aPILogsServer struct {
	grpc.ServerStream
}

func (x *aPILogsServer) Send(m *Log) error {
	return x.ServerStream.SendMsg(m)
}

var _API_serviceDesc = grpc.ServiceDesc{
	ServiceName: "types.API",
	HandlerType: (*APIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StartJob",
			Handler:    _API_StartJob_Handler,
		},
		{
			MethodName: "DeleteJob",
			Handler:    _API_DeleteJob_Handler,
		},
		{
			MethodName: "ListJobs",
			Handler:    _API_ListJobs_Handler,
		},
		{
			MethodName: "State",
			Handler:    _API_State_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Logs",
			Handler:       _API_Logs_Handler,
			ServerStreams: true,
		},
	},
}

var fileDescriptor0 = []byte{
	// 379 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x53, 0x41, 0x4f, 0xf2, 0x40,
	0x10, 0xa5, 0xdd, 0x42, 0x60, 0xf8, 0xf8, 0x80, 0xe1, 0x0b, 0x34, 0xcd, 0x17, 0x42, 0xf6, 0x22,
	0xf1, 0x40, 0x0c, 0xea, 0xd1, 0x44, 0x13, 0x2f, 0x1a, 0x0e, 0xa6, 0x1e, 0x3c, 0x17, 0x5d, 0x48,
	0x4d, 0x65, 0x6b, 0x77, 0x89, 0xf1, 0x17, 0xf9, 0x37, 0xdd, 0x5d, 0xb7, 0x2d, 0x14, 0xb9, 0xed,
	0xbc, 0x9d, 0x79, 0xfb, 0xde, 0x9b, 0x16, 0x5a, 0x51, 0x1a, 0xcf, 0xd2, 0x8c, 0x4b, 0x8e, 0x75,
	0xf9, 0x99, 0x32, 0x41, 0x9f, 0xa0, 0xfb, 0x28, 0xa3, 0x4c, 0xde, 0xf3, 0x65, 0xc8, 0xde, 0xb7,
	0x4c, 0x48, 0x44, 0xf0, 0x36, 0xd1, 0x1b, 0xf3, 0x9d, 0x89, 0x33, 0x6d, 0x85, 0xe6, 0xac, 0xb1,
	0x28, 0x5b, 0x0b, 0xdf, 0x9d, 0x10, 0x8d, 0xe9, 0x33, 0xfe, 0x57, 0x74, 0x99, 0x8c, 0x57, 0xd1,
	0xb3, 0x14, 0x3e, 0x31, 0xcd, 0x25, 0x40, 0x29, 0xf4, 0x4a, 0x62, 0x91, 0xf2, 0x8d, 0x60, 0xf8,
	0x17, 0xdc, 0xf8, 0xc5, 0xf0, 0x76, 0x42, 0x75, 0xd2, 0x3d, 0xb7, 0x2c, 0x61, 0x92, 0xed, 0xbc,
	0x5e, 0xed, 0x19, 0x40, 0x7f, 0xa7, 0xe7, 0x87, 0x88, 0xf6, 0xa1, 0xbb, 0x88, 0x85, 0xe6, 0x16,
	0x76, 0x8e, 0x0a, 0x20, 0xaa, 0xac, 0x8e, 0x17, 0x66, 0xdc, 0x5f, 0xcc, 0x90, 0x63, 0x66, 0xbc,
	0x8a, 0x19, 0x1c, 0x42, 0x43, 0xc8, 0x48, 0x6e, 0x85, 0x5f, 0x37, 0x57, 0xb6, 0xa2, 0x73, 0xe8,
	0x95, 0x3a, 0xac, 0xc9, 0x31, 0x78, 0xaf, 0xaa, 0x56, 0x1a, 0xc8, 0xb4, 0x3d, 0x87, 0x99, 0xc9,
	0x79, 0xa6, 0xd5, 0x1b, 0x9c, 0x8e, 0xe1, 0x8f, 0x0a, 0x46, 0xb2, 0x63, 0x86, 0x4f, 0xa0, 0x63,
	0xef, 0x2d, 0x61, 0xf9, 0xb8, 0xb3, 0xf7, 0xf8, 0x25, 0xb4, 0x17, 0x7c, 0x2d, 0x8e, 0xf0, 0xe8,
	0xb1, 0x15, 0x4f, 0x12, 0xfe, 0x61, 0xbc, 0x37, 0x43, 0x5b, 0xd1, 0x11, 0x10, 0x35, 0x86, 0x3d,
	0x20, 0x09, 0x5f, 0x5b, 0x4a, 0x7d, 0x9c, 0x7f, 0xb9, 0x40, 0x6e, 0x1e, 0xee, 0xf0, 0x0a, 0x9a,
	0xf9, 0xe6, 0x70, 0x68, 0xe5, 0x57, 0xbe, 0x91, 0x60, 0x74, 0x80, 0xdb, 0xcd, 0xd4, 0xf0, 0x1a,
	0x5a, 0xc5, 0xc2, 0x30, 0xef, 0xab, 0xae, 0x39, 0xf0, 0x0f, 0x2f, 0x0a, 0x06, 0x25, 0x20, 0x4f,
	0xb5, 0x10, 0x50, 0x59, 0x77, 0x21, 0xa0, 0x1a, 0xbf, 0x1a, 0xbf, 0x80, 0xba, 0x09, 0x10, 0x07,
	0xa5, 0xc8, 0x22, 0xee, 0xe0, 0xdf, 0x3e, 0x58, 0x4c, 0x9d, 0x82, 0xa7, 0xd3, 0x44, 0xcc, 0x89,
	0xcb, 0x68, 0x03, 0x28, 0x31, 0x5a, 0x3b, 0x73, 0x96, 0x0d, 0xf3, 0x0b, 0x9d, 0x7f, 0x07, 0x00,
	0x00, 0xff, 0xff, 0x51, 0x84, 0x2b, 0x66, 0x4f, 0x03, 0x00, 0x00,
}
