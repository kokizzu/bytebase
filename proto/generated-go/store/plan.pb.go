// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: store/plan.proto

package store

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

// Type is the database change type.
type PlanConfig_ChangeDatabaseConfig_Type int32

const (
	PlanConfig_ChangeDatabaseConfig_TYPE_UNSPECIFIED PlanConfig_ChangeDatabaseConfig_Type = 0
	// Used for establishing schema baseline, this is used when
	// 1. Onboard the database into Bytebase since Bytebase needs to know the current database schema.
	// 2. Had schema drift and need to re-establish the baseline.
	PlanConfig_ChangeDatabaseConfig_BASELINE PlanConfig_ChangeDatabaseConfig_Type = 1
	// Used for DDL changes including CREATE DATABASE.
	PlanConfig_ChangeDatabaseConfig_MIGRATE PlanConfig_ChangeDatabaseConfig_Type = 2
	// Used for schema changes via state-based schema migration including CREATE DATABASE.
	PlanConfig_ChangeDatabaseConfig_MIGRATE_SDL PlanConfig_ChangeDatabaseConfig_Type = 3
	// Used for DDL changes using gh-ost.
	PlanConfig_ChangeDatabaseConfig_MIGRATE_GHOST PlanConfig_ChangeDatabaseConfig_Type = 4
	// Used when restoring from a backup (the restored database branched from the original backup).
	PlanConfig_ChangeDatabaseConfig_BRANCH PlanConfig_ChangeDatabaseConfig_Type = 5
	// Used for DML change.
	PlanConfig_ChangeDatabaseConfig_DATA PlanConfig_ChangeDatabaseConfig_Type = 6
)

// Enum value maps for PlanConfig_ChangeDatabaseConfig_Type.
var (
	PlanConfig_ChangeDatabaseConfig_Type_name = map[int32]string{
		0: "TYPE_UNSPECIFIED",
		1: "BASELINE",
		2: "MIGRATE",
		3: "MIGRATE_SDL",
		4: "MIGRATE_GHOST",
		5: "BRANCH",
		6: "DATA",
	}
	PlanConfig_ChangeDatabaseConfig_Type_value = map[string]int32{
		"TYPE_UNSPECIFIED": 0,
		"BASELINE":         1,
		"MIGRATE":          2,
		"MIGRATE_SDL":      3,
		"MIGRATE_GHOST":    4,
		"BRANCH":           5,
		"DATA":             6,
	}
)

func (x PlanConfig_ChangeDatabaseConfig_Type) Enum() *PlanConfig_ChangeDatabaseConfig_Type {
	p := new(PlanConfig_ChangeDatabaseConfig_Type)
	*p = x
	return p
}

func (x PlanConfig_ChangeDatabaseConfig_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PlanConfig_ChangeDatabaseConfig_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_store_plan_proto_enumTypes[0].Descriptor()
}

func (PlanConfig_ChangeDatabaseConfig_Type) Type() protoreflect.EnumType {
	return &file_store_plan_proto_enumTypes[0]
}

func (x PlanConfig_ChangeDatabaseConfig_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PlanConfig_ChangeDatabaseConfig_Type.Descriptor instead.
func (PlanConfig_ChangeDatabaseConfig_Type) EnumDescriptor() ([]byte, []int) {
	return file_store_plan_proto_rawDescGZIP(), []int{0, 3, 0}
}

type PlanConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Steps []*PlanConfig_Step `protobuf:"bytes,1,rep,name=steps,proto3" json:"steps,omitempty"`
}

func (x *PlanConfig) Reset() {
	*x = PlanConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_plan_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlanConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlanConfig) ProtoMessage() {}

func (x *PlanConfig) ProtoReflect() protoreflect.Message {
	mi := &file_store_plan_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlanConfig.ProtoReflect.Descriptor instead.
func (*PlanConfig) Descriptor() ([]byte, []int) {
	return file_store_plan_proto_rawDescGZIP(), []int{0}
}

func (x *PlanConfig) GetSteps() []*PlanConfig_Step {
	if x != nil {
		return x.Steps
	}
	return nil
}

type PlanConfig_Step struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Specs []*PlanConfig_Spec `protobuf:"bytes,1,rep,name=specs,proto3" json:"specs,omitempty"`
}

func (x *PlanConfig_Step) Reset() {
	*x = PlanConfig_Step{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_plan_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlanConfig_Step) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlanConfig_Step) ProtoMessage() {}

func (x *PlanConfig_Step) ProtoReflect() protoreflect.Message {
	mi := &file_store_plan_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlanConfig_Step.ProtoReflect.Descriptor instead.
func (*PlanConfig_Step) Descriptor() ([]byte, []int) {
	return file_store_plan_proto_rawDescGZIP(), []int{0, 0}
}

func (x *PlanConfig_Step) GetSpecs() []*PlanConfig_Spec {
	if x != nil {
		return x.Specs
	}
	return nil
}

type PlanConfig_Spec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// earliest_allowed_time the earliest execution time of the change.
	EarliestAllowedTime *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=earliest_allowed_time,json=earliestAllowedTime,proto3" json:"earliest_allowed_time,omitempty"`
	// Types that are assignable to Config:
	//
	//	*PlanConfig_Spec_CreateDatabaseConfig
	//	*PlanConfig_Spec_ChangeDatabaseConfig
	//	*PlanConfig_Spec_RestoreDatabaseConfig
	Config isPlanConfig_Spec_Config `protobuf_oneof:"config"`
}

func (x *PlanConfig_Spec) Reset() {
	*x = PlanConfig_Spec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_plan_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlanConfig_Spec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlanConfig_Spec) ProtoMessage() {}

func (x *PlanConfig_Spec) ProtoReflect() protoreflect.Message {
	mi := &file_store_plan_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlanConfig_Spec.ProtoReflect.Descriptor instead.
func (*PlanConfig_Spec) Descriptor() ([]byte, []int) {
	return file_store_plan_proto_rawDescGZIP(), []int{0, 1}
}

func (x *PlanConfig_Spec) GetEarliestAllowedTime() *timestamppb.Timestamp {
	if x != nil {
		return x.EarliestAllowedTime
	}
	return nil
}

func (m *PlanConfig_Spec) GetConfig() isPlanConfig_Spec_Config {
	if m != nil {
		return m.Config
	}
	return nil
}

func (x *PlanConfig_Spec) GetCreateDatabaseConfig() *PlanConfig_CreateDatabaseConfig {
	if x, ok := x.GetConfig().(*PlanConfig_Spec_CreateDatabaseConfig); ok {
		return x.CreateDatabaseConfig
	}
	return nil
}

func (x *PlanConfig_Spec) GetChangeDatabaseConfig() *PlanConfig_ChangeDatabaseConfig {
	if x, ok := x.GetConfig().(*PlanConfig_Spec_ChangeDatabaseConfig); ok {
		return x.ChangeDatabaseConfig
	}
	return nil
}

func (x *PlanConfig_Spec) GetRestoreDatabaseConfig() *PlanConfig_RestoreDatabaseConfig {
	if x, ok := x.GetConfig().(*PlanConfig_Spec_RestoreDatabaseConfig); ok {
		return x.RestoreDatabaseConfig
	}
	return nil
}

type isPlanConfig_Spec_Config interface {
	isPlanConfig_Spec_Config()
}

type PlanConfig_Spec_CreateDatabaseConfig struct {
	CreateDatabaseConfig *PlanConfig_CreateDatabaseConfig `protobuf:"bytes,1,opt,name=create_database_config,json=createDatabaseConfig,proto3,oneof"`
}

type PlanConfig_Spec_ChangeDatabaseConfig struct {
	ChangeDatabaseConfig *PlanConfig_ChangeDatabaseConfig `protobuf:"bytes,2,opt,name=change_database_config,json=changeDatabaseConfig,proto3,oneof"`
}

type PlanConfig_Spec_RestoreDatabaseConfig struct {
	RestoreDatabaseConfig *PlanConfig_RestoreDatabaseConfig `protobuf:"bytes,3,opt,name=restore_database_config,json=restoreDatabaseConfig,proto3,oneof"`
}

func (*PlanConfig_Spec_CreateDatabaseConfig) isPlanConfig_Spec_Config() {}

func (*PlanConfig_Spec_ChangeDatabaseConfig) isPlanConfig_Spec_Config() {}

func (*PlanConfig_Spec_RestoreDatabaseConfig) isPlanConfig_Spec_Config() {}

type PlanConfig_CreateDatabaseConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The resource name of the instance on which the database is created.
	// Format: instances/{instance}
	Target string `protobuf:"bytes,1,opt,name=target,proto3" json:"target,omitempty"`
	// The name of the database to create.
	Database string `protobuf:"bytes,2,opt,name=database,proto3" json:"database,omitempty"`
	// table is the name of the table, if it is not empty, Bytebase should create a table after creating the database.
	// For example, in MongoDB, it only creates the database when we first store data in that database.
	Table string `protobuf:"bytes,3,opt,name=table,proto3" json:"table,omitempty"`
	// character_set is the character set of the database.
	CharacterSet string `protobuf:"bytes,4,opt,name=character_set,json=characterSet,proto3" json:"character_set,omitempty"`
	// collation is the collation of the database.
	Collation string `protobuf:"bytes,5,opt,name=collation,proto3" json:"collation,omitempty"`
	// cluster is the cluster of the database. This is only applicable to ClickHouse for "ON CLUSTER <<cluster>>".
	Cluster string `protobuf:"bytes,6,opt,name=cluster,proto3" json:"cluster,omitempty"`
	// owner is the owner of the database. This is only applicable to Postgres for "WITH OWNER <<owner>>".
	Owner string `protobuf:"bytes,7,opt,name=owner,proto3" json:"owner,omitempty"`
	// backup is the resource name of the backup.
	// FIXME: backup v1 API is not ready yet, write the format here when it's ready.
	Backup string `protobuf:"bytes,8,opt,name=backup,proto3" json:"backup,omitempty"`
	// labels of the database.
	Labels map[string]string `protobuf:"bytes,9,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *PlanConfig_CreateDatabaseConfig) Reset() {
	*x = PlanConfig_CreateDatabaseConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_plan_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlanConfig_CreateDatabaseConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlanConfig_CreateDatabaseConfig) ProtoMessage() {}

func (x *PlanConfig_CreateDatabaseConfig) ProtoReflect() protoreflect.Message {
	mi := &file_store_plan_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlanConfig_CreateDatabaseConfig.ProtoReflect.Descriptor instead.
func (*PlanConfig_CreateDatabaseConfig) Descriptor() ([]byte, []int) {
	return file_store_plan_proto_rawDescGZIP(), []int{0, 2}
}

func (x *PlanConfig_CreateDatabaseConfig) GetTarget() string {
	if x != nil {
		return x.Target
	}
	return ""
}

func (x *PlanConfig_CreateDatabaseConfig) GetDatabase() string {
	if x != nil {
		return x.Database
	}
	return ""
}

func (x *PlanConfig_CreateDatabaseConfig) GetTable() string {
	if x != nil {
		return x.Table
	}
	return ""
}

func (x *PlanConfig_CreateDatabaseConfig) GetCharacterSet() string {
	if x != nil {
		return x.CharacterSet
	}
	return ""
}

func (x *PlanConfig_CreateDatabaseConfig) GetCollation() string {
	if x != nil {
		return x.Collation
	}
	return ""
}

func (x *PlanConfig_CreateDatabaseConfig) GetCluster() string {
	if x != nil {
		return x.Cluster
	}
	return ""
}

func (x *PlanConfig_CreateDatabaseConfig) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *PlanConfig_CreateDatabaseConfig) GetBackup() string {
	if x != nil {
		return x.Backup
	}
	return ""
}

func (x *PlanConfig_CreateDatabaseConfig) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

type PlanConfig_ChangeDatabaseConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The resource name of the target.
	// Format: projects/{project}/logicalDatabases/{ldb1}.
	// Format: projects/{project}/logicalDatabases/{ldb1}/logicalTables/{ltb1}.
	// Format: instances/{xxx}/databases/{db1}.
	Target string `protobuf:"bytes,1,opt,name=target,proto3" json:"target,omitempty"`
	// The resource name of the sheet.
	// Format: sheets/{sheet}
	Sheet string                               `protobuf:"bytes,2,opt,name=sheet,proto3" json:"sheet,omitempty"`
	Type  PlanConfig_ChangeDatabaseConfig_Type `protobuf:"varint,3,opt,name=type,proto3,enum=bytebase.store.PlanConfig_ChangeDatabaseConfig_Type" json:"type,omitempty"`
	// schema_version is parsed from VCS file name.
	// It is automatically generated in the UI workflow.
	SchemaVersion string `protobuf:"bytes,4,opt,name=schema_version,json=schemaVersion,proto3" json:"schema_version,omitempty"`
	// If RollbackEnabled, build the RollbackSheetID of the task.
	RollbackEnabled bool `protobuf:"varint,5,opt,name=rollback_enabled,json=rollbackEnabled,proto3" json:"rollback_enabled,omitempty"`
}

func (x *PlanConfig_ChangeDatabaseConfig) Reset() {
	*x = PlanConfig_ChangeDatabaseConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_plan_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlanConfig_ChangeDatabaseConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlanConfig_ChangeDatabaseConfig) ProtoMessage() {}

func (x *PlanConfig_ChangeDatabaseConfig) ProtoReflect() protoreflect.Message {
	mi := &file_store_plan_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlanConfig_ChangeDatabaseConfig.ProtoReflect.Descriptor instead.
func (*PlanConfig_ChangeDatabaseConfig) Descriptor() ([]byte, []int) {
	return file_store_plan_proto_rawDescGZIP(), []int{0, 3}
}

func (x *PlanConfig_ChangeDatabaseConfig) GetTarget() string {
	if x != nil {
		return x.Target
	}
	return ""
}

func (x *PlanConfig_ChangeDatabaseConfig) GetSheet() string {
	if x != nil {
		return x.Sheet
	}
	return ""
}

func (x *PlanConfig_ChangeDatabaseConfig) GetType() PlanConfig_ChangeDatabaseConfig_Type {
	if x != nil {
		return x.Type
	}
	return PlanConfig_ChangeDatabaseConfig_TYPE_UNSPECIFIED
}

func (x *PlanConfig_ChangeDatabaseConfig) GetSchemaVersion() string {
	if x != nil {
		return x.SchemaVersion
	}
	return ""
}

func (x *PlanConfig_ChangeDatabaseConfig) GetRollbackEnabled() bool {
	if x != nil {
		return x.RollbackEnabled
	}
	return false
}

type PlanConfig_RestoreDatabaseConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The resource name of the target to restore.
	// Format: instances/{instance}/databases/{database}
	Target string `protobuf:"bytes,1,opt,name=target,proto3" json:"target,omitempty"`
	// create_database_config is present if the user wants to restore to a new database.
	CreateDatabaseConfig *PlanConfig_CreateDatabaseConfig `protobuf:"bytes,2,opt,name=create_database_config,json=createDatabaseConfig,proto3,oneof" json:"create_database_config,omitempty"`
	// source determines how to restore the database.
	// 1. from a backup
	// 2. from a point in time
	//
	// Types that are assignable to Source:
	//
	//	*PlanConfig_RestoreDatabaseConfig_Backup
	//	*PlanConfig_RestoreDatabaseConfig_PointInTime
	Source isPlanConfig_RestoreDatabaseConfig_Source `protobuf_oneof:"source"`
}

func (x *PlanConfig_RestoreDatabaseConfig) Reset() {
	*x = PlanConfig_RestoreDatabaseConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_plan_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlanConfig_RestoreDatabaseConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlanConfig_RestoreDatabaseConfig) ProtoMessage() {}

func (x *PlanConfig_RestoreDatabaseConfig) ProtoReflect() protoreflect.Message {
	mi := &file_store_plan_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlanConfig_RestoreDatabaseConfig.ProtoReflect.Descriptor instead.
func (*PlanConfig_RestoreDatabaseConfig) Descriptor() ([]byte, []int) {
	return file_store_plan_proto_rawDescGZIP(), []int{0, 4}
}

func (x *PlanConfig_RestoreDatabaseConfig) GetTarget() string {
	if x != nil {
		return x.Target
	}
	return ""
}

func (x *PlanConfig_RestoreDatabaseConfig) GetCreateDatabaseConfig() *PlanConfig_CreateDatabaseConfig {
	if x != nil {
		return x.CreateDatabaseConfig
	}
	return nil
}

func (m *PlanConfig_RestoreDatabaseConfig) GetSource() isPlanConfig_RestoreDatabaseConfig_Source {
	if m != nil {
		return m.Source
	}
	return nil
}

func (x *PlanConfig_RestoreDatabaseConfig) GetBackup() string {
	if x, ok := x.GetSource().(*PlanConfig_RestoreDatabaseConfig_Backup); ok {
		return x.Backup
	}
	return ""
}

func (x *PlanConfig_RestoreDatabaseConfig) GetPointInTime() *timestamppb.Timestamp {
	if x, ok := x.GetSource().(*PlanConfig_RestoreDatabaseConfig_PointInTime); ok {
		return x.PointInTime
	}
	return nil
}

type isPlanConfig_RestoreDatabaseConfig_Source interface {
	isPlanConfig_RestoreDatabaseConfig_Source()
}

type PlanConfig_RestoreDatabaseConfig_Backup struct {
	// FIXME: format TBD
	// Restore from a backup.
	Backup string `protobuf:"bytes,3,opt,name=backup,proto3,oneof"`
}

type PlanConfig_RestoreDatabaseConfig_PointInTime struct {
	// After the PITR operations, the database will be recovered to the state at this time.
	PointInTime *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=point_in_time,json=pointInTime,proto3,oneof"`
}

func (*PlanConfig_RestoreDatabaseConfig_Backup) isPlanConfig_RestoreDatabaseConfig_Source() {}

func (*PlanConfig_RestoreDatabaseConfig_PointInTime) isPlanConfig_RestoreDatabaseConfig_Source() {}

var File_store_plan_proto protoreflect.FileDescriptor

var file_store_plan_proto_rawDesc = []byte{
	0x0a, 0x10, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x96, 0x0c, 0x0a, 0x0a, 0x50, 0x6c, 0x61, 0x6e, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x12, 0x35, 0x0a, 0x05, 0x73, 0x74, 0x65, 0x70, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1f, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x53, 0x74,
	0x65, 0x70, 0x52, 0x05, 0x73, 0x74, 0x65, 0x70, 0x73, 0x1a, 0x3d, 0x0a, 0x04, 0x53, 0x74, 0x65,
	0x70, 0x12, 0x35, 0x0a, 0x05, 0x73, 0x70, 0x65, 0x63, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1f, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x53, 0x70, 0x65,
	0x63, 0x52, 0x05, 0x73, 0x70, 0x65, 0x63, 0x73, 0x1a, 0x9e, 0x03, 0x0a, 0x04, 0x53, 0x70, 0x65,
	0x63, 0x12, 0x4e, 0x0a, 0x15, 0x65, 0x61, 0x72, 0x6c, 0x69, 0x65, 0x73, 0x74, 0x5f, 0x61, 0x6c,
	0x6c, 0x6f, 0x77, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x13, 0x65, 0x61,
	0x72, 0x6c, 0x69, 0x65, 0x73, 0x74, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x67, 0x0a, 0x16, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x61,
	0x62, 0x61, 0x73, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x2f, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x48, 0x00, 0x52, 0x14, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61,
	0x62, 0x61, 0x73, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x67, 0x0a, 0x16, 0x63, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x62, 0x79, 0x74,
	0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x50, 0x6c, 0x61, 0x6e,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x44, 0x61, 0x74,
	0x61, 0x62, 0x61, 0x73, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x48, 0x00, 0x52, 0x14, 0x63,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x12, 0x6a, 0x0a, 0x17, 0x72, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x64,
	0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x52, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x48, 0x00, 0x52, 0x15, 0x72, 0x65, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42,
	0x08, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0xfb, 0x02, 0x0a, 0x14, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x61,
	0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x61,
	0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x23, 0x0a, 0x0d,
	0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x53, 0x65,
	0x74, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e,
	0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12,
	0x16, 0x0a, 0x06, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x12, 0x53, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c,
	0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3b, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61,
	0x73, 0x65, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61,
	0x73, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x1a, 0x39, 0x0a, 0x0b,
	0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0xd3, 0x02, 0x0a, 0x14, 0x43, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x12, 0x16, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x68, 0x65, 0x65,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x68, 0x65, 0x65, 0x74, 0x12, 0x48,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x34, 0x2e, 0x62,
	0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x50, 0x6c,
	0x61, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x44,
	0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x63, 0x68, 0x65,
	0x6d, 0x61, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x29, 0x0a, 0x10, 0x72, 0x6f, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x5f, 0x65, 0x6e, 0x61, 0x62,
	0x6c, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x72, 0x6f, 0x6c, 0x6c, 0x62,
	0x61, 0x63, 0x6b, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x22, 0x71, 0x0a, 0x04, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x14, 0x0a, 0x10, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x42, 0x41, 0x53, 0x45,
	0x4c, 0x49, 0x4e, 0x45, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x4d, 0x49, 0x47, 0x52, 0x41, 0x54,
	0x45, 0x10, 0x02, 0x12, 0x0f, 0x0a, 0x0b, 0x4d, 0x49, 0x47, 0x52, 0x41, 0x54, 0x45, 0x5f, 0x53,
	0x44, 0x4c, 0x10, 0x03, 0x12, 0x11, 0x0a, 0x0d, 0x4d, 0x49, 0x47, 0x52, 0x41, 0x54, 0x45, 0x5f,
	0x47, 0x48, 0x4f, 0x53, 0x54, 0x10, 0x04, 0x12, 0x0a, 0x0a, 0x06, 0x42, 0x52, 0x41, 0x4e, 0x43,
	0x48, 0x10, 0x05, 0x12, 0x08, 0x0a, 0x04, 0x44, 0x41, 0x54, 0x41, 0x10, 0x06, 0x1a, 0x9c, 0x02,
	0x0a, 0x15, 0x52, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73,
	0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12,
	0x6a, 0x0a, 0x16, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61,
	0x73, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x2f, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x48, 0x01, 0x52, 0x14, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61,
	0x73, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x88, 0x01, 0x01, 0x12, 0x18, 0x0a, 0x06, 0x62,
	0x61, 0x63, 0x6b, 0x75, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x06, 0x62,
	0x61, 0x63, 0x6b, 0x75, 0x70, 0x12, 0x40, 0x0a, 0x0d, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x69,
	0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x00, 0x52, 0x0b, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x49, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x42, 0x08, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x42, 0x19, 0x0a, 0x17, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x64, 0x61, 0x74,
	0x61, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x14, 0x5a, 0x12,
	0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2d, 0x67, 0x6f, 0x2f, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_store_plan_proto_rawDescOnce sync.Once
	file_store_plan_proto_rawDescData = file_store_plan_proto_rawDesc
)

func file_store_plan_proto_rawDescGZIP() []byte {
	file_store_plan_proto_rawDescOnce.Do(func() {
		file_store_plan_proto_rawDescData = protoimpl.X.CompressGZIP(file_store_plan_proto_rawDescData)
	})
	return file_store_plan_proto_rawDescData
}

var file_store_plan_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_store_plan_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_store_plan_proto_goTypes = []interface{}{
	(PlanConfig_ChangeDatabaseConfig_Type)(0), // 0: bytebase.store.PlanConfig.ChangeDatabaseConfig.Type
	(*PlanConfig)(nil),                        // 1: bytebase.store.PlanConfig
	(*PlanConfig_Step)(nil),                   // 2: bytebase.store.PlanConfig.Step
	(*PlanConfig_Spec)(nil),                   // 3: bytebase.store.PlanConfig.Spec
	(*PlanConfig_CreateDatabaseConfig)(nil),   // 4: bytebase.store.PlanConfig.CreateDatabaseConfig
	(*PlanConfig_ChangeDatabaseConfig)(nil),   // 5: bytebase.store.PlanConfig.ChangeDatabaseConfig
	(*PlanConfig_RestoreDatabaseConfig)(nil),  // 6: bytebase.store.PlanConfig.RestoreDatabaseConfig
	nil,                                       // 7: bytebase.store.PlanConfig.CreateDatabaseConfig.LabelsEntry
	(*timestamppb.Timestamp)(nil),             // 8: google.protobuf.Timestamp
}
var file_store_plan_proto_depIdxs = []int32{
	2,  // 0: bytebase.store.PlanConfig.steps:type_name -> bytebase.store.PlanConfig.Step
	3,  // 1: bytebase.store.PlanConfig.Step.specs:type_name -> bytebase.store.PlanConfig.Spec
	8,  // 2: bytebase.store.PlanConfig.Spec.earliest_allowed_time:type_name -> google.protobuf.Timestamp
	4,  // 3: bytebase.store.PlanConfig.Spec.create_database_config:type_name -> bytebase.store.PlanConfig.CreateDatabaseConfig
	5,  // 4: bytebase.store.PlanConfig.Spec.change_database_config:type_name -> bytebase.store.PlanConfig.ChangeDatabaseConfig
	6,  // 5: bytebase.store.PlanConfig.Spec.restore_database_config:type_name -> bytebase.store.PlanConfig.RestoreDatabaseConfig
	7,  // 6: bytebase.store.PlanConfig.CreateDatabaseConfig.labels:type_name -> bytebase.store.PlanConfig.CreateDatabaseConfig.LabelsEntry
	0,  // 7: bytebase.store.PlanConfig.ChangeDatabaseConfig.type:type_name -> bytebase.store.PlanConfig.ChangeDatabaseConfig.Type
	4,  // 8: bytebase.store.PlanConfig.RestoreDatabaseConfig.create_database_config:type_name -> bytebase.store.PlanConfig.CreateDatabaseConfig
	8,  // 9: bytebase.store.PlanConfig.RestoreDatabaseConfig.point_in_time:type_name -> google.protobuf.Timestamp
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_store_plan_proto_init() }
func file_store_plan_proto_init() {
	if File_store_plan_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_store_plan_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlanConfig); i {
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
		file_store_plan_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlanConfig_Step); i {
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
		file_store_plan_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlanConfig_Spec); i {
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
		file_store_plan_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlanConfig_CreateDatabaseConfig); i {
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
		file_store_plan_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlanConfig_ChangeDatabaseConfig); i {
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
		file_store_plan_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlanConfig_RestoreDatabaseConfig); i {
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
	file_store_plan_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*PlanConfig_Spec_CreateDatabaseConfig)(nil),
		(*PlanConfig_Spec_ChangeDatabaseConfig)(nil),
		(*PlanConfig_Spec_RestoreDatabaseConfig)(nil),
	}
	file_store_plan_proto_msgTypes[5].OneofWrappers = []interface{}{
		(*PlanConfig_RestoreDatabaseConfig_Backup)(nil),
		(*PlanConfig_RestoreDatabaseConfig_PointInTime)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_store_plan_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_store_plan_proto_goTypes,
		DependencyIndexes: file_store_plan_proto_depIdxs,
		EnumInfos:         file_store_plan_proto_enumTypes,
		MessageInfos:      file_store_plan_proto_msgTypes,
	}.Build()
	File_store_plan_proto = out.File
	file_store_plan_proto_rawDesc = nil
	file_store_plan_proto_goTypes = nil
	file_store_plan_proto_depIdxs = nil
}
