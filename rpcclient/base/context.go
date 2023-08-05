package base

import (
	"context"
	"time"
)

type ContextWithTracing interface {
	WithTracing(ctx context.Context) context.Context
}

const (
	// ContextKey 上下文变量存储键名，前后端系统共享
	ContextKey = "ContextKey"
	// TrackKey 上下文传递
	TrackKey = "TraceId"
	// TraceingEnabled 当前请求是否开启
	TraceingEnabled = "TraceingEnabled"
)

func ContextWithTimeoutAndTracing(ctx context.Context, t time.Duration) (context.Context, context.CancelFunc) {
	ctxTimeout, cancel := context.WithTimeout(ctx, t)
	value := ctx.Value(ContextKey)
	if value == nil {
		return ctxTimeout, cancel
	}

	localCtx, ok := value.(ContextWithTracing)
	if !ok {
		return ctxTimeout, cancel
	}

	ctx = localCtx.WithTracing(ctxTimeout)
	return ctx, cancel
}
