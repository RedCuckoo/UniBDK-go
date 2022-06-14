package btc_indexer

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/RedCuckoo/UniBDK-go/proto/generated/proto"
	"github.com/RedCuckoo/UniBDK-go/unindexer/fdbutils"
	"github.com/apple/foundationdb/bindings/go/src/fdb"
	"gitlab.com/distributed_lab/running"
	"log"
	"math"
	"time"
)

func (s *Service) Run(ctx context.Context) {
	s.log.Info("Running btc indexer service")

	startBlock, err := s.FoundationDB.ReadLastBlock(proto.Blockchains_Bitcoin)
	if err != nil {
		panic("failed to read last btc block from db")
	}
	s.StartBlock = startBlock

	running.WithBackOff(ctx, s.log, "new-checker-service", func(ctx context.Context) error {
		if s.StartBlock == 0 {
			s.StartBlock = 2055584
		}

		if s.EndBlock == 0 {
			s.EndBlock = s.StartBlock + 1
		} else {
			s.StartBlock = s.EndBlock
			s.EndBlock = s.EndBlock + 1
		}

		s.log.Infof("processing from %d to %d", s.StartBlock, s.EndBlock)

		blockHash, err := s.GetBlockHash(s.StartBlock)

		if err != nil {
			s.log.Error(err)
			return err
		}

		blockResponse, err := s.GetBlockByHash(blockHash.Result)

		if err != nil {
			s.log.Error(err)
			return err
		}

		for _, txId := range blockResponse.Result.Tx {
			tx, err := s.GetRawTransaction(txId)
			if err != nil {
				return err
			}

			_, err = s.FoundationDB.Transact(func(transaction fdb.Transaction) (interface{}, error) {
				marshal, err := json.Marshal(tx)
				if err != nil {
					s.log.Error(err)
					return nil, err
				}
				transaction.Set(fdb.Key(fdbutils.BtcTransactionKey(txId)), marshal)
				return nil, nil
			})
			if err != nil {
				s.log.Error(err)
				return err
			}

			log.Printf("transaction %s\n", tx.Result.TxId)
			fee := 0.0

			for _, vin := range tx.Result.Vin {
				if vin.Coinbase == "" {
					inTx, err := s.GetRawTransaction(vin.TxId)
					fee += inTx.Result.Vout[vin.Vout].Value
					if err != nil {
						return err
					}
					log.Printf("\tfrom %s amount %.8f", inTx.Result.Vout[vin.Vout].ScriptPubKey.Address, inTx.Result.Vout[vin.Vout].Value)

					_, err = s.FoundationDB.Transact(func(transaction fdb.Transaction) (interface{}, error) {
						transaction.Clear(fdb.Key(fdbutils.BtcAddressKey(inTx.Result.Vout[vin.Vout].ScriptPubKey.Address, vin.TxId)))
						return nil, nil
					})
					if err != nil {
						return err
					}
				} else {
					fee += tx.Result.Vout[0].Value
					if err != nil {
						return err
					}
					log.Printf("\tminted %.8f", tx.Result.Vout[0].Value)
				}
			}

			for i, vout := range tx.Result.Vout {
				fee -= vout.Value
				switch vout.ScriptPubKey.Type {
				case "pubkeyhash", "scripthash", "witness_v0_keyhash", "witness_v0_scripthash":
					address := vout.ScriptPubKey.Address
					log.Printf("\tto %s amount %.8f", address, tx.Result.Vout[i].Value)

					_, err = s.FoundationDB.Transact(func(transaction fdb.Transaction) (interface{}, error) {
						transaction.Set(fdb.Key(fdbutils.BtcAddressKey(address, txId)), fdbutils.EncodeUInt64(math.Float64bits(tx.Result.Vout[i].Value)))
						return nil, nil
					})
				case "nulldata":
					continue
				default:
					return errors.New(fmt.Sprintf("unsupported address type %s", vout.ScriptPubKey.Type))
				}
			}

			log.Printf("\tfee %.8f", fee)

			log.Println()

		}

		return nil
	}, 1*time.Second, 5*time.Second, 10*time.Second)
}
