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
// 	protoc-gen-go v1.28.1
// 	protoc        v5.28.2
// source: TraceAdmin.proto

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

type ListSpanReceiversRequestProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListSpanReceiversRequestProto) Reset() {
	*x = ListSpanReceiversRequestProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_TraceAdmin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSpanReceiversRequestProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSpanReceiversRequestProto) ProtoMessage() {}

func (x *ListSpanReceiversRequestProto) ProtoReflect() protoreflect.Message {
	mi := &file_TraceAdmin_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSpanReceiversRequestProto.ProtoReflect.Descriptor instead.
func (*ListSpanReceiversRequestProto) Descriptor() ([]byte, []int) {
	return file_TraceAdmin_proto_rawDescGZIP(), []int{0}
}

type SpanReceiverListInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        *int64  `protobuf:"varint,1,req,name=id" json:"id,omitempty"`
	ClassName *string `protobuf:"bytes,2,req,name=className" json:"className,omitempty"`
}

func (x *SpanReceiverListInfo) Reset() {
	*x = SpanReceiverListInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_TraceAdmin_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SpanReceiverListInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpanReceiverListInfo) ProtoMessage() {}

func (x *SpanReceiverListInfo) ProtoReflect() protoreflect.Message {
	mi := &file_TraceAdmin_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpanReceiverListInfo.ProtoReflect.Descriptor instead.
func (*SpanReceiverListInfo) Descriptor() ([]byte, []int) {
	return file_TraceAdmin_proto_rawDescGZIP(), []int{1}
}

func (x *SpanReceiverListInfo) GetId() int64 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *SpanReceiverListInfo) GetClassName() string {
	if x != nil && x.ClassName != nil {
		return *x.ClassName
	}
	return ""
}

type ListSpanReceiversResponseProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Descriptions []*SpanReceiverListInfo `protobuf:"bytes,1,rep,name=descriptions" json:"descriptions,omitempty"`
}

func (x *ListSpanReceiversResponseProto) Reset() {
	*x = ListSpanReceiversResponseProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_TraceAdmin_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSpanReceiversResponseProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSpanReceiversResponseProto) ProtoMessage() {}

func (x *ListSpanReceiversResponseProto) ProtoReflect() protoreflect.Message {
	mi := &file_TraceAdmin_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSpanReceiversResponseProto.ProtoReflect.Descriptor instead.
func (*ListSpanReceiversResponseProto) Descriptor() ([]byte, []int) {
	return file_TraceAdmin_proto_rawDescGZIP(), []int{2}
}

func (x *ListSpanReceiversResponseProto) GetDescriptions() []*SpanReceiverListInfo {
	if x != nil {
		return x.Descriptions
	}
	return nil
}

type ConfigPair struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   *string `protobuf:"bytes,1,req,name=key" json:"key,omitempty"`
	Value *string `protobuf:"bytes,2,req,name=value" json:"value,omitempty"`
}

func (x *ConfigPair) Reset() {
	*x = ConfigPair{}
	if protoimpl.UnsafeEnabled {
		mi := &file_TraceAdmin_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigPair) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigPair) ProtoMessage() {}

func (x *ConfigPair) ProtoReflect() protoreflect.Message {
	mi := &file_TraceAdmin_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigPair.ProtoReflect.Descriptor instead.
func (*ConfigPair) Descriptor() ([]byte, []int) {
	return file_TraceAdmin_proto_rawDescGZIP(), []int{3}
}

func (x *ConfigPair) GetKey() string {
	if x != nil && x.Key != nil {
		return *x.Key
	}
	return ""
}

func (x *ConfigPair) GetValue() string {
	if x != nil && x.Value != nil {
		return *x.Value
	}
	return ""
}

type AddSpanReceiverRequestProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClassName *string       `protobuf:"bytes,1,req,name=className" json:"className,omitempty"`
	Config    []*ConfigPair `protobuf:"bytes,2,rep,name=config" json:"config,omitempty"`
}

func (x *AddSpanReceiverRequestProto) Reset() {
	*x = AddSpanReceiverRequestProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_TraceAdmin_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddSpanReceiverRequestProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddSpanReceiverRequestProto) ProtoMessage() {}

func (x *AddSpanReceiverRequestProto) ProtoReflect() protoreflect.Message {
	mi := &file_TraceAdmin_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddSpanReceiverRequestProto.ProtoReflect.Descriptor instead.
func (*AddSpanReceiverRequestProto) Descriptor() ([]byte, []int) {
	return file_TraceAdmin_proto_rawDescGZIP(), []int{4}
}

func (x *AddSpanReceiverRequestProto) GetClassName() string {
	if x != nil && x.ClassName != nil {
		return *x.ClassName
	}
	return ""
}

func (x *AddSpanReceiverRequestProto) GetConfig() []*ConfigPair {
	if x != nil {
		return x.Config
	}
	return nil
}

type AddSpanReceiverResponseProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id *int64 `protobuf:"varint,1,req,name=id" json:"id,omitempty"`
}

func (x *AddSpanReceiverResponseProto) Reset() {
	*x = AddSpanReceiverResponseProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_TraceAdmin_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddSpanReceiverResponseProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddSpanReceiverResponseProto) ProtoMessage() {}

func (x *AddSpanReceiverResponseProto) ProtoReflect() protoreflect.Message {
	mi := &file_TraceAdmin_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddSpanReceiverResponseProto.ProtoReflect.Descriptor instead.
func (*AddSpanReceiverResponseProto) Descriptor() ([]byte, []int) {
	return file_TraceAdmin_proto_rawDescGZIP(), []int{5}
}

func (x *AddSpanReceiverResponseProto) GetId() int64 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

type RemoveSpanReceiverRequestProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id *int64 `protobuf:"varint,1,req,name=id" json:"id,omitempty"`
}

func (x *RemoveSpanReceiverRequestProto) Reset() {
	*x = RemoveSpanReceiverRequestProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_TraceAdmin_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveSpanReceiverRequestProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveSpanReceiverRequestProto) ProtoMessage() {}

func (x *RemoveSpanReceiverRequestProto) ProtoReflect() protoreflect.Message {
	mi := &file_TraceAdmin_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveSpanReceiverRequestProto.ProtoReflect.Descriptor instead.
func (*RemoveSpanReceiverRequestProto) Descriptor() ([]byte, []int) {
	return file_TraceAdmin_proto_rawDescGZIP(), []int{6}
}

func (x *RemoveSpanReceiverRequestProto) GetId() int64 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

type RemoveSpanReceiverResponseProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RemoveSpanReceiverResponseProto) Reset() {
	*x = RemoveSpanReceiverResponseProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_TraceAdmin_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveSpanReceiverResponseProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveSpanReceiverResponseProto) ProtoMessage() {}

func (x *RemoveSpanReceiverResponseProto) ProtoReflect() protoreflect.Message {
	mi := &file_TraceAdmin_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveSpanReceiverResponseProto.ProtoReflect.Descriptor instead.
func (*RemoveSpanReceiverResponseProto) Descriptor() ([]byte, []int) {
	return file_TraceAdmin_proto_rawDescGZIP(), []int{7}
}

var File_TraceAdmin_proto protoreflect.FileDescriptor

var file_TraceAdmin_proto_rawDesc = []byte{
	0x0a, 0x10, 0x54, 0x72, 0x61, 0x63, 0x65, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0d, 0x68, 0x61, 0x64, 0x6f, 0x6f, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x22, 0x1f, 0x0a, 0x1d, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x70, 0x61, 0x6e, 0x52, 0x65, 0x63,
	0x65, 0x69, 0x76, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x44, 0x0a, 0x14, 0x53, 0x70, 0x61, 0x6e, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76,
	0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x02, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6c,
	0x61, 0x73, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x02, 0x28, 0x09, 0x52, 0x09, 0x63,
	0x6c, 0x61, 0x73, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x69, 0x0a, 0x1e, 0x4c, 0x69, 0x73, 0x74,
	0x53, 0x70, 0x61, 0x6e, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x47, 0x0a, 0x0c, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x23, 0x2e, 0x68, 0x61, 0x64, 0x6f, 0x6f, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x53, 0x70, 0x61, 0x6e, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x4c, 0x69, 0x73,
	0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0c, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x22, 0x34, 0x0a, 0x0a, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x50, 0x61, 0x69,
	0x72, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x02, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x02,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x6e, 0x0a, 0x1b, 0x41, 0x64, 0x64,
	0x53, 0x70, 0x61, 0x6e, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6c, 0x61, 0x73,
	0x73, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x02, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6c, 0x61,
	0x73, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x31, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x68, 0x61, 0x64, 0x6f, 0x6f, 0x70, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x50, 0x61, 0x69,
	0x72, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x2e, 0x0a, 0x1c, 0x41, 0x64, 0x64,
	0x53, 0x70, 0x61, 0x6e, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x02, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x30, 0x0a, 0x1e, 0x52, 0x65, 0x6d,
	0x6f, 0x76, 0x65, 0x53, 0x70, 0x61, 0x6e, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x02, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x21, 0x0a, 0x1f, 0x52,
	0x65, 0x6d, 0x6f, 0x76, 0x65, 0x53, 0x70, 0x61, 0x6e, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xe6,
	0x02, 0x0a, 0x11, 0x54, 0x72, 0x61, 0x63, 0x65, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x70, 0x0a, 0x11, 0x6c, 0x69, 0x73, 0x74, 0x53, 0x70, 0x61, 0x6e,
	0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x73, 0x12, 0x2c, 0x2e, 0x68, 0x61, 0x64, 0x6f,
	0x6f, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x70,
	0x61, 0x6e, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2d, 0x2e, 0x68, 0x61, 0x64, 0x6f, 0x6f, 0x70,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x70, 0x61, 0x6e,
	0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x6a, 0x0a, 0x0f, 0x61, 0x64, 0x64, 0x53, 0x70, 0x61,
	0x6e, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x12, 0x2a, 0x2e, 0x68, 0x61, 0x64, 0x6f,
	0x6f, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x41, 0x64, 0x64, 0x53, 0x70, 0x61,
	0x6e, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2b, 0x2e, 0x68, 0x61, 0x64, 0x6f, 0x6f, 0x70, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x41, 0x64, 0x64, 0x53, 0x70, 0x61, 0x6e, 0x52, 0x65, 0x63,
	0x65, 0x69, 0x76, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x73, 0x0a, 0x12, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x53, 0x70, 0x61, 0x6e,
	0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x12, 0x2d, 0x2e, 0x68, 0x61, 0x64, 0x6f, 0x6f,
	0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x53,
	0x70, 0x61, 0x6e, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x2e, 0x68, 0x61, 0x64, 0x6f, 0x6f, 0x70,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x53, 0x70,
	0x61, 0x6e, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x42, 0x6d, 0x0a, 0x19, 0x6f, 0x72, 0x67, 0x2e, 0x61,
	0x70, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x68, 0x61, 0x64, 0x6f, 0x6f, 0x70, 0x2e, 0x74, 0x72, 0x61,
	0x63, 0x69, 0x6e, 0x67, 0x42, 0x0c, 0x54, 0x72, 0x61, 0x63, 0x65, 0x41, 0x64, 0x6d, 0x69, 0x6e,
	0x50, 0x42, 0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63,
	0x6f, 0x6c, 0x69, 0x6e, 0x6d, 0x61, 0x72, 0x63, 0x2f, 0x68, 0x64, 0x66, 0x73, 0x2f, 0x76, 0x32,
	0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x2f, 0x68, 0x61, 0x64, 0x6f, 0x6f, 0x70, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x88, 0x01, 0x01, 0xa0, 0x01, 0x01,
}

var (
	file_TraceAdmin_proto_rawDescOnce sync.Once
	file_TraceAdmin_proto_rawDescData = file_TraceAdmin_proto_rawDesc
)

func file_TraceAdmin_proto_rawDescGZIP() []byte {
	file_TraceAdmin_proto_rawDescOnce.Do(func() {
		file_TraceAdmin_proto_rawDescData = protoimpl.X.CompressGZIP(file_TraceAdmin_proto_rawDescData)
	})
	return file_TraceAdmin_proto_rawDescData
}

var file_TraceAdmin_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_TraceAdmin_proto_goTypes = []interface{}{
	(*ListSpanReceiversRequestProto)(nil),   // 0: hadoop.common.ListSpanReceiversRequestProto
	(*SpanReceiverListInfo)(nil),            // 1: hadoop.common.SpanReceiverListInfo
	(*ListSpanReceiversResponseProto)(nil),  // 2: hadoop.common.ListSpanReceiversResponseProto
	(*ConfigPair)(nil),                      // 3: hadoop.common.ConfigPair
	(*AddSpanReceiverRequestProto)(nil),     // 4: hadoop.common.AddSpanReceiverRequestProto
	(*AddSpanReceiverResponseProto)(nil),    // 5: hadoop.common.AddSpanReceiverResponseProto
	(*RemoveSpanReceiverRequestProto)(nil),  // 6: hadoop.common.RemoveSpanReceiverRequestProto
	(*RemoveSpanReceiverResponseProto)(nil), // 7: hadoop.common.RemoveSpanReceiverResponseProto
}
var file_TraceAdmin_proto_depIdxs = []int32{
	1, // 0: hadoop.common.ListSpanReceiversResponseProto.descriptions:type_name -> hadoop.common.SpanReceiverListInfo
	3, // 1: hadoop.common.AddSpanReceiverRequestProto.config:type_name -> hadoop.common.ConfigPair
	0, // 2: hadoop.common.TraceAdminService.listSpanReceivers:input_type -> hadoop.common.ListSpanReceiversRequestProto
	4, // 3: hadoop.common.TraceAdminService.addSpanReceiver:input_type -> hadoop.common.AddSpanReceiverRequestProto
	6, // 4: hadoop.common.TraceAdminService.removeSpanReceiver:input_type -> hadoop.common.RemoveSpanReceiverRequestProto
	2, // 5: hadoop.common.TraceAdminService.listSpanReceivers:output_type -> hadoop.common.ListSpanReceiversResponseProto
	5, // 6: hadoop.common.TraceAdminService.addSpanReceiver:output_type -> hadoop.common.AddSpanReceiverResponseProto
	7, // 7: hadoop.common.TraceAdminService.removeSpanReceiver:output_type -> hadoop.common.RemoveSpanReceiverResponseProto
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_TraceAdmin_proto_init() }
func file_TraceAdmin_proto_init() {
	if File_TraceAdmin_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_TraceAdmin_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSpanReceiversRequestProto); i {
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
		file_TraceAdmin_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SpanReceiverListInfo); i {
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
		file_TraceAdmin_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSpanReceiversResponseProto); i {
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
		file_TraceAdmin_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigPair); i {
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
		file_TraceAdmin_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddSpanReceiverRequestProto); i {
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
		file_TraceAdmin_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddSpanReceiverResponseProto); i {
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
		file_TraceAdmin_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveSpanReceiverRequestProto); i {
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
		file_TraceAdmin_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveSpanReceiverResponseProto); i {
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
			RawDescriptor: file_TraceAdmin_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_TraceAdmin_proto_goTypes,
		DependencyIndexes: file_TraceAdmin_proto_depIdxs,
		MessageInfos:      file_TraceAdmin_proto_msgTypes,
	}.Build()
	File_TraceAdmin_proto = out.File
	file_TraceAdmin_proto_rawDesc = nil
	file_TraceAdmin_proto_goTypes = nil
	file_TraceAdmin_proto_depIdxs = nil
}
