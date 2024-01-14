// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.1
// source: gateway/riot/account/v1/account.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	v1 "meep.gg/protos/scylla/riot/account/v1"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GatewayAccountReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Region   string `protobuf:"bytes,1,opt,name=region,proto3" json:"region,omitempty"`
	Puuid    string `protobuf:"bytes,2,opt,name=puuid,proto3" json:"puuid,omitempty"`
	GameName string `protobuf:"bytes,3,opt,name=gameName,proto3" json:"gameName,omitempty"`
	TagLine  string `protobuf:"bytes,4,opt,name=tagLine,proto3" json:"tagLine,omitempty"`
}

func (x *GatewayAccountReq) Reset() {
	*x = GatewayAccountReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_riot_account_v1_account_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GatewayAccountReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GatewayAccountReq) ProtoMessage() {}

func (x *GatewayAccountReq) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_riot_account_v1_account_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GatewayAccountReq.ProtoReflect.Descriptor instead.
func (*GatewayAccountReq) Descriptor() ([]byte, []int) {
	return file_gateway_riot_account_v1_account_proto_rawDescGZIP(), []int{0}
}

func (x *GatewayAccountReq) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *GatewayAccountReq) GetPuuid() string {
	if x != nil {
		return x.Puuid
	}
	return ""
}

func (x *GatewayAccountReq) GetGameName() string {
	if x != nil {
		return x.GameName
	}
	return ""
}

func (x *GatewayAccountReq) GetTagLine() string {
	if x != nil {
		return x.TagLine
	}
	return ""
}

var File_gateway_riot_account_v1_account_proto protoreflect.FileDescriptor

var file_gateway_riot_account_v1_account_proto_rawDesc = []byte{
	0x0a, 0x25, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x72, 0x69, 0x6f, 0x74, 0x2f, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x2e, 0x72, 0x69, 0x6f, 0x74, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76, 0x31,
	0x1a, 0x24, 0x73, 0x63, 0x79, 0x6c, 0x6c, 0x61, 0x2f, 0x72, 0x69, 0x6f, 0x74, 0x2f, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x77, 0x0a, 0x11, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x72,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x67,
	0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x75, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x70, 0x75, 0x75, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x67, 0x61, 0x6d,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x67, 0x61, 0x6d,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x61, 0x67, 0x4c, 0x69, 0x6e, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x61, 0x67, 0x4c, 0x69, 0x6e, 0x65, 0x32,
	0xe0, 0x01, 0x0a, 0x15, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x61, 0x0a, 0x0a, 0x47, 0x65, 0x74,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2a, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x2e, 0x72, 0x69, 0x6f, 0x74, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x1a, 0x25, 0x2e, 0x73, 0x63, 0x79, 0x6c, 0x6c, 0x61, 0x2e, 0x72, 0x69, 0x6f,
	0x74, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x63, 0x79,
	0x6c, 0x6c, 0x61, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x00, 0x12, 0x64, 0x0a, 0x0d,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2a, 0x2e,
	0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x72, 0x69, 0x6f, 0x74, 0x2e, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x25, 0x2e, 0x73, 0x63, 0x79, 0x6c,
	0x6c, 0x61, 0x2e, 0x72, 0x69, 0x6f, 0x74, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x63, 0x79, 0x6c, 0x6c, 0x61, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x22, 0x00, 0x42, 0x28, 0x5a, 0x26, 0x6d, 0x65, 0x65, 0x70, 0x2e, 0x67, 0x67, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x72, 0x69, 0x6f,
	0x74, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gateway_riot_account_v1_account_proto_rawDescOnce sync.Once
	file_gateway_riot_account_v1_account_proto_rawDescData = file_gateway_riot_account_v1_account_proto_rawDesc
)

func file_gateway_riot_account_v1_account_proto_rawDescGZIP() []byte {
	file_gateway_riot_account_v1_account_proto_rawDescOnce.Do(func() {
		file_gateway_riot_account_v1_account_proto_rawDescData = protoimpl.X.CompressGZIP(file_gateway_riot_account_v1_account_proto_rawDescData)
	})
	return file_gateway_riot_account_v1_account_proto_rawDescData
}

var file_gateway_riot_account_v1_account_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_gateway_riot_account_v1_account_proto_goTypes = []interface{}{
	(*GatewayAccountReq)(nil), // 0: gateway.riot.account.v1.GatewayAccountReq
	(*v1.ScyllaAccount)(nil),  // 1: scylla.riot.account.v1.ScyllaAccount
}
var file_gateway_riot_account_v1_account_proto_depIdxs = []int32{
	0, // 0: gateway.riot.account.v1.GatewayAccountService.GetAccount:input_type -> gateway.riot.account.v1.GatewayAccountReq
	0, // 1: gateway.riot.account.v1.GatewayAccountService.UpdateAccount:input_type -> gateway.riot.account.v1.GatewayAccountReq
	1, // 2: gateway.riot.account.v1.GatewayAccountService.GetAccount:output_type -> scylla.riot.account.v1.ScyllaAccount
	1, // 3: gateway.riot.account.v1.GatewayAccountService.UpdateAccount:output_type -> scylla.riot.account.v1.ScyllaAccount
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_gateway_riot_account_v1_account_proto_init() }
func file_gateway_riot_account_v1_account_proto_init() {
	if File_gateway_riot_account_v1_account_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gateway_riot_account_v1_account_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GatewayAccountReq); i {
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
			RawDescriptor: file_gateway_riot_account_v1_account_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_gateway_riot_account_v1_account_proto_goTypes,
		DependencyIndexes: file_gateway_riot_account_v1_account_proto_depIdxs,
		MessageInfos:      file_gateway_riot_account_v1_account_proto_msgTypes,
	}.Build()
	File_gateway_riot_account_v1_account_proto = out.File
	file_gateway_riot_account_v1_account_proto_rawDesc = nil
	file_gateway_riot_account_v1_account_proto_goTypes = nil
	file_gateway_riot_account_v1_account_proto_depIdxs = nil
}