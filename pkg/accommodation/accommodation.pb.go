// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pkg/accommodation/accommodation.proto

package accommodation

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type AccommodationRequest struct {
	HostId               int64                 `protobuf:"varint,1,opt,name=host_id,json=hostId,proto3" json:"host_id,omitempty"`
	Name                 string                `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Town                 string                `protobuf:"bytes,3,opt,name=town,proto3" json:"town,omitempty"`
	Municipality         string                `protobuf:"bytes,4,opt,name=municipality,proto3" json:"municipality,omitempty"`
	Country              string                `protobuf:"bytes,5,opt,name=country,proto3" json:"country,omitempty"`
	Amenities            string                `protobuf:"bytes,6,opt,name=amenities,proto3" json:"amenities,omitempty"`
	MinGuests            int64                 `protobuf:"varint,7,opt,name=min_guests,json=minGuests,proto3" json:"min_guests,omitempty"`
	MaxGuests            int64                 `protobuf:"varint,8,opt,name=max_guests,json=maxGuests,proto3" json:"max_guests,omitempty"`
	PricingModel         string                `protobuf:"bytes,9,opt,name=pricing_model,json=pricingModel,proto3" json:"pricing_model,omitempty"`
	ReservationModel     string                `protobuf:"bytes,10,opt,name=reservation_model,json=reservationModel,proto3" json:"reservation_model,omitempty"`
	Images               []*AccommodationImage `protobuf:"bytes,11,rep,name=images,proto3" json:"images,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *AccommodationRequest) Reset()         { *m = AccommodationRequest{} }
func (m *AccommodationRequest) String() string { return proto.CompactTextString(m) }
func (*AccommodationRequest) ProtoMessage()    {}
func (*AccommodationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_60c033c297702363, []int{0}
}

func (m *AccommodationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccommodationRequest.Unmarshal(m, b)
}
func (m *AccommodationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccommodationRequest.Marshal(b, m, deterministic)
}
func (m *AccommodationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccommodationRequest.Merge(m, src)
}
func (m *AccommodationRequest) XXX_Size() int {
	return xxx_messageInfo_AccommodationRequest.Size(m)
}
func (m *AccommodationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AccommodationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AccommodationRequest proto.InternalMessageInfo

func (m *AccommodationRequest) GetHostId() int64 {
	if m != nil {
		return m.HostId
	}
	return 0
}

func (m *AccommodationRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AccommodationRequest) GetTown() string {
	if m != nil {
		return m.Town
	}
	return ""
}

func (m *AccommodationRequest) GetMunicipality() string {
	if m != nil {
		return m.Municipality
	}
	return ""
}

func (m *AccommodationRequest) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

func (m *AccommodationRequest) GetAmenities() string {
	if m != nil {
		return m.Amenities
	}
	return ""
}

func (m *AccommodationRequest) GetMinGuests() int64 {
	if m != nil {
		return m.MinGuests
	}
	return 0
}

func (m *AccommodationRequest) GetMaxGuests() int64 {
	if m != nil {
		return m.MaxGuests
	}
	return 0
}

func (m *AccommodationRequest) GetPricingModel() string {
	if m != nil {
		return m.PricingModel
	}
	return ""
}

func (m *AccommodationRequest) GetReservationModel() string {
	if m != nil {
		return m.ReservationModel
	}
	return ""
}

func (m *AccommodationRequest) GetImages() []*AccommodationImage {
	if m != nil {
		return m.Images
	}
	return nil
}

type AccommodationImage struct {
	B64Img               string   `protobuf:"bytes,1,opt,name=b64img,proto3" json:"b64img,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccommodationImage) Reset()         { *m = AccommodationImage{} }
func (m *AccommodationImage) String() string { return proto.CompactTextString(m) }
func (*AccommodationImage) ProtoMessage()    {}
func (*AccommodationImage) Descriptor() ([]byte, []int) {
	return fileDescriptor_60c033c297702363, []int{1}
}

func (m *AccommodationImage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccommodationImage.Unmarshal(m, b)
}
func (m *AccommodationImage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccommodationImage.Marshal(b, m, deterministic)
}
func (m *AccommodationImage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccommodationImage.Merge(m, src)
}
func (m *AccommodationImage) XXX_Size() int {
	return xxx_messageInfo_AccommodationImage.Size(m)
}
func (m *AccommodationImage) XXX_DiscardUnknown() {
	xxx_messageInfo_AccommodationImage.DiscardUnknown(m)
}

var xxx_messageInfo_AccommodationImage proto.InternalMessageInfo

func (m *AccommodationImage) GetB64Img() string {
	if m != nil {
		return m.B64Img
	}
	return ""
}

type AccommodationResponse struct {
	AccommodationId      int64    `protobuf:"varint,1,opt,name=accommodation_id,json=accommodationId,proto3" json:"accommodation_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccommodationResponse) Reset()         { *m = AccommodationResponse{} }
func (m *AccommodationResponse) String() string { return proto.CompactTextString(m) }
func (*AccommodationResponse) ProtoMessage()    {}
func (*AccommodationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_60c033c297702363, []int{2}
}

func (m *AccommodationResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccommodationResponse.Unmarshal(m, b)
}
func (m *AccommodationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccommodationResponse.Marshal(b, m, deterministic)
}
func (m *AccommodationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccommodationResponse.Merge(m, src)
}
func (m *AccommodationResponse) XXX_Size() int {
	return xxx_messageInfo_AccommodationResponse.Size(m)
}
func (m *AccommodationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AccommodationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AccommodationResponse proto.InternalMessageInfo

func (m *AccommodationResponse) GetAccommodationId() int64 {
	if m != nil {
		return m.AccommodationId
	}
	return 0
}

type AccommodationInfo struct {
	HostId               int64                         `protobuf:"varint,1,opt,name=host_id,json=hostId,proto3" json:"host_id,omitempty"`
	Name                 string                        `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Town                 string                        `protobuf:"bytes,3,opt,name=town,proto3" json:"town,omitempty"`
	Municipality         string                        `protobuf:"bytes,4,opt,name=municipality,proto3" json:"municipality,omitempty"`
	Country              string                        `protobuf:"bytes,5,opt,name=country,proto3" json:"country,omitempty"`
	Amenities            string                        `protobuf:"bytes,6,opt,name=amenities,proto3" json:"amenities,omitempty"`
	MinGuests            int64                         `protobuf:"varint,7,opt,name=min_guests,json=minGuests,proto3" json:"min_guests,omitempty"`
	MaxGuests            int64                         `protobuf:"varint,8,opt,name=max_guests,json=maxGuests,proto3" json:"max_guests,omitempty"`
	PricingModel         string                        `protobuf:"bytes,9,opt,name=pricing_model,json=pricingModel,proto3" json:"pricing_model,omitempty"`
	ReservationModel     string                        `protobuf:"bytes,10,opt,name=reservation_model,json=reservationModel,proto3" json:"reservation_model,omitempty"`
	Images               []*AccommodationImageResponse `protobuf:"bytes,11,rep,name=images,proto3" json:"images,omitempty"`
	Id                   int64                         `protobuf:"varint,12,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *AccommodationInfo) Reset()         { *m = AccommodationInfo{} }
func (m *AccommodationInfo) String() string { return proto.CompactTextString(m) }
func (*AccommodationInfo) ProtoMessage()    {}
func (*AccommodationInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_60c033c297702363, []int{3}
}

func (m *AccommodationInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccommodationInfo.Unmarshal(m, b)
}
func (m *AccommodationInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccommodationInfo.Marshal(b, m, deterministic)
}
func (m *AccommodationInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccommodationInfo.Merge(m, src)
}
func (m *AccommodationInfo) XXX_Size() int {
	return xxx_messageInfo_AccommodationInfo.Size(m)
}
func (m *AccommodationInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_AccommodationInfo.DiscardUnknown(m)
}

var xxx_messageInfo_AccommodationInfo proto.InternalMessageInfo

func (m *AccommodationInfo) GetHostId() int64 {
	if m != nil {
		return m.HostId
	}
	return 0
}

func (m *AccommodationInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AccommodationInfo) GetTown() string {
	if m != nil {
		return m.Town
	}
	return ""
}

func (m *AccommodationInfo) GetMunicipality() string {
	if m != nil {
		return m.Municipality
	}
	return ""
}

func (m *AccommodationInfo) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

func (m *AccommodationInfo) GetAmenities() string {
	if m != nil {
		return m.Amenities
	}
	return ""
}

func (m *AccommodationInfo) GetMinGuests() int64 {
	if m != nil {
		return m.MinGuests
	}
	return 0
}

func (m *AccommodationInfo) GetMaxGuests() int64 {
	if m != nil {
		return m.MaxGuests
	}
	return 0
}

func (m *AccommodationInfo) GetPricingModel() string {
	if m != nil {
		return m.PricingModel
	}
	return ""
}

func (m *AccommodationInfo) GetReservationModel() string {
	if m != nil {
		return m.ReservationModel
	}
	return ""
}

func (m *AccommodationInfo) GetImages() []*AccommodationImageResponse {
	if m != nil {
		return m.Images
	}
	return nil
}

func (m *AccommodationInfo) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type AccommodationImageResponse struct {
	B64Img               string   `protobuf:"bytes,1,opt,name=b64img,proto3" json:"b64img,omitempty"`
	ImageId              int64    `protobuf:"varint,2,opt,name=image_id,json=imageId,proto3" json:"image_id,omitempty"`
	AccommodationId      int64    `protobuf:"varint,3,opt,name=accommodation_id,json=accommodationId,proto3" json:"accommodation_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccommodationImageResponse) Reset()         { *m = AccommodationImageResponse{} }
func (m *AccommodationImageResponse) String() string { return proto.CompactTextString(m) }
func (*AccommodationImageResponse) ProtoMessage()    {}
func (*AccommodationImageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_60c033c297702363, []int{4}
}

func (m *AccommodationImageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccommodationImageResponse.Unmarshal(m, b)
}
func (m *AccommodationImageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccommodationImageResponse.Marshal(b, m, deterministic)
}
func (m *AccommodationImageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccommodationImageResponse.Merge(m, src)
}
func (m *AccommodationImageResponse) XXX_Size() int {
	return xxx_messageInfo_AccommodationImageResponse.Size(m)
}
func (m *AccommodationImageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AccommodationImageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AccommodationImageResponse proto.InternalMessageInfo

func (m *AccommodationImageResponse) GetB64Img() string {
	if m != nil {
		return m.B64Img
	}
	return ""
}

func (m *AccommodationImageResponse) GetImageId() int64 {
	if m != nil {
		return m.ImageId
	}
	return 0
}

func (m *AccommodationImageResponse) GetAccommodationId() int64 {
	if m != nil {
		return m.AccommodationId
	}
	return 0
}

type AccommodationSearchRequest struct {
	Town                 string   `protobuf:"bytes,1,opt,name=town,proto3" json:"town,omitempty"`
	Municipality         string   `protobuf:"bytes,2,opt,name=municipality,proto3" json:"municipality,omitempty"`
	Country              string   `protobuf:"bytes,3,opt,name=country,proto3" json:"country,omitempty"`
	MinGuests            int64    `protobuf:"varint,4,opt,name=min_guests,json=minGuests,proto3" json:"min_guests,omitempty"`
	MaxGuests            int64    `protobuf:"varint,5,opt,name=max_guests,json=maxGuests,proto3" json:"max_guests,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccommodationSearchRequest) Reset()         { *m = AccommodationSearchRequest{} }
func (m *AccommodationSearchRequest) String() string { return proto.CompactTextString(m) }
func (*AccommodationSearchRequest) ProtoMessage()    {}
func (*AccommodationSearchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_60c033c297702363, []int{5}
}

func (m *AccommodationSearchRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccommodationSearchRequest.Unmarshal(m, b)
}
func (m *AccommodationSearchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccommodationSearchRequest.Marshal(b, m, deterministic)
}
func (m *AccommodationSearchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccommodationSearchRequest.Merge(m, src)
}
func (m *AccommodationSearchRequest) XXX_Size() int {
	return xxx_messageInfo_AccommodationSearchRequest.Size(m)
}
func (m *AccommodationSearchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AccommodationSearchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AccommodationSearchRequest proto.InternalMessageInfo

func (m *AccommodationSearchRequest) GetTown() string {
	if m != nil {
		return m.Town
	}
	return ""
}

func (m *AccommodationSearchRequest) GetMunicipality() string {
	if m != nil {
		return m.Municipality
	}
	return ""
}

func (m *AccommodationSearchRequest) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

func (m *AccommodationSearchRequest) GetMinGuests() int64 {
	if m != nil {
		return m.MinGuests
	}
	return 0
}

func (m *AccommodationSearchRequest) GetMaxGuests() int64 {
	if m != nil {
		return m.MaxGuests
	}
	return 0
}

type AccommodationSearchResponse struct {
	AccomomodationList   []*AccommodationInfo `protobuf:"bytes,1,rep,name=accomomodation_list,json=accomomodationList,proto3" json:"accomomodation_list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *AccommodationSearchResponse) Reset()         { *m = AccommodationSearchResponse{} }
func (m *AccommodationSearchResponse) String() string { return proto.CompactTextString(m) }
func (*AccommodationSearchResponse) ProtoMessage()    {}
func (*AccommodationSearchResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_60c033c297702363, []int{6}
}

func (m *AccommodationSearchResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccommodationSearchResponse.Unmarshal(m, b)
}
func (m *AccommodationSearchResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccommodationSearchResponse.Marshal(b, m, deterministic)
}
func (m *AccommodationSearchResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccommodationSearchResponse.Merge(m, src)
}
func (m *AccommodationSearchResponse) XXX_Size() int {
	return xxx_messageInfo_AccommodationSearchResponse.Size(m)
}
func (m *AccommodationSearchResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AccommodationSearchResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AccommodationSearchResponse proto.InternalMessageInfo

func (m *AccommodationSearchResponse) GetAccomomodationList() []*AccommodationInfo {
	if m != nil {
		return m.AccomomodationList
	}
	return nil
}

func init() {
	proto.RegisterType((*AccommodationRequest)(nil), "accommodation.AccommodationRequest")
	proto.RegisterType((*AccommodationImage)(nil), "accommodation.AccommodationImage")
	proto.RegisterType((*AccommodationResponse)(nil), "accommodation.AccommodationResponse")
	proto.RegisterType((*AccommodationInfo)(nil), "accommodation.AccommodationInfo")
	proto.RegisterType((*AccommodationImageResponse)(nil), "accommodation.AccommodationImageResponse")
	proto.RegisterType((*AccommodationSearchRequest)(nil), "accommodation.AccommodationSearchRequest")
	proto.RegisterType((*AccommodationSearchResponse)(nil), "accommodation.AccommodationSearchResponse")
}

func init() {
	proto.RegisterFile("pkg/accommodation/accommodation.proto", fileDescriptor_60c033c297702363)
}

var fileDescriptor_60c033c297702363 = []byte{
	// 563 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x95, 0xc1, 0x6e, 0xd3, 0x4c,
	0x10, 0xc7, 0x6b, 0xbb, 0x4d, 0x9a, 0x69, 0xfb, 0x7d, 0xed, 0xa4, 0xc0, 0x12, 0x40, 0x0a, 0x2e,
	0x48, 0x29, 0x54, 0x89, 0x54, 0x10, 0x12, 0xc7, 0x86, 0x43, 0x15, 0x09, 0x0e, 0xb8, 0x37, 0x2e,
	0xc1, 0xb1, 0xb7, 0xce, 0x8a, 0xec, 0xae, 0xf1, 0x6e, 0x4a, 0xd3, 0xd7, 0xe0, 0x35, 0x78, 0x14,
	0x9e, 0x09, 0x21, 0x6f, 0x6c, 0xda, 0x4d, 0x53, 0xb7, 0xdc, 0xb9, 0x79, 0xff, 0xff, 0x19, 0xef,
	0xcc, 0xfc, 0x46, 0x36, 0x3c, 0x4f, 0xbf, 0x24, 0xbd, 0x30, 0x8a, 0x24, 0xe7, 0x32, 0x0e, 0x35,
	0x93, 0xc2, 0x3e, 0x75, 0xd3, 0x4c, 0x6a, 0x89, 0x5b, 0x96, 0xe8, 0xff, 0x72, 0x61, 0xf7, 0xe8,
	0xaa, 0x12, 0xd0, 0xaf, 0x53, 0xaa, 0x34, 0x3e, 0x80, 0xfa, 0x58, 0x2a, 0x3d, 0x64, 0x31, 0x71,
	0xda, 0x4e, 0xc7, 0x0b, 0x6a, 0xf9, 0x71, 0x10, 0x23, 0xc2, 0xaa, 0x08, 0x39, 0x25, 0x6e, 0xdb,
	0xe9, 0x34, 0x02, 0xf3, 0x9c, 0x6b, 0x5a, 0x7e, 0x13, 0xc4, 0x9b, 0x6b, 0xf9, 0x33, 0xfa, 0xb0,
	0xc9, 0xa7, 0x82, 0x45, 0x2c, 0x0d, 0x27, 0x4c, 0xcf, 0xc8, 0xaa, 0xf1, 0x2c, 0x0d, 0x09, 0xd4,
	0x23, 0x39, 0x15, 0x3a, 0x9b, 0x91, 0x35, 0x63, 0x97, 0x47, 0x7c, 0x0c, 0x8d, 0x90, 0x53, 0xc1,
	0x34, 0xa3, 0x8a, 0xd4, 0x8c, 0x77, 0x29, 0xe0, 0x13, 0x00, 0xce, 0xc4, 0x30, 0xc9, 0x2b, 0x55,
	0xa4, 0x6e, 0xea, 0x6b, 0x70, 0x26, 0x8e, 0x8d, 0x60, 0xec, 0xf0, 0xbc, 0xb4, 0xd7, 0x0b, 0x3b,
	0x3c, 0x2f, 0xec, 0x3d, 0xd8, 0x4a, 0x33, 0x16, 0x31, 0x91, 0x0c, 0xb9, 0x8c, 0xe9, 0x84, 0x34,
	0xe6, 0xa5, 0x15, 0xe2, 0x87, 0x5c, 0xc3, 0x97, 0xb0, 0x93, 0x51, 0x45, 0xb3, 0x33, 0x33, 0x95,
	0x22, 0x10, 0x4c, 0xe0, 0xf6, 0x15, 0x63, 0x1e, 0xfc, 0x16, 0x6a, 0x8c, 0x87, 0x09, 0x55, 0x64,
	0xa3, 0xed, 0x75, 0x36, 0x0e, 0x9f, 0x76, 0xed, 0xd1, 0x5b, 0x13, 0x1e, 0xe4, 0x91, 0x41, 0x91,
	0xe0, 0x1f, 0x00, 0x5e, 0x77, 0xf1, 0x3e, 0xd4, 0x46, 0x6f, 0x5e, 0x33, 0x9e, 0x98, 0xe1, 0x37,
	0x82, 0xe2, 0xe4, 0xf7, 0xe1, 0xde, 0x02, 0x2d, 0x95, 0x4a, 0xa1, 0x28, 0xee, 0xc3, 0xb6, 0x75,
	0xe5, 0x25, 0xb7, 0xff, 0x2d, 0x7d, 0x10, 0xfb, 0xdf, 0x3d, 0xd8, 0xb1, 0xaf, 0x14, 0xa7, 0xf2,
	0x1f, 0xef, 0xbf, 0xe1, 0x7d, 0xb4, 0xc0, 0x7b, 0xff, 0x76, 0xde, 0x05, 0xa8, 0x92, 0x3b, 0xfe,
	0x07, 0x2e, 0x8b, 0xc9, 0xa6, 0xa9, 0xd5, 0x65, 0xb1, 0x7f, 0x01, 0xad, 0x9b, 0xb3, 0x6e, 0xda,
	0x07, 0x7c, 0x08, 0xeb, 0xe6, 0x7d, 0x39, 0x36, 0xd7, 0xbc, 0xab, 0x6e, 0xce, 0x83, 0x78, 0xe9,
	0x46, 0x78, 0xcb, 0x37, 0xe2, 0x87, 0xb3, 0x70, 0xf9, 0x09, 0x0d, 0xb3, 0x68, 0x5c, 0x7e, 0x0a,
	0x4a, 0xda, 0x4e, 0x05, 0x6d, 0xb7, 0x9a, 0xb6, 0x67, 0xd3, 0xb6, 0x79, 0xae, 0x56, 0xf3, 0x5c,
	0x5b, 0xe0, 0xe9, 0xa7, 0xf0, 0x68, 0x69, 0xb5, 0xc5, 0xac, 0x3e, 0x42, 0xd3, 0x34, 0x28, 0xff,
	0x74, 0x3e, 0x61, 0x4a, 0x13, 0xc7, 0x90, 0x6a, 0x57, 0x92, 0x12, 0xa7, 0x32, 0x40, 0x3b, 0xf9,
	0x3d, 0x53, 0xfa, 0xf0, 0xe7, 0xe2, 0x57, 0xf2, 0x84, 0x66, 0x67, 0x2c, 0xa2, 0x38, 0x82, 0xe6,
	0xbb, 0x8c, 0x86, 0x9a, 0x5a, 0x2e, 0xee, 0x55, 0xdd, 0x52, 0x8c, 0xb5, 0xf5, 0xac, 0x3a, 0x68,
	0xde, 0x8d, 0xbf, 0x82, 0x9f, 0x61, 0xf7, 0x98, 0x6a, 0xcb, 0xed, 0xcf, 0x06, 0x31, 0xde, 0x29,
	0xbf, 0x75, 0x6b, 0xc3, 0xfe, 0x0a, 0x0a, 0x68, 0x2e, 0x19, 0x28, 0x56, 0x6e, 0xb5, 0xb5, 0x22,
	0xad, 0x17, 0x77, 0x09, 0x2d, 0x3b, 0xea, 0x77, 0x3f, 0x1d, 0x24, 0x4c, 0x8f, 0xa7, 0xa3, 0x6e,
	0x24, 0x79, 0x2f, 0xbe, 0x90, 0x82, 0xc9, 0x54, 0xf5, 0x4c, 0x8a, 0xe8, 0x5d, 0xfb, 0x8f, 0x8d,
	0x6a, 0xe6, 0xd7, 0xf5, 0xea, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x29, 0x5b, 0xbd, 0xd1, 0xe3,
	0x06, 0x00, 0x00,
}
