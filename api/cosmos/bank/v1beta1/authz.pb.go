// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: cosmos/bank/v1beta1/authz.proto

package bankv1beta1

import (
	_ "github.com/cosmos/cosmos-proto"
	v1beta1 "github.com/cosmos/cosmos-sdk/api/cosmos/base/v1beta1"
	_ "github.com/gogo/protobuf/gogoproto"
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

// SendAuthorization allows the grantee to spend up to spend_limit coins from
// the granter's account.
//
// Since: cosmos-sdk 0.43
type SendAuthorization struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SpendLimit []*v1beta1.Coin `protobuf:"bytes,1,rep,name=spend_limit,json=spendLimit,proto3" json:"spend_limit,omitempty"`
}

func (x *SendAuthorization) Reset() {
	*x = SendAuthorization{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_bank_v1beta1_authz_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendAuthorization) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendAuthorization) ProtoMessage() {}

func (x *SendAuthorization) ProtoReflect() protoreflect.Message {
	mi := &file_cosmos_bank_v1beta1_authz_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendAuthorization.ProtoReflect.Descriptor instead.
func (*SendAuthorization) Descriptor() ([]byte, []int) {
	return file_cosmos_bank_v1beta1_authz_proto_rawDescGZIP(), []int{0}
}

func (x *SendAuthorization) GetSpendLimit() []*v1beta1.Coin {
	if x != nil {
		return x.SpendLimit
	}
	return nil
}

var File_cosmos_bank_v1beta1_authz_proto protoreflect.FileDescriptor

var file_cosmos_bank_v1beta1_authz_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x62, 0x61, 0x6e, 0x6b, 0x2f, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x13, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x14, 0x67, 0x6f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x63, 0x6f,
	0x73, 0x6d, 0x6f, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f,
	0x62, 0x61, 0x73, 0x65, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x63, 0x6f, 0x69,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x94, 0x01, 0x0a, 0x11, 0x53, 0x65, 0x6e, 0x64,
	0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x6c, 0x0a,
	0x0b, 0x73, 0x70, 0x65, 0x6e, 0x64, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61, 0x73, 0x65,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x43, 0x6f, 0x69, 0x6e, 0x42, 0x30, 0xc8,
	0xde, 0x1f, 0x00, 0xaa, 0xdf, 0x1f, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2d,
	0x73, 0x64, 0x6b, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x69, 0x6e, 0x73, 0x52,
	0x0a, 0x73, 0x70, 0x65, 0x6e, 0x64, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x3a, 0x11, 0xca, 0xb4, 0x2d,
	0x0d, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0xd5,
	0x01, 0x0a, 0x17, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61,
	0x6e, 0x6b, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x42, 0x0a, 0x41, 0x75, 0x74, 0x68,
	0x7a, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x73, 0x6d,
	0x6f, 0x73, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f,
	0x73, 0x2f, 0x62, 0x61, 0x6e, 0x6b, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x3b, 0x62,
	0x61, 0x6e, 0x6b, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x42, 0x58,
	0xaa, 0x02, 0x13, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x42, 0x61, 0x6e, 0x6b, 0x2e, 0x56,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xca, 0x02, 0x13, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x5c,
	0x42, 0x61, 0x6e, 0x6b, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xe2, 0x02, 0x1f, 0x43,
	0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x5c, 0x42, 0x61, 0x6e, 0x6b, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x15, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x3a, 0x3a, 0x42, 0x61, 0x6e, 0x6b, 0x3a, 0x3a, 0x56,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cosmos_bank_v1beta1_authz_proto_rawDescOnce sync.Once
	file_cosmos_bank_v1beta1_authz_proto_rawDescData = file_cosmos_bank_v1beta1_authz_proto_rawDesc
)

func file_cosmos_bank_v1beta1_authz_proto_rawDescGZIP() []byte {
	file_cosmos_bank_v1beta1_authz_proto_rawDescOnce.Do(func() {
		file_cosmos_bank_v1beta1_authz_proto_rawDescData = protoimpl.X.CompressGZIP(file_cosmos_bank_v1beta1_authz_proto_rawDescData)
	})
	return file_cosmos_bank_v1beta1_authz_proto_rawDescData
}

var file_cosmos_bank_v1beta1_authz_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_cosmos_bank_v1beta1_authz_proto_goTypes = []interface{}{
	(*SendAuthorization)(nil), // 0: cosmos.bank.v1beta1.SendAuthorization
	(*v1beta1.Coin)(nil),      // 1: cosmos.base.v1beta1.Coin
}
var file_cosmos_bank_v1beta1_authz_proto_depIdxs = []int32{
	1, // 0: cosmos.bank.v1beta1.SendAuthorization.spend_limit:type_name -> cosmos.base.v1beta1.Coin
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_cosmos_bank_v1beta1_authz_proto_init() }
func file_cosmos_bank_v1beta1_authz_proto_init() {
	if File_cosmos_bank_v1beta1_authz_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cosmos_bank_v1beta1_authz_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendAuthorization); i {
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
			RawDescriptor: file_cosmos_bank_v1beta1_authz_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cosmos_bank_v1beta1_authz_proto_goTypes,
		DependencyIndexes: file_cosmos_bank_v1beta1_authz_proto_depIdxs,
		MessageInfos:      file_cosmos_bank_v1beta1_authz_proto_msgTypes,
	}.Build()
	File_cosmos_bank_v1beta1_authz_proto = out.File
	file_cosmos_bank_v1beta1_authz_proto_rawDesc = nil
	file_cosmos_bank_v1beta1_authz_proto_goTypes = nil
	file_cosmos_bank_v1beta1_authz_proto_depIdxs = nil
}