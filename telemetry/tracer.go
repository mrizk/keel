package telemetry

import (
	"go.opentelemetry.io/otel/trace"
)

// Tracer returns a trace.Tracer instance configured with optional TracerOptions.
// The instrumentation scope version defaults to the keel module version; pass a
// trace.WithInstrumentationVersion option to override it.
func Tracer(opts ...trace.TracerOption) trace.Tracer {
	return TracerProvider().Tracer(Name, append([]trace.TracerOption{trace.WithInstrumentationVersion(Version())}, opts...)...)
}
