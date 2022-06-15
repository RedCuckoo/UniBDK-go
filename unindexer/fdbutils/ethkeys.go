package fdbutils

import (
	"github.com/RedCuckoo/UniBDK-go/proto/generated/proto"
	"github.com/apple/foundationdb/bindings/go/src/fdb/tuple"
	"github.com/ethereum/go-ethereum/common"
)

// key : storage namespace
// value : marshaled Int64
func EthLastBlockKey() []byte {
	return tuple.Tuple{
		int(proto.NamespaceFdb_EthLastBlock),
	}.Pack()
}

// key : storage namespace + address + txId
// value : amount left
func EthBalanceKey(address, contractAddress common.Address) []byte {
	return tuple.Tuple{
		int(proto.NamespaceFdb_EthBalance),
		address.Bytes(),
		contractAddress.Bytes(),
	}.Pack()
}

// key : storage namespace + address
// value : amount left
func EthBalanceKeyPrefix(address common.Address) []byte {
	return tuple.Tuple{
		int(proto.NamespaceFdb_EthBalance),
		address.Bytes(),
	}.Pack()
}

// key : storage namespace + address
// value : amount left
func EthBalanceKeyGetContractAddress(key []byte) common.Address {
	unpack, err := tuple.Unpack(key)
	if err != nil {
		return common.Address{}
	}

	elements := []tuple.TupleElement(unpack)

	return common.BytesToAddress(elements[2].([]byte))
}
