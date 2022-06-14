package btc_indexer

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/RedCuckoo/UniBDK-go/unindexer/fdbutils"
	"github.com/apple/foundationdb/bindings/go/src/fdb"
	"gitlab.com/distributed_lab/running"
	"log"
	"time"
)

type Address struct {
	// txids
	UnspentTransactions []string
}

func (s *Service) Run(ctx context.Context) {
	s.log.Info("Running btc indexer service")

	startBlock, err := s.FoundationDB.ReadLastBlock()
	if err != nil {
		panic("failed to read last btc block from db")
	}
	s.StartBlock = startBlock

	addresses := make(map[string]*Address)

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
				marshal, err := json.Marshal(transaction)
				if err != nil {
					return nil, err
				}
				transaction.Set(fdb.Key(fdbutils.BtcTransactionKey(&tx.Result)), marshal)
				return nil, nil
			})
			if err != nil {
				s.log.Error(err)
				return err
			}

			var address string

			for _, out := range tx.Result.Vout {
				switch out.ScriptPubKey.Type {
				case "pubkeyhash", "scripthash", "witness_v0_keyhash", "witness_v0_scripthash":
					address = out.ScriptPubKey.Address
				case "nulldata":
					continue
				default:
					return errors.New(fmt.Sprintf("unsupported address type %s", out.ScriptPubKey.Type))
				}
			}
			if len(tx.Result.Vin) == 1 && tx.Result.Vin[0].Coinbase != "" {
				if addresses[address] == nil {
					addresses[address] = &Address{UnspentTransactions: []string{address}}
				} else {
					addresses[address].UnspentTransactions = append(addresses[address].UnspentTransactions, address)
				}
			} else {
				log.Printf("transaction %s\n", tx.Result.TxId)

				fee := 0.0
				for _, vin := range tx.Result.Vin {
					tx, err := s.GetRawTransaction(vin.TxId)
					if err != nil {
						return err
					}
					log.Printf("\tfrom %s amount %.8f", tx.Result.Vout[vin.Vout].ScriptPubKey.Address, tx.Result.Vout[vin.Vout].Value)
					fee += tx.Result.Vout[vin.Vout].Value
				}

				for _, vout := range tx.Result.Vout {
					log.Printf("\tto %s amount %.8f", vout.ScriptPubKey.Address, vout.Value)
					fee -= vout.Value
				}

				log.Printf("\tfee %.8f", fee)

				log.Println()
			}

		}

		return nil
	}, 1*time.Second, 5*time.Second, 10*time.Second)
}
