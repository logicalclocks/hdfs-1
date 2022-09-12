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
<<<<<<< HEAD
// 	protoc-gen-go v1.25.0
// 	protoc        (unknown)
=======
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
>>>>>>> upstream-colinmarc/master
// source: IpcConnectionContext.proto

package hadoop_common

import (
<<<<<<< HEAD
	proto "github.com/golang/protobuf/proto"
=======
>>>>>>> upstream-colinmarc/master
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

<<<<<<< HEAD
// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

=======
>>>>>>> upstream-colinmarc/master
//*
// Spec for UserInformationProto is specified in ProtoUtil#makeIpcConnectionContext
type UserInformationProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EffectiveUser *string `protobuf:"bytes,1,opt,name=effectiveUser" json:"effectiveUser,omitempty"`
	RealUser      *string `protobuf:"bytes,2,opt,name=realUser" json:"realUser,omitempty"`
}

func (x *UserInformationProto) Reset() {
	*x = UserInformationProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_IpcConnectionContext_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInformationProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInformationProto) ProtoMessage() {}

func (x *UserInformationProto) ProtoReflect() protoreflect.Message {
	mi := &file_IpcConnectionContext_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInformationProto.ProtoReflect.Descriptor instead.
func (*UserInformationProto) Descriptor() ([]byte, []int) {
	return file_IpcConnectionContext_proto_rawDescGZIP(), []int{0}
}

func (x *UserInformationProto) GetEffectiveUser() string {
	if x != nil && x.EffectiveUser != nil {
		return *x.EffectiveUser
	}
	return ""
}

func (x *UserInformationProto) GetRealUser() string {
	if x != nil && x.RealUser != nil {
		return *x.RealUser
	}
	return ""
}

//*
// The connection context is sent as part of the connection establishment.
// It establishes the context for ALL Rpc calls within the connection.
type IpcConnectionContextProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// UserInfo beyond what is determined as part of security handshake
	// at connection time (kerberos, tokens etc).
	UserInfo *UserInformationProto `protobuf:"bytes,2,opt,name=userInfo" json:"userInfo,omitempty"`
	// Protocol name for next rpc layer.
	// The client created a proxy with this protocol name
	Protocol *string `protobuf:"bytes,3,opt,name=protocol" json:"protocol,omitempty"`
}

func (x *IpcConnectionContextProto) Reset() {
	*x = IpcConnectionContextProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_IpcConnectionContext_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IpcConnectionContextProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IpcConnectionContextProto) ProtoMessage() {}

func (x *IpcConnectionContextProto) ProtoReflect() protoreflect.Message {
	mi := &file_IpcConnectionContext_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IpcConnectionContextProto.ProtoReflect.Descriptor instead.
func (*IpcConnectionContextProto) Descriptor() ([]byte, []int) {
	return file_IpcConnectionContext_proto_rawDescGZIP(), []int{1}
}

func (x *IpcConnectionContextProto) GetUserInfo() *UserInformationProto {
	if x != nil {
		return x.UserInfo
	}
	return nil
}

func (x *IpcConnectionContextProto) GetProtocol() string {
	if x != nil && x.Protocol != nil {
		return *x.Protocol
	}
	return ""
}

var File_IpcConnectionContext_proto protoreflect.FileDescriptor

var file_IpcConnectionContext_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x49, 0x70, 0x63, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x68, 0x61,
<<<<<<< HEAD
	0x64, 0x6f, 0x6f, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x22, 0x3f, 0x0a, 0x14, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x15, 0x0a, 0x0d, 0x65, 0x66, 0x66, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x12, 0x10, 0x0a, 0x08, 0x72, 0x65,
	0x61, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x22, 0x64, 0x0a, 0x19,
	0x49, 0x70, 0x63, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x78, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x35, 0x0a, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x68, 0x61,
	0x64, 0x6f, 0x6f, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x10, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x3f, 0x0a, 0x1e, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65,
	0x2e, 0x68, 0x61, 0x64, 0x6f, 0x6f, 0x70, 0x2e, 0x69, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x42, 0x1a, 0x49, 0x70, 0x63, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0xa0, 0x01, 0x01,
=======
	0x64, 0x6f, 0x6f, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x22, 0x58, 0x0a, 0x14, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x24, 0x0a, 0x0d, 0x65, 0x66, 0x66, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x65, 0x66, 0x66, 0x65,
	0x63, 0x74, 0x69, 0x76, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x61,
	0x6c, 0x55, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x61,
	0x6c, 0x55, 0x73, 0x65, 0x72, 0x22, 0x78, 0x0a, 0x19, 0x49, 0x70, 0x63, 0x43, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x3f, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x68, 0x61, 0x64, 0x6f, 0x6f, 0x70, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x42,
	0x7d, 0x0a, 0x1e, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x68, 0x61,
	0x64, 0x6f, 0x6f, 0x70, 0x2e, 0x69, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x42, 0x1a, 0x49, 0x70, 0x63, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x5a, 0x3c, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x6c, 0x69, 0x6e, 0x6d,
	0x61, 0x72, 0x63, 0x2f, 0x68, 0x64, 0x66, 0x73, 0x2f, 0x76, 0x32, 0x2f, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x68, 0x61,
	0x64, 0x6f, 0x6f, 0x70, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0xa0, 0x01, 0x01,
>>>>>>> upstream-colinmarc/master
}

var (
	file_IpcConnectionContext_proto_rawDescOnce sync.Once
	file_IpcConnectionContext_proto_rawDescData = file_IpcConnectionContext_proto_rawDesc
)

func file_IpcConnectionContext_proto_rawDescGZIP() []byte {
	file_IpcConnectionContext_proto_rawDescOnce.Do(func() {
		file_IpcConnectionContext_proto_rawDescData = protoimpl.X.CompressGZIP(file_IpcConnectionContext_proto_rawDescData)
	})
	return file_IpcConnectionContext_proto_rawDescData
}

var file_IpcConnectionContext_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_IpcConnectionContext_proto_goTypes = []interface{}{
	(*UserInformationProto)(nil),      // 0: hadoop.common.UserInformationProto
	(*IpcConnectionContextProto)(nil), // 1: hadoop.common.IpcConnectionContextProto
}
var file_IpcConnectionContext_proto_depIdxs = []int32{
	0, // 0: hadoop.common.IpcConnectionContextProto.userInfo:type_name -> hadoop.common.UserInformationProto
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_IpcConnectionContext_proto_init() }
func file_IpcConnectionContext_proto_init() {
	if File_IpcConnectionContext_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_IpcConnectionContext_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserInformationProto); i {
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
		file_IpcConnectionContext_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IpcConnectionContextProto); i {
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
			RawDescriptor: file_IpcConnectionContext_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_IpcConnectionContext_proto_goTypes,
		DependencyIndexes: file_IpcConnectionContext_proto_depIdxs,
		MessageInfos:      file_IpcConnectionContext_proto_msgTypes,
	}.Build()
	File_IpcConnectionContext_proto = out.File
	file_IpcConnectionContext_proto_rawDesc = nil
	file_IpcConnectionContext_proto_goTypes = nil
	file_IpcConnectionContext_proto_depIdxs = nil
}
