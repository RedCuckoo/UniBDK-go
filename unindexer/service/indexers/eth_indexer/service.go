package eth_indexer

import (
	"context"
	"github.com/RedCuckoo/UniBDK-go/proto/generated/proto"
	"github.com/RedCuckoo/UniBDK-go/unindexer/fdbutils"
	"github.com/RedCuckoo/UniBDK-go/unindexer/service/indexers/eth_indexer/abi"
	"github.com/apple/foundationdb/bindings/go/src/fdb"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"gitlab.com/distributed_lab/running"
	"math/big"
	"sync"
	"time"
)

func (s *Service) Run(ctx context.Context) {
	s.log.Info("Running eth indexer service")

	startBlock, err := s.FoundationDB.ReadLastBlock(proto.Blockchains_Ethereum)
	if err != nil {
		panic("failed to read last eth block from db")
	}
	s.StartBlock = startBlock

	running.WithBackOff(ctx, s.log, "new-eth-indexer", func(ctx context.Context) error {
		if s.StartBlock == 0 {
			s.StartBlock = 13627170
		}

		if s.EndBlock == 0 {
			s.EndBlock = s.StartBlock + 1
		} else {
			s.StartBlock = s.EndBlock
			s.EndBlock = s.EndBlock + 1
		}

		s.log.Infof("processing from %d to %d", s.StartBlock, s.EndBlock)
		logs, err := s.client.FilterLogs(ctx, ethereum.FilterQuery{
			BlockHash: nil,
			FromBlock: big.NewInt(s.StartBlock),
			ToBlock:   big.NewInt(s.EndBlock),
			Topics:    [][]common.Hash{{common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef")}},
		})
		if err != nil {
			return err
		}

		var wg sync.WaitGroup
		batches := splitToBatches(s.batchSize, logs)
		for _, batch := range batches {
			wg.Add(1)
			go s.ProcessLog(&wg, batch)
		}
		wg.Wait()

		return nil
	}, 1*time.Second, 5*time.Second, 10*time.Second)
}

func (s *Service) ProcessLog(wg *sync.WaitGroup, transferLogs []types.Log) {
	defer wg.Done()
	for _, transferLog := range transferLogs {
		var erc20Filterer, _ = abi.NewIERC20Filterer(common.Address{}, nil)
		transfer, err := erc20Filterer.ParseTransfer(transferLog)
		if err != nil {
			continue
		}
		//log.Println(transfer.Raw.Address.String())
		var erc20Caller, _ = abi.NewIERC20Caller(transfer.Raw.Address, s.client)
		balance, err := erc20Caller.BalanceOf(&bind.CallOpts{}, transfer.From)
		if err == nil {
			_, err := s.FoundationDB.Transact(func(transaction fdb.Transaction) (interface{}, error) {
				if balance.Cmp(big.NewInt(0)) == 0 {
					transaction.Clear(fdb.Key(fdbutils.EthBalanceKey(transfer.From, transfer.Raw.Address)))
				} else {
					transaction.Set(fdb.Key(fdbutils.EthBalanceKey(transfer.From, transfer.Raw.Address)), balance.Bytes())
				}
				return nil, nil
			})
			if err != nil {
				s.log.Fatal(err)
				return
			}
		}
		//log.Println(balance, transfer.From.String())
		balance, err = erc20Caller.BalanceOf(&bind.CallOpts{}, transfer.To)
		if err == nil {
			_, err := s.FoundationDB.Transact(func(transaction fdb.Transaction) (interface{}, error) {
				if balance.Cmp(big.NewInt(0)) == 0 {
					transaction.Clear(fdb.Key(fdbutils.EthBalanceKey(transfer.To, transfer.Raw.Address)))
				} else {
					transaction.Set(fdb.Key(fdbutils.EthBalanceKey(transfer.To, transfer.Raw.Address)), balance.Bytes())
				}
				return nil, nil
			})
			if err != nil {
				s.log.Fatal(err)
				return
			}
		}
	}
}

func splitToBatches(batchSize int, transferLogs []types.Log) [][]types.Log {
	inserted := 0
	res := make([][]types.Log, 0)

	for inserted < len(transferLogs) {
		if inserted+batchSize < len(transferLogs) {
			res = append(res, transferLogs[inserted:inserted+batchSize])
		} else {
			if inserted == len(transferLogs) {
				break
			}
			res = append(res, transferLogs[inserted:])
		}
		inserted += batchSize
	}

	return res
}
