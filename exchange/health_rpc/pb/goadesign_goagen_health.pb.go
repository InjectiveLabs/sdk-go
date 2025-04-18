// Code generated with goa v3.7.0, DO NOT EDIT.
//
// health protocol buffer definition
//
// Command:
// $ goa gen github.com/InjectiveLabs/injective-indexer/api/design -o ../

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.24.4
// source: goadesign_goagen_health.proto

package api_v1pb

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

type GetStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetStatusRequest) Reset() {
	*x = GetStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goadesign_goagen_health_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStatusRequest) ProtoMessage() {}

func (x *GetStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_goadesign_goagen_health_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStatusRequest.ProtoReflect.Descriptor instead.
func (*GetStatusRequest) Descriptor() ([]byte, []int) {
	return file_goadesign_goagen_health_proto_rawDescGZIP(), []int{0}
}

type GetStatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Status of the response.
	S string `protobuf:"bytes,1,opt,name=s,proto3" json:"s,omitempty"`
	// Error message.
	Errmsg string        `protobuf:"bytes,2,opt,name=errmsg,proto3" json:"errmsg,omitempty"`
	Data   *HealthStatus `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	Status string        `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *GetStatusResponse) Reset() {
	*x = GetStatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goadesign_goagen_health_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStatusResponse) ProtoMessage() {}

func (x *GetStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_goadesign_goagen_health_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStatusResponse.ProtoReflect.Descriptor instead.
func (*GetStatusResponse) Descriptor() ([]byte, []int) {
	return file_goadesign_goagen_health_proto_rawDescGZIP(), []int{1}
}

func (x *GetStatusResponse) GetS() string {
	if x != nil {
		return x.S
	}
	return ""
}

func (x *GetStatusResponse) GetErrmsg() string {
	if x != nil {
		return x.Errmsg
	}
	return ""
}

func (x *GetStatusResponse) GetData() *HealthStatus {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *GetStatusResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

// Status defines the structure for health information
type HealthStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Block height from local mongo exchange db.
	LocalHeight int32 `protobuf:"zigzag32,1,opt,name=local_height,json=localHeight,proto3" json:"local_height,omitempty"`
	// block timestamp from local mongo exchange db.
	LocalTimestamp int32 `protobuf:"zigzag32,2,opt,name=local_timestamp,json=localTimestamp,proto3" json:"local_timestamp,omitempty"`
	// block height from Horacle service.
	HoracleHeight int32 `protobuf:"zigzag32,3,opt,name=horacle_height,json=horacleHeight,proto3" json:"horacle_height,omitempty"`
	// block timestamp from Horacle service.
	HoracleTimestamp int32 `protobuf:"zigzag32,4,opt,name=horacle_timestamp,json=horacleTimestamp,proto3" json:"horacle_timestamp,omitempty"`
	// Migration version of the database.
	MigrationLastVersion int32 `protobuf:"zigzag32,5,opt,name=migration_last_version,json=migrationLastVersion,proto3" json:"migration_last_version,omitempty"`
	// Block height from event provider service.
	EpHeight int32 `protobuf:"zigzag32,6,opt,name=ep_height,json=epHeight,proto3" json:"ep_height,omitempty"`
	// Block UNIX timestamp from event provider service.
	EpTimestamp int32 `protobuf:"zigzag32,7,opt,name=ep_timestamp,json=epTimestamp,proto3" json:"ep_timestamp,omitempty"`
}

func (x *HealthStatus) Reset() {
	*x = HealthStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goadesign_goagen_health_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthStatus) ProtoMessage() {}

func (x *HealthStatus) ProtoReflect() protoreflect.Message {
	mi := &file_goadesign_goagen_health_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthStatus.ProtoReflect.Descriptor instead.
func (*HealthStatus) Descriptor() ([]byte, []int) {
	return file_goadesign_goagen_health_proto_rawDescGZIP(), []int{2}
}

func (x *HealthStatus) GetLocalHeight() int32 {
	if x != nil {
		return x.LocalHeight
	}
	return 0
}

func (x *HealthStatus) GetLocalTimestamp() int32 {
	if x != nil {
		return x.LocalTimestamp
	}
	return 0
}

func (x *HealthStatus) GetHoracleHeight() int32 {
	if x != nil {
		return x.HoracleHeight
	}
	return 0
}

func (x *HealthStatus) GetHoracleTimestamp() int32 {
	if x != nil {
		return x.HoracleTimestamp
	}
	return 0
}

func (x *HealthStatus) GetMigrationLastVersion() int32 {
	if x != nil {
		return x.MigrationLastVersion
	}
	return 0
}

func (x *HealthStatus) GetEpHeight() int32 {
	if x != nil {
		return x.EpHeight
	}
	return 0
}

func (x *HealthStatus) GetEpTimestamp() int32 {
	if x != nil {
		return x.EpTimestamp
	}
	return 0
}

var File_goadesign_goagen_health_proto protoreflect.FileDescriptor

var file_goadesign_goagen_health_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x67, 0x6f, 0x61, 0x64, 0x65, 0x73, 0x69, 0x67, 0x6e, 0x5f, 0x67, 0x6f, 0x61, 0x67,
	0x65, 0x6e, 0x5f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x06, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x22, 0x12, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x7b, 0x0a, 0x11, 0x47,
	0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x0c, 0x0a, 0x01, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01, 0x73, 0x12, 0x16,
	0x0a, 0x06, 0x65, 0x72, 0x72, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x65, 0x72, 0x72, 0x6d, 0x73, 0x67, 0x12, 0x28, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x65,
	0x61, 0x6c, 0x74, 0x68, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0xa4, 0x02, 0x0a, 0x0c, 0x48, 0x65, 0x61,
	0x6c, 0x74, 0x68, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x6c, 0x6f, 0x63,
	0x61, 0x6c, 0x5f, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x11, 0x52,
	0x0b, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x27, 0x0a, 0x0f,
	0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x11, 0x52, 0x0e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x25, 0x0a, 0x0e, 0x68, 0x6f, 0x72, 0x61, 0x63, 0x6c, 0x65,
	0x5f, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x11, 0x52, 0x0d, 0x68,
	0x6f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x2b, 0x0a, 0x11,
	0x68, 0x6f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x11, 0x52, 0x10, 0x68, 0x6f, 0x72, 0x61, 0x63, 0x6c, 0x65,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x34, 0x0a, 0x16, 0x6d, 0x69, 0x67,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x11, 0x52, 0x14, 0x6d, 0x69, 0x67, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x61, 0x73, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x1b, 0x0a, 0x09, 0x65, 0x70, 0x5f, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x11, 0x52, 0x08, 0x65, 0x70, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x21, 0x0a, 0x0c,
	0x65, 0x70, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x11, 0x52, 0x0b, 0x65, 0x70, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x32,
	0x4a, 0x0a, 0x06, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x12, 0x40, 0x0a, 0x09, 0x47, 0x65, 0x74,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0b, 0x5a, 0x09, 0x2f,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_goadesign_goagen_health_proto_rawDescOnce sync.Once
	file_goadesign_goagen_health_proto_rawDescData = file_goadesign_goagen_health_proto_rawDesc
)

func file_goadesign_goagen_health_proto_rawDescGZIP() []byte {
	file_goadesign_goagen_health_proto_rawDescOnce.Do(func() {
		file_goadesign_goagen_health_proto_rawDescData = protoimpl.X.CompressGZIP(file_goadesign_goagen_health_proto_rawDescData)
	})
	return file_goadesign_goagen_health_proto_rawDescData
}

var file_goadesign_goagen_health_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_goadesign_goagen_health_proto_goTypes = []interface{}{
	(*GetStatusRequest)(nil),  // 0: api.v1.GetStatusRequest
	(*GetStatusResponse)(nil), // 1: api.v1.GetStatusResponse
	(*HealthStatus)(nil),      // 2: api.v1.HealthStatus
}
var file_goadesign_goagen_health_proto_depIdxs = []int32{
	2, // 0: api.v1.GetStatusResponse.data:type_name -> api.v1.HealthStatus
	0, // 1: api.v1.Health.GetStatus:input_type -> api.v1.GetStatusRequest
	1, // 2: api.v1.Health.GetStatus:output_type -> api.v1.GetStatusResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_goadesign_goagen_health_proto_init() }
func file_goadesign_goagen_health_proto_init() {
	if File_goadesign_goagen_health_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_goadesign_goagen_health_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStatusRequest); i {
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
		file_goadesign_goagen_health_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStatusResponse); i {
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
		file_goadesign_goagen_health_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealthStatus); i {
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
			RawDescriptor: file_goadesign_goagen_health_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_goadesign_goagen_health_proto_goTypes,
		DependencyIndexes: file_goadesign_goagen_health_proto_depIdxs,
		MessageInfos:      file_goadesign_goagen_health_proto_msgTypes,
	}.Build()
	File_goadesign_goagen_health_proto = out.File
	file_goadesign_goagen_health_proto_rawDesc = nil
	file_goadesign_goagen_health_proto_goTypes = nil
	file_goadesign_goagen_health_proto_depIdxs = nil
}
