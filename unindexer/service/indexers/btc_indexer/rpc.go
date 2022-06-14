package btc_indexer

import (
	"bytes"
	"encoding/json"
	"github.com/RedCuckoo/UniBDK-go/unindexer/fdbutils"
	"github.com/apple/foundationdb/bindings/go/src/fdb"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/pkg/errors"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
)

func (s *Service) GetBlockHash(blockNumber int64) (GetBlockHashResponse, error) {
	var result GetBlockHashResponse
	response, err := s.Request(&rpc.BatchElem{
		Method: "getblockhash",
		Args:   []interface{}{blockNumber},
	})
	if err != nil {
		return GetBlockHashResponse{}, err
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return GetBlockHashResponse{}, err
	}

	return result, nil
}

func (s *Service) GetBlockByHash(blockHash string) (GetBlockResponse, error) {
	var result GetBlockResponse
	response, err := s.Request(&rpc.BatchElem{
		Method: "getblock",
		Args:   []interface{}{blockHash, 1},
	})
	if err != nil {
		return GetBlockResponse{}, err
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return GetBlockResponse{}, err
	}

	return result, nil
}

func (s *Service) GetRawTransaction(txId string) (GetRawTransactionResponse, error) {
	var response []byte
	var err error
	var result GetRawTransactionResponse

	transactionResponse, err := s.FoundationDB.ReadTransact(func(transaction fdb.ReadTransaction) (interface{}, error) {
		return transaction.Get(fdb.Key(fdbutils.BtcTransactionKey(txId))).Get()
	})
	if err != nil {
		return GetRawTransactionResponse{}, err
	}

	if transactionResponse.([]byte) != nil {
		response = transactionResponse.([]byte)
	} else {
		response, err = s.Request(&rpc.BatchElem{
			Method: "getrawtransaction",
			Args:   []interface{}{txId, 2},
		})
		if err != nil {
			return GetRawTransactionResponse{}, err
		}
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return GetRawTransactionResponse{}, err
	}

	return result, nil
}

func (s *Service) Request(requestElem *rpc.BatchElem) ([]byte, error) {
	startId := uint64(rand.Int31())
	var body []byte
	var err error
	id := strconv.AppendUint(nil, startId, 10)
	msg := &jsonrpcMessage{Version: "2.0", ID: id, Method: requestElem.Method}
	if requestElem.Args != nil {
		var err error
		msg.Params, err = json.Marshal(requestElem.Args)
		if err != nil {
			return nil, nil
		}
	}

	body, err = json.Marshal(msg)
	if err != nil {
		return nil, nil
	}
	request, err := http.NewRequest("POST", s.requestUrl, ioutil.NopCloser(bytes.NewReader(body)))
	if err != nil {
		return nil, errors.Wrapf(err, "failed to build request")
	}

	headers := make(http.Header, 2)
	headers.Set("accept", "application/json")
	headers.Set("content-type", "application/json")
	headers.Set("x-api-key", s.requestApiKey)
	request.Header = headers
	response, err := s.httpClient.RoundTrip(request)

	if err != nil {
		return nil, errors.Wrapf(err, "failed to round trip")
	}
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return responseBody, errors.Wrapf(err, "failed to ioutil.ReadAll(response), status: %s", response.Status)
	}
	if response.StatusCode != 200 {
		return nil, errors.Errorf("http request returned not OK %s %s", response.Status, string(responseBody))
	}

	var responseMessage *jsonrpcMessage
	if err := json.Unmarshal(responseBody, &responseMessage); err != nil {
		var tempResp jsonrpcMessage
		err = json.Unmarshal(responseBody, &tempResp)
		if err != nil {
			return responseBody, errors.Wrapf(err, "failed to unmarshal response \"%s\"", string(responseBody))
		}
		if tempResp.Error != nil {
			return responseBody, errors.Wrapf(err, "failed to process request by node with error \"%s\"", tempResp.Error.Message)
		}
	}

	requestElem.Result = responseMessage.Result

	return responseBody, nil
}
