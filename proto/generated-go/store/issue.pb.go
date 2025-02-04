// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        (unknown)
// source: store/issue.proto

package store

import (
	expr "google.golang.org/genproto/googleapis/type/expr"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
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

type IssuePayload struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Approval      *IssuePayloadApproval  `protobuf:"bytes,1,opt,name=approval,proto3" json:"approval,omitempty"`
	GrantRequest  *GrantRequest          `protobuf:"bytes,2,opt,name=grant_request,json=grantRequest,proto3" json:"grant_request,omitempty"`
	Labels        []string               `protobuf:"bytes,3,rep,name=labels,proto3" json:"labels,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *IssuePayload) Reset() {
	*x = IssuePayload{}
	mi := &file_store_issue_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IssuePayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IssuePayload) ProtoMessage() {}

func (x *IssuePayload) ProtoReflect() protoreflect.Message {
	mi := &file_store_issue_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IssuePayload.ProtoReflect.Descriptor instead.
func (*IssuePayload) Descriptor() ([]byte, []int) {
	return file_store_issue_proto_rawDescGZIP(), []int{0}
}

func (x *IssuePayload) GetApproval() *IssuePayloadApproval {
	if x != nil {
		return x.Approval
	}
	return nil
}

func (x *IssuePayload) GetGrantRequest() *GrantRequest {
	if x != nil {
		return x.GrantRequest
	}
	return nil
}

func (x *IssuePayload) GetLabels() []string {
	if x != nil {
		return x.Labels
	}
	return nil
}

type GrantRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The requested role.
	// Format: roles/EXPORTER.
	Role string `protobuf:"bytes,1,opt,name=role,proto3" json:"role,omitempty"`
	// The user to be granted.
	// Format: users/{userUID}.
	User          string               `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	Condition     *expr.Expr           `protobuf:"bytes,3,opt,name=condition,proto3" json:"condition,omitempty"`
	Expiration    *durationpb.Duration `protobuf:"bytes,4,opt,name=expiration,proto3" json:"expiration,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GrantRequest) Reset() {
	*x = GrantRequest{}
	mi := &file_store_issue_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GrantRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GrantRequest) ProtoMessage() {}

func (x *GrantRequest) ProtoReflect() protoreflect.Message {
	mi := &file_store_issue_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GrantRequest.ProtoReflect.Descriptor instead.
func (*GrantRequest) Descriptor() ([]byte, []int) {
	return file_store_issue_proto_rawDescGZIP(), []int{1}
}

func (x *GrantRequest) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *GrantRequest) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *GrantRequest) GetCondition() *expr.Expr {
	if x != nil {
		return x.Condition
	}
	return nil
}

func (x *GrantRequest) GetExpiration() *durationpb.Duration {
	if x != nil {
		return x.Expiration
	}
	return nil
}

var File_store_issue_proto protoreflect.FileDescriptor

var file_store_issue_proto_rawDesc = string([]byte{
	0x0a, 0x11, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x69, 0x73, 0x73, 0x75, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x74, 0x79, 0x70, 0x65,
	0x2f, 0x65, 0x78, 0x70, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x2f, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xab, 0x01, 0x0a, 0x0c, 0x49, 0x73, 0x73, 0x75, 0x65, 0x50, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x12, 0x40, 0x0a, 0x08, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x49, 0x73, 0x73, 0x75, 0x65, 0x50, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x52, 0x08, 0x61, 0x70, 0x70, 0x72,
	0x6f, 0x76, 0x61, 0x6c, 0x12, 0x41, 0x0a, 0x0d, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x5f, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x62, 0x79,
	0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x47, 0x72, 0x61,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0c, 0x67, 0x72, 0x61, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x22,
	0xa2, 0x01, 0x0a, 0x0c, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x72, 0x6f, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x2f, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x64,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x52, 0x09,
	0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x39, 0x0a, 0x0a, 0x65, 0x78, 0x70,
	0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x42, 0x14, 0x5a, 0x12, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65,
	0x64, 0x2d, 0x67, 0x6f, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
})

var (
	file_store_issue_proto_rawDescOnce sync.Once
	file_store_issue_proto_rawDescData []byte
)

func file_store_issue_proto_rawDescGZIP() []byte {
	file_store_issue_proto_rawDescOnce.Do(func() {
		file_store_issue_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_store_issue_proto_rawDesc), len(file_store_issue_proto_rawDesc)))
	})
	return file_store_issue_proto_rawDescData
}

var file_store_issue_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_store_issue_proto_goTypes = []any{
	(*IssuePayload)(nil),         // 0: bytebase.store.IssuePayload
	(*GrantRequest)(nil),         // 1: bytebase.store.GrantRequest
	(*IssuePayloadApproval)(nil), // 2: bytebase.store.IssuePayloadApproval
	(*expr.Expr)(nil),            // 3: google.type.Expr
	(*durationpb.Duration)(nil),  // 4: google.protobuf.Duration
}
var file_store_issue_proto_depIdxs = []int32{
	2, // 0: bytebase.store.IssuePayload.approval:type_name -> bytebase.store.IssuePayloadApproval
	1, // 1: bytebase.store.IssuePayload.grant_request:type_name -> bytebase.store.GrantRequest
	3, // 2: bytebase.store.GrantRequest.condition:type_name -> google.type.Expr
	4, // 3: bytebase.store.GrantRequest.expiration:type_name -> google.protobuf.Duration
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_store_issue_proto_init() }
func file_store_issue_proto_init() {
	if File_store_issue_proto != nil {
		return
	}
	file_store_approval_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_store_issue_proto_rawDesc), len(file_store_issue_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_store_issue_proto_goTypes,
		DependencyIndexes: file_store_issue_proto_depIdxs,
		MessageInfos:      file_store_issue_proto_msgTypes,
	}.Build()
	File_store_issue_proto = out.File
	file_store_issue_proto_goTypes = nil
	file_store_issue_proto_depIdxs = nil
}
