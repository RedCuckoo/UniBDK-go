package fdbutils

import (
	"github.com/RedCuckoo/UniBDK-go/proto/generated/proto"
	"github.com/apple/foundationdb/bindings/go/src/fdb/tuple"
)

// key : storage namespace
// value : marshaled Int64
func BtcLastBlockKey() []byte {
	return tuple.Tuple{
		int(proto.NamespaceFdb_BtcLastBlock),
	}.Pack()
}

// key : storage namespace + transaction id
// value : marshaled btc_indexer.Transaction
func BtcTransactionKey(txId string) []byte {
	return tuple.Tuple{
		int(proto.NamespaceFdb_BtcTransaction),
		txId,
	}.Pack()
}

// key : storage namespace + address + txId
// value : amount left
func BtcAddressKey(address, txId string) []byte {
	return tuple.Tuple{
		int(proto.NamespaceFdb_BtcAddress),
		address,
		txId,
	}.Pack()
}

// key : storage namespace + address
// value : amount left
func BtcAddressKeyPrefix(address string) []byte {
	return tuple.Tuple{
		int(proto.NamespaceFdb_BtcAddress),
		address,
	}.Pack()
}
