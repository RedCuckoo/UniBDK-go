package api

import (
	"context"
	"github.com/deliveroo/jsonrpc-go"
	"gitlab.com/distributed_lab/logan/v3"
	"time"
)

func LoggingMiddleware(logger *logan.Entry) jsonrpc.Middleware {
	return func(next jsonrpc.Next) jsonrpc.Next {
		return func(ctx context.Context, params interface{}) (interface{}, error) {
			method := jsonrpc.MethodFromContext(ctx)
			start := time.Now()
			logger = logger.WithFields(logan.F{
				"method": method,
			})

			defer func() {
				dur := time.Since(start)
				logger = logger.WithFields(logan.F{
					"duration": dur,
				})

				logger.Info("request finished")
			}()
			logger.Info("request started")
			return next(ctx, params)
		}
	}
}

func CtxMiddleware(extenders ...func(context.Context) context.Context) jsonrpc.Middleware {
	return func(next jsonrpc.Next) jsonrpc.Next {
		return func(ctx context.Context, params interface{}) (interface{}, error) {
			for _, extender := range extenders {
				ctx = extender(ctx)
			}

			return next(ctx, params)
		}
	}
}
