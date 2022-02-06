//
// * Copyright 2021 The Feast Authors
// *
// * Licensed under the Apache License, Version 2.0 (the "License");
// * you may not use this file except in compliance with the License.
// * You may obtain a copy of the License at
// *
// *     https://www.apache.org/licenses/LICENSE-2.0
// *
// * Unless required by applicable law or agreed to in writing, software
// * distributed under the License is distributed on an "AS IS" BASIS,
// * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// * See the License for the specific language governing permissions and
// * limitations under the License.
//

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: feast/core/DynamoDBTable.proto

package core

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

// Represents a DynamoDB table
type DynamoDBTable struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the table
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Region of the table
	Region string `protobuf:"bytes,2,opt,name=region,proto3" json:"region,omitempty"`
}

func (x *DynamoDBTable) Reset() {
	*x = DynamoDBTable{}
	if protoimpl.UnsafeEnabled {
		mi := &file_feast_core_DynamoDBTable_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DynamoDBTable) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DynamoDBTable) ProtoMessage() {}

func (x *DynamoDBTable) ProtoReflect() protoreflect.Message {
	mi := &file_feast_core_DynamoDBTable_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DynamoDBTable.ProtoReflect.Descriptor instead.
func (*DynamoDBTable) Descriptor() ([]byte, []int) {
	return file_feast_core_DynamoDBTable_proto_rawDescGZIP(), []int{0}
}

func (x *DynamoDBTable) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DynamoDBTable) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

var File_feast_core_DynamoDBTable_proto protoreflect.FileDescriptor

var file_feast_core_DynamoDBTable_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x66, 0x65, 0x61, 0x73, 0x74, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x44, 0x79, 0x6e,
	0x61, 0x6d, 0x6f, 0x44, 0x42, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0a, 0x66, 0x65, 0x61, 0x73, 0x74, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x22, 0x3b, 0x0a, 0x0d,
	0x44, 0x79, 0x6e, 0x61, 0x6d, 0x6f, 0x44, 0x42, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x42, 0x57, 0x0a, 0x10, 0x66, 0x65, 0x61,
	0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x42, 0x12, 0x44,
	0x79, 0x6e, 0x61, 0x6d, 0x6f, 0x44, 0x42, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x65,
	0x61, 0x73, 0x74, 0x2d, 0x64, 0x65, 0x76, 0x2f, 0x66, 0x65, 0x61, 0x73, 0x74, 0x2f, 0x67, 0x6f,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x66, 0x65, 0x61, 0x73, 0x74, 0x2f, 0x63, 0x6f,
	0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_feast_core_DynamoDBTable_proto_rawDescOnce sync.Once
	file_feast_core_DynamoDBTable_proto_rawDescData = file_feast_core_DynamoDBTable_proto_rawDesc
)

func file_feast_core_DynamoDBTable_proto_rawDescGZIP() []byte {
	file_feast_core_DynamoDBTable_proto_rawDescOnce.Do(func() {
		file_feast_core_DynamoDBTable_proto_rawDescData = protoimpl.X.CompressGZIP(file_feast_core_DynamoDBTable_proto_rawDescData)
	})
	return file_feast_core_DynamoDBTable_proto_rawDescData
}

var file_feast_core_DynamoDBTable_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_feast_core_DynamoDBTable_proto_goTypes = []interface{}{
	(*DynamoDBTable)(nil), // 0: feast.core.DynamoDBTable
}
var file_feast_core_DynamoDBTable_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_feast_core_DynamoDBTable_proto_init() }
func file_feast_core_DynamoDBTable_proto_init() {
	if File_feast_core_DynamoDBTable_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_feast_core_DynamoDBTable_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DynamoDBTable); i {
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
			RawDescriptor: file_feast_core_DynamoDBTable_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_feast_core_DynamoDBTable_proto_goTypes,
		DependencyIndexes: file_feast_core_DynamoDBTable_proto_depIdxs,
		MessageInfos:      file_feast_core_DynamoDBTable_proto_msgTypes,
	}.Build()
	File_feast_core_DynamoDBTable_proto = out.File
	file_feast_core_DynamoDBTable_proto_rawDesc = nil
	file_feast_core_DynamoDBTable_proto_goTypes = nil
	file_feast_core_DynamoDBTable_proto_depIdxs = nil
}
