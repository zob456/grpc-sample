// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: airport-service/proto/air.proto

package proto

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

type AirportDetailsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AirportID string `protobuf:"bytes,1,opt,name=AirportID,proto3" json:"AirportID,omitempty"`
}

func (x *AirportDetailsReq) Reset() {
	*x = AirportDetailsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_airport_service_proto_air_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AirportDetailsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AirportDetailsReq) ProtoMessage() {}

func (x *AirportDetailsReq) ProtoReflect() protoreflect.Message {
	mi := &file_airport_service_proto_air_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AirportDetailsReq.ProtoReflect.Descriptor instead.
func (*AirportDetailsReq) Descriptor() ([]byte, []int) {
	return file_airport_service_proto_air_proto_rawDescGZIP(), []int{0}
}

func (x *AirportDetailsReq) GetAirportID() string {
	if x != nil {
		return x.AirportID
	}
	return ""
}

type AirportDetailsRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AirportID   string  `protobuf:"bytes,1,opt,name=AirportID,proto3" json:"AirportID,omitempty"`
	AirportName string  `protobuf:"bytes,2,opt,name=AirportName,proto3" json:"AirportName,omitempty"`
	City        string  `protobuf:"bytes,3,opt,name=City,proto3" json:"City,omitempty"`
	Country     string  `protobuf:"bytes,4,opt,name=Country,proto3" json:"Country,omitempty"`
	IATACode    string  `protobuf:"bytes,5,opt,name=IATACode,proto3" json:"IATACode,omitempty"`
	ICAOCode    string  `protobuf:"bytes,6,opt,name=ICAOCode,proto3" json:"ICAOCode,omitempty"`
	Latitude    float64 `protobuf:"fixed64,7,opt,name=Latitude,proto3" json:"Latitude,omitempty"`
	Longitude   float64 `protobuf:"fixed64,8,opt,name=Longitude,proto3" json:"Longitude,omitempty"`
	Altitude    float64 `protobuf:"fixed64,9,opt,name=Altitude,proto3" json:"Altitude,omitempty"`
	TimeZone    string  `protobuf:"bytes,10,opt,name=TimeZone,proto3" json:"TimeZone,omitempty"`
}

func (x *AirportDetailsRes) Reset() {
	*x = AirportDetailsRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_airport_service_proto_air_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AirportDetailsRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AirportDetailsRes) ProtoMessage() {}

func (x *AirportDetailsRes) ProtoReflect() protoreflect.Message {
	mi := &file_airport_service_proto_air_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AirportDetailsRes.ProtoReflect.Descriptor instead.
func (*AirportDetailsRes) Descriptor() ([]byte, []int) {
	return file_airport_service_proto_air_proto_rawDescGZIP(), []int{1}
}

func (x *AirportDetailsRes) GetAirportID() string {
	if x != nil {
		return x.AirportID
	}
	return ""
}

func (x *AirportDetailsRes) GetAirportName() string {
	if x != nil {
		return x.AirportName
	}
	return ""
}

func (x *AirportDetailsRes) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *AirportDetailsRes) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *AirportDetailsRes) GetIATACode() string {
	if x != nil {
		return x.IATACode
	}
	return ""
}

func (x *AirportDetailsRes) GetICAOCode() string {
	if x != nil {
		return x.ICAOCode
	}
	return ""
}

func (x *AirportDetailsRes) GetLatitude() float64 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *AirportDetailsRes) GetLongitude() float64 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

func (x *AirportDetailsRes) GetAltitude() float64 {
	if x != nil {
		return x.Altitude
	}
	return 0
}

func (x *AirportDetailsRes) GetTimeZone() string {
	if x != nil {
		return x.TimeZone
	}
	return ""
}

type AirportDistanceReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstAirportID  string `protobuf:"bytes,1,opt,name=FirstAirportID,proto3" json:"FirstAirportID,omitempty"`
	SecondAirportID string `protobuf:"bytes,2,opt,name=SecondAirportID,proto3" json:"SecondAirportID,omitempty"`
}

func (x *AirportDistanceReq) Reset() {
	*x = AirportDistanceReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_airport_service_proto_air_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AirportDistanceReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AirportDistanceReq) ProtoMessage() {}

func (x *AirportDistanceReq) ProtoReflect() protoreflect.Message {
	mi := &file_airport_service_proto_air_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AirportDistanceReq.ProtoReflect.Descriptor instead.
func (*AirportDistanceReq) Descriptor() ([]byte, []int) {
	return file_airport_service_proto_air_proto_rawDescGZIP(), []int{2}
}

func (x *AirportDistanceReq) GetFirstAirportID() string {
	if x != nil {
		return x.FirstAirportID
	}
	return ""
}

func (x *AirportDistanceReq) GetSecondAirportID() string {
	if x != nil {
		return x.SecondAirportID
	}
	return ""
}

type AirportDistanceRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Distance float64 `protobuf:"fixed64,1,opt,name=Distance,proto3" json:"Distance,omitempty"`
}

func (x *AirportDistanceRes) Reset() {
	*x = AirportDistanceRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_airport_service_proto_air_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AirportDistanceRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AirportDistanceRes) ProtoMessage() {}

func (x *AirportDistanceRes) ProtoReflect() protoreflect.Message {
	mi := &file_airport_service_proto_air_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AirportDistanceRes.ProtoReflect.Descriptor instead.
func (*AirportDistanceRes) Descriptor() ([]byte, []int) {
	return file_airport_service_proto_air_proto_rawDescGZIP(), []int{3}
}

func (x *AirportDistanceRes) GetDistance() float64 {
	if x != nil {
		return x.Distance
	}
	return 0
}

var File_airport_service_proto_air_proto protoreflect.FileDescriptor

var file_airport_service_proto_air_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x61, 0x69, 0x72, 0x70, 0x6f, 0x72, 0x74, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x69, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x03, 0x61, 0x69, 0x72, 0x22, 0x31, 0x0a, 0x11, 0x41, 0x69, 0x72, 0x70, 0x6f, 0x72,
	0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x12, 0x1c, 0x0a, 0x09, 0x41,
	0x69, 0x72, 0x70, 0x6f, 0x72, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x41, 0x69, 0x72, 0x70, 0x6f, 0x72, 0x74, 0x49, 0x44, 0x22, 0xab, 0x02, 0x0a, 0x11, 0x41, 0x69,
	0x72, 0x70, 0x6f, 0x72, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x12,
	0x1c, 0x0a, 0x09, 0x41, 0x69, 0x72, 0x70, 0x6f, 0x72, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x41, 0x69, 0x72, 0x70, 0x6f, 0x72, 0x74, 0x49, 0x44, 0x12, 0x20, 0x0a,
	0x0b, 0x41, 0x69, 0x72, 0x70, 0x6f, 0x72, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x41, 0x69, 0x72, 0x70, 0x6f, 0x72, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x43, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x43,
	0x69, 0x74, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x1a, 0x0a,
	0x08, 0x49, 0x41, 0x54, 0x41, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x49, 0x41, 0x54, 0x41, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x49, 0x43, 0x41,
	0x4f, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x49, 0x43, 0x41,
	0x4f, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x4c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x4c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64,
	0x65, 0x12, 0x1c, 0x0a, 0x09, 0x4c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x4c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x41, 0x6c, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x08, 0x41, 0x6c, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x54,
	0x69, 0x6d, 0x65, 0x5a, 0x6f, 0x6e, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x54,
	0x69, 0x6d, 0x65, 0x5a, 0x6f, 0x6e, 0x65, 0x22, 0x66, 0x0a, 0x12, 0x41, 0x69, 0x72, 0x70, 0x6f,
	0x72, 0x74, 0x44, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x12, 0x26, 0x0a,
	0x0e, 0x46, 0x69, 0x72, 0x73, 0x74, 0x41, 0x69, 0x72, 0x70, 0x6f, 0x72, 0x74, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x46, 0x69, 0x72, 0x73, 0x74, 0x41, 0x69, 0x72, 0x70,
	0x6f, 0x72, 0x74, 0x49, 0x44, 0x12, 0x28, 0x0a, 0x0f, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x41,
	0x69, 0x72, 0x70, 0x6f, 0x72, 0x74, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f,
	0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x41, 0x69, 0x72, 0x70, 0x6f, 0x72, 0x74, 0x49, 0x44, 0x22,
	0x30, 0x0a, 0x12, 0x41, 0x69, 0x72, 0x70, 0x6f, 0x72, 0x74, 0x44, 0x69, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x52, 0x65, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x44, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x44, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x32, 0x93, 0x01, 0x0a, 0x0b, 0x41, 0x69, 0x72, 0x70, 0x6f, 0x72, 0x74, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x43, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x41, 0x69, 0x72, 0x70, 0x6f, 0x72, 0x74, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x16, 0x2e, 0x61, 0x69, 0x72, 0x2e, 0x41, 0x69, 0x72,
	0x70, 0x6f, 0x72, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x16,
	0x2e, 0x61, 0x69, 0x72, 0x2e, 0x41, 0x69, 0x72, 0x70, 0x6f, 0x72, 0x74, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x12, 0x3f, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x44, 0x69, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x17, 0x2e, 0x61, 0x69, 0x72, 0x2e, 0x41, 0x69, 0x72, 0x70,
	0x6f, 0x72, 0x74, 0x44, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x17,
	0x2e, 0x61, 0x69, 0x72, 0x2e, 0x41, 0x69, 0x72, 0x70, 0x6f, 0x72, 0x74, 0x44, 0x69, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x42, 0x3b, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x7a, 0x6f, 0x62, 0x34, 0x35, 0x36, 0x2f, 0x61, 0x69, 0x72,
	0x70, 0x6f, 0x72, 0x74, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x69, 0x72, 0x70, 0x6f, 0x72, 0x74,
	0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_airport_service_proto_air_proto_rawDescOnce sync.Once
	file_airport_service_proto_air_proto_rawDescData = file_airport_service_proto_air_proto_rawDesc
)

func file_airport_service_proto_air_proto_rawDescGZIP() []byte {
	file_airport_service_proto_air_proto_rawDescOnce.Do(func() {
		file_airport_service_proto_air_proto_rawDescData = protoimpl.X.CompressGZIP(file_airport_service_proto_air_proto_rawDescData)
	})
	return file_airport_service_proto_air_proto_rawDescData
}

var file_airport_service_proto_air_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_airport_service_proto_air_proto_goTypes = []interface{}{
	(*AirportDetailsReq)(nil),  // 0: air.AirportDetailsReq
	(*AirportDetailsRes)(nil),  // 1: air.AirportDetailsRes
	(*AirportDistanceReq)(nil), // 2: air.AirportDistanceReq
	(*AirportDistanceRes)(nil), // 3: air.AirportDistanceRes
}
var file_airport_service_proto_air_proto_depIdxs = []int32{
	0, // 0: air.AirportData.GetAirportDetails:input_type -> air.AirportDetailsReq
	2, // 1: air.AirportData.GetDistance:input_type -> air.AirportDistanceReq
	1, // 2: air.AirportData.GetAirportDetails:output_type -> air.AirportDetailsRes
	3, // 3: air.AirportData.GetDistance:output_type -> air.AirportDistanceRes
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_airport_service_proto_air_proto_init() }
func file_airport_service_proto_air_proto_init() {
	if File_airport_service_proto_air_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_airport_service_proto_air_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AirportDetailsReq); i {
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
		file_airport_service_proto_air_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AirportDetailsRes); i {
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
		file_airport_service_proto_air_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AirportDistanceReq); i {
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
		file_airport_service_proto_air_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AirportDistanceRes); i {
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
			RawDescriptor: file_airport_service_proto_air_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_airport_service_proto_air_proto_goTypes,
		DependencyIndexes: file_airport_service_proto_air_proto_depIdxs,
		MessageInfos:      file_airport_service_proto_air_proto_msgTypes,
	}.Build()
	File_airport_service_proto_air_proto = out.File
	file_airport_service_proto_air_proto_rawDesc = nil
	file_airport_service_proto_air_proto_goTypes = nil
	file_airport_service_proto_air_proto_depIdxs = nil
}
