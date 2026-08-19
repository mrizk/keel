package telemetry_test

import (
	"testing"

	"github.com/foomo/keel/telemetry"
	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/otel/metric"
)

// TestNewMetricInstruments verifies the typed constructors build non-nil
// instruments and accept the shared/typed options that compile for each kind.
func TestNewMetricInstruments(t *testing.T) {
	t.Parallel()

	// Shared InstrumentOption (WithDescription/WithUnit) applies to every kind;
	// WithExplicitBucketBoundaries is histogram-only. Wrong-kind options would
	// fail to compile — that is the point of the typed signatures.
	assert.NotNil(t, telemetry.NewIntHistogram("h.int", metric.WithUnit("ms"), metric.WithExplicitBucketBoundaries(1, 2, 5)))
	assert.NotNil(t, telemetry.NewFloatHistogram("h.float", metric.WithExplicitBucketBoundaries(0.1, 0.5)))
	assert.NotNil(t, telemetry.NewIntGauge("g.int", metric.WithDescription("d")))
	assert.NotNil(t, telemetry.NewFloatGauge("g.float"))
	assert.NotNil(t, telemetry.NewIntCounter("c.int", metric.WithUnit("{req}")))
	assert.NotNil(t, telemetry.NewFloatCounter("c.float"))
	assert.NotNil(t, telemetry.NewIntUpDownCounter("ud.int"))
	assert.NotNil(t, telemetry.NewFloatUpDownCounter("ud.float"))
}
