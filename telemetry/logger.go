package telemetry

import (
	"context"
	"path"

	goruntime "github.com/foomo/go/runtime"
	"github.com/foomo/keel/log"
	foomosemconv "github.com/foomo/opentelemetry-go/semconv"
	"go.opentelemetry.io/otel/attribute"
	semconv "go.opentelemetry.io/otel/semconv/v1.41.0"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func LogWarn(ctx context.Context, msg string, kv ...attribute.KeyValue) {
	Log(ctx, zapcore.WarnLevel, msg, 1, kv...)
}

func LogError(ctx context.Context, msg string, kv ...attribute.KeyValue) {
	Log(ctx, zapcore.ErrorLevel, msg, 1, kv...)
}

func LogDebug(ctx context.Context, msg string, kv ...attribute.KeyValue) {
	Log(ctx, zapcore.DebugLevel, msg, 1, kv...)
}

func LogInfo(ctx context.Context, msg string, kv ...attribute.KeyValue) {
	Log(ctx, zapcore.InfoLevel, msg, 1, kv...)
}

func Log(ctx context.Context, lvl zapcore.Level, msg string, skip int, kv ...attribute.KeyValue) {
	if !zap.L().Core().Enabled(lvl) {
		return
	}

	attrs := make([]attribute.KeyValue, 0, len(kv)+5)
	attrs = append(attrs, kv...)

	if spanCtx := trace.SpanContextFromContext(ctx); spanCtx.IsValid() {
		attrs = append(attrs,
			foomosemconv.TraceID(spanCtx.TraceID().String()),
			foomosemconv.SpanID(spanCtx.SpanID().String()),
		)
	}

	if fr := goruntime.CallFrame(skip + 1); !fr.Zero() {
		attrs = append(attrs,
			semconv.CodeFunctionName(fr.Short()),
			semconv.CodeFilePath(path.Join(path.Base(path.Dir(fr.File)), path.Base(fr.File))),
			semconv.CodeLineNumber(fr.Line),
		)
	}

	zap.L().WithOptions(zap.WithCaller(false)).Log(lvl, msg, log.Attributes(attrs...)...)
}
