// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: indexer.proto

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

type Blockchains int32

const (
	Blockchains_Bitcoin  Blockchains = 0
	Blockchains_Ethereum Blockchains = 1
)

// Enum value maps for Blockchains.
var (
	Blockchains_name = map[int32]string{
		0: "Bitcoin",
		1: "Ethereum",
	}
	Blockchains_value = map[string]int32{
		"Bitcoin":  0,
		"Ethereum": 1,
	}
)

func (x Blockchains) Enum() *Blockchains {
	p := new(Blockchains)
	*p = x
	return p
}

func (x Blockchains) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Blockchains) Descriptor() protoreflect.EnumDescriptor {
	return file_indexer_proto_enumTypes[0].Descriptor()
}

func (Blockchains) Type() protoreflect.EnumType {
	return &file_indexer_proto_enumTypes[0]
}

func (x Blockchains) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Blockchains.Descriptor instead.
func (Blockchains) EnumDescriptor() ([]byte, []int) {
	return file_indexer_proto_rawDescGZIP(), []int{0}
}

type AddressBalanceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address    string      `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Blockchain Blockchains `protobuf:"varint,2,opt,name=blockchain,proto3,enum=unindexer.btcindexer.Blockchains" json:"blockchain,omitempty"`
}

func (x *AddressBalanceRequest) Reset() {
	*x = AddressBalanceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_indexer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddressBalanceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddressBalanceRequest) ProtoMessage() {}

func (x *AddressBalanceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_indexer_proto_msgTypes[0]
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
	return file_indexer_proto_rawDescGZIP(), []int{0}
}

func (x *AddressBalanceRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *AddressBalanceRequest) GetBlockchain() Blockchains {
	if x != nil {
		return x.Blockchain
	}
	return Blockchains_Bitcoin
}

type AddressBalanceReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address  string         `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Balance  string         `protobuf:"bytes,2,opt,name=balance,proto3" json:"balance,omitempty"`
	Balances []*EthBalances `protobuf:"bytes,3,rep,name=balances,proto3" json:"balances,omitempty"`
}

func (x *AddressBalanceReply) Reset() {
	*x = AddressBalanceReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_indexer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddressBalanceReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddressBalanceReply) ProtoMessage() {}

func (x *AddressBalanceReply) ProtoReflect() protoreflect.Message {
	mi := &file_indexer_proto_msgTypes[1]
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
	return file_indexer_proto_rawDescGZIP(), []int{1}
}

func (x *AddressBalanceReply) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *AddressBalanceReply) GetBalance() string {
	if x != nil {
		return x.Balance
	}
	return ""
}

func (x *AddressBalanceReply) GetBalances() []*EthBalances {
	if x != nil {
		return x.Balances
	}
	return nil
}

type EthBalances struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ContractAddress string `protobuf:"bytes,1,opt,name=contractAddress,proto3" json:"contractAddress,omitempty"`
	Balance         string `protobuf:"bytes,2,opt,name=balance,proto3" json:"balance,omitempty"`
}

func (x *EthBalances) Reset() {
	*x = EthBalances{}
	if protoimpl.UnsafeEnabled {
		mi := &file_indexer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EthBalances) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EthBalances) ProtoMessage() {}

func (x *EthBalances) ProtoReflect() protoreflect.Message {
	mi := &file_indexer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EthBalances.ProtoReflect.Descriptor instead.
func (*EthBalances) Descriptor() ([]byte, []int) {
	return file_indexer_proto_rawDescGZIP(), []int{2}
}

func (x *EthBalances) GetContractAddress() string {
	if x != nil {
		return x.ContractAddress
	}
	return ""
}

func (x *EthBalances) GetBalance() string {
	if x != nil {
		return x.Balance
	}
	return ""
}

var File_indexer_proto protoreflect.FileDescriptor

var file_indexer_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x14, 0x75, 0x6e, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x72, 0x2e, 0x62, 0x74, 0x63, 0x69, 0x6e,
	0x64, 0x65, 0x78, 0x65, 0x72, 0x22, 0x74, 0x0a, 0x15, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18,
	0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x41, 0x0a, 0x0a, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x75,
	0x6e, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x72, 0x2e, 0x62, 0x74, 0x63, 0x69, 0x6e, 0x64, 0x65,
	0x78, 0x65, 0x72, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x73, 0x52,
	0x0a, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x22, 0x88, 0x01, 0x0a, 0x13,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a,
	0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x3d, 0x0a, 0x08, 0x62, 0x61, 0x6c, 0x61, 0x6e,
	0x63, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x75, 0x6e, 0x69, 0x6e,
	0x64, 0x65, 0x78, 0x65, 0x72, 0x2e, 0x62, 0x74, 0x63, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x72,
	0x2e, 0x45, 0x74, 0x68, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x52, 0x08, 0x62, 0x61,
	0x6c, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x22, 0x51, 0x0a, 0x0b, 0x45, 0x74, 0x68, 0x42, 0x61, 0x6c,
	0x61, 0x6e, 0x63, 0x65, 0x73, 0x12, 0x28, 0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x18, 0x0a, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x2a, 0x28, 0x0a, 0x0b, 0x42, 0x6c, 0x6f,
	0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x42, 0x69, 0x74, 0x63,
	0x6f, 0x69, 0x6e, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x45, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75,
	0x6d, 0x10, 0x01, 0x32, 0x73, 0x0a, 0x07, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x72, 0x12, 0x68,
	0x0a, 0x0e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65,
	0x12, 0x2b, 0x2e, 0x75, 0x6e, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x72, 0x2e, 0x62, 0x74, 0x63,
	0x69, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x72, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x42,
	0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e,
	0x75, 0x6e, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x72, 0x2e, 0x62, 0x74, 0x63, 0x69, 0x6e, 0x64,
	0x65, 0x78, 0x65, 0x72, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x42, 0x61, 0x6c, 0x61,
	0x6e, 0x63, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_indexer_proto_rawDescOnce sync.Once
	file_indexer_proto_rawDescData = file_indexer_proto_rawDesc
)

func file_indexer_proto_rawDescGZIP() []byte {
	file_indexer_proto_rawDescOnce.Do(func() {
		file_indexer_proto_rawDescData = protoimpl.X.CompressGZIP(file_indexer_proto_rawDescData)
	})
	return file_indexer_proto_rawDescData
}

var file_indexer_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_indexer_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_indexer_proto_goTypes = []interface{}{
	(Blockchains)(0),              // 0: unindexer.btcindexer.Blockchains
	(*AddressBalanceRequest)(nil), // 1: unindexer.btcindexer.AddressBalanceRequest
	(*AddressBalanceReply)(nil),   // 2: unindexer.btcindexer.AddressBalanceReply
	(*EthBalances)(nil),           // 3: unindexer.btcindexer.EthBalances
}
var file_indexer_proto_depIdxs = []int32{
	0, // 0: unindexer.btcindexer.AddressBalanceRequest.blockchain:type_name -> unindexer.btcindexer.Blockchains
	3, // 1: unindexer.btcindexer.AddressBalanceReply.balances:type_name -> unindexer.btcindexer.EthBalances
	1, // 2: unindexer.btcindexer.Indexer.AddressBalance:input_type -> unindexer.btcindexer.AddressBalanceRequest
	2, // 3: unindexer.btcindexer.Indexer.AddressBalance:output_type -> unindexer.btcindexer.AddressBalanceReply
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_indexer_proto_init() }
func file_indexer_proto_init() {
	if File_indexer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_indexer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_indexer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_indexer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EthBalances); i {
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
			RawDescriptor: file_indexer_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_indexer_proto_goTypes,
		DependencyIndexes: file_indexer_proto_depIdxs,
		EnumInfos:         file_indexer_proto_enumTypes,
		MessageInfos:      file_indexer_proto_msgTypes,
	}.Build()
	File_indexer_proto = out.File
	file_indexer_proto_rawDesc = nil
	file_indexer_proto_goTypes = nil
	file_indexer_proto_depIdxs = nil
}
