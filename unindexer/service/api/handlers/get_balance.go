package handlers
//
//import (
//	"context"
//	"github.com/deliveroo/jsonrpc-go"
//	"github.com/stellar/go/clients/horizonclient"
//	"github.com/stellar/go/protocols/horizon/base"
//	"github.com/stellar/go/txnbuild"
//	"gitlab.com/tokend/psims/stellar/internal/service/api/helpers"
//	"gitlab.com/tokend/psims/stellar/internal/service/api/requests"
//	"gitlab.com/tokend/psims/stellar/resources"
//)
//
//func GetBalance(ctx context.Context, params resources.Params) (interface{}, error) {
//	assetString, err := requests.GetCurrencyParam(params)
//
//	if err != nil {
//		helpers.LogFromCtx(ctx).WithError(err).Error("Failed to get asset from request")
//		return "", jsonrpc.Error("-1", "Failed to get asset parameters from request")
//	}
//
//	horizonClient := helpers.HorizonFromCtx(ctx)
//	info, err := horizonClient.AccountDetail(horizonclient.AccountRequest{
//		AccountID: helpers.AccountFromCtx(ctx).StockExchangeAccount,
//	})
//
//	if err != nil {
//		helpers.LogFromCtx(ctx).WithError(err).Error("Failed to get info")
//		return "", jsonrpc.Error("-1", "Failed to get info from blockchain")
//	}
//
//	asset, err := GetAssetFromString(assetString)
//
//	if err != nil {
//		helpers.LogFromCtx(ctx).WithError(err).Error("Failed to parse asset")
//		return "", jsonrpc.Error("-1", "Failed to parse asset")
//	}
//
//	var result string
//
//	for _, value := range info.Balances {
//		if assetsEqual(value.Asset, asset) {
//			result = value.Balance
//			break
//		}
//	}
//
//	if result == "" {
//		helpers.LogFromCtx(ctx).WithError(err).Error("No balance found for such currency")
//		return nil, jsonrpc.Error("-1", "No balance found for such currency")
//	}
//
//	return result, nil
//}
//
//func assetsEqual(asset1 base.Asset, asset2 base.Asset) bool {
//	return asset1.Code == asset2.Code && asset1.Type == asset2.Type && asset1.Issuer == asset2.Issuer
//}
//
//func GetAssetFromString(assetString string) (base.Asset, error) {
//	assetTxnBuild, err := txnbuild.ParseAssetString(assetString)
//
//	if err != nil {
//		return base.Asset{}, jsonrpc.Error("-1", "Failed to parse asset from string")
//	}
//
//	// from http://github.com/stellar/go/xdr/xdr_generated.go
//	//
//	// AssetType is an XDR Enum defines as:
//	//
//	//   enum AssetType
//	//    {
//	//        ASSET_TYPE_NATIVE = 0,
//	//        ASSET_TYPE_CREDIT_ALPHANUM4 = 1,
//	//        ASSET_TYPE_CREDIT_ALPHANUM12 = 2
//	//    };
//	//
//
//	assetType, err := assetTxnBuild.GetType()
//
//	if err != nil {
//		return base.Asset{}, jsonrpc.Error("-1", "Failed to get asset type")
//	}
//
//	var assetTypeString string
//
//	switch assetType {
//	case 0:
//		assetTypeString = "native"
//		break
//	case 1:
//		assetTypeString = "credit_alphanum4"
//		break
//	case 2:
//		assetTypeString = "credit_alphanum12"
//		break
//	}
//
//	return base.Asset{
//		Type:   assetTypeString,
//		Code:   assetTxnBuild.GetCode(),
//		Issuer: assetTxnBuild.GetIssuer(),
//	}, nil
//}
