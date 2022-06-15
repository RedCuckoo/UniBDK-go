package fdbutils

import (
	"context"
	"fmt"
	"github.com/RedCuckoo/UniBDK-go/proto/generated/proto"
	"github.com/apple/foundationdb/bindings/go/src/fdb"
	"github.com/ethereum/go-ethereum/common"
	"math"
	"math/big"
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

func (db *FDB) GetAddressBalance(ctx context.Context, request *proto.AddressBalanceRequest) (*proto.AddressBalanceReply, error) {
	reply, err := db.ReadTransact(func(transaction fdb.ReadTransaction) (interface{}, error) {
		var prefixRange fdb.KeyRange
		var err error
		balances := make([]*proto.EthBalances, 0)
		balance := float32(0.0)
		var resbalance string
		var address string
		switch request.Blockchain {
		case proto.Blockchains_Bitcoin:
			address = request.Address
			prefixRange, err = fdb.PrefixRange(BtcAddressKeyPrefix(request.Address))
			if err != nil {
				return nil, err
			}
			it := transaction.GetRange(prefixRange, fdb.RangeOptions{}).Iterator()

			for it.Advance() {
				balance += float32(math.Float64frombits(DecodeUInt64(it.MustGet().Value)))
			}
		case proto.Blockchains_Ethereum:
			address = common.HexToAddress(request.Address).String()
			prefixRange, err = fdb.PrefixRange(EthBalanceKeyPrefix(common.HexToAddress(request.Address)))
			if err != nil {
				return nil, err
			}
			ethBalance, err := db.cfg.EthAuth().EthClient.BalanceAt(ctx, common.HexToAddress(request.Address), nil)
			if err != nil {
				return nil, err
			}
			resbalance = ethBalance.String()
			it := transaction.GetRange(prefixRange, fdb.RangeOptions{}).Iterator()
			for it.Advance() {
				keyValue := it.MustGet()
				contractAddress := EthBalanceKeyGetContractAddress(keyValue.Key)
				balances = append(balances, &proto.EthBalances{
					ContractAddress: contractAddress.String(),
					Balance:         new(big.Int).SetBytes(keyValue.Value).String(),
				})
			}
		}

		switch request.Blockchain {
		case proto.Blockchains_Bitcoin:
			resbalance = fmt.Sprintf("%.8f", balance)
		}

		return &proto.AddressBalanceReply{
			Address:  address,
			Balance:  resbalance,
			Balances: balances,
		}, nil
	})
	if err != nil {
		return nil, err
	}

	return reply.(*proto.AddressBalanceReply), nil
}
