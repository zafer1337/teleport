// Copyright 2023 Gravitational, Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        (unknown)
// source: teleport/dbobjectimportrule/v1/dbobjectimportrule_service.proto

package dbobjectimportrulev1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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

// The request for CreateDatabaseObjectImportRule.
type CreateDatabaseObjectImportRuleRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The DatabaseObjectImportRule to create.
	Rule          *DatabaseObjectImportRule `protobuf:"bytes,1,opt,name=rule,proto3" json:"rule,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateDatabaseObjectImportRuleRequest) Reset() {
	*x = CreateDatabaseObjectImportRuleRequest{}
	mi := &file_teleport_dbobjectimportrule_v1_dbobjectimportrule_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateDatabaseObjectImportRuleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDatabaseObjectImportRuleRequest) ProtoMessage() {}

func (x *CreateDatabaseObjectImportRuleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_dbobjectimportrule_v1_dbobjectimportrule_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDatabaseObjectImportRuleRequest.ProtoReflect.Descriptor instead.
func (*CreateDatabaseObjectImportRuleRequest) Descriptor() ([]byte, []int) {
	return file_teleport_dbobjectimportrule_v1_dbobjectimportrule_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateDatabaseObjectImportRuleRequest) GetRule() *DatabaseObjectImportRule {
	if x != nil {
		return x.Rule
	}
	return nil
}

// The request for GetDatabaseObjectImportRule.
type GetDatabaseObjectImportRuleRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The name of the DatabaseObjectImportRule to fetch.
	Name          string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetDatabaseObjectImportRuleRequest) Reset() {
	*x = GetDatabaseObjectImportRuleRequest{}
	mi := &file_teleport_dbobjectimportrule_v1_dbobjectimportrule_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetDatabaseObjectImportRuleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDatabaseObjectImportRuleRequest) ProtoMessage() {}

func (x *GetDatabaseObjectImportRuleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_dbobjectimportrule_v1_dbobjectimportrule_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDatabaseObjectImportRuleRequest.ProtoReflect.Descriptor instead.
func (*GetDatabaseObjectImportRuleRequest) Descriptor() ([]byte, []int) {
	return file_teleport_dbobjectimportrule_v1_dbobjectimportrule_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetDatabaseObjectImportRuleRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// The request for ListDatabaseObjectImportRules.
type ListDatabaseObjectImportRulesRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The maximum number of items to return.
	// The server may impose a different page size at its discretion.
	PageSize int32 `protobuf:"varint,1,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// The page_token is the next_page_token value returned from a previous List request, if any.
	PageToken     string `protobuf:"bytes,2,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListDatabaseObjectImportRulesRequest) Reset() {
	*x = ListDatabaseObjectImportRulesRequest{}
	mi := &file_teleport_dbobjectimportrule_v1_dbobjectimportrule_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListDatabaseObjectImportRulesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDatabaseObjectImportRulesRequest) ProtoMessage() {}

func (x *ListDatabaseObjectImportRulesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_dbobjectimportrule_v1_dbobjectimportrule_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDatabaseObjectImportRulesRequest.ProtoReflect.Descriptor instead.
func (*ListDatabaseObjectImportRulesRequest) Descriptor() ([]byte, []int) {
	return file_teleport_dbobjectimportrule_v1_dbobjectimportrule_service_proto_rawDescGZIP(), []int{2}
}

func (x *ListDatabaseObjectImportRulesRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListDatabaseObjectImportRulesRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

// The response for ListDatabaseObjectImportRules.
type ListDatabaseObjectImportRulesResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The page of DatabaseObjectImportRules that matched the request.
	Rules []*DatabaseObjectImportRule `protobuf:"bytes,1,rep,name=rules,proto3" json:"rules,omitempty"`
	// Token to retrieve the next page of results, or empty if there are no
	// more results in the list.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListDatabaseObjectImportRulesResponse) Reset() {
	*x = ListDatabaseObjectImportRulesResponse{}
	mi := &file_teleport_dbobjectimportrule_v1_dbobjectimportrule_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListDatabaseObjectImportRulesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDatabaseObjectImportRulesResponse) ProtoMessage() {}

func (x *ListDatabaseObjectImportRulesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_dbobjectimportrule_v1_dbobjectimportrule_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDatabaseObjectImportRulesResponse.ProtoReflect.Descriptor instead.
func (*ListDatabaseObjectImportRulesResponse) Descriptor() ([]byte, []int) {
	return file_teleport_dbobjectimportrule_v1_dbobjectimportrule_service_proto_rawDescGZIP(), []int{3}
}

func (x *ListDatabaseObjectImportRulesResponse) GetRules() []*DatabaseObjectImportRule {
	if x != nil {
		return x.Rules
	}
	return nil
}

func (x *ListDatabaseObjectImportRulesResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

// The request for UpdateDatabaseObjectImportRule.
type UpdateDatabaseObjectImportRuleRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The DatabaseObjectImportRule to replace.
	Rule          *DatabaseObjectImportRule `protobuf:"bytes,1,opt,name=rule,proto3" json:"rule,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateDatabaseObjectImportRuleRequest) Reset() {
	*x = UpdateDatabaseObjectImportRuleRequest{}
	mi := &file_teleport_dbobjectimportrule_v1_dbobjectimportrule_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateDatabaseObjectImportRuleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateDatabaseObjectImportRuleRequest) ProtoMessage() {}

func (x *UpdateDatabaseObjectImportRuleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_dbobjectimportrule_v1_dbobjectimportrule_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateDatabaseObjectImportRuleRequest.ProtoReflect.Descriptor instead.
func (*UpdateDatabaseObjectImportRuleRequest) Descriptor() ([]byte, []int) {
	return file_teleport_dbobjectimportrule_v1_dbobjectimportrule_service_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateDatabaseObjectImportRuleRequest) GetRule() *DatabaseObjectImportRule {
	if x != nil {
		return x.Rule
	}
	return nil
}

// The request for UpsertDatabaseObjectImportRule.
type UpsertDatabaseObjectImportRuleRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The DatabaseObjectImportRule to create or replace.
	Rule          *DatabaseObjectImportRule `protobuf:"bytes,1,opt,name=rule,proto3" json:"rule,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpsertDatabaseObjectImportRuleRequest) Reset() {
	*x = UpsertDatabaseObjectImportRuleRequest{}
	mi := &file_teleport_dbobjectimportrule_v1_dbobjectimportrule_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpsertDatabaseObjectImportRuleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpsertDatabaseObjectImportRuleRequest) ProtoMessage() {}

func (x *UpsertDatabaseObjectImportRuleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_dbobjectimportrule_v1_dbobjectimportrule_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpsertDatabaseObjectImportRuleRequest.ProtoReflect.Descriptor instead.
func (*UpsertDatabaseObjectImportRuleRequest) Descriptor() ([]byte, []int) {
	return file_teleport_dbobjectimportrule_v1_dbobjectimportrule_service_proto_rawDescGZIP(), []int{5}
}

func (x *UpsertDatabaseObjectImportRuleRequest) GetRule() *DatabaseObjectImportRule {
	if x != nil {
		return x.Rule
	}
	return nil
}

// The request for DeleteDatabaseObjectImportRule.
type DeleteDatabaseObjectImportRuleRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The name of the DatabaseObjectImportRule to delete.
	Name          string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteDatabaseObjectImportRuleRequest) Reset() {
	*x = DeleteDatabaseObjectImportRuleRequest{}
	mi := &file_teleport_dbobjectimportrule_v1_dbobjectimportrule_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteDatabaseObjectImportRuleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteDatabaseObjectImportRuleRequest) ProtoMessage() {}

func (x *DeleteDatabaseObjectImportRuleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_dbobjectimportrule_v1_dbobjectimportrule_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteDatabaseObjectImportRuleRequest.ProtoReflect.Descriptor instead.
func (*DeleteDatabaseObjectImportRuleRequest) Descriptor() ([]byte, []int) {
	return file_teleport_dbobjectimportrule_v1_dbobjectimportrule_service_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteDatabaseObjectImportRuleRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_teleport_dbobjectimportrule_v1_dbobjectimportrule_service_proto protoreflect.FileDescriptor

var file_teleport_dbobjectimportrule_v1_dbobjectimportrule_service_proto_rawDesc = string([]byte{
	0x0a, 0x3f, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x64, 0x62, 0x6f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x72, 0x75, 0x6c, 0x65, 0x2f, 0x76, 0x31,
	0x2f, 0x64, 0x62, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x72,
	0x75, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x1e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x64, 0x62, 0x6f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x72, 0x75, 0x6c, 0x65, 0x2e, 0x76,
	0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x37,
	0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x64, 0x62, 0x6f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x72, 0x75, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x64,
	0x62, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x72, 0x75, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x75, 0x0a, 0x25, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49,
	0x6d, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x4c, 0x0a, 0x04, 0x72, 0x75, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x38,
	0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x64, 0x62, 0x6f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x72, 0x75, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x6d,
	0x70, 0x6f, 0x72, 0x74, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x04, 0x72, 0x75, 0x6c, 0x65, 0x22, 0x38,
	0x0a, 0x22, 0x47, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x4f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x62, 0x0a, 0x24, 0x4c, 0x69, 0x73, 0x74,
	0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x6d,
	0x70, 0x6f, 0x72, 0x74, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x9f, 0x01, 0x0a,
	0x25, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x4f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4e, 0x0a, 0x05, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x2e, 0x64, 0x62, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x72,
	0x75, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x4f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x75, 0x6c, 0x65, 0x52,
	0x05, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70,
	0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x75,
	0x0a, 0x25, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65,
	0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x75, 0x6c, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4c, 0x0a, 0x04, 0x72, 0x75, 0x6c, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x2e, 0x64, 0x62, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x72,
	0x75, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x4f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x75, 0x6c, 0x65, 0x52,
	0x04, 0x72, 0x75, 0x6c, 0x65, 0x22, 0x75, 0x0a, 0x25, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x44,
	0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x6d, 0x70,
	0x6f, 0x72, 0x74, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4c,
	0x0a, 0x04, 0x72, 0x75, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x74,
	0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x64, 0x62, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x72, 0x75, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61,
	0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x6d, 0x70, 0x6f,
	0x72, 0x74, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x04, 0x72, 0x75, 0x6c, 0x65, 0x22, 0x3b, 0x0a, 0x25,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x4f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x32, 0xdb, 0x07, 0x0a, 0x1f, 0x44, 0x61,
	0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x6d, 0x70, 0x6f,
	0x72, 0x74, 0x52, 0x75, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x9b, 0x01,
	0x0a, 0x1b, 0x47, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x4f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x75, 0x6c, 0x65, 0x12, 0x42, 0x2e,
	0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x64, 0x62, 0x6f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x72, 0x75, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x38, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x64, 0x62, 0x6f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x72, 0x75, 0x6c, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x75, 0x6c, 0x65, 0x12, 0xac, 0x01, 0x0a, 0x1d,
	0x4c, 0x69, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x4f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x12, 0x44, 0x2e,
	0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x64, 0x62, 0x6f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x72, 0x75, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x45, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x64,
	0x62, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x72, 0x75, 0x6c,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73,
	0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x75, 0x6c,
	0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0xa1, 0x01, 0x0a, 0x1e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x4f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x75, 0x6c, 0x65, 0x12, 0x45, 0x2e,
	0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x64, 0x62, 0x6f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x72, 0x75, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x4f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x38, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e,
	0x64, 0x62, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x72, 0x75,
	0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x4f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x75, 0x6c, 0x65, 0x12, 0xa1,
	0x01, 0x0a, 0x1e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73,
	0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x75, 0x6c,
	0x65, 0x12, 0x45, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x64, 0x62, 0x6f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x72, 0x75, 0x6c, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73,
	0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x75, 0x6c,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x38, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x2e, 0x64, 0x62, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x6d, 0x70, 0x6f,
	0x72, 0x74, 0x72, 0x75, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61,
	0x73, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x75,
	0x6c, 0x65, 0x12, 0xa1, 0x01, 0x0a, 0x1e, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x44, 0x61, 0x74,
	0x61, 0x62, 0x61, 0x73, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x6d, 0x70, 0x6f, 0x72,
	0x74, 0x52, 0x75, 0x6c, 0x65, 0x12, 0x45, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x2e, 0x64, 0x62, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x72,
	0x75, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x44, 0x61, 0x74,
	0x61, 0x62, 0x61, 0x73, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x6d, 0x70, 0x6f, 0x72,
	0x74, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x38, 0x2e, 0x74,
	0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x64, 0x62, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x72, 0x75, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61,
	0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x6d, 0x70, 0x6f,
	0x72, 0x74, 0x52, 0x75, 0x6c, 0x65, 0x12, 0x7f, 0x0a, 0x1e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x6d,
	0x70, 0x6f, 0x72, 0x74, 0x52, 0x75, 0x6c, 0x65, 0x12, 0x45, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x2e, 0x64, 0x62, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x6d, 0x70, 0x6f,
	0x72, 0x74, 0x72, 0x75, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x6d,
	0x70, 0x6f, 0x72, 0x74, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x68, 0x5a, 0x66, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x72, 0x61, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x74, 0x65,
	0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x64, 0x62, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x69,
	0x6d, 0x70, 0x6f, 0x72, 0x74, 0x72, 0x75, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x64, 0x62, 0x6f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x72, 0x75, 0x6c, 0x65, 0x76,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_teleport_dbobjectimportrule_v1_dbobjectimportrule_service_proto_rawDescOnce sync.Once
	file_teleport_dbobjectimportrule_v1_dbobjectimportrule_service_proto_rawDescData []byte
)

func file_teleport_dbobjectimportrule_v1_dbobjectimportrule_service_proto_rawDescGZIP() []byte {
	file_teleport_dbobjectimportrule_v1_dbobjectimportrule_service_proto_rawDescOnce.Do(func() {
		file_teleport_dbobjectimportrule_v1_dbobjectimportrule_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_teleport_dbobjectimportrule_v1_dbobjectimportrule_service_proto_rawDesc), len(file_teleport_dbobjectimportrule_v1_dbobjectimportrule_service_proto_rawDesc)))
	})
	return file_teleport_dbobjectimportrule_v1_dbobjectimportrule_service_proto_rawDescData
}

var file_teleport_dbobjectimportrule_v1_dbobjectimportrule_service_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_teleport_dbobjectimportrule_v1_dbobjectimportrule_service_proto_goTypes = []any{
	(*CreateDatabaseObjectImportRuleRequest)(nil), // 0: teleport.dbobjectimportrule.v1.CreateDatabaseObjectImportRuleRequest
	(*GetDatabaseObjectImportRuleRequest)(nil),    // 1: teleport.dbobjectimportrule.v1.GetDatabaseObjectImportRuleRequest
	(*ListDatabaseObjectImportRulesRequest)(nil),  // 2: teleport.dbobjectimportrule.v1.ListDatabaseObjectImportRulesRequest
	(*ListDatabaseObjectImportRulesResponse)(nil), // 3: teleport.dbobjectimportrule.v1.ListDatabaseObjectImportRulesResponse
	(*UpdateDatabaseObjectImportRuleRequest)(nil), // 4: teleport.dbobjectimportrule.v1.UpdateDatabaseObjectImportRuleRequest
	(*UpsertDatabaseObjectImportRuleRequest)(nil), // 5: teleport.dbobjectimportrule.v1.UpsertDatabaseObjectImportRuleRequest
	(*DeleteDatabaseObjectImportRuleRequest)(nil), // 6: teleport.dbobjectimportrule.v1.DeleteDatabaseObjectImportRuleRequest
	(*DatabaseObjectImportRule)(nil),              // 7: teleport.dbobjectimportrule.v1.DatabaseObjectImportRule
	(*emptypb.Empty)(nil),                         // 8: google.protobuf.Empty
}
var file_teleport_dbobjectimportrule_v1_dbobjectimportrule_service_proto_depIdxs = []int32{
	7,  // 0: teleport.dbobjectimportrule.v1.CreateDatabaseObjectImportRuleRequest.rule:type_name -> teleport.dbobjectimportrule.v1.DatabaseObjectImportRule
	7,  // 1: teleport.dbobjectimportrule.v1.ListDatabaseObjectImportRulesResponse.rules:type_name -> teleport.dbobjectimportrule.v1.DatabaseObjectImportRule
	7,  // 2: teleport.dbobjectimportrule.v1.UpdateDatabaseObjectImportRuleRequest.rule:type_name -> teleport.dbobjectimportrule.v1.DatabaseObjectImportRule
	7,  // 3: teleport.dbobjectimportrule.v1.UpsertDatabaseObjectImportRuleRequest.rule:type_name -> teleport.dbobjectimportrule.v1.DatabaseObjectImportRule
	1,  // 4: teleport.dbobjectimportrule.v1.DatabaseObjectImportRuleService.GetDatabaseObjectImportRule:input_type -> teleport.dbobjectimportrule.v1.GetDatabaseObjectImportRuleRequest
	2,  // 5: teleport.dbobjectimportrule.v1.DatabaseObjectImportRuleService.ListDatabaseObjectImportRules:input_type -> teleport.dbobjectimportrule.v1.ListDatabaseObjectImportRulesRequest
	0,  // 6: teleport.dbobjectimportrule.v1.DatabaseObjectImportRuleService.CreateDatabaseObjectImportRule:input_type -> teleport.dbobjectimportrule.v1.CreateDatabaseObjectImportRuleRequest
	4,  // 7: teleport.dbobjectimportrule.v1.DatabaseObjectImportRuleService.UpdateDatabaseObjectImportRule:input_type -> teleport.dbobjectimportrule.v1.UpdateDatabaseObjectImportRuleRequest
	5,  // 8: teleport.dbobjectimportrule.v1.DatabaseObjectImportRuleService.UpsertDatabaseObjectImportRule:input_type -> teleport.dbobjectimportrule.v1.UpsertDatabaseObjectImportRuleRequest
	6,  // 9: teleport.dbobjectimportrule.v1.DatabaseObjectImportRuleService.DeleteDatabaseObjectImportRule:input_type -> teleport.dbobjectimportrule.v1.DeleteDatabaseObjectImportRuleRequest
	7,  // 10: teleport.dbobjectimportrule.v1.DatabaseObjectImportRuleService.GetDatabaseObjectImportRule:output_type -> teleport.dbobjectimportrule.v1.DatabaseObjectImportRule
	3,  // 11: teleport.dbobjectimportrule.v1.DatabaseObjectImportRuleService.ListDatabaseObjectImportRules:output_type -> teleport.dbobjectimportrule.v1.ListDatabaseObjectImportRulesResponse
	7,  // 12: teleport.dbobjectimportrule.v1.DatabaseObjectImportRuleService.CreateDatabaseObjectImportRule:output_type -> teleport.dbobjectimportrule.v1.DatabaseObjectImportRule
	7,  // 13: teleport.dbobjectimportrule.v1.DatabaseObjectImportRuleService.UpdateDatabaseObjectImportRule:output_type -> teleport.dbobjectimportrule.v1.DatabaseObjectImportRule
	7,  // 14: teleport.dbobjectimportrule.v1.DatabaseObjectImportRuleService.UpsertDatabaseObjectImportRule:output_type -> teleport.dbobjectimportrule.v1.DatabaseObjectImportRule
	8,  // 15: teleport.dbobjectimportrule.v1.DatabaseObjectImportRuleService.DeleteDatabaseObjectImportRule:output_type -> google.protobuf.Empty
	10, // [10:16] is the sub-list for method output_type
	4,  // [4:10] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_teleport_dbobjectimportrule_v1_dbobjectimportrule_service_proto_init() }
func file_teleport_dbobjectimportrule_v1_dbobjectimportrule_service_proto_init() {
	if File_teleport_dbobjectimportrule_v1_dbobjectimportrule_service_proto != nil {
		return
	}
	file_teleport_dbobjectimportrule_v1_dbobjectimportrule_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_teleport_dbobjectimportrule_v1_dbobjectimportrule_service_proto_rawDesc), len(file_teleport_dbobjectimportrule_v1_dbobjectimportrule_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_teleport_dbobjectimportrule_v1_dbobjectimportrule_service_proto_goTypes,
		DependencyIndexes: file_teleport_dbobjectimportrule_v1_dbobjectimportrule_service_proto_depIdxs,
		MessageInfos:      file_teleport_dbobjectimportrule_v1_dbobjectimportrule_service_proto_msgTypes,
	}.Build()
	File_teleport_dbobjectimportrule_v1_dbobjectimportrule_service_proto = out.File
	file_teleport_dbobjectimportrule_v1_dbobjectimportrule_service_proto_goTypes = nil
	file_teleport_dbobjectimportrule_v1_dbobjectimportrule_service_proto_depIdxs = nil
}
