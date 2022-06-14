package fdbutils

import (
	"github.com/RedCuckoo/UniBDK-go/unindexer/config"
	"github.com/apple/foundationdb/bindings/go/src/fdb"
	"github.com/pkg/errors"
	"os"
	"strings"
	"time"
)

type FDB struct {
	fdb.Database
}

func Connect(config config.Config) (*FDB, error) {
	if err := fdb.APIVersion(630); err != nil {
		return nil, errors.Wrapf(err, "api version of fdb client must be 630")
	}
	fdbConfig := config.FDB()
	if err := checkClusterFile(fdbConfig.ClusterFile); err != nil {
		return nil, errors.Wrapf(err, "failed to check cluster file")
	}
	database, err := fdb.OpenDatabase(fdbConfig.ClusterFile)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to open database")
	}
	db := &FDB{database}
	if err = db.healthCheck(time.Second * 5); err != nil {
		return nil, errors.Wrapf(err, "health check failed")
	}

	return db, nil
}

func checkClusterFile(clusterFile string) error {
	url, err := os.ReadFile(clusterFile)
	if err != nil {
		return errors.Wrapf(err, "failed to read cluster file")
	}
	splitUrl := strings.Split(string(url), "@")
	if len(splitUrl) == 0 {
		return errors.Errorf("foundation db url does not contain @")
	}
	return nil
}

func (db *FDB) healthCheck(timeout time.Duration) error {
	errChan := make(chan error)
	go func() {
		keys := []fdb.KeyConvertible{fdb.Key([]byte{0})}

		keyRanges := make([]fdb.KeyRange, 0)
		for _, key := range keys {
			keyRange, err := fdb.PrefixRange(key.FDBKey())
			if err != nil {
				continue
			}
			keyRanges = append(keyRanges, keyRange)
		}

		_, err := db.ReadTransact(func(transaction fdb.ReadTransaction) (interface{}, error) {
			rangeResults := make([]fdb.RangeResult, 0)
			for _, keyRange := range keyRanges {
				rangeResults = append(rangeResults, transaction.Snapshot().GetRange(keyRange, fdb.RangeOptions{Limit: 1, Reverse: false}))
			}
			results := make([]fdb.KeyValue, 0)
			for _, rangeResult := range rangeResults {
				slice, err := rangeResult.GetSliceWithError()
				if err != nil {
					return nil, err
				}
				results = append(results, slice...)
			}
			return results, nil
		})

		errChan <- err
	}()
	select {
	case err := <-errChan:
		return err
	case <-time.After(timeout):
		return errors.Errorf("timeout")
	}
}
