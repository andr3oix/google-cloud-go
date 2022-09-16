// Copyright 2020 Google LLC
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
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.5
// source: google/cloud/gaming/v1beta/game_server_configs_service.proto

package gamingpb

import (
	context "context"
	reflect "reflect"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	longrunning "google.golang.org/genproto/googleapis/longrunning"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_google_cloud_gaming_v1beta_game_server_configs_service_proto protoreflect.FileDescriptor

var file_google_cloud_gaming_v1beta_game_server_configs_service_proto_rawDesc = []byte{
	0x0a, 0x3c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67,
	0x61, 0x6d, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2f, 0x67, 0x61, 0x6d,
	0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x61, 0x6d,
	0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x34, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x67, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2f, 0x67, 0x61,
	0x6d, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x6c, 0x6f, 0x6e, 0x67, 0x72, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x2f, 0x6f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xc4, 0x08, 0x0a,
	0x18, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xe6, 0x01, 0x0a, 0x15, 0x4c, 0x69,
	0x73, 0x74, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x73, 0x12, 0x38, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x67, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x39, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x61, 0x6d,
	0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x47,
	0x61, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x58, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x49,
	0x12, 0x47, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2f, 0x7b, 0x70, 0x61, 0x72, 0x65, 0x6e,
	0x74, 0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x2a, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x2a,
	0x7d, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0xda, 0x41, 0x06, 0x70, 0x61, 0x72, 0x65,
	0x6e, 0x74, 0x12, 0xd3, 0x01, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x36, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x61, 0x6d, 0x69, 0x6e, 0x67,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x61, 0x6d, 0x65, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x67, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2e,
	0x47, 0x61, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x22, 0x56, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x49, 0x12, 0x47, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73,
	0x2f, 0x2a, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x2a, 0x2f, 0x67,
	0x61, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x2f, 0x2a,
	0x7d, 0xda, 0x41, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x9c, 0x02, 0x0a, 0x16, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x12, 0x39, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x67, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6c, 0x6f, 0x6e, 0x67, 0x72, 0x75, 0x6e, 0x6e,
	0x69, 0x6e, 0x67, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xa7, 0x01,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x5d, 0x22, 0x47, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2f,
	0x7b, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73,
	0x2f, 0x2a, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x2a, 0x2f, 0x67,
	0x61, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x2f, 0x2a, 0x7d, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x3a,
	0x12, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0xda, 0x41, 0x19, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x2c, 0x67, 0x61, 0x6d,
	0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0xca,
	0x41, 0x25, 0x0a, 0x10, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x11, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0xf8, 0x01, 0x0a, 0x16, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x12, 0x39, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x67, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6c, 0x6f, 0x6e, 0x67, 0x72, 0x75, 0x6e, 0x6e, 0x69,
	0x6e, 0x67, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x83, 0x01, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x49, 0x2a, 0x47, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2f, 0x7b,
	0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a, 0x2f,
	0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x2a, 0x2f, 0x67, 0x61, 0x6d, 0x65,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x2f, 0x2a, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x2f, 0x2a, 0x7d, 0xda, 0x41,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0xca, 0x41, 0x2a, 0x0a, 0x15, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12,
	0x11, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x1a, 0x4f, 0xca, 0x41, 0x1b, 0x67, 0x61, 0x6d, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63,
	0x6f, 0x6d, 0xd2, 0x41, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x77, 0x77, 0x77,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x61, 0x75, 0x74, 0x68, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x70, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x42, 0x81, 0x01, 0x0a, 0x1e, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x2e,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x50, 0x01, 0x5a, 0x40, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x3b, 0x67, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0xca, 0x02, 0x1a, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x47, 0x61, 0x6d, 0x69, 0x6e, 0x67,
	0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_google_cloud_gaming_v1beta_game_server_configs_service_proto_goTypes = []interface{}{
	(*ListGameServerConfigsRequest)(nil),  // 0: google.cloud.gaming.v1beta.ListGameServerConfigsRequest
	(*GetGameServerConfigRequest)(nil),    // 1: google.cloud.gaming.v1beta.GetGameServerConfigRequest
	(*CreateGameServerConfigRequest)(nil), // 2: google.cloud.gaming.v1beta.CreateGameServerConfigRequest
	(*DeleteGameServerConfigRequest)(nil), // 3: google.cloud.gaming.v1beta.DeleteGameServerConfigRequest
	(*ListGameServerConfigsResponse)(nil), // 4: google.cloud.gaming.v1beta.ListGameServerConfigsResponse
	(*GameServerConfig)(nil),              // 5: google.cloud.gaming.v1beta.GameServerConfig
	(*longrunning.Operation)(nil),         // 6: google.longrunning.Operation
}
var file_google_cloud_gaming_v1beta_game_server_configs_service_proto_depIdxs = []int32{
	0, // 0: google.cloud.gaming.v1beta.GameServerConfigsService.ListGameServerConfigs:input_type -> google.cloud.gaming.v1beta.ListGameServerConfigsRequest
	1, // 1: google.cloud.gaming.v1beta.GameServerConfigsService.GetGameServerConfig:input_type -> google.cloud.gaming.v1beta.GetGameServerConfigRequest
	2, // 2: google.cloud.gaming.v1beta.GameServerConfigsService.CreateGameServerConfig:input_type -> google.cloud.gaming.v1beta.CreateGameServerConfigRequest
	3, // 3: google.cloud.gaming.v1beta.GameServerConfigsService.DeleteGameServerConfig:input_type -> google.cloud.gaming.v1beta.DeleteGameServerConfigRequest
	4, // 4: google.cloud.gaming.v1beta.GameServerConfigsService.ListGameServerConfigs:output_type -> google.cloud.gaming.v1beta.ListGameServerConfigsResponse
	5, // 5: google.cloud.gaming.v1beta.GameServerConfigsService.GetGameServerConfig:output_type -> google.cloud.gaming.v1beta.GameServerConfig
	6, // 6: google.cloud.gaming.v1beta.GameServerConfigsService.CreateGameServerConfig:output_type -> google.longrunning.Operation
	6, // 7: google.cloud.gaming.v1beta.GameServerConfigsService.DeleteGameServerConfig:output_type -> google.longrunning.Operation
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_cloud_gaming_v1beta_game_server_configs_service_proto_init() }
func file_google_cloud_gaming_v1beta_game_server_configs_service_proto_init() {
	if File_google_cloud_gaming_v1beta_game_server_configs_service_proto != nil {
		return
	}
	file_google_cloud_gaming_v1beta_game_server_configs_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_gaming_v1beta_game_server_configs_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_cloud_gaming_v1beta_game_server_configs_service_proto_goTypes,
		DependencyIndexes: file_google_cloud_gaming_v1beta_game_server_configs_service_proto_depIdxs,
	}.Build()
	File_google_cloud_gaming_v1beta_game_server_configs_service_proto = out.File
	file_google_cloud_gaming_v1beta_game_server_configs_service_proto_rawDesc = nil
	file_google_cloud_gaming_v1beta_game_server_configs_service_proto_goTypes = nil
	file_google_cloud_gaming_v1beta_game_server_configs_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// GameServerConfigsServiceClient is the client API for GameServerConfigsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GameServerConfigsServiceClient interface {
	// Lists game server configs in a given project, location, and game server
	// deployment.
	ListGameServerConfigs(ctx context.Context, in *ListGameServerConfigsRequest, opts ...grpc.CallOption) (*ListGameServerConfigsResponse, error)
	// Gets details of a single game server config.
	GetGameServerConfig(ctx context.Context, in *GetGameServerConfigRequest, opts ...grpc.CallOption) (*GameServerConfig, error)
	// Creates a new game server config in a given project, location, and game
	// server deployment. Game server configs are immutable, and are not applied
	// until referenced in the game server deployment rollout resource.
	CreateGameServerConfig(ctx context.Context, in *CreateGameServerConfigRequest, opts ...grpc.CallOption) (*longrunning.Operation, error)
	// Deletes a single game server config. The deletion will fail if the game
	// server config is referenced in a game server deployment rollout.
	DeleteGameServerConfig(ctx context.Context, in *DeleteGameServerConfigRequest, opts ...grpc.CallOption) (*longrunning.Operation, error)
}

type gameServerConfigsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGameServerConfigsServiceClient(cc grpc.ClientConnInterface) GameServerConfigsServiceClient {
	return &gameServerConfigsServiceClient{cc}
}

func (c *gameServerConfigsServiceClient) ListGameServerConfigs(ctx context.Context, in *ListGameServerConfigsRequest, opts ...grpc.CallOption) (*ListGameServerConfigsResponse, error) {
	out := new(ListGameServerConfigsResponse)
	err := c.cc.Invoke(ctx, "/google.cloud.gaming.v1beta.GameServerConfigsService/ListGameServerConfigs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameServerConfigsServiceClient) GetGameServerConfig(ctx context.Context, in *GetGameServerConfigRequest, opts ...grpc.CallOption) (*GameServerConfig, error) {
	out := new(GameServerConfig)
	err := c.cc.Invoke(ctx, "/google.cloud.gaming.v1beta.GameServerConfigsService/GetGameServerConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameServerConfigsServiceClient) CreateGameServerConfig(ctx context.Context, in *CreateGameServerConfigRequest, opts ...grpc.CallOption) (*longrunning.Operation, error) {
	out := new(longrunning.Operation)
	err := c.cc.Invoke(ctx, "/google.cloud.gaming.v1beta.GameServerConfigsService/CreateGameServerConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameServerConfigsServiceClient) DeleteGameServerConfig(ctx context.Context, in *DeleteGameServerConfigRequest, opts ...grpc.CallOption) (*longrunning.Operation, error) {
	out := new(longrunning.Operation)
	err := c.cc.Invoke(ctx, "/google.cloud.gaming.v1beta.GameServerConfigsService/DeleteGameServerConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GameServerConfigsServiceServer is the server API for GameServerConfigsService service.
type GameServerConfigsServiceServer interface {
	// Lists game server configs in a given project, location, and game server
	// deployment.
	ListGameServerConfigs(context.Context, *ListGameServerConfigsRequest) (*ListGameServerConfigsResponse, error)
	// Gets details of a single game server config.
	GetGameServerConfig(context.Context, *GetGameServerConfigRequest) (*GameServerConfig, error)
	// Creates a new game server config in a given project, location, and game
	// server deployment. Game server configs are immutable, and are not applied
	// until referenced in the game server deployment rollout resource.
	CreateGameServerConfig(context.Context, *CreateGameServerConfigRequest) (*longrunning.Operation, error)
	// Deletes a single game server config. The deletion will fail if the game
	// server config is referenced in a game server deployment rollout.
	DeleteGameServerConfig(context.Context, *DeleteGameServerConfigRequest) (*longrunning.Operation, error)
}

// UnimplementedGameServerConfigsServiceServer can be embedded to have forward compatible implementations.
type UnimplementedGameServerConfigsServiceServer struct {
}

func (*UnimplementedGameServerConfigsServiceServer) ListGameServerConfigs(context.Context, *ListGameServerConfigsRequest) (*ListGameServerConfigsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListGameServerConfigs not implemented")
}
func (*UnimplementedGameServerConfigsServiceServer) GetGameServerConfig(context.Context, *GetGameServerConfigRequest) (*GameServerConfig, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGameServerConfig not implemented")
}
func (*UnimplementedGameServerConfigsServiceServer) CreateGameServerConfig(context.Context, *CreateGameServerConfigRequest) (*longrunning.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGameServerConfig not implemented")
}
func (*UnimplementedGameServerConfigsServiceServer) DeleteGameServerConfig(context.Context, *DeleteGameServerConfigRequest) (*longrunning.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteGameServerConfig not implemented")
}

func RegisterGameServerConfigsServiceServer(s *grpc.Server, srv GameServerConfigsServiceServer) {
	s.RegisterService(&_GameServerConfigsService_serviceDesc, srv)
}

func _GameServerConfigsService_ListGameServerConfigs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListGameServerConfigsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServerConfigsServiceServer).ListGameServerConfigs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.gaming.v1beta.GameServerConfigsService/ListGameServerConfigs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServerConfigsServiceServer).ListGameServerConfigs(ctx, req.(*ListGameServerConfigsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GameServerConfigsService_GetGameServerConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGameServerConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServerConfigsServiceServer).GetGameServerConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.gaming.v1beta.GameServerConfigsService/GetGameServerConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServerConfigsServiceServer).GetGameServerConfig(ctx, req.(*GetGameServerConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GameServerConfigsService_CreateGameServerConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGameServerConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServerConfigsServiceServer).CreateGameServerConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.gaming.v1beta.GameServerConfigsService/CreateGameServerConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServerConfigsServiceServer).CreateGameServerConfig(ctx, req.(*CreateGameServerConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GameServerConfigsService_DeleteGameServerConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteGameServerConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServerConfigsServiceServer).DeleteGameServerConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.gaming.v1beta.GameServerConfigsService/DeleteGameServerConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServerConfigsServiceServer).DeleteGameServerConfig(ctx, req.(*DeleteGameServerConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GameServerConfigsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.cloud.gaming.v1beta.GameServerConfigsService",
	HandlerType: (*GameServerConfigsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListGameServerConfigs",
			Handler:    _GameServerConfigsService_ListGameServerConfigs_Handler,
		},
		{
			MethodName: "GetGameServerConfig",
			Handler:    _GameServerConfigsService_GetGameServerConfig_Handler,
		},
		{
			MethodName: "CreateGameServerConfig",
			Handler:    _GameServerConfigsService_CreateGameServerConfig_Handler,
		},
		{
			MethodName: "DeleteGameServerConfig",
			Handler:    _GameServerConfigsService_DeleteGameServerConfig_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/cloud/gaming/v1beta/game_server_configs_service.proto",
}
