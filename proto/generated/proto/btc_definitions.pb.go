// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: btc_definitions.proto

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

type Transaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InActiveChain bool               `protobuf:"varint,1,opt,name=InActiveChain,proto3" json:"InActiveChain,omitempty"`
	Hex           string             `protobuf:"bytes,2,opt,name=Hex,proto3" json:"Hex,omitempty"`
	TxId          string             `protobuf:"bytes,3,opt,name=TxId,proto3" json:"TxId,omitempty"`
	Hash          string             `protobuf:"bytes,4,opt,name=Hash,proto3" json:"Hash,omitempty"`
	Size          int64              `protobuf:"varint,5,opt,name=Size,proto3" json:"Size,omitempty"`
	VSize         int64              `protobuf:"varint,6,opt,name=VSize,proto3" json:"VSize,omitempty"`
	Weight        int64              `protobuf:"varint,7,opt,name=Weight,proto3" json:"Weight,omitempty"`
	Version       int64              `protobuf:"varint,8,opt,name=Version,proto3" json:"Version,omitempty"`
	LockTime      int64              `protobuf:"varint,9,opt,name=LockTime,proto3" json:"LockTime,omitempty"`
	BlockHash     string             `protobuf:"bytes,10,opt,name=BlockHash,proto3" json:"BlockHash,omitempty"`
	Confirmations int64              `protobuf:"varint,11,opt,name=Confirmations,proto3" json:"Confirmations,omitempty"`
	BlockTime     int64              `protobuf:"varint,12,opt,name=BlockTime,proto3" json:"BlockTime,omitempty"`
	Time          int64              `protobuf:"varint,13,opt,name=Time,proto3" json:"Time,omitempty"`
	Vin           []*TransactionVIn  `protobuf:"bytes,14,rep,name=Vin,proto3" json:"Vin,omitempty"`
	Vout          []*TransactionVOut `protobuf:"bytes,15,rep,name=Vout,proto3" json:"Vout,omitempty"`
}

func (x *Transaction) Reset() {
	*x = Transaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_btc_definitions_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transaction) ProtoMessage() {}

func (x *Transaction) ProtoReflect() protoreflect.Message {
	mi := &file_btc_definitions_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transaction.ProtoReflect.Descriptor instead.
func (*Transaction) Descriptor() ([]byte, []int) {
	return file_btc_definitions_proto_rawDescGZIP(), []int{0}
}

func (x *Transaction) GetInActiveChain() bool {
	if x != nil {
		return x.InActiveChain
	}
	return false
}

func (x *Transaction) GetHex() string {
	if x != nil {
		return x.Hex
	}
	return ""
}

func (x *Transaction) GetTxId() string {
	if x != nil {
		return x.TxId
	}
	return ""
}

func (x *Transaction) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

func (x *Transaction) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *Transaction) GetVSize() int64 {
	if x != nil {
		return x.VSize
	}
	return 0
}

func (x *Transaction) GetWeight() int64 {
	if x != nil {
		return x.Weight
	}
	return 0
}

func (x *Transaction) GetVersion() int64 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *Transaction) GetLockTime() int64 {
	if x != nil {
		return x.LockTime
	}
	return 0
}

func (x *Transaction) GetBlockHash() string {
	if x != nil {
		return x.BlockHash
	}
	return ""
}

func (x *Transaction) GetConfirmations() int64 {
	if x != nil {
		return x.Confirmations
	}
	return 0
}

func (x *Transaction) GetBlockTime() int64 {
	if x != nil {
		return x.BlockTime
	}
	return 0
}

func (x *Transaction) GetTime() int64 {
	if x != nil {
		return x.Time
	}
	return 0
}

func (x *Transaction) GetVin() []*TransactionVIn {
	if x != nil {
		return x.Vin
	}
	return nil
}

func (x *Transaction) GetVout() []*TransactionVOut {
	if x != nil {
		return x.Vout
	}
	return nil
}

type TransactionVIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TxId        string       `protobuf:"bytes,1,opt,name=TxId,proto3" json:"TxId,omitempty"`
	Vout        int64        `protobuf:"varint,2,opt,name=Vout,proto3" json:"Vout,omitempty"`
	ScriptSig   *TxScriptSig `protobuf:"bytes,3,opt,name=ScriptSig,proto3" json:"ScriptSig,omitempty"`
	Sequence    int64        `protobuf:"varint,4,opt,name=Sequence,proto3" json:"Sequence,omitempty"`
	TxInWitness []string     `protobuf:"bytes,5,rep,name=TxInWitness,proto3" json:"TxInWitness,omitempty"`
}

func (x *TransactionVIn) Reset() {
	*x = TransactionVIn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_btc_definitions_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionVIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionVIn) ProtoMessage() {}

func (x *TransactionVIn) ProtoReflect() protoreflect.Message {
	mi := &file_btc_definitions_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionVIn.ProtoReflect.Descriptor instead.
func (*TransactionVIn) Descriptor() ([]byte, []int) {
	return file_btc_definitions_proto_rawDescGZIP(), []int{1}
}

func (x *TransactionVIn) GetTxId() string {
	if x != nil {
		return x.TxId
	}
	return ""
}

func (x *TransactionVIn) GetVout() int64 {
	if x != nil {
		return x.Vout
	}
	return 0
}

func (x *TransactionVIn) GetScriptSig() *TxScriptSig {
	if x != nil {
		return x.ScriptSig
	}
	return nil
}

func (x *TransactionVIn) GetSequence() int64 {
	if x != nil {
		return x.Sequence
	}
	return 0
}

func (x *TransactionVIn) GetTxInWitness() []string {
	if x != nil {
		return x.TxInWitness
	}
	return nil
}

type TxScriptSig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Asm string `protobuf:"bytes,1,opt,name=Asm,proto3" json:"Asm,omitempty"`
	Hex string `protobuf:"bytes,2,opt,name=Hex,proto3" json:"Hex,omitempty"`
}

func (x *TxScriptSig) Reset() {
	*x = TxScriptSig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_btc_definitions_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TxScriptSig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TxScriptSig) ProtoMessage() {}

func (x *TxScriptSig) ProtoReflect() protoreflect.Message {
	mi := &file_btc_definitions_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TxScriptSig.ProtoReflect.Descriptor instead.
func (*TxScriptSig) Descriptor() ([]byte, []int) {
	return file_btc_definitions_proto_rawDescGZIP(), []int{2}
}

func (x *TxScriptSig) GetAsm() string {
	if x != nil {
		return x.Asm
	}
	return ""
}

func (x *TxScriptSig) GetHex() string {
	if x != nil {
		return x.Hex
	}
	return ""
}

type TransactionVOut struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value        float64                  `protobuf:"fixed64,1,opt,name=Value,proto3" json:"Value,omitempty"`
	N            int64                    `protobuf:"varint,2,opt,name=N,proto3" json:"N,omitempty"`
	ScriptPubKey *TransactionScriptPubKey `protobuf:"bytes,3,opt,name=ScriptPubKey,proto3" json:"ScriptPubKey,omitempty"`
}

func (x *TransactionVOut) Reset() {
	*x = TransactionVOut{}
	if protoimpl.UnsafeEnabled {
		mi := &file_btc_definitions_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionVOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionVOut) ProtoMessage() {}

func (x *TransactionVOut) ProtoReflect() protoreflect.Message {
	mi := &file_btc_definitions_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionVOut.ProtoReflect.Descriptor instead.
func (*TransactionVOut) Descriptor() ([]byte, []int) {
	return file_btc_definitions_proto_rawDescGZIP(), []int{3}
}

func (x *TransactionVOut) GetValue() float64 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *TransactionVOut) GetN() int64 {
	if x != nil {
		return x.N
	}
	return 0
}

func (x *TransactionVOut) GetScriptPubKey() *TransactionScriptPubKey {
	if x != nil {
		return x.ScriptPubKey
	}
	return nil
}

type TransactionScriptPubKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Asm       string   `protobuf:"bytes,1,opt,name=Asm,proto3" json:"Asm,omitempty"`
	Hex       string   `protobuf:"bytes,2,opt,name=Hex,proto3" json:"Hex,omitempty"`
	ReqSigs   int64    `protobuf:"varint,3,opt,name=ReqSigs,proto3" json:"ReqSigs,omitempty"`
	Type      string   `protobuf:"bytes,4,opt,name=Type,proto3" json:"Type,omitempty"`
	Addresses []string `protobuf:"bytes,5,rep,name=Addresses,proto3" json:"Addresses,omitempty"`
}

func (x *TransactionScriptPubKey) Reset() {
	*x = TransactionScriptPubKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_btc_definitions_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionScriptPubKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionScriptPubKey) ProtoMessage() {}

func (x *TransactionScriptPubKey) ProtoReflect() protoreflect.Message {
	mi := &file_btc_definitions_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionScriptPubKey.ProtoReflect.Descriptor instead.
func (*TransactionScriptPubKey) Descriptor() ([]byte, []int) {
	return file_btc_definitions_proto_rawDescGZIP(), []int{4}
}

func (x *TransactionScriptPubKey) GetAsm() string {
	if x != nil {
		return x.Asm
	}
	return ""
}

func (x *TransactionScriptPubKey) GetHex() string {
	if x != nil {
		return x.Hex
	}
	return ""
}

func (x *TransactionScriptPubKey) GetReqSigs() int64 {
	if x != nil {
		return x.ReqSigs
	}
	return 0
}

func (x *TransactionScriptPubKey) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *TransactionScriptPubKey) GetAddresses() []string {
	if x != nil {
		return x.Addresses
	}
	return nil
}

var File_btc_definitions_proto protoreflect.FileDescriptor

var file_btc_definitions_proto_rawDesc = []byte{
	0x0a, 0x15, 0x62, 0x74, 0x63, 0x5f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x75, 0x6e, 0x69, 0x6e, 0x64, 0x65, 0x78,
	0x65, 0x72, 0x2e, 0x62, 0x74, 0x63, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x22, 0xd6, 0x03, 0x0a, 0x0b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x24, 0x0a, 0x0d, 0x49, 0x6e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x43, 0x68, 0x61,
	0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x49, 0x6e, 0x41, 0x63, 0x74, 0x69,
	0x76, 0x65, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x48, 0x65, 0x78, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x48, 0x65, 0x78, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x78, 0x49,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x78, 0x49, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x48, 0x61, 0x73, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x48, 0x61, 0x73,
	0x68, 0x12, 0x12, 0x0a, 0x04, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x04, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x56, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x56, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x57,
	0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x57, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a,
	0x08, 0x4c, 0x6f, 0x63, 0x6b, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x4c, 0x6f, 0x63, 0x6b, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x42, 0x6c, 0x6f,
	0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x12, 0x24, 0x0a, 0x0d, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1c, 0x0a,
	0x09, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x54,
	0x69, 0x6d, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x3a, 0x0a, 0x03, 0x56, 0x69, 0x6e, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x75,
	0x6e, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x72, 0x2e, 0x62, 0x74, 0x63, 0x64, 0x65, 0x66, 0x69,
	0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x56, 0x49, 0x6e, 0x52, 0x03, 0x56, 0x69, 0x6e, 0x12, 0x3d, 0x0a, 0x04, 0x56,
	0x6f, 0x75, 0x74, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x75, 0x6e, 0x69, 0x6e,
	0x64, 0x65, 0x78, 0x65, 0x72, 0x2e, 0x62, 0x74, 0x63, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x56, 0x4f, 0x75, 0x74, 0x52, 0x04, 0x56, 0x6f, 0x75, 0x74, 0x22, 0xbb, 0x01, 0x0a, 0x0e, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x56, 0x49, 0x6e, 0x12, 0x12, 0x0a,
	0x04, 0x54, 0x78, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x78, 0x49,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x56, 0x6f, 0x75, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x04, 0x56, 0x6f, 0x75, 0x74, 0x12, 0x43, 0x0a, 0x09, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x53,
	0x69, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x75, 0x6e, 0x69, 0x6e, 0x64,
	0x65, 0x78, 0x65, 0x72, 0x2e, 0x62, 0x74, 0x63, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x54, 0x78, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x53, 0x69, 0x67, 0x52,
	0x09, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x53, 0x69, 0x67, 0x12, 0x1a, 0x0a, 0x08, 0x53, 0x65,
	0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x53, 0x65,
	0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x54, 0x78, 0x49, 0x6e, 0x57, 0x69,
	0x74, 0x6e, 0x65, 0x73, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x54, 0x78, 0x49,
	0x6e, 0x57, 0x69, 0x74, 0x6e, 0x65, 0x73, 0x73, 0x22, 0x31, 0x0a, 0x0b, 0x54, 0x78, 0x53, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x53, 0x69, 0x67, 0x12, 0x10, 0x0a, 0x03, 0x41, 0x73, 0x6d, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x41, 0x73, 0x6d, 0x12, 0x10, 0x0a, 0x03, 0x48, 0x65, 0x78,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x48, 0x65, 0x78, 0x22, 0x8c, 0x01, 0x0a, 0x0f,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x56, 0x4f, 0x75, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x0c, 0x0a, 0x01, 0x4e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x01, 0x4e, 0x12, 0x55, 0x0a, 0x0c, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x50, 0x75, 0x62,
	0x4b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x75, 0x6e, 0x69, 0x6e,
	0x64, 0x65, 0x78, 0x65, 0x72, 0x2e, 0x62, 0x74, 0x63, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x50, 0x75, 0x62, 0x4b, 0x65, 0x79, 0x52, 0x0c, 0x53, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x50, 0x75, 0x62, 0x4b, 0x65, 0x79, 0x22, 0x89, 0x01, 0x0a, 0x17, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x50, 0x75, 0x62, 0x4b, 0x65, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x41, 0x73, 0x6d, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x41, 0x73, 0x6d, 0x12, 0x10, 0x0a, 0x03, 0x48, 0x65, 0x78, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x48, 0x65, 0x78, 0x12, 0x18, 0x0a, 0x07, 0x52, 0x65,
	0x71, 0x53, 0x69, 0x67, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x52, 0x65, 0x71,
	0x53, 0x69, 0x67, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_btc_definitions_proto_rawDescOnce sync.Once
	file_btc_definitions_proto_rawDescData = file_btc_definitions_proto_rawDesc
)

func file_btc_definitions_proto_rawDescGZIP() []byte {
	file_btc_definitions_proto_rawDescOnce.Do(func() {
		file_btc_definitions_proto_rawDescData = protoimpl.X.CompressGZIP(file_btc_definitions_proto_rawDescData)
	})
	return file_btc_definitions_proto_rawDescData
}

var file_btc_definitions_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_btc_definitions_proto_goTypes = []interface{}{
	(*Transaction)(nil),             // 0: unindexer.btcdefinitions.Transaction
	(*TransactionVIn)(nil),          // 1: unindexer.btcdefinitions.TransactionVIn
	(*TxScriptSig)(nil),             // 2: unindexer.btcdefinitions.TxScriptSig
	(*TransactionVOut)(nil),         // 3: unindexer.btcdefinitions.TransactionVOut
	(*TransactionScriptPubKey)(nil), // 4: unindexer.btcdefinitions.TransactionScriptPubKey
}
var file_btc_definitions_proto_depIdxs = []int32{
	1, // 0: unindexer.btcdefinitions.Transaction.Vin:type_name -> unindexer.btcdefinitions.TransactionVIn
	3, // 1: unindexer.btcdefinitions.Transaction.Vout:type_name -> unindexer.btcdefinitions.TransactionVOut
	2, // 2: unindexer.btcdefinitions.TransactionVIn.ScriptSig:type_name -> unindexer.btcdefinitions.TxScriptSig
	4, // 3: unindexer.btcdefinitions.TransactionVOut.ScriptPubKey:type_name -> unindexer.btcdefinitions.TransactionScriptPubKey
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_btc_definitions_proto_init() }
func file_btc_definitions_proto_init() {
	if File_btc_definitions_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_btc_definitions_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Transaction); i {
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
		file_btc_definitions_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionVIn); i {
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
		file_btc_definitions_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TxScriptSig); i {
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
		file_btc_definitions_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionVOut); i {
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
		file_btc_definitions_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionScriptPubKey); i {
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
			RawDescriptor: file_btc_definitions_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_btc_definitions_proto_goTypes,
		DependencyIndexes: file_btc_definitions_proto_depIdxs,
		MessageInfos:      file_btc_definitions_proto_msgTypes,
	}.Build()
	File_btc_definitions_proto = out.File
	file_btc_definitions_proto_rawDesc = nil
	file_btc_definitions_proto_goTypes = nil
	file_btc_definitions_proto_depIdxs = nil
}
