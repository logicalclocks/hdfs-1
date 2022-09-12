//*
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//*
// These .proto interfaces are private and stable.
// Please see http://wiki.apache.org/hadoop/Compatibility
// for what changes are allowed for a *stable* .proto interface.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.2
// source: RefreshUserMappingsProtocol.proto

package hadoop_common

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

//*
//  Refresh user to group mappings request.
type RefreshUserToGroupsMappingsRequestProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RefreshUserToGroupsMappingsRequestProto) Reset() {
	*x = RefreshUserToGroupsMappingsRequestProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_RefreshUserMappingsProtocol_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RefreshUserToGroupsMappingsRequestProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefreshUserToGroupsMappingsRequestProto) ProtoMessage() {}

func (x *RefreshUserToGroupsMappingsRequestProto) ProtoReflect() protoreflect.Message {
	mi := &file_RefreshUserMappingsProtocol_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefreshUserToGroupsMappingsRequestProto.ProtoReflect.Descriptor instead.
func (*RefreshUserToGroupsMappingsRequestProto) Descriptor() ([]byte, []int) {
	return file_RefreshUserMappingsProtocol_proto_rawDescGZIP(), []int{0}
}

//*
// void response
type RefreshUserToGroupsMappingsResponseProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RefreshUserToGroupsMappingsResponseProto) Reset() {
	*x = RefreshUserToGroupsMappingsResponseProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_RefreshUserMappingsProtocol_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RefreshUserToGroupsMappingsResponseProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefreshUserToGroupsMappingsResponseProto) ProtoMessage() {}

func (x *RefreshUserToGroupsMappingsResponseProto) ProtoReflect() protoreflect.Message {
	mi := &file_RefreshUserMappingsProtocol_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefreshUserToGroupsMappingsResponseProto.ProtoReflect.Descriptor instead.
func (*RefreshUserToGroupsMappingsResponseProto) Descriptor() ([]byte, []int) {
	return file_RefreshUserMappingsProtocol_proto_rawDescGZIP(), []int{1}
}

//*
// Refresh superuser configuration request.
type RefreshSuperUserGroupsConfigurationRequestProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RefreshSuperUserGroupsConfigurationRequestProto) Reset() {
	*x = RefreshSuperUserGroupsConfigurationRequestProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_RefreshUserMappingsProtocol_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RefreshSuperUserGroupsConfigurationRequestProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefreshSuperUserGroupsConfigurationRequestProto) ProtoMessage() {}

func (x *RefreshSuperUserGroupsConfigurationRequestProto) ProtoReflect() protoreflect.Message {
	mi := &file_RefreshUserMappingsProtocol_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefreshSuperUserGroupsConfigurationRequestProto.ProtoReflect.Descriptor instead.
func (*RefreshSuperUserGroupsConfigurationRequestProto) Descriptor() ([]byte, []int) {
	return file_RefreshUserMappingsProtocol_proto_rawDescGZIP(), []int{2}
}

//*
// void response
type RefreshSuperUserGroupsConfigurationResponseProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RefreshSuperUserGroupsConfigurationResponseProto) Reset() {
	*x = RefreshSuperUserGroupsConfigurationResponseProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_RefreshUserMappingsProtocol_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RefreshSuperUserGroupsConfigurationResponseProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefreshSuperUserGroupsConfigurationResponseProto) ProtoMessage() {}

func (x *RefreshSuperUserGroupsConfigurationResponseProto) ProtoReflect() protoreflect.Message {
	mi := &file_RefreshUserMappingsProtocol_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefreshSuperUserGroupsConfigurationResponseProto.ProtoReflect.Descriptor instead.
func (*RefreshSuperUserGroupsConfigurationResponseProto) Descriptor() ([]byte, []int) {
	return file_RefreshUserMappingsProtocol_proto_rawDescGZIP(), []int{3}
}

var File_RefreshUserMappingsProtocol_proto protoreflect.FileDescriptor

var file_RefreshUserMappingsProtocol_proto_rawDesc = []byte{
	0x0a, 0x21, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x61, 0x70,
	0x70, 0x69, 0x6e, 0x67, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x68, 0x61, 0x64, 0x6f, 0x6f, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x22, 0x29, 0x0a, 0x27, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x55, 0x73, 0x65,
	0x72, 0x54, 0x6f, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2a, 0x0a,
	0x28, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x55, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x73, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x31, 0x0a, 0x2f, 0x52, 0x65, 0x66,
	0x72, 0x65, 0x73, 0x68, 0x53, 0x75, 0x70, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x32, 0x0a, 0x30,
	0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x53, 0x75, 0x70, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x32, 0xde, 0x02, 0x0a, 0x22, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x55, 0x73, 0x65, 0x72,
	0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x8e, 0x01, 0x0a, 0x1b, 0x72, 0x65, 0x66, 0x72,
	0x65, 0x73, 0x68, 0x55, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x4d,
	0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x36, 0x2e, 0x68, 0x61, 0x64, 0x6f, 0x6f, 0x70,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x55,
	0x73, 0x65, 0x72, 0x54, 0x6f, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x4d, 0x61, 0x70, 0x70, 0x69,
	0x6e, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x37, 0x2e, 0x68, 0x61, 0x64, 0x6f, 0x6f, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x55, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x73, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0xa6, 0x01, 0x0a, 0x23, 0x72, 0x65, 0x66,
	0x72, 0x65, 0x73, 0x68, 0x53, 0x75, 0x70, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x3e, 0x2e, 0x68, 0x61, 0x64, 0x6f, 0x6f, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x53, 0x75, 0x70, 0x65, 0x72, 0x55, 0x73, 0x65,
	0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x3f, 0x2e, 0x68, 0x61, 0x64, 0x6f, 0x6f, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x53, 0x75, 0x70, 0x65, 0x72, 0x55, 0x73, 0x65,
	0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x42, 0x89, 0x01, 0x0a, 0x20, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65,
	0x2e, 0x68, 0x61, 0x64, 0x6f, 0x6f, 0x70, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x42, 0x21, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x55,
	0x73, 0x65, 0x72, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x6c, 0x69, 0x6e, 0x6d, 0x61, 0x72, 0x63, 0x2f,
	0x68, 0x64, 0x66, 0x73, 0x2f, 0x76, 0x32, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x68, 0x61, 0x64, 0x6f, 0x6f, 0x70,
	0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0xa0, 0x01, 0x01,
}

var (
	file_RefreshUserMappingsProtocol_proto_rawDescOnce sync.Once
	file_RefreshUserMappingsProtocol_proto_rawDescData = file_RefreshUserMappingsProtocol_proto_rawDesc
)

func file_RefreshUserMappingsProtocol_proto_rawDescGZIP() []byte {
	file_RefreshUserMappingsProtocol_proto_rawDescOnce.Do(func() {
		file_RefreshUserMappingsProtocol_proto_rawDescData = protoimpl.X.CompressGZIP(file_RefreshUserMappingsProtocol_proto_rawDescData)
	})
	return file_RefreshUserMappingsProtocol_proto_rawDescData
}

var file_RefreshUserMappingsProtocol_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_RefreshUserMappingsProtocol_proto_goTypes = []interface{}{
	(*RefreshUserToGroupsMappingsRequestProto)(nil),          // 0: hadoop.common.RefreshUserToGroupsMappingsRequestProto
	(*RefreshUserToGroupsMappingsResponseProto)(nil),         // 1: hadoop.common.RefreshUserToGroupsMappingsResponseProto
	(*RefreshSuperUserGroupsConfigurationRequestProto)(nil),  // 2: hadoop.common.RefreshSuperUserGroupsConfigurationRequestProto
	(*RefreshSuperUserGroupsConfigurationResponseProto)(nil), // 3: hadoop.common.RefreshSuperUserGroupsConfigurationResponseProto
}
var file_RefreshUserMappingsProtocol_proto_depIdxs = []int32{
	0, // 0: hadoop.common.RefreshUserMappingsProtocolService.refreshUserToGroupsMappings:input_type -> hadoop.common.RefreshUserToGroupsMappingsRequestProto
	2, // 1: hadoop.common.RefreshUserMappingsProtocolService.refreshSuperUserGroupsConfiguration:input_type -> hadoop.common.RefreshSuperUserGroupsConfigurationRequestProto
	1, // 2: hadoop.common.RefreshUserMappingsProtocolService.refreshUserToGroupsMappings:output_type -> hadoop.common.RefreshUserToGroupsMappingsResponseProto
	3, // 3: hadoop.common.RefreshUserMappingsProtocolService.refreshSuperUserGroupsConfiguration:output_type -> hadoop.common.RefreshSuperUserGroupsConfigurationResponseProto
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_RefreshUserMappingsProtocol_proto_init() }
func file_RefreshUserMappingsProtocol_proto_init() {
	if File_RefreshUserMappingsProtocol_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_RefreshUserMappingsProtocol_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RefreshUserToGroupsMappingsRequestProto); i {
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
		file_RefreshUserMappingsProtocol_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RefreshUserToGroupsMappingsResponseProto); i {
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
		file_RefreshUserMappingsProtocol_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RefreshSuperUserGroupsConfigurationRequestProto); i {
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
		file_RefreshUserMappingsProtocol_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RefreshSuperUserGroupsConfigurationResponseProto); i {
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
			RawDescriptor: file_RefreshUserMappingsProtocol_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_RefreshUserMappingsProtocol_proto_goTypes,
		DependencyIndexes: file_RefreshUserMappingsProtocol_proto_depIdxs,
		MessageInfos:      file_RefreshUserMappingsProtocol_proto_msgTypes,
	}.Build()
	File_RefreshUserMappingsProtocol_proto = out.File
	file_RefreshUserMappingsProtocol_proto_rawDesc = nil
	file_RefreshUserMappingsProtocol_proto_goTypes = nil
	file_RefreshUserMappingsProtocol_proto_depIdxs = nil
}
