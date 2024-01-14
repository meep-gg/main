// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.1
// source: scylla/riot/account/v1/account.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ScyllaAccount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Puuid    string `protobuf:"bytes,1,opt,name=puuid,proto3" json:"puuid,omitempty"`
	GameName string `protobuf:"bytes,2,opt,name=gameName,proto3" json:"gameName,omitempty"`
	TagLine  string `protobuf:"bytes,3,opt,name=tagLine,proto3" json:"tagLine,omitempty"`
}

func (x *ScyllaAccount) Reset() {
	*x = ScyllaAccount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scylla_riot_account_v1_account_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScyllaAccount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScyllaAccount) ProtoMessage() {}

func (x *ScyllaAccount) ProtoReflect() protoreflect.Message {
	mi := &file_scylla_riot_account_v1_account_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScyllaAccount.ProtoReflect.Descriptor instead.
func (*ScyllaAccount) Descriptor() ([]byte, []int) {
	return file_scylla_riot_account_v1_account_proto_rawDescGZIP(), []int{0}
}

func (x *ScyllaAccount) GetPuuid() string {
	if x != nil {
		return x.Puuid
	}
	return ""
}

func (x *ScyllaAccount) GetGameName() string {
	if x != nil {
		return x.GameName
	}
	return ""
}

func (x *ScyllaAccount) GetTagLine() string {
	if x != nil {
		return x.TagLine
	}
	return ""
}

type ScyllaAccountReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Region string `protobuf:"bytes,1,opt,name=region,proto3" json:"region,omitempty"`
	Puuid  string `protobuf:"bytes,2,opt,name=puuid,proto3" json:"puuid,omitempty"`
}

func (x *ScyllaAccountReq) Reset() {
	*x = ScyllaAccountReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scylla_riot_account_v1_account_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScyllaAccountReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScyllaAccountReq) ProtoMessage() {}

func (x *ScyllaAccountReq) ProtoReflect() protoreflect.Message {
	mi := &file_scylla_riot_account_v1_account_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScyllaAccountReq.ProtoReflect.Descriptor instead.
func (*ScyllaAccountReq) Descriptor() ([]byte, []int) {
	return file_scylla_riot_account_v1_account_proto_rawDescGZIP(), []int{1}
}

func (x *ScyllaAccountReq) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *ScyllaAccountReq) GetPuuid() string {
	if x != nil {
		return x.Puuid
	}
	return ""
}

type ScyllaUpsertAccountReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Region  string         `protobuf:"bytes,1,opt,name=region,proto3" json:"region,omitempty"`
	Account *ScyllaAccount `protobuf:"bytes,2,opt,name=account,proto3" json:"account,omitempty"`
}

func (x *ScyllaUpsertAccountReq) Reset() {
	*x = ScyllaUpsertAccountReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scylla_riot_account_v1_account_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScyllaUpsertAccountReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScyllaUpsertAccountReq) ProtoMessage() {}

func (x *ScyllaUpsertAccountReq) ProtoReflect() protoreflect.Message {
	mi := &file_scylla_riot_account_v1_account_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScyllaUpsertAccountReq.ProtoReflect.Descriptor instead.
func (*ScyllaUpsertAccountReq) Descriptor() ([]byte, []int) {
	return file_scylla_riot_account_v1_account_proto_rawDescGZIP(), []int{2}
}

func (x *ScyllaUpsertAccountReq) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *ScyllaUpsertAccountReq) GetAccount() *ScyllaAccount {
	if x != nil {
		return x.Account
	}
	return nil
}

var File_scylla_riot_account_v1_account_proto protoreflect.FileDescriptor

var file_scylla_riot_account_v1_account_proto_rawDesc = []byte{
	0x0a, 0x24, 0x73, 0x63, 0x79, 0x6c, 0x6c, 0x61, 0x2f, 0x72, 0x69, 0x6f, 0x74, 0x2f, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x73, 0x63, 0x79, 0x6c, 0x6c, 0x61, 0x2e, 0x72,
	0x69, 0x6f, 0x74, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x1b,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5b, 0x0a, 0x0d, 0x53,
	0x63, 0x79, 0x6c, 0x6c, 0x61, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x70, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x75, 0x75,
	0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x67, 0x61, 0x6d, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x67, 0x61, 0x6d, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x74, 0x61, 0x67, 0x4c, 0x69, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x74, 0x61, 0x67, 0x4c, 0x69, 0x6e, 0x65, 0x22, 0x40, 0x0a, 0x10, 0x53, 0x63, 0x79, 0x6c,
	0x6c, 0x61, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06,
	0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x75, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x75, 0x75, 0x69, 0x64, 0x22, 0x71, 0x0a, 0x16, 0x53, 0x63,
	0x79, 0x6c, 0x6c, 0x61, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x3f, 0x0a, 0x07,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e,
	0x73, 0x63, 0x79, 0x6c, 0x6c, 0x61, 0x2e, 0x72, 0x69, 0x6f, 0x74, 0x2e, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x63, 0x79, 0x6c, 0x6c, 0x61, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x32, 0xa7, 0x02,
	0x0a, 0x14, 0x53, 0x63, 0x79, 0x6c, 0x6c, 0x61, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5f, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x28, 0x2e, 0x73, 0x63, 0x79, 0x6c, 0x6c, 0x61, 0x2e, 0x72, 0x69,
	0x6f, 0x74, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x63,
	0x79, 0x6c, 0x6c, 0x61, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x25,
	0x2e, 0x73, 0x63, 0x79, 0x6c, 0x6c, 0x61, 0x2e, 0x72, 0x69, 0x6f, 0x74, 0x2e, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x63, 0x79, 0x6c, 0x6c, 0x61, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x00, 0x12, 0x59, 0x0a, 0x0d, 0x55, 0x70, 0x73, 0x65, 0x72,
	0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2e, 0x2e, 0x73, 0x63, 0x79, 0x6c, 0x6c,
	0x61, 0x2e, 0x72, 0x69, 0x6f, 0x74, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x63, 0x79, 0x6c, 0x6c, 0x61, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x00, 0x12, 0x53, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x28, 0x2e, 0x73, 0x63, 0x79, 0x6c, 0x6c, 0x61, 0x2e, 0x72, 0x69, 0x6f,
	0x74, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x63, 0x79,
	0x6c, 0x6c, 0x61, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x27, 0x5a, 0x25, 0x6d, 0x65, 0x65, 0x70, 0x2e,
	0x67, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x73, 0x63, 0x79, 0x6c, 0x6c, 0x61,
	0x2f, 0x72, 0x69, 0x6f, 0x74, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_scylla_riot_account_v1_account_proto_rawDescOnce sync.Once
	file_scylla_riot_account_v1_account_proto_rawDescData = file_scylla_riot_account_v1_account_proto_rawDesc
)

func file_scylla_riot_account_v1_account_proto_rawDescGZIP() []byte {
	file_scylla_riot_account_v1_account_proto_rawDescOnce.Do(func() {
		file_scylla_riot_account_v1_account_proto_rawDescData = protoimpl.X.CompressGZIP(file_scylla_riot_account_v1_account_proto_rawDescData)
	})
	return file_scylla_riot_account_v1_account_proto_rawDescData
}

var file_scylla_riot_account_v1_account_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_scylla_riot_account_v1_account_proto_goTypes = []interface{}{
	(*ScyllaAccount)(nil),          // 0: scylla.riot.account.v1.ScyllaAccount
	(*ScyllaAccountReq)(nil),       // 1: scylla.riot.account.v1.ScyllaAccountReq
	(*ScyllaUpsertAccountReq)(nil), // 2: scylla.riot.account.v1.ScyllaUpsertAccountReq
	(*emptypb.Empty)(nil),          // 3: google.protobuf.Empty
}
var file_scylla_riot_account_v1_account_proto_depIdxs = []int32{
	0, // 0: scylla.riot.account.v1.ScyllaUpsertAccountReq.account:type_name -> scylla.riot.account.v1.ScyllaAccount
	1, // 1: scylla.riot.account.v1.ScyllaAccountService.GetAccount:input_type -> scylla.riot.account.v1.ScyllaAccountReq
	2, // 2: scylla.riot.account.v1.ScyllaAccountService.UpsertAccount:input_type -> scylla.riot.account.v1.ScyllaUpsertAccountReq
	1, // 3: scylla.riot.account.v1.ScyllaAccountService.DeleteAccount:input_type -> scylla.riot.account.v1.ScyllaAccountReq
	0, // 4: scylla.riot.account.v1.ScyllaAccountService.GetAccount:output_type -> scylla.riot.account.v1.ScyllaAccount
	3, // 5: scylla.riot.account.v1.ScyllaAccountService.UpsertAccount:output_type -> google.protobuf.Empty
	3, // 6: scylla.riot.account.v1.ScyllaAccountService.DeleteAccount:output_type -> google.protobuf.Empty
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_scylla_riot_account_v1_account_proto_init() }
func file_scylla_riot_account_v1_account_proto_init() {
	if File_scylla_riot_account_v1_account_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_scylla_riot_account_v1_account_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScyllaAccount); i {
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
		file_scylla_riot_account_v1_account_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScyllaAccountReq); i {
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
		file_scylla_riot_account_v1_account_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScyllaUpsertAccountReq); i {
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
			RawDescriptor: file_scylla_riot_account_v1_account_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_scylla_riot_account_v1_account_proto_goTypes,
		DependencyIndexes: file_scylla_riot_account_v1_account_proto_depIdxs,
		MessageInfos:      file_scylla_riot_account_v1_account_proto_msgTypes,
	}.Build()
	File_scylla_riot_account_v1_account_proto = out.File
	file_scylla_riot_account_v1_account_proto_rawDesc = nil
	file_scylla_riot_account_v1_account_proto_goTypes = nil
	file_scylla_riot_account_v1_account_proto_depIdxs = nil
}
