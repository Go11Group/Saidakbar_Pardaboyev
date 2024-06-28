// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: transport.proto

package transport

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

type Transport struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BusNumber int32    `protobuf:"varint,1,opt,name=bus_number,json=busNumber,proto3" json:"bus_number,omitempty"`
	Stations  []string `protobuf:"bytes,2,rep,name=stations,proto3" json:"stations,omitempty"`
}

func (x *Transport) Reset() {
	*x = Transport{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transport) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transport) ProtoMessage() {}

func (x *Transport) ProtoReflect() protoreflect.Message {
	mi := &file_transport_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transport.ProtoReflect.Descriptor instead.
func (*Transport) Descriptor() ([]byte, []int) {
	return file_transport_proto_rawDescGZIP(), []int{0}
}

func (x *Transport) GetBusNumber() int32 {
	if x != nil {
		return x.BusNumber
	}
	return 0
}

func (x *Transport) GetStations() []string {
	if x != nil {
		return x.Stations
	}
	return nil
}

type RequestBusSchedule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BusNumber int32 `protobuf:"varint,1,opt,name=busNumber,proto3" json:"busNumber,omitempty"`
}

func (x *RequestBusSchedule) Reset() {
	*x = RequestBusSchedule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestBusSchedule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestBusSchedule) ProtoMessage() {}

func (x *RequestBusSchedule) ProtoReflect() protoreflect.Message {
	mi := &file_transport_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestBusSchedule.ProtoReflect.Descriptor instead.
func (*RequestBusSchedule) Descriptor() ([]byte, []int) {
	return file_transport_proto_rawDescGZIP(), []int{1}
}

func (x *RequestBusSchedule) GetBusNumber() int32 {
	if x != nil {
		return x.BusNumber
	}
	return 0
}

type ResponceBusSchedule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Schedules []string `protobuf:"bytes,1,rep,name=schedules,proto3" json:"schedules,omitempty"`
}

func (x *ResponceBusSchedule) Reset() {
	*x = ResponceBusSchedule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponceBusSchedule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponceBusSchedule) ProtoMessage() {}

func (x *ResponceBusSchedule) ProtoReflect() protoreflect.Message {
	mi := &file_transport_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponceBusSchedule.ProtoReflect.Descriptor instead.
func (*ResponceBusSchedule) Descriptor() ([]byte, []int) {
	return file_transport_proto_rawDescGZIP(), []int{2}
}

func (x *ResponceBusSchedule) GetSchedules() []string {
	if x != nil {
		return x.Schedules
	}
	return nil
}

type RequestBusLocation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BusNumber int32 `protobuf:"varint,1,opt,name=busNumber,proto3" json:"busNumber,omitempty"`
}

func (x *RequestBusLocation) Reset() {
	*x = RequestBusLocation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestBusLocation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestBusLocation) ProtoMessage() {}

func (x *RequestBusLocation) ProtoReflect() protoreflect.Message {
	mi := &file_transport_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestBusLocation.ProtoReflect.Descriptor instead.
func (*RequestBusLocation) Descriptor() ([]byte, []int) {
	return file_transport_proto_rawDescGZIP(), []int{3}
}

func (x *RequestBusLocation) GetBusNumber() int32 {
	if x != nil {
		return x.BusNumber
	}
	return 0
}

type ResponceBusLocation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StationName string `protobuf:"bytes,1,opt,name=stationName,proto3" json:"stationName,omitempty"`
}

func (x *ResponceBusLocation) Reset() {
	*x = ResponceBusLocation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponceBusLocation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponceBusLocation) ProtoMessage() {}

func (x *ResponceBusLocation) ProtoReflect() protoreflect.Message {
	mi := &file_transport_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponceBusLocation.ProtoReflect.Descriptor instead.
func (*ResponceBusLocation) Descriptor() ([]byte, []int) {
	return file_transport_proto_rawDescGZIP(), []int{4}
}

func (x *ResponceBusLocation) GetStationName() string {
	if x != nil {
		return x.StationName
	}
	return ""
}

type RequestTrafficJam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Report string `protobuf:"bytes,1,opt,name=report,proto3" json:"report,omitempty"`
}

func (x *RequestTrafficJam) Reset() {
	*x = RequestTrafficJam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestTrafficJam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestTrafficJam) ProtoMessage() {}

func (x *RequestTrafficJam) ProtoReflect() protoreflect.Message {
	mi := &file_transport_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestTrafficJam.ProtoReflect.Descriptor instead.
func (*RequestTrafficJam) Descriptor() ([]byte, []int) {
	return file_transport_proto_rawDescGZIP(), []int{5}
}

func (x *RequestTrafficJam) GetReport() string {
	if x != nil {
		return x.Report
	}
	return ""
}

type ResponceTrafficJam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *ResponceTrafficJam) Reset() {
	*x = ResponceTrafficJam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponceTrafficJam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponceTrafficJam) ProtoMessage() {}

func (x *ResponceTrafficJam) ProtoReflect() protoreflect.Message {
	mi := &file_transport_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponceTrafficJam.ProtoReflect.Descriptor instead.
func (*ResponceTrafficJam) Descriptor() ([]byte, []int) {
	return file_transport_proto_rawDescGZIP(), []int{6}
}

func (x *ResponceTrafficJam) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_transport_proto protoreflect.FileDescriptor

var file_transport_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x46, 0x0a, 0x09, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x1d,
	0x0a, 0x0a, 0x62, 0x75, 0x73, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x09, 0x62, 0x75, 0x73, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1a, 0x0a,
	0x08, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x08, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x32, 0x0a, 0x12, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x42, 0x75, 0x73, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x62, 0x75, 0x73, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x09, 0x62, 0x75, 0x73, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x33, 0x0a,
	0x13, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x63, 0x65, 0x42, 0x75, 0x73, 0x53, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c,
	0x65, 0x73, 0x22, 0x32, 0x0a, 0x12, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x75, 0x73,
	0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x62, 0x75, 0x73, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x62, 0x75, 0x73,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x37, 0x0a, 0x13, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x63, 0x65, 0x42, 0x75, 0x73, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a,
	0x0b, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x22,
	0x2b, 0x0a, 0x11, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69,
	0x63, 0x4a, 0x61, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x22, 0x2e, 0x0a, 0x12,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x63, 0x65, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x4a,
	0x61, 0x6d, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x32, 0xd0, 0x01, 0x0a,
	0x0f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x12, 0x3d, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x42, 0x75, 0x73, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75,
	0x6c, 0x65, 0x12, 0x13, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x75, 0x73, 0x53,
	0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x1a, 0x14, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x63, 0x65, 0x42, 0x75, 0x73, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x22, 0x00, 0x12,
	0x3f, 0x0a, 0x10, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x42, 0x75, 0x73, 0x4c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x13, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x75, 0x73,
	0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x14, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x63, 0x65, 0x42, 0x75, 0x73, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x00,
	0x12, 0x3d, 0x0a, 0x10, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69,
	0x63, 0x4a, 0x61, 0x6d, 0x12, 0x12, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x54, 0x72,
	0x61, 0x66, 0x66, 0x69, 0x63, 0x4a, 0x61, 0x6d, 0x1a, 0x13, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x63, 0x65, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x4a, 0x61, 0x6d, 0x22, 0x00, 0x42,
	0x15, 0x5a, 0x13, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_transport_proto_rawDescOnce sync.Once
	file_transport_proto_rawDescData = file_transport_proto_rawDesc
)

func file_transport_proto_rawDescGZIP() []byte {
	file_transport_proto_rawDescOnce.Do(func() {
		file_transport_proto_rawDescData = protoimpl.X.CompressGZIP(file_transport_proto_rawDescData)
	})
	return file_transport_proto_rawDescData
}

var file_transport_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_transport_proto_goTypes = []interface{}{
	(*Transport)(nil),           // 0: Transport
	(*RequestBusSchedule)(nil),  // 1: RequestBusSchedule
	(*ResponceBusSchedule)(nil), // 2: ResponceBusSchedule
	(*RequestBusLocation)(nil),  // 3: RequestBusLocation
	(*ResponceBusLocation)(nil), // 4: ResponceBusLocation
	(*RequestTrafficJam)(nil),   // 5: RequestTrafficJam
	(*ResponceTrafficJam)(nil),  // 6: ResponceTrafficJam
}
var file_transport_proto_depIdxs = []int32{
	1, // 0: TransportServer.GetBusSchedule:input_type -> RequestBusSchedule
	3, // 1: TransportServer.TrackBusLocation:input_type -> RequestBusLocation
	5, // 2: TransportServer.ReportTrafficJam:input_type -> RequestTrafficJam
	2, // 3: TransportServer.GetBusSchedule:output_type -> ResponceBusSchedule
	4, // 4: TransportServer.TrackBusLocation:output_type -> ResponceBusLocation
	6, // 5: TransportServer.ReportTrafficJam:output_type -> ResponceTrafficJam
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_transport_proto_init() }
func file_transport_proto_init() {
	if File_transport_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_transport_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Transport); i {
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
		file_transport_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestBusSchedule); i {
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
		file_transport_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponceBusSchedule); i {
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
		file_transport_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestBusLocation); i {
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
		file_transport_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponceBusLocation); i {
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
		file_transport_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestTrafficJam); i {
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
		file_transport_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponceTrafficJam); i {
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
			RawDescriptor: file_transport_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_transport_proto_goTypes,
		DependencyIndexes: file_transport_proto_depIdxs,
		MessageInfos:      file_transport_proto_msgTypes,
	}.Build()
	File_transport_proto = out.File
	file_transport_proto_rawDesc = nil
	file_transport_proto_goTypes = nil
	file_transport_proto_depIdxs = nil
}
