// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.15.2
// source: v1/service.proto

package protov1

import (
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

type ListSomethingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Next  int64 `protobuf:"varint,1,opt,name=next,proto3" json:"next,omitempty"`
	Limit int64 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *ListSomethingRequest) Reset() {
	*x = ListSomethingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSomethingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSomethingRequest) ProtoMessage() {}

func (x *ListSomethingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSomethingRequest.ProtoReflect.Descriptor instead.
func (*ListSomethingRequest) Descriptor() ([]byte, []int) {
	return file_v1_service_proto_rawDescGZIP(), []int{0}
}

func (x *ListSomethingRequest) GetNext() int64 {
	if x != nil {
		return x.Next
	}
	return 0
}

func (x *ListSomethingRequest) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type ListSomethingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Foofoo []*ListSomethingResponse_Data `protobuf:"bytes,1,rep,name=foofoo,proto3" json:"foofoo,omitempty"`
}

func (x *ListSomethingResponse) Reset() {
	*x = ListSomethingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSomethingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSomethingResponse) ProtoMessage() {}

func (x *ListSomethingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSomethingResponse.ProtoReflect.Descriptor instead.
func (*ListSomethingResponse) Descriptor() ([]byte, []int) {
	return file_v1_service_proto_rawDescGZIP(), []int{1}
}

func (x *ListSomethingResponse) GetFoofoo() []*ListSomethingResponse_Data {
	if x != nil {
		return x.Foofoo
	}
	return nil
}

type ListSomethingResponse_Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id  int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Foo string `protobuf:"bytes,2,opt,name=foo,proto3" json:"foo,omitempty"`
}

func (x *ListSomethingResponse_Data) Reset() {
	*x = ListSomethingResponse_Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSomethingResponse_Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSomethingResponse_Data) ProtoMessage() {}

func (x *ListSomethingResponse_Data) ProtoReflect() protoreflect.Message {
	mi := &file_v1_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSomethingResponse_Data.ProtoReflect.Descriptor instead.
func (*ListSomethingResponse_Data) Descriptor() ([]byte, []int) {
	return file_v1_service_proto_rawDescGZIP(), []int{1, 0}
}

func (x *ListSomethingResponse_Data) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ListSomethingResponse_Data) GetFoo() string {
	if x != nil {
		return x.Foo
	}
	return ""
}

var File_v1_service_proto protoreflect.FileDescriptor

var file_v1_service_proto_rawDesc = []byte{
	0x0a, 0x10, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x13, 0x67, 0x72, 0x70, 0x63, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x22, 0x40, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x53,
	0x6f, 0x6d, 0x65, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x6e,
	0x65, 0x78, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x8a, 0x01, 0x0a, 0x15, 0x4c, 0x69,
	0x73, 0x74, 0x53, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x06, 0x66, 0x6f, 0x6f, 0x66, 0x6f, 0x6f, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x6f,
	0x6d, 0x65, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e,
	0x44, 0x61, 0x74, 0x61, 0x52, 0x06, 0x66, 0x6f, 0x6f, 0x66, 0x6f, 0x6f, 0x1a, 0x28, 0x0a, 0x04,
	0x44, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x66, 0x6f, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x66, 0x6f, 0x6f, 0x32, 0x7e, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x6f,
	0x6d, 0x65, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x66,
	0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x12,
	0x29, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x6f, 0x6d, 0x65, 0x74, 0x68,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_service_proto_rawDescOnce sync.Once
	file_v1_service_proto_rawDescData = file_v1_service_proto_rawDesc
)

func file_v1_service_proto_rawDescGZIP() []byte {
	file_v1_service_proto_rawDescOnce.Do(func() {
		file_v1_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_service_proto_rawDescData)
	})
	return file_v1_service_proto_rawDescData
}

var file_v1_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_v1_service_proto_goTypes = []interface{}{
	(*ListSomethingRequest)(nil),       // 0: grpcserver.proto.v1.ListSomethingRequest
	(*ListSomethingResponse)(nil),      // 1: grpcserver.proto.v1.ListSomethingResponse
	(*ListSomethingResponse_Data)(nil), // 2: grpcserver.proto.v1.ListSomethingResponse.Data
}
var file_v1_service_proto_depIdxs = []int32{
	2, // 0: grpcserver.proto.v1.ListSomethingResponse.foofoo:type_name -> grpcserver.proto.v1.ListSomethingResponse.Data
	0, // 1: grpcserver.proto.v1.ListSomethingService.ListSomething:input_type -> grpcserver.proto.v1.ListSomethingRequest
	1, // 2: grpcserver.proto.v1.ListSomethingService.ListSomething:output_type -> grpcserver.proto.v1.ListSomethingResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_v1_service_proto_init() }
func file_v1_service_proto_init() {
	if File_v1_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSomethingRequest); i {
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
		file_v1_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSomethingResponse); i {
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
		file_v1_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSomethingResponse_Data); i {
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
			RawDescriptor: file_v1_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_service_proto_goTypes,
		DependencyIndexes: file_v1_service_proto_depIdxs,
		MessageInfos:      file_v1_service_proto_msgTypes,
	}.Build()
	File_v1_service_proto = out.File
	file_v1_service_proto_rawDesc = nil
	file_v1_service_proto_goTypes = nil
	file_v1_service_proto_depIdxs = nil
}
