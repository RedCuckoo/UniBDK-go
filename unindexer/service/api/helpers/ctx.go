package helpers

import (
	"context"
	"gitlab.com/distributed_lab/logan/v3"
)

type ctxKey int

const (
	logCtxKey ctxKey = iota
	depositDBCtxKey
	withdrawDBCtxKey
	horizonCtxKey
	accountCtxKey
	withdrawCommandCtxKey
)

func CtxLog(entry *logan.Entry) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, logCtxKey, entry)
	}
}

//func CtxDepositsDB(depositQ data.DepositQ) func(context.Context) context.Context {
//	return func(ctx context.Context) context.Context {
//		return context.WithValue(ctx, depositDBCtxKey, depositQ)
//	}
//}

//func CtxAccount(account config.Account) func(context.Context) context.Context {
//	return func(ctx context.Context) context.Context {
//		return context.WithValue(ctx, accountCtxKey, account)
//	}
//}

func LogFromCtx(ctx context.Context) *logan.Entry {
	return ctx.Value(logCtxKey).(*logan.Entry)
}

//func DepositsDBFromCtx(ctx context.Context) data.DepositQ {
//	return ctx.Value(depositDBCtxKey).(data.DepositQ)
//}
//
//func AccountFromCtx(ctx context.Context) config.Account {
//	return ctx.Value(accountCtxKey).(config.Account)
//}
