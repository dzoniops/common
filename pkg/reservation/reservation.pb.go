// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.23.4
// source: pkg/reservation/reservation.proto

package reservation

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ReservationStatus int32

const (
	ReservationStatus_RESERVATION_STATUS_UNSPECIFIED ReservationStatus = 0
	ReservationStatus_RESERVATION_STATUS_ACTIVE      ReservationStatus = 1
	ReservationStatus_RESERVATION_STATUS_ACCEPTED    ReservationStatus = 2
	ReservationStatus_RESERVATION_STATUS_DECLINED    ReservationStatus = 3
)

// Enum value maps for ReservationStatus.
var (
	ReservationStatus_name = map[int32]string{
		0: "RESERVATION_STATUS_UNSPECIFIED",
		1: "RESERVATION_STATUS_ACTIVE",
		2: "RESERVATION_STATUS_ACCEPTED",
		3: "RESERVATION_STATUS_DECLINED",
	}
	ReservationStatus_value = map[string]int32{
		"RESERVATION_STATUS_UNSPECIFIED": 0,
		"RESERVATION_STATUS_ACTIVE":      1,
		"RESERVATION_STATUS_ACCEPTED":    2,
		"RESERVATION_STATUS_DECLINED":    3,
	}
)

func (x ReservationStatus) Enum() *ReservationStatus {
	p := new(ReservationStatus)
	*p = x
	return p
}

func (x ReservationStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ReservationStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_pkg_reservation_reservation_proto_enumTypes[0].Descriptor()
}

func (ReservationStatus) Type() protoreflect.EnumType {
	return &file_pkg_reservation_reservation_proto_enumTypes[0]
}

func (x ReservationStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ReservationStatus.Descriptor instead.
func (ReservationStatus) EnumDescriptor() ([]byte, []int) {
	return file_pkg_reservation_reservation_proto_rawDescGZIP(), []int{0}
}

type IdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *IdRequest) Reset() {
	*x = IdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_reservation_reservation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdRequest) ProtoMessage() {}

func (x *IdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_reservation_reservation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdRequest.ProtoReflect.Descriptor instead.
func (*IdRequest) Descriptor() ([]byte, []int) {
	return file_pkg_reservation_reservation_proto_rawDescGZIP(), []int{0}
}

func (x *IdRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type Reservation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccommodationId int64                  `protobuf:"varint,1,opt,name=accommodation_id,json=accommodationId,proto3" json:"accommodation_id,omitempty"`
	StartDate       *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"`
	EndDate         *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=end_date,json=endDate,proto3" json:"end_date,omitempty"`
	NumberOfGuests  int64                  `protobuf:"varint,4,opt,name=number_of_guests,json=numberOfGuests,proto3" json:"number_of_guests,omitempty"`
	Status          ReservationStatus      `protobuf:"varint,5,opt,name=status,proto3,enum=reservation.ReservationStatus" json:"status,omitempty"`
}

func (x *Reservation) Reset() {
	*x = Reservation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_reservation_reservation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Reservation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Reservation) ProtoMessage() {}

func (x *Reservation) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_reservation_reservation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Reservation.ProtoReflect.Descriptor instead.
func (*Reservation) Descriptor() ([]byte, []int) {
	return file_pkg_reservation_reservation_proto_rawDescGZIP(), []int{1}
}

func (x *Reservation) GetAccommodationId() int64 {
	if x != nil {
		return x.AccommodationId
	}
	return 0
}

func (x *Reservation) GetStartDate() *timestamppb.Timestamp {
	if x != nil {
		return x.StartDate
	}
	return nil
}

func (x *Reservation) GetEndDate() *timestamppb.Timestamp {
	if x != nil {
		return x.EndDate
	}
	return nil
}

func (x *Reservation) GetNumberOfGuests() int64 {
	if x != nil {
		return x.NumberOfGuests
	}
	return 0
}

func (x *Reservation) GetStatus() ReservationStatus {
	if x != nil {
		return x.Status
	}
	return ReservationStatus_RESERVATION_STATUS_UNSPECIFIED
}

type ActiveReservationsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reservations []*Reservation `protobuf:"bytes,1,rep,name=reservations,proto3" json:"reservations,omitempty"`
}

func (x *ActiveReservationsResponse) Reset() {
	*x = ActiveReservationsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_reservation_reservation_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActiveReservationsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActiveReservationsResponse) ProtoMessage() {}

func (x *ActiveReservationsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_reservation_reservation_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActiveReservationsResponse.ProtoReflect.Descriptor instead.
func (*ActiveReservationsResponse) Descriptor() ([]byte, []int) {
	return file_pkg_reservation_reservation_proto_rawDescGZIP(), []int{2}
}

func (x *ActiveReservationsResponse) GetReservations() []*Reservation {
	if x != nil {
		return x.Reservations
	}
	return nil
}

var File_pkg_reservation_reservation_proto protoreflect.FileDescriptor

var file_pkg_reservation_reservation_proto_rawDesc = []byte{
	0x0a, 0x21, 0x70, 0x6b, 0x67, 0x2f, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x1b, 0x0a, 0x09, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x8c,
	0x02, 0x0a, 0x0b, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x29,
	0x0a, 0x10, 0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x39, 0x0a, 0x0a, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x44, 0x61, 0x74, 0x65, 0x12, 0x35, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x28, 0x0a, 0x10, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x6f, 0x66, 0x5f, 0x67, 0x75, 0x65, 0x73, 0x74, 0x73, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x4f, 0x66, 0x47,
	0x75, 0x65, 0x73, 0x74, 0x73, 0x12, 0x36, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x5a, 0x0a,
	0x1a, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x0c, 0x72,
	0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x18, 0x2e, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x72, 0x65, 0x73,
	0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2a, 0x98, 0x01, 0x0a, 0x11, 0x52, 0x65,
	0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x22, 0x0a, 0x1e, 0x52, 0x45, 0x53, 0x45, 0x52, 0x56, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53,
	0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x1d, 0x0a, 0x19, 0x52, 0x45, 0x53, 0x45, 0x52, 0x56, 0x41, 0x54, 0x49,
	0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x56, 0x45,
	0x10, 0x01, 0x12, 0x1f, 0x0a, 0x1b, 0x52, 0x45, 0x53, 0x45, 0x52, 0x56, 0x41, 0x54, 0x49, 0x4f,
	0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x41, 0x43, 0x43, 0x45, 0x50, 0x54, 0x45,
	0x44, 0x10, 0x02, 0x12, 0x1f, 0x0a, 0x1b, 0x52, 0x45, 0x53, 0x45, 0x52, 0x56, 0x41, 0x54, 0x49,
	0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x44, 0x45, 0x43, 0x4c, 0x49, 0x4e,
	0x45, 0x44, 0x10, 0x03, 0x32, 0xd3, 0x01, 0x0a, 0x12, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5e, 0x0a, 0x19, 0x41,
	0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x47, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x2e, 0x72, 0x65, 0x73, 0x65, 0x72,
	0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x27, 0x2e, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x41,
	0x63, 0x74, 0x69, 0x76, 0x65, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5d, 0x0a, 0x18, 0x41,
	0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x48, 0x6f, 0x73, 0x74, 0x12, 0x16, 0x2e, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x27, 0x2e, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x41, 0x63,
	0x74, 0x69, 0x76, 0x65, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x7a, 0x6f, 0x6e, 0x69, 0x6f, 0x70,
	0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x72,
	0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_pkg_reservation_reservation_proto_rawDescOnce sync.Once
	file_pkg_reservation_reservation_proto_rawDescData = file_pkg_reservation_reservation_proto_rawDesc
)

func file_pkg_reservation_reservation_proto_rawDescGZIP() []byte {
	file_pkg_reservation_reservation_proto_rawDescOnce.Do(func() {
		file_pkg_reservation_reservation_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_reservation_reservation_proto_rawDescData)
	})
	return file_pkg_reservation_reservation_proto_rawDescData
}

var file_pkg_reservation_reservation_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_pkg_reservation_reservation_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_pkg_reservation_reservation_proto_goTypes = []interface{}{
	(ReservationStatus)(0),             // 0: reservation.ReservationStatus
	(*IdRequest)(nil),                  // 1: reservation.IdRequest
	(*Reservation)(nil),                // 2: reservation.Reservation
	(*ActiveReservationsResponse)(nil), // 3: reservation.ActiveReservationsResponse
	(*timestamppb.Timestamp)(nil),      // 4: google.protobuf.Timestamp
}
var file_pkg_reservation_reservation_proto_depIdxs = []int32{
	4, // 0: reservation.Reservation.start_date:type_name -> google.protobuf.Timestamp
	4, // 1: reservation.Reservation.end_date:type_name -> google.protobuf.Timestamp
	0, // 2: reservation.Reservation.status:type_name -> reservation.ReservationStatus
	2, // 3: reservation.ActiveReservationsResponse.reservations:type_name -> reservation.Reservation
	1, // 4: reservation.ReservationService.ActivateReservationsGuest:input_type -> reservation.IdRequest
	1, // 5: reservation.ReservationService.ActivateReservationsHost:input_type -> reservation.IdRequest
	3, // 6: reservation.ReservationService.ActivateReservationsGuest:output_type -> reservation.ActiveReservationsResponse
	3, // 7: reservation.ReservationService.ActivateReservationsHost:output_type -> reservation.ActiveReservationsResponse
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_pkg_reservation_reservation_proto_init() }
func file_pkg_reservation_reservation_proto_init() {
	if File_pkg_reservation_reservation_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_reservation_reservation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IdRequest); i {
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
		file_pkg_reservation_reservation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Reservation); i {
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
		file_pkg_reservation_reservation_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActiveReservationsResponse); i {
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
			RawDescriptor: file_pkg_reservation_reservation_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_reservation_reservation_proto_goTypes,
		DependencyIndexes: file_pkg_reservation_reservation_proto_depIdxs,
		EnumInfos:         file_pkg_reservation_reservation_proto_enumTypes,
		MessageInfos:      file_pkg_reservation_reservation_proto_msgTypes,
	}.Build()
	File_pkg_reservation_reservation_proto = out.File
	file_pkg_reservation_reservation_proto_rawDesc = nil
	file_pkg_reservation_reservation_proto_goTypes = nil
	file_pkg_reservation_reservation_proto_depIdxs = nil
}
