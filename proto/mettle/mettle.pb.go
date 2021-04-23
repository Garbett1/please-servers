// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.11.4
// source: proto/mettle.proto

package build_please_remote_mettle

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Job struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID         string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Current    []byte `protobuf:"bytes,2,opt,name=Current,proto3" json:"Current,omitempty"`
	LastUpdate int64  `protobuf:"varint,5,opt,name=LastUpdate,proto3" json:"LastUpdate,omitempty"`
	SentFirst  bool   `protobuf:"varint,3,opt,name=SentFirst,proto3" json:"SentFirst,omitempty"`
	Done       bool   `protobuf:"varint,4,opt,name=Done,proto3" json:"Done,omitempty"`
}

func (x *Job) Reset() {
	*x = Job{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_mettle_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Job) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Job) ProtoMessage() {}

func (x *Job) ProtoReflect() protoreflect.Message {
	mi := &file_proto_mettle_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Job.ProtoReflect.Descriptor instead.
func (*Job) Descriptor() ([]byte, []int) {
	return file_proto_mettle_proto_rawDescGZIP(), []int{0}
}

func (x *Job) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *Job) GetCurrent() []byte {
	if x != nil {
		return x.Current
	}
	return nil
}

func (x *Job) GetLastUpdate() int64 {
	if x != nil {
		return x.LastUpdate
	}
	return 0
}

func (x *Job) GetSentFirst() bool {
	if x != nil {
		return x.SentFirst
	}
	return false
}

func (x *Job) GetDone() bool {
	if x != nil {
		return x.Done
	}
	return false
}

type ServeExecutionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ServeExecutionsRequest) Reset() {
	*x = ServeExecutionsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_mettle_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServeExecutionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServeExecutionsRequest) ProtoMessage() {}

func (x *ServeExecutionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_mettle_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServeExecutionsRequest.ProtoReflect.Descriptor instead.
func (*ServeExecutionsRequest) Descriptor() ([]byte, []int) {
	return file_proto_mettle_proto_rawDescGZIP(), []int{1}
}

type ServeExecutionsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Jobs []*Job `protobuf:"bytes,1,rep,name=jobs,proto3" json:"jobs,omitempty"`
}

func (x *ServeExecutionsResponse) Reset() {
	*x = ServeExecutionsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_mettle_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServeExecutionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServeExecutionsResponse) ProtoMessage() {}

func (x *ServeExecutionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_mettle_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServeExecutionsResponse.ProtoReflect.Descriptor instead.
func (*ServeExecutionsResponse) Descriptor() ([]byte, []int) {
	return file_proto_mettle_proto_rawDescGZIP(), []int{2}
}

func (x *ServeExecutionsResponse) GetJobs() []*Job {
	if x != nil {
		return x.Jobs
	}
	return nil
}

var File_proto_mettle_proto protoreflect.FileDescriptor

var file_proto_mettle_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x70, 0x6c, 0x65, 0x61,
	0x73, 0x65, 0x2e, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2e, 0x6d, 0x65, 0x74, 0x74, 0x6c, 0x65,
	0x22, 0x81, 0x01, 0x0a, 0x03, 0x4a, 0x6f, 0x62, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x43, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x4c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x4c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x65, 0x6e, 0x74, 0x46, 0x69, 0x72, 0x73, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x53, 0x65, 0x6e, 0x74, 0x46, 0x69, 0x72, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x44, 0x6f, 0x6e, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04,
	0x44, 0x6f, 0x6e, 0x65, 0x22, 0x18, 0x0a, 0x16, 0x53, 0x65, 0x72, 0x76, 0x65, 0x45, 0x78, 0x65,
	0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x4e,
	0x0a, 0x17, 0x53, 0x65, 0x72, 0x76, 0x65, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x04, 0x6a, 0x6f, 0x62,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e,
	0x70, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x2e, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2e, 0x6d, 0x65,
	0x74, 0x74, 0x6c, 0x65, 0x2e, 0x4a, 0x6f, 0x62, 0x52, 0x04, 0x6a, 0x6f, 0x62, 0x73, 0x32, 0x87,
	0x01, 0x0a, 0x09, 0x42, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x12, 0x7a, 0x0a, 0x0f,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x32, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x70, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x2e, 0x72,
	0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2e, 0x6d, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x2e, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x33, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x70, 0x6c, 0x65, 0x61,
	0x73, 0x65, 0x2e, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2e, 0x6d, 0x65, 0x74, 0x74, 0x6c, 0x65,
	0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_mettle_proto_rawDescOnce sync.Once
	file_proto_mettle_proto_rawDescData = file_proto_mettle_proto_rawDesc
)

func file_proto_mettle_proto_rawDescGZIP() []byte {
	file_proto_mettle_proto_rawDescOnce.Do(func() {
		file_proto_mettle_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_mettle_proto_rawDescData)
	})
	return file_proto_mettle_proto_rawDescData
}

var file_proto_mettle_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_mettle_proto_goTypes = []interface{}{
	(*Job)(nil),                     // 0: build.please.remote.mettle.Job
	(*ServeExecutionsRequest)(nil),  // 1: build.please.remote.mettle.ServeExecutionsRequest
	(*ServeExecutionsResponse)(nil), // 2: build.please.remote.mettle.ServeExecutionsResponse
}
var file_proto_mettle_proto_depIdxs = []int32{
	0, // 0: build.please.remote.mettle.ServeExecutionsResponse.jobs:type_name -> build.please.remote.mettle.Job
	1, // 1: build.please.remote.mettle.Bootstrap.ServeExecutions:input_type -> build.please.remote.mettle.ServeExecutionsRequest
	2, // 2: build.please.remote.mettle.Bootstrap.ServeExecutions:output_type -> build.please.remote.mettle.ServeExecutionsResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_mettle_proto_init() }
func file_proto_mettle_proto_init() {
	if File_proto_mettle_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_mettle_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Job); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_mettle_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServeExecutionsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_mettle_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServeExecutionsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_mettle_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_mettle_proto_goTypes,
		DependencyIndexes: file_proto_mettle_proto_depIdxs,
		MessageInfos:      file_proto_mettle_proto_msgTypes,
	}.Build()
	File_proto_mettle_proto = out.File
	file_proto_mettle_proto_rawDesc = nil
	file_proto_mettle_proto_goTypes = nil
	file_proto_mettle_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// BootstrapClient is the client API for Bootstrap service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BootstrapClient interface {
	ServeExecutions(ctx context.Context, in *ServeExecutionsRequest, opts ...grpc.CallOption) (*ServeExecutionsResponse, error)
}

type bootstrapClient struct {
	cc grpc.ClientConnInterface
}

func NewBootstrapClient(cc grpc.ClientConnInterface) BootstrapClient {
	return &bootstrapClient{cc}
}

func (c *bootstrapClient) ServeExecutions(ctx context.Context, in *ServeExecutionsRequest, opts ...grpc.CallOption) (*ServeExecutionsResponse, error) {
	out := new(ServeExecutionsResponse)
	err := c.cc.Invoke(ctx, "/build.please.remote.mettle.Bootstrap/ServeExecutions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BootstrapServer is the server API for Bootstrap service.
type BootstrapServer interface {
	ServeExecutions(context.Context, *ServeExecutionsRequest) (*ServeExecutionsResponse, error)
}

// UnimplementedBootstrapServer can be embedded to have forward compatible implementations.
type UnimplementedBootstrapServer struct {
}

func (*UnimplementedBootstrapServer) ServeExecutions(context.Context, *ServeExecutionsRequest) (*ServeExecutionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ServeExecutions not implemented")
}

func RegisterBootstrapServer(s *grpc.Server, srv BootstrapServer) {
	s.RegisterService(&_Bootstrap_serviceDesc, srv)
}

func _Bootstrap_ServeExecutions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServeExecutionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BootstrapServer).ServeExecutions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/build.please.remote.mettle.Bootstrap/ServeExecutions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BootstrapServer).ServeExecutions(ctx, req.(*ServeExecutionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Bootstrap_serviceDesc = grpc.ServiceDesc{
	ServiceName: "build.please.remote.mettle.Bootstrap",
	HandlerType: (*BootstrapServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ServeExecutions",
			Handler:    _Bootstrap_ServeExecutions_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/mettle.proto",
}
