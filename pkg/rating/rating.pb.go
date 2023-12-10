// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.24.4
// source: pkg/rating/rating.proto

package rating

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

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_rating_rating_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_rating_rating_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_pkg_rating_rating_proto_rawDescGZIP(), []int{0}
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_rating_rating_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_rating_rating_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_pkg_rating_rating_proto_rawDescGZIP(), []int{1}
}

type RateHostRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	HostId int64 `protobuf:"varint,2,opt,name=hostId,proto3" json:"hostId,omitempty"`
	Grade  int64 `protobuf:"varint,3,opt,name=grade,proto3" json:"grade,omitempty"`
}

func (x *RateHostRequest) Reset() {
	*x = RateHostRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_rating_rating_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RateHostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RateHostRequest) ProtoMessage() {}

func (x *RateHostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_rating_rating_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RateHostRequest.ProtoReflect.Descriptor instead.
func (*RateHostRequest) Descriptor() ([]byte, []int) {
	return file_pkg_rating_rating_proto_rawDescGZIP(), []int{2}
}

func (x *RateHostRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *RateHostRequest) GetHostId() int64 {
	if x != nil {
		return x.HostId
	}
	return 0
}

func (x *RateHostRequest) GetGrade() int64 {
	if x != nil {
		return x.Grade
	}
	return 0
}

type RateHostResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HostId int64 `protobuf:"varint,1,opt,name=hostId,proto3" json:"hostId,omitempty"`
}

func (x *RateHostResponse) Reset() {
	*x = RateHostResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_rating_rating_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RateHostResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RateHostResponse) ProtoMessage() {}

func (x *RateHostResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_rating_rating_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RateHostResponse.ProtoReflect.Descriptor instead.
func (*RateHostResponse) Descriptor() ([]byte, []int) {
	return file_pkg_rating_rating_proto_rawDescGZIP(), []int{3}
}

func (x *RateHostResponse) GetHostId() int64 {
	if x != nil {
		return x.HostId
	}
	return 0
}

var File_pkg_rating_rating_proto protoreflect.FileDescriptor

var file_pkg_rating_rating_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x6b, 0x67, 0x2f, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x72, 0x61, 0x74,
	0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e,
	0x67, 0x22, 0x09, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x0a, 0x0a, 0x08,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x57, 0x0a, 0x0f, 0x52, 0x61, 0x74, 0x65,
	0x48, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x68, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x67,
	0x72, 0x61, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x67, 0x72, 0x61, 0x64,
	0x65, 0x22, 0x2a, 0x0a, 0x10, 0x52, 0x61, 0x74, 0x65, 0x48, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x68, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x32, 0x8a, 0x01,
	0x0a, 0x0d, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x3f, 0x0a, 0x08, 0x52, 0x61, 0x74, 0x65, 0x48, 0x6f, 0x73, 0x74, 0x12, 0x17, 0x2e, 0x72, 0x61,
	0x74, 0x69, 0x6e, 0x67, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x48, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x52, 0x61,
	0x74, 0x65, 0x48, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x38, 0x0a, 0x11, 0x52, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0f, 0x2e, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x2e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x7a, 0x6f, 0x6e, 0x69, 0x6f, 0x70,
	0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x72,
	0x61, 0x74, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_rating_rating_proto_rawDescOnce sync.Once
	file_pkg_rating_rating_proto_rawDescData = file_pkg_rating_rating_proto_rawDesc
)

func file_pkg_rating_rating_proto_rawDescGZIP() []byte {
	file_pkg_rating_rating_proto_rawDescOnce.Do(func() {
		file_pkg_rating_rating_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_rating_rating_proto_rawDescData)
	})
	return file_pkg_rating_rating_proto_rawDescData
}

var file_pkg_rating_rating_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_pkg_rating_rating_proto_goTypes = []interface{}{
	(*Request)(nil),          // 0: rating.Request
	(*Response)(nil),         // 1: rating.Response
	(*RateHostRequest)(nil),  // 2: rating.RateHostRequest
	(*RateHostResponse)(nil), // 3: rating.RateHostResponse
}
var file_pkg_rating_rating_proto_depIdxs = []int32{
	2, // 0: rating.RatingService.RateHost:input_type -> rating.RateHostRequest
	0, // 1: rating.RatingService.RateAccommodation:input_type -> rating.Request
	3, // 2: rating.RatingService.RateHost:output_type -> rating.RateHostResponse
	1, // 3: rating.RatingService.RateAccommodation:output_type -> rating.Response
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_rating_rating_proto_init() }
func file_pkg_rating_rating_proto_init() {
	if File_pkg_rating_rating_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_rating_rating_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_pkg_rating_rating_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_pkg_rating_rating_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RateHostRequest); i {
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
		file_pkg_rating_rating_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RateHostResponse); i {
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
			RawDescriptor: file_pkg_rating_rating_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_rating_rating_proto_goTypes,
		DependencyIndexes: file_pkg_rating_rating_proto_depIdxs,
		MessageInfos:      file_pkg_rating_rating_proto_msgTypes,
	}.Build()
	File_pkg_rating_rating_proto = out.File
	file_pkg_rating_rating_proto_rawDesc = nil
	file_pkg_rating_rating_proto_goTypes = nil
	file_pkg_rating_rating_proto_depIdxs = nil
}
