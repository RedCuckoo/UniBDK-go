package btc_indexer

import (
	"encoding/json"
	"github.com/RedCuckoo/UniBDK-go/unindexer/fdbutils"
)

type GetBlockHashResponse struct {
	Result string `json:"result"`
}

type GetRawTransactionResponse struct {
	Result fdbutils.Transaction `json:"result"`
}

type GetBlockResponse struct {
	Result fdbutils.Block `json:"result"`
}

type jsonrpcMessage struct {
	Version string          `json:"jsonrpc,omitempty"`
	ID      json.RawMessage `json:"id,omitempty"`
	Method  string          `json:"method,omitempty"`
	Params  json.RawMessage `json:"params,omitempty"`
	Error   *jsonError      `json:"error,omitempty"`
	Result  json.RawMessage `json:"result,omitempty"`
}
type jsonError struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

//func (tx *Transaction) MarshalToFDB() []byte {
//	return tuple.Tuple{
//		int(proto.NamespaceFdb_BtcTransaction),
//		tx.Hash,
//		tx.Time,
//		tx.Vin,
//	}.Pack()
//}
//
//func (tin *TransactionVIn) MarshalToFDB() []byte {
//	return tuple.Tuple{
//		tin.TxId,
//		tin.Vout,
//		tin.ScriptSig,
//		tin.Sequence,
//		tin.TxInWitness,
//	}.Pack()
//}
//
//func (tin *TransactionVIn) MarshalToFDB() []byte {
//	return tuple.Tuple{
//		tin.TxId,
//		tin.Vout,
//		tin.ScriptSig,
//		tin.Sequence,
//		tin.TxInWitness,
//	}.Pack()
//}
