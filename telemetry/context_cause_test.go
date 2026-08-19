package telemetry_test

import (
	"testing"
	"time"

	"github.com/foomo/keel/telemetry"
	pkgerrors "github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/otel/codes"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/sdk/trace/tracetest"
)

type selfCauser struct{ msg string }

func (e *selfCauser) Error() string { return e.msg }
func (e *selfCauser) Cause() error  { return e }

type cycleCauser struct {
	msg  string
	next error
}

func (e *cycleCauser) Error() string { return e.msg }
func (e *cycleCauser) Cause() error  { return e.next }

type nilCauser struct{ msg string }

func (e *nilCauser) Error() string { return e.msg }
func (e *nilCauser) Cause() error  { return nil }

func mustTerminate(t *testing.T, name string, fn func()) {
	t.Helper()

	done := make(chan struct{})

	go func() {
		defer close(done)
		fn()
	}()

	select {
	case <-done:
	case <-time.After(5 * time.Second):
		t.Fatalf("%s did not terminate within 5s: a cyclic Cause() chain is spinning", name)
	}
}

func TestRootCause(t *testing.T) {
	t.Parallel()

	t.Run("self referential chain terminates", func(t *testing.T) {
		t.Parallel()

		err := &selfCauser{msg: "rpc failed"}

		var got error

		mustTerminate(t, "RootCause", func() { got = telemetry.RootCause(err) })

		require.NotNil(t, got)
		assert.Equal(t, "rpc failed", got.Error())
	})

	t.Run("two node cycle terminates", func(t *testing.T) {
		t.Parallel()

		a := &cycleCauser{msg: "a"}
		b := &cycleCauser{msg: "b"}
		a.next, b.next = b, a

		mustTerminate(t, "RootCause", func() { _ = telemetry.RootCause(a) })
	})

	t.Run("unwraps a regular chain", func(t *testing.T) {
		t.Parallel()

		root := pkgerrors.New("root")
		wrapped := pkgerrors.Wrap(pkgerrors.Wrap(root, "middle"), "outer")

		assert.Equal(t, root, telemetry.RootCause(wrapped))
	})

	t.Run("never returns nil for a non nil error", func(t *testing.T) {
		t.Parallel()

		err := &nilCauser{msg: "no cause"}

		// pkg/errors.Cause returns nil here, so callers doing
		// Cause(err).Error() panic. RootCause returns the last non-nil error.
		require.Nil(t, pkgerrors.Cause(err)) //nolint:testifylint

		got := telemetry.RootCause(err)
		require.NotNil(t, got)
		assert.Equal(t, "no cause", got.Error())
	})
}

func TestEndSpanWithSelfReferentialCause(t *testing.T) {
	t.Parallel()

	spanRecorder := tracetest.NewSpanRecorder()
	tracerProvider := sdktrace.NewTracerProvider(sdktrace.WithSpanProcessor(spanRecorder))

	ctx, _ := tracerProvider.Tracer("test").Start(t.Context(), "test")

	mustTerminate(t, "EndSpan", func() {
		telemetry.Ctx(ctx).EndSpan(&selfCauser{msg: "rpc failed"})
	})

	spans := spanRecorder.Ended()
	require.Len(t, spans, 1)
	assert.Equal(t, codes.Error, spans[0].Status().Code)
	assert.Equal(t, "rpc failed", spans[0].Status().Description)
}
