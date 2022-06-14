package requests
//
//import (
//	"github.com/deliveroo/jsonrpc-go"
//	"gitlab.com/tokend/psims/stellar/resources"
//)
//
//// GetCurrencyParam returns baseAsset, error
//func GetCurrencyParam(params resources.Params) (string, error) {
//	if len(params) != 1 {
//		return "", jsonrpc.InvalidParams("method expects one params")
//	}
//
//	assetString, ok := params[0].(string)
//
//	if !ok {
//		return "", jsonrpc.InvalidParams("asset string should be string")
//	}
//
//	return assetString, nil
////}
