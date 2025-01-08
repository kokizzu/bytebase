// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        (unknown)
// source: store/plan_check_run.proto

package store

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

type PlanCheckRunConfig_ChangeDatabaseType int32

const (
	PlanCheckRunConfig_CHANGE_DATABASE_TYPE_UNSPECIFIED PlanCheckRunConfig_ChangeDatabaseType = 0
	PlanCheckRunConfig_DDL                              PlanCheckRunConfig_ChangeDatabaseType = 1
	PlanCheckRunConfig_DML                              PlanCheckRunConfig_ChangeDatabaseType = 2
	PlanCheckRunConfig_SDL                              PlanCheckRunConfig_ChangeDatabaseType = 3
	PlanCheckRunConfig_DDL_GHOST                        PlanCheckRunConfig_ChangeDatabaseType = 4
	PlanCheckRunConfig_SQL_EDITOR                       PlanCheckRunConfig_ChangeDatabaseType = 5
)

// Enum value maps for PlanCheckRunConfig_ChangeDatabaseType.
var (
	PlanCheckRunConfig_ChangeDatabaseType_name = map[int32]string{
		0: "CHANGE_DATABASE_TYPE_UNSPECIFIED",
		1: "DDL",
		2: "DML",
		3: "SDL",
		4: "DDL_GHOST",
		5: "SQL_EDITOR",
	}
	PlanCheckRunConfig_ChangeDatabaseType_value = map[string]int32{
		"CHANGE_DATABASE_TYPE_UNSPECIFIED": 0,
		"DDL":                              1,
		"DML":                              2,
		"SDL":                              3,
		"DDL_GHOST":                        4,
		"SQL_EDITOR":                       5,
	}
)

func (x PlanCheckRunConfig_ChangeDatabaseType) Enum() *PlanCheckRunConfig_ChangeDatabaseType {
	p := new(PlanCheckRunConfig_ChangeDatabaseType)
	*p = x
	return p
}

func (x PlanCheckRunConfig_ChangeDatabaseType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PlanCheckRunConfig_ChangeDatabaseType) Descriptor() protoreflect.EnumDescriptor {
	return file_store_plan_check_run_proto_enumTypes[0].Descriptor()
}

func (PlanCheckRunConfig_ChangeDatabaseType) Type() protoreflect.EnumType {
	return &file_store_plan_check_run_proto_enumTypes[0]
}

func (x PlanCheckRunConfig_ChangeDatabaseType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PlanCheckRunConfig_ChangeDatabaseType.Descriptor instead.
func (PlanCheckRunConfig_ChangeDatabaseType) EnumDescriptor() ([]byte, []int) {
	return file_store_plan_check_run_proto_rawDescGZIP(), []int{1, 0}
}

type PlanCheckRunResult_Result_Status int32

const (
	PlanCheckRunResult_Result_STATUS_UNSPECIFIED PlanCheckRunResult_Result_Status = 0
	PlanCheckRunResult_Result_ERROR              PlanCheckRunResult_Result_Status = 1
	PlanCheckRunResult_Result_WARNING            PlanCheckRunResult_Result_Status = 2
	PlanCheckRunResult_Result_SUCCESS            PlanCheckRunResult_Result_Status = 3
)

// Enum value maps for PlanCheckRunResult_Result_Status.
var (
	PlanCheckRunResult_Result_Status_name = map[int32]string{
		0: "STATUS_UNSPECIFIED",
		1: "ERROR",
		2: "WARNING",
		3: "SUCCESS",
	}
	PlanCheckRunResult_Result_Status_value = map[string]int32{
		"STATUS_UNSPECIFIED": 0,
		"ERROR":              1,
		"WARNING":            2,
		"SUCCESS":            3,
	}
)

func (x PlanCheckRunResult_Result_Status) Enum() *PlanCheckRunResult_Result_Status {
	p := new(PlanCheckRunResult_Result_Status)
	*p = x
	return p
}

func (x PlanCheckRunResult_Result_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PlanCheckRunResult_Result_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_store_plan_check_run_proto_enumTypes[1].Descriptor()
}

func (PlanCheckRunResult_Result_Status) Type() protoreflect.EnumType {
	return &file_store_plan_check_run_proto_enumTypes[1]
}

func (x PlanCheckRunResult_Result_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PlanCheckRunResult_Result_Status.Descriptor instead.
func (PlanCheckRunResult_Result_Status) EnumDescriptor() ([]byte, []int) {
	return file_store_plan_check_run_proto_rawDescGZIP(), []int{2, 0, 0}
}

type PreUpdateBackupDetail struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The database for keeping the backup data.
	// Format: instances/{instance}/databases/{database}
	Database      string `protobuf:"bytes,1,opt,name=database,proto3" json:"database,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PreUpdateBackupDetail) Reset() {
	*x = PreUpdateBackupDetail{}
	mi := &file_store_plan_check_run_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PreUpdateBackupDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PreUpdateBackupDetail) ProtoMessage() {}

func (x *PreUpdateBackupDetail) ProtoReflect() protoreflect.Message {
	mi := &file_store_plan_check_run_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PreUpdateBackupDetail.ProtoReflect.Descriptor instead.
func (*PreUpdateBackupDetail) Descriptor() ([]byte, []int) {
	return file_store_plan_check_run_proto_rawDescGZIP(), []int{0}
}

func (x *PreUpdateBackupDetail) GetDatabase() string {
	if x != nil {
		return x.Database
	}
	return ""
}

type PlanCheckRunConfig struct {
	state              protoimpl.MessageState                `protogen:"open.v1"`
	SheetUid           int32                                 `protobuf:"varint,1,opt,name=sheet_uid,json=sheetUid,proto3" json:"sheet_uid,omitempty"`
	ChangeDatabaseType PlanCheckRunConfig_ChangeDatabaseType `protobuf:"varint,2,opt,name=change_database_type,json=changeDatabaseType,proto3,enum=bytebase.store.PlanCheckRunConfig_ChangeDatabaseType" json:"change_database_type,omitempty"`
	InstanceUid        int32                                 `protobuf:"varint,3,opt,name=instance_uid,json=instanceUid,proto3" json:"instance_uid,omitempty"`
	DatabaseName       string                                `protobuf:"bytes,4,opt,name=database_name,json=databaseName,proto3" json:"database_name,omitempty"`
	// Deprecated: Marked as deprecated in store/plan_check_run.proto.
	DatabaseGroupUid *int64            `protobuf:"varint,5,opt,name=database_group_uid,json=databaseGroupUid,proto3,oneof" json:"database_group_uid,omitempty"`
	GhostFlags       map[string]string `protobuf:"bytes,6,rep,name=ghost_flags,json=ghostFlags,proto3" json:"ghost_flags,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// If set, a backup of the modified data will be created automatically before any changes are applied.
	PreUpdateBackupDetail *PreUpdateBackupDetail `protobuf:"bytes,7,opt,name=pre_update_backup_detail,json=preUpdateBackupDetail,proto3,oneof" json:"pre_update_backup_detail,omitempty"`
	unknownFields         protoimpl.UnknownFields
	sizeCache             protoimpl.SizeCache
}

func (x *PlanCheckRunConfig) Reset() {
	*x = PlanCheckRunConfig{}
	mi := &file_store_plan_check_run_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PlanCheckRunConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlanCheckRunConfig) ProtoMessage() {}

func (x *PlanCheckRunConfig) ProtoReflect() protoreflect.Message {
	mi := &file_store_plan_check_run_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlanCheckRunConfig.ProtoReflect.Descriptor instead.
func (*PlanCheckRunConfig) Descriptor() ([]byte, []int) {
	return file_store_plan_check_run_proto_rawDescGZIP(), []int{1}
}

func (x *PlanCheckRunConfig) GetSheetUid() int32 {
	if x != nil {
		return x.SheetUid
	}
	return 0
}

func (x *PlanCheckRunConfig) GetChangeDatabaseType() PlanCheckRunConfig_ChangeDatabaseType {
	if x != nil {
		return x.ChangeDatabaseType
	}
	return PlanCheckRunConfig_CHANGE_DATABASE_TYPE_UNSPECIFIED
}

func (x *PlanCheckRunConfig) GetInstanceUid() int32 {
	if x != nil {
		return x.InstanceUid
	}
	return 0
}

func (x *PlanCheckRunConfig) GetDatabaseName() string {
	if x != nil {
		return x.DatabaseName
	}
	return ""
}

// Deprecated: Marked as deprecated in store/plan_check_run.proto.
func (x *PlanCheckRunConfig) GetDatabaseGroupUid() int64 {
	if x != nil && x.DatabaseGroupUid != nil {
		return *x.DatabaseGroupUid
	}
	return 0
}

func (x *PlanCheckRunConfig) GetGhostFlags() map[string]string {
	if x != nil {
		return x.GhostFlags
	}
	return nil
}

func (x *PlanCheckRunConfig) GetPreUpdateBackupDetail() *PreUpdateBackupDetail {
	if x != nil {
		return x.PreUpdateBackupDetail
	}
	return nil
}

type PlanCheckRunResult struct {
	state         protoimpl.MessageState       `protogen:"open.v1"`
	Results       []*PlanCheckRunResult_Result `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	Error         string                       `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PlanCheckRunResult) Reset() {
	*x = PlanCheckRunResult{}
	mi := &file_store_plan_check_run_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PlanCheckRunResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlanCheckRunResult) ProtoMessage() {}

func (x *PlanCheckRunResult) ProtoReflect() protoreflect.Message {
	mi := &file_store_plan_check_run_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlanCheckRunResult.ProtoReflect.Descriptor instead.
func (*PlanCheckRunResult) Descriptor() ([]byte, []int) {
	return file_store_plan_check_run_proto_rawDescGZIP(), []int{2}
}

func (x *PlanCheckRunResult) GetResults() []*PlanCheckRunResult_Result {
	if x != nil {
		return x.Results
	}
	return nil
}

func (x *PlanCheckRunResult) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type PlanCheckRunResult_Result struct {
	state   protoimpl.MessageState           `protogen:"open.v1"`
	Status  PlanCheckRunResult_Result_Status `protobuf:"varint,1,opt,name=status,proto3,enum=bytebase.store.PlanCheckRunResult_Result_Status" json:"status,omitempty"`
	Title   string                           `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Content string                           `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	Code    int32                            `protobuf:"varint,4,opt,name=code,proto3" json:"code,omitempty"`
	// Types that are valid to be assigned to Report:
	//
	//	*PlanCheckRunResult_Result_SqlSummaryReport_
	//	*PlanCheckRunResult_Result_SqlReviewReport_
	Report        isPlanCheckRunResult_Result_Report `protobuf_oneof:"report"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PlanCheckRunResult_Result) Reset() {
	*x = PlanCheckRunResult_Result{}
	mi := &file_store_plan_check_run_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PlanCheckRunResult_Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlanCheckRunResult_Result) ProtoMessage() {}

func (x *PlanCheckRunResult_Result) ProtoReflect() protoreflect.Message {
	mi := &file_store_plan_check_run_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlanCheckRunResult_Result.ProtoReflect.Descriptor instead.
func (*PlanCheckRunResult_Result) Descriptor() ([]byte, []int) {
	return file_store_plan_check_run_proto_rawDescGZIP(), []int{2, 0}
}

func (x *PlanCheckRunResult_Result) GetStatus() PlanCheckRunResult_Result_Status {
	if x != nil {
		return x.Status
	}
	return PlanCheckRunResult_Result_STATUS_UNSPECIFIED
}

func (x *PlanCheckRunResult_Result) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *PlanCheckRunResult_Result) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *PlanCheckRunResult_Result) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *PlanCheckRunResult_Result) GetReport() isPlanCheckRunResult_Result_Report {
	if x != nil {
		return x.Report
	}
	return nil
}

func (x *PlanCheckRunResult_Result) GetSqlSummaryReport() *PlanCheckRunResult_Result_SqlSummaryReport {
	if x != nil {
		if x, ok := x.Report.(*PlanCheckRunResult_Result_SqlSummaryReport_); ok {
			return x.SqlSummaryReport
		}
	}
	return nil
}

func (x *PlanCheckRunResult_Result) GetSqlReviewReport() *PlanCheckRunResult_Result_SqlReviewReport {
	if x != nil {
		if x, ok := x.Report.(*PlanCheckRunResult_Result_SqlReviewReport_); ok {
			return x.SqlReviewReport
		}
	}
	return nil
}

type isPlanCheckRunResult_Result_Report interface {
	isPlanCheckRunResult_Result_Report()
}

type PlanCheckRunResult_Result_SqlSummaryReport_ struct {
	SqlSummaryReport *PlanCheckRunResult_Result_SqlSummaryReport `protobuf:"bytes,5,opt,name=sql_summary_report,json=sqlSummaryReport,proto3,oneof"`
}

type PlanCheckRunResult_Result_SqlReviewReport_ struct {
	SqlReviewReport *PlanCheckRunResult_Result_SqlReviewReport `protobuf:"bytes,6,opt,name=sql_review_report,json=sqlReviewReport,proto3,oneof"`
}

func (*PlanCheckRunResult_Result_SqlSummaryReport_) isPlanCheckRunResult_Result_Report() {}

func (*PlanCheckRunResult_Result_SqlReviewReport_) isPlanCheckRunResult_Result_Report() {}

type PlanCheckRunResult_Result_SqlSummaryReport struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	Code  int32                  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	// statement_types are the types of statements that are found in the sql.
	StatementTypes   []string          `protobuf:"bytes,2,rep,name=statement_types,json=statementTypes,proto3" json:"statement_types,omitempty"`
	AffectedRows     int32             `protobuf:"varint,3,opt,name=affected_rows,json=affectedRows,proto3" json:"affected_rows,omitempty"`
	ChangedResources *ChangedResources `protobuf:"bytes,4,opt,name=changed_resources,json=changedResources,proto3" json:"changed_resources,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *PlanCheckRunResult_Result_SqlSummaryReport) Reset() {
	*x = PlanCheckRunResult_Result_SqlSummaryReport{}
	mi := &file_store_plan_check_run_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PlanCheckRunResult_Result_SqlSummaryReport) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlanCheckRunResult_Result_SqlSummaryReport) ProtoMessage() {}

func (x *PlanCheckRunResult_Result_SqlSummaryReport) ProtoReflect() protoreflect.Message {
	mi := &file_store_plan_check_run_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlanCheckRunResult_Result_SqlSummaryReport.ProtoReflect.Descriptor instead.
func (*PlanCheckRunResult_Result_SqlSummaryReport) Descriptor() ([]byte, []int) {
	return file_store_plan_check_run_proto_rawDescGZIP(), []int{2, 0, 0}
}

func (x *PlanCheckRunResult_Result_SqlSummaryReport) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *PlanCheckRunResult_Result_SqlSummaryReport) GetStatementTypes() []string {
	if x != nil {
		return x.StatementTypes
	}
	return nil
}

func (x *PlanCheckRunResult_Result_SqlSummaryReport) GetAffectedRows() int32 {
	if x != nil {
		return x.AffectedRows
	}
	return 0
}

func (x *PlanCheckRunResult_Result_SqlSummaryReport) GetChangedResources() *ChangedResources {
	if x != nil {
		return x.ChangedResources
	}
	return nil
}

type PlanCheckRunResult_Result_SqlReviewReport struct {
	state  protoimpl.MessageState `protogen:"open.v1"`
	Line   int32                  `protobuf:"varint,1,opt,name=line,proto3" json:"line,omitempty"`
	Column int32                  `protobuf:"varint,2,opt,name=column,proto3" json:"column,omitempty"`
	Detail string                 `protobuf:"bytes,3,opt,name=detail,proto3" json:"detail,omitempty"`
	// Code from sql review.
	Code int32 `protobuf:"varint,4,opt,name=code,proto3" json:"code,omitempty"`
	// 1-based Position of the SQL statement.
	// To supersede `line` and `column` above.
	StartPosition *Position `protobuf:"bytes,8,opt,name=start_position,json=startPosition,proto3" json:"start_position,omitempty"`
	EndPosition   *Position `protobuf:"bytes,9,opt,name=end_position,json=endPosition,proto3" json:"end_position,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PlanCheckRunResult_Result_SqlReviewReport) Reset() {
	*x = PlanCheckRunResult_Result_SqlReviewReport{}
	mi := &file_store_plan_check_run_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PlanCheckRunResult_Result_SqlReviewReport) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlanCheckRunResult_Result_SqlReviewReport) ProtoMessage() {}

func (x *PlanCheckRunResult_Result_SqlReviewReport) ProtoReflect() protoreflect.Message {
	mi := &file_store_plan_check_run_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlanCheckRunResult_Result_SqlReviewReport.ProtoReflect.Descriptor instead.
func (*PlanCheckRunResult_Result_SqlReviewReport) Descriptor() ([]byte, []int) {
	return file_store_plan_check_run_proto_rawDescGZIP(), []int{2, 0, 1}
}

func (x *PlanCheckRunResult_Result_SqlReviewReport) GetLine() int32 {
	if x != nil {
		return x.Line
	}
	return 0
}

func (x *PlanCheckRunResult_Result_SqlReviewReport) GetColumn() int32 {
	if x != nil {
		return x.Column
	}
	return 0
}

func (x *PlanCheckRunResult_Result_SqlReviewReport) GetDetail() string {
	if x != nil {
		return x.Detail
	}
	return ""
}

func (x *PlanCheckRunResult_Result_SqlReviewReport) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *PlanCheckRunResult_Result_SqlReviewReport) GetStartPosition() *Position {
	if x != nil {
		return x.StartPosition
	}
	return nil
}

func (x *PlanCheckRunResult_Result_SqlReviewReport) GetEndPosition() *Position {
	if x != nil {
		return x.EndPosition
	}
	return nil
}

var File_store_plan_check_run_proto protoreflect.FileDescriptor

var file_store_plan_check_run_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x5f, 0x63, 0x68, 0x65,
	0x63, 0x6b, 0x5f, 0x72, 0x75, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x62, 0x79,
	0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x1a, 0x15, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x2f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x33, 0x0a, 0x15, 0x50, 0x72, 0x65, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x12, 0x1a, 0x0a, 0x08, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x22, 0xbc, 0x05, 0x0a,
	0x12, 0x50, 0x6c, 0x61, 0x6e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x75, 0x6e, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x68, 0x65, 0x65, 0x74, 0x5f, 0x75, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x73, 0x68, 0x65, 0x65, 0x74, 0x55, 0x69, 0x64,
	0x12, 0x67, 0x0a, 0x14, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x62,
	0x61, 0x73, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x35,
	0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e,
	0x50, 0x6c, 0x61, 0x6e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x75, 0x6e, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x12, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x44, 0x61, 0x74,
	0x61, 0x62, 0x61, 0x73, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x69, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x75, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0b, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x55, 0x69, 0x64, 0x12, 0x23, 0x0a, 0x0d,
	0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x35, 0x0a, 0x12, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x5f, 0x75, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x42, 0x02, 0x18,
	0x01, 0x48, 0x00, 0x52, 0x10, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x55, 0x69, 0x64, 0x88, 0x01, 0x01, 0x12, 0x53, 0x0a, 0x0b, 0x67, 0x68, 0x6f, 0x73,
	0x74, 0x5f, 0x66, 0x6c, 0x61, 0x67, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x32, 0x2e,
	0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x50,
	0x6c, 0x61, 0x6e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x75, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x47, 0x68, 0x6f, 0x73, 0x74, 0x46, 0x6c, 0x61, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x0a, 0x67, 0x68, 0x6f, 0x73, 0x74, 0x46, 0x6c, 0x61, 0x67, 0x73, 0x12, 0x63, 0x0a,
	0x18, 0x70, 0x72, 0x65, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x62, 0x61, 0x63, 0x6b,
	0x75, 0x70, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x25, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x2e, 0x50, 0x72, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x48, 0x01, 0x52, 0x15, 0x70, 0x72, 0x65, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x88,
	0x01, 0x01, 0x1a, 0x3d, 0x0a, 0x0f, 0x47, 0x68, 0x6f, 0x73, 0x74, 0x46, 0x6c, 0x61, 0x67, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x22, 0x74, 0x0a, 0x12, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x44, 0x61, 0x74, 0x61, 0x62,
	0x61, 0x73, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x24, 0x0a, 0x20, 0x43, 0x48, 0x41, 0x4e, 0x47,
	0x45, 0x5f, 0x44, 0x41, 0x54, 0x41, 0x42, 0x41, 0x53, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x07, 0x0a,
	0x03, 0x44, 0x44, 0x4c, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x44, 0x4d, 0x4c, 0x10, 0x02, 0x12,
	0x07, 0x0a, 0x03, 0x53, 0x44, 0x4c, 0x10, 0x03, 0x12, 0x0d, 0x0a, 0x09, 0x44, 0x44, 0x4c, 0x5f,
	0x47, 0x48, 0x4f, 0x53, 0x54, 0x10, 0x04, 0x12, 0x0e, 0x0a, 0x0a, 0x53, 0x51, 0x4c, 0x5f, 0x45,
	0x44, 0x49, 0x54, 0x4f, 0x52, 0x10, 0x05, 0x42, 0x15, 0x0a, 0x13, 0x5f, 0x64, 0x61, 0x74, 0x61,
	0x62, 0x61, 0x73, 0x65, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x75, 0x69, 0x64, 0x42, 0x1b,
	0x0a, 0x19, 0x5f, 0x70, 0x72, 0x65, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x62, 0x61,
	0x63, 0x6b, 0x75, 0x70, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x22, 0xde, 0x07, 0x0a, 0x12,
	0x50, 0x6c, 0x61, 0x6e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x75, 0x6e, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x12, 0x43, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x75,
	0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x07,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x1a, 0xec, 0x06,
	0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x48, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x30, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62,
	0x61, 0x73, 0x65, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x52, 0x75, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x6a, 0x0a, 0x12, 0x73, 0x71, 0x6c, 0x5f, 0x73, 0x75,
	0x6d, 0x6d, 0x61, 0x72, 0x79, 0x5f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x3a, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x75, 0x6e,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x53, 0x71,
	0x6c, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x48, 0x00,
	0x52, 0x10, 0x73, 0x71, 0x6c, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x52, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x12, 0x67, 0x0a, 0x11, 0x73, 0x71, 0x6c, 0x5f, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77,
	0x5f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x39, 0x2e,
	0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x50,
	0x6c, 0x61, 0x6e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x75, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x53, 0x71, 0x6c, 0x52, 0x65, 0x76, 0x69,
	0x65, 0x77, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x48, 0x00, 0x52, 0x0f, 0x73, 0x71, 0x6c, 0x52,
	0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x1a, 0xc3, 0x01, 0x0a, 0x10,
	0x53, 0x71, 0x6c, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0e, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x73, 0x12, 0x23, 0x0a,
	0x0d, 0x61, 0x66, 0x66, 0x65, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x72, 0x6f, 0x77, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x61, 0x66, 0x66, 0x65, 0x63, 0x74, 0x65, 0x64, 0x52, 0x6f,
	0x77, 0x73, 0x12, 0x4d, 0x0a, 0x11, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x5f, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e,
	0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x43,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x52,
	0x10, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x1a, 0xe7, 0x01, 0x0a, 0x0f, 0x53, 0x71, 0x6c, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x6c,
	0x75, 0x6d, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x63, 0x6f, 0x6c, 0x75, 0x6d,
	0x6e, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x3f, 0x0a,
	0x0e, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65,
	0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x0d, 0x73, 0x74, 0x61, 0x72, 0x74, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3b,
	0x0a, 0x0c, 0x65, 0x6e, 0x64, 0x5f, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b,
	0x65, 0x6e, 0x64, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x45, 0x0a, 0x06, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f,
	0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x09, 0x0a,
	0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x57, 0x41, 0x52, 0x4e,
	0x49, 0x4e, 0x47, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53,
	0x10, 0x03, 0x42, 0x08, 0x0a, 0x06, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x42, 0x14, 0x5a, 0x12,
	0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2d, 0x67, 0x6f, 0x2f, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_store_plan_check_run_proto_rawDescOnce sync.Once
	file_store_plan_check_run_proto_rawDescData = file_store_plan_check_run_proto_rawDesc
)

func file_store_plan_check_run_proto_rawDescGZIP() []byte {
	file_store_plan_check_run_proto_rawDescOnce.Do(func() {
		file_store_plan_check_run_proto_rawDescData = protoimpl.X.CompressGZIP(file_store_plan_check_run_proto_rawDescData)
	})
	return file_store_plan_check_run_proto_rawDescData
}

var file_store_plan_check_run_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_store_plan_check_run_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_store_plan_check_run_proto_goTypes = []any{
	(PlanCheckRunConfig_ChangeDatabaseType)(0),         // 0: bytebase.store.PlanCheckRunConfig.ChangeDatabaseType
	(PlanCheckRunResult_Result_Status)(0),              // 1: bytebase.store.PlanCheckRunResult.Result.Status
	(*PreUpdateBackupDetail)(nil),                      // 2: bytebase.store.PreUpdateBackupDetail
	(*PlanCheckRunConfig)(nil),                         // 3: bytebase.store.PlanCheckRunConfig
	(*PlanCheckRunResult)(nil),                         // 4: bytebase.store.PlanCheckRunResult
	nil,                                                // 5: bytebase.store.PlanCheckRunConfig.GhostFlagsEntry
	(*PlanCheckRunResult_Result)(nil),                  // 6: bytebase.store.PlanCheckRunResult.Result
	(*PlanCheckRunResult_Result_SqlSummaryReport)(nil), // 7: bytebase.store.PlanCheckRunResult.Result.SqlSummaryReport
	(*PlanCheckRunResult_Result_SqlReviewReport)(nil),  // 8: bytebase.store.PlanCheckRunResult.Result.SqlReviewReport
	(*ChangedResources)(nil),                           // 9: bytebase.store.ChangedResources
	(*Position)(nil),                                   // 10: bytebase.store.Position
}
var file_store_plan_check_run_proto_depIdxs = []int32{
	0,  // 0: bytebase.store.PlanCheckRunConfig.change_database_type:type_name -> bytebase.store.PlanCheckRunConfig.ChangeDatabaseType
	5,  // 1: bytebase.store.PlanCheckRunConfig.ghost_flags:type_name -> bytebase.store.PlanCheckRunConfig.GhostFlagsEntry
	2,  // 2: bytebase.store.PlanCheckRunConfig.pre_update_backup_detail:type_name -> bytebase.store.PreUpdateBackupDetail
	6,  // 3: bytebase.store.PlanCheckRunResult.results:type_name -> bytebase.store.PlanCheckRunResult.Result
	1,  // 4: bytebase.store.PlanCheckRunResult.Result.status:type_name -> bytebase.store.PlanCheckRunResult.Result.Status
	7,  // 5: bytebase.store.PlanCheckRunResult.Result.sql_summary_report:type_name -> bytebase.store.PlanCheckRunResult.Result.SqlSummaryReport
	8,  // 6: bytebase.store.PlanCheckRunResult.Result.sql_review_report:type_name -> bytebase.store.PlanCheckRunResult.Result.SqlReviewReport
	9,  // 7: bytebase.store.PlanCheckRunResult.Result.SqlSummaryReport.changed_resources:type_name -> bytebase.store.ChangedResources
	10, // 8: bytebase.store.PlanCheckRunResult.Result.SqlReviewReport.start_position:type_name -> bytebase.store.Position
	10, // 9: bytebase.store.PlanCheckRunResult.Result.SqlReviewReport.end_position:type_name -> bytebase.store.Position
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_store_plan_check_run_proto_init() }
func file_store_plan_check_run_proto_init() {
	if File_store_plan_check_run_proto != nil {
		return
	}
	file_store_changelog_proto_init()
	file_store_common_proto_init()
	file_store_plan_check_run_proto_msgTypes[1].OneofWrappers = []any{}
	file_store_plan_check_run_proto_msgTypes[4].OneofWrappers = []any{
		(*PlanCheckRunResult_Result_SqlSummaryReport_)(nil),
		(*PlanCheckRunResult_Result_SqlReviewReport_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_store_plan_check_run_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_store_plan_check_run_proto_goTypes,
		DependencyIndexes: file_store_plan_check_run_proto_depIdxs,
		EnumInfos:         file_store_plan_check_run_proto_enumTypes,
		MessageInfos:      file_store_plan_check_run_proto_msgTypes,
	}.Build()
	File_store_plan_check_run_proto = out.File
	file_store_plan_check_run_proto_rawDesc = nil
	file_store_plan_check_run_proto_goTypes = nil
	file_store_plan_check_run_proto_depIdxs = nil
}
