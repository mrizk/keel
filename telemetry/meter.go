package telemetry

import (
	"go.opentelemetry.io/otel/metric"
)

// Meter returns a metric.Meter instance with the provided options, using the default MeterProvider and the defined Name.
// The instrumentation scope version defaults to the keel module version; pass a
// metric.WithInstrumentationVersion option to override it.
func Meter(opts ...metric.MeterOption) metric.Meter {
	return MeterProvider().Meter(Name, append([]metric.MeterOption{metric.WithInstrumentationVersion(Version())}, opts...)...)
}
