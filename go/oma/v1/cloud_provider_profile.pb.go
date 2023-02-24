// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: oma/v1/cloud_provider_profile.proto

package omav1

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

// A `CloudProviderProfile` associates a set of cloud provider-specific
// configuration settings with an `Account`
type CloudProviderProfile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Id is a globally unique identifier
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Provider identifies the cloud provider
	CloudProvider CloudProvider `protobuf:"varint,2,opt,name=cloud_provider,json=cloudProvider,proto3,enum=oma.v1.CloudProvider" json:"cloud_provider,omitempty"`
	// AccountId is the `any.cloud` Account identifier for this profile
	AccountId string `protobuf:"bytes,3,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	// Tags contains a set of free-form key/value pairs
	Tags map[string]string `protobuf:"bytes,15,rep,name=tags,proto3" json:"tags,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *CloudProviderProfile) Reset() {
	*x = CloudProviderProfile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oma_v1_cloud_provider_profile_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CloudProviderProfile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloudProviderProfile) ProtoMessage() {}

func (x *CloudProviderProfile) ProtoReflect() protoreflect.Message {
	mi := &file_oma_v1_cloud_provider_profile_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloudProviderProfile.ProtoReflect.Descriptor instead.
func (*CloudProviderProfile) Descriptor() ([]byte, []int) {
	return file_oma_v1_cloud_provider_profile_proto_rawDescGZIP(), []int{0}
}

func (x *CloudProviderProfile) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CloudProviderProfile) GetCloudProvider() CloudProvider {
	if x != nil {
		return x.CloudProvider
	}
	return CloudProvider_AWS
}

func (x *CloudProviderProfile) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

func (x *CloudProviderProfile) GetTags() map[string]string {
	if x != nil {
		return x.Tags
	}
	return nil
}

var File_oma_v1_cloud_provider_profile_proto protoreflect.FileDescriptor

var file_oma_v1_cloud_provider_profile_proto_rawDesc = []byte{
	0x0a, 0x23, 0x6f, 0x6d, 0x61, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x6f, 0x6d, 0x61, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x6f,
	0x6d, 0x61, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x70, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf8, 0x01, 0x0a, 0x14, 0x43,
	0x6c, 0x6f, 0x75, 0x64, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x3c, 0x0a, 0x0e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x70, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x6f, 0x6d,
	0x61, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x52, 0x0d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64,
	0x12, 0x3a, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26,
	0x2e, 0x6f, 0x6d, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x50, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x54, 0x61, 0x67,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x1a, 0x37, 0x0a, 0x09,
	0x54, 0x61, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x8c, 0x01, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x2e, 0x6f, 0x6d,
	0x61, 0x2e, 0x76, 0x31, 0x42, 0x19, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x50, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6e,
	0x79, 0x64, 0x6f, 0x74, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6f,
	0x2f, 0x6f, 0x6d, 0x61, 0x2f, 0x76, 0x31, 0x3b, 0x6f, 0x6d, 0x61, 0x76, 0x31, 0xa2, 0x02, 0x03,
	0x4f, 0x58, 0x58, 0xaa, 0x02, 0x06, 0x4f, 0x6d, 0x61, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x06, 0x4f,
	0x6d, 0x61, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x12, 0x4f, 0x6d, 0x61, 0x5c, 0x56, 0x31, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x07, 0x4f, 0x6d, 0x61,
	0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_oma_v1_cloud_provider_profile_proto_rawDescOnce sync.Once
	file_oma_v1_cloud_provider_profile_proto_rawDescData = file_oma_v1_cloud_provider_profile_proto_rawDesc
)

func file_oma_v1_cloud_provider_profile_proto_rawDescGZIP() []byte {
	file_oma_v1_cloud_provider_profile_proto_rawDescOnce.Do(func() {
		file_oma_v1_cloud_provider_profile_proto_rawDescData = protoimpl.X.CompressGZIP(file_oma_v1_cloud_provider_profile_proto_rawDescData)
	})
	return file_oma_v1_cloud_provider_profile_proto_rawDescData
}

var file_oma_v1_cloud_provider_profile_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_oma_v1_cloud_provider_profile_proto_goTypes = []interface{}{
	(*CloudProviderProfile)(nil), // 0: oma.v1.CloudProviderProfile
	nil,                          // 1: oma.v1.CloudProviderProfile.TagsEntry
	(CloudProvider)(0),           // 2: oma.v1.CloudProvider
}
var file_oma_v1_cloud_provider_profile_proto_depIdxs = []int32{
	2, // 0: oma.v1.CloudProviderProfile.cloud_provider:type_name -> oma.v1.CloudProvider
	1, // 1: oma.v1.CloudProviderProfile.tags:type_name -> oma.v1.CloudProviderProfile.TagsEntry
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_oma_v1_cloud_provider_profile_proto_init() }
func file_oma_v1_cloud_provider_profile_proto_init() {
	if File_oma_v1_cloud_provider_profile_proto != nil {
		return
	}
	file_oma_v1_cloud_provider_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_oma_v1_cloud_provider_profile_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CloudProviderProfile); i {
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
			RawDescriptor: file_oma_v1_cloud_provider_profile_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_oma_v1_cloud_provider_profile_proto_goTypes,
		DependencyIndexes: file_oma_v1_cloud_provider_profile_proto_depIdxs,
		MessageInfos:      file_oma_v1_cloud_provider_profile_proto_msgTypes,
	}.Build()
	File_oma_v1_cloud_provider_profile_proto = out.File
	file_oma_v1_cloud_provider_profile_proto_rawDesc = nil
	file_oma_v1_cloud_provider_profile_proto_goTypes = nil
	file_oma_v1_cloud_provider_profile_proto_depIdxs = nil
}