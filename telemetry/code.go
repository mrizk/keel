package telemetry

import (
	goruntime "github.com/foomo/go/runtime"
	"go.opentelemetry.io/otel/attribute"
	semconv "go.opentelemetry.io/otel/semconv/v1.41.0"
)

func CodeCaller(skip int) []attribute.KeyValue {
	if fr := goruntime.CallFrame(skip + 1); !fr.Zero() {
		return []attribute.KeyValue{
			semconv.CodeFunctionName(fr.Name()),
			semconv.CodeFilePath(fr.File),
			semconv.CodeLineNumber(fr.Line),
		}
	}

	return nil
}

func CodeStacktrace(num, skip int) attribute.KeyValue {
	return semconv.CodeStacktrace(goruntime.StackTrace(num, skip+1))
}
