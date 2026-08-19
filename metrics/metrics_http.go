package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

// Deprecated: use otelhttp
func NewHTTPRequestSizeSummaryVec(namespace string) *prometheus.SummaryVec {
	return NewRequestSizeSummaryVec(namespace, "http", []string{"method", "code"})
}

// Deprecated: use otelhttp
func NewHTTPResponseSizeSummaryVec(namespace string) *prometheus.SummaryVec {
	return NewResponseSizeSummaryVec(namespace, "http", []string{"method", "code"})
}

// Deprecated: use otelhttp
func NewHTTPRequestsCounterVec(namespace string) *prometheus.CounterVec {
	return NewRequestsCounterVec(namespace, "http", []string{"method", "code"})
}

// Deprecated: use otelhttp
func NewHTTPRequestDurationHistogram(namespace string) prometheus.Histogram {
	return NewRequestDurationHistogram(namespace, "http")
}
