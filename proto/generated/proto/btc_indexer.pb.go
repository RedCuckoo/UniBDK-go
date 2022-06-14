// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: btc_indexer.proto

package proto

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

type AddressBalanceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *AddressBalanceRequest) Reset() {
	*x = AddressBalanceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_btc_indexer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddressBalanceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddressBalanceRequest) ProtoMessage() {}

func (x *AddressBalanceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_btc_indexer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddressBalanceRequest.ProtoReflect.Descriptor instead.
func (*AddressBalanceRequest) Descriptor() ([]byte, []int) {
	return file_btc_indexer_proto_rawDescGZIP(), []int{0}
}

func (x *AddressBalanceRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type AddressBalanceReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Balance float32 `protobuf:"fixed32,1,opt,name=balance,proto3" json:"balance,omitempty"`
}

func (x *AddressBalanceReply) Reset() {
	*x = AddressBalanceReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_btc_indexer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddressBalanceReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddressBalanceReply) ProtoMessage() {}

func (x *AddressBalanceReply) ProtoReflect() protoreflect.Message {
	mi := &file_btc_indexer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddressBalanceReply.ProtoReflect.Descriptor instead.
func (*AddressBalanceReply) Descriptor() ([]byte, []int) {
	return file_btc_indexer_proto_rawDescGZIP(), []int{1}
}

func (x *AddressBalanceReply) GetBalance() float32 {
	if x != nil {
		return x.Balance
	}
	return 0
}

var File_btc_indexer_proto protoreflect.FileDescriptor

var file_btc_indexer_proto_rawDesc = []byte{
	0x0a, 0x11, 0x62, 0x74, 0x63, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x14, 0x75, 0x6e, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x72, 0x2e, 0x62,
	0x74, 0x63, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x72, 0x22, 0x31, 0x0a, 0x15, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x2f, 0x0a, 0x13,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x32, 0x76, 0x0a,
	0x0a, 0x42, 0x74, 0x63, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x72, 0x12, 0x68, 0x0a, 0x0e, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x2b, 0x2e,
	0x75, 0x6e, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x72, 0x2e, 0x62, 0x74, 0x63, 0x69, 0x6e, 0x64,
	0x65, 0x78, 0x65, 0x72, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x42, 0x61, 0x6c, 0x61,
	0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x75, 0x6e, 0x69,
	0x6e, 0x64, 0x65, 0x78, 0x65, 0x72, 0x2e, 0x62, 0x74, 0x63, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x65,
	0x72, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_btc_indexer_proto_rawDescOnce sync.Once
	file_btc_indexer_proto_rawDescData = file_btc_indexer_proto_rawDesc
)

func file_btc_indexer_proto_rawDescGZIP() []byte {
	file_btc_indexer_proto_rawDescOnce.Do(func() {
		file_btc_indexer_proto_rawDescData = protoimpl.X.CompressGZIP(file_btc_indexer_proto_rawDescData)
	})
	return file_btc_indexer_proto_rawDescData
}

var file_btc_indexer_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_btc_indexer_proto_goTypes = []interface{}{
	(*AddressBalanceRequest)(nil), // 0: unindexer.btcindexer.AddressBalanceRequest
	(*AddressBalanceReply)(nil),   // 1: unindexer.btcindexer.AddressBalanceReply
}
var file_btc_indexer_proto_depIdxs = []int32{
	0, // 0: unindexer.btcindexer.BtcIndexer.AddressBalance:input_type -> unindexer.btcindexer.AddressBalanceRequest
	1, // 1: unindexer.btcindexer.BtcIndexer.AddressBalance:output_type -> unindexer.btcindexer.AddressBalanceReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_btc_indexer_proto_init() }
func file_btc_indexer_proto_init() {
	if File_btc_indexer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_btc_indexer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddressBalanceRequest); i {
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
		file_btc_indexer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddressBalanceReply); i {
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
			RawDescriptor: file_btc_indexer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_btc_indexer_proto_goTypes,
		DependencyIndexes: file_btc_indexer_proto_depIdxs,
		MessageInfos:      file_btc_indexer_proto_msgTypes,
	}.Build()
	File_btc_indexer_proto = out.File
	file_btc_indexer_proto_rawDesc = nil
	file_btc_indexer_proto_goTypes = nil
	file_btc_indexer_proto_depIdxs = nil
}