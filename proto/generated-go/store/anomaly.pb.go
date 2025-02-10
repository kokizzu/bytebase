// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: store/anomaly.proto

package store

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AnomalyConnectionPayload struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Connection failure detail
	Detail        string `protobuf:"bytes,1,opt,name=detail,proto3" json:"detail,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AnomalyConnectionPayload) Reset() {
	*x = AnomalyConnectionPayload{}
	mi := &file_store_anomaly_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AnomalyConnectionPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnomalyConnectionPayload) ProtoMessage() {}

func (x *AnomalyConnectionPayload) ProtoReflect() protoreflect.Message {
	mi := &file_store_anomaly_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnomalyConnectionPayload.ProtoReflect.Descriptor instead.
func (*AnomalyConnectionPayload) Descriptor() ([]byte, []int) {
	return file_store_anomaly_proto_rawDescGZIP(), []int{0}
}

func (x *AnomalyConnectionPayload) GetDetail() string {
	if x != nil {
		return x.Detail
	}
	return ""
}

type AnomalyDatabaseSchemaDriftPayload struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The schema version corresponds to the expected schema
	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	// The expected latest schema stored in the migration history table
	Expect string `protobuf:"bytes,2,opt,name=expect,proto3" json:"expect,omitempty"`
	// The actual schema dumped from the database
	Actual        string `protobuf:"bytes,3,opt,name=actual,proto3" json:"actual,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AnomalyDatabaseSchemaDriftPayload) Reset() {
	*x = AnomalyDatabaseSchemaDriftPayload{}
	mi := &file_store_anomaly_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AnomalyDatabaseSchemaDriftPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnomalyDatabaseSchemaDriftPayload) ProtoMessage() {}

func (x *AnomalyDatabaseSchemaDriftPayload) ProtoReflect() protoreflect.Message {
	mi := &file_store_anomaly_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnomalyDatabaseSchemaDriftPayload.ProtoReflect.Descriptor instead.
func (*AnomalyDatabaseSchemaDriftPayload) Descriptor() ([]byte, []int) {
	return file_store_anomaly_proto_rawDescGZIP(), []int{1}
}

func (x *AnomalyDatabaseSchemaDriftPayload) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *AnomalyDatabaseSchemaDriftPayload) GetExpect() string {
	if x != nil {
		return x.Expect
	}
	return ""
}

func (x *AnomalyDatabaseSchemaDriftPayload) GetActual() string {
	if x != nil {
		return x.Actual
	}
	return ""
}

var File_store_anomaly_proto protoreflect.FileDescriptor

var file_store_anomaly_proto_rawDesc = string([]byte{
	0x0a, 0x13, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x61, 0x6e, 0x6f, 0x6d, 0x61, 0x6c, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x22, 0x32, 0x0a, 0x18, 0x41, 0x6e, 0x6f, 0x6d, 0x61, 0x6c, 0x79,
	0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x22, 0x6d, 0x0a, 0x21, 0x41, 0x6e, 0x6f,
	0x6d, 0x61, 0x6c, 0x79, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x53, 0x63, 0x68, 0x65,
	0x6d, 0x61, 0x44, 0x72, 0x69, 0x66, 0x74, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x18,
	0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x78, 0x70, 0x65,
	0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x65, 0x78, 0x70, 0x65, 0x63, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x75, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x61, 0x63, 0x74, 0x75, 0x61, 0x6c, 0x42, 0x14, 0x5a, 0x12, 0x67, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x64, 0x2d, 0x67, 0x6f, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_store_anomaly_proto_rawDescOnce sync.Once
	file_store_anomaly_proto_rawDescData []byte
)

func file_store_anomaly_proto_rawDescGZIP() []byte {
	file_store_anomaly_proto_rawDescOnce.Do(func() {
		file_store_anomaly_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_store_anomaly_proto_rawDesc), len(file_store_anomaly_proto_rawDesc)))
	})
	return file_store_anomaly_proto_rawDescData
}

var file_store_anomaly_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_store_anomaly_proto_goTypes = []any{
	(*AnomalyConnectionPayload)(nil),          // 0: bytebase.store.AnomalyConnectionPayload
	(*AnomalyDatabaseSchemaDriftPayload)(nil), // 1: bytebase.store.AnomalyDatabaseSchemaDriftPayload
}
var file_store_anomaly_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_store_anomaly_proto_init() }
func file_store_anomaly_proto_init() {
	if File_store_anomaly_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_store_anomaly_proto_rawDesc), len(file_store_anomaly_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_store_anomaly_proto_goTypes,
		DependencyIndexes: file_store_anomaly_proto_depIdxs,
		MessageInfos:      file_store_anomaly_proto_msgTypes,
	}.Build()
	File_store_anomaly_proto = out.File
	file_store_anomaly_proto_goTypes = nil
	file_store_anomaly_proto_depIdxs = nil
}
