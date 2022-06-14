package fdbutils

import (
	"bytes"
	"encoding/binary"
	"github.com/RedCuckoo/UniBDK-go/proto/generated/proto"
	"github.com/apple/foundationdb/bindings/go/src/fdb"
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
func BtcTransactionKey(transaction *Transaction) []byte {
	return tuple.Tuple{
		int(proto.NamespaceFdb_BtcTransaction),
		transaction.TxId,
	}.Pack()
}

func EncodeInt64(v int64) []byte {
	var buf bytes.Buffer

	data := make([]byte, binary.MaxVarintLen64)
	binary.PutVarint(data, v)

	buf.Write(data)

	return buf.Bytes()
}

func DecodeInt64(value []byte, dest *int64) (int, error) {
	res, _ := binary.Varint(value)
	*dest = res
	return binary.MaxVarintLen64, nil
}

func (db *FDB) ReadLastBlock() (int64, error) {
	blockNumber, err := db.ReadTransact(func(transaction fdb.ReadTransaction) (interface{}, error) {
		futureKey := transaction.Get(fdb.Key(BtcLastBlockKey()))
		if blockNumber, err := futureKey.Get(); err == nil {
			if blockNumber == nil {
				return int64(0), nil
			}
			var res int64
			decodeInt64, err := DecodeInt64(blockNumber, &res)
			if err != nil {
				return nil, err
			}
			return decodeInt64, nil
		} else {
			return nil, err
		}
	})

	if err != nil {
		return 0, err
	}

	return blockNumber.(int64), err
}
