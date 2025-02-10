// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: store/sheet.proto

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

type SheetPayload struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The SQL dialect.
	Engine Engine `protobuf:"varint,3,opt,name=engine,proto3,enum=bytebase.store.Engine" json:"engine,omitempty"`
	// The start and end position of each command in the sheet statement.
	Commands      []*SheetCommand `protobuf:"bytes,4,rep,name=commands,proto3" json:"commands,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SheetPayload) Reset() {
	*x = SheetPayload{}
	mi := &file_store_sheet_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SheetPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SheetPayload) ProtoMessage() {}

func (x *SheetPayload) ProtoReflect() protoreflect.Message {
	mi := &file_store_sheet_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SheetPayload.ProtoReflect.Descriptor instead.
func (*SheetPayload) Descriptor() ([]byte, []int) {
	return file_store_sheet_proto_rawDescGZIP(), []int{0}
}

func (x *SheetPayload) GetEngine() Engine {
	if x != nil {
		return x.Engine
	}
	return Engine_ENGINE_UNSPECIFIED
}

func (x *SheetPayload) GetCommands() []*SheetCommand {
	if x != nil {
		return x.Commands
	}
	return nil
}

type SheetCommand struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Start         int32                  `protobuf:"varint,1,opt,name=start,proto3" json:"start,omitempty"`
	End           int32                  `protobuf:"varint,2,opt,name=end,proto3" json:"end,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SheetCommand) Reset() {
	*x = SheetCommand{}
	mi := &file_store_sheet_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SheetCommand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SheetCommand) ProtoMessage() {}

func (x *SheetCommand) ProtoReflect() protoreflect.Message {
	mi := &file_store_sheet_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SheetCommand.ProtoReflect.Descriptor instead.
func (*SheetCommand) Descriptor() ([]byte, []int) {
	return file_store_sheet_proto_rawDescGZIP(), []int{1}
}

func (x *SheetCommand) GetStart() int32 {
	if x != nil {
		return x.Start
	}
	return 0
}

func (x *SheetCommand) GetEnd() int32 {
	if x != nil {
		return x.End
	}
	return 0
}

var File_store_sheet_proto protoreflect.FileDescriptor

var file_store_sheet_proto_rawDesc = string([]byte{
	0x0a, 0x11, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x73, 0x68, 0x65, 0x65, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x1a, 0x12, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x78, 0x0a, 0x0c, 0x53, 0x68, 0x65, 0x65, 0x74,
	0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x2e, 0x0a, 0x06, 0x65, 0x6e, 0x67, 0x69, 0x6e,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61,
	0x73, 0x65, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x52,
	0x06, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x12, 0x38, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x62, 0x79, 0x74, 0x65,
	0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x53, 0x68, 0x65, 0x65, 0x74,
	0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x73, 0x22, 0x36, 0x0a, 0x0c, 0x53, 0x68, 0x65, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x65, 0x6e, 0x64, 0x42, 0x14, 0x5a, 0x12, 0x67, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2d, 0x67, 0x6f, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_store_sheet_proto_rawDescOnce sync.Once
	file_store_sheet_proto_rawDescData []byte
)

func file_store_sheet_proto_rawDescGZIP() []byte {
	file_store_sheet_proto_rawDescOnce.Do(func() {
		file_store_sheet_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_store_sheet_proto_rawDesc), len(file_store_sheet_proto_rawDesc)))
	})
	return file_store_sheet_proto_rawDescData
}

var file_store_sheet_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_store_sheet_proto_goTypes = []any{
	(*SheetPayload)(nil), // 0: bytebase.store.SheetPayload
	(*SheetCommand)(nil), // 1: bytebase.store.SheetCommand
	(Engine)(0),          // 2: bytebase.store.Engine
}
var file_store_sheet_proto_depIdxs = []int32{
	2, // 0: bytebase.store.SheetPayload.engine:type_name -> bytebase.store.Engine
	1, // 1: bytebase.store.SheetPayload.commands:type_name -> bytebase.store.SheetCommand
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_store_sheet_proto_init() }
func file_store_sheet_proto_init() {
	if File_store_sheet_proto != nil {
		return
	}
	file_store_common_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_store_sheet_proto_rawDesc), len(file_store_sheet_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_store_sheet_proto_goTypes,
		DependencyIndexes: file_store_sheet_proto_depIdxs,
		MessageInfos:      file_store_sheet_proto_msgTypes,
	}.Build()
	File_store_sheet_proto = out.File
	file_store_sheet_proto_goTypes = nil
	file_store_sheet_proto_depIdxs = nil
}
