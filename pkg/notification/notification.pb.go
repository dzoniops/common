// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.23.4
// source: pkg/notification/notification.proto

package notification

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

type UserPreferences struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CreateNewReservation bool  `protobuf:"varint,1,opt,name=create_new_reservation,json=createNewReservation,proto3" json:"create_new_reservation,omitempty"`
	CancelReservation    bool  `protobuf:"varint,2,opt,name=cancel_reservation,json=cancelReservation,proto3" json:"cancel_reservation,omitempty"`
	RateHost             bool  `protobuf:"varint,3,opt,name=rate_host,json=rateHost,proto3" json:"rate_host,omitempty"`
	RateAccommodation    bool  `protobuf:"varint,4,opt,name=rate_accommodation,json=rateAccommodation,proto3" json:"rate_accommodation,omitempty"`
	ReservationAnswer    bool  `protobuf:"varint,5,opt,name=reservation_answer,json=reservationAnswer,proto3" json:"reservation_answer,omitempty"`
	UserId               int64 `protobuf:"varint,6,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *UserPreferences) Reset() {
	*x = UserPreferences{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_notification_notification_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserPreferences) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserPreferences) ProtoMessage() {}

func (x *UserPreferences) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_notification_notification_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserPreferences.ProtoReflect.Descriptor instead.
func (*UserPreferences) Descriptor() ([]byte, []int) {
	return file_pkg_notification_notification_proto_rawDescGZIP(), []int{0}
}

func (x *UserPreferences) GetCreateNewReservation() bool {
	if x != nil {
		return x.CreateNewReservation
	}
	return false
}

func (x *UserPreferences) GetCancelReservation() bool {
	if x != nil {
		return x.CancelReservation
	}
	return false
}

func (x *UserPreferences) GetRateHost() bool {
	if x != nil {
		return x.RateHost
	}
	return false
}

func (x *UserPreferences) GetRateAccommodation() bool {
	if x != nil {
		return x.RateAccommodation
	}
	return false
}

func (x *UserPreferences) GetReservationAnswer() bool {
	if x != nil {
		return x.ReservationAnswer
	}
	return false
}

func (x *UserPreferences) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type RequestReservationNotification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HostId        int64  `protobuf:"varint,1,opt,name=host_id,json=hostId,proto3" json:"host_id,omitempty"`
	Accommodation string `protobuf:"bytes,2,opt,name=accommodation,proto3" json:"accommodation,omitempty"`
}

func (x *RequestReservationNotification) Reset() {
	*x = RequestReservationNotification{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_notification_notification_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestReservationNotification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestReservationNotification) ProtoMessage() {}

func (x *RequestReservationNotification) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_notification_notification_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestReservationNotification.ProtoReflect.Descriptor instead.
func (*RequestReservationNotification) Descriptor() ([]byte, []int) {
	return file_pkg_notification_notification_proto_rawDescGZIP(), []int{1}
}

func (x *RequestReservationNotification) GetHostId() int64 {
	if x != nil {
		return x.HostId
	}
	return 0
}

func (x *RequestReservationNotification) GetAccommodation() string {
	if x != nil {
		return x.Accommodation
	}
	return ""
}

type NotificationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *NotificationRequest) Reset() {
	*x = NotificationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_notification_notification_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotificationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotificationRequest) ProtoMessage() {}

func (x *NotificationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_notification_notification_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotificationRequest.ProtoReflect.Descriptor instead.
func (*NotificationRequest) Descriptor() ([]byte, []int) {
	return file_pkg_notification_notification_proto_rawDescGZIP(), []int{2}
}

func (x *NotificationRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type NotificationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId  int64  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *NotificationResponse) Reset() {
	*x = NotificationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_notification_notification_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotificationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotificationResponse) ProtoMessage() {}

func (x *NotificationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_notification_notification_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotificationResponse.ProtoReflect.Descriptor instead.
func (*NotificationResponse) Descriptor() ([]byte, []int) {
	return file_pkg_notification_notification_proto_rawDescGZIP(), []int{3}
}

func (x *NotificationResponse) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *NotificationResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_pkg_notification_notification_proto protoreflect.FileDescriptor

var file_pkg_notification_notification_proto_rawDesc = []byte{
	0x0a, 0x23, 0x70, 0x6b, 0x67, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x22, 0x8a, 0x02, 0x0a, 0x0f, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x65, 0x66,
	0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x12, 0x34, 0x0a, 0x16, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x5f, 0x6e, 0x65, 0x77, 0x5f, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x14, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e,
	0x65, 0x77, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2d, 0x0a,
	0x12, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x5f, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x63, 0x61, 0x6e, 0x63, 0x65,
	0x6c, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09,
	0x72, 0x61, 0x74, 0x65, 0x5f, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x08, 0x72, 0x61, 0x74, 0x65, 0x48, 0x6f, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x12, 0x72, 0x61, 0x74,
	0x65, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x72, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2d, 0x0a, 0x12, 0x72, 0x65, 0x73, 0x65,
	0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x22, 0x5f, 0x0a, 0x1e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x65, 0x72,
	0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x17, 0x0a, 0x07, 0x68, 0x6f, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x68, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x61,
	0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0x2e, 0x0a, 0x13, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x22, 0x49, 0x0a, 0x14, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xb1, 0x02, 0x0a,
	0x13, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x68, 0x0a, 0x12, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2c, 0x2e, 0x6e, 0x6f, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x22, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x57,
	0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x65, 0x66,
	0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x12, 0x1d, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x65, 0x66, 0x65,
	0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x1a, 0x1d, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x65, 0x66, 0x65, 0x72,
	0x65, 0x6e, 0x63, 0x65, 0x73, 0x22, 0x00, 0x12, 0x57, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73,
	0x12, 0x1d, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x1a,
	0x1d, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x50, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x22, 0x00,
	0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64,
	0x7a, 0x6f, 0x6e, 0x69, 0x6f, 0x70, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70,
	0x6b, 0x67, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_notification_notification_proto_rawDescOnce sync.Once
	file_pkg_notification_notification_proto_rawDescData = file_pkg_notification_notification_proto_rawDesc
)

func file_pkg_notification_notification_proto_rawDescGZIP() []byte {
	file_pkg_notification_notification_proto_rawDescOnce.Do(func() {
		file_pkg_notification_notification_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_notification_notification_proto_rawDescData)
	})
	return file_pkg_notification_notification_proto_rawDescData
}

var file_pkg_notification_notification_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_pkg_notification_notification_proto_goTypes = []interface{}{
	(*UserPreferences)(nil),                // 0: notification.UserPreferences
	(*RequestReservationNotification)(nil), // 1: notification.RequestReservationNotification
	(*NotificationRequest)(nil),            // 2: notification.NotificationRequest
	(*NotificationResponse)(nil),           // 3: notification.NotificationResponse
}
var file_pkg_notification_notification_proto_depIdxs = []int32{
	1, // 0: notification.NotificationService.RequestReservation:input_type -> notification.RequestReservationNotification
	0, // 1: notification.NotificationService.UpdateUserPreferences:input_type -> notification.UserPreferences
	0, // 2: notification.NotificationService.CreateUserPreferences:input_type -> notification.UserPreferences
	3, // 3: notification.NotificationService.RequestReservation:output_type -> notification.NotificationResponse
	0, // 4: notification.NotificationService.UpdateUserPreferences:output_type -> notification.UserPreferences
	0, // 5: notification.NotificationService.CreateUserPreferences:output_type -> notification.UserPreferences
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_notification_notification_proto_init() }
func file_pkg_notification_notification_proto_init() {
	if File_pkg_notification_notification_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_notification_notification_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserPreferences); i {
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
		file_pkg_notification_notification_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestReservationNotification); i {
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
		file_pkg_notification_notification_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotificationRequest); i {
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
		file_pkg_notification_notification_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotificationResponse); i {
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
			RawDescriptor: file_pkg_notification_notification_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_notification_notification_proto_goTypes,
		DependencyIndexes: file_pkg_notification_notification_proto_depIdxs,
		MessageInfos:      file_pkg_notification_notification_proto_msgTypes,
	}.Build()
	File_pkg_notification_notification_proto = out.File
	file_pkg_notification_notification_proto_rawDesc = nil
	file_pkg_notification_notification_proto_goTypes = nil
	file_pkg_notification_notification_proto_depIdxs = nil
}
