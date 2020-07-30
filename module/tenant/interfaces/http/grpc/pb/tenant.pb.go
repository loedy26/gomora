// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: tenant.proto

package tenant

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type TenantRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *TenantRequest) Reset() {
	*x = TenantRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tenant_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TenantRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TenantRequest) ProtoMessage() {}

func (x *TenantRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tenant_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TenantRequest.ProtoReflect.Descriptor instead.
func (*TenantRequest) Descriptor() ([]byte, []int) {
	return file_tenant_proto_rawDescGZIP(), []int{0}
}

func (x *TenantRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type TenantResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string               `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Alias         string               `protobuf:"bytes,3,opt,name=alias,proto3" json:"alias,omitempty"`
	Email         string               `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	Code          string               `protobuf:"bytes,5,opt,name=code,proto3" json:"code,omitempty"`
	Address       string               `protobuf:"bytes,6,opt,name=address,proto3" json:"address,omitempty"`
	ContactNumber string               `protobuf:"bytes,7,opt,name=contactNumber,proto3" json:"contactNumber,omitempty"`
	Avatar        string               `protobuf:"bytes,8,opt,name=avatar,proto3" json:"avatar,omitempty"`
	IsActive      bool                 `protobuf:"varint,9,opt,name=isActive,proto3" json:"isActive,omitempty"`
	CreatedAt     *timestamp.Timestamp `protobuf:"bytes,10,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     *timestamp.Timestamp `protobuf:"bytes,11,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *TenantResponse) Reset() {
	*x = TenantResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tenant_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TenantResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TenantResponse) ProtoMessage() {}

func (x *TenantResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tenant_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TenantResponse.ProtoReflect.Descriptor instead.
func (*TenantResponse) Descriptor() ([]byte, []int) {
	return file_tenant_proto_rawDescGZIP(), []int{1}
}

func (x *TenantResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *TenantResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TenantResponse) GetAlias() string {
	if x != nil {
		return x.Alias
	}
	return ""
}

func (x *TenantResponse) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *TenantResponse) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *TenantResponse) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *TenantResponse) GetContactNumber() string {
	if x != nil {
		return x.ContactNumber
	}
	return ""
}

func (x *TenantResponse) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *TenantResponse) GetIsActive() bool {
	if x != nil {
		return x.IsActive
	}
	return false
}

func (x *TenantResponse) GetCreatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *TenantResponse) GetUpdatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

var File_tenant_proto protoreflect.FileDescriptor

var file_tenant_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1f, 0x0a, 0x0d, 0x54, 0x65, 0x6e, 0x61, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0xde, 0x02, 0x0a, 0x0e, 0x54, 0x65, 0x6e,
	0x61, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x61, 0x6c, 0x69, 0x61, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x24, 0x0a, 0x0d, 0x63, 0x6f, 0x6e,
	0x74, 0x61, 0x63, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12,
	0x16, 0x0a, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x73, 0x41, 0x63, 0x74,
	0x69, 0x76, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x41, 0x63, 0x74,
	0x69, 0x76, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39,
	0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x32, 0x51, 0x0a, 0x0d, 0x54, 0x65, 0x6e,
	0x61, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x40, 0x0a, 0x0d, 0x47, 0x65,
	0x74, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x42, 0x79, 0x49, 0x44, 0x12, 0x15, 0x2e, 0x74, 0x65,
	0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x54, 0x65, 0x6e, 0x61,
	0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tenant_proto_rawDescOnce sync.Once
	file_tenant_proto_rawDescData = file_tenant_proto_rawDesc
)

func file_tenant_proto_rawDescGZIP() []byte {
	file_tenant_proto_rawDescOnce.Do(func() {
		file_tenant_proto_rawDescData = protoimpl.X.CompressGZIP(file_tenant_proto_rawDescData)
	})
	return file_tenant_proto_rawDescData
}

var file_tenant_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_tenant_proto_goTypes = []interface{}{
	(*TenantRequest)(nil),       // 0: tenant.TenantRequest
	(*TenantResponse)(nil),      // 1: tenant.TenantResponse
	(*timestamp.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_tenant_proto_depIdxs = []int32{
	2, // 0: tenant.TenantResponse.created_at:type_name -> google.protobuf.Timestamp
	2, // 1: tenant.TenantResponse.updated_at:type_name -> google.protobuf.Timestamp
	0, // 2: tenant.TenantService.GetTenantByID:input_type -> tenant.TenantRequest
	1, // 3: tenant.TenantService.GetTenantByID:output_type -> tenant.TenantResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_tenant_proto_init() }
func file_tenant_proto_init() {
	if File_tenant_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tenant_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TenantRequest); i {
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
		file_tenant_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TenantResponse); i {
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
			RawDescriptor: file_tenant_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tenant_proto_goTypes,
		DependencyIndexes: file_tenant_proto_depIdxs,
		MessageInfos:      file_tenant_proto_msgTypes,
	}.Build()
	File_tenant_proto = out.File
	file_tenant_proto_rawDesc = nil
	file_tenant_proto_goTypes = nil
	file_tenant_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TenantServiceClient is the client API for TenantService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TenantServiceClient interface {
	GetTenantByID(ctx context.Context, in *TenantRequest, opts ...grpc.CallOption) (*TenantResponse, error)
}

type tenantServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTenantServiceClient(cc grpc.ClientConnInterface) TenantServiceClient {
	return &tenantServiceClient{cc}
}

func (c *tenantServiceClient) GetTenantByID(ctx context.Context, in *TenantRequest, opts ...grpc.CallOption) (*TenantResponse, error) {
	out := new(TenantResponse)
	err := c.cc.Invoke(ctx, "/tenant.TenantService/GetTenantByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TenantServiceServer is the server API for TenantService service.
type TenantServiceServer interface {
	GetTenantByID(context.Context, *TenantRequest) (*TenantResponse, error)
}

// UnimplementedTenantServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTenantServiceServer struct {
}

func (*UnimplementedTenantServiceServer) GetTenantByID(context.Context, *TenantRequest) (*TenantResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTenantByID not implemented")
}

func RegisterTenantServiceServer(s *grpc.Server, srv TenantServiceServer) {
	s.RegisterService(&_TenantService_serviceDesc, srv)
}

func _TenantService_GetTenantByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TenantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantServiceServer).GetTenantByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tenant.TenantService/GetTenantByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantServiceServer).GetTenantByID(ctx, req.(*TenantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TenantService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tenant.TenantService",
	HandlerType: (*TenantServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTenantByID",
			Handler:    _TenantService_GetTenantByID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tenant.proto",
}