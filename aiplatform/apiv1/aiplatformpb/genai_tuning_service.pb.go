// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.2
// source: google/cloud/aiplatform/v1/genai_tuning_service.proto

package aiplatformpb

import (
	context "context"
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Request message for
// [GenAiTuningService.CreateTuningJob][google.cloud.aiplatform.v1.GenAiTuningService.CreateTuningJob].
type CreateTuningJobRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The resource name of the Location to create the TuningJob in.
	// Format: `projects/{project}/locations/{location}`
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// Required. The TuningJob to create.
	TuningJob *TuningJob `protobuf:"bytes,2,opt,name=tuning_job,json=tuningJob,proto3" json:"tuning_job,omitempty"`
}

func (x *CreateTuningJobRequest) Reset() {
	*x = CreateTuningJobRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_aiplatform_v1_genai_tuning_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTuningJobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTuningJobRequest) ProtoMessage() {}

func (x *CreateTuningJobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1_genai_tuning_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTuningJobRequest.ProtoReflect.Descriptor instead.
func (*CreateTuningJobRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1_genai_tuning_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateTuningJobRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *CreateTuningJobRequest) GetTuningJob() *TuningJob {
	if x != nil {
		return x.TuningJob
	}
	return nil
}

// Request message for
// [GenAiTuningService.GetTuningJob][google.cloud.aiplatform.v1.GenAiTuningService.GetTuningJob].
type GetTuningJobRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The name of the TuningJob resource. Format:
	// `projects/{project}/locations/{location}/tuningJobs/{tuning_job}`
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetTuningJobRequest) Reset() {
	*x = GetTuningJobRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_aiplatform_v1_genai_tuning_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTuningJobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTuningJobRequest) ProtoMessage() {}

func (x *GetTuningJobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1_genai_tuning_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTuningJobRequest.ProtoReflect.Descriptor instead.
func (*GetTuningJobRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1_genai_tuning_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetTuningJobRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// Request message for
// [GenAiTuningService.ListTuningJobs][google.cloud.aiplatform.v1.GenAiTuningService.ListTuningJobs].
type ListTuningJobsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The resource name of the Location to list the TuningJobs from.
	// Format: `projects/{project}/locations/{location}`
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// Optional. The standard list filter.
	Filter string `protobuf:"bytes,2,opt,name=filter,proto3" json:"filter,omitempty"`
	// Optional. The standard list page size.
	PageSize int32 `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Optional. The standard list page token.
	// Typically obtained via [ListTuningJob.next_page_token][] of the
	// previous GenAiTuningService.ListTuningJob][] call.
	PageToken string `protobuf:"bytes,4,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
}

func (x *ListTuningJobsRequest) Reset() {
	*x = ListTuningJobsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_aiplatform_v1_genai_tuning_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTuningJobsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTuningJobsRequest) ProtoMessage() {}

func (x *ListTuningJobsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1_genai_tuning_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTuningJobsRequest.ProtoReflect.Descriptor instead.
func (*ListTuningJobsRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1_genai_tuning_service_proto_rawDescGZIP(), []int{2}
}

func (x *ListTuningJobsRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *ListTuningJobsRequest) GetFilter() string {
	if x != nil {
		return x.Filter
	}
	return ""
}

func (x *ListTuningJobsRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListTuningJobsRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

// Response message for
// [GenAiTuningService.ListTuningJobs][google.cloud.aiplatform.v1.GenAiTuningService.ListTuningJobs]
type ListTuningJobsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// List of TuningJobs in the requested page.
	TuningJobs []*TuningJob `protobuf:"bytes,1,rep,name=tuning_jobs,json=tuningJobs,proto3" json:"tuning_jobs,omitempty"`
	// A token to retrieve the next page of results.
	// Pass to
	// [ListTuningJobsRequest.page_token][google.cloud.aiplatform.v1.ListTuningJobsRequest.page_token]
	// to obtain that page.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ListTuningJobsResponse) Reset() {
	*x = ListTuningJobsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_aiplatform_v1_genai_tuning_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTuningJobsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTuningJobsResponse) ProtoMessage() {}

func (x *ListTuningJobsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1_genai_tuning_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTuningJobsResponse.ProtoReflect.Descriptor instead.
func (*ListTuningJobsResponse) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1_genai_tuning_service_proto_rawDescGZIP(), []int{3}
}

func (x *ListTuningJobsResponse) GetTuningJobs() []*TuningJob {
	if x != nil {
		return x.TuningJobs
	}
	return nil
}

func (x *ListTuningJobsResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

// Request message for
// [GenAiTuningService.CancelTuningJob][google.cloud.aiplatform.v1.GenAiTuningService.CancelTuningJob].
type CancelTuningJobRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The name of the TuningJob to cancel. Format:
	// `projects/{project}/locations/{location}/tuningJobs/{tuning_job}`
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CancelTuningJobRequest) Reset() {
	*x = CancelTuningJobRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_aiplatform_v1_genai_tuning_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CancelTuningJobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelTuningJobRequest) ProtoMessage() {}

func (x *CancelTuningJobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1_genai_tuning_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelTuningJobRequest.ProtoReflect.Descriptor instead.
func (*CancelTuningJobRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1_genai_tuning_service_proto_rawDescGZIP(), []int{4}
}

func (x *CancelTuningJobRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_google_cloud_aiplatform_v1_genai_tuning_service_proto protoreflect.FileDescriptor

var file_google_cloud_aiplatform_v1_genai_tuning_service_proto_rawDesc = []byte{
	0x0a, 0x35, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61,
	0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x6e,
	0x61, 0x69, 0x5f, 0x74, 0x75, 0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68,
	0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f,
	0x76, 0x31, 0x2f, 0x74, 0x75, 0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x6a, 0x6f, 0x62, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xa6, 0x01, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x75, 0x6e, 0x69, 0x6e,
	0x67, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x41, 0x0a, 0x06, 0x70,
	0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x29, 0xe0, 0x41, 0x02,
	0xfa, 0x41, 0x23, 0x0a, 0x21, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x49,
	0x0a, 0x0a, 0x74, 0x75, 0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x6a, 0x6f, 0x62, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x25, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e,
	0x54, 0x75, 0x6e, 0x69, 0x6e, 0x67, 0x4a, 0x6f, 0x62, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x09,
	0x74, 0x75, 0x6e, 0x69, 0x6e, 0x67, 0x4a, 0x6f, 0x62, 0x22, 0x56, 0x0a, 0x13, 0x47, 0x65, 0x74,
	0x54, 0x75, 0x6e, 0x69, 0x6e, 0x67, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x3f, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2b,
	0xe0, 0x41, 0x02, 0xfa, 0x41, 0x25, 0x0a, 0x23, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x54, 0x75, 0x6e, 0x69, 0x6e, 0x67, 0x4a, 0x6f, 0x62, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0xbd, 0x01, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x75, 0x6e, 0x69, 0x6e, 0x67,
	0x4a, 0x6f, 0x62, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x41, 0x0a, 0x06, 0x70,
	0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x29, 0xe0, 0x41, 0x02,
	0xfa, 0x41, 0x23, 0x0a, 0x21, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x1b,
	0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03,
	0xe0, 0x41, 0x01, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x20, 0x0a, 0x09, 0x70,
	0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x42, 0x03,
	0xe0, 0x41, 0x01, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x22, 0x0a,
	0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x22, 0x88, 0x01, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x75, 0x6e, 0x69, 0x6e, 0x67,
	0x4a, 0x6f, 0x62, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a, 0x0b,
	0x74, 0x75, 0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x6a, 0x6f, 0x62, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x25, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x54,
	0x75, 0x6e, 0x69, 0x6e, 0x67, 0x4a, 0x6f, 0x62, 0x52, 0x0a, 0x74, 0x75, 0x6e, 0x69, 0x6e, 0x67,
	0x4a, 0x6f, 0x62, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67,
	0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e,
	0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x59, 0x0a, 0x16,
	0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x54, 0x75, 0x6e, 0x69, 0x6e, 0x67, 0x4a, 0x6f, 0x62, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3f, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x2b, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x25, 0x0a, 0x23, 0x61, 0x69,
	0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x54, 0x75, 0x6e, 0x69, 0x6e, 0x67, 0x4a, 0x6f,
	0x62, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x32, 0xb6, 0x06, 0x0a, 0x12, 0x47, 0x65, 0x6e, 0x41,
	0x69, 0x54, 0x75, 0x6e, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xc4,
	0x01, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x75, 0x6e, 0x69, 0x6e, 0x67, 0x4a,
	0x6f, 0x62, 0x12, 0x32, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x75, 0x6e, 0x69, 0x6e, 0x67, 0x4a, 0x6f, 0x62, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x2e, 0x76, 0x31, 0x2e, 0x54, 0x75, 0x6e, 0x69, 0x6e, 0x67, 0x4a, 0x6f, 0x62, 0x22, 0x56, 0xda,
	0x41, 0x11, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x2c, 0x74, 0x75, 0x6e, 0x69, 0x6e, 0x67, 0x5f,
	0x6a, 0x6f, 0x62, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x3c, 0x3a, 0x0a, 0x74, 0x75, 0x6e, 0x69, 0x6e,
	0x67, 0x5f, 0x6a, 0x6f, 0x62, 0x22, 0x2e, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x70, 0x61, 0x72, 0x65,
	0x6e, 0x74, 0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x2a, 0x7d, 0x2f, 0x74, 0x75, 0x6e, 0x69, 0x6e,
	0x67, 0x4a, 0x6f, 0x62, 0x73, 0x12, 0xa5, 0x01, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x54, 0x75, 0x6e,
	0x69, 0x6e, 0x67, 0x4a, 0x6f, 0x62, 0x12, 0x2f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x75, 0x6e, 0x69, 0x6e, 0x67, 0x4a, 0x6f, 0x62,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x75, 0x6e, 0x69, 0x6e, 0x67, 0x4a, 0x6f, 0x62, 0x22, 0x3d,
	0xda, 0x41, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x30, 0x12, 0x2e, 0x2f,
	0x76, 0x31, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x73, 0x2f, 0x2a, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x2a, 0x2f,
	0x74, 0x75, 0x6e, 0x69, 0x6e, 0x67, 0x4a, 0x6f, 0x62, 0x73, 0x2f, 0x2a, 0x7d, 0x12, 0xb8, 0x01,
	0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x75, 0x6e, 0x69, 0x6e, 0x67, 0x4a, 0x6f, 0x62, 0x73,
	0x12, 0x31, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x54, 0x75, 0x6e, 0x69, 0x6e, 0x67, 0x4a, 0x6f, 0x62, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x75, 0x6e, 0x69, 0x6e, 0x67, 0x4a, 0x6f, 0x62, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3f, 0xda, 0x41, 0x06, 0x70, 0x61, 0x72, 0x65,
	0x6e, 0x74, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x30, 0x12, 0x2e, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x70,
	0x61, 0x72, 0x65, 0x6e, 0x74, 0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a,
	0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x2a, 0x7d, 0x2f, 0x74, 0x75,
	0x6e, 0x69, 0x6e, 0x67, 0x4a, 0x6f, 0x62, 0x73, 0x12, 0xa6, 0x01, 0x0a, 0x0f, 0x43, 0x61, 0x6e,
	0x63, 0x65, 0x6c, 0x54, 0x75, 0x6e, 0x69, 0x6e, 0x67, 0x4a, 0x6f, 0x62, 0x12, 0x32, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c,
	0x54, 0x75, 0x6e, 0x69, 0x6e, 0x67, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x47, 0xda, 0x41, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x3a, 0x3a, 0x01, 0x2a, 0x22, 0x35, 0x2f, 0x76, 0x31, 0x2f,
	0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a,
	0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x2a, 0x2f, 0x74, 0x75, 0x6e,
	0x69, 0x6e, 0x67, 0x4a, 0x6f, 0x62, 0x73, 0x2f, 0x2a, 0x7d, 0x3a, 0x63, 0x61, 0x6e, 0x63, 0x65,
	0x6c, 0x1a, 0x4d, 0xca, 0x41, 0x19, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0xd2,
	0x41, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x77, 0x77, 0x77, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x75, 0x74,
	0x68, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x42, 0xd5, 0x01, 0x0a, 0x1e, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x2e, 0x76, 0x31, 0x42, 0x17, 0x47, 0x65, 0x6e, 0x41, 0x69, 0x54, 0x75, 0x6e, 0x69, 0x6e, 0x67,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x67, 0x6f, 0x2f, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x61,
	0x70, 0x69, 0x76, 0x31, 0x2f, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x70,
	0x62, 0x3b, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x70, 0x62, 0xaa, 0x02,
	0x1a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x41, 0x49,
	0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x1a, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x41, 0x49, 0x50, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x5c, 0x56, 0x31, 0xea, 0x02, 0x1d, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x41, 0x49, 0x50, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_aiplatform_v1_genai_tuning_service_proto_rawDescOnce sync.Once
	file_google_cloud_aiplatform_v1_genai_tuning_service_proto_rawDescData = file_google_cloud_aiplatform_v1_genai_tuning_service_proto_rawDesc
)

func file_google_cloud_aiplatform_v1_genai_tuning_service_proto_rawDescGZIP() []byte {
	file_google_cloud_aiplatform_v1_genai_tuning_service_proto_rawDescOnce.Do(func() {
		file_google_cloud_aiplatform_v1_genai_tuning_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_aiplatform_v1_genai_tuning_service_proto_rawDescData)
	})
	return file_google_cloud_aiplatform_v1_genai_tuning_service_proto_rawDescData
}

var file_google_cloud_aiplatform_v1_genai_tuning_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_google_cloud_aiplatform_v1_genai_tuning_service_proto_goTypes = []interface{}{
	(*CreateTuningJobRequest)(nil), // 0: google.cloud.aiplatform.v1.CreateTuningJobRequest
	(*GetTuningJobRequest)(nil),    // 1: google.cloud.aiplatform.v1.GetTuningJobRequest
	(*ListTuningJobsRequest)(nil),  // 2: google.cloud.aiplatform.v1.ListTuningJobsRequest
	(*ListTuningJobsResponse)(nil), // 3: google.cloud.aiplatform.v1.ListTuningJobsResponse
	(*CancelTuningJobRequest)(nil), // 4: google.cloud.aiplatform.v1.CancelTuningJobRequest
	(*TuningJob)(nil),              // 5: google.cloud.aiplatform.v1.TuningJob
	(*emptypb.Empty)(nil),          // 6: google.protobuf.Empty
}
var file_google_cloud_aiplatform_v1_genai_tuning_service_proto_depIdxs = []int32{
	5, // 0: google.cloud.aiplatform.v1.CreateTuningJobRequest.tuning_job:type_name -> google.cloud.aiplatform.v1.TuningJob
	5, // 1: google.cloud.aiplatform.v1.ListTuningJobsResponse.tuning_jobs:type_name -> google.cloud.aiplatform.v1.TuningJob
	0, // 2: google.cloud.aiplatform.v1.GenAiTuningService.CreateTuningJob:input_type -> google.cloud.aiplatform.v1.CreateTuningJobRequest
	1, // 3: google.cloud.aiplatform.v1.GenAiTuningService.GetTuningJob:input_type -> google.cloud.aiplatform.v1.GetTuningJobRequest
	2, // 4: google.cloud.aiplatform.v1.GenAiTuningService.ListTuningJobs:input_type -> google.cloud.aiplatform.v1.ListTuningJobsRequest
	4, // 5: google.cloud.aiplatform.v1.GenAiTuningService.CancelTuningJob:input_type -> google.cloud.aiplatform.v1.CancelTuningJobRequest
	5, // 6: google.cloud.aiplatform.v1.GenAiTuningService.CreateTuningJob:output_type -> google.cloud.aiplatform.v1.TuningJob
	5, // 7: google.cloud.aiplatform.v1.GenAiTuningService.GetTuningJob:output_type -> google.cloud.aiplatform.v1.TuningJob
	3, // 8: google.cloud.aiplatform.v1.GenAiTuningService.ListTuningJobs:output_type -> google.cloud.aiplatform.v1.ListTuningJobsResponse
	6, // 9: google.cloud.aiplatform.v1.GenAiTuningService.CancelTuningJob:output_type -> google.protobuf.Empty
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_google_cloud_aiplatform_v1_genai_tuning_service_proto_init() }
func file_google_cloud_aiplatform_v1_genai_tuning_service_proto_init() {
	if File_google_cloud_aiplatform_v1_genai_tuning_service_proto != nil {
		return
	}
	file_google_cloud_aiplatform_v1_tuning_job_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_aiplatform_v1_genai_tuning_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTuningJobRequest); i {
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
		file_google_cloud_aiplatform_v1_genai_tuning_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTuningJobRequest); i {
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
		file_google_cloud_aiplatform_v1_genai_tuning_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTuningJobsRequest); i {
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
		file_google_cloud_aiplatform_v1_genai_tuning_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTuningJobsResponse); i {
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
		file_google_cloud_aiplatform_v1_genai_tuning_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CancelTuningJobRequest); i {
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
			RawDescriptor: file_google_cloud_aiplatform_v1_genai_tuning_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_cloud_aiplatform_v1_genai_tuning_service_proto_goTypes,
		DependencyIndexes: file_google_cloud_aiplatform_v1_genai_tuning_service_proto_depIdxs,
		MessageInfos:      file_google_cloud_aiplatform_v1_genai_tuning_service_proto_msgTypes,
	}.Build()
	File_google_cloud_aiplatform_v1_genai_tuning_service_proto = out.File
	file_google_cloud_aiplatform_v1_genai_tuning_service_proto_rawDesc = nil
	file_google_cloud_aiplatform_v1_genai_tuning_service_proto_goTypes = nil
	file_google_cloud_aiplatform_v1_genai_tuning_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// GenAiTuningServiceClient is the client API for GenAiTuningService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GenAiTuningServiceClient interface {
	// Creates a TuningJob. A created TuningJob right away will be attempted to
	// be run.
	CreateTuningJob(ctx context.Context, in *CreateTuningJobRequest, opts ...grpc.CallOption) (*TuningJob, error)
	// Gets a TuningJob.
	GetTuningJob(ctx context.Context, in *GetTuningJobRequest, opts ...grpc.CallOption) (*TuningJob, error)
	// Lists TuningJobs in a Location.
	ListTuningJobs(ctx context.Context, in *ListTuningJobsRequest, opts ...grpc.CallOption) (*ListTuningJobsResponse, error)
	// Cancels a TuningJob.
	// Starts asynchronous cancellation on the TuningJob. The server makes a best
	// effort to cancel the job, but success is not guaranteed. Clients can use
	// [GenAiTuningService.GetTuningJob][google.cloud.aiplatform.v1.GenAiTuningService.GetTuningJob]
	// or other methods to check whether the cancellation succeeded or whether the
	// job completed despite cancellation. On successful cancellation, the
	// TuningJob is not deleted; instead it becomes a job with a
	// [TuningJob.error][google.cloud.aiplatform.v1.TuningJob.error] value with a
	// [google.rpc.Status.code][google.rpc.Status.code] of 1, corresponding to
	// `Code.CANCELLED`, and
	// [TuningJob.state][google.cloud.aiplatform.v1.TuningJob.state] is set to
	// `CANCELLED`.
	CancelTuningJob(ctx context.Context, in *CancelTuningJobRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type genAiTuningServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGenAiTuningServiceClient(cc grpc.ClientConnInterface) GenAiTuningServiceClient {
	return &genAiTuningServiceClient{cc}
}

func (c *genAiTuningServiceClient) CreateTuningJob(ctx context.Context, in *CreateTuningJobRequest, opts ...grpc.CallOption) (*TuningJob, error) {
	out := new(TuningJob)
	err := c.cc.Invoke(ctx, "/google.cloud.aiplatform.v1.GenAiTuningService/CreateTuningJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *genAiTuningServiceClient) GetTuningJob(ctx context.Context, in *GetTuningJobRequest, opts ...grpc.CallOption) (*TuningJob, error) {
	out := new(TuningJob)
	err := c.cc.Invoke(ctx, "/google.cloud.aiplatform.v1.GenAiTuningService/GetTuningJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *genAiTuningServiceClient) ListTuningJobs(ctx context.Context, in *ListTuningJobsRequest, opts ...grpc.CallOption) (*ListTuningJobsResponse, error) {
	out := new(ListTuningJobsResponse)
	err := c.cc.Invoke(ctx, "/google.cloud.aiplatform.v1.GenAiTuningService/ListTuningJobs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *genAiTuningServiceClient) CancelTuningJob(ctx context.Context, in *CancelTuningJobRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/google.cloud.aiplatform.v1.GenAiTuningService/CancelTuningJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GenAiTuningServiceServer is the server API for GenAiTuningService service.
type GenAiTuningServiceServer interface {
	// Creates a TuningJob. A created TuningJob right away will be attempted to
	// be run.
	CreateTuningJob(context.Context, *CreateTuningJobRequest) (*TuningJob, error)
	// Gets a TuningJob.
	GetTuningJob(context.Context, *GetTuningJobRequest) (*TuningJob, error)
	// Lists TuningJobs in a Location.
	ListTuningJobs(context.Context, *ListTuningJobsRequest) (*ListTuningJobsResponse, error)
	// Cancels a TuningJob.
	// Starts asynchronous cancellation on the TuningJob. The server makes a best
	// effort to cancel the job, but success is not guaranteed. Clients can use
	// [GenAiTuningService.GetTuningJob][google.cloud.aiplatform.v1.GenAiTuningService.GetTuningJob]
	// or other methods to check whether the cancellation succeeded or whether the
	// job completed despite cancellation. On successful cancellation, the
	// TuningJob is not deleted; instead it becomes a job with a
	// [TuningJob.error][google.cloud.aiplatform.v1.TuningJob.error] value with a
	// [google.rpc.Status.code][google.rpc.Status.code] of 1, corresponding to
	// `Code.CANCELLED`, and
	// [TuningJob.state][google.cloud.aiplatform.v1.TuningJob.state] is set to
	// `CANCELLED`.
	CancelTuningJob(context.Context, *CancelTuningJobRequest) (*emptypb.Empty, error)
}

// UnimplementedGenAiTuningServiceServer can be embedded to have forward compatible implementations.
type UnimplementedGenAiTuningServiceServer struct {
}

func (*UnimplementedGenAiTuningServiceServer) CreateTuningJob(context.Context, *CreateTuningJobRequest) (*TuningJob, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTuningJob not implemented")
}
func (*UnimplementedGenAiTuningServiceServer) GetTuningJob(context.Context, *GetTuningJobRequest) (*TuningJob, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTuningJob not implemented")
}
func (*UnimplementedGenAiTuningServiceServer) ListTuningJobs(context.Context, *ListTuningJobsRequest) (*ListTuningJobsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTuningJobs not implemented")
}
func (*UnimplementedGenAiTuningServiceServer) CancelTuningJob(context.Context, *CancelTuningJobRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelTuningJob not implemented")
}

func RegisterGenAiTuningServiceServer(s *grpc.Server, srv GenAiTuningServiceServer) {
	s.RegisterService(&_GenAiTuningService_serviceDesc, srv)
}

func _GenAiTuningService_CreateTuningJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTuningJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GenAiTuningServiceServer).CreateTuningJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.aiplatform.v1.GenAiTuningService/CreateTuningJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GenAiTuningServiceServer).CreateTuningJob(ctx, req.(*CreateTuningJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GenAiTuningService_GetTuningJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTuningJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GenAiTuningServiceServer).GetTuningJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.aiplatform.v1.GenAiTuningService/GetTuningJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GenAiTuningServiceServer).GetTuningJob(ctx, req.(*GetTuningJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GenAiTuningService_ListTuningJobs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTuningJobsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GenAiTuningServiceServer).ListTuningJobs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.aiplatform.v1.GenAiTuningService/ListTuningJobs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GenAiTuningServiceServer).ListTuningJobs(ctx, req.(*ListTuningJobsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GenAiTuningService_CancelTuningJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelTuningJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GenAiTuningServiceServer).CancelTuningJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.aiplatform.v1.GenAiTuningService/CancelTuningJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GenAiTuningServiceServer).CancelTuningJob(ctx, req.(*CancelTuningJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GenAiTuningService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.cloud.aiplatform.v1.GenAiTuningService",
	HandlerType: (*GenAiTuningServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTuningJob",
			Handler:    _GenAiTuningService_CreateTuningJob_Handler,
		},
		{
			MethodName: "GetTuningJob",
			Handler:    _GenAiTuningService_GetTuningJob_Handler,
		},
		{
			MethodName: "ListTuningJobs",
			Handler:    _GenAiTuningService_ListTuningJobs_Handler,
		},
		{
			MethodName: "CancelTuningJob",
			Handler:    _GenAiTuningService_CancelTuningJob_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/cloud/aiplatform/v1/genai_tuning_service.proto",
}