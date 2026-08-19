package telemetry

import (
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/metric/noop"
)

// NewIntHistogram creates and returns a Int64Histogram metric instrument with the specified name and options.
func NewIntHistogram(name string, opts ...metric.Int64HistogramOption) metric.Int64Histogram {
	m, err := Meter().Int64Histogram(name, opts...)
	if err != nil {
		otel.Handle(err)
		return noop.Int64Histogram{}
	}

	return m
}

// NewFloatHistogram creates and returns a Float64Histogram metric instrument with the specified name and options.
func NewFloatHistogram(name string, opts ...metric.Float64HistogramOption) metric.Float64Histogram {
	m, err := Meter().Float64Histogram(name, opts...)
	if err != nil {
		otel.Handle(err)
		return noop.Float64Histogram{}
	}

	return m
}

// NewIntGauge creates and returns a Int64Gauge metric instrument with the specified name and options.
func NewIntGauge(name string, opts ...metric.Int64GaugeOption) metric.Int64Gauge {
	m, err := Meter().Int64Gauge(name, opts...)
	if err != nil {
		otel.Handle(err)
		return noop.Int64Gauge{}
	}

	return m
}

// NewFloatGauge creates and returns a Float64Gauge metric instrument with the specified name and options.
func NewFloatGauge(name string, opts ...metric.Float64GaugeOption) metric.Float64Gauge {
	m, err := Meter().Float64Gauge(name, opts...)
	if err != nil {
		otel.Handle(err)
		return noop.Float64Gauge{}
	}

	return m
}

// NewIntCounter creates and returns a Int64Counter metric instrument with the specified name and options.
func NewIntCounter(name string, opts ...metric.Int64CounterOption) metric.Int64Counter {
	m, err := Meter().Int64Counter(name, opts...)
	if err != nil {
		otel.Handle(err)
		return noop.Int64Counter{}
	}

	return m
}

// NewFloatCounter creates and returns a Float64Counter metric instrument with the specified name and options.
func NewFloatCounter(name string, opts ...metric.Float64CounterOption) metric.Float64Counter {
	m, err := Meter().Float64Counter(name, opts...)
	if err != nil {
		otel.Handle(err)
		return noop.Float64Counter{}
	}

	return m
}

// NewIntUpDownCounter creates and returns a Int64UpDownCounter metric instrument with the specified name and options.
func NewIntUpDownCounter(name string, opts ...metric.Int64UpDownCounterOption) metric.Int64UpDownCounter {
	m, err := Meter().Int64UpDownCounter(name, opts...)
	if err != nil {
		otel.Handle(err)
		return noop.Int64UpDownCounter{}
	}

	return m
}

// NewFloatUpDownCounter creates and returns a Float64UpDownCounter metric instrument with the specified name and options.
func NewFloatUpDownCounter(name string, opts ...metric.Float64UpDownCounterOption) metric.Float64UpDownCounter {
	m, err := Meter().Float64UpDownCounter(name, opts...)
	if err != nil {
		otel.Handle(err)
		return noop.Float64UpDownCounter{}
	}

	return m
}
