package fdbutils

import (
	"github.com/RedCuckoo/UniBDK-go/proto/generated/proto"
	"github.com/apple/foundationdb/bindings/go/src/fdb"
	"math"
)

func (db *FDB) ReadLastBlock(blockchain proto.Blockchains) (int64, error) {
	blockNumber, err := db.ReadTransact(func(transaction fdb.ReadTransaction) (interface{}, error) {
		var key fdb.Key
		switch blockchain {
		case proto.Blockchains_Bitcoin:
			key = BtcLastBlockKey()
		case proto.Blockchains_Ethereum:
			key = EthLastBlockKey()
		}
		futureKey := transaction.Get(key)
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

func (db *FDB) GetAddressBalance(address string) (float32, error) {
	balance, err := db.ReadTransact(func(transaction fdb.ReadTransaction) (interface{}, error) {
		prefixRange, err := fdb.PrefixRange(BtcAddressKeyPrefix(address))
		if err != nil {
			return nil, err
		}
		it := transaction.GetRange(prefixRange, fdb.RangeOptions{}).Iterator()

		balance := 0.0

		for it.Advance() {
			balance += math.Float64frombits(DecodeUInt64(it.MustGet().Value))
		}

		return balance, nil
	})
	if err != nil {
		return 0, err
	}

	return float32(balance.(float64)), nil
}
